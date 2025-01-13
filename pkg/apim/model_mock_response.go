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

// checks if the MockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MockResponse{}

// MockResponse struct for MockResponse
type MockResponse struct {
	Body            *string          `json:"body,omitempty"`
	Code            *int32           `json:"code,omitempty"`
	Enabled         *bool            `json:"enabled,omitempty"`
	FromOASExamples *FromOASExamples `json:"fromOASExamples,omitempty"`
	Headers         []Header         `json:"headers,omitempty"`
}

// NewMockResponse instantiates a new MockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMockResponse() *MockResponse {
	this := MockResponse{}
	return &this
}

// NewMockResponseWithDefaults instantiates a new MockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMockResponseWithDefaults() *MockResponse {
	this := MockResponse{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *MockResponse) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MockResponse) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *MockResponse) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *MockResponse) SetBody(v string) {
	o.Body = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MockResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MockResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MockResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *MockResponse) SetCode(v int32) {
	o.Code = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MockResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MockResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MockResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MockResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFromOASExamples returns the FromOASExamples field value if set, zero value otherwise.
func (o *MockResponse) GetFromOASExamples() FromOASExamples {
	if o == nil || IsNil(o.FromOASExamples) {
		var ret FromOASExamples
		return ret
	}
	return *o.FromOASExamples
}

// GetFromOASExamplesOk returns a tuple with the FromOASExamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MockResponse) GetFromOASExamplesOk() (*FromOASExamples, bool) {
	if o == nil || IsNil(o.FromOASExamples) {
		return nil, false
	}
	return o.FromOASExamples, true
}

// HasFromOASExamples returns a boolean if a field has been set.
func (o *MockResponse) HasFromOASExamples() bool {
	if o != nil && !IsNil(o.FromOASExamples) {
		return true
	}

	return false
}

// SetFromOASExamples gets a reference to the given FromOASExamples and assigns it to the FromOASExamples field.
func (o *MockResponse) SetFromOASExamples(v FromOASExamples) {
	o.FromOASExamples = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *MockResponse) GetHeaders() []Header {
	if o == nil || IsNil(o.Headers) {
		var ret []Header
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MockResponse) GetHeadersOk() ([]Header, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *MockResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *MockResponse) SetHeaders(v []Header) {
	o.Headers = v
}

func (o MockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.FromOASExamples) {
		toSerialize["fromOASExamples"] = o.FromOASExamples
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableMockResponse struct {
	value *MockResponse
	isSet bool
}

func (v NullableMockResponse) Get() *MockResponse {
	return v.value
}

func (v *NullableMockResponse) Set(val *MockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMockResponse(val *MockResponse) *NullableMockResponse {
	return &NullableMockResponse{value: val, isSet: true}
}

func (v NullableMockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
