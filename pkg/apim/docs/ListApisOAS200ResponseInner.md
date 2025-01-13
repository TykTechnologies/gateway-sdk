# ListApisOAS200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | [**Info**](Info.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server1**](Server1.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components**](Components.md) |  | [optional] 
**Middleware** | Pointer to [**Middleware**](Middleware.md) |  | [optional] 
**Server** | Pointer to [**Server**](Server.md) |  | [optional] 
**Upstream** | Pointer to [**Upstream**](Upstream.md) |  | [optional] 

## Methods

### NewListApisOAS200ResponseInner

`func NewListApisOAS200ResponseInner(openapi string, info Info, paths map[string]interface{}, ) *ListApisOAS200ResponseInner`

NewListApisOAS200ResponseInner instantiates a new ListApisOAS200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApisOAS200ResponseInnerWithDefaults

`func NewListApisOAS200ResponseInnerWithDefaults() *ListApisOAS200ResponseInner`

NewListApisOAS200ResponseInnerWithDefaults instantiates a new ListApisOAS200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *ListApisOAS200ResponseInner) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *ListApisOAS200ResponseInner) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *ListApisOAS200ResponseInner) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *ListApisOAS200ResponseInner) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ListApisOAS200ResponseInner) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ListApisOAS200ResponseInner) SetInfo(v Info)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *ListApisOAS200ResponseInner) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *ListApisOAS200ResponseInner) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *ListApisOAS200ResponseInner) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *ListApisOAS200ResponseInner) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *ListApisOAS200ResponseInner) GetServers() []Server1`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListApisOAS200ResponseInner) GetServersOk() (*[]Server1, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListApisOAS200ResponseInner) SetServers(v []Server1)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListApisOAS200ResponseInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *ListApisOAS200ResponseInner) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ListApisOAS200ResponseInner) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ListApisOAS200ResponseInner) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ListApisOAS200ResponseInner) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *ListApisOAS200ResponseInner) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListApisOAS200ResponseInner) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListApisOAS200ResponseInner) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListApisOAS200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *ListApisOAS200ResponseInner) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ListApisOAS200ResponseInner) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ListApisOAS200ResponseInner) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *ListApisOAS200ResponseInner) GetComponents() Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ListApisOAS200ResponseInner) GetComponentsOk() (*Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ListApisOAS200ResponseInner) SetComponents(v Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ListApisOAS200ResponseInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetMiddleware

`func (o *ListApisOAS200ResponseInner) GetMiddleware() Middleware`

GetMiddleware returns the Middleware field if non-nil, zero value otherwise.

### GetMiddlewareOk

`func (o *ListApisOAS200ResponseInner) GetMiddlewareOk() (*Middleware, bool)`

GetMiddlewareOk returns a tuple with the Middleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleware

`func (o *ListApisOAS200ResponseInner) SetMiddleware(v Middleware)`

SetMiddleware sets Middleware field to given value.

### HasMiddleware

`func (o *ListApisOAS200ResponseInner) HasMiddleware() bool`

HasMiddleware returns a boolean if a field has been set.

### GetServer

`func (o *ListApisOAS200ResponseInner) GetServer() Server`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ListApisOAS200ResponseInner) GetServerOk() (*Server, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ListApisOAS200ResponseInner) SetServer(v Server)`

SetServer sets Server field to given value.

### HasServer

`func (o *ListApisOAS200ResponseInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpstream

`func (o *ListApisOAS200ResponseInner) GetUpstream() Upstream`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ListApisOAS200ResponseInner) GetUpstreamOk() (*Upstream, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ListApisOAS200ResponseInner) SetUpstream(v Upstream)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ListApisOAS200ResponseInner) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


