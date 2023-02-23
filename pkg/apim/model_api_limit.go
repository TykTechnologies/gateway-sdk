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

// checks if the APILimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APILimit{}

// APILimit APILimit stores quota and rate limit on ACL level (per API)
type APILimit struct {
	Per *float64 `json:"per,omitempty"`
	QuotaMax *int64 `json:"quota_max,omitempty"`
	QuotaRemaining *int64 `json:"quota_remaining,omitempty"`
	QuotaRenewalRate *int64 `json:"quota_renewal_rate,omitempty"`
	QuotaRenews *int64 `json:"quota_renews,omitempty"`
	Rate *float64 `json:"rate,omitempty"`
	SetByPolicy *bool `json:"set_by_policy,omitempty"`
	ThrottleInterval *float64 `json:"throttle_interval,omitempty"`
	ThrottleRetryLimit *int64 `json:"throttle_retry_limit,omitempty"`
}

// NewAPILimit instantiates a new APILimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPILimit() *APILimit {
	this := APILimit{}
	return &this
}

// NewAPILimitWithDefaults instantiates a new APILimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPILimitWithDefaults() *APILimit {
	this := APILimit{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *APILimit) GetPer() float64 {
	if o == nil || IsNil(o.Per) {
		var ret float64
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetPerOk() (*float64, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *APILimit) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given float64 and assigns it to the Per field.
func (o *APILimit) SetPer(v float64) {
	o.Per = &v
}

// GetQuotaMax returns the QuotaMax field value if set, zero value otherwise.
func (o *APILimit) GetQuotaMax() int64 {
	if o == nil || IsNil(o.QuotaMax) {
		var ret int64
		return ret
	}
	return *o.QuotaMax
}

// GetQuotaMaxOk returns a tuple with the QuotaMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetQuotaMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaMax) {
		return nil, false
	}
	return o.QuotaMax, true
}

// HasQuotaMax returns a boolean if a field has been set.
func (o *APILimit) HasQuotaMax() bool {
	if o != nil && !IsNil(o.QuotaMax) {
		return true
	}

	return false
}

// SetQuotaMax gets a reference to the given int64 and assigns it to the QuotaMax field.
func (o *APILimit) SetQuotaMax(v int64) {
	o.QuotaMax = &v
}

// GetQuotaRemaining returns the QuotaRemaining field value if set, zero value otherwise.
func (o *APILimit) GetQuotaRemaining() int64 {
	if o == nil || IsNil(o.QuotaRemaining) {
		var ret int64
		return ret
	}
	return *o.QuotaRemaining
}

// GetQuotaRemainingOk returns a tuple with the QuotaRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetQuotaRemainingOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaRemaining) {
		return nil, false
	}
	return o.QuotaRemaining, true
}

// HasQuotaRemaining returns a boolean if a field has been set.
func (o *APILimit) HasQuotaRemaining() bool {
	if o != nil && !IsNil(o.QuotaRemaining) {
		return true
	}

	return false
}

// SetQuotaRemaining gets a reference to the given int64 and assigns it to the QuotaRemaining field.
func (o *APILimit) SetQuotaRemaining(v int64) {
	o.QuotaRemaining = &v
}

// GetQuotaRenewalRate returns the QuotaRenewalRate field value if set, zero value otherwise.
func (o *APILimit) GetQuotaRenewalRate() int64 {
	if o == nil || IsNil(o.QuotaRenewalRate) {
		var ret int64
		return ret
	}
	return *o.QuotaRenewalRate
}

// GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetQuotaRenewalRateOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaRenewalRate) {
		return nil, false
	}
	return o.QuotaRenewalRate, true
}

// HasQuotaRenewalRate returns a boolean if a field has been set.
func (o *APILimit) HasQuotaRenewalRate() bool {
	if o != nil && !IsNil(o.QuotaRenewalRate) {
		return true
	}

	return false
}

// SetQuotaRenewalRate gets a reference to the given int64 and assigns it to the QuotaRenewalRate field.
func (o *APILimit) SetQuotaRenewalRate(v int64) {
	o.QuotaRenewalRate = &v
}

// GetQuotaRenews returns the QuotaRenews field value if set, zero value otherwise.
func (o *APILimit) GetQuotaRenews() int64 {
	if o == nil || IsNil(o.QuotaRenews) {
		var ret int64
		return ret
	}
	return *o.QuotaRenews
}

// GetQuotaRenewsOk returns a tuple with the QuotaRenews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetQuotaRenewsOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaRenews) {
		return nil, false
	}
	return o.QuotaRenews, true
}

// HasQuotaRenews returns a boolean if a field has been set.
func (o *APILimit) HasQuotaRenews() bool {
	if o != nil && !IsNil(o.QuotaRenews) {
		return true
	}

	return false
}

// SetQuotaRenews gets a reference to the given int64 and assigns it to the QuotaRenews field.
func (o *APILimit) SetQuotaRenews(v int64) {
	o.QuotaRenews = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *APILimit) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *APILimit) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *APILimit) SetRate(v float64) {
	o.Rate = &v
}

// GetSetByPolicy returns the SetByPolicy field value if set, zero value otherwise.
func (o *APILimit) GetSetByPolicy() bool {
	if o == nil || IsNil(o.SetByPolicy) {
		var ret bool
		return ret
	}
	return *o.SetByPolicy
}

// GetSetByPolicyOk returns a tuple with the SetByPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetSetByPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.SetByPolicy) {
		return nil, false
	}
	return o.SetByPolicy, true
}

// HasSetByPolicy returns a boolean if a field has been set.
func (o *APILimit) HasSetByPolicy() bool {
	if o != nil && !IsNil(o.SetByPolicy) {
		return true
	}

	return false
}

// SetSetByPolicy gets a reference to the given bool and assigns it to the SetByPolicy field.
func (o *APILimit) SetSetByPolicy(v bool) {
	o.SetByPolicy = &v
}

// GetThrottleInterval returns the ThrottleInterval field value if set, zero value otherwise.
func (o *APILimit) GetThrottleInterval() float64 {
	if o == nil || IsNil(o.ThrottleInterval) {
		var ret float64
		return ret
	}
	return *o.ThrottleInterval
}

// GetThrottleIntervalOk returns a tuple with the ThrottleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetThrottleIntervalOk() (*float64, bool) {
	if o == nil || IsNil(o.ThrottleInterval) {
		return nil, false
	}
	return o.ThrottleInterval, true
}

// HasThrottleInterval returns a boolean if a field has been set.
func (o *APILimit) HasThrottleInterval() bool {
	if o != nil && !IsNil(o.ThrottleInterval) {
		return true
	}

	return false
}

// SetThrottleInterval gets a reference to the given float64 and assigns it to the ThrottleInterval field.
func (o *APILimit) SetThrottleInterval(v float64) {
	o.ThrottleInterval = &v
}

// GetThrottleRetryLimit returns the ThrottleRetryLimit field value if set, zero value otherwise.
func (o *APILimit) GetThrottleRetryLimit() int64 {
	if o == nil || IsNil(o.ThrottleRetryLimit) {
		var ret int64
		return ret
	}
	return *o.ThrottleRetryLimit
}

// GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILimit) GetThrottleRetryLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.ThrottleRetryLimit) {
		return nil, false
	}
	return o.ThrottleRetryLimit, true
}

// HasThrottleRetryLimit returns a boolean if a field has been set.
func (o *APILimit) HasThrottleRetryLimit() bool {
	if o != nil && !IsNil(o.ThrottleRetryLimit) {
		return true
	}

	return false
}

// SetThrottleRetryLimit gets a reference to the given int64 and assigns it to the ThrottleRetryLimit field.
func (o *APILimit) SetThrottleRetryLimit(v int64) {
	o.ThrottleRetryLimit = &v
}

func (o APILimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APILimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.QuotaMax) {
		toSerialize["quota_max"] = o.QuotaMax
	}
	if !IsNil(o.QuotaRemaining) {
		toSerialize["quota_remaining"] = o.QuotaRemaining
	}
	if !IsNil(o.QuotaRenewalRate) {
		toSerialize["quota_renewal_rate"] = o.QuotaRenewalRate
	}
	if !IsNil(o.QuotaRenews) {
		toSerialize["quota_renews"] = o.QuotaRenews
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.SetByPolicy) {
		toSerialize["set_by_policy"] = o.SetByPolicy
	}
	if !IsNil(o.ThrottleInterval) {
		toSerialize["throttle_interval"] = o.ThrottleInterval
	}
	if !IsNil(o.ThrottleRetryLimit) {
		toSerialize["throttle_retry_limit"] = o.ThrottleRetryLimit
	}
	return toSerialize, nil
}

type NullableAPILimit struct {
	value *APILimit
	isSet bool
}

func (v NullableAPILimit) Get() *APILimit {
	return v.value
}

func (v *NullableAPILimit) Set(val *APILimit) {
	v.value = val
	v.isSet = true
}

func (v NullableAPILimit) IsSet() bool {
	return v.isSet
}

func (v *NullableAPILimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPILimit(val *APILimit) *NullableAPILimit {
	return &NullableAPILimit{value: val, isSet: true}
}

func (v NullableAPILimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPILimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


