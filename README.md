# Gateway SDK ( Alpha )

- This is the gateway sdk generated from the open Api specs.
- Report any issue you may encounter will using this Sdk as we plan to perfect them before making them public.

## How to install

To install the sdk run

`go get github.com/TykTechnologies/gateway-sdk`

## Sample Usage
In these samples Tyk Gateway is running on localhost port 8080.
The BaseUrl is: `http://localhost:8080` and the secret key is `secret`, you should change this to your actual key.

## Creating the Api client

In this sample we are going to client a and then send a request to the reload endpoint.
Hot Reload endpoint path is at `/tyk/reload/`.

 ```go 
package main

import (
	"context"
	"errors"
	"log"

	"github.com/antihax/optional"

	"github.com/TykTechnologies/gateway-sdk/gateway"
)

var (
	BaseUrl = "http://localhost:8080"
)

func main() {
	// You can set the base url and headers in the configuration.
	conf := gateway.Configuration{
		BasePath: BaseUrl,
		Host:     "",
		Scheme:   "",
		DefaultHeader: map[string]string{
			"X-Tyk-Authorization": "secret",
		},
		UserAgent:  "",
		HTTPClient: nil,
	}

	// Create Tyk gateway SDK client and pass configuration we created.
	// You can use this client now, to make any request to the gateway.
	client := gateway.NewAPIClient(&conf)

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
```

## List Policies
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/TykTechnologies/gateway-sdk/gateway"
)

func main() {
	// You can set the base url and headers in the configuration.
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

	// Create Tyk gateway SDK client and pass configuration we created.
	// You can use this client now, to make any request to the gateway.
	client := gateway.NewAPIClient(&conf)

	// For example, in bellow case we use the client to list policies.
	policy, err := ListPolicy(context.Background(), client)
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
		// Do something here.
		log.Println(errors.New(resp.Status))

		return nil, errors.New(resp.Status)
	}

	return policies, nil
}
```

## Documentation

For documentation please [look here](https://github.com/TykTechnologies/gateway-sdk/blob/main/gateway/README.md).
