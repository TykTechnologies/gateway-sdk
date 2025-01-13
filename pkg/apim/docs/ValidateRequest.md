# ValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorResponseCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewValidateRequest

`func NewValidateRequest() *ValidateRequest`

NewValidateRequest instantiates a new ValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateRequestWithDefaults

`func NewValidateRequestWithDefaults() *ValidateRequest`

NewValidateRequestWithDefaults instantiates a new ValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ValidateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ValidateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ValidateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ValidateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorResponseCode

`func (o *ValidateRequest) GetErrorResponseCode() int32`

GetErrorResponseCode returns the ErrorResponseCode field if non-nil, zero value otherwise.

### GetErrorResponseCodeOk

`func (o *ValidateRequest) GetErrorResponseCodeOk() (*int32, bool)`

GetErrorResponseCodeOk returns a tuple with the ErrorResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponseCode

`func (o *ValidateRequest) SetErrorResponseCode(v int32)`

SetErrorResponseCode sets ErrorResponseCode field to given value.

### HasErrorResponseCode

`func (o *ValidateRequest) HasErrorResponseCode() bool`

HasErrorResponseCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


