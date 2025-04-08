# Scopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to [**ScopeClaim**](ScopeClaim.md) |  | [optional] 
**Oidc** | Pointer to [**ScopeClaim**](ScopeClaim.md) |  | [optional] 

## Methods

### NewScopes

`func NewScopes() *Scopes`

NewScopes instantiates a new Scopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesWithDefaults

`func NewScopesWithDefaults() *Scopes`

NewScopesWithDefaults instantiates a new Scopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *Scopes) GetJwt() ScopeClaim`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *Scopes) GetJwtOk() (*ScopeClaim, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *Scopes) SetJwt(v ScopeClaim)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *Scopes) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetOidc

`func (o *Scopes) GetOidc() ScopeClaim`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *Scopes) GetOidcOk() (*ScopeClaim, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *Scopes) SetOidc(v ScopeClaim)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *Scopes) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


