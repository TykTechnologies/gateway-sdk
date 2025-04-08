# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseIdentityProvider** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to [**CustomPluginAuthentication**](CustomPluginAuthentication.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Hmac** | Pointer to [**HMAC**](HMAC.md) |  | [optional] 
**Oidc** | Pointer to [**OIDC**](OIDC.md) |  | [optional] 
**SecuritySchemes** | Pointer to **map[string]interface{}** |  | [optional] 
**StripAuthorizationData** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseIdentityProvider

`func (o *Authentication) GetBaseIdentityProvider() string`

GetBaseIdentityProvider returns the BaseIdentityProvider field if non-nil, zero value otherwise.

### GetBaseIdentityProviderOk

`func (o *Authentication) GetBaseIdentityProviderOk() (*string, bool)`

GetBaseIdentityProviderOk returns a tuple with the BaseIdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIdentityProvider

`func (o *Authentication) SetBaseIdentityProvider(v string)`

SetBaseIdentityProvider sets BaseIdentityProvider field to given value.

### HasBaseIdentityProvider

`func (o *Authentication) HasBaseIdentityProvider() bool`

HasBaseIdentityProvider returns a boolean if a field has been set.

### GetCustom

`func (o *Authentication) GetCustom() CustomPluginAuthentication`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *Authentication) GetCustomOk() (*CustomPluginAuthentication, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *Authentication) SetCustom(v CustomPluginAuthentication)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *Authentication) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEnabled

`func (o *Authentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Authentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Authentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Authentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHmac

`func (o *Authentication) GetHmac() HMAC`

GetHmac returns the Hmac field if non-nil, zero value otherwise.

### GetHmacOk

`func (o *Authentication) GetHmacOk() (*HMAC, bool)`

GetHmacOk returns a tuple with the Hmac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmac

`func (o *Authentication) SetHmac(v HMAC)`

SetHmac sets Hmac field to given value.

### HasHmac

`func (o *Authentication) HasHmac() bool`

HasHmac returns a boolean if a field has been set.

### GetOidc

`func (o *Authentication) GetOidc() OIDC`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *Authentication) GetOidcOk() (*OIDC, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *Authentication) SetOidc(v OIDC)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *Authentication) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetSecuritySchemes

`func (o *Authentication) GetSecuritySchemes() map[string]interface{}`

GetSecuritySchemes returns the SecuritySchemes field if non-nil, zero value otherwise.

### GetSecuritySchemesOk

`func (o *Authentication) GetSecuritySchemesOk() (*map[string]interface{}, bool)`

GetSecuritySchemesOk returns a tuple with the SecuritySchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySchemes

`func (o *Authentication) SetSecuritySchemes(v map[string]interface{})`

SetSecuritySchemes sets SecuritySchemes field to given value.

### HasSecuritySchemes

`func (o *Authentication) HasSecuritySchemes() bool`

HasSecuritySchemes returns a boolean if a field has been set.

### GetStripAuthorizationData

`func (o *Authentication) GetStripAuthorizationData() bool`

GetStripAuthorizationData returns the StripAuthorizationData field if non-nil, zero value otherwise.

### GetStripAuthorizationDataOk

`func (o *Authentication) GetStripAuthorizationDataOk() (*bool, bool)`

GetStripAuthorizationDataOk returns a tuple with the StripAuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripAuthorizationData

`func (o *Authentication) SetStripAuthorizationData(v bool)`

SetStripAuthorizationData sets StripAuthorizationData field to given value.

### HasStripAuthorizationData

`func (o *Authentication) HasStripAuthorizationData() bool`

HasStripAuthorizationData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


