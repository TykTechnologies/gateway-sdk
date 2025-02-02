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

// checks if the CircuitBreakerMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitBreakerMeta{}

// CircuitBreakerMeta struct for CircuitBreakerMeta
type CircuitBreakerMeta struct {
	DisableHalfOpenState *bool    `json:"disable_half_open_state,omitempty"`
	Disabled             *bool    `json:"disabled,omitempty"`
	Method               *string  `json:"method,omitempty"`
	Path                 *string  `json:"path,omitempty"`
	ReturnToServiceAfter *int32   `json:"return_to_service_after,omitempty"`
	Samples              *int64   `json:"samples,omitempty"`
	ThresholdPercent     *float32 `json:"threshold_percent,omitempty"`
}

// NewCircuitBreakerMeta instantiates a new CircuitBreakerMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitBreakerMeta() *CircuitBreakerMeta {
	this := CircuitBreakerMeta{}
	return &this
}

// NewCircuitBreakerMetaWithDefaults instantiates a new CircuitBreakerMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitBreakerMetaWithDefaults() *CircuitBreakerMeta {
	this := CircuitBreakerMeta{}
	return &this
}

// GetDisableHalfOpenState returns the DisableHalfOpenState field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetDisableHalfOpenState() bool {
	if o == nil || IsNil(o.DisableHalfOpenState) {
		var ret bool
		return ret
	}
	return *o.DisableHalfOpenState
}

// GetDisableHalfOpenStateOk returns a tuple with the DisableHalfOpenState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetDisableHalfOpenStateOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableHalfOpenState) {
		return nil, false
	}
	return o.DisableHalfOpenState, true
}

// HasDisableHalfOpenState returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasDisableHalfOpenState() bool {
	if o != nil && !IsNil(o.DisableHalfOpenState) {
		return true
	}

	return false
}

// SetDisableHalfOpenState gets a reference to the given bool and assigns it to the DisableHalfOpenState field.
func (o *CircuitBreakerMeta) SetDisableHalfOpenState(v bool) {
	o.DisableHalfOpenState = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *CircuitBreakerMeta) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CircuitBreakerMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CircuitBreakerMeta) SetPath(v string) {
	o.Path = &v
}

// GetReturnToServiceAfter returns the ReturnToServiceAfter field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetReturnToServiceAfter() int32 {
	if o == nil || IsNil(o.ReturnToServiceAfter) {
		var ret int32
		return ret
	}
	return *o.ReturnToServiceAfter
}

// GetReturnToServiceAfterOk returns a tuple with the ReturnToServiceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetReturnToServiceAfterOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnToServiceAfter) {
		return nil, false
	}
	return o.ReturnToServiceAfter, true
}

// HasReturnToServiceAfter returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasReturnToServiceAfter() bool {
	if o != nil && !IsNil(o.ReturnToServiceAfter) {
		return true
	}

	return false
}

// SetReturnToServiceAfter gets a reference to the given int32 and assigns it to the ReturnToServiceAfter field.
func (o *CircuitBreakerMeta) SetReturnToServiceAfter(v int32) {
	o.ReturnToServiceAfter = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetSamples() int64 {
	if o == nil || IsNil(o.Samples) {
		var ret int64
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetSamplesOk() (*int64, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given int64 and assigns it to the Samples field.
func (o *CircuitBreakerMeta) SetSamples(v int64) {
	o.Samples = &v
}

// GetThresholdPercent returns the ThresholdPercent field value if set, zero value otherwise.
func (o *CircuitBreakerMeta) GetThresholdPercent() float32 {
	if o == nil || IsNil(o.ThresholdPercent) {
		var ret float32
		return ret
	}
	return *o.ThresholdPercent
}

// GetThresholdPercentOk returns a tuple with the ThresholdPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitBreakerMeta) GetThresholdPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.ThresholdPercent) {
		return nil, false
	}
	return o.ThresholdPercent, true
}

// HasThresholdPercent returns a boolean if a field has been set.
func (o *CircuitBreakerMeta) HasThresholdPercent() bool {
	if o != nil && !IsNil(o.ThresholdPercent) {
		return true
	}

	return false
}

// SetThresholdPercent gets a reference to the given float32 and assigns it to the ThresholdPercent field.
func (o *CircuitBreakerMeta) SetThresholdPercent(v float32) {
	o.ThresholdPercent = &v
}

func (o CircuitBreakerMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitBreakerMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableHalfOpenState) {
		toSerialize["disable_half_open_state"] = o.DisableHalfOpenState
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ReturnToServiceAfter) {
		toSerialize["return_to_service_after"] = o.ReturnToServiceAfter
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	if !IsNil(o.ThresholdPercent) {
		toSerialize["threshold_percent"] = o.ThresholdPercent
	}
	return toSerialize, nil
}

type NullableCircuitBreakerMeta struct {
	value *CircuitBreakerMeta
	isSet bool
}

func (v NullableCircuitBreakerMeta) Get() *CircuitBreakerMeta {
	return v.value
}

func (v *NullableCircuitBreakerMeta) Set(val *CircuitBreakerMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitBreakerMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitBreakerMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitBreakerMeta(val *CircuitBreakerMeta) *NullableCircuitBreakerMeta {
	return &NullableCircuitBreakerMeta{value: val, isSet: true}
}

func (v NullableCircuitBreakerMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitBreakerMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
