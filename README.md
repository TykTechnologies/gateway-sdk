# Gateway SDK (Alpha )

- This is the gateway sdk generated from the open Api specs.
- Report any issue you may encounter will using this Sdk as we plan to perfect them before making them public.

## How to install

To install the sdk run

`go get github.com/TykTechnologies/gateway-sdk`

## Documentation

For documentation please [look here](https://github.com/TykTechnologies/gateway-sdk/blob/main/gateway/README.md).

## Sample Usage
In these samples tyk is running on localhost  port 8080 .[The BaseUrl is](http://localhost:8080) and the secret key is foo(change this to your actual key).

## Creating the Api client

In this sample we are going to client a and then send a request to the reload endpoint.(Hot Reload
/tyk/reload/)

 ```go 
package main

import (
	"context"
	"errors"
	"github.com/TykTechnologies/gateway-sdk/gateway"
	"github.com/antihax/optional"
	"log"
)

var (
	BaseUrl = "http://localhost:8080"
)

func main() {
	///You can set the baseurl and headers in the configuration.
	conf := gateway.Configuration{
		BasePath: BaseUrl,
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
	////for example in this case we use the client to call the reload endpoint.
	status, resp, err := client.HotReloadApi.HotReload(context.Background(), &gateway.HotReloadApiHotReloadOpts{
		Block: optional.NewBool(true),
	})
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		//Do something here
		log.Println(errors.New(resp.Status))
		return
	}
	log.Println(status.Status)
	log.Println(status.Message)
}
  
  ```


