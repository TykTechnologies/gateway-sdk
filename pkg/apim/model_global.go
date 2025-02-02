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

// checks if the Global type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Global{}

// Global struct for Global
type Global struct {
	Cache                     *Cache                    `json:"cache,omitempty"`
	ContextVariables          *ContextVariables         `json:"contextVariables,omitempty"`
	Cors                      *CORS                     `json:"cors,omitempty"`
	PluginConfig              *PluginConfig             `json:"pluginConfig,omitempty"`
	PostAuthenticationPlugin  *PostAuthenticationPlugin `json:"postAuthenticationPlugin,omitempty"`
	PostAuthenticationPlugins []CustomPlugin            `json:"postAuthenticationPlugins,omitempty"`
	PostPlugin                *PostPlugin               `json:"postPlugin,omitempty"`
	PostPlugins               []CustomPlugin            `json:"postPlugins,omitempty"`
	PrePlugin                 *PrePlugin                `json:"prePlugin,omitempty"`
	PrePlugins                []CustomPlugin            `json:"prePlugins,omitempty"`
	ResponsePlugin            *ResponsePlugin           `json:"responsePlugin,omitempty"`
	ResponsePlugins           []CustomPlugin            `json:"responsePlugins,omitempty"`
	TrafficLogs               *TrafficLogs              `json:"trafficLogs,omitempty"`
	TransformRequestHeaders   *TransformHeaders         `json:"transformRequestHeaders,omitempty"`
	TransformResponseHeaders  *TransformHeaders         `json:"transformResponseHeaders,omitempty"`
}

// NewGlobal instantiates a new Global object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobal() *Global {
	this := Global{}
	return &this
}

// NewGlobalWithDefaults instantiates a new Global object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalWithDefaults() *Global {
	this := Global{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *Global) GetCache() Cache {
	if o == nil || IsNil(o.Cache) {
		var ret Cache
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetCacheOk() (*Cache, bool) {
	if o == nil || IsNil(o.Cache) {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *Global) HasCache() bool {
	if o != nil && !IsNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given Cache and assigns it to the Cache field.
func (o *Global) SetCache(v Cache) {
	o.Cache = &v
}

// GetContextVariables returns the ContextVariables field value if set, zero value otherwise.
func (o *Global) GetContextVariables() ContextVariables {
	if o == nil || IsNil(o.ContextVariables) {
		var ret ContextVariables
		return ret
	}
	return *o.ContextVariables
}

// GetContextVariablesOk returns a tuple with the ContextVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetContextVariablesOk() (*ContextVariables, bool) {
	if o == nil || IsNil(o.ContextVariables) {
		return nil, false
	}
	return o.ContextVariables, true
}

// HasContextVariables returns a boolean if a field has been set.
func (o *Global) HasContextVariables() bool {
	if o != nil && !IsNil(o.ContextVariables) {
		return true
	}

	return false
}

// SetContextVariables gets a reference to the given ContextVariables and assigns it to the ContextVariables field.
func (o *Global) SetContextVariables(v ContextVariables) {
	o.ContextVariables = &v
}

// GetCors returns the Cors field value if set, zero value otherwise.
func (o *Global) GetCors() CORS {
	if o == nil || IsNil(o.Cors) {
		var ret CORS
		return ret
	}
	return *o.Cors
}

// GetCorsOk returns a tuple with the Cors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetCorsOk() (*CORS, bool) {
	if o == nil || IsNil(o.Cors) {
		return nil, false
	}
	return o.Cors, true
}

// HasCors returns a boolean if a field has been set.
func (o *Global) HasCors() bool {
	if o != nil && !IsNil(o.Cors) {
		return true
	}

	return false
}

// SetCors gets a reference to the given CORS and assigns it to the Cors field.
func (o *Global) SetCors(v CORS) {
	o.Cors = &v
}

// GetPluginConfig returns the PluginConfig field value if set, zero value otherwise.
func (o *Global) GetPluginConfig() PluginConfig {
	if o == nil || IsNil(o.PluginConfig) {
		var ret PluginConfig
		return ret
	}
	return *o.PluginConfig
}

// GetPluginConfigOk returns a tuple with the PluginConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPluginConfigOk() (*PluginConfig, bool) {
	if o == nil || IsNil(o.PluginConfig) {
		return nil, false
	}
	return o.PluginConfig, true
}

// HasPluginConfig returns a boolean if a field has been set.
func (o *Global) HasPluginConfig() bool {
	if o != nil && !IsNil(o.PluginConfig) {
		return true
	}

	return false
}

// SetPluginConfig gets a reference to the given PluginConfig and assigns it to the PluginConfig field.
func (o *Global) SetPluginConfig(v PluginConfig) {
	o.PluginConfig = &v
}

// GetPostAuthenticationPlugin returns the PostAuthenticationPlugin field value if set, zero value otherwise.
func (o *Global) GetPostAuthenticationPlugin() PostAuthenticationPlugin {
	if o == nil || IsNil(o.PostAuthenticationPlugin) {
		var ret PostAuthenticationPlugin
		return ret
	}
	return *o.PostAuthenticationPlugin
}

// GetPostAuthenticationPluginOk returns a tuple with the PostAuthenticationPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPostAuthenticationPluginOk() (*PostAuthenticationPlugin, bool) {
	if o == nil || IsNil(o.PostAuthenticationPlugin) {
		return nil, false
	}
	return o.PostAuthenticationPlugin, true
}

// HasPostAuthenticationPlugin returns a boolean if a field has been set.
func (o *Global) HasPostAuthenticationPlugin() bool {
	if o != nil && !IsNil(o.PostAuthenticationPlugin) {
		return true
	}

	return false
}

// SetPostAuthenticationPlugin gets a reference to the given PostAuthenticationPlugin and assigns it to the PostAuthenticationPlugin field.
func (o *Global) SetPostAuthenticationPlugin(v PostAuthenticationPlugin) {
	o.PostAuthenticationPlugin = &v
}

// GetPostAuthenticationPlugins returns the PostAuthenticationPlugins field value if set, zero value otherwise.
func (o *Global) GetPostAuthenticationPlugins() []CustomPlugin {
	if o == nil || IsNil(o.PostAuthenticationPlugins) {
		var ret []CustomPlugin
		return ret
	}
	return o.PostAuthenticationPlugins
}

// GetPostAuthenticationPluginsOk returns a tuple with the PostAuthenticationPlugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPostAuthenticationPluginsOk() ([]CustomPlugin, bool) {
	if o == nil || IsNil(o.PostAuthenticationPlugins) {
		return nil, false
	}
	return o.PostAuthenticationPlugins, true
}

// HasPostAuthenticationPlugins returns a boolean if a field has been set.
func (o *Global) HasPostAuthenticationPlugins() bool {
	if o != nil && !IsNil(o.PostAuthenticationPlugins) {
		return true
	}

	return false
}

// SetPostAuthenticationPlugins gets a reference to the given []CustomPlugin and assigns it to the PostAuthenticationPlugins field.
func (o *Global) SetPostAuthenticationPlugins(v []CustomPlugin) {
	o.PostAuthenticationPlugins = v
}

// GetPostPlugin returns the PostPlugin field value if set, zero value otherwise.
func (o *Global) GetPostPlugin() PostPlugin {
	if o == nil || IsNil(o.PostPlugin) {
		var ret PostPlugin
		return ret
	}
	return *o.PostPlugin
}

// GetPostPluginOk returns a tuple with the PostPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPostPluginOk() (*PostPlugin, bool) {
	if o == nil || IsNil(o.PostPlugin) {
		return nil, false
	}
	return o.PostPlugin, true
}

// HasPostPlugin returns a boolean if a field has been set.
func (o *Global) HasPostPlugin() bool {
	if o != nil && !IsNil(o.PostPlugin) {
		return true
	}

	return false
}

// SetPostPlugin gets a reference to the given PostPlugin and assigns it to the PostPlugin field.
func (o *Global) SetPostPlugin(v PostPlugin) {
	o.PostPlugin = &v
}

// GetPostPlugins returns the PostPlugins field value if set, zero value otherwise.
func (o *Global) GetPostPlugins() []CustomPlugin {
	if o == nil || IsNil(o.PostPlugins) {
		var ret []CustomPlugin
		return ret
	}
	return o.PostPlugins
}

// GetPostPluginsOk returns a tuple with the PostPlugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPostPluginsOk() ([]CustomPlugin, bool) {
	if o == nil || IsNil(o.PostPlugins) {
		return nil, false
	}
	return o.PostPlugins, true
}

// HasPostPlugins returns a boolean if a field has been set.
func (o *Global) HasPostPlugins() bool {
	if o != nil && !IsNil(o.PostPlugins) {
		return true
	}

	return false
}

// SetPostPlugins gets a reference to the given []CustomPlugin and assigns it to the PostPlugins field.
func (o *Global) SetPostPlugins(v []CustomPlugin) {
	o.PostPlugins = v
}

// GetPrePlugin returns the PrePlugin field value if set, zero value otherwise.
func (o *Global) GetPrePlugin() PrePlugin {
	if o == nil || IsNil(o.PrePlugin) {
		var ret PrePlugin
		return ret
	}
	return *o.PrePlugin
}

// GetPrePluginOk returns a tuple with the PrePlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPrePluginOk() (*PrePlugin, bool) {
	if o == nil || IsNil(o.PrePlugin) {
		return nil, false
	}
	return o.PrePlugin, true
}

// HasPrePlugin returns a boolean if a field has been set.
func (o *Global) HasPrePlugin() bool {
	if o != nil && !IsNil(o.PrePlugin) {
		return true
	}

	return false
}

// SetPrePlugin gets a reference to the given PrePlugin and assigns it to the PrePlugin field.
func (o *Global) SetPrePlugin(v PrePlugin) {
	o.PrePlugin = &v
}

// GetPrePlugins returns the PrePlugins field value if set, zero value otherwise.
func (o *Global) GetPrePlugins() []CustomPlugin {
	if o == nil || IsNil(o.PrePlugins) {
		var ret []CustomPlugin
		return ret
	}
	return o.PrePlugins
}

// GetPrePluginsOk returns a tuple with the PrePlugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetPrePluginsOk() ([]CustomPlugin, bool) {
	if o == nil || IsNil(o.PrePlugins) {
		return nil, false
	}
	return o.PrePlugins, true
}

// HasPrePlugins returns a boolean if a field has been set.
func (o *Global) HasPrePlugins() bool {
	if o != nil && !IsNil(o.PrePlugins) {
		return true
	}

	return false
}

// SetPrePlugins gets a reference to the given []CustomPlugin and assigns it to the PrePlugins field.
func (o *Global) SetPrePlugins(v []CustomPlugin) {
	o.PrePlugins = v
}

// GetResponsePlugin returns the ResponsePlugin field value if set, zero value otherwise.
func (o *Global) GetResponsePlugin() ResponsePlugin {
	if o == nil || IsNil(o.ResponsePlugin) {
		var ret ResponsePlugin
		return ret
	}
	return *o.ResponsePlugin
}

// GetResponsePluginOk returns a tuple with the ResponsePlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetResponsePluginOk() (*ResponsePlugin, bool) {
	if o == nil || IsNil(o.ResponsePlugin) {
		return nil, false
	}
	return o.ResponsePlugin, true
}

// HasResponsePlugin returns a boolean if a field has been set.
func (o *Global) HasResponsePlugin() bool {
	if o != nil && !IsNil(o.ResponsePlugin) {
		return true
	}

	return false
}

// SetResponsePlugin gets a reference to the given ResponsePlugin and assigns it to the ResponsePlugin field.
func (o *Global) SetResponsePlugin(v ResponsePlugin) {
	o.ResponsePlugin = &v
}

// GetResponsePlugins returns the ResponsePlugins field value if set, zero value otherwise.
func (o *Global) GetResponsePlugins() []CustomPlugin {
	if o == nil || IsNil(o.ResponsePlugins) {
		var ret []CustomPlugin
		return ret
	}
	return o.ResponsePlugins
}

// GetResponsePluginsOk returns a tuple with the ResponsePlugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetResponsePluginsOk() ([]CustomPlugin, bool) {
	if o == nil || IsNil(o.ResponsePlugins) {
		return nil, false
	}
	return o.ResponsePlugins, true
}

// HasResponsePlugins returns a boolean if a field has been set.
func (o *Global) HasResponsePlugins() bool {
	if o != nil && !IsNil(o.ResponsePlugins) {
		return true
	}

	return false
}

// SetResponsePlugins gets a reference to the given []CustomPlugin and assigns it to the ResponsePlugins field.
func (o *Global) SetResponsePlugins(v []CustomPlugin) {
	o.ResponsePlugins = v
}

// GetTrafficLogs returns the TrafficLogs field value if set, zero value otherwise.
func (o *Global) GetTrafficLogs() TrafficLogs {
	if o == nil || IsNil(o.TrafficLogs) {
		var ret TrafficLogs
		return ret
	}
	return *o.TrafficLogs
}

// GetTrafficLogsOk returns a tuple with the TrafficLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetTrafficLogsOk() (*TrafficLogs, bool) {
	if o == nil || IsNil(o.TrafficLogs) {
		return nil, false
	}
	return o.TrafficLogs, true
}

// HasTrafficLogs returns a boolean if a field has been set.
func (o *Global) HasTrafficLogs() bool {
	if o != nil && !IsNil(o.TrafficLogs) {
		return true
	}

	return false
}

// SetTrafficLogs gets a reference to the given TrafficLogs and assigns it to the TrafficLogs field.
func (o *Global) SetTrafficLogs(v TrafficLogs) {
	o.TrafficLogs = &v
}

// GetTransformRequestHeaders returns the TransformRequestHeaders field value if set, zero value otherwise.
func (o *Global) GetTransformRequestHeaders() TransformHeaders {
	if o == nil || IsNil(o.TransformRequestHeaders) {
		var ret TransformHeaders
		return ret
	}
	return *o.TransformRequestHeaders
}

// GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetTransformRequestHeadersOk() (*TransformHeaders, bool) {
	if o == nil || IsNil(o.TransformRequestHeaders) {
		return nil, false
	}
	return o.TransformRequestHeaders, true
}

// HasTransformRequestHeaders returns a boolean if a field has been set.
func (o *Global) HasTransformRequestHeaders() bool {
	if o != nil && !IsNil(o.TransformRequestHeaders) {
		return true
	}

	return false
}

// SetTransformRequestHeaders gets a reference to the given TransformHeaders and assigns it to the TransformRequestHeaders field.
func (o *Global) SetTransformRequestHeaders(v TransformHeaders) {
	o.TransformRequestHeaders = &v
}

// GetTransformResponseHeaders returns the TransformResponseHeaders field value if set, zero value otherwise.
func (o *Global) GetTransformResponseHeaders() TransformHeaders {
	if o == nil || IsNil(o.TransformResponseHeaders) {
		var ret TransformHeaders
		return ret
	}
	return *o.TransformResponseHeaders
}

// GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetTransformResponseHeadersOk() (*TransformHeaders, bool) {
	if o == nil || IsNil(o.TransformResponseHeaders) {
		return nil, false
	}
	return o.TransformResponseHeaders, true
}

// HasTransformResponseHeaders returns a boolean if a field has been set.
func (o *Global) HasTransformResponseHeaders() bool {
	if o != nil && !IsNil(o.TransformResponseHeaders) {
		return true
	}

	return false
}

// SetTransformResponseHeaders gets a reference to the given TransformHeaders and assigns it to the TransformResponseHeaders field.
func (o *Global) SetTransformResponseHeaders(v TransformHeaders) {
	o.TransformResponseHeaders = &v
}

func (o Global) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Global) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cache) {
		toSerialize["cache"] = o.Cache
	}
	if !IsNil(o.ContextVariables) {
		toSerialize["contextVariables"] = o.ContextVariables
	}
	if !IsNil(o.Cors) {
		toSerialize["cors"] = o.Cors
	}
	if !IsNil(o.PluginConfig) {
		toSerialize["pluginConfig"] = o.PluginConfig
	}
	if !IsNil(o.PostAuthenticationPlugin) {
		toSerialize["postAuthenticationPlugin"] = o.PostAuthenticationPlugin
	}
	if !IsNil(o.PostAuthenticationPlugins) {
		toSerialize["postAuthenticationPlugins"] = o.PostAuthenticationPlugins
	}
	if !IsNil(o.PostPlugin) {
		toSerialize["postPlugin"] = o.PostPlugin
	}
	if !IsNil(o.PostPlugins) {
		toSerialize["postPlugins"] = o.PostPlugins
	}
	if !IsNil(o.PrePlugin) {
		toSerialize["prePlugin"] = o.PrePlugin
	}
	if !IsNil(o.PrePlugins) {
		toSerialize["prePlugins"] = o.PrePlugins
	}
	if !IsNil(o.ResponsePlugin) {
		toSerialize["responsePlugin"] = o.ResponsePlugin
	}
	if !IsNil(o.ResponsePlugins) {
		toSerialize["responsePlugins"] = o.ResponsePlugins
	}
	if !IsNil(o.TrafficLogs) {
		toSerialize["trafficLogs"] = o.TrafficLogs
	}
	if !IsNil(o.TransformRequestHeaders) {
		toSerialize["transformRequestHeaders"] = o.TransformRequestHeaders
	}
	if !IsNil(o.TransformResponseHeaders) {
		toSerialize["transformResponseHeaders"] = o.TransformResponseHeaders
	}
	return toSerialize, nil
}

type NullableGlobal struct {
	value *Global
	isSet bool
}

func (v NullableGlobal) Get() *Global {
	return v.value
}

func (v *NullableGlobal) Set(val *Global) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobal) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobal(val *Global) *NullableGlobal {
	return &NullableGlobal{value: val, isSet: true}
}

func (v NullableGlobal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
