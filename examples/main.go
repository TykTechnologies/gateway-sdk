package main

import (
	"context"

	gate2 "github.com/TykTechnologies/gateway-sdk/gateway-dev"

	"log"
)

var (
	baseUrl = "http://localhost:8080/hello"
)

func main() {
	conf := gate2.Configuration{
		Host:             "",
		Scheme:           "",
		DefaultHeader:    nil,
		UserAgent:        "",
		Debug:            false,
		Servers:          gate2.ServerConfigurations{},
		OperationServers: nil,
		HTTPClient:       nil,
	}
	client := gate2.NewAPIClient(&conf)

	ctx := context.Background()
	api, resp, err := client.HealthCheckingApi.HelloAPI(ctx, "hello")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}
	log.Println(api)

}
