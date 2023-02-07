/*
Tyk Gateway API

Testing PoliciesApiService

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

func Test_gate_PoliciesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PoliciesApiService AddPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PoliciesApi.AddPolicy(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PoliciesApiService DeletePolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polID string

        resp, httpRes, err := apiClient.PoliciesApi.DeletePolicy(context.Background(), polID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PoliciesApiService GetPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polID string

        resp, httpRes, err := apiClient.PoliciesApi.GetPolicy(context.Background(), polID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PoliciesApiService ListPolicies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PoliciesApi.ListPolicies(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PoliciesApiService UpdatePolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polID string

        resp, httpRes, err := apiClient.PoliciesApi.UpdatePolicy(context.Background(), polID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
