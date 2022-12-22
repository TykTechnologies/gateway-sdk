# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | [**Info**](Info.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components**](Components.md) |  | [optional] 

## Methods

### NewSchema

`func NewSchema(openapi string, info Info, paths map[string]interface{}, ) *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *Schema) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *Schema) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *Schema) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *Schema) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Schema) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Schema) SetInfo(v Info)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *Schema) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *Schema) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *Schema) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *Schema) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *Schema) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Schema) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Schema) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Schema) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *Schema) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *Schema) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *Schema) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *Schema) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *Schema) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Schema) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Schema) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Schema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *Schema) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *Schema) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *Schema) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *Schema) GetComponents() Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Schema) GetComponentsOk() (*Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Schema) SetComponents(v Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Schema) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


