# TransformHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to [**[]Header**](Header.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Remove** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTransformHeaders

`func NewTransformHeaders() *TransformHeaders`

NewTransformHeaders instantiates a new TransformHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformHeadersWithDefaults

`func NewTransformHeadersWithDefaults() *TransformHeaders`

NewTransformHeadersWithDefaults instantiates a new TransformHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *TransformHeaders) GetAdd() []Header`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *TransformHeaders) GetAddOk() (*[]Header, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *TransformHeaders) SetAdd(v []Header)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *TransformHeaders) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetEnabled

`func (o *TransformHeaders) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TransformHeaders) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TransformHeaders) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TransformHeaders) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRemove

`func (o *TransformHeaders) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *TransformHeaders) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *TransformHeaders) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *TransformHeaders) HasRemove() bool`

HasRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


