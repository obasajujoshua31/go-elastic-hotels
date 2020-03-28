package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SearchForResultSucceeds(t *testing.T) {
	esclient := NewElasticSearchMock()

	hotels, err := SearchForResult(context.TODO(), esclient, "any", "value")
	assert.NoError(t, err)
	assert.NotNil(t, hotels)
	assert.Equal(t, len(hotels), 1)
}

func Test_SearchForResultFails(t *testing.T) {
	esclient := NewElasticSearchConnectionFails()

	hotels, err := SearchForResult(context.TODO(), esclient, "any", "value")
	assert.Error(t, err)
	assert.Nil(t, hotels)
}
