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

// checks if the OIDProviderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OIDProviderConfig{}

// OIDProviderConfig struct for OIDProviderConfig
type OIDProviderConfig struct {
	ClientIds map[string]string `json:"client_ids,omitempty"`
	Issuer    *string           `json:"issuer,omitempty"`
}

// NewOIDProviderConfig instantiates a new OIDProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDProviderConfig() *OIDProviderConfig {
	this := OIDProviderConfig{}
	return &this
}

// NewOIDProviderConfigWithDefaults instantiates a new OIDProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDProviderConfigWithDefaults() *OIDProviderConfig {
	this := OIDProviderConfig{}
	return &this
}

// GetClientIds returns the ClientIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OIDProviderConfig) GetClientIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ClientIds
}

// GetClientIdsOk returns a tuple with the ClientIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OIDProviderConfig) GetClientIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ClientIds) {
		return nil, false
	}
	return &o.ClientIds, true
}

// HasClientIds returns a boolean if a field has been set.
func (o *OIDProviderConfig) HasClientIds() bool {
	if o != nil && !IsNil(o.ClientIds) {
		return true
	}

	return false
}

// SetClientIds gets a reference to the given map[string]string and assigns it to the ClientIds field.
func (o *OIDProviderConfig) SetClientIds(v map[string]string) {
	o.ClientIds = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OIDProviderConfig) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDProviderConfig) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OIDProviderConfig) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OIDProviderConfig) SetIssuer(v string) {
	o.Issuer = &v
}

func (o OIDProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OIDProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientIds != nil {
		toSerialize["client_ids"] = o.ClientIds
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	return toSerialize, nil
}

type NullableOIDProviderConfig struct {
	value *OIDProviderConfig
	isSet bool
}

func (v NullableOIDProviderConfig) Get() *OIDProviderConfig {
	return v.value
}

func (v *NullableOIDProviderConfig) Set(val *OIDProviderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDProviderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDProviderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDProviderConfig(val *OIDProviderConfig) *NullableOIDProviderConfig {
	return &NullableOIDProviderConfig{value: val, isSet: true}
}

func (v NullableOIDProviderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDProviderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
