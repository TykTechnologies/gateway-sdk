/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation). * Managing and listing policies. * Managing and listing API Definitions (only when not using the Tyk Dashboard). * Hot reloads / reloading a cluster configuration. * OAuth client creation (only when not using the Tyk Dashboard).  In order to use the Gateway API, you'll need to set the **secret** parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  **x-tyk-authorization: <your-secret>*** <br/>  <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the Scopes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scopes{}

// Scopes struct for Scopes
type Scopes struct {
	Jwt  *ScopeClaim `json:"jwt,omitempty"`
	Oidc *ScopeClaim `json:"oidc,omitempty"`
}

// NewScopes instantiates a new Scopes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopes() *Scopes {
	this := Scopes{}
	return &this
}

// NewScopesWithDefaults instantiates a new Scopes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopesWithDefaults() *Scopes {
	this := Scopes{}
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *Scopes) GetJwt() ScopeClaim {
	if o == nil || IsNil(o.Jwt) {
		var ret ScopeClaim
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scopes) GetJwtOk() (*ScopeClaim, bool) {
	if o == nil || IsNil(o.Jwt) {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *Scopes) HasJwt() bool {
	if o != nil && !IsNil(o.Jwt) {
		return true
	}

	return false
}

// SetJwt gets a reference to the given ScopeClaim and assigns it to the Jwt field.
func (o *Scopes) SetJwt(v ScopeClaim) {
	o.Jwt = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *Scopes) GetOidc() ScopeClaim {
	if o == nil || IsNil(o.Oidc) {
		var ret ScopeClaim
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scopes) GetOidcOk() (*ScopeClaim, bool) {
	if o == nil || IsNil(o.Oidc) {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *Scopes) HasOidc() bool {
	if o != nil && !IsNil(o.Oidc) {
		return true
	}

	return false
}

// SetOidc gets a reference to the given ScopeClaim and assigns it to the Oidc field.
func (o *Scopes) SetOidc(v ScopeClaim) {
	o.Oidc = &v
}

func (o Scopes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scopes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jwt) {
		toSerialize["jwt"] = o.Jwt
	}
	if !IsNil(o.Oidc) {
		toSerialize["oidc"] = o.Oidc
	}
	return toSerialize, nil
}

type NullableScopes struct {
	value *Scopes
	isSet bool
}

func (v NullableScopes) Get() *Scopes {
	return v.value
}

func (v *NullableScopes) Set(val *Scopes) {
	v.value = val
	v.isSet = true
}

func (v NullableScopes) IsSet() bool {
	return v.isSet
}

func (v *NullableScopes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopes(val *Scopes) *NullableScopes {
	return &NullableScopes{value: val, isSet: true}
}

func (v NullableScopes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
