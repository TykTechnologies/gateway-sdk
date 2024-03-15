/*
Tyk Gateway API

 The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway

API version: 5.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the JWTValidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JWTValidation{}

// JWTValidation struct for JWTValidation
type JWTValidation struct {
	Enabled *bool `json:"enabled,omitempty"`
	ExpiresAtValidationSkew *int32 `json:"expires_at_validation_skew,omitempty"`
	IdentityBaseField *string `json:"identity_base_field,omitempty"`
	IssuedAtValidationSkew *int32 `json:"issued_at_validation_skew,omitempty"`
	NotBeforeValidationSkew *int32 `json:"not_before_validation_skew,omitempty"`
	SigningMethod *string `json:"signing_method,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewJWTValidation instantiates a new JWTValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWTValidation() *JWTValidation {
	this := JWTValidation{}
	return &this
}

// NewJWTValidationWithDefaults instantiates a new JWTValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTValidationWithDefaults() *JWTValidation {
	this := JWTValidation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JWTValidation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JWTValidation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JWTValidation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpiresAtValidationSkew returns the ExpiresAtValidationSkew field value if set, zero value otherwise.
func (o *JWTValidation) GetExpiresAtValidationSkew() int32 {
	if o == nil || IsNil(o.ExpiresAtValidationSkew) {
		var ret int32
		return ret
	}
	return *o.ExpiresAtValidationSkew
}

// GetExpiresAtValidationSkewOk returns a tuple with the ExpiresAtValidationSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetExpiresAtValidationSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAtValidationSkew) {
		return nil, false
	}
	return o.ExpiresAtValidationSkew, true
}

// HasExpiresAtValidationSkew returns a boolean if a field has been set.
func (o *JWTValidation) HasExpiresAtValidationSkew() bool {
	if o != nil && !IsNil(o.ExpiresAtValidationSkew) {
		return true
	}

	return false
}

// SetExpiresAtValidationSkew gets a reference to the given int32 and assigns it to the ExpiresAtValidationSkew field.
func (o *JWTValidation) SetExpiresAtValidationSkew(v int32) {
	o.ExpiresAtValidationSkew = &v
}

// GetIdentityBaseField returns the IdentityBaseField field value if set, zero value otherwise.
func (o *JWTValidation) GetIdentityBaseField() string {
	if o == nil || IsNil(o.IdentityBaseField) {
		var ret string
		return ret
	}
	return *o.IdentityBaseField
}

// GetIdentityBaseFieldOk returns a tuple with the IdentityBaseField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetIdentityBaseFieldOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityBaseField) {
		return nil, false
	}
	return o.IdentityBaseField, true
}

// HasIdentityBaseField returns a boolean if a field has been set.
func (o *JWTValidation) HasIdentityBaseField() bool {
	if o != nil && !IsNil(o.IdentityBaseField) {
		return true
	}

	return false
}

// SetIdentityBaseField gets a reference to the given string and assigns it to the IdentityBaseField field.
func (o *JWTValidation) SetIdentityBaseField(v string) {
	o.IdentityBaseField = &v
}

// GetIssuedAtValidationSkew returns the IssuedAtValidationSkew field value if set, zero value otherwise.
func (o *JWTValidation) GetIssuedAtValidationSkew() int32 {
	if o == nil || IsNil(o.IssuedAtValidationSkew) {
		var ret int32
		return ret
	}
	return *o.IssuedAtValidationSkew
}

// GetIssuedAtValidationSkewOk returns a tuple with the IssuedAtValidationSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetIssuedAtValidationSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.IssuedAtValidationSkew) {
		return nil, false
	}
	return o.IssuedAtValidationSkew, true
}

// HasIssuedAtValidationSkew returns a boolean if a field has been set.
func (o *JWTValidation) HasIssuedAtValidationSkew() bool {
	if o != nil && !IsNil(o.IssuedAtValidationSkew) {
		return true
	}

	return false
}

// SetIssuedAtValidationSkew gets a reference to the given int32 and assigns it to the IssuedAtValidationSkew field.
func (o *JWTValidation) SetIssuedAtValidationSkew(v int32) {
	o.IssuedAtValidationSkew = &v
}

// GetNotBeforeValidationSkew returns the NotBeforeValidationSkew field value if set, zero value otherwise.
func (o *JWTValidation) GetNotBeforeValidationSkew() int32 {
	if o == nil || IsNil(o.NotBeforeValidationSkew) {
		var ret int32
		return ret
	}
	return *o.NotBeforeValidationSkew
}

// GetNotBeforeValidationSkewOk returns a tuple with the NotBeforeValidationSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetNotBeforeValidationSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.NotBeforeValidationSkew) {
		return nil, false
	}
	return o.NotBeforeValidationSkew, true
}

// HasNotBeforeValidationSkew returns a boolean if a field has been set.
func (o *JWTValidation) HasNotBeforeValidationSkew() bool {
	if o != nil && !IsNil(o.NotBeforeValidationSkew) {
		return true
	}

	return false
}

// SetNotBeforeValidationSkew gets a reference to the given int32 and assigns it to the NotBeforeValidationSkew field.
func (o *JWTValidation) SetNotBeforeValidationSkew(v int32) {
	o.NotBeforeValidationSkew = &v
}

// GetSigningMethod returns the SigningMethod field value if set, zero value otherwise.
func (o *JWTValidation) GetSigningMethod() string {
	if o == nil || IsNil(o.SigningMethod) {
		var ret string
		return ret
	}
	return *o.SigningMethod
}

// GetSigningMethodOk returns a tuple with the SigningMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetSigningMethodOk() (*string, bool) {
	if o == nil || IsNil(o.SigningMethod) {
		return nil, false
	}
	return o.SigningMethod, true
}

// HasSigningMethod returns a boolean if a field has been set.
func (o *JWTValidation) HasSigningMethod() bool {
	if o != nil && !IsNil(o.SigningMethod) {
		return true
	}

	return false
}

// SetSigningMethod gets a reference to the given string and assigns it to the SigningMethod field.
func (o *JWTValidation) SetSigningMethod(v string) {
	o.SigningMethod = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JWTValidation) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWTValidation) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JWTValidation) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *JWTValidation) SetSource(v string) {
	o.Source = &v
}

func (o JWTValidation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JWTValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ExpiresAtValidationSkew) {
		toSerialize["expires_at_validation_skew"] = o.ExpiresAtValidationSkew
	}
	if !IsNil(o.IdentityBaseField) {
		toSerialize["identity_base_field"] = o.IdentityBaseField
	}
	if !IsNil(o.IssuedAtValidationSkew) {
		toSerialize["issued_at_validation_skew"] = o.IssuedAtValidationSkew
	}
	if !IsNil(o.NotBeforeValidationSkew) {
		toSerialize["not_before_validation_skew"] = o.NotBeforeValidationSkew
	}
	if !IsNil(o.SigningMethod) {
		toSerialize["signing_method"] = o.SigningMethod
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableJWTValidation struct {
	value *JWTValidation
	isSet bool
}

func (v NullableJWTValidation) Get() *JWTValidation {
	return v.value
}

func (v *NullableJWTValidation) Set(val *JWTValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableJWTValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableJWTValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWTValidation(val *JWTValidation) *NullableJWTValidation {
	return &NullableJWTValidation{value: val, isSet: true}
}

func (v NullableJWTValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWTValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


