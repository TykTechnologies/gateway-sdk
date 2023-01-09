package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/gateway-sdk/gateway"
	"log"
)

func main() {
	ctx := context.Background()
	conf := gateway.Configuration{
		BasePath: "http://localhost:8080",
		Host:     "",
		Scheme:   "",
		DefaultHeader: map[string]string{
			"X-Tyk-Authorization": "foo",
		},
		UserAgent:  "",
		HTTPClient: nil,
	}
	///here we create our client and pass the config file we created above.
	///you can use this client to make any request to the gateway
	client := gateway.NewAPIClient(&conf)
	policy, err := ListPolicy(ctx, client)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", policy)
}

func ListPolicy(ctx context.Context, client *gateway.APIClient) ([]gateway.Policy, error) {
	policies, resp, err := client.PoliciesApi.ListPolicies(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		//Do something here
		return nil, errors.New(resp.Status)

	}

	return policies, nil
}
