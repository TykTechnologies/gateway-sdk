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

// checks if the ProxyConfigTransport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfigTransport{}

// ProxyConfigTransport struct for ProxyConfigTransport
type ProxyConfigTransport struct {
	ProxyUrl                *string  `json:"proxy_url,omitempty"`
	SslCiphers              []string `json:"ssl_ciphers,omitempty"`
	SslForceCommonNameCheck *bool    `json:"ssl_force_common_name_check,omitempty"`
	SslInsecureSkipVerify   *bool    `json:"ssl_insecure_skip_verify,omitempty"`
	SslMaxVersion           *int32   `json:"ssl_max_version,omitempty"`
	SslMinVersion           *int32   `json:"ssl_min_version,omitempty"`
}

// NewProxyConfigTransport instantiates a new ProxyConfigTransport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfigTransport() *ProxyConfigTransport {
	this := ProxyConfigTransport{}
	return &this
}

// NewProxyConfigTransportWithDefaults instantiates a new ProxyConfigTransport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigTransportWithDefaults() *ProxyConfigTransport {
	this := ProxyConfigTransport{}
	return &this
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise.
func (o *ProxyConfigTransport) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl) {
		var ret string
		return ret
	}
	return *o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfigTransport) GetProxyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyUrl) {
		return nil, false
	}
	return o.ProxyUrl, true
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasProxyUrl() bool {
	if o != nil && !IsNil(o.ProxyUrl) {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given string and assigns it to the ProxyUrl field.
func (o *ProxyConfigTransport) SetProxyUrl(v string) {
	o.ProxyUrl = &v
}

// GetSslCiphers returns the SslCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfigTransport) GetSslCiphers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SslCiphers
}

// GetSslCiphersOk returns a tuple with the SslCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfigTransport) GetSslCiphersOk() ([]string, bool) {
	if o == nil || IsNil(o.SslCiphers) {
		return nil, false
	}
	return o.SslCiphers, true
}

// HasSslCiphers returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasSslCiphers() bool {
	if o != nil && !IsNil(o.SslCiphers) {
		return true
	}

	return false
}

// SetSslCiphers gets a reference to the given []string and assigns it to the SslCiphers field.
func (o *ProxyConfigTransport) SetSslCiphers(v []string) {
	o.SslCiphers = v
}

// GetSslForceCommonNameCheck returns the SslForceCommonNameCheck field value if set, zero value otherwise.
func (o *ProxyConfigTransport) GetSslForceCommonNameCheck() bool {
	if o == nil || IsNil(o.SslForceCommonNameCheck) {
		var ret bool
		return ret
	}
	return *o.SslForceCommonNameCheck
}

// GetSslForceCommonNameCheckOk returns a tuple with the SslForceCommonNameCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfigTransport) GetSslForceCommonNameCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.SslForceCommonNameCheck) {
		return nil, false
	}
	return o.SslForceCommonNameCheck, true
}

// HasSslForceCommonNameCheck returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasSslForceCommonNameCheck() bool {
	if o != nil && !IsNil(o.SslForceCommonNameCheck) {
		return true
	}

	return false
}

// SetSslForceCommonNameCheck gets a reference to the given bool and assigns it to the SslForceCommonNameCheck field.
func (o *ProxyConfigTransport) SetSslForceCommonNameCheck(v bool) {
	o.SslForceCommonNameCheck = &v
}

// GetSslInsecureSkipVerify returns the SslInsecureSkipVerify field value if set, zero value otherwise.
func (o *ProxyConfigTransport) GetSslInsecureSkipVerify() bool {
	if o == nil || IsNil(o.SslInsecureSkipVerify) {
		var ret bool
		return ret
	}
	return *o.SslInsecureSkipVerify
}

// GetSslInsecureSkipVerifyOk returns a tuple with the SslInsecureSkipVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfigTransport) GetSslInsecureSkipVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.SslInsecureSkipVerify) {
		return nil, false
	}
	return o.SslInsecureSkipVerify, true
}

// HasSslInsecureSkipVerify returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasSslInsecureSkipVerify() bool {
	if o != nil && !IsNil(o.SslInsecureSkipVerify) {
		return true
	}

	return false
}

// SetSslInsecureSkipVerify gets a reference to the given bool and assigns it to the SslInsecureSkipVerify field.
func (o *ProxyConfigTransport) SetSslInsecureSkipVerify(v bool) {
	o.SslInsecureSkipVerify = &v
}

// GetSslMaxVersion returns the SslMaxVersion field value if set, zero value otherwise.
func (o *ProxyConfigTransport) GetSslMaxVersion() int32 {
	if o == nil || IsNil(o.SslMaxVersion) {
		var ret int32
		return ret
	}
	return *o.SslMaxVersion
}

// GetSslMaxVersionOk returns a tuple with the SslMaxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfigTransport) GetSslMaxVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SslMaxVersion) {
		return nil, false
	}
	return o.SslMaxVersion, true
}

// HasSslMaxVersion returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasSslMaxVersion() bool {
	if o != nil && !IsNil(o.SslMaxVersion) {
		return true
	}

	return false
}

// SetSslMaxVersion gets a reference to the given int32 and assigns it to the SslMaxVersion field.
func (o *ProxyConfigTransport) SetSslMaxVersion(v int32) {
	o.SslMaxVersion = &v
}

// GetSslMinVersion returns the SslMinVersion field value if set, zero value otherwise.
func (o *ProxyConfigTransport) GetSslMinVersion() int32 {
	if o == nil || IsNil(o.SslMinVersion) {
		var ret int32
		return ret
	}
	return *o.SslMinVersion
}

// GetSslMinVersionOk returns a tuple with the SslMinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfigTransport) GetSslMinVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SslMinVersion) {
		return nil, false
	}
	return o.SslMinVersion, true
}

// HasSslMinVersion returns a boolean if a field has been set.
func (o *ProxyConfigTransport) HasSslMinVersion() bool {
	if o != nil && !IsNil(o.SslMinVersion) {
		return true
	}

	return false
}

// SetSslMinVersion gets a reference to the given int32 and assigns it to the SslMinVersion field.
func (o *ProxyConfigTransport) SetSslMinVersion(v int32) {
	o.SslMinVersion = &v
}

func (o ProxyConfigTransport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfigTransport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProxyUrl) {
		toSerialize["proxy_url"] = o.ProxyUrl
	}
	if o.SslCiphers != nil {
		toSerialize["ssl_ciphers"] = o.SslCiphers
	}
	if !IsNil(o.SslForceCommonNameCheck) {
		toSerialize["ssl_force_common_name_check"] = o.SslForceCommonNameCheck
	}
	if !IsNil(o.SslInsecureSkipVerify) {
		toSerialize["ssl_insecure_skip_verify"] = o.SslInsecureSkipVerify
	}
	if !IsNil(o.SslMaxVersion) {
		toSerialize["ssl_max_version"] = o.SslMaxVersion
	}
	if !IsNil(o.SslMinVersion) {
		toSerialize["ssl_min_version"] = o.SslMinVersion
	}
	return toSerialize, nil
}

type NullableProxyConfigTransport struct {
	value *ProxyConfigTransport
	isSet bool
}

func (v NullableProxyConfigTransport) Get() *ProxyConfigTransport {
	return v.value
}

func (v *NullableProxyConfigTransport) Set(val *ProxyConfigTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfigTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfigTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfigTransport(val *ProxyConfigTransport) *NullableProxyConfigTransport {
	return &NullableProxyConfigTransport{value: val, isSet: true}
}

func (v NullableProxyConfigTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfigTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
