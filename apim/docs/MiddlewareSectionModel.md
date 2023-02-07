# MiddlewareSectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCheck** | Pointer to [**MiddlewareDefinitionModelModel**](MiddlewareDefinitionModel.md) |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**IdExtractor** | Pointer to [**MiddlewareIdExtractorModelModel**](MiddlewareIdExtractorModel.md) |  | [optional] 
**Post** | Pointer to [**[]MiddlewareDefinitionModelModel**](MiddlewareDefinitionModelModel.md) |  | [optional] 
**PostKeyAuth** | Pointer to [**[]MiddlewareDefinitionModelModel**](MiddlewareDefinitionModelModel.md) |  | [optional] 
**Pre** | Pointer to [**[]MiddlewareDefinitionModelModel**](MiddlewareDefinitionModelModel.md) |  | [optional] 
**Response** | Pointer to [**[]MiddlewareDefinitionModelModel**](MiddlewareDefinitionModelModel.md) |  | [optional] 

## Methods

### NewMiddlewareSectionModel

`func NewMiddlewareSectionModel() *MiddlewareSectionModel`

NewMiddlewareSectionModel instantiates a new MiddlewareSectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareSectionModelWithDefaults

`func NewMiddlewareSectionModelWithDefaults() *MiddlewareSectionModel`

NewMiddlewareSectionModelWithDefaults instantiates a new MiddlewareSectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCheck

`func (o *MiddlewareSectionModel) GetAuthCheck() MiddlewareDefinitionModelModel`

GetAuthCheck returns the AuthCheck field if non-nil, zero value otherwise.

### GetAuthCheckOk

`func (o *MiddlewareSectionModel) GetAuthCheckOk() (*MiddlewareDefinitionModelModel, bool)`

GetAuthCheckOk returns a tuple with the AuthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCheck

`func (o *MiddlewareSectionModel) SetAuthCheck(v MiddlewareDefinitionModelModel)`

SetAuthCheck sets AuthCheck field to given value.

### HasAuthCheck

`func (o *MiddlewareSectionModel) HasAuthCheck() bool`

HasAuthCheck returns a boolean if a field has been set.

### GetDriver

`func (o *MiddlewareSectionModel) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *MiddlewareSectionModel) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *MiddlewareSectionModel) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *MiddlewareSectionModel) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetIdExtractor

`func (o *MiddlewareSectionModel) GetIdExtractor() MiddlewareIdExtractorModelModel`

GetIdExtractor returns the IdExtractor field if non-nil, zero value otherwise.

### GetIdExtractorOk

`func (o *MiddlewareSectionModel) GetIdExtractorOk() (*MiddlewareIdExtractorModelModel, bool)`

GetIdExtractorOk returns a tuple with the IdExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractor

`func (o *MiddlewareSectionModel) SetIdExtractor(v MiddlewareIdExtractorModelModel)`

SetIdExtractor sets IdExtractor field to given value.

### HasIdExtractor

`func (o *MiddlewareSectionModel) HasIdExtractor() bool`

HasIdExtractor returns a boolean if a field has been set.

### GetPost

`func (o *MiddlewareSectionModel) GetPost() []MiddlewareDefinitionModelModel`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *MiddlewareSectionModel) GetPostOk() (*[]MiddlewareDefinitionModelModel, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *MiddlewareSectionModel) SetPost(v []MiddlewareDefinitionModelModel)`

SetPost sets Post field to given value.

### HasPost

`func (o *MiddlewareSectionModel) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPostKeyAuth

`func (o *MiddlewareSectionModel) GetPostKeyAuth() []MiddlewareDefinitionModelModel`

GetPostKeyAuth returns the PostKeyAuth field if non-nil, zero value otherwise.

### GetPostKeyAuthOk

`func (o *MiddlewareSectionModel) GetPostKeyAuthOk() (*[]MiddlewareDefinitionModelModel, bool)`

GetPostKeyAuthOk returns a tuple with the PostKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostKeyAuth

`func (o *MiddlewareSectionModel) SetPostKeyAuth(v []MiddlewareDefinitionModelModel)`

SetPostKeyAuth sets PostKeyAuth field to given value.

### HasPostKeyAuth

`func (o *MiddlewareSectionModel) HasPostKeyAuth() bool`

HasPostKeyAuth returns a boolean if a field has been set.

### GetPre

`func (o *MiddlewareSectionModel) GetPre() []MiddlewareDefinitionModelModel`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *MiddlewareSectionModel) GetPreOk() (*[]MiddlewareDefinitionModelModel, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *MiddlewareSectionModel) SetPre(v []MiddlewareDefinitionModelModel)`

SetPre sets Pre field to given value.

### HasPre

`func (o *MiddlewareSectionModel) HasPre() bool`

HasPre returns a boolean if a field has been set.

### GetResponse

`func (o *MiddlewareSectionModel) GetResponse() []MiddlewareDefinitionModelModel`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MiddlewareSectionModel) GetResponseOk() (*[]MiddlewareDefinitionModelModel, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MiddlewareSectionModel) SetResponse(v []MiddlewareDefinitionModelModel)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MiddlewareSectionModel) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


