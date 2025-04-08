# EndPointMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
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

### GetDisabled

`func (o *EndPointMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *EndPointMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *EndPointMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *EndPointMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *EndPointMeta) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *EndPointMeta) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *EndPointMeta) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *EndPointMeta) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetMethod

`func (o *EndPointMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EndPointMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EndPointMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *EndPointMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

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


