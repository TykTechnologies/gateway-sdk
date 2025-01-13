/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation). * Managing and listing policies. * Managing and listing API Definitions (only when not using the Tyk Dashboard). * Hot reloads / reloading a cluster configuration. * OAuth client creation (only when not using the Tyk Dashboard).  In order to use the Gateway API, you'll need to set the **secret** parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  **x-tyk-authorization: <your-secret>*** <br/>  <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the EndpointMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointMethod{}

// EndpointMethod struct for EndpointMethod
type EndpointMethod struct {
	Limit *RateLimitType2 `json:"limit,omitempty"`
	Name  *string         `json:"name,omitempty"`
}

// NewEndpointMethod instantiates a new EndpointMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMethod() *EndpointMethod {
	this := EndpointMethod{}
	return &this
}

// NewEndpointMethodWithDefaults instantiates a new EndpointMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMethodWithDefaults() *EndpointMethod {
	this := EndpointMethod{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *EndpointMethod) GetLimit() RateLimitType2 {
	if o == nil || IsNil(o.Limit) {
		var ret RateLimitType2
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMethod) GetLimitOk() (*RateLimitType2, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *EndpointMethod) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given RateLimitType2 and assigns it to the Limit field.
func (o *EndpointMethod) SetLimit(v RateLimitType2) {
	o.Limit = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointMethod) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMethod) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointMethod) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointMethod) SetName(v string) {
	o.Name = &v
}

func (o EndpointMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEndpointMethod struct {
	value *EndpointMethod
	isSet bool
}

func (v NullableEndpointMethod) Get() *EndpointMethod {
	return v.value
}

func (v *NullableEndpointMethod) Set(val *EndpointMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMethod(val *EndpointMethod) *NullableEndpointMethod {
	return &NullableEndpointMethod{value: val, isSet: true}
}

func (v NullableEndpointMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
