package weaviatex

import (
	"context"
	"errors"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/fault"
)

type WeaviateClient struct {
	*weaviate.Client
	ClassName string
}

func GetWeaviateError(err error) error {
	if err != nil {
		return errors.New(getErrorMsg(err))
	}

	return err
}

func getErrorMsg(err error) string {
	msg := ""
	if err != nil {
		msg = err.Error()
		weaviateError, ok := err.(*fault.WeaviateClientError)
		if ok {
			msg = fmt.Sprintf("%s %v", msg, getErrorMsg(weaviateError.DerivedFromError))
		}
	}

	return msg
}

func NewWeaviateClient(host string) (*WeaviateClient, error) {
	cfg := weaviate.Config{
		Host:   host,
		Scheme: "http",
	}

	client := weaviate.New(cfg)
	isAlive, err := client.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		return nil, GetWeaviateError(err)
	}
	if !isAlive {
		return nil, errors.New("Weaviate is not alive")
	}

	return &WeaviateClient{
		Client: client,
	}, nil
}
