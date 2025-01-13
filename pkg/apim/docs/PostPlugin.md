# PostPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]CustomPlugin**](CustomPlugin.md) |  | [optional] 

## Methods

### NewPostPlugin

`func NewPostPlugin() *PostPlugin`

NewPostPlugin instantiates a new PostPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPluginWithDefaults

`func NewPostPluginWithDefaults() *PostPlugin`

NewPostPluginWithDefaults instantiates a new PostPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *PostPlugin) GetPlugins() []CustomPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PostPlugin) GetPluginsOk() (*[]CustomPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PostPlugin) SetPlugins(v []CustomPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PostPlugin) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


