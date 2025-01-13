# CustomPluginAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSources** | Pointer to [**AuthSources**](AuthSources.md) |  | [optional] 
**Config** | Pointer to [**AuthenticationPlugin**](AuthenticationPlugin.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomPluginAuthentication

`func NewCustomPluginAuthentication() *CustomPluginAuthentication`

NewCustomPluginAuthentication instantiates a new CustomPluginAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPluginAuthenticationWithDefaults

`func NewCustomPluginAuthenticationWithDefaults() *CustomPluginAuthentication`

NewCustomPluginAuthenticationWithDefaults instantiates a new CustomPluginAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSources

`func (o *CustomPluginAuthentication) GetAuthSources() AuthSources`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *CustomPluginAuthentication) GetAuthSourcesOk() (*AuthSources, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *CustomPluginAuthentication) SetAuthSources(v AuthSources)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *CustomPluginAuthentication) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetConfig

`func (o *CustomPluginAuthentication) GetConfig() AuthenticationPlugin`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CustomPluginAuthentication) GetConfigOk() (*AuthenticationPlugin, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CustomPluginAuthentication) SetConfig(v AuthenticationPlugin)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CustomPluginAuthentication) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomPluginAuthentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomPluginAuthentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomPluginAuthentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomPluginAuthentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


