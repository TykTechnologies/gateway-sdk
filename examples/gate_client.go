package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/gateway-sdk/gate"
	"log"
)

func main() {
	cfg := &gate.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: gate.ServerConfigurations{
			{
				URL:         "http://localhost:8080",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]gate.ServerConfigurations{},
	}
	client := gate.NewAPIClient(cfg)
	ctx := context.Background()
	statusMessage, resp, err := client.PoliciesApi.DeletePolicy(ctx, "ooiiii").Execute()
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		//Do something here
		log.Println(errors.New(resp.Status))
		return
	}
	fmt.Printf("%+v\n", statusMessage)

}
