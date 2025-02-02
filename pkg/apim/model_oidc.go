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

// checks if the OIDC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OIDC{}

// OIDC struct for OIDC
type OIDC struct {
	AuthSources         *AuthSources    `json:"AuthSources,omitempty"`
	Enabled             *bool           `json:"enabled,omitempty"`
	Providers           []ProviderType2 `json:"providers,omitempty"`
	Scopes              *ScopesType2    `json:"scopes,omitempty"`
	SegregateByClientId *bool           `json:"segregateByClientId,omitempty"`
}

// NewOIDC instantiates a new OIDC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDC() *OIDC {
	this := OIDC{}
	return &this
}

// NewOIDCWithDefaults instantiates a new OIDC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWithDefaults() *OIDC {
	this := OIDC{}
	return &this
}

// GetAuthSources returns the AuthSources field value if set, zero value otherwise.
func (o *OIDC) GetAuthSources() AuthSources {
	if o == nil || IsNil(o.AuthSources) {
		var ret AuthSources
		return ret
	}
	return *o.AuthSources
}

// GetAuthSourcesOk returns a tuple with the AuthSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDC) GetAuthSourcesOk() (*AuthSources, bool) {
	if o == nil || IsNil(o.AuthSources) {
		return nil, false
	}
	return o.AuthSources, true
}

// HasAuthSources returns a boolean if a field has been set.
func (o *OIDC) HasAuthSources() bool {
	if o != nil && !IsNil(o.AuthSources) {
		return true
	}

	return false
}

// SetAuthSources gets a reference to the given AuthSources and assigns it to the AuthSources field.
func (o *OIDC) SetAuthSources(v AuthSources) {
	o.AuthSources = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OIDC) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDC) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OIDC) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OIDC) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *OIDC) GetProviders() []ProviderType2 {
	if o == nil || IsNil(o.Providers) {
		var ret []ProviderType2
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDC) GetProvidersOk() ([]ProviderType2, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *OIDC) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []ProviderType2 and assigns it to the Providers field.
func (o *OIDC) SetProviders(v []ProviderType2) {
	o.Providers = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OIDC) GetScopes() ScopesType2 {
	if o == nil || IsNil(o.Scopes) {
		var ret ScopesType2
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDC) GetScopesOk() (*ScopesType2, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OIDC) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given ScopesType2 and assigns it to the Scopes field.
func (o *OIDC) SetScopes(v ScopesType2) {
	o.Scopes = &v
}

// GetSegregateByClientId returns the SegregateByClientId field value if set, zero value otherwise.
func (o *OIDC) GetSegregateByClientId() bool {
	if o == nil || IsNil(o.SegregateByClientId) {
		var ret bool
		return ret
	}
	return *o.SegregateByClientId
}

// GetSegregateByClientIdOk returns a tuple with the SegregateByClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDC) GetSegregateByClientIdOk() (*bool, bool) {
	if o == nil || IsNil(o.SegregateByClientId) {
		return nil, false
	}
	return o.SegregateByClientId, true
}

// HasSegregateByClientId returns a boolean if a field has been set.
func (o *OIDC) HasSegregateByClientId() bool {
	if o != nil && !IsNil(o.SegregateByClientId) {
		return true
	}

	return false
}

// SetSegregateByClientId gets a reference to the given bool and assigns it to the SegregateByClientId field.
func (o *OIDC) SetSegregateByClientId(v bool) {
	o.SegregateByClientId = &v
}

func (o OIDC) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OIDC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthSources) {
		toSerialize["AuthSources"] = o.AuthSources
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SegregateByClientId) {
		toSerialize["segregateByClientId"] = o.SegregateByClientId
	}
	return toSerialize, nil
}

type NullableOIDC struct {
	value *OIDC
	isSet bool
}

func (v NullableOIDC) Get() *OIDC {
	return v.value
}

func (v *NullableOIDC) Set(val *OIDC) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDC) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDC(val *OIDC) *NullableOIDC {
	return &NullableOIDC{value: val, isSet: true}
}

func (v NullableOIDC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
