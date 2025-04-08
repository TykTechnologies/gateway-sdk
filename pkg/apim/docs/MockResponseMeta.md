# MockResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewMockResponseMeta

`func NewMockResponseMeta() *MockResponseMeta`

NewMockResponseMeta instantiates a new MockResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockResponseMetaWithDefaults

`func NewMockResponseMetaWithDefaults() *MockResponseMeta`

NewMockResponseMetaWithDefaults instantiates a new MockResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MockResponseMeta) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MockResponseMeta) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MockResponseMeta) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MockResponseMeta) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCode

`func (o *MockResponseMeta) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MockResponseMeta) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MockResponseMeta) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MockResponseMeta) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisabled

`func (o *MockResponseMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MockResponseMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MockResponseMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MockResponseMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHeaders

`func (o *MockResponseMeta) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MockResponseMeta) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MockResponseMeta) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MockResponseMeta) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *MockResponseMeta) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *MockResponseMeta) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetIgnoreCase

`func (o *MockResponseMeta) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *MockResponseMeta) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *MockResponseMeta) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *MockResponseMeta) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetMethod

`func (o *MockResponseMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MockResponseMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MockResponseMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MockResponseMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *MockResponseMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MockResponseMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MockResponseMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MockResponseMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


