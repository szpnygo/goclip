package controller

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/szpnygo/goclip/clip"
	"github.com/szpnygo/goclip/logic"
)

func Search(c *logic.LogicContext) {
	results, _ := c.ClipHelper.SearchText(c.Query("search"), clip.WithPageSize(50), clip.WithPageNum(1))
	list := make([]string, len(results))
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			item := results[index]
			if _, err := c.S3Client.HeadObject(&s3.HeadObjectInput{
				Bucket: &c.Bucket,
				Key:    &item.UID,
			}); err == nil {
				list[index] = item.UID
			}
		}(i)
	}
	wg.Wait()

	result := []string{}
	for i := range list {
		if len(list[i]) != 0 {
			result = append(result, list[i])
		}
	}

	c.Result(0, result, "")
}
