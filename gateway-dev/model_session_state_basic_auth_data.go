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

// SessionStateBasicAuthData struct for SessionStateBasicAuthData
type SessionStateBasicAuthData struct {
	HashType *string `json:"hash_type,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewSessionStateBasicAuthData instantiates a new SessionStateBasicAuthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStateBasicAuthData() *SessionStateBasicAuthData {
	this := SessionStateBasicAuthData{}
	return &this
}

// NewSessionStateBasicAuthDataWithDefaults instantiates a new SessionStateBasicAuthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStateBasicAuthDataWithDefaults() *SessionStateBasicAuthData {
	this := SessionStateBasicAuthData{}
	return &this
}

// GetHashType returns the HashType field value if set, zero value otherwise.
func (o *SessionStateBasicAuthData) GetHashType() string {
	if o == nil || o.HashType == nil {
		var ret string
		return ret
	}
	return *o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthData) GetHashTypeOk() (*string, bool) {
	if o == nil || o.HashType == nil {
		return nil, false
	}
	return o.HashType, true
}

// HasHashType returns a boolean if a field has been set.
func (o *SessionStateBasicAuthData) HasHashType() bool {
	if o != nil && o.HashType != nil {
		return true
	}

	return false
}

// SetHashType gets a reference to the given string and assigns it to the HashType field.
func (o *SessionStateBasicAuthData) SetHashType(v string) {
	o.HashType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SessionStateBasicAuthData) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthData) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SessionStateBasicAuthData) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SessionStateBasicAuthData) SetPassword(v string) {
	o.Password = &v
}

func (o SessionStateBasicAuthData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HashType != nil {
		toSerialize["hash_type"] = o.HashType
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSessionStateBasicAuthData struct {
	value *SessionStateBasicAuthData
	isSet bool
}

func (v NullableSessionStateBasicAuthData) Get() *SessionStateBasicAuthData {
	return v.value
}

func (v *NullableSessionStateBasicAuthData) Set(val *SessionStateBasicAuthData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStateBasicAuthData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStateBasicAuthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStateBasicAuthData(val *SessionStateBasicAuthData) *NullableSessionStateBasicAuthData {
	return &NullableSessionStateBasicAuthData{value: val, isSet: true}
}

func (v NullableSessionStateBasicAuthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStateBasicAuthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


