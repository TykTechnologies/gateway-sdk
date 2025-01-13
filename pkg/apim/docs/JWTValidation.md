# JWTValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ExpiresAtValidationSkew** | Pointer to **int32** |  | [optional] 
**IdentityBaseField** | Pointer to **string** |  | [optional] 
**IssuedAtValidationSkew** | Pointer to **int32** |  | [optional] 
**NotBeforeValidationSkew** | Pointer to **int32** |  | [optional] 
**SigningMethod** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewJWTValidation

`func NewJWTValidation() *JWTValidation`

NewJWTValidation instantiates a new JWTValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTValidationWithDefaults

`func NewJWTValidationWithDefaults() *JWTValidation`

NewJWTValidationWithDefaults instantiates a new JWTValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JWTValidation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JWTValidation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JWTValidation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JWTValidation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiresAtValidationSkew

`func (o *JWTValidation) GetExpiresAtValidationSkew() int32`

GetExpiresAtValidationSkew returns the ExpiresAtValidationSkew field if non-nil, zero value otherwise.

### GetExpiresAtValidationSkewOk

`func (o *JWTValidation) GetExpiresAtValidationSkewOk() (*int32, bool)`

GetExpiresAtValidationSkewOk returns a tuple with the ExpiresAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtValidationSkew

`func (o *JWTValidation) SetExpiresAtValidationSkew(v int32)`

SetExpiresAtValidationSkew sets ExpiresAtValidationSkew field to given value.

### HasExpiresAtValidationSkew

`func (o *JWTValidation) HasExpiresAtValidationSkew() bool`

HasExpiresAtValidationSkew returns a boolean if a field has been set.

### GetIdentityBaseField

`func (o *JWTValidation) GetIdentityBaseField() string`

GetIdentityBaseField returns the IdentityBaseField field if non-nil, zero value otherwise.

### GetIdentityBaseFieldOk

`func (o *JWTValidation) GetIdentityBaseFieldOk() (*string, bool)`

GetIdentityBaseFieldOk returns a tuple with the IdentityBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityBaseField

`func (o *JWTValidation) SetIdentityBaseField(v string)`

SetIdentityBaseField sets IdentityBaseField field to given value.

### HasIdentityBaseField

`func (o *JWTValidation) HasIdentityBaseField() bool`

HasIdentityBaseField returns a boolean if a field has been set.

### GetIssuedAtValidationSkew

`func (o *JWTValidation) GetIssuedAtValidationSkew() int32`

GetIssuedAtValidationSkew returns the IssuedAtValidationSkew field if non-nil, zero value otherwise.

### GetIssuedAtValidationSkewOk

`func (o *JWTValidation) GetIssuedAtValidationSkewOk() (*int32, bool)`

GetIssuedAtValidationSkewOk returns a tuple with the IssuedAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtValidationSkew

`func (o *JWTValidation) SetIssuedAtValidationSkew(v int32)`

SetIssuedAtValidationSkew sets IssuedAtValidationSkew field to given value.

### HasIssuedAtValidationSkew

`func (o *JWTValidation) HasIssuedAtValidationSkew() bool`

HasIssuedAtValidationSkew returns a boolean if a field has been set.

### GetNotBeforeValidationSkew

`func (o *JWTValidation) GetNotBeforeValidationSkew() int32`

GetNotBeforeValidationSkew returns the NotBeforeValidationSkew field if non-nil, zero value otherwise.

### GetNotBeforeValidationSkewOk

`func (o *JWTValidation) GetNotBeforeValidationSkewOk() (*int32, bool)`

GetNotBeforeValidationSkewOk returns a tuple with the NotBeforeValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeValidationSkew

`func (o *JWTValidation) SetNotBeforeValidationSkew(v int32)`

SetNotBeforeValidationSkew sets NotBeforeValidationSkew field to given value.

### HasNotBeforeValidationSkew

`func (o *JWTValidation) HasNotBeforeValidationSkew() bool`

HasNotBeforeValidationSkew returns a boolean if a field has been set.

### GetSigningMethod

`func (o *JWTValidation) GetSigningMethod() string`

GetSigningMethod returns the SigningMethod field if non-nil, zero value otherwise.

### GetSigningMethodOk

`func (o *JWTValidation) GetSigningMethodOk() (*string, bool)`

GetSigningMethodOk returns a tuple with the SigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningMethod

`func (o *JWTValidation) SetSigningMethod(v string)`

SetSigningMethod sets SigningMethod field to given value.

### HasSigningMethod

`func (o *JWTValidation) HasSigningMethod() bool`

HasSigningMethod returns a boolean if a field has been set.

### GetSource

`func (o *JWTValidation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JWTValidation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JWTValidation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *JWTValidation) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


