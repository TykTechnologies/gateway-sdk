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

// RegexpModel Regexp is a wrapper around regexp.Regexp but with caching
type RegexpModel struct {
	FromCache *bool `json:"FromCache,omitempty"`
}

// NewRegexpModel instantiates a new RegexpModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegexpModel() *RegexpModel {
	this := RegexpModel{}
	return &this
}

// NewRegexpModelWithDefaults instantiates a new RegexpModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegexpModelWithDefaults() *RegexpModel {
	this := RegexpModel{}
	return &this
}

// GetFromCache returns the FromCache field value if set, zero value otherwise.
func (o *RegexpModel) GetFromCache() bool {
	if o == nil || o.FromCache == nil {
		var ret bool
		return ret
	}
	return *o.FromCache
}

// GetFromCacheOk returns a tuple with the FromCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexpModel) GetFromCacheOk() (*bool, bool) {
	if o == nil || o.FromCache == nil {
		return nil, false
	}
	return o.FromCache, true
}

// HasFromCache returns a boolean if a field has been set.
func (o *RegexpModel) HasFromCache() bool {
	if o != nil && o.FromCache != nil {
		return true
	}

	return false
}

// SetFromCache gets a reference to the given bool and assigns it to the FromCache field.
func (o *RegexpModel) SetFromCache(v bool) {
	o.FromCache = &v
}

func (o RegexpModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromCache != nil {
		toSerialize["FromCache"] = o.FromCache
	}
	return json.Marshal(toSerialize)
}

type NullableRegexpModel struct {
	value *RegexpModel
	isSet bool
}

func (v NullableRegexpModel) Get() *RegexpModel {
	return v.value
}

func (v *NullableRegexpModel) Set(val *RegexpModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRegexpModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRegexpModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegexpModel(val *RegexpModel) *NullableRegexpModel {
	return &NullableRegexpModel{value: val, isSet: true}
}

func (v NullableRegexpModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegexpModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


