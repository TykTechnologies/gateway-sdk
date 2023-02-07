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

// LicenseModel struct for LicenseModel
type LicenseModel struct {
	Name string `json:"name"`
	Url *string `json:"url,omitempty"`
}

// NewLicenseModel instantiates a new LicenseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseModel(name string) *LicenseModel {
	this := LicenseModel{}
	this.Name = name
	return &this
}

// NewLicenseModelWithDefaults instantiates a new LicenseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseModelWithDefaults() *LicenseModel {
	this := LicenseModel{}
	return &this
}

// GetName returns the Name field value
func (o *LicenseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LicenseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LicenseModel) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LicenseModel) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseModel) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LicenseModel) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LicenseModel) SetUrl(v string) {
	o.Url = &v
}

func (o LicenseModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseModel struct {
	value *LicenseModel
	isSet bool
}

func (v NullableLicenseModel) Get() *LicenseModel {
	return v.value
}

func (v *NullableLicenseModel) Set(val *LicenseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseModel(val *LicenseModel) *NullableLicenseModel {
	return &NullableLicenseModel{value: val, isSet: true}
}

func (v NullableLicenseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


