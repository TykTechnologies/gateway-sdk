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

// checks if the UptimeTestsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UptimeTestsConfig{}

// UptimeTestsConfig struct for UptimeTestsConfig
type UptimeTestsConfig struct {
	ExpireUtimeAfter *int32                         `json:"expire_utime_after,omitempty"`
	RecheckWait      *int32                         `json:"recheck_wait,omitempty"`
	ServiceDiscovery *ServiceDiscoveryConfiguration `json:"service_discovery,omitempty"`
}

// NewUptimeTestsConfig instantiates a new UptimeTestsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUptimeTestsConfig() *UptimeTestsConfig {
	this := UptimeTestsConfig{}
	return &this
}

// NewUptimeTestsConfigWithDefaults instantiates a new UptimeTestsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUptimeTestsConfigWithDefaults() *UptimeTestsConfig {
	this := UptimeTestsConfig{}
	return &this
}

// GetExpireUtimeAfter returns the ExpireUtimeAfter field value if set, zero value otherwise.
func (o *UptimeTestsConfig) GetExpireUtimeAfter() int32 {
	if o == nil || IsNil(o.ExpireUtimeAfter) {
		var ret int32
		return ret
	}
	return *o.ExpireUtimeAfter
}

// GetExpireUtimeAfterOk returns a tuple with the ExpireUtimeAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UptimeTestsConfig) GetExpireUtimeAfterOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireUtimeAfter) {
		return nil, false
	}
	return o.ExpireUtimeAfter, true
}

// HasExpireUtimeAfter returns a boolean if a field has been set.
func (o *UptimeTestsConfig) HasExpireUtimeAfter() bool {
	if o != nil && !IsNil(o.ExpireUtimeAfter) {
		return true
	}

	return false
}

// SetExpireUtimeAfter gets a reference to the given int32 and assigns it to the ExpireUtimeAfter field.
func (o *UptimeTestsConfig) SetExpireUtimeAfter(v int32) {
	o.ExpireUtimeAfter = &v
}

// GetRecheckWait returns the RecheckWait field value if set, zero value otherwise.
func (o *UptimeTestsConfig) GetRecheckWait() int32 {
	if o == nil || IsNil(o.RecheckWait) {
		var ret int32
		return ret
	}
	return *o.RecheckWait
}

// GetRecheckWaitOk returns a tuple with the RecheckWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UptimeTestsConfig) GetRecheckWaitOk() (*int32, bool) {
	if o == nil || IsNil(o.RecheckWait) {
		return nil, false
	}
	return o.RecheckWait, true
}

// HasRecheckWait returns a boolean if a field has been set.
func (o *UptimeTestsConfig) HasRecheckWait() bool {
	if o != nil && !IsNil(o.RecheckWait) {
		return true
	}

	return false
}

// SetRecheckWait gets a reference to the given int32 and assigns it to the RecheckWait field.
func (o *UptimeTestsConfig) SetRecheckWait(v int32) {
	o.RecheckWait = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise.
func (o *UptimeTestsConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret ServiceDiscoveryConfiguration
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UptimeTestsConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}
	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *UptimeTestsConfig) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given ServiceDiscoveryConfiguration and assigns it to the ServiceDiscovery field.
func (o *UptimeTestsConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration) {
	o.ServiceDiscovery = &v
}

func (o UptimeTestsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UptimeTestsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpireUtimeAfter) {
		toSerialize["expire_utime_after"] = o.ExpireUtimeAfter
	}
	if !IsNil(o.RecheckWait) {
		toSerialize["recheck_wait"] = o.RecheckWait
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["service_discovery"] = o.ServiceDiscovery
	}
	return toSerialize, nil
}

type NullableUptimeTestsConfig struct {
	value *UptimeTestsConfig
	isSet bool
}

func (v NullableUptimeTestsConfig) Get() *UptimeTestsConfig {
	return v.value
}

func (v *NullableUptimeTestsConfig) Set(val *UptimeTestsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUptimeTestsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUptimeTestsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUptimeTestsConfig(val *UptimeTestsConfig) *NullableUptimeTestsConfig {
	return &NullableUptimeTestsConfig{value: val, isSet: true}
}

func (v NullableUptimeTestsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUptimeTestsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
