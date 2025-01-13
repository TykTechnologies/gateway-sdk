# OpenIDOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]OIDProviderConfig**](OIDProviderConfig.md) |  | [optional] 
**SegregateByClient** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenIDOptions

`func NewOpenIDOptions() *OpenIDOptions`

NewOpenIDOptions instantiates a new OpenIDOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDOptionsWithDefaults

`func NewOpenIDOptionsWithDefaults() *OpenIDOptions`

NewOpenIDOptionsWithDefaults instantiates a new OpenIDOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *OpenIDOptions) GetProviders() []OIDProviderConfig`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OpenIDOptions) GetProvidersOk() (*[]OIDProviderConfig, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OpenIDOptions) SetProviders(v []OIDProviderConfig)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *OpenIDOptions) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *OpenIDOptions) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *OpenIDOptions) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetSegregateByClient

`func (o *OpenIDOptions) GetSegregateByClient() bool`

GetSegregateByClient returns the SegregateByClient field if non-nil, zero value otherwise.

### GetSegregateByClientOk

`func (o *OpenIDOptions) GetSegregateByClientOk() (*bool, bool)`

GetSegregateByClientOk returns a tuple with the SegregateByClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegregateByClient

`func (o *OpenIDOptions) SetSegregateByClient(v bool)`

SetSegregateByClient sets SegregateByClient field to given value.

### HasSegregateByClient

`func (o *OpenIDOptions) HasSegregateByClient() bool`

HasSegregateByClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


