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

// checks if the ApiAllKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAllKeys{}

// ApiAllKeys struct for ApiAllKeys
type ApiAllKeys struct {
	Keys []string `json:"keys,omitempty"`
}

// NewApiAllKeys instantiates a new ApiAllKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAllKeys() *ApiAllKeys {
	this := ApiAllKeys{}
	return &this
}

// NewApiAllKeysWithDefaults instantiates a new ApiAllKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAllKeysWithDefaults() *ApiAllKeys {
	this := ApiAllKeys{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAllKeys) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAllKeys) GetKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *ApiAllKeys) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *ApiAllKeys) SetKeys(v []string) {
	o.Keys = v
}

func (o ApiAllKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAllKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return toSerialize, nil
}

type NullableApiAllKeys struct {
	value *ApiAllKeys
	isSet bool
}

func (v NullableApiAllKeys) Get() *ApiAllKeys {
	return v.value
}

func (v *NullableApiAllKeys) Set(val *ApiAllKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAllKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAllKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAllKeys(val *ApiAllKeys) *NullableApiAllKeys {
	return &NullableApiAllKeys{value: val, isSet: true}
}

func (v NullableApiAllKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAllKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
