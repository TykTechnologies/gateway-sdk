package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"log"
)

var (
	BaseUrl           = "http://localhost:8080"
	xTykAuthorization = "x-tyk-authorization"
)

func main() {
	apiConfig := apim.Configuration{
		DefaultHeader: map[string]string{
			xTykAuthorization: "<Your TYK SECRET HERE>",
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
	policies, resp, err := client.PoliciesAPI.ListPolicies(ctx).Execute()
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		// Do something here.
		log.Println(errors.New(resp.Status))

		return
	}
	fmt.Printf("%+v\n", policies)
}
