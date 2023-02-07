# SchemaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | [**InfoModelModel**](InfoModel.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentationModelModel**](ExternalDocumentationModel.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerModelModel**](ServerModelModel.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]TagModelModel**](TagModelModel.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**ComponentsModelModel**](ComponentsModel.md) |  | [optional] 

## Methods

### NewSchemaModel

`func NewSchemaModel(openapi string, info InfoModelModel, paths map[string]interface{}, ) *SchemaModel`

NewSchemaModel instantiates a new SchemaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaModelWithDefaults

`func NewSchemaModelWithDefaults() *SchemaModel`

NewSchemaModelWithDefaults instantiates a new SchemaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *SchemaModel) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *SchemaModel) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *SchemaModel) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *SchemaModel) GetInfo() InfoModelModel`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SchemaModel) GetInfoOk() (*InfoModelModel, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SchemaModel) SetInfo(v InfoModelModel)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *SchemaModel) GetExternalDocs() ExternalDocumentationModelModel`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *SchemaModel) GetExternalDocsOk() (*ExternalDocumentationModelModel, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *SchemaModel) SetExternalDocs(v ExternalDocumentationModelModel)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *SchemaModel) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *SchemaModel) GetServers() []ServerModelModel`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SchemaModel) GetServersOk() (*[]ServerModelModel, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SchemaModel) SetServers(v []ServerModelModel)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SchemaModel) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *SchemaModel) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SchemaModel) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SchemaModel) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SchemaModel) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *SchemaModel) GetTags() []TagModelModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SchemaModel) GetTagsOk() (*[]TagModelModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SchemaModel) SetTags(v []TagModelModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SchemaModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *SchemaModel) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *SchemaModel) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *SchemaModel) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *SchemaModel) GetComponents() ComponentsModelModel`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SchemaModel) GetComponentsOk() (*ComponentsModelModel, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SchemaModel) SetComponents(v ComponentsModelModel)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SchemaModel) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


