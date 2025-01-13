package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

var (
	BaseUrl           = "http://tyk-gateway.localhost:8080"
	xTykAuthorization = "x-tyk-authorization"
)

func main() {
	apiConfig := apim.Configuration{
		DefaultHeader: map[string]string{
			xTykAuthorization: "<YOUR GATEWAY AUTHORIZATION KEY HERE>",
		},
		Debug: false,
		Servers: apim.ServerConfigurations{
			{
				URL: BaseUrl,
			},
		},
		HTTPClient: nil,
	}
	client := apim.NewAPIClient(&apiConfig)
	ctx := context.Background()
	apis, resp, err := client.APIsAPI.ListApis(ctx).Execute()
	if err != nil || resp.StatusCode != 200 {
		body, err := ReadResponseBody(resp)
		if err != nil {
			log.Println(resp.Status)
			return
		}
		log.Println(body)
		return
	}
	for _, item := range apis {
		fmt.Printf("%+v\n", item.GetApiId())
	}
}

func ReadResponseBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert the body to a string and return
	return string(body), nil
}
