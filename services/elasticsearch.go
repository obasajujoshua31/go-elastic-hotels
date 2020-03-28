package services

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

type ESSearch interface {
	SearchClient(ctx context.Context, key, value string) ([]Hotel, error)
}

type ESClient struct {
	*elastic.Client
}

func ConnectToESServer() (*ESClient, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	if err != nil {
		return nil, err
	}

	return &ESClient{client}, nil
}

func SearchForResult(ctx context.Context, esSearch ESSearch, key, value string) (hotels []Hotel, err error) {
	return esSearch.SearchClient(ctx, key, value)
}

func (es *ESClient) SearchClient(ctx context.Context, key, value string) (hotels []Hotel, err error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	searchService := es.Search().Index("hotels").SearchSource(searchSource)
	searchResult, err := searchService.Do(ctx)

	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		hotel := Hotel{}
		if err := json.Unmarshal(hit.Source, &hotel); err != nil {
			return nil, err
		}

		hotels = append(hotels, hotel)
	}
	return hotels, nil
}
