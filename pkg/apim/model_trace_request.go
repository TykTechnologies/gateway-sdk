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

// checks if the TraceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceRequest{}

// TraceRequest struct for TraceRequest
type TraceRequest struct {
	Request *TraceHttpRequest `json:"request,omitempty"`
	Spec    *APIDefinition    `json:"spec,omitempty"`
}

// NewTraceRequest instantiates a new TraceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceRequest() *TraceRequest {
	this := TraceRequest{}
	return &this
}

// NewTraceRequestWithDefaults instantiates a new TraceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceRequestWithDefaults() *TraceRequest {
	this := TraceRequest{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TraceRequest) GetRequest() TraceHttpRequest {
	if o == nil || IsNil(o.Request) {
		var ret TraceHttpRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceRequest) GetRequestOk() (*TraceHttpRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TraceRequest) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given TraceHttpRequest and assigns it to the Request field.
func (o *TraceRequest) SetRequest(v TraceHttpRequest) {
	o.Request = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *TraceRequest) GetSpec() APIDefinition {
	if o == nil || IsNil(o.Spec) {
		var ret APIDefinition
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceRequest) GetSpecOk() (*APIDefinition, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *TraceRequest) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given APIDefinition and assigns it to the Spec field.
func (o *TraceRequest) SetSpec(v APIDefinition) {
	o.Spec = &v
}

func (o TraceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableTraceRequest struct {
	value *TraceRequest
	isSet bool
}

func (v NullableTraceRequest) Get() *TraceRequest {
	return v.value
}

func (v *NullableTraceRequest) Set(val *TraceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceRequest(val *TraceRequest) *NullableTraceRequest {
	return &NullableTraceRequest{value: val, isSet: true}
}

func (v NullableTraceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
