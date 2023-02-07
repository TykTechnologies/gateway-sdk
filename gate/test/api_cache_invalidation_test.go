/*
Tyk Gateway API

Testing CacheInvalidationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gate

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_gate_CacheInvalidationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CacheInvalidationApiService InvalidateCache", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var apiID string

        resp, httpRes, err := apiClient.CacheInvalidationApi.InvalidateCache(context.Background(), apiID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
