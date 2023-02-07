# OpenIDOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]OIDProviderConfigModelModel**](OIDProviderConfigModelModel.md) |  | [optional] 
**SegregateByClient** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenIDOptionsModel

`func NewOpenIDOptionsModel() *OpenIDOptionsModel`

NewOpenIDOptionsModel instantiates a new OpenIDOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDOptionsModelWithDefaults

`func NewOpenIDOptionsModelWithDefaults() *OpenIDOptionsModel`

NewOpenIDOptionsModelWithDefaults instantiates a new OpenIDOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *OpenIDOptionsModel) GetProviders() []OIDProviderConfigModelModel`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OpenIDOptionsModel) GetProvidersOk() (*[]OIDProviderConfigModelModel, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OpenIDOptionsModel) SetProviders(v []OIDProviderConfigModelModel)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *OpenIDOptionsModel) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetSegregateByClient

`func (o *OpenIDOptionsModel) GetSegregateByClient() bool`

GetSegregateByClient returns the SegregateByClient field if non-nil, zero value otherwise.

### GetSegregateByClientOk

`func (o *OpenIDOptionsModel) GetSegregateByClientOk() (*bool, bool)`

GetSegregateByClientOk returns a tuple with the SegregateByClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegregateByClient

`func (o *OpenIDOptionsModel) SetSegregateByClient(v bool)`

SetSegregateByClient sets SegregateByClient field to given value.

### HasSegregateByClient

`func (o *OpenIDOptionsModel) HasSegregateByClient() bool`

HasSegregateByClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


