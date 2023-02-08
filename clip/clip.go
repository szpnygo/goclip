package clip

import (
	"context"
	"errors"

	"github.com/szpnygo/goclip/weaviatex"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

type ImageData struct {
	Image string
	MD5   string
	UID   string
}

type SearchResult struct {
	ID       string
	MD5      string
	UID      string
	Distance float64
	Vector   []float32
}

type ClipHelper struct {
	*weaviatex.WeaviateClient
	ClassName string
}

func NewClipHelper(client *weaviatex.WeaviateClient) *ClipHelper {
	return &ClipHelper{
		WeaviateClient: client,
	}
}

// Init the clip class in weaviate
func (c *ClipHelper) Init(name string) error {
	schema, err := c.Schema().
		ClassGetter().
		WithClassName(name).
		Do(context.Background())

	if err == nil && schema.Class == name {
		c.ClassName = name
		return nil
	}

	newClass := c.createClipClass(name)

	err = c.Schema().
		ClassCreator().
		WithClass(newClass).
		Do(context.Background())
	c.ClassName = name

	return err
}

// create a new class for the clip in weaviate
func (c *ClipHelper) createClipClass(name string) *models.Class {
	class := &models.Class{
		Class:           name,
		VectorIndexType: "hnsw",
		Vectorizer:      "multi2vec-clip",
		ModuleConfig: map[string]interface{}{
			"multi2vec-clip": map[string]interface{}{
				"imageFields": []string{"image"},
				"textFields":  []string{"name"},
				"weight": map[string][]float32{
					"textFields":  {0.7},
					"imageFields": {0.3},
				},
			},
		},
		Properties: []*models.Property{
			{
				DataType: []string{"string"},
				Name:     "name",
			},
			{
				DataType: []string{"blob"},
				Name:     "image",
			},
			{
				DataType: []string{"string"},
				Name:     "md5", // md5 of the image
			},
			{
				DataType: []string{"string"},
				Name:     "uid", // uid of the image, generated by snowflake
			},
		},
	}

	return class
}

// Create a new clip object in weaviate
func (c *ClipHelper) CreateObject(data ImageData) (*SearchResult, error) {
	object := map[string]any{
		"image": data.Image,
		"md5":   data.MD5,
		"uid":   data.UID,
	}

	created, err := c.Data().
		Creator().
		WithClassName(c.ClassName).
		WithProperties(object).
		Do(context.Background())

	if err != nil {
		return nil, weaviatex.GetWeaviateError(err)
	}

	result := &SearchResult{
		ID:     created.Object.ID.String(),
		MD5:    data.MD5,
		UID:    data.UID,
		Vector: created.Object.Vector,
	}

	return result, nil
}

// Get object by some key
func (c *ClipHelper) GetObject(key string, value string) ([]*SearchResult, error) {
	additional := graphql.Field{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "id"},
		},
	}
	fields := []graphql.Field{{Name: "md5"}, {Name: "uid"}, additional}

	filter := filters.Where().WithPath([]string{key}).WithOperator(filters.Equal).WithValueString(value)

	data, err := c.GraphQL().Get().
		WithClassName(c.ClassName).
		WithFields(fields...).
		WithWhere(filter).
		Do(context.Background())

	if err != nil {
		return nil, weaviatex.GetWeaviateError(err)
	}

	return c.parseData(data)
}

// Get object by id
func (c *ClipHelper) GetObjectByID(id string) ([]*SearchResult, error) {
	data, err := c.Data().ObjectsGetter().
		WithClassName(c.ClassName).
		WithID(id).
		WithVector().
		WithConsistencyLevel(replication.ConsistencyLevel.QUORUM).
		Do(context.Background())

	if err != nil {
		return nil, weaviatex.GetWeaviateError(err)
	}

	result := []*SearchResult{}
	for _, v := range data {
		if properties, ok := v.Properties.(map[string]interface{}); ok {
			s := &SearchResult{
				ID:     v.ID.String(),
				Vector: v.Vector,
			}
			if d, ok := properties["md5"].(string); ok {
				s.MD5 = d
			}
			if d, ok := properties["uid"].(string); ok {
				s.UID = d
			}
			result = append(result, s)
		}
	}

	return result, nil
}

// Search for objects by text
func (c *ClipHelper) SearchText(text string, optFunc ...SearchOptionFunc) ([]*SearchResult, error) {
	opts := &searchOptions{}
	for _, v := range optFunc {
		v(opts)
	}
	if opts.pageNum <= 0 {
		opts.pageNum = 1
	}
	if opts.pageSize <= 0 {
		opts.pageSize = 10
	}

	nearText := c.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{text})

	additional := graphql.Field{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "id"},
			{Name: "distance"},
		},
	}
	fields := []graphql.Field{{Name: "md5"}, {Name: "uid"}, additional}

	data, err := c.GraphQL().Get().
		WithClassName(c.ClassName).
		WithFields(fields...).
		WithNearText(nearText).
		WithOffset(int((opts.pageNum - 1) * opts.pageSize)).
		WithLimit(int(opts.pageSize)).
		Do(context.Background())

	if err != nil {
		return nil, weaviatex.GetWeaviateError(err)
	}

	return c.parseData(data)
}

// Search for objects by base64 image
func (c *ClipHelper) SearchImage(image string, optFunc ...SearchOptionFunc) ([]*SearchResult, error) {
	opts := &searchOptions{}
	for _, v := range optFunc {
		v(opts)
	}
	if opts.pageNum <= 0 {
		opts.pageNum = 1
	}
	if opts.pageSize <= 0 {
		opts.pageSize = 10
	}

	nearImage := c.GraphQL().
		NearImageArgBuilder().
		WithImage(image)

	additional := graphql.Field{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "id"},
			{Name: "distance"},
		},
	}
	fields := []graphql.Field{{Name: "md5"}, {Name: "uid"}, additional}

	data, err := c.GraphQL().Get().
		WithClassName(c.ClassName).
		WithFields(fields...).
		WithNearImage(nearImage).
		WithOffset(int((opts.pageNum - 1) * opts.pageSize)).
		WithLimit(int(opts.pageSize)).
		Do(context.Background())

	if err != nil {
		return nil, weaviatex.GetWeaviateError(err)
	}

	return c.parseData(data)
}

func (c *ClipHelper) parseData(response *models.GraphQLResponse) ([]*SearchResult, error) {
	for _, err := range response.Errors {
		return nil, errors.New(err.Message)
	}

	result := []*SearchResult{}

	for _, d := range response.Data["Get"].(map[string]interface{}) {
		for _, v := range d.([]interface{}) {
			if vv, ok := v.(map[string]interface{}); ok {
				searchData := SearchResult{
					MD5: vv["md5"].(string),
					UID: vv["uid"].(string),
				}
				if _additional, ok := vv["_additional"].(map[string]interface{}); ok {
					searchData.ID = _additional["id"].(string)
					if f, ok := _additional["distance"].(float64); ok {
						searchData.Distance = f
					}
				}
				result = append(result, &searchData)
			}
		}
	}

	return result, nil
}
