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

// checks if the PolicyPartitions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyPartitions{}

// PolicyPartitions struct for PolicyPartitions
type PolicyPartitions struct {
	Acl *bool `json:"acl,omitempty"`
	Complexity *bool `json:"complexity,omitempty"`
	PerApi *bool `json:"per_api,omitempty"`
	Quota *bool `json:"quota,omitempty"`
	RateLimit *bool `json:"rate_limit,omitempty"`
}

// NewPolicyPartitions instantiates a new PolicyPartitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyPartitions() *PolicyPartitions {
	this := PolicyPartitions{}
	return &this
}

// NewPolicyPartitionsWithDefaults instantiates a new PolicyPartitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyPartitionsWithDefaults() *PolicyPartitions {
	this := PolicyPartitions{}
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *PolicyPartitions) GetAcl() bool {
	if o == nil || IsNil(o.Acl) {
		var ret bool
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPartitions) GetAclOk() (*bool, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *PolicyPartitions) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given bool and assigns it to the Acl field.
func (o *PolicyPartitions) SetAcl(v bool) {
	o.Acl = &v
}

// GetComplexity returns the Complexity field value if set, zero value otherwise.
func (o *PolicyPartitions) GetComplexity() bool {
	if o == nil || IsNil(o.Complexity) {
		var ret bool
		return ret
	}
	return *o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPartitions) GetComplexityOk() (*bool, bool) {
	if o == nil || IsNil(o.Complexity) {
		return nil, false
	}
	return o.Complexity, true
}

// HasComplexity returns a boolean if a field has been set.
func (o *PolicyPartitions) HasComplexity() bool {
	if o != nil && !IsNil(o.Complexity) {
		return true
	}

	return false
}

// SetComplexity gets a reference to the given bool and assigns it to the Complexity field.
func (o *PolicyPartitions) SetComplexity(v bool) {
	o.Complexity = &v
}

// GetPerApi returns the PerApi field value if set, zero value otherwise.
func (o *PolicyPartitions) GetPerApi() bool {
	if o == nil || IsNil(o.PerApi) {
		var ret bool
		return ret
	}
	return *o.PerApi
}

// GetPerApiOk returns a tuple with the PerApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPartitions) GetPerApiOk() (*bool, bool) {
	if o == nil || IsNil(o.PerApi) {
		return nil, false
	}
	return o.PerApi, true
}

// HasPerApi returns a boolean if a field has been set.
func (o *PolicyPartitions) HasPerApi() bool {
	if o != nil && !IsNil(o.PerApi) {
		return true
	}

	return false
}

// SetPerApi gets a reference to the given bool and assigns it to the PerApi field.
func (o *PolicyPartitions) SetPerApi(v bool) {
	o.PerApi = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *PolicyPartitions) GetQuota() bool {
	if o == nil || IsNil(o.Quota) {
		var ret bool
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPartitions) GetQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *PolicyPartitions) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given bool and assigns it to the Quota field.
func (o *PolicyPartitions) SetQuota(v bool) {
	o.Quota = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *PolicyPartitions) GetRateLimit() bool {
	if o == nil || IsNil(o.RateLimit) {
		var ret bool
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPartitions) GetRateLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.RateLimit) {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *PolicyPartitions) HasRateLimit() bool {
	if o != nil && !IsNil(o.RateLimit) {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given bool and assigns it to the RateLimit field.
func (o *PolicyPartitions) SetRateLimit(v bool) {
	o.RateLimit = &v
}

func (o PolicyPartitions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyPartitions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.Complexity) {
		toSerialize["complexity"] = o.Complexity
	}
	if !IsNil(o.PerApi) {
		toSerialize["per_api"] = o.PerApi
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.RateLimit) {
		toSerialize["rate_limit"] = o.RateLimit
	}
	return toSerialize, nil
}

type NullablePolicyPartitions struct {
	value *PolicyPartitions
	isSet bool
}

func (v NullablePolicyPartitions) Get() *PolicyPartitions {
	return v.value
}

func (v *NullablePolicyPartitions) Set(val *PolicyPartitions) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyPartitions) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyPartitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyPartitions(val *PolicyPartitions) *NullablePolicyPartitions {
	return &NullablePolicyPartitions{value: val, isSet: true}
}

func (v NullablePolicyPartitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyPartitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


