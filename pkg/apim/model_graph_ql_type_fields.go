/*
Tyk Gateway API

 The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway

API version: 5.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the GraphQLTypeFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLTypeFields{}

// GraphQLTypeFields struct for GraphQLTypeFields
type GraphQLTypeFields struct {
	Fields []string `json:"fields,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGraphQLTypeFields instantiates a new GraphQLTypeFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLTypeFields() *GraphQLTypeFields {
	this := GraphQLTypeFields{}
	return &this
}

// NewGraphQLTypeFieldsWithDefaults instantiates a new GraphQLTypeFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLTypeFieldsWithDefaults() *GraphQLTypeFields {
	this := GraphQLTypeFields{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLTypeFields) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLTypeFields) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *GraphQLTypeFields) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *GraphQLTypeFields) SetFields(v []string) {
	o.Fields = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GraphQLTypeFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLTypeFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GraphQLTypeFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GraphQLTypeFields) SetType(v string) {
	o.Type = &v
}

func (o GraphQLTypeFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLTypeFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGraphQLTypeFields struct {
	value *GraphQLTypeFields
	isSet bool
}

func (v NullableGraphQLTypeFields) Get() *GraphQLTypeFields {
	return v.value
}

func (v *NullableGraphQLTypeFields) Set(val *GraphQLTypeFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLTypeFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLTypeFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLTypeFields(val *GraphQLTypeFields) *NullableGraphQLTypeFields {
	return &NullableGraphQLTypeFields{value: val, isSet: true}
}

func (v NullableGraphQLTypeFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLTypeFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


