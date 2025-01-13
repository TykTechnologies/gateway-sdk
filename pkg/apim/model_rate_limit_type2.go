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

// checks if the RateLimitType2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitType2{}

// RateLimitType2 struct for RateLimitType2
type RateLimitType2 struct {
	Per       *float32            `json:"per,omitempty"`
	Rate      *float32            `json:"rate,omitempty"`
	Smoothing *RateLimitSmoothing `json:"smoothing,omitempty"`
}

// NewRateLimitType2 instantiates a new RateLimitType2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitType2() *RateLimitType2 {
	this := RateLimitType2{}
	return &this
}

// NewRateLimitType2WithDefaults instantiates a new RateLimitType2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitType2WithDefaults() *RateLimitType2 {
	this := RateLimitType2{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *RateLimitType2) GetPer() float32 {
	if o == nil || IsNil(o.Per) {
		var ret float32
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitType2) GetPerOk() (*float32, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *RateLimitType2) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given float32 and assigns it to the Per field.
func (o *RateLimitType2) SetPer(v float32) {
	o.Per = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *RateLimitType2) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitType2) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *RateLimitType2) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *RateLimitType2) SetRate(v float32) {
	o.Rate = &v
}

// GetSmoothing returns the Smoothing field value if set, zero value otherwise.
func (o *RateLimitType2) GetSmoothing() RateLimitSmoothing {
	if o == nil || IsNil(o.Smoothing) {
		var ret RateLimitSmoothing
		return ret
	}
	return *o.Smoothing
}

// GetSmoothingOk returns a tuple with the Smoothing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitType2) GetSmoothingOk() (*RateLimitSmoothing, bool) {
	if o == nil || IsNil(o.Smoothing) {
		return nil, false
	}
	return o.Smoothing, true
}

// HasSmoothing returns a boolean if a field has been set.
func (o *RateLimitType2) HasSmoothing() bool {
	if o != nil && !IsNil(o.Smoothing) {
		return true
	}

	return false
}

// SetSmoothing gets a reference to the given RateLimitSmoothing and assigns it to the Smoothing field.
func (o *RateLimitType2) SetSmoothing(v RateLimitSmoothing) {
	o.Smoothing = &v
}

func (o RateLimitType2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitType2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Smoothing) {
		toSerialize["smoothing"] = o.Smoothing
	}
	return toSerialize, nil
}

type NullableRateLimitType2 struct {
	value *RateLimitType2
	isSet bool
}

func (v NullableRateLimitType2) Get() *RateLimitType2 {
	return v.value
}

func (v *NullableRateLimitType2) Set(val *RateLimitType2) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitType2) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitType2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitType2(val *RateLimitType2) *NullableRateLimitType2 {
	return &NullableRateLimitType2{value: val, isSet: true}
}

func (v NullableRateLimitType2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitType2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
