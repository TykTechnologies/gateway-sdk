# AuthenticationPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**FunctionName** | Pointer to **string** |  | [optional] 
**IdExtractor** | Pointer to [**IDExtractor**](IDExtractor.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RawBodyOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthenticationPlugin

`func NewAuthenticationPlugin() *AuthenticationPlugin`

NewAuthenticationPlugin instantiates a new AuthenticationPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPluginWithDefaults

`func NewAuthenticationPluginWithDefaults() *AuthenticationPlugin`

NewAuthenticationPluginWithDefaults instantiates a new AuthenticationPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuthenticationPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthenticationPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthenticationPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthenticationPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFunctionName

`func (o *AuthenticationPlugin) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AuthenticationPlugin) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AuthenticationPlugin) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *AuthenticationPlugin) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetIdExtractor

`func (o *AuthenticationPlugin) GetIdExtractor() IDExtractor`

GetIdExtractor returns the IdExtractor field if non-nil, zero value otherwise.

### GetIdExtractorOk

`func (o *AuthenticationPlugin) GetIdExtractorOk() (*IDExtractor, bool)`

GetIdExtractorOk returns a tuple with the IdExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractor

`func (o *AuthenticationPlugin) SetIdExtractor(v IDExtractor)`

SetIdExtractor sets IdExtractor field to given value.

### HasIdExtractor

`func (o *AuthenticationPlugin) HasIdExtractor() bool`

HasIdExtractor returns a boolean if a field has been set.

### GetPath

`func (o *AuthenticationPlugin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AuthenticationPlugin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AuthenticationPlugin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AuthenticationPlugin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRawBodyOnly

`func (o *AuthenticationPlugin) GetRawBodyOnly() bool`

GetRawBodyOnly returns the RawBodyOnly field if non-nil, zero value otherwise.

### GetRawBodyOnlyOk

`func (o *AuthenticationPlugin) GetRawBodyOnlyOk() (*bool, bool)`

GetRawBodyOnlyOk returns a tuple with the RawBodyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBodyOnly

`func (o *AuthenticationPlugin) SetRawBodyOnly(v bool)`

SetRawBodyOnly sets RawBodyOnly field to given value.

### HasRawBodyOnly

`func (o *AuthenticationPlugin) HasRawBodyOnly() bool`

HasRawBodyOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


