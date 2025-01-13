# GraphQLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Engine** | Pointer to [**GraphQLEngineConfig**](GraphQLEngineConfig.md) |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Introspection** | Pointer to [**GraphQLIntrospectionConfig**](GraphQLIntrospectionConfig.md) |  | [optional] 
**LastSchemaUpdate** | Pointer to **NullableTime** |  | [optional] 
**Playground** | Pointer to [**GraphQLPlayground**](GraphQLPlayground.md) |  | [optional] 
**Proxy** | Pointer to [**GraphQLProxyConfig**](GraphQLProxyConfig.md) |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Subgraph** | Pointer to [**GraphQLSubgraphConfig**](GraphQLSubgraphConfig.md) |  | [optional] 
**Supergraph** | Pointer to [**GraphQLSupergraphConfig**](GraphQLSupergraphConfig.md) |  | [optional] 
**TypeFieldConfigurations** | Pointer to [**[]DatasourceTypeFieldConfiguration**](DatasourceTypeFieldConfiguration.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphQLConfig

`func NewGraphQLConfig() *GraphQLConfig`

NewGraphQLConfig instantiates a new GraphQLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLConfigWithDefaults

`func NewGraphQLConfigWithDefaults() *GraphQLConfig`

NewGraphQLConfigWithDefaults instantiates a new GraphQLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GraphQLConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GraphQLConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GraphQLConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GraphQLConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngine

`func (o *GraphQLConfig) GetEngine() GraphQLEngineConfig`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *GraphQLConfig) GetEngineOk() (*GraphQLEngineConfig, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *GraphQLConfig) SetEngine(v GraphQLEngineConfig)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *GraphQLConfig) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetExecutionMode

`func (o *GraphQLConfig) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *GraphQLConfig) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *GraphQLConfig) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *GraphQLConfig) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetIntrospection

`func (o *GraphQLConfig) GetIntrospection() GraphQLIntrospectionConfig`

GetIntrospection returns the Introspection field if non-nil, zero value otherwise.

### GetIntrospectionOk

`func (o *GraphQLConfig) GetIntrospectionOk() (*GraphQLIntrospectionConfig, bool)`

GetIntrospectionOk returns a tuple with the Introspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospection

`func (o *GraphQLConfig) SetIntrospection(v GraphQLIntrospectionConfig)`

SetIntrospection sets Introspection field to given value.

### HasIntrospection

`func (o *GraphQLConfig) HasIntrospection() bool`

HasIntrospection returns a boolean if a field has been set.

### GetLastSchemaUpdate

`func (o *GraphQLConfig) GetLastSchemaUpdate() time.Time`

GetLastSchemaUpdate returns the LastSchemaUpdate field if non-nil, zero value otherwise.

### GetLastSchemaUpdateOk

`func (o *GraphQLConfig) GetLastSchemaUpdateOk() (*time.Time, bool)`

GetLastSchemaUpdateOk returns a tuple with the LastSchemaUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchemaUpdate

`func (o *GraphQLConfig) SetLastSchemaUpdate(v time.Time)`

SetLastSchemaUpdate sets LastSchemaUpdate field to given value.

### HasLastSchemaUpdate

`func (o *GraphQLConfig) HasLastSchemaUpdate() bool`

HasLastSchemaUpdate returns a boolean if a field has been set.

### SetLastSchemaUpdateNil

`func (o *GraphQLConfig) SetLastSchemaUpdateNil(b bool)`

 SetLastSchemaUpdateNil sets the value for LastSchemaUpdate to be an explicit nil

### UnsetLastSchemaUpdate
`func (o *GraphQLConfig) UnsetLastSchemaUpdate()`

UnsetLastSchemaUpdate ensures that no value is present for LastSchemaUpdate, not even an explicit nil
### GetPlayground

`func (o *GraphQLConfig) GetPlayground() GraphQLPlayground`

GetPlayground returns the Playground field if non-nil, zero value otherwise.

### GetPlaygroundOk

`func (o *GraphQLConfig) GetPlaygroundOk() (*GraphQLPlayground, bool)`

GetPlaygroundOk returns a tuple with the Playground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayground

`func (o *GraphQLConfig) SetPlayground(v GraphQLPlayground)`

SetPlayground sets Playground field to given value.

### HasPlayground

`func (o *GraphQLConfig) HasPlayground() bool`

HasPlayground returns a boolean if a field has been set.

### GetProxy

`func (o *GraphQLConfig) GetProxy() GraphQLProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GraphQLConfig) GetProxyOk() (*GraphQLProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GraphQLConfig) SetProxy(v GraphQLProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GraphQLConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSchema

`func (o *GraphQLConfig) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GraphQLConfig) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GraphQLConfig) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GraphQLConfig) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSubgraph

`func (o *GraphQLConfig) GetSubgraph() GraphQLSubgraphConfig`

GetSubgraph returns the Subgraph field if non-nil, zero value otherwise.

### GetSubgraphOk

`func (o *GraphQLConfig) GetSubgraphOk() (*GraphQLSubgraphConfig, bool)`

GetSubgraphOk returns a tuple with the Subgraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgraph

`func (o *GraphQLConfig) SetSubgraph(v GraphQLSubgraphConfig)`

SetSubgraph sets Subgraph field to given value.

### HasSubgraph

`func (o *GraphQLConfig) HasSubgraph() bool`

HasSubgraph returns a boolean if a field has been set.

### GetSupergraph

`func (o *GraphQLConfig) GetSupergraph() GraphQLSupergraphConfig`

GetSupergraph returns the Supergraph field if non-nil, zero value otherwise.

### GetSupergraphOk

`func (o *GraphQLConfig) GetSupergraphOk() (*GraphQLSupergraphConfig, bool)`

GetSupergraphOk returns a tuple with the Supergraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupergraph

`func (o *GraphQLConfig) SetSupergraph(v GraphQLSupergraphConfig)`

SetSupergraph sets Supergraph field to given value.

### HasSupergraph

`func (o *GraphQLConfig) HasSupergraph() bool`

HasSupergraph returns a boolean if a field has been set.

### GetTypeFieldConfigurations

`func (o *GraphQLConfig) GetTypeFieldConfigurations() []DatasourceTypeFieldConfiguration`

GetTypeFieldConfigurations returns the TypeFieldConfigurations field if non-nil, zero value otherwise.

### GetTypeFieldConfigurationsOk

`func (o *GraphQLConfig) GetTypeFieldConfigurationsOk() (*[]DatasourceTypeFieldConfiguration, bool)`

GetTypeFieldConfigurationsOk returns a tuple with the TypeFieldConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFieldConfigurations

`func (o *GraphQLConfig) SetTypeFieldConfigurations(v []DatasourceTypeFieldConfiguration)`

SetTypeFieldConfigurations sets TypeFieldConfigurations field to given value.

### HasTypeFieldConfigurations

`func (o *GraphQLConfig) HasTypeFieldConfigurations() bool`

HasTypeFieldConfigurations returns a boolean if a field has been set.

### SetTypeFieldConfigurationsNil

`func (o *GraphQLConfig) SetTypeFieldConfigurationsNil(b bool)`

 SetTypeFieldConfigurationsNil sets the value for TypeFieldConfigurations to be an explicit nil

### UnsetTypeFieldConfigurations
`func (o *GraphQLConfig) UnsetTypeFieldConfigurations()`

UnsetTypeFieldConfigurations ensures that no value is present for TypeFieldConfigurations, not even an explicit nil
### GetVersion

`func (o *GraphQLConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GraphQLConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GraphQLConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GraphQLConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


