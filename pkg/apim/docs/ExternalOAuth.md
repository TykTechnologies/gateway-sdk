# ExternalOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**[]Provider**](Provider.md) |  | [optional] 

## Methods

### NewExternalOAuth

`func NewExternalOAuth() *ExternalOAuth`

NewExternalOAuth instantiates a new ExternalOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalOAuthWithDefaults

`func NewExternalOAuthWithDefaults() *ExternalOAuth`

NewExternalOAuthWithDefaults instantiates a new ExternalOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ExternalOAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExternalOAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExternalOAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExternalOAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviders

`func (o *ExternalOAuth) GetProviders() []Provider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ExternalOAuth) GetProvidersOk() (*[]Provider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ExternalOAuth) SetProviders(v []Provider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ExternalOAuth) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *ExternalOAuth) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *ExternalOAuth) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


