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

// checks if the TemplateMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMeta{}

// TemplateMeta struct for TemplateMeta
type TemplateMeta struct {
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
	TemplateData *TemplateData `json:"template_data,omitempty"`
}

// NewTemplateMeta instantiates a new TemplateMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateMeta() *TemplateMeta {
	this := TemplateMeta{}
	return &this
}

// NewTemplateMetaWithDefaults instantiates a new TemplateMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMetaWithDefaults() *TemplateMeta {
	this := TemplateMeta{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TemplateMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TemplateMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *TemplateMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TemplateMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TemplateMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TemplateMeta) SetPath(v string) {
	o.Path = &v
}

// GetTemplateData returns the TemplateData field value if set, zero value otherwise.
func (o *TemplateMeta) GetTemplateData() TemplateData {
	if o == nil || IsNil(o.TemplateData) {
		var ret TemplateData
		return ret
	}
	return *o.TemplateData
}

// GetTemplateDataOk returns a tuple with the TemplateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMeta) GetTemplateDataOk() (*TemplateData, bool) {
	if o == nil || IsNil(o.TemplateData) {
		return nil, false
	}
	return o.TemplateData, true
}

// HasTemplateData returns a boolean if a field has been set.
func (o *TemplateMeta) HasTemplateData() bool {
	if o != nil && !IsNil(o.TemplateData) {
		return true
	}

	return false
}

// SetTemplateData gets a reference to the given TemplateData and assigns it to the TemplateData field.
func (o *TemplateMeta) SetTemplateData(v TemplateData) {
	o.TemplateData = &v
}

func (o TemplateMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.TemplateData) {
		toSerialize["template_data"] = o.TemplateData
	}
	return toSerialize, nil
}

type NullableTemplateMeta struct {
	value *TemplateMeta
	isSet bool
}

func (v NullableTemplateMeta) Get() *TemplateMeta {
	return v.value
}

func (v *NullableTemplateMeta) Set(val *TemplateMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMeta(val *TemplateMeta) *NullableTemplateMeta {
	return &NullableTemplateMeta{value: val, isSet: true}
}

func (v NullableTemplateMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


