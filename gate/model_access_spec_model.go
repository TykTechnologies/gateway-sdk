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

// AccessSpecModel AccessSpecs define what URLS a user has access to an what methods are enabled
type AccessSpecModel struct {
	Methods []string `json:"methods,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAccessSpecModel instantiates a new AccessSpecModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessSpecModel() *AccessSpecModel {
	this := AccessSpecModel{}
	return &this
}

// NewAccessSpecModelWithDefaults instantiates a new AccessSpecModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessSpecModelWithDefaults() *AccessSpecModel {
	this := AccessSpecModel{}
	return &this
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *AccessSpecModel) GetMethods() []string {
	if o == nil || o.Methods == nil {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSpecModel) GetMethodsOk() ([]string, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *AccessSpecModel) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *AccessSpecModel) SetMethods(v []string) {
	o.Methods = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AccessSpecModel) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSpecModel) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AccessSpecModel) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AccessSpecModel) SetUrl(v string) {
	o.Url = &v
}

func (o AccessSpecModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableAccessSpecModel struct {
	value *AccessSpecModel
	isSet bool
}

func (v NullableAccessSpecModel) Get() *AccessSpecModel {
	return v.value
}

func (v *NullableAccessSpecModel) Set(val *AccessSpecModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessSpecModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessSpecModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessSpecModel(val *AccessSpecModel) *NullableAccessSpecModel {
	return &NullableAccessSpecModel{value: val, isSet: true}
}

func (v NullableAccessSpecModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessSpecModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


