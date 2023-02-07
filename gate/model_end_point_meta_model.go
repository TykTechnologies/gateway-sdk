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

// EndPointMetaModel struct for EndPointMetaModel
type EndPointMetaModel struct {
	MethodActions *map[string]EndpointMethodMetaModelModel `json:"method_actions,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewEndPointMetaModel instantiates a new EndPointMetaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndPointMetaModel() *EndPointMetaModel {
	this := EndPointMetaModel{}
	return &this
}

// NewEndPointMetaModelWithDefaults instantiates a new EndPointMetaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndPointMetaModelWithDefaults() *EndPointMetaModel {
	this := EndPointMetaModel{}
	return &this
}

// GetMethodActions returns the MethodActions field value if set, zero value otherwise.
func (o *EndPointMetaModel) GetMethodActions() map[string]EndpointMethodMetaModelModel {
	if o == nil || o.MethodActions == nil {
		var ret map[string]EndpointMethodMetaModelModel
		return ret
	}
	return *o.MethodActions
}

// GetMethodActionsOk returns a tuple with the MethodActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndPointMetaModel) GetMethodActionsOk() (*map[string]EndpointMethodMetaModelModel, bool) {
	if o == nil || o.MethodActions == nil {
		return nil, false
	}
	return o.MethodActions, true
}

// HasMethodActions returns a boolean if a field has been set.
func (o *EndPointMetaModel) HasMethodActions() bool {
	if o != nil && o.MethodActions != nil {
		return true
	}

	return false
}

// SetMethodActions gets a reference to the given map[string]EndpointMethodMetaModelModel and assigns it to the MethodActions field.
func (o *EndPointMetaModel) SetMethodActions(v map[string]EndpointMethodMetaModelModel) {
	o.MethodActions = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *EndPointMetaModel) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndPointMetaModel) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *EndPointMetaModel) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *EndPointMetaModel) SetPath(v string) {
	o.Path = &v
}

func (o EndPointMetaModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MethodActions != nil {
		toSerialize["method_actions"] = o.MethodActions
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableEndPointMetaModel struct {
	value *EndPointMetaModel
	isSet bool
}

func (v NullableEndPointMetaModel) Get() *EndPointMetaModel {
	return v.value
}

func (v *NullableEndPointMetaModel) Set(val *EndPointMetaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEndPointMetaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEndPointMetaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndPointMetaModel(val *EndPointMetaModel) *NullableEndPointMetaModel {
	return &NullableEndPointMetaModel{value: val, isSet: true}
}

func (v NullableEndPointMetaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndPointMetaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


