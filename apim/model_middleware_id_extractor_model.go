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

// MiddlewareIdExtractorModel struct for MiddlewareIdExtractorModel
type MiddlewareIdExtractorModel struct {
	ExtractFrom *string `json:"extract_from,omitempty"`
	ExtractWith *string `json:"extract_with,omitempty"`
	ExtractorConfig map[string]map[string]interface{} `json:"extractor_config,omitempty"`
}

// NewMiddlewareIdExtractorModel instantiates a new MiddlewareIdExtractorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareIdExtractorModel() *MiddlewareIdExtractorModel {
	this := MiddlewareIdExtractorModel{}
	return &this
}

// NewMiddlewareIdExtractorModelWithDefaults instantiates a new MiddlewareIdExtractorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareIdExtractorModelWithDefaults() *MiddlewareIdExtractorModel {
	this := MiddlewareIdExtractorModel{}
	return &this
}

// GetExtractFrom returns the ExtractFrom field value if set, zero value otherwise.
func (o *MiddlewareIdExtractorModel) GetExtractFrom() string {
	if o == nil || o.ExtractFrom == nil {
		var ret string
		return ret
	}
	return *o.ExtractFrom
}

// GetExtractFromOk returns a tuple with the ExtractFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractorModel) GetExtractFromOk() (*string, bool) {
	if o == nil || o.ExtractFrom == nil {
		return nil, false
	}
	return o.ExtractFrom, true
}

// HasExtractFrom returns a boolean if a field has been set.
func (o *MiddlewareIdExtractorModel) HasExtractFrom() bool {
	if o != nil && o.ExtractFrom != nil {
		return true
	}

	return false
}

// SetExtractFrom gets a reference to the given string and assigns it to the ExtractFrom field.
func (o *MiddlewareIdExtractorModel) SetExtractFrom(v string) {
	o.ExtractFrom = &v
}

// GetExtractWith returns the ExtractWith field value if set, zero value otherwise.
func (o *MiddlewareIdExtractorModel) GetExtractWith() string {
	if o == nil || o.ExtractWith == nil {
		var ret string
		return ret
	}
	return *o.ExtractWith
}

// GetExtractWithOk returns a tuple with the ExtractWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractorModel) GetExtractWithOk() (*string, bool) {
	if o == nil || o.ExtractWith == nil {
		return nil, false
	}
	return o.ExtractWith, true
}

// HasExtractWith returns a boolean if a field has been set.
func (o *MiddlewareIdExtractorModel) HasExtractWith() bool {
	if o != nil && o.ExtractWith != nil {
		return true
	}

	return false
}

// SetExtractWith gets a reference to the given string and assigns it to the ExtractWith field.
func (o *MiddlewareIdExtractorModel) SetExtractWith(v string) {
	o.ExtractWith = &v
}

// GetExtractorConfig returns the ExtractorConfig field value if set, zero value otherwise.
func (o *MiddlewareIdExtractorModel) GetExtractorConfig() map[string]map[string]interface{} {
	if o == nil || o.ExtractorConfig == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtractorConfig
}

// GetExtractorConfigOk returns a tuple with the ExtractorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractorModel) GetExtractorConfigOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ExtractorConfig == nil {
		return nil, false
	}
	return o.ExtractorConfig, true
}

// HasExtractorConfig returns a boolean if a field has been set.
func (o *MiddlewareIdExtractorModel) HasExtractorConfig() bool {
	if o != nil && o.ExtractorConfig != nil {
		return true
	}

	return false
}

// SetExtractorConfig gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtractorConfig field.
func (o *MiddlewareIdExtractorModel) SetExtractorConfig(v map[string]map[string]interface{}) {
	o.ExtractorConfig = v
}

func (o MiddlewareIdExtractorModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtractFrom != nil {
		toSerialize["extract_from"] = o.ExtractFrom
	}
	if o.ExtractWith != nil {
		toSerialize["extract_with"] = o.ExtractWith
	}
	if o.ExtractorConfig != nil {
		toSerialize["extractor_config"] = o.ExtractorConfig
	}
	return json.Marshal(toSerialize)
}

type NullableMiddlewareIdExtractorModel struct {
	value *MiddlewareIdExtractorModel
	isSet bool
}

func (v NullableMiddlewareIdExtractorModel) Get() *MiddlewareIdExtractorModel {
	return v.value
}

func (v *NullableMiddlewareIdExtractorModel) Set(val *MiddlewareIdExtractorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareIdExtractorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareIdExtractorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareIdExtractorModel(val *MiddlewareIdExtractorModel) *NullableMiddlewareIdExtractorModel {
	return &NullableMiddlewareIdExtractorModel{value: val, isSet: true}
}

func (v NullableMiddlewareIdExtractorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareIdExtractorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


