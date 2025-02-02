/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation). * Managing and listing policies. * Managing and listing API Definitions (only when not using the Tyk Dashboard). * Hot reloads / reloading a cluster configuration. * OAuth client creation (only when not using the Tyk Dashboard).  In order to use the Gateway API, you'll need to set the **secret** parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  **x-tyk-authorization: <your-secret>*** <br/>  <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Server1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Server1{}

// Server1 struct for Server1
type Server1 struct {
	Url         string                     `json:"url"`
	Description *string                    `json:"description,omitempty"`
	Variables   *map[string]ServerVariable `json:"variables,omitempty"`
}

type _Server1 Server1

// NewServer1 instantiates a new Server1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer1(url string) *Server1 {
	this := Server1{}
	this.Url = url
	return &this
}

// NewServer1WithDefaults instantiates a new Server1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServer1WithDefaults() *Server1 {
	this := Server1{}
	return &this
}

// GetUrl returns the Url field value
func (o *Server1) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Server1) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Server1) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Server1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Server1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Server1) SetDescription(v string) {
	o.Description = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *Server1) GetVariables() map[string]ServerVariable {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]ServerVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server1) GetVariablesOk() (*map[string]ServerVariable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *Server1) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]ServerVariable and assigns it to the Variables field.
func (o *Server1) SetVariables(v map[string]ServerVariable) {
	o.Variables = &v
}

func (o Server1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Server1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *Server1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServer1 := _Server1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServer1)

	if err != nil {
		return err
	}

	*o = Server1(varServer1)

	return err
}

type NullableServer1 struct {
	value *Server1
	isSet bool
}

func (v NullableServer1) Get() *Server1 {
	return v.value
}

func (v *NullableServer1) Set(val *Server1) {
	v.value = val
	v.isSet = true
}

func (v NullableServer1) IsSet() bool {
	return v.isSet
}

func (v *NullableServer1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer1(val *Server1) *NullableServer1 {
	return &NullableServer1{value: val, isSet: true}
}

func (v NullableServer1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
