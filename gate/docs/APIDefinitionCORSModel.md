# APIDefinitionCORSModel

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

### NewAPIDefinitionCORSModel

`func NewAPIDefinitionCORSModel() *APIDefinitionCORSModel`

NewAPIDefinitionCORSModel instantiates a new APIDefinitionCORSModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionCORSModelWithDefaults

`func NewAPIDefinitionCORSModelWithDefaults() *APIDefinitionCORSModel`

NewAPIDefinitionCORSModelWithDefaults instantiates a new APIDefinitionCORSModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCredentials

`func (o *APIDefinitionCORSModel) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *APIDefinitionCORSModel) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *APIDefinitionCORSModel) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *APIDefinitionCORSModel) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetAllowedHeaders

`func (o *APIDefinitionCORSModel) GetAllowedHeaders() []string`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *APIDefinitionCORSModel) GetAllowedHeadersOk() (*[]string, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *APIDefinitionCORSModel) SetAllowedHeaders(v []string)`

SetAllowedHeaders sets AllowedHeaders field to given value.

### HasAllowedHeaders

`func (o *APIDefinitionCORSModel) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *APIDefinitionCORSModel) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *APIDefinitionCORSModel) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *APIDefinitionCORSModel) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *APIDefinitionCORSModel) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *APIDefinitionCORSModel) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *APIDefinitionCORSModel) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *APIDefinitionCORSModel) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *APIDefinitionCORSModel) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetDebug

`func (o *APIDefinitionCORSModel) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *APIDefinitionCORSModel) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *APIDefinitionCORSModel) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *APIDefinitionCORSModel) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetEnable

`func (o *APIDefinitionCORSModel) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *APIDefinitionCORSModel) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *APIDefinitionCORSModel) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *APIDefinitionCORSModel) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExposedHeaders

`func (o *APIDefinitionCORSModel) GetExposedHeaders() []string`

GetExposedHeaders returns the ExposedHeaders field if non-nil, zero value otherwise.

### GetExposedHeadersOk

`func (o *APIDefinitionCORSModel) GetExposedHeadersOk() (*[]string, bool)`

GetExposedHeadersOk returns a tuple with the ExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedHeaders

`func (o *APIDefinitionCORSModel) SetExposedHeaders(v []string)`

SetExposedHeaders sets ExposedHeaders field to given value.

### HasExposedHeaders

`func (o *APIDefinitionCORSModel) HasExposedHeaders() bool`

HasExposedHeaders returns a boolean if a field has been set.

### GetMaxAge

`func (o *APIDefinitionCORSModel) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *APIDefinitionCORSModel) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *APIDefinitionCORSModel) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *APIDefinitionCORSModel) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetOptionsPassthrough

`func (o *APIDefinitionCORSModel) GetOptionsPassthrough() bool`

GetOptionsPassthrough returns the OptionsPassthrough field if non-nil, zero value otherwise.

### GetOptionsPassthroughOk

`func (o *APIDefinitionCORSModel) GetOptionsPassthroughOk() (*bool, bool)`

GetOptionsPassthroughOk returns a tuple with the OptionsPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsPassthrough

`func (o *APIDefinitionCORSModel) SetOptionsPassthrough(v bool)`

SetOptionsPassthrough sets OptionsPassthrough field to given value.

### HasOptionsPassthrough

`func (o *APIDefinitionCORSModel) HasOptionsPassthrough() bool`

HasOptionsPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


