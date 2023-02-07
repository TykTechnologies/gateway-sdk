# APIDefinitionCORS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCredentials** | Pointer to **bool** |  | [optional] 
**AllowedHeaders** | Pointer to **[]string** |  | [optional] 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 
**AllowedOrigins** | Pointer to **[]string** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**ExposedHeaders** | Pointer to **[]string** |  | [optional] 
**MaxAge** | Pointer to **int64** |  | [optional] 
**OptionsPassthrough** | Pointer to **bool** |  | [optional] 

## Methods

### NewAPIDefinitionCORS

`func NewAPIDefinitionCORS() *APIDefinitionCORS`

NewAPIDefinitionCORS instantiates a new APIDefinitionCORS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionCORSWithDefaults

`func NewAPIDefinitionCORSWithDefaults() *APIDefinitionCORS`

NewAPIDefinitionCORSWithDefaults instantiates a new APIDefinitionCORS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCredentials

`func (o *APIDefinitionCORS) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *APIDefinitionCORS) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *APIDefinitionCORS) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *APIDefinitionCORS) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetAllowedHeaders

`func (o *APIDefinitionCORS) GetAllowedHeaders() []string`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *APIDefinitionCORS) GetAllowedHeadersOk() (*[]string, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *APIDefinitionCORS) SetAllowedHeaders(v []string)`

SetAllowedHeaders sets AllowedHeaders field to given value.

### HasAllowedHeaders

`func (o *APIDefinitionCORS) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *APIDefinitionCORS) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *APIDefinitionCORS) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *APIDefinitionCORS) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *APIDefinitionCORS) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *APIDefinitionCORS) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *APIDefinitionCORS) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *APIDefinitionCORS) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *APIDefinitionCORS) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetDebug

`func (o *APIDefinitionCORS) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *APIDefinitionCORS) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *APIDefinitionCORS) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *APIDefinitionCORS) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetEnable

`func (o *APIDefinitionCORS) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *APIDefinitionCORS) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *APIDefinitionCORS) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *APIDefinitionCORS) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExposedHeaders

`func (o *APIDefinitionCORS) GetExposedHeaders() []string`

GetExposedHeaders returns the ExposedHeaders field if non-nil, zero value otherwise.

### GetExposedHeadersOk

`func (o *APIDefinitionCORS) GetExposedHeadersOk() (*[]string, bool)`

GetExposedHeadersOk returns a tuple with the ExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedHeaders

`func (o *APIDefinitionCORS) SetExposedHeaders(v []string)`

SetExposedHeaders sets ExposedHeaders field to given value.

### HasExposedHeaders

`func (o *APIDefinitionCORS) HasExposedHeaders() bool`

HasExposedHeaders returns a boolean if a field has been set.

### GetMaxAge

`func (o *APIDefinitionCORS) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *APIDefinitionCORS) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *APIDefinitionCORS) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *APIDefinitionCORS) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetOptionsPassthrough

`func (o *APIDefinitionCORS) GetOptionsPassthrough() bool`

GetOptionsPassthrough returns the OptionsPassthrough field if non-nil, zero value otherwise.

### GetOptionsPassthroughOk

`func (o *APIDefinitionCORS) GetOptionsPassthroughOk() (*bool, bool)`

GetOptionsPassthroughOk returns a tuple with the OptionsPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsPassthrough

`func (o *APIDefinitionCORS) SetOptionsPassthrough(v bool)`

SetOptionsPassthrough sets OptionsPassthrough field to given value.

### HasOptionsPassthrough

`func (o *APIDefinitionCORS) HasOptionsPassthrough() bool`

HasOptionsPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


