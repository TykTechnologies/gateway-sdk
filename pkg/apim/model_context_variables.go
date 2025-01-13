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

// checks if the ContextVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextVariables{}

// ContextVariables struct for ContextVariables
type ContextVariables struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewContextVariables instantiates a new ContextVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextVariables() *ContextVariables {
	this := ContextVariables{}
	return &this
}

// NewContextVariablesWithDefaults instantiates a new ContextVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextVariablesWithDefaults() *ContextVariables {
	this := ContextVariables{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ContextVariables) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextVariables) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ContextVariables) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ContextVariables) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ContextVariables) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableContextVariables struct {
	value *ContextVariables
	isSet bool
}

func (v NullableContextVariables) Get() *ContextVariables {
	return v.value
}

func (v *NullableContextVariables) Set(val *ContextVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableContextVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableContextVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextVariables(val *ContextVariables) *NullableContextVariables {
	return &NullableContextVariables{value: val, isSet: true}
}

func (v NullableContextVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
