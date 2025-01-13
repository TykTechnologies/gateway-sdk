package main

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		// Do something here.
		log.Println(errors.New(resp.Status))

		return
	}
	for _, item := range apis {
		fmt.Printf("%+v\n", item.GetApiId())
	}
}
