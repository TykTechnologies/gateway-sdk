# HeaderInjectionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActOn** | Pointer to **bool** |  | [optional] 
**AddHeaders** | Pointer to **map[string]string** |  | [optional] 
**DeleteHeaders** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewHeaderInjectionMeta

`func NewHeaderInjectionMeta() *HeaderInjectionMeta`

NewHeaderInjectionMeta instantiates a new HeaderInjectionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderInjectionMetaWithDefaults

`func NewHeaderInjectionMetaWithDefaults() *HeaderInjectionMeta`

NewHeaderInjectionMetaWithDefaults instantiates a new HeaderInjectionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActOn

`func (o *HeaderInjectionMeta) GetActOn() bool`

GetActOn returns the ActOn field if non-nil, zero value otherwise.

### GetActOnOk

`func (o *HeaderInjectionMeta) GetActOnOk() (*bool, bool)`

GetActOnOk returns a tuple with the ActOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActOn

`func (o *HeaderInjectionMeta) SetActOn(v bool)`

SetActOn sets ActOn field to given value.

### HasActOn

`func (o *HeaderInjectionMeta) HasActOn() bool`

HasActOn returns a boolean if a field has been set.

### GetAddHeaders

`func (o *HeaderInjectionMeta) GetAddHeaders() map[string]string`

GetAddHeaders returns the AddHeaders field if non-nil, zero value otherwise.

### GetAddHeadersOk

`func (o *HeaderInjectionMeta) GetAddHeadersOk() (*map[string]string, bool)`

GetAddHeadersOk returns a tuple with the AddHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHeaders

`func (o *HeaderInjectionMeta) SetAddHeaders(v map[string]string)`

SetAddHeaders sets AddHeaders field to given value.

### HasAddHeaders

`func (o *HeaderInjectionMeta) HasAddHeaders() bool`

HasAddHeaders returns a boolean if a field has been set.

### SetAddHeadersNil

`func (o *HeaderInjectionMeta) SetAddHeadersNil(b bool)`

 SetAddHeadersNil sets the value for AddHeaders to be an explicit nil

### UnsetAddHeaders
`func (o *HeaderInjectionMeta) UnsetAddHeaders()`

UnsetAddHeaders ensures that no value is present for AddHeaders, not even an explicit nil
### GetDeleteHeaders

`func (o *HeaderInjectionMeta) GetDeleteHeaders() []string`

GetDeleteHeaders returns the DeleteHeaders field if non-nil, zero value otherwise.

### GetDeleteHeadersOk

`func (o *HeaderInjectionMeta) GetDeleteHeadersOk() (*[]string, bool)`

GetDeleteHeadersOk returns a tuple with the DeleteHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHeaders

`func (o *HeaderInjectionMeta) SetDeleteHeaders(v []string)`

SetDeleteHeaders sets DeleteHeaders field to given value.

### HasDeleteHeaders

`func (o *HeaderInjectionMeta) HasDeleteHeaders() bool`

HasDeleteHeaders returns a boolean if a field has been set.

### SetDeleteHeadersNil

`func (o *HeaderInjectionMeta) SetDeleteHeadersNil(b bool)`

 SetDeleteHeadersNil sets the value for DeleteHeaders to be an explicit nil

### UnsetDeleteHeaders
`func (o *HeaderInjectionMeta) UnsetDeleteHeaders()`

UnsetDeleteHeaders ensures that no value is present for DeleteHeaders, not even an explicit nil
### GetDisabled

`func (o *HeaderInjectionMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *HeaderInjectionMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *HeaderInjectionMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *HeaderInjectionMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMethod

`func (o *HeaderInjectionMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HeaderInjectionMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HeaderInjectionMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HeaderInjectionMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *HeaderInjectionMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeaderInjectionMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeaderInjectionMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeaderInjectionMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


