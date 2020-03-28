package services

import (
	"context"
	"errors"
)

type ElasticSearchMock struct{}

type ElasticSearchFailMock struct{}

func NewElasticSearchConnectionFails() *ElasticSearchFailMock {
	return &ElasticSearchFailMock{}
}

func NewElasticSearchMock() *ElasticSearchMock {
	return &ElasticSearchMock{}
}

func (es *ElasticSearchMock) SearchClient(ctx context.Context, key, value string) ([]Hotel, error) {
	return []Hotel{
		{
			Name:    "joshua",
			Address: "23 idris",
			Stars:   "123",
			Contact: "9084744",
			Phone:   "+098477333",
			URI:     "www.google.com",
		},
	}, nil
}

func (esf *ElasticSearchFailMock) SearchClient(ctx context.Context, key, value string) ([]Hotel, error) {
	return nil, errors.New("search result fails with error")
}
