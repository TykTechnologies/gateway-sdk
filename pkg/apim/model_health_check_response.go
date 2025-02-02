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

// checks if the HealthCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckResponse{}

// HealthCheckResponse struct for HealthCheckResponse
type HealthCheckResponse struct {
	Description *string                     `json:"description,omitempty"`
	Details     *map[string]HealthCheckItem `json:"details,omitempty"`
	Output      *string                     `json:"output,omitempty"`
	Status      *string                     `json:"status,omitempty"`
	Version     *string                     `json:"version,omitempty"`
}

// NewHealthCheckResponse instantiates a new HealthCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResponse() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResponseWithDefaults() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HealthCheckResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetDetails() map[string]HealthCheckItem {
	if o == nil || IsNil(o.Details) {
		var ret map[string]HealthCheckItem
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetDetailsOk() (*map[string]HealthCheckItem, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]HealthCheckItem and assigns it to the Details field.
func (o *HealthCheckResponse) SetDetails(v map[string]HealthCheckItem) {
	o.Details = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetOutput() string {
	if o == nil || IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *HealthCheckResponse) SetOutput(v string) {
	o.Output = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthCheckResponse) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HealthCheckResponse) SetVersion(v string) {
	o.Version = &v
}

func (o HealthCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableHealthCheckResponse struct {
	value *HealthCheckResponse
	isSet bool
}

func (v NullableHealthCheckResponse) Get() *HealthCheckResponse {
	return v.value
}

func (v *NullableHealthCheckResponse) Set(val *HealthCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResponse(val *HealthCheckResponse) *NullableHealthCheckResponse {
	return &NullableHealthCheckResponse{value: val, isSet: true}
}

func (v NullableHealthCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
