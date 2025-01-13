# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Introspection** | Pointer to [**Introspection**](Introspection.md) |  | [optional] 
**Jwt** | Pointer to [**JWTValidation**](JWTValidation.md) |  | [optional] 

## Methods

### NewProvider

`func NewProvider() *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntrospection

`func (o *Provider) GetIntrospection() Introspection`

GetIntrospection returns the Introspection field if non-nil, zero value otherwise.

### GetIntrospectionOk

`func (o *Provider) GetIntrospectionOk() (*Introspection, bool)`

GetIntrospectionOk returns a tuple with the Introspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospection

`func (o *Provider) SetIntrospection(v Introspection)`

SetIntrospection sets Introspection field to given value.

### HasIntrospection

`func (o *Provider) HasIntrospection() bool`

HasIntrospection returns a boolean if a field has been set.

### GetJwt

`func (o *Provider) GetJwt() JWTValidation`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *Provider) GetJwtOk() (*JWTValidation, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *Provider) SetJwt(v JWTValidation)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *Provider) HasJwt() bool`

HasJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


