/*
Tyk Gateway API

Testing HealthCheckingApiService

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

func Test_gate_HealthCheckingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test HealthCheckingApiService Hello", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.HealthCheckingApi.Hello(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test HealthCheckingApiService HelloAPI", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var listenPath string

        resp, httpRes, err := apiClient.HealthCheckingApi.HelloAPI(context.Background(), listenPath).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
