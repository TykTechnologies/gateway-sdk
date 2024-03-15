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

// checks if the TrackEndpointMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackEndpointMeta{}

// TrackEndpointMeta struct for TrackEndpointMeta
type TrackEndpointMeta struct {
	Disabled *bool `json:"disabled,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewTrackEndpointMeta instantiates a new TrackEndpointMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackEndpointMeta() *TrackEndpointMeta {
	this := TrackEndpointMeta{}
	return &this
}

// NewTrackEndpointMetaWithDefaults instantiates a new TrackEndpointMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackEndpointMetaWithDefaults() *TrackEndpointMeta {
	this := TrackEndpointMeta{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *TrackEndpointMeta) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackEndpointMeta) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *TrackEndpointMeta) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *TrackEndpointMeta) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TrackEndpointMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackEndpointMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TrackEndpointMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *TrackEndpointMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TrackEndpointMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackEndpointMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TrackEndpointMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TrackEndpointMeta) SetPath(v string) {
	o.Path = &v
}

func (o TrackEndpointMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackEndpointMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableTrackEndpointMeta struct {
	value *TrackEndpointMeta
	isSet bool
}

func (v NullableTrackEndpointMeta) Get() *TrackEndpointMeta {
	return v.value
}

func (v *NullableTrackEndpointMeta) Set(val *TrackEndpointMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackEndpointMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackEndpointMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackEndpointMeta(val *TrackEndpointMeta) *NullableTrackEndpointMeta {
	return &NullableTrackEndpointMeta{value: val, isSet: true}
}

func (v NullableTrackEndpointMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackEndpointMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


