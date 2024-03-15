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

// checks if the VersionInfoPaths type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfoPaths{}

// VersionInfoPaths struct for VersionInfoPaths
type VersionInfoPaths struct {
	BlackList []string `json:"black_list,omitempty"`
	Ignored []string `json:"ignored,omitempty"`
	WhiteList []string `json:"white_list,omitempty"`
}

// NewVersionInfoPaths instantiates a new VersionInfoPaths object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfoPaths() *VersionInfoPaths {
	this := VersionInfoPaths{}
	return &this
}

// NewVersionInfoPathsWithDefaults instantiates a new VersionInfoPaths object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoPathsWithDefaults() *VersionInfoPaths {
	this := VersionInfoPaths{}
	return &this
}

// GetBlackList returns the BlackList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfoPaths) GetBlackList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BlackList
}

// GetBlackListOk returns a tuple with the BlackList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfoPaths) GetBlackListOk() ([]string, bool) {
	if o == nil || IsNil(o.BlackList) {
		return nil, false
	}
	return o.BlackList, true
}

// HasBlackList returns a boolean if a field has been set.
func (o *VersionInfoPaths) HasBlackList() bool {
	if o != nil && !IsNil(o.BlackList) {
		return true
	}

	return false
}

// SetBlackList gets a reference to the given []string and assigns it to the BlackList field.
func (o *VersionInfoPaths) SetBlackList(v []string) {
	o.BlackList = v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfoPaths) GetIgnored() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfoPaths) GetIgnoredOk() ([]string, bool) {
	if o == nil || IsNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *VersionInfoPaths) HasIgnored() bool {
	if o != nil && !IsNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given []string and assigns it to the Ignored field.
func (o *VersionInfoPaths) SetIgnored(v []string) {
	o.Ignored = v
}

// GetWhiteList returns the WhiteList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfoPaths) GetWhiteList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WhiteList
}

// GetWhiteListOk returns a tuple with the WhiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfoPaths) GetWhiteListOk() ([]string, bool) {
	if o == nil || IsNil(o.WhiteList) {
		return nil, false
	}
	return o.WhiteList, true
}

// HasWhiteList returns a boolean if a field has been set.
func (o *VersionInfoPaths) HasWhiteList() bool {
	if o != nil && !IsNil(o.WhiteList) {
		return true
	}

	return false
}

// SetWhiteList gets a reference to the given []string and assigns it to the WhiteList field.
func (o *VersionInfoPaths) SetWhiteList(v []string) {
	o.WhiteList = v
}

func (o VersionInfoPaths) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfoPaths) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlackList != nil {
		toSerialize["black_list"] = o.BlackList
	}
	if o.Ignored != nil {
		toSerialize["ignored"] = o.Ignored
	}
	if o.WhiteList != nil {
		toSerialize["white_list"] = o.WhiteList
	}
	return toSerialize, nil
}

type NullableVersionInfoPaths struct {
	value *VersionInfoPaths
	isSet bool
}

func (v NullableVersionInfoPaths) Get() *VersionInfoPaths {
	return v.value
}

func (v *NullableVersionInfoPaths) Set(val *VersionInfoPaths) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfoPaths) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfoPaths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfoPaths(val *VersionInfoPaths) *NullableVersionInfoPaths {
	return &NullableVersionInfoPaths{value: val, isSet: true}
}

func (v NullableVersionInfoPaths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfoPaths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


