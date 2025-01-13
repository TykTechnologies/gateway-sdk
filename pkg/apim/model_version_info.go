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

// checks if the VersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfo{}

// VersionInfo struct for VersionInfo
type VersionInfo struct {
	Expires                       *string           `json:"expires,omitempty"`
	ExtendedPaths                 *ExtendedPathsSet `json:"extended_paths,omitempty"`
	GlobalHeaders                 map[string]string `json:"global_headers,omitempty"`
	GlobalHeadersDisabled         *bool             `json:"global_headers_disabled,omitempty"`
	GlobalHeadersRemove           []string          `json:"global_headers_remove,omitempty"`
	GlobalResponseHeaders         map[string]string `json:"global_response_headers,omitempty"`
	GlobalResponseHeadersDisabled *bool             `json:"global_response_headers_disabled,omitempty"`
	GlobalResponseHeadersRemove   []string          `json:"global_response_headers_remove,omitempty"`
	GlobalSizeLimit               *int64            `json:"global_size_limit,omitempty"`
	IgnoreEndpointCase            *bool             `json:"ignore_endpoint_case,omitempty"`
	Name                          *string           `json:"name,omitempty"`
	OverrideTarget                *string           `json:"override_target,omitempty"`
	Paths                         *VersionInfoPaths `json:"paths,omitempty"`
	UseExtendedPaths              *bool             `json:"use_extended_paths,omitempty"`
}

// NewVersionInfo instantiates a new VersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfo() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// NewVersionInfoWithDefaults instantiates a new VersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoWithDefaults() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *VersionInfo) GetExpires() string {
	if o == nil || IsNil(o.Expires) {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetExpiresOk() (*string, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *VersionInfo) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *VersionInfo) SetExpires(v string) {
	o.Expires = &v
}

// GetExtendedPaths returns the ExtendedPaths field value if set, zero value otherwise.
func (o *VersionInfo) GetExtendedPaths() ExtendedPathsSet {
	if o == nil || IsNil(o.ExtendedPaths) {
		var ret ExtendedPathsSet
		return ret
	}
	return *o.ExtendedPaths
}

// GetExtendedPathsOk returns a tuple with the ExtendedPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetExtendedPathsOk() (*ExtendedPathsSet, bool) {
	if o == nil || IsNil(o.ExtendedPaths) {
		return nil, false
	}
	return o.ExtendedPaths, true
}

// HasExtendedPaths returns a boolean if a field has been set.
func (o *VersionInfo) HasExtendedPaths() bool {
	if o != nil && !IsNil(o.ExtendedPaths) {
		return true
	}

	return false
}

// SetExtendedPaths gets a reference to the given ExtendedPathsSet and assigns it to the ExtendedPaths field.
func (o *VersionInfo) SetExtendedPaths(v ExtendedPathsSet) {
	o.ExtendedPaths = &v
}

// GetGlobalHeaders returns the GlobalHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetGlobalHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.GlobalHeaders
}

// GetGlobalHeadersOk returns a tuple with the GlobalHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetGlobalHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.GlobalHeaders) {
		return nil, false
	}
	return &o.GlobalHeaders, true
}

// HasGlobalHeaders returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalHeaders() bool {
	if o != nil && !IsNil(o.GlobalHeaders) {
		return true
	}

	return false
}

// SetGlobalHeaders gets a reference to the given map[string]string and assigns it to the GlobalHeaders field.
func (o *VersionInfo) SetGlobalHeaders(v map[string]string) {
	o.GlobalHeaders = v
}

// GetGlobalHeadersDisabled returns the GlobalHeadersDisabled field value if set, zero value otherwise.
func (o *VersionInfo) GetGlobalHeadersDisabled() bool {
	if o == nil || IsNil(o.GlobalHeadersDisabled) {
		var ret bool
		return ret
	}
	return *o.GlobalHeadersDisabled
}

// GetGlobalHeadersDisabledOk returns a tuple with the GlobalHeadersDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetGlobalHeadersDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalHeadersDisabled) {
		return nil, false
	}
	return o.GlobalHeadersDisabled, true
}

// HasGlobalHeadersDisabled returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalHeadersDisabled() bool {
	if o != nil && !IsNil(o.GlobalHeadersDisabled) {
		return true
	}

	return false
}

// SetGlobalHeadersDisabled gets a reference to the given bool and assigns it to the GlobalHeadersDisabled field.
func (o *VersionInfo) SetGlobalHeadersDisabled(v bool) {
	o.GlobalHeadersDisabled = &v
}

// GetGlobalHeadersRemove returns the GlobalHeadersRemove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetGlobalHeadersRemove() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GlobalHeadersRemove
}

// GetGlobalHeadersRemoveOk returns a tuple with the GlobalHeadersRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetGlobalHeadersRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.GlobalHeadersRemove) {
		return nil, false
	}
	return o.GlobalHeadersRemove, true
}

// HasGlobalHeadersRemove returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalHeadersRemove() bool {
	if o != nil && !IsNil(o.GlobalHeadersRemove) {
		return true
	}

	return false
}

// SetGlobalHeadersRemove gets a reference to the given []string and assigns it to the GlobalHeadersRemove field.
func (o *VersionInfo) SetGlobalHeadersRemove(v []string) {
	o.GlobalHeadersRemove = v
}

// GetGlobalResponseHeaders returns the GlobalResponseHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetGlobalResponseHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.GlobalResponseHeaders
}

// GetGlobalResponseHeadersOk returns a tuple with the GlobalResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetGlobalResponseHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.GlobalResponseHeaders) {
		return nil, false
	}
	return &o.GlobalResponseHeaders, true
}

// HasGlobalResponseHeaders returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalResponseHeaders() bool {
	if o != nil && !IsNil(o.GlobalResponseHeaders) {
		return true
	}

	return false
}

// SetGlobalResponseHeaders gets a reference to the given map[string]string and assigns it to the GlobalResponseHeaders field.
func (o *VersionInfo) SetGlobalResponseHeaders(v map[string]string) {
	o.GlobalResponseHeaders = v
}

// GetGlobalResponseHeadersDisabled returns the GlobalResponseHeadersDisabled field value if set, zero value otherwise.
func (o *VersionInfo) GetGlobalResponseHeadersDisabled() bool {
	if o == nil || IsNil(o.GlobalResponseHeadersDisabled) {
		var ret bool
		return ret
	}
	return *o.GlobalResponseHeadersDisabled
}

// GetGlobalResponseHeadersDisabledOk returns a tuple with the GlobalResponseHeadersDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetGlobalResponseHeadersDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalResponseHeadersDisabled) {
		return nil, false
	}
	return o.GlobalResponseHeadersDisabled, true
}

// HasGlobalResponseHeadersDisabled returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalResponseHeadersDisabled() bool {
	if o != nil && !IsNil(o.GlobalResponseHeadersDisabled) {
		return true
	}

	return false
}

// SetGlobalResponseHeadersDisabled gets a reference to the given bool and assigns it to the GlobalResponseHeadersDisabled field.
func (o *VersionInfo) SetGlobalResponseHeadersDisabled(v bool) {
	o.GlobalResponseHeadersDisabled = &v
}

// GetGlobalResponseHeadersRemove returns the GlobalResponseHeadersRemove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetGlobalResponseHeadersRemove() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GlobalResponseHeadersRemove
}

// GetGlobalResponseHeadersRemoveOk returns a tuple with the GlobalResponseHeadersRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetGlobalResponseHeadersRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.GlobalResponseHeadersRemove) {
		return nil, false
	}
	return o.GlobalResponseHeadersRemove, true
}

// HasGlobalResponseHeadersRemove returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalResponseHeadersRemove() bool {
	if o != nil && !IsNil(o.GlobalResponseHeadersRemove) {
		return true
	}

	return false
}

// SetGlobalResponseHeadersRemove gets a reference to the given []string and assigns it to the GlobalResponseHeadersRemove field.
func (o *VersionInfo) SetGlobalResponseHeadersRemove(v []string) {
	o.GlobalResponseHeadersRemove = v
}

// GetGlobalSizeLimit returns the GlobalSizeLimit field value if set, zero value otherwise.
func (o *VersionInfo) GetGlobalSizeLimit() int64 {
	if o == nil || IsNil(o.GlobalSizeLimit) {
		var ret int64
		return ret
	}
	return *o.GlobalSizeLimit
}

// GetGlobalSizeLimitOk returns a tuple with the GlobalSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetGlobalSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalSizeLimit) {
		return nil, false
	}
	return o.GlobalSizeLimit, true
}

// HasGlobalSizeLimit returns a boolean if a field has been set.
func (o *VersionInfo) HasGlobalSizeLimit() bool {
	if o != nil && !IsNil(o.GlobalSizeLimit) {
		return true
	}

	return false
}

// SetGlobalSizeLimit gets a reference to the given int64 and assigns it to the GlobalSizeLimit field.
func (o *VersionInfo) SetGlobalSizeLimit(v int64) {
	o.GlobalSizeLimit = &v
}

// GetIgnoreEndpointCase returns the IgnoreEndpointCase field value if set, zero value otherwise.
func (o *VersionInfo) GetIgnoreEndpointCase() bool {
	if o == nil || IsNil(o.IgnoreEndpointCase) {
		var ret bool
		return ret
	}
	return *o.IgnoreEndpointCase
}

// GetIgnoreEndpointCaseOk returns a tuple with the IgnoreEndpointCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetIgnoreEndpointCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreEndpointCase) {
		return nil, false
	}
	return o.IgnoreEndpointCase, true
}

// HasIgnoreEndpointCase returns a boolean if a field has been set.
func (o *VersionInfo) HasIgnoreEndpointCase() bool {
	if o != nil && !IsNil(o.IgnoreEndpointCase) {
		return true
	}

	return false
}

// SetIgnoreEndpointCase gets a reference to the given bool and assigns it to the IgnoreEndpointCase field.
func (o *VersionInfo) SetIgnoreEndpointCase(v bool) {
	o.IgnoreEndpointCase = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VersionInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VersionInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VersionInfo) SetName(v string) {
	o.Name = &v
}

// GetOverrideTarget returns the OverrideTarget field value if set, zero value otherwise.
func (o *VersionInfo) GetOverrideTarget() string {
	if o == nil || IsNil(o.OverrideTarget) {
		var ret string
		return ret
	}
	return *o.OverrideTarget
}

// GetOverrideTargetOk returns a tuple with the OverrideTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetOverrideTargetOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideTarget) {
		return nil, false
	}
	return o.OverrideTarget, true
}

// HasOverrideTarget returns a boolean if a field has been set.
func (o *VersionInfo) HasOverrideTarget() bool {
	if o != nil && !IsNil(o.OverrideTarget) {
		return true
	}

	return false
}

// SetOverrideTarget gets a reference to the given string and assigns it to the OverrideTarget field.
func (o *VersionInfo) SetOverrideTarget(v string) {
	o.OverrideTarget = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *VersionInfo) GetPaths() VersionInfoPaths {
	if o == nil || IsNil(o.Paths) {
		var ret VersionInfoPaths
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetPathsOk() (*VersionInfoPaths, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *VersionInfo) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given VersionInfoPaths and assigns it to the Paths field.
func (o *VersionInfo) SetPaths(v VersionInfoPaths) {
	o.Paths = &v
}

// GetUseExtendedPaths returns the UseExtendedPaths field value if set, zero value otherwise.
func (o *VersionInfo) GetUseExtendedPaths() bool {
	if o == nil || IsNil(o.UseExtendedPaths) {
		var ret bool
		return ret
	}
	return *o.UseExtendedPaths
}

// GetUseExtendedPathsOk returns a tuple with the UseExtendedPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetUseExtendedPathsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseExtendedPaths) {
		return nil, false
	}
	return o.UseExtendedPaths, true
}

// HasUseExtendedPaths returns a boolean if a field has been set.
func (o *VersionInfo) HasUseExtendedPaths() bool {
	if o != nil && !IsNil(o.UseExtendedPaths) {
		return true
	}

	return false
}

// SetUseExtendedPaths gets a reference to the given bool and assigns it to the UseExtendedPaths field.
func (o *VersionInfo) SetUseExtendedPaths(v bool) {
	o.UseExtendedPaths = &v
}

func (o VersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !IsNil(o.ExtendedPaths) {
		toSerialize["extended_paths"] = o.ExtendedPaths
	}
	if o.GlobalHeaders != nil {
		toSerialize["global_headers"] = o.GlobalHeaders
	}
	if !IsNil(o.GlobalHeadersDisabled) {
		toSerialize["global_headers_disabled"] = o.GlobalHeadersDisabled
	}
	if o.GlobalHeadersRemove != nil {
		toSerialize["global_headers_remove"] = o.GlobalHeadersRemove
	}
	if o.GlobalResponseHeaders != nil {
		toSerialize["global_response_headers"] = o.GlobalResponseHeaders
	}
	if !IsNil(o.GlobalResponseHeadersDisabled) {
		toSerialize["global_response_headers_disabled"] = o.GlobalResponseHeadersDisabled
	}
	if o.GlobalResponseHeadersRemove != nil {
		toSerialize["global_response_headers_remove"] = o.GlobalResponseHeadersRemove
	}
	if !IsNil(o.GlobalSizeLimit) {
		toSerialize["global_size_limit"] = o.GlobalSizeLimit
	}
	if !IsNil(o.IgnoreEndpointCase) {
		toSerialize["ignore_endpoint_case"] = o.IgnoreEndpointCase
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OverrideTarget) {
		toSerialize["override_target"] = o.OverrideTarget
	}
	if !IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	if !IsNil(o.UseExtendedPaths) {
		toSerialize["use_extended_paths"] = o.UseExtendedPaths
	}
	return toSerialize, nil
}

type NullableVersionInfo struct {
	value *VersionInfo
	isSet bool
}

func (v NullableVersionInfo) Get() *VersionInfo {
	return v.value
}

func (v *NullableVersionInfo) Set(val *VersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfo(val *VersionInfo) *NullableVersionInfo {
	return &NullableVersionInfo{value: val, isSet: true}
}

func (v NullableVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
