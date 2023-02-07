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

// TagModel struct for TagModel
type TagModel struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationModelModel `json:"externalDocs,omitempty"`
}

// NewTagModel instantiates a new TagModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModel(name string) *TagModel {
	this := TagModel{}
	this.Name = name
	return &this
}

// NewTagModelWithDefaults instantiates a new TagModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelWithDefaults() *TagModel {
	this := TagModel{}
	return &this
}

// GetName returns the Name field value
func (o *TagModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagModel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagModel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagModel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagModel) SetDescription(v string) {
	o.Description = &v
}

// GetExternalDocs returns the ExternalDocs field value if set, zero value otherwise.
func (o *TagModel) GetExternalDocs() ExternalDocumentationModelModel {
	if o == nil || o.ExternalDocs == nil {
		var ret ExternalDocumentationModelModel
		return ret
	}
	return *o.ExternalDocs
}

// GetExternalDocsOk returns a tuple with the ExternalDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagModel) GetExternalDocsOk() (*ExternalDocumentationModelModel, bool) {
	if o == nil || o.ExternalDocs == nil {
		return nil, false
	}
	return o.ExternalDocs, true
}

// HasExternalDocs returns a boolean if a field has been set.
func (o *TagModel) HasExternalDocs() bool {
	if o != nil && o.ExternalDocs != nil {
		return true
	}

	return false
}

// SetExternalDocs gets a reference to the given ExternalDocumentationModelModel and assigns it to the ExternalDocs field.
func (o *TagModel) SetExternalDocs(v ExternalDocumentationModelModel) {
	o.ExternalDocs = &v
}

func (o TagModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExternalDocs != nil {
		toSerialize["externalDocs"] = o.ExternalDocs
	}
	return json.Marshal(toSerialize)
}

type NullableTagModel struct {
	value *TagModel
	isSet bool
}

func (v NullableTagModel) Get() *TagModel {
	return v.value
}

func (v *NullableTagModel) Set(val *TagModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModel(val *TagModel) *NullableTagModel {
	return &NullableTagModel{value: val, isSet: true}
}

func (v NullableTagModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


