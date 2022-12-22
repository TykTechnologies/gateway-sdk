package main

import (
	"context"
	"github.com/TykTechnologies/gateway-sdk/gateway"

	"log"
)

var (
	baseUrl = "http://localhost:8080"
)

func main() {
	conf := gateway.Configuration{
		BasePath:      baseUrl,
		Host:          "",
		Scheme:        "",
		DefaultHeader: map[string]string{},
		UserAgent:     "",
		HTTPClient:    nil,
	}
	client := gateway.NewAPIClient(&conf)
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
