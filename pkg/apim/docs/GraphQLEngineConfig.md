# GraphQLEngineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSources** | Pointer to [**[]GraphQLEngineDataSource**](GraphQLEngineDataSource.md) |  | [optional] 
**FieldConfigs** | Pointer to [**[]GraphQLFieldConfig**](GraphQLFieldConfig.md) |  | [optional] 
**GlobalHeaders** | Pointer to [**[]UDGGlobalHeader**](UDGGlobalHeader.md) |  | [optional] 

## Methods

### NewGraphQLEngineConfig

`func NewGraphQLEngineConfig() *GraphQLEngineConfig`

NewGraphQLEngineConfig instantiates a new GraphQLEngineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLEngineConfigWithDefaults

`func NewGraphQLEngineConfigWithDefaults() *GraphQLEngineConfig`

NewGraphQLEngineConfigWithDefaults instantiates a new GraphQLEngineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSources

`func (o *GraphQLEngineConfig) GetDataSources() []GraphQLEngineDataSource`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *GraphQLEngineConfig) GetDataSourcesOk() (*[]GraphQLEngineDataSource, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *GraphQLEngineConfig) SetDataSources(v []GraphQLEngineDataSource)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *GraphQLEngineConfig) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### SetDataSourcesNil

`func (o *GraphQLEngineConfig) SetDataSourcesNil(b bool)`

 SetDataSourcesNil sets the value for DataSources to be an explicit nil

### UnsetDataSources
`func (o *GraphQLEngineConfig) UnsetDataSources()`

UnsetDataSources ensures that no value is present for DataSources, not even an explicit nil
### GetFieldConfigs

`func (o *GraphQLEngineConfig) GetFieldConfigs() []GraphQLFieldConfig`

GetFieldConfigs returns the FieldConfigs field if non-nil, zero value otherwise.

### GetFieldConfigsOk

`func (o *GraphQLEngineConfig) GetFieldConfigsOk() (*[]GraphQLFieldConfig, bool)`

GetFieldConfigsOk returns a tuple with the FieldConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigs

`func (o *GraphQLEngineConfig) SetFieldConfigs(v []GraphQLFieldConfig)`

SetFieldConfigs sets FieldConfigs field to given value.

### HasFieldConfigs

`func (o *GraphQLEngineConfig) HasFieldConfigs() bool`

HasFieldConfigs returns a boolean if a field has been set.

### SetFieldConfigsNil

`func (o *GraphQLEngineConfig) SetFieldConfigsNil(b bool)`

 SetFieldConfigsNil sets the value for FieldConfigs to be an explicit nil

### UnsetFieldConfigs
`func (o *GraphQLEngineConfig) UnsetFieldConfigs()`

UnsetFieldConfigs ensures that no value is present for FieldConfigs, not even an explicit nil
### GetGlobalHeaders

`func (o *GraphQLEngineConfig) GetGlobalHeaders() []UDGGlobalHeader`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *GraphQLEngineConfig) GetGlobalHeadersOk() (*[]UDGGlobalHeader, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *GraphQLEngineConfig) SetGlobalHeaders(v []UDGGlobalHeader)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *GraphQLEngineConfig) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### SetGlobalHeadersNil

`func (o *GraphQLEngineConfig) SetGlobalHeadersNil(b bool)`

 SetGlobalHeadersNil sets the value for GlobalHeaders to be an explicit nil

### UnsetGlobalHeaders
`func (o *GraphQLEngineConfig) UnsetGlobalHeaders()`

UnsetGlobalHeaders ensures that no value is present for GlobalHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


