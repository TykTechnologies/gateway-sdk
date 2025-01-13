# OIDC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSources** | Pointer to [**AuthSources**](AuthSources.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**[]ProviderType2**](ProviderType2.md) |  | [optional] 
**Scopes** | Pointer to [**ScopesType2**](ScopesType2.md) |  | [optional] 
**SegregateByClientId** | Pointer to **bool** |  | [optional] 

## Methods

### NewOIDC

`func NewOIDC() *OIDC`

NewOIDC instantiates a new OIDC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWithDefaults

`func NewOIDCWithDefaults() *OIDC`

NewOIDCWithDefaults instantiates a new OIDC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSources

`func (o *OIDC) GetAuthSources() AuthSources`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *OIDC) GetAuthSourcesOk() (*AuthSources, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *OIDC) SetAuthSources(v AuthSources)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *OIDC) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetEnabled

`func (o *OIDC) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OIDC) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OIDC) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OIDC) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviders

`func (o *OIDC) GetProviders() []ProviderType2`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OIDC) GetProvidersOk() (*[]ProviderType2, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OIDC) SetProviders(v []ProviderType2)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *OIDC) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetScopes

`func (o *OIDC) GetScopes() ScopesType2`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OIDC) GetScopesOk() (*ScopesType2, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OIDC) SetScopes(v ScopesType2)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OIDC) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSegregateByClientId

`func (o *OIDC) GetSegregateByClientId() bool`

GetSegregateByClientId returns the SegregateByClientId field if non-nil, zero value otherwise.

### GetSegregateByClientIdOk

`func (o *OIDC) GetSegregateByClientIdOk() (*bool, bool)`

GetSegregateByClientIdOk returns a tuple with the SegregateByClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegregateByClientId

`func (o *OIDC) SetSegregateByClientId(v bool)`

SetSegregateByClientId sets SegregateByClientId field to given value.

### HasSegregateByClientId

`func (o *OIDC) HasSegregateByClientId() bool`

HasSegregateByClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


