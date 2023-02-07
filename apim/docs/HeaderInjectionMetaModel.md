# HeaderInjectionMetaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActOn** | Pointer to **bool** |  | [optional] 
**AddHeaders** | Pointer to **map[string]string** |  | [optional] 
**DeleteHeaders** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewHeaderInjectionMetaModel

`func NewHeaderInjectionMetaModel() *HeaderInjectionMetaModel`

NewHeaderInjectionMetaModel instantiates a new HeaderInjectionMetaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderInjectionMetaModelWithDefaults

`func NewHeaderInjectionMetaModelWithDefaults() *HeaderInjectionMetaModel`

NewHeaderInjectionMetaModelWithDefaults instantiates a new HeaderInjectionMetaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActOn

`func (o *HeaderInjectionMetaModel) GetActOn() bool`

GetActOn returns the ActOn field if non-nil, zero value otherwise.

### GetActOnOk

`func (o *HeaderInjectionMetaModel) GetActOnOk() (*bool, bool)`

GetActOnOk returns a tuple with the ActOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActOn

`func (o *HeaderInjectionMetaModel) SetActOn(v bool)`

SetActOn sets ActOn field to given value.

### HasActOn

`func (o *HeaderInjectionMetaModel) HasActOn() bool`

HasActOn returns a boolean if a field has been set.

### GetAddHeaders

`func (o *HeaderInjectionMetaModel) GetAddHeaders() map[string]string`

GetAddHeaders returns the AddHeaders field if non-nil, zero value otherwise.

### GetAddHeadersOk

`func (o *HeaderInjectionMetaModel) GetAddHeadersOk() (*map[string]string, bool)`

GetAddHeadersOk returns a tuple with the AddHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHeaders

`func (o *HeaderInjectionMetaModel) SetAddHeaders(v map[string]string)`

SetAddHeaders sets AddHeaders field to given value.

### HasAddHeaders

`func (o *HeaderInjectionMetaModel) HasAddHeaders() bool`

HasAddHeaders returns a boolean if a field has been set.

### GetDeleteHeaders

`func (o *HeaderInjectionMetaModel) GetDeleteHeaders() []string`

GetDeleteHeaders returns the DeleteHeaders field if non-nil, zero value otherwise.

### GetDeleteHeadersOk

`func (o *HeaderInjectionMetaModel) GetDeleteHeadersOk() (*[]string, bool)`

GetDeleteHeadersOk returns a tuple with the DeleteHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHeaders

`func (o *HeaderInjectionMetaModel) SetDeleteHeaders(v []string)`

SetDeleteHeaders sets DeleteHeaders field to given value.

### HasDeleteHeaders

`func (o *HeaderInjectionMetaModel) HasDeleteHeaders() bool`

HasDeleteHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *HeaderInjectionMetaModel) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HeaderInjectionMetaModel) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HeaderInjectionMetaModel) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HeaderInjectionMetaModel) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *HeaderInjectionMetaModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeaderInjectionMetaModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeaderInjectionMetaModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeaderInjectionMetaModel) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


