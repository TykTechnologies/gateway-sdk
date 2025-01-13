# GraphQLSupergraphConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableQueryBatching** | Pointer to **bool** |  | [optional] 
**GlobalHeaders** | Pointer to **map[string]string** |  | [optional] 
**MergedSdl** | Pointer to **string** |  | [optional] 
**Subgraphs** | Pointer to [**[]GraphQLSubgraphEntity**](GraphQLSubgraphEntity.md) |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGraphQLSupergraphConfig

`func NewGraphQLSupergraphConfig() *GraphQLSupergraphConfig`

NewGraphQLSupergraphConfig instantiates a new GraphQLSupergraphConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLSupergraphConfigWithDefaults

`func NewGraphQLSupergraphConfigWithDefaults() *GraphQLSupergraphConfig`

NewGraphQLSupergraphConfigWithDefaults instantiates a new GraphQLSupergraphConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableQueryBatching

`func (o *GraphQLSupergraphConfig) GetDisableQueryBatching() bool`

GetDisableQueryBatching returns the DisableQueryBatching field if non-nil, zero value otherwise.

### GetDisableQueryBatchingOk

`func (o *GraphQLSupergraphConfig) GetDisableQueryBatchingOk() (*bool, bool)`

GetDisableQueryBatchingOk returns a tuple with the DisableQueryBatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableQueryBatching

`func (o *GraphQLSupergraphConfig) SetDisableQueryBatching(v bool)`

SetDisableQueryBatching sets DisableQueryBatching field to given value.

### HasDisableQueryBatching

`func (o *GraphQLSupergraphConfig) HasDisableQueryBatching() bool`

HasDisableQueryBatching returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *GraphQLSupergraphConfig) GetGlobalHeaders() map[string]string`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *GraphQLSupergraphConfig) GetGlobalHeadersOk() (*map[string]string, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *GraphQLSupergraphConfig) SetGlobalHeaders(v map[string]string)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *GraphQLSupergraphConfig) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### SetGlobalHeadersNil

`func (o *GraphQLSupergraphConfig) SetGlobalHeadersNil(b bool)`

 SetGlobalHeadersNil sets the value for GlobalHeaders to be an explicit nil

### UnsetGlobalHeaders
`func (o *GraphQLSupergraphConfig) UnsetGlobalHeaders()`

UnsetGlobalHeaders ensures that no value is present for GlobalHeaders, not even an explicit nil
### GetMergedSdl

`func (o *GraphQLSupergraphConfig) GetMergedSdl() string`

GetMergedSdl returns the MergedSdl field if non-nil, zero value otherwise.

### GetMergedSdlOk

`func (o *GraphQLSupergraphConfig) GetMergedSdlOk() (*string, bool)`

GetMergedSdlOk returns a tuple with the MergedSdl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedSdl

`func (o *GraphQLSupergraphConfig) SetMergedSdl(v string)`

SetMergedSdl sets MergedSdl field to given value.

### HasMergedSdl

`func (o *GraphQLSupergraphConfig) HasMergedSdl() bool`

HasMergedSdl returns a boolean if a field has been set.

### GetSubgraphs

`func (o *GraphQLSupergraphConfig) GetSubgraphs() []GraphQLSubgraphEntity`

GetSubgraphs returns the Subgraphs field if non-nil, zero value otherwise.

### GetSubgraphsOk

`func (o *GraphQLSupergraphConfig) GetSubgraphsOk() (*[]GraphQLSubgraphEntity, bool)`

GetSubgraphsOk returns a tuple with the Subgraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgraphs

`func (o *GraphQLSupergraphConfig) SetSubgraphs(v []GraphQLSubgraphEntity)`

SetSubgraphs sets Subgraphs field to given value.

### HasSubgraphs

`func (o *GraphQLSupergraphConfig) HasSubgraphs() bool`

HasSubgraphs returns a boolean if a field has been set.

### SetSubgraphsNil

`func (o *GraphQLSupergraphConfig) SetSubgraphsNil(b bool)`

 SetSubgraphsNil sets the value for Subgraphs to be an explicit nil

### UnsetSubgraphs
`func (o *GraphQLSupergraphConfig) UnsetSubgraphs()`

UnsetSubgraphs ensures that no value is present for Subgraphs, not even an explicit nil
### GetUpdatedAt

`func (o *GraphQLSupergraphConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GraphQLSupergraphConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GraphQLSupergraphConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GraphQLSupergraphConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GraphQLSupergraphConfig) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GraphQLSupergraphConfig) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


