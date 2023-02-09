package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/disintegration/imaging"
	"github.com/szpnygo/goclip/clip"
	"github.com/szpnygo/goclip/logic"
)

func UploadImage(c *logic.LogicContext) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Result(1, nil, "get file error")
		return
	}
	imageFile, err := file.Open()
	if err != nil {
		c.Result(1, nil, "open file error")
		return
	}

	hash := md5.New()
	_, err = io.Copy(hash, imageFile)
	if err != nil {
		c.Result(1, nil, "hash file error")
		return
	}
	imageMD5 := fmt.Sprintf("%x", hash.Sum(nil))
	fileName := fmt.Sprintf("%s%s", imageMD5, filepath.Ext(file.Filename))

	// upload image to s3
	_, _ = imageFile.Seek(0, io.SeekStart)
	_, err = c.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: &c.Bucket,
		Key:    &fileName,
		Body:   imageFile,
	})
	if err != nil {
		c.Result(1, nil, "upload file error to s3 ")
		return
	}

	_, _ = imageFile.Seek(0, io.SeekStart)
	image, err := imaging.Decode(imageFile)
	if err != nil {
		c.Result(1, nil, "decode file error")
		return
	}

	// clip need 224*224 image
	smallImage := imaging.Resize(image, 224, 224, imaging.Lanczos)
	var buf bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	if err := imaging.Encode(encoder, smallImage, imaging.PNG); err != nil {
		c.Result(1, nil, "encode file error")
		return
	}
	if err := encoder.Close(); err != nil {
		c.Result(1, nil, "encode close error")
		return
	}

	// insert clip
	results, err := c.ClipHelper.GetObject("md5", imageMD5)
	if err != nil {
		c.Result(1, nil, "get clip error")
		return
	}
	if len(results) == 0 {
		clipImageData := clip.ImageData{
			Image: buf.String(),
			MD5:   imageMD5,
			UID:   fileName,
		}
		if _, err := c.ClipHelper.CreateObject(clipImageData); err != nil {
			c.Result(1, nil, "create clip error")
			log.Printf("create clip %s error: %v\n", imageMD5, err)
			return
		} else {
			log.Printf("create clip %s success\n", imageMD5)
		}
	}

	log.Printf("upload file %s success\n", fileName)
	c.Result(0, nil, "success")
}

type ListImagesRequest struct {
	Token string `json:"token"`
	Size  int64  `json:"size"`
}

type ListImagesResponse struct {
	Images []string `json:"images"`
	Token  string   `json:"token"`
}

func ListImage(c *logic.LogicContext) {
	var request ListImagesRequest
	_ = c.BindJSON(&request)
	if request.Size == 0 {
		request.Size = 50
	}
	result, err := c.S3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:            &c.Bucket,
		MaxKeys:           aws.Int64(request.Size),
		ContinuationToken: aws.String(request.Token),
	})
	if err != nil {
		c.Result(1, nil, "list image error "+err.Error())
		return
	}

	response := ListImagesResponse{
		Images: []string{},
		Token:  request.Token,
	}
	for _, item := range result.Contents {
		response.Images = append(response.Images, *item.Key)
	}
	if *result.IsTruncated {
		response.Token = *result.NextContinuationToken
	}

	c.Result(0, response, "success")
}
