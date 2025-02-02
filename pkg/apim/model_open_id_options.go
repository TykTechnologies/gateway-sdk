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

// checks if the OpenIDOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDOptions{}

// OpenIDOptions struct for OpenIDOptions
type OpenIDOptions struct {
	Providers         []OIDProviderConfig `json:"providers,omitempty"`
	SegregateByClient *bool               `json:"segregate_by_client,omitempty"`
}

// NewOpenIDOptions instantiates a new OpenIDOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDOptions() *OpenIDOptions {
	this := OpenIDOptions{}
	return &this
}

// NewOpenIDOptionsWithDefaults instantiates a new OpenIDOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDOptionsWithDefaults() *OpenIDOptions {
	this := OpenIDOptions{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenIDOptions) GetProviders() []OIDProviderConfig {
	if o == nil {
		var ret []OIDProviderConfig
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenIDOptions) GetProvidersOk() ([]OIDProviderConfig, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *OpenIDOptions) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []OIDProviderConfig and assigns it to the Providers field.
func (o *OpenIDOptions) SetProviders(v []OIDProviderConfig) {
	o.Providers = v
}

// GetSegregateByClient returns the SegregateByClient field value if set, zero value otherwise.
func (o *OpenIDOptions) GetSegregateByClient() bool {
	if o == nil || IsNil(o.SegregateByClient) {
		var ret bool
		return ret
	}
	return *o.SegregateByClient
}

// GetSegregateByClientOk returns a tuple with the SegregateByClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDOptions) GetSegregateByClientOk() (*bool, bool) {
	if o == nil || IsNil(o.SegregateByClient) {
		return nil, false
	}
	return o.SegregateByClient, true
}

// HasSegregateByClient returns a boolean if a field has been set.
func (o *OpenIDOptions) HasSegregateByClient() bool {
	if o != nil && !IsNil(o.SegregateByClient) {
		return true
	}

	return false
}

// SetSegregateByClient gets a reference to the given bool and assigns it to the SegregateByClient field.
func (o *OpenIDOptions) SetSegregateByClient(v bool) {
	o.SegregateByClient = &v
}

func (o OpenIDOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if !IsNil(o.SegregateByClient) {
		toSerialize["segregate_by_client"] = o.SegregateByClient
	}
	return toSerialize, nil
}

type NullableOpenIDOptions struct {
	value *OpenIDOptions
	isSet bool
}

func (v NullableOpenIDOptions) Get() *OpenIDOptions {
	return v.value
}

func (v *NullableOpenIDOptions) Set(val *OpenIDOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDOptions(val *OpenIDOptions) *NullableOpenIDOptions {
	return &NullableOpenIDOptions{value: val, isSet: true}
}

func (v NullableOpenIDOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
