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

// SessionStateJwtDataModel struct for SessionStateJwtDataModel
type SessionStateJwtDataModel struct {
	Secret *string `json:"secret,omitempty"`
}

// NewSessionStateJwtDataModel instantiates a new SessionStateJwtDataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStateJwtDataModel() *SessionStateJwtDataModel {
	this := SessionStateJwtDataModel{}
	return &this
}

// NewSessionStateJwtDataModelWithDefaults instantiates a new SessionStateJwtDataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStateJwtDataModelWithDefaults() *SessionStateJwtDataModel {
	this := SessionStateJwtDataModel{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SessionStateJwtDataModel) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateJwtDataModel) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SessionStateJwtDataModel) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SessionStateJwtDataModel) SetSecret(v string) {
	o.Secret = &v
}

func (o SessionStateJwtDataModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableSessionStateJwtDataModel struct {
	value *SessionStateJwtDataModel
	isSet bool
}

func (v NullableSessionStateJwtDataModel) Get() *SessionStateJwtDataModel {
	return v.value
}

func (v *NullableSessionStateJwtDataModel) Set(val *SessionStateJwtDataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStateJwtDataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStateJwtDataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStateJwtDataModel(val *SessionStateJwtDataModel) *NullableSessionStateJwtDataModel {
	return &NullableSessionStateJwtDataModel{value: val, isSet: true}
}

func (v NullableSessionStateJwtDataModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStateJwtDataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


