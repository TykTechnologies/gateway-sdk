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

// ListOrgKeys200Response struct for ListOrgKeys200Response
type ListOrgKeys200Response struct {
	Keys []string `json:"keys,omitempty"`
}

// NewListOrgKeys200Response instantiates a new ListOrgKeys200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrgKeys200Response() *ListOrgKeys200Response {
	this := ListOrgKeys200Response{}
	return &this
}

// NewListOrgKeys200ResponseWithDefaults instantiates a new ListOrgKeys200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrgKeys200ResponseWithDefaults() *ListOrgKeys200Response {
	this := ListOrgKeys200Response{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *ListOrgKeys200Response) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgKeys200Response) GetKeysOk() ([]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *ListOrgKeys200Response) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *ListOrgKeys200Response) SetKeys(v []string) {
	o.Keys = v
}

func (o ListOrgKeys200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableListOrgKeys200Response struct {
	value *ListOrgKeys200Response
	isSet bool
}

func (v NullableListOrgKeys200Response) Get() *ListOrgKeys200Response {
	return v.value
}

func (v *NullableListOrgKeys200Response) Set(val *ListOrgKeys200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrgKeys200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrgKeys200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrgKeys200Response(val *ListOrgKeys200Response) *NullableListOrgKeys200Response {
	return &NullableListOrgKeys200Response{value: val, isSet: true}
}

func (v NullableListOrgKeys200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrgKeys200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


