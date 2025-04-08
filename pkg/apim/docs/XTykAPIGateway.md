# XTykAPIGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**Info**](Info.md) |  | [optional] 
**Middleware** | Pointer to [**Middleware**](Middleware.md) |  | [optional] 
**Server** | Pointer to [**Server**](Server.md) |  | [optional] 
**Upstream** | Pointer to [**Upstream**](Upstream.md) |  | [optional] 

## Methods

### NewXTykAPIGateway

`func NewXTykAPIGateway() *XTykAPIGateway`

NewXTykAPIGateway instantiates a new XTykAPIGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXTykAPIGatewayWithDefaults

`func NewXTykAPIGatewayWithDefaults() *XTykAPIGateway`

NewXTykAPIGatewayWithDefaults instantiates a new XTykAPIGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *XTykAPIGateway) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *XTykAPIGateway) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *XTykAPIGateway) SetInfo(v Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *XTykAPIGateway) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetMiddleware

`func (o *XTykAPIGateway) GetMiddleware() Middleware`

GetMiddleware returns the Middleware field if non-nil, zero value otherwise.

### GetMiddlewareOk

`func (o *XTykAPIGateway) GetMiddlewareOk() (*Middleware, bool)`

GetMiddlewareOk returns a tuple with the Middleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleware

`func (o *XTykAPIGateway) SetMiddleware(v Middleware)`

SetMiddleware sets Middleware field to given value.

### HasMiddleware

`func (o *XTykAPIGateway) HasMiddleware() bool`

HasMiddleware returns a boolean if a field has been set.

### GetServer

`func (o *XTykAPIGateway) GetServer() Server`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *XTykAPIGateway) GetServerOk() (*Server, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *XTykAPIGateway) SetServer(v Server)`

SetServer sets Server field to given value.

### HasServer

`func (o *XTykAPIGateway) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpstream

`func (o *XTykAPIGateway) GetUpstream() Upstream`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *XTykAPIGateway) GetUpstreamOk() (*Upstream, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *XTykAPIGateway) SetUpstream(v Upstream)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *XTykAPIGateway) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


