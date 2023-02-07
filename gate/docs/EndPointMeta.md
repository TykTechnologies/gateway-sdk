# EndPointMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodActions** | Pointer to [**map[string]EndpointMethodMeta**](EndpointMethodMeta.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewEndPointMeta

`func NewEndPointMeta() *EndPointMeta`

NewEndPointMeta instantiates a new EndPointMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndPointMetaWithDefaults

`func NewEndPointMetaWithDefaults() *EndPointMeta`

NewEndPointMetaWithDefaults instantiates a new EndPointMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodActions

`func (o *EndPointMeta) GetMethodActions() map[string]EndpointMethodMeta`

GetMethodActions returns the MethodActions field if non-nil, zero value otherwise.

### GetMethodActionsOk

`func (o *EndPointMeta) GetMethodActionsOk() (*map[string]EndpointMethodMeta, bool)`

GetMethodActionsOk returns a tuple with the MethodActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodActions

`func (o *EndPointMeta) SetMethodActions(v map[string]EndpointMethodMeta)`

SetMethodActions sets MethodActions field to given value.

### HasMethodActions

`func (o *EndPointMeta) HasMethodActions() bool`

HasMethodActions returns a boolean if a field has been set.

### GetPath

`func (o *EndPointMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EndPointMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EndPointMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EndPointMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


