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

// SessionStateBasicAuthDataModel struct for SessionStateBasicAuthDataModel
type SessionStateBasicAuthDataModel struct {
	HashType *string `json:"hash_type,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewSessionStateBasicAuthDataModel instantiates a new SessionStateBasicAuthDataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStateBasicAuthDataModel() *SessionStateBasicAuthDataModel {
	this := SessionStateBasicAuthDataModel{}
	return &this
}

// NewSessionStateBasicAuthDataModelWithDefaults instantiates a new SessionStateBasicAuthDataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStateBasicAuthDataModelWithDefaults() *SessionStateBasicAuthDataModel {
	this := SessionStateBasicAuthDataModel{}
	return &this
}

// GetHashType returns the HashType field value if set, zero value otherwise.
func (o *SessionStateBasicAuthDataModel) GetHashType() string {
	if o == nil || o.HashType == nil {
		var ret string
		return ret
	}
	return *o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthDataModel) GetHashTypeOk() (*string, bool) {
	if o == nil || o.HashType == nil {
		return nil, false
	}
	return o.HashType, true
}

// HasHashType returns a boolean if a field has been set.
func (o *SessionStateBasicAuthDataModel) HasHashType() bool {
	if o != nil && o.HashType != nil {
		return true
	}

	return false
}

// SetHashType gets a reference to the given string and assigns it to the HashType field.
func (o *SessionStateBasicAuthDataModel) SetHashType(v string) {
	o.HashType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SessionStateBasicAuthDataModel) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthDataModel) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SessionStateBasicAuthDataModel) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SessionStateBasicAuthDataModel) SetPassword(v string) {
	o.Password = &v
}

func (o SessionStateBasicAuthDataModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HashType != nil {
		toSerialize["hash_type"] = o.HashType
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSessionStateBasicAuthDataModel struct {
	value *SessionStateBasicAuthDataModel
	isSet bool
}

func (v NullableSessionStateBasicAuthDataModel) Get() *SessionStateBasicAuthDataModel {
	return v.value
}

func (v *NullableSessionStateBasicAuthDataModel) Set(val *SessionStateBasicAuthDataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStateBasicAuthDataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStateBasicAuthDataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStateBasicAuthDataModel(val *SessionStateBasicAuthDataModel) *NullableSessionStateBasicAuthDataModel {
	return &NullableSessionStateBasicAuthDataModel{value: val, isSet: true}
}

func (v NullableSessionStateBasicAuthDataModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStateBasicAuthDataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


