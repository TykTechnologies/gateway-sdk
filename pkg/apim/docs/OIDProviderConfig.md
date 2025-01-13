# OIDProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIds** | Pointer to **map[string]string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewOIDProviderConfig

`func NewOIDProviderConfig() *OIDProviderConfig`

NewOIDProviderConfig instantiates a new OIDProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDProviderConfigWithDefaults

`func NewOIDProviderConfigWithDefaults() *OIDProviderConfig`

NewOIDProviderConfigWithDefaults instantiates a new OIDProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIds

`func (o *OIDProviderConfig) GetClientIds() map[string]string`

GetClientIds returns the ClientIds field if non-nil, zero value otherwise.

### GetClientIdsOk

`func (o *OIDProviderConfig) GetClientIdsOk() (*map[string]string, bool)`

GetClientIdsOk returns a tuple with the ClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIds

`func (o *OIDProviderConfig) SetClientIds(v map[string]string)`

SetClientIds sets ClientIds field to given value.

### HasClientIds

`func (o *OIDProviderConfig) HasClientIds() bool`

HasClientIds returns a boolean if a field has been set.

### SetClientIdsNil

`func (o *OIDProviderConfig) SetClientIdsNil(b bool)`

 SetClientIdsNil sets the value for ClientIds to be an explicit nil

### UnsetClientIds
`func (o *OIDProviderConfig) UnsetClientIds()`

UnsetClientIds ensures that no value is present for ClientIds, not even an explicit nil
### GetIssuer

`func (o *OIDProviderConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDProviderConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDProviderConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OIDProviderConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


