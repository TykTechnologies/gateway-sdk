# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to [**[]EndpointMethod**](EndpointMethod.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *Endpoint) GetMethods() []EndpointMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Endpoint) GetMethodsOk() (*[]EndpointMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Endpoint) SetMethods(v []EndpointMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *Endpoint) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetPath

`func (o *Endpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Endpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Endpoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Endpoint) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


