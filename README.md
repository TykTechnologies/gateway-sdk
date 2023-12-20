# Gateway SDK ( Alpha )

- This is the gateway sdk generated from the open Api specs.
- Report any issue you may encounter will using this Sdk as we plan to perfect them before making them public.

## How to install

To install the sdk run

`go get github.com/TykTechnologies/gateway-sdk`

## Sample Usage
In these samples Tyk Gateway is running on localhost port 8080.
The BaseUrl is: `http://localhost:8080` and the secret key is `secret`, you should change this to your actual key.

## List Policies
```go
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
```

## Documentation

For documentation please [look here](https://github.com/TykTechnologies/gateway-sdk/blob/main/pkg/apim/README.md).
