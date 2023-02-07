# OIDProviderConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIds** | Pointer to **map[string]string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewOIDProviderConfigModel

`func NewOIDProviderConfigModel() *OIDProviderConfigModel`

NewOIDProviderConfigModel instantiates a new OIDProviderConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDProviderConfigModelWithDefaults

`func NewOIDProviderConfigModelWithDefaults() *OIDProviderConfigModel`

NewOIDProviderConfigModelWithDefaults instantiates a new OIDProviderConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIds

`func (o *OIDProviderConfigModel) GetClientIds() map[string]string`

GetClientIds returns the ClientIds field if non-nil, zero value otherwise.

### GetClientIdsOk

`func (o *OIDProviderConfigModel) GetClientIdsOk() (*map[string]string, bool)`

GetClientIdsOk returns a tuple with the ClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIds

`func (o *OIDProviderConfigModel) SetClientIds(v map[string]string)`

SetClientIds sets ClientIds field to given value.

### HasClientIds

`func (o *OIDProviderConfigModel) HasClientIds() bool`

HasClientIds returns a boolean if a field has been set.

### GetIssuer

`func (o *OIDProviderConfigModel) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDProviderConfigModel) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDProviderConfigModel) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OIDProviderConfigModel) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


