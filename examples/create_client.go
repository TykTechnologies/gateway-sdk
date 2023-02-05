package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/antihax/optional"

	"github.com/TykTechnologies/gateway-sdk/gateway"
)

var (
	BaseUrl = "http://localhost:8080"
)

func main() {
	// You can set the base url and headers in the configuration.
	cfg := gateway.NewConfiguration()

	cfg.BasePath = BaseUrl
	cfg.AddDefaultHeader("X-Tyk-Authorization", "secret")
	cfg.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}

	// Create Tyk gateway SDK client and pass configuration we created.
	// You can use this client now, to make any request to the gateway.
	client := gateway.NewAPIClient(cfg)

	// For example, in bellow case we use the client to call the reload endpoint.
	status, resp, err := client.HotReloadApi.HotReload(context.Background(), &gateway.HotReloadApiHotReloadOpts{
		Block: optional.NewBool(true),
	})
	if err != nil {
		log.Println(err)

		return
	}

	if resp.StatusCode != 200 {
		// Do something here.
		log.Println(errors.New(resp.Status))
		
		return
	}

	log.Println(status.Status)
	log.Println(status.Message)
}
