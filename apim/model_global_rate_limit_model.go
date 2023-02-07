/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// GlobalRateLimitModel struct for GlobalRateLimitModel
type GlobalRateLimitModel struct {
	Per *float64 `json:"per,omitempty"`
	Rate *float64 `json:"rate,omitempty"`
}

// NewGlobalRateLimitModel instantiates a new GlobalRateLimitModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRateLimitModel() *GlobalRateLimitModel {
	this := GlobalRateLimitModel{}
	return &this
}

// NewGlobalRateLimitModelWithDefaults instantiates a new GlobalRateLimitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRateLimitModelWithDefaults() *GlobalRateLimitModel {
	this := GlobalRateLimitModel{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *GlobalRateLimitModel) GetPer() float64 {
	if o == nil || o.Per == nil {
		var ret float64
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRateLimitModel) GetPerOk() (*float64, bool) {
	if o == nil || o.Per == nil {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *GlobalRateLimitModel) HasPer() bool {
	if o != nil && o.Per != nil {
		return true
	}

	return false
}

// SetPer gets a reference to the given float64 and assigns it to the Per field.
func (o *GlobalRateLimitModel) SetPer(v float64) {
	o.Per = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GlobalRateLimitModel) GetRate() float64 {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRateLimitModel) GetRateOk() (*float64, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GlobalRateLimitModel) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *GlobalRateLimitModel) SetRate(v float64) {
	o.Rate = &v
}

func (o GlobalRateLimitModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Per != nil {
		toSerialize["per"] = o.Per
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalRateLimitModel struct {
	value *GlobalRateLimitModel
	isSet bool
}

func (v NullableGlobalRateLimitModel) Get() *GlobalRateLimitModel {
	return v.value
}

func (v *NullableGlobalRateLimitModel) Set(val *GlobalRateLimitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRateLimitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRateLimitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRateLimitModel(val *GlobalRateLimitModel) *NullableGlobalRateLimitModel {
	return &NullableGlobalRateLimitModel{value: val, isSet: true}
}

func (v NullableGlobalRateLimitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRateLimitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


