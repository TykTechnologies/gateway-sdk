# Global

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**Cache**](Cache.md) |  | [optional] 
**ContextVariables** | Pointer to [**ContextVariables**](ContextVariables.md) |  | [optional] 
**Cors** | Pointer to [**CORS**](CORS.md) |  | [optional] 
**PluginConfig** | Pointer to [**PluginConfig**](PluginConfig.md) |  | [optional] 
**PostAuthenticationPlugin** | Pointer to [**PostAuthenticationPlugin**](PostAuthenticationPlugin.md) |  | [optional] 
**PostAuthenticationPlugins** | Pointer to [**[]CustomPlugin**](CustomPlugin.md) |  | [optional] 
**PostPlugin** | Pointer to [**PostPlugin**](PostPlugin.md) |  | [optional] 
**PostPlugins** | Pointer to [**[]CustomPlugin**](CustomPlugin.md) |  | [optional] 
**PrePlugin** | Pointer to [**PrePlugin**](PrePlugin.md) |  | [optional] 
**PrePlugins** | Pointer to [**[]CustomPlugin**](CustomPlugin.md) |  | [optional] 
**ResponsePlugin** | Pointer to [**ResponsePlugin**](ResponsePlugin.md) |  | [optional] 
**ResponsePlugins** | Pointer to [**[]CustomPlugin**](CustomPlugin.md) |  | [optional] 
**TrafficLogs** | Pointer to [**TrafficLogs**](TrafficLogs.md) |  | [optional] 
**TransformRequestHeaders** | Pointer to [**TransformHeaders**](TransformHeaders.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**TransformHeaders**](TransformHeaders.md) |  | [optional] 

## Methods

### NewGlobal

`func NewGlobal() *Global`

NewGlobal instantiates a new Global object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalWithDefaults

`func NewGlobalWithDefaults() *Global`

NewGlobalWithDefaults instantiates a new Global object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *Global) GetCache() Cache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *Global) GetCacheOk() (*Cache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *Global) SetCache(v Cache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *Global) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetContextVariables

`func (o *Global) GetContextVariables() ContextVariables`

GetContextVariables returns the ContextVariables field if non-nil, zero value otherwise.

### GetContextVariablesOk

`func (o *Global) GetContextVariablesOk() (*ContextVariables, bool)`

GetContextVariablesOk returns a tuple with the ContextVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextVariables

`func (o *Global) SetContextVariables(v ContextVariables)`

SetContextVariables sets ContextVariables field to given value.

### HasContextVariables

`func (o *Global) HasContextVariables() bool`

HasContextVariables returns a boolean if a field has been set.

### GetCors

`func (o *Global) GetCors() CORS`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *Global) GetCorsOk() (*CORS, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *Global) SetCors(v CORS)`

SetCors sets Cors field to given value.

### HasCors

`func (o *Global) HasCors() bool`

HasCors returns a boolean if a field has been set.

### GetPluginConfig

`func (o *Global) GetPluginConfig() PluginConfig`

GetPluginConfig returns the PluginConfig field if non-nil, zero value otherwise.

### GetPluginConfigOk

`func (o *Global) GetPluginConfigOk() (*PluginConfig, bool)`

GetPluginConfigOk returns a tuple with the PluginConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginConfig

`func (o *Global) SetPluginConfig(v PluginConfig)`

SetPluginConfig sets PluginConfig field to given value.

### HasPluginConfig

`func (o *Global) HasPluginConfig() bool`

HasPluginConfig returns a boolean if a field has been set.

### GetPostAuthenticationPlugin

`func (o *Global) GetPostAuthenticationPlugin() PostAuthenticationPlugin`

GetPostAuthenticationPlugin returns the PostAuthenticationPlugin field if non-nil, zero value otherwise.

### GetPostAuthenticationPluginOk

`func (o *Global) GetPostAuthenticationPluginOk() (*PostAuthenticationPlugin, bool)`

GetPostAuthenticationPluginOk returns a tuple with the PostAuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthenticationPlugin

`func (o *Global) SetPostAuthenticationPlugin(v PostAuthenticationPlugin)`

SetPostAuthenticationPlugin sets PostAuthenticationPlugin field to given value.

### HasPostAuthenticationPlugin

`func (o *Global) HasPostAuthenticationPlugin() bool`

HasPostAuthenticationPlugin returns a boolean if a field has been set.

### GetPostAuthenticationPlugins

`func (o *Global) GetPostAuthenticationPlugins() []CustomPlugin`

GetPostAuthenticationPlugins returns the PostAuthenticationPlugins field if non-nil, zero value otherwise.

### GetPostAuthenticationPluginsOk

`func (o *Global) GetPostAuthenticationPluginsOk() (*[]CustomPlugin, bool)`

GetPostAuthenticationPluginsOk returns a tuple with the PostAuthenticationPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthenticationPlugins

`func (o *Global) SetPostAuthenticationPlugins(v []CustomPlugin)`

SetPostAuthenticationPlugins sets PostAuthenticationPlugins field to given value.

### HasPostAuthenticationPlugins

`func (o *Global) HasPostAuthenticationPlugins() bool`

HasPostAuthenticationPlugins returns a boolean if a field has been set.

### GetPostPlugin

`func (o *Global) GetPostPlugin() PostPlugin`

GetPostPlugin returns the PostPlugin field if non-nil, zero value otherwise.

### GetPostPluginOk

`func (o *Global) GetPostPluginOk() (*PostPlugin, bool)`

GetPostPluginOk returns a tuple with the PostPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugin

`func (o *Global) SetPostPlugin(v PostPlugin)`

SetPostPlugin sets PostPlugin field to given value.

### HasPostPlugin

`func (o *Global) HasPostPlugin() bool`

HasPostPlugin returns a boolean if a field has been set.

### GetPostPlugins

`func (o *Global) GetPostPlugins() []CustomPlugin`

GetPostPlugins returns the PostPlugins field if non-nil, zero value otherwise.

### GetPostPluginsOk

`func (o *Global) GetPostPluginsOk() (*[]CustomPlugin, bool)`

GetPostPluginsOk returns a tuple with the PostPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugins

`func (o *Global) SetPostPlugins(v []CustomPlugin)`

SetPostPlugins sets PostPlugins field to given value.

### HasPostPlugins

`func (o *Global) HasPostPlugins() bool`

HasPostPlugins returns a boolean if a field has been set.

### GetPrePlugin

`func (o *Global) GetPrePlugin() PrePlugin`

GetPrePlugin returns the PrePlugin field if non-nil, zero value otherwise.

### GetPrePluginOk

`func (o *Global) GetPrePluginOk() (*PrePlugin, bool)`

GetPrePluginOk returns a tuple with the PrePlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePlugin

`func (o *Global) SetPrePlugin(v PrePlugin)`

SetPrePlugin sets PrePlugin field to given value.

### HasPrePlugin

`func (o *Global) HasPrePlugin() bool`

HasPrePlugin returns a boolean if a field has been set.

### GetPrePlugins

`func (o *Global) GetPrePlugins() []CustomPlugin`

GetPrePlugins returns the PrePlugins field if non-nil, zero value otherwise.

### GetPrePluginsOk

`func (o *Global) GetPrePluginsOk() (*[]CustomPlugin, bool)`

GetPrePluginsOk returns a tuple with the PrePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePlugins

`func (o *Global) SetPrePlugins(v []CustomPlugin)`

SetPrePlugins sets PrePlugins field to given value.

### HasPrePlugins

`func (o *Global) HasPrePlugins() bool`

HasPrePlugins returns a boolean if a field has been set.

### GetResponsePlugin

`func (o *Global) GetResponsePlugin() ResponsePlugin`

GetResponsePlugin returns the ResponsePlugin field if non-nil, zero value otherwise.

### GetResponsePluginOk

`func (o *Global) GetResponsePluginOk() (*ResponsePlugin, bool)`

GetResponsePluginOk returns a tuple with the ResponsePlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePlugin

`func (o *Global) SetResponsePlugin(v ResponsePlugin)`

SetResponsePlugin sets ResponsePlugin field to given value.

### HasResponsePlugin

`func (o *Global) HasResponsePlugin() bool`

HasResponsePlugin returns a boolean if a field has been set.

### GetResponsePlugins

`func (o *Global) GetResponsePlugins() []CustomPlugin`

GetResponsePlugins returns the ResponsePlugins field if non-nil, zero value otherwise.

### GetResponsePluginsOk

`func (o *Global) GetResponsePluginsOk() (*[]CustomPlugin, bool)`

GetResponsePluginsOk returns a tuple with the ResponsePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePlugins

`func (o *Global) SetResponsePlugins(v []CustomPlugin)`

SetResponsePlugins sets ResponsePlugins field to given value.

### HasResponsePlugins

`func (o *Global) HasResponsePlugins() bool`

HasResponsePlugins returns a boolean if a field has been set.

### GetTrafficLogs

`func (o *Global) GetTrafficLogs() TrafficLogs`

GetTrafficLogs returns the TrafficLogs field if non-nil, zero value otherwise.

### GetTrafficLogsOk

`func (o *Global) GetTrafficLogsOk() (*TrafficLogs, bool)`

GetTrafficLogsOk returns a tuple with the TrafficLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLogs

`func (o *Global) SetTrafficLogs(v TrafficLogs)`

SetTrafficLogs sets TrafficLogs field to given value.

### HasTrafficLogs

`func (o *Global) HasTrafficLogs() bool`

HasTrafficLogs returns a boolean if a field has been set.

### GetTransformRequestHeaders

`func (o *Global) GetTransformRequestHeaders() TransformHeaders`

GetTransformRequestHeaders returns the TransformRequestHeaders field if non-nil, zero value otherwise.

### GetTransformRequestHeadersOk

`func (o *Global) GetTransformRequestHeadersOk() (*TransformHeaders, bool)`

GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestHeaders

`func (o *Global) SetTransformRequestHeaders(v TransformHeaders)`

SetTransformRequestHeaders sets TransformRequestHeaders field to given value.

### HasTransformRequestHeaders

`func (o *Global) HasTransformRequestHeaders() bool`

HasTransformRequestHeaders returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *Global) GetTransformResponseHeaders() TransformHeaders`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *Global) GetTransformResponseHeadersOk() (*TransformHeaders, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *Global) SetTransformResponseHeaders(v TransformHeaders)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *Global) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


