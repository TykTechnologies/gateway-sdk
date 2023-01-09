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
	client := gateway.NewAPIClient(&conf)
	policy, err := CreatePolicy(ctx, client)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", policy)
}

func CreatePolicy(ctx context.Context, client *gateway.APIClient) (*gateway.ApiModifyKeySuccess, error) {
	addPolicyOpt := gateway.PoliciesApiAddPolicyOpts{}
	policy, resp, err := client.PoliciesApi.AddPolicy(ctx, &addPolicyOpt)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		//Do something here
		return nil, errors.New(resp.Status)

	}
	return &policy, nil

}
