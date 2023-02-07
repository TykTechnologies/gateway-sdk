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

// CacheMetaModel struct for CacheMetaModel
type CacheMetaModel struct {
	CacheResponseCodes []int64 `json:"cache_response_codes,omitempty"`
	CacheKeyRegex *string `json:"cache_key_regex,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewCacheMetaModel instantiates a new CacheMetaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheMetaModel() *CacheMetaModel {
	this := CacheMetaModel{}
	return &this
}

// NewCacheMetaModelWithDefaults instantiates a new CacheMetaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheMetaModelWithDefaults() *CacheMetaModel {
	this := CacheMetaModel{}
	return &this
}

// GetCacheResponseCodes returns the CacheResponseCodes field value if set, zero value otherwise.
func (o *CacheMetaModel) GetCacheResponseCodes() []int64 {
	if o == nil || o.CacheResponseCodes == nil {
		var ret []int64
		return ret
	}
	return o.CacheResponseCodes
}

// GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMetaModel) GetCacheResponseCodesOk() ([]int64, bool) {
	if o == nil || o.CacheResponseCodes == nil {
		return nil, false
	}
	return o.CacheResponseCodes, true
}

// HasCacheResponseCodes returns a boolean if a field has been set.
func (o *CacheMetaModel) HasCacheResponseCodes() bool {
	if o != nil && o.CacheResponseCodes != nil {
		return true
	}

	return false
}

// SetCacheResponseCodes gets a reference to the given []int64 and assigns it to the CacheResponseCodes field.
func (o *CacheMetaModel) SetCacheResponseCodes(v []int64) {
	o.CacheResponseCodes = v
}

// GetCacheKeyRegex returns the CacheKeyRegex field value if set, zero value otherwise.
func (o *CacheMetaModel) GetCacheKeyRegex() string {
	if o == nil || o.CacheKeyRegex == nil {
		var ret string
		return ret
	}
	return *o.CacheKeyRegex
}

// GetCacheKeyRegexOk returns a tuple with the CacheKeyRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMetaModel) GetCacheKeyRegexOk() (*string, bool) {
	if o == nil || o.CacheKeyRegex == nil {
		return nil, false
	}
	return o.CacheKeyRegex, true
}

// HasCacheKeyRegex returns a boolean if a field has been set.
func (o *CacheMetaModel) HasCacheKeyRegex() bool {
	if o != nil && o.CacheKeyRegex != nil {
		return true
	}

	return false
}

// SetCacheKeyRegex gets a reference to the given string and assigns it to the CacheKeyRegex field.
func (o *CacheMetaModel) SetCacheKeyRegex(v string) {
	o.CacheKeyRegex = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CacheMetaModel) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMetaModel) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CacheMetaModel) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CacheMetaModel) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CacheMetaModel) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMetaModel) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CacheMetaModel) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CacheMetaModel) SetPath(v string) {
	o.Path = &v
}

func (o CacheMetaModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheResponseCodes != nil {
		toSerialize["cache_response_codes"] = o.CacheResponseCodes
	}
	if o.CacheKeyRegex != nil {
		toSerialize["cache_key_regex"] = o.CacheKeyRegex
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableCacheMetaModel struct {
	value *CacheMetaModel
	isSet bool
}

func (v NullableCacheMetaModel) Get() *CacheMetaModel {
	return v.value
}

func (v *NullableCacheMetaModel) Set(val *CacheMetaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheMetaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheMetaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheMetaModel(val *CacheMetaModel) *NullableCacheMetaModel {
	return &NullableCacheMetaModel{value: val, isSet: true}
}

func (v NullableCacheMetaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheMetaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


