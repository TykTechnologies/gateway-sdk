# Gateway SDK ( Alpha )

- This is the gateway sdk generated from the open Api specs.
- Report any issue you may encounter while using this Sdk.

## Versioning

- For versionning since the sdks are still experimental we are using a tag like: v0.0.1-gateway.5.7.2
 The -gateway.5.7.2 indicates the gateway that this sdks were generated for.

## How to install

To install the sdk run

`go get github.com/TykTechnologies/gateway-sdk` (For master branch)

For a branch like 5.7.2 You would install as follows:

 go get github.com/TykTechnologies/gateway-sdk@v0.0.1-gateway.5.7.2`



## Sample Usage
In these samples Tyk Gateway is running on localhost port 8080.
The BaseUrl is: `http://localhost:8080` .

## Example to list Apis in the gateway
```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

var (
	BaseUrl           = "http://tyk-gateway.localhost:8080"
	xTykAuthorization = "x-tyk-authorization"
)

func main() {
	apiConfig := apim.Configuration{
		DefaultHeader: map[string]string{
			xTykAuthorization: "<YOUR GATEWAY AUTHORIZATION KEY HERE>",
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
	apis, resp, err := client.APIsAPI.ListApis(ctx).Execute()
	if err != nil || resp.StatusCode != 200 {
		body, err := ReadResponseBody(resp)
		if err != nil {
			log.Println(resp.Status)
			return
		}
		log.Println(body)
		return
	}
	for _, item := range apis {
		fmt.Printf("%+v\n", item.GetApiId())
	}
}

func ReadResponseBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert the body to a string and return
	return string(body), nil
}
```

## Documentation

For documentation please [look here](https://github.com/TykTechnologies/gateway-sdk/blob/main/pkg/apim/README.md).
