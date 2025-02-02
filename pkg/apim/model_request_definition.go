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

// checks if the RequestDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestDefinition{}

// RequestDefinition struct for RequestDefinition
type RequestDefinition struct {
	Body        *string           `json:"body,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Method      *string           `json:"method,omitempty"`
	RelativeUrl *string           `json:"relative_url,omitempty"`
}

// NewRequestDefinition instantiates a new RequestDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestDefinition() *RequestDefinition {
	this := RequestDefinition{}
	return &this
}

// NewRequestDefinitionWithDefaults instantiates a new RequestDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestDefinitionWithDefaults() *RequestDefinition {
	this := RequestDefinition{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *RequestDefinition) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDefinition) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *RequestDefinition) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *RequestDefinition) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestDefinition) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestDefinition) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RequestDefinition) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *RequestDefinition) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RequestDefinition) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDefinition) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RequestDefinition) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RequestDefinition) SetMethod(v string) {
	o.Method = &v
}

// GetRelativeUrl returns the RelativeUrl field value if set, zero value otherwise.
func (o *RequestDefinition) GetRelativeUrl() string {
	if o == nil || IsNil(o.RelativeUrl) {
		var ret string
		return ret
	}
	return *o.RelativeUrl
}

// GetRelativeUrlOk returns a tuple with the RelativeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDefinition) GetRelativeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RelativeUrl) {
		return nil, false
	}
	return o.RelativeUrl, true
}

// HasRelativeUrl returns a boolean if a field has been set.
func (o *RequestDefinition) HasRelativeUrl() bool {
	if o != nil && !IsNil(o.RelativeUrl) {
		return true
	}

	return false
}

// SetRelativeUrl gets a reference to the given string and assigns it to the RelativeUrl field.
func (o *RequestDefinition) SetRelativeUrl(v string) {
	o.RelativeUrl = &v
}

func (o RequestDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.RelativeUrl) {
		toSerialize["relative_url"] = o.RelativeUrl
	}
	return toSerialize, nil
}

type NullableRequestDefinition struct {
	value *RequestDefinition
	isSet bool
}

func (v NullableRequestDefinition) Get() *RequestDefinition {
	return v.value
}

func (v *NullableRequestDefinition) Set(val *RequestDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestDefinition(val *RequestDefinition) *NullableRequestDefinition {
	return &NullableRequestDefinition{value: val, isSet: true}
}

func (v NullableRequestDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
