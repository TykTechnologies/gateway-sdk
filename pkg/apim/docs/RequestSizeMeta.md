# RequestSizeMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SizeLimit** | Pointer to **int64** |  | [optional] 

## Methods

### NewRequestSizeMeta

`func NewRequestSizeMeta() *RequestSizeMeta`

NewRequestSizeMeta instantiates a new RequestSizeMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSizeMetaWithDefaults

`func NewRequestSizeMetaWithDefaults() *RequestSizeMeta`

NewRequestSizeMetaWithDefaults instantiates a new RequestSizeMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *RequestSizeMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RequestSizeMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RequestSizeMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RequestSizeMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMethod

`func (o *RequestSizeMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestSizeMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestSizeMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestSizeMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *RequestSizeMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestSizeMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestSizeMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RequestSizeMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSizeLimit

`func (o *RequestSizeMeta) GetSizeLimit() int64`

GetSizeLimit returns the SizeLimit field if non-nil, zero value otherwise.

### GetSizeLimitOk

`func (o *RequestSizeMeta) GetSizeLimitOk() (*int64, bool)`

GetSizeLimitOk returns a tuple with the SizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimit

`func (o *RequestSizeMeta) SetSizeLimit(v int64)`

SetSizeLimit sets SizeLimit field to given value.

### HasSizeLimit

`func (o *RequestSizeMeta) HasSizeLimit() bool`

HasSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


