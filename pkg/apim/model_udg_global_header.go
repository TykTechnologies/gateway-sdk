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

// checks if the UDGGlobalHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UDGGlobalHeader{}

// UDGGlobalHeader struct for UDGGlobalHeader
type UDGGlobalHeader struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewUDGGlobalHeader instantiates a new UDGGlobalHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUDGGlobalHeader() *UDGGlobalHeader {
	this := UDGGlobalHeader{}
	return &this
}

// NewUDGGlobalHeaderWithDefaults instantiates a new UDGGlobalHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUDGGlobalHeaderWithDefaults() *UDGGlobalHeader {
	this := UDGGlobalHeader{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UDGGlobalHeader) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UDGGlobalHeader) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UDGGlobalHeader) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UDGGlobalHeader) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UDGGlobalHeader) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UDGGlobalHeader) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UDGGlobalHeader) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UDGGlobalHeader) SetValue(v string) {
	o.Value = &v
}

func (o UDGGlobalHeader) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UDGGlobalHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableUDGGlobalHeader struct {
	value *UDGGlobalHeader
	isSet bool
}

func (v NullableUDGGlobalHeader) Get() *UDGGlobalHeader {
	return v.value
}

func (v *NullableUDGGlobalHeader) Set(val *UDGGlobalHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableUDGGlobalHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableUDGGlobalHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUDGGlobalHeader(val *UDGGlobalHeader) *NullableUDGGlobalHeader {
	return &NullableUDGGlobalHeader{value: val, isSet: true}
}

func (v NullableUDGGlobalHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUDGGlobalHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
