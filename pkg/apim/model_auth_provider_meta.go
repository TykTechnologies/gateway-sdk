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

// checks if the AuthProviderMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthProviderMeta{}

// AuthProviderMeta struct for AuthProviderMeta
type AuthProviderMeta struct {
	Meta          map[string]interface{} `json:"meta,omitempty"`
	Name          *string                `json:"name,omitempty"`
	StorageEngine *string                `json:"storage_engine,omitempty"`
}

// NewAuthProviderMeta instantiates a new AuthProviderMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthProviderMeta() *AuthProviderMeta {
	this := AuthProviderMeta{}
	return &this
}

// NewAuthProviderMetaWithDefaults instantiates a new AuthProviderMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthProviderMetaWithDefaults() *AuthProviderMeta {
	this := AuthProviderMeta{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthProviderMeta) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthProviderMeta) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AuthProviderMeta) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *AuthProviderMeta) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthProviderMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProviderMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthProviderMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthProviderMeta) SetName(v string) {
	o.Name = &v
}

// GetStorageEngine returns the StorageEngine field value if set, zero value otherwise.
func (o *AuthProviderMeta) GetStorageEngine() string {
	if o == nil || IsNil(o.StorageEngine) {
		var ret string
		return ret
	}
	return *o.StorageEngine
}

// GetStorageEngineOk returns a tuple with the StorageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProviderMeta) GetStorageEngineOk() (*string, bool) {
	if o == nil || IsNil(o.StorageEngine) {
		return nil, false
	}
	return o.StorageEngine, true
}

// HasStorageEngine returns a boolean if a field has been set.
func (o *AuthProviderMeta) HasStorageEngine() bool {
	if o != nil && !IsNil(o.StorageEngine) {
		return true
	}

	return false
}

// SetStorageEngine gets a reference to the given string and assigns it to the StorageEngine field.
func (o *AuthProviderMeta) SetStorageEngine(v string) {
	o.StorageEngine = &v
}

func (o AuthProviderMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthProviderMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StorageEngine) {
		toSerialize["storage_engine"] = o.StorageEngine
	}
	return toSerialize, nil
}

type NullableAuthProviderMeta struct {
	value *AuthProviderMeta
	isSet bool
}

func (v NullableAuthProviderMeta) Get() *AuthProviderMeta {
	return v.value
}

func (v *NullableAuthProviderMeta) Set(val *AuthProviderMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProviderMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProviderMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProviderMeta(val *AuthProviderMeta) *NullableAuthProviderMeta {
	return &NullableAuthProviderMeta{value: val, isSet: true}
}

func (v NullableAuthProviderMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProviderMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
