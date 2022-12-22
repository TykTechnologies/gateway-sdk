/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gate

import (
	"encoding/json"
)

// SessionStateMonitor struct for SessionStateMonitor
type SessionStateMonitor struct {
	TriggerLimits []float64 `json:"trigger_limits,omitempty"`
}

// NewSessionStateMonitor instantiates a new SessionStateMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStateMonitor() *SessionStateMonitor {
	this := SessionStateMonitor{}
	return &this
}

// NewSessionStateMonitorWithDefaults instantiates a new SessionStateMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStateMonitorWithDefaults() *SessionStateMonitor {
	this := SessionStateMonitor{}
	return &this
}

// GetTriggerLimits returns the TriggerLimits field value if set, zero value otherwise.
func (o *SessionStateMonitor) GetTriggerLimits() []float64 {
	if o == nil || o.TriggerLimits == nil {
		var ret []float64
		return ret
	}
	return o.TriggerLimits
}

// GetTriggerLimitsOk returns a tuple with the TriggerLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateMonitor) GetTriggerLimitsOk() ([]float64, bool) {
	if o == nil || o.TriggerLimits == nil {
		return nil, false
	}
	return o.TriggerLimits, true
}

// HasTriggerLimits returns a boolean if a field has been set.
func (o *SessionStateMonitor) HasTriggerLimits() bool {
	if o != nil && o.TriggerLimits != nil {
		return true
	}

	return false
}

// SetTriggerLimits gets a reference to the given []float64 and assigns it to the TriggerLimits field.
func (o *SessionStateMonitor) SetTriggerLimits(v []float64) {
	o.TriggerLimits = v
}

func (o SessionStateMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TriggerLimits != nil {
		toSerialize["trigger_limits"] = o.TriggerLimits
	}
	return json.Marshal(toSerialize)
}

type NullableSessionStateMonitor struct {
	value *SessionStateMonitor
	isSet bool
}

func (v NullableSessionStateMonitor) Get() *SessionStateMonitor {
	return v.value
}

func (v *NullableSessionStateMonitor) Set(val *SessionStateMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStateMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStateMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStateMonitor(val *SessionStateMonitor) *NullableSessionStateMonitor {
	return &NullableSessionStateMonitor{value: val, isSet: true}
}

func (v NullableSessionStateMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStateMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


