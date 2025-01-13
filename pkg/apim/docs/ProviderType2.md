# ProviderType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToPolicyMapping** | Pointer to [**[]ClientToPolicy**](ClientToPolicy.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewProviderType2

`func NewProviderType2() *ProviderType2`

NewProviderType2 instantiates a new ProviderType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderType2WithDefaults

`func NewProviderType2WithDefaults() *ProviderType2`

NewProviderType2WithDefaults instantiates a new ProviderType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToPolicyMapping

`func (o *ProviderType2) GetClientToPolicyMapping() []ClientToPolicy`

GetClientToPolicyMapping returns the ClientToPolicyMapping field if non-nil, zero value otherwise.

### GetClientToPolicyMappingOk

`func (o *ProviderType2) GetClientToPolicyMappingOk() (*[]ClientToPolicy, bool)`

GetClientToPolicyMappingOk returns a tuple with the ClientToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToPolicyMapping

`func (o *ProviderType2) SetClientToPolicyMapping(v []ClientToPolicy)`

SetClientToPolicyMapping sets ClientToPolicyMapping field to given value.

### HasClientToPolicyMapping

`func (o *ProviderType2) HasClientToPolicyMapping() bool`

HasClientToPolicyMapping returns a boolean if a field has been set.

### GetIssuer

`func (o *ProviderType2) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ProviderType2) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ProviderType2) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ProviderType2) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


