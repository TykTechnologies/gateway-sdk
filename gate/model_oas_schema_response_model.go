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

// OASSchemaResponseModel OAS schema endpoint response
type OASSchemaResponseModel struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	// <OAS schema definition>
	Schema *string `json:"schema,omitempty"`
}

// NewOASSchemaResponseModel instantiates a new OASSchemaResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOASSchemaResponseModel() *OASSchemaResponseModel {
	this := OASSchemaResponseModel{}
	return &this
}

// NewOASSchemaResponseModelWithDefaults instantiates a new OASSchemaResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOASSchemaResponseModelWithDefaults() *OASSchemaResponseModel {
	this := OASSchemaResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OASSchemaResponseModel) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OASSchemaResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OASSchemaResponseModel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OASSchemaResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OASSchemaResponseModel) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OASSchemaResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OASSchemaResponseModel) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OASSchemaResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OASSchemaResponseModel) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OASSchemaResponseModel) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OASSchemaResponseModel) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OASSchemaResponseModel) SetSchema(v string) {
	o.Schema = &v
}

func (o OASSchemaResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableOASSchemaResponseModel struct {
	value *OASSchemaResponseModel
	isSet bool
}

func (v NullableOASSchemaResponseModel) Get() *OASSchemaResponseModel {
	return v.value
}

func (v *NullableOASSchemaResponseModel) Set(val *OASSchemaResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOASSchemaResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOASSchemaResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOASSchemaResponseModel(val *OASSchemaResponseModel) *NullableOASSchemaResponseModel {
	return &NullableOASSchemaResponseModel{value: val, isSet: true}
}

func (v NullableOASSchemaResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOASSchemaResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


