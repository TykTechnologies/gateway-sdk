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

// checks if the Introspection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Introspection{}

// Introspection struct for Introspection
type Introspection struct {
	Cache             *IntrospectionCache `json:"cache,omitempty"`
	ClientId          *string             `json:"client_id,omitempty"`
	ClientSecret      *string             `json:"client_secret,omitempty"`
	Enabled           *bool               `json:"enabled,omitempty"`
	IdentityBaseField *string             `json:"identity_base_field,omitempty"`
	Url               *string             `json:"url,omitempty"`
}

// NewIntrospection instantiates a new Introspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntrospection() *Introspection {
	this := Introspection{}
	return &this
}

// NewIntrospectionWithDefaults instantiates a new Introspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntrospectionWithDefaults() *Introspection {
	this := Introspection{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *Introspection) GetCache() IntrospectionCache {
	if o == nil || IsNil(o.Cache) {
		var ret IntrospectionCache
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetCacheOk() (*IntrospectionCache, bool) {
	if o == nil || IsNil(o.Cache) {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *Introspection) HasCache() bool {
	if o != nil && !IsNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given IntrospectionCache and assigns it to the Cache field.
func (o *Introspection) SetCache(v IntrospectionCache) {
	o.Cache = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Introspection) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Introspection) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Introspection) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *Introspection) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *Introspection) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *Introspection) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Introspection) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Introspection) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Introspection) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIdentityBaseField returns the IdentityBaseField field value if set, zero value otherwise.
func (o *Introspection) GetIdentityBaseField() string {
	if o == nil || IsNil(o.IdentityBaseField) {
		var ret string
		return ret
	}
	return *o.IdentityBaseField
}

// GetIdentityBaseFieldOk returns a tuple with the IdentityBaseField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetIdentityBaseFieldOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityBaseField) {
		return nil, false
	}
	return o.IdentityBaseField, true
}

// HasIdentityBaseField returns a boolean if a field has been set.
func (o *Introspection) HasIdentityBaseField() bool {
	if o != nil && !IsNil(o.IdentityBaseField) {
		return true
	}

	return false
}

// SetIdentityBaseField gets a reference to the given string and assigns it to the IdentityBaseField field.
func (o *Introspection) SetIdentityBaseField(v string) {
	o.IdentityBaseField = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Introspection) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Introspection) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Introspection) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Introspection) SetUrl(v string) {
	o.Url = &v
}

func (o Introspection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Introspection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cache) {
		toSerialize["cache"] = o.Cache
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IdentityBaseField) {
		toSerialize["identity_base_field"] = o.IdentityBaseField
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableIntrospection struct {
	value *Introspection
	isSet bool
}

func (v NullableIntrospection) Get() *Introspection {
	return v.value
}

func (v *NullableIntrospection) Set(val *Introspection) {
	v.value = val
	v.isSet = true
}

func (v NullableIntrospection) IsSet() bool {
	return v.isSet
}

func (v *NullableIntrospection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntrospection(val *Introspection) *NullableIntrospection {
	return &NullableIntrospection{value: val, isSet: true}
}

func (v NullableIntrospection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntrospection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
