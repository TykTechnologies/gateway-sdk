package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TykTechnologies/gateway-sdk/gateway"
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

	policy, err := CreatePolicy(context.Background(), client)
	if err != nil {
		log.Println(err)

		return
	}

	fmt.Printf("%+v\n", policy)
}

func CreatePolicy(ctx context.Context, client *gateway.APIClient) (*gateway.ApiModifyKeySuccess, error) {
	policy := gateway.Policy{
		Name:  "name-1",
		OrgId: "org-1",
	}

	res, resp, err := client.PoliciesApi.AddPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		// Do something here.
		return nil, errors.New(resp.Status)
	}

	return &res, nil
}
