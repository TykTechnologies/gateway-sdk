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

// checks if the ProxyConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfig{}

// ProxyConfig struct for ProxyConfig
type ProxyConfig struct {
	CheckHostAgainstUptimeTests *bool                          `json:"check_host_against_uptime_tests,omitempty"`
	DisableStripSlash           *bool                          `json:"disable_strip_slash,omitempty"`
	EnableLoadBalancing         *bool                          `json:"enable_load_balancing,omitempty"`
	ListenPath                  *string                        `json:"listen_path,omitempty"`
	PreserveHostHeader          *bool                          `json:"preserve_host_header,omitempty"`
	ServiceDiscovery            *ServiceDiscoveryConfiguration `json:"service_discovery,omitempty"`
	StripListenPath             *bool                          `json:"strip_listen_path,omitempty"`
	TargetList                  []string                       `json:"target_list,omitempty"`
	TargetUrl                   *string                        `json:"target_url,omitempty"`
	Transport                   *ProxyConfigTransport          `json:"transport,omitempty"`
}

// NewProxyConfig instantiates a new ProxyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfig() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// NewProxyConfigWithDefaults instantiates a new ProxyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigWithDefaults() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field value if set, zero value otherwise.
func (o *ProxyConfig) GetCheckHostAgainstUptimeTests() bool {
	if o == nil || IsNil(o.CheckHostAgainstUptimeTests) {
		var ret bool
		return ret
	}
	return *o.CheckHostAgainstUptimeTests
}

// GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetCheckHostAgainstUptimeTestsOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckHostAgainstUptimeTests) {
		return nil, false
	}
	return o.CheckHostAgainstUptimeTests, true
}

// HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.
func (o *ProxyConfig) HasCheckHostAgainstUptimeTests() bool {
	if o != nil && !IsNil(o.CheckHostAgainstUptimeTests) {
		return true
	}

	return false
}

// SetCheckHostAgainstUptimeTests gets a reference to the given bool and assigns it to the CheckHostAgainstUptimeTests field.
func (o *ProxyConfig) SetCheckHostAgainstUptimeTests(v bool) {
	o.CheckHostAgainstUptimeTests = &v
}

// GetDisableStripSlash returns the DisableStripSlash field value if set, zero value otherwise.
func (o *ProxyConfig) GetDisableStripSlash() bool {
	if o == nil || IsNil(o.DisableStripSlash) {
		var ret bool
		return ret
	}
	return *o.DisableStripSlash
}

// GetDisableStripSlashOk returns a tuple with the DisableStripSlash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetDisableStripSlashOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableStripSlash) {
		return nil, false
	}
	return o.DisableStripSlash, true
}

// HasDisableStripSlash returns a boolean if a field has been set.
func (o *ProxyConfig) HasDisableStripSlash() bool {
	if o != nil && !IsNil(o.DisableStripSlash) {
		return true
	}

	return false
}

// SetDisableStripSlash gets a reference to the given bool and assigns it to the DisableStripSlash field.
func (o *ProxyConfig) SetDisableStripSlash(v bool) {
	o.DisableStripSlash = &v
}

// GetEnableLoadBalancing returns the EnableLoadBalancing field value if set, zero value otherwise.
func (o *ProxyConfig) GetEnableLoadBalancing() bool {
	if o == nil || IsNil(o.EnableLoadBalancing) {
		var ret bool
		return ret
	}
	return *o.EnableLoadBalancing
}

// GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetEnableLoadBalancingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLoadBalancing) {
		return nil, false
	}
	return o.EnableLoadBalancing, true
}

// HasEnableLoadBalancing returns a boolean if a field has been set.
func (o *ProxyConfig) HasEnableLoadBalancing() bool {
	if o != nil && !IsNil(o.EnableLoadBalancing) {
		return true
	}

	return false
}

// SetEnableLoadBalancing gets a reference to the given bool and assigns it to the EnableLoadBalancing field.
func (o *ProxyConfig) SetEnableLoadBalancing(v bool) {
	o.EnableLoadBalancing = &v
}

// GetListenPath returns the ListenPath field value if set, zero value otherwise.
func (o *ProxyConfig) GetListenPath() string {
	if o == nil || IsNil(o.ListenPath) {
		var ret string
		return ret
	}
	return *o.ListenPath
}

// GetListenPathOk returns a tuple with the ListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetListenPathOk() (*string, bool) {
	if o == nil || IsNil(o.ListenPath) {
		return nil, false
	}
	return o.ListenPath, true
}

// HasListenPath returns a boolean if a field has been set.
func (o *ProxyConfig) HasListenPath() bool {
	if o != nil && !IsNil(o.ListenPath) {
		return true
	}

	return false
}

// SetListenPath gets a reference to the given string and assigns it to the ListenPath field.
func (o *ProxyConfig) SetListenPath(v string) {
	o.ListenPath = &v
}

// GetPreserveHostHeader returns the PreserveHostHeader field value if set, zero value otherwise.
func (o *ProxyConfig) GetPreserveHostHeader() bool {
	if o == nil || IsNil(o.PreserveHostHeader) {
		var ret bool
		return ret
	}
	return *o.PreserveHostHeader
}

// GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetPreserveHostHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveHostHeader) {
		return nil, false
	}
	return o.PreserveHostHeader, true
}

// HasPreserveHostHeader returns a boolean if a field has been set.
func (o *ProxyConfig) HasPreserveHostHeader() bool {
	if o != nil && !IsNil(o.PreserveHostHeader) {
		return true
	}

	return false
}

// SetPreserveHostHeader gets a reference to the given bool and assigns it to the PreserveHostHeader field.
func (o *ProxyConfig) SetPreserveHostHeader(v bool) {
	o.PreserveHostHeader = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise.
func (o *ProxyConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret ServiceDiscoveryConfiguration
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}
	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *ProxyConfig) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given ServiceDiscoveryConfiguration and assigns it to the ServiceDiscovery field.
func (o *ProxyConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration) {
	o.ServiceDiscovery = &v
}

// GetStripListenPath returns the StripListenPath field value if set, zero value otherwise.
func (o *ProxyConfig) GetStripListenPath() bool {
	if o == nil || IsNil(o.StripListenPath) {
		var ret bool
		return ret
	}
	return *o.StripListenPath
}

// GetStripListenPathOk returns a tuple with the StripListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetStripListenPathOk() (*bool, bool) {
	if o == nil || IsNil(o.StripListenPath) {
		return nil, false
	}
	return o.StripListenPath, true
}

// HasStripListenPath returns a boolean if a field has been set.
func (o *ProxyConfig) HasStripListenPath() bool {
	if o != nil && !IsNil(o.StripListenPath) {
		return true
	}

	return false
}

// SetStripListenPath gets a reference to the given bool and assigns it to the StripListenPath field.
func (o *ProxyConfig) SetStripListenPath(v bool) {
	o.StripListenPath = &v
}

// GetTargetList returns the TargetList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfig) GetTargetList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetList
}

// GetTargetListOk returns a tuple with the TargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfig) GetTargetListOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetList) {
		return nil, false
	}
	return o.TargetList, true
}

// HasTargetList returns a boolean if a field has been set.
func (o *ProxyConfig) HasTargetList() bool {
	if o != nil && !IsNil(o.TargetList) {
		return true
	}

	return false
}

// SetTargetList gets a reference to the given []string and assigns it to the TargetList field.
func (o *ProxyConfig) SetTargetList(v []string) {
	o.TargetList = v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *ProxyConfig) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *ProxyConfig) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *ProxyConfig) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *ProxyConfig) GetTransport() ProxyConfigTransport {
	if o == nil || IsNil(o.Transport) {
		var ret ProxyConfigTransport
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetTransportOk() (*ProxyConfigTransport, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *ProxyConfig) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given ProxyConfigTransport and assigns it to the Transport field.
func (o *ProxyConfig) SetTransport(v ProxyConfigTransport) {
	o.Transport = &v
}

func (o ProxyConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckHostAgainstUptimeTests) {
		toSerialize["check_host_against_uptime_tests"] = o.CheckHostAgainstUptimeTests
	}
	if !IsNil(o.DisableStripSlash) {
		toSerialize["disable_strip_slash"] = o.DisableStripSlash
	}
	if !IsNil(o.EnableLoadBalancing) {
		toSerialize["enable_load_balancing"] = o.EnableLoadBalancing
	}
	if !IsNil(o.ListenPath) {
		toSerialize["listen_path"] = o.ListenPath
	}
	if !IsNil(o.PreserveHostHeader) {
		toSerialize["preserve_host_header"] = o.PreserveHostHeader
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["service_discovery"] = o.ServiceDiscovery
	}
	if !IsNil(o.StripListenPath) {
		toSerialize["strip_listen_path"] = o.StripListenPath
	}
	if o.TargetList != nil {
		toSerialize["target_list"] = o.TargetList
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["target_url"] = o.TargetUrl
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	return toSerialize, nil
}

type NullableProxyConfig struct {
	value *ProxyConfig
	isSet bool
}

func (v NullableProxyConfig) Get() *ProxyConfig {
	return v.value
}

func (v *NullableProxyConfig) Set(val *ProxyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfig(val *ProxyConfig) *NullableProxyConfig {
	return &NullableProxyConfig{value: val, isSet: true}
}

func (v NullableProxyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
