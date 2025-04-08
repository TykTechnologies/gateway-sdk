# EndpointMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to [**RateLimitType2**](RateLimitType2.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewEndpointMethod

`func NewEndpointMethod() *EndpointMethod`

NewEndpointMethod instantiates a new EndpointMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointMethodWithDefaults

`func NewEndpointMethodWithDefaults() *EndpointMethod`

NewEndpointMethodWithDefaults instantiates a new EndpointMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *EndpointMethod) GetLimit() RateLimitType2`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EndpointMethod) GetLimitOk() (*RateLimitType2, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EndpointMethod) SetLimit(v RateLimitType2)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EndpointMethod) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *EndpointMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointMethod) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


