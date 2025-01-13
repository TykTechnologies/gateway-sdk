# GraphQLFieldConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableDefaultMapping** | Pointer to **bool** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **[]string** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphQLFieldConfig

`func NewGraphQLFieldConfig() *GraphQLFieldConfig`

NewGraphQLFieldConfig instantiates a new GraphQLFieldConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLFieldConfigWithDefaults

`func NewGraphQLFieldConfigWithDefaults() *GraphQLFieldConfig`

NewGraphQLFieldConfigWithDefaults instantiates a new GraphQLFieldConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableDefaultMapping

`func (o *GraphQLFieldConfig) GetDisableDefaultMapping() bool`

GetDisableDefaultMapping returns the DisableDefaultMapping field if non-nil, zero value otherwise.

### GetDisableDefaultMappingOk

`func (o *GraphQLFieldConfig) GetDisableDefaultMappingOk() (*bool, bool)`

GetDisableDefaultMappingOk returns a tuple with the DisableDefaultMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDefaultMapping

`func (o *GraphQLFieldConfig) SetDisableDefaultMapping(v bool)`

SetDisableDefaultMapping sets DisableDefaultMapping field to given value.

### HasDisableDefaultMapping

`func (o *GraphQLFieldConfig) HasDisableDefaultMapping() bool`

HasDisableDefaultMapping returns a boolean if a field has been set.

### GetFieldName

`func (o *GraphQLFieldConfig) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *GraphQLFieldConfig) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *GraphQLFieldConfig) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *GraphQLFieldConfig) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetPath

`func (o *GraphQLFieldConfig) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GraphQLFieldConfig) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GraphQLFieldConfig) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GraphQLFieldConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GraphQLFieldConfig) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GraphQLFieldConfig) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTypeName

`func (o *GraphQLFieldConfig) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *GraphQLFieldConfig) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *GraphQLFieldConfig) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *GraphQLFieldConfig) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


