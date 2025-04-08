# MiddlewareSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCheck** | Pointer to [**MiddlewareDefinition**](MiddlewareDefinition.md) |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**IdExtractor** | Pointer to [**MiddlewareIdExtractor**](MiddlewareIdExtractor.md) |  | [optional] 
**Post** | Pointer to [**[]MiddlewareDefinition**](MiddlewareDefinition.md) |  | [optional] 
**PostKeyAuth** | Pointer to [**[]MiddlewareDefinition**](MiddlewareDefinition.md) |  | [optional] 
**Pre** | Pointer to [**[]MiddlewareDefinition**](MiddlewareDefinition.md) |  | [optional] 
**Response** | Pointer to [**[]MiddlewareDefinition**](MiddlewareDefinition.md) |  | [optional] 

## Methods

### NewMiddlewareSection

`func NewMiddlewareSection() *MiddlewareSection`

NewMiddlewareSection instantiates a new MiddlewareSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareSectionWithDefaults

`func NewMiddlewareSectionWithDefaults() *MiddlewareSection`

NewMiddlewareSectionWithDefaults instantiates a new MiddlewareSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCheck

`func (o *MiddlewareSection) GetAuthCheck() MiddlewareDefinition`

GetAuthCheck returns the AuthCheck field if non-nil, zero value otherwise.

### GetAuthCheckOk

`func (o *MiddlewareSection) GetAuthCheckOk() (*MiddlewareDefinition, bool)`

GetAuthCheckOk returns a tuple with the AuthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCheck

`func (o *MiddlewareSection) SetAuthCheck(v MiddlewareDefinition)`

SetAuthCheck sets AuthCheck field to given value.

### HasAuthCheck

`func (o *MiddlewareSection) HasAuthCheck() bool`

HasAuthCheck returns a boolean if a field has been set.

### GetDriver

`func (o *MiddlewareSection) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *MiddlewareSection) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *MiddlewareSection) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *MiddlewareSection) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetIdExtractor

`func (o *MiddlewareSection) GetIdExtractor() MiddlewareIdExtractor`

GetIdExtractor returns the IdExtractor field if non-nil, zero value otherwise.

### GetIdExtractorOk

`func (o *MiddlewareSection) GetIdExtractorOk() (*MiddlewareIdExtractor, bool)`

GetIdExtractorOk returns a tuple with the IdExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractor

`func (o *MiddlewareSection) SetIdExtractor(v MiddlewareIdExtractor)`

SetIdExtractor sets IdExtractor field to given value.

### HasIdExtractor

`func (o *MiddlewareSection) HasIdExtractor() bool`

HasIdExtractor returns a boolean if a field has been set.

### GetPost

`func (o *MiddlewareSection) GetPost() []MiddlewareDefinition`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *MiddlewareSection) GetPostOk() (*[]MiddlewareDefinition, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *MiddlewareSection) SetPost(v []MiddlewareDefinition)`

SetPost sets Post field to given value.

### HasPost

`func (o *MiddlewareSection) HasPost() bool`

HasPost returns a boolean if a field has been set.

### SetPostNil

`func (o *MiddlewareSection) SetPostNil(b bool)`

 SetPostNil sets the value for Post to be an explicit nil

### UnsetPost
`func (o *MiddlewareSection) UnsetPost()`

UnsetPost ensures that no value is present for Post, not even an explicit nil
### GetPostKeyAuth

`func (o *MiddlewareSection) GetPostKeyAuth() []MiddlewareDefinition`

GetPostKeyAuth returns the PostKeyAuth field if non-nil, zero value otherwise.

### GetPostKeyAuthOk

`func (o *MiddlewareSection) GetPostKeyAuthOk() (*[]MiddlewareDefinition, bool)`

GetPostKeyAuthOk returns a tuple with the PostKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostKeyAuth

`func (o *MiddlewareSection) SetPostKeyAuth(v []MiddlewareDefinition)`

SetPostKeyAuth sets PostKeyAuth field to given value.

### HasPostKeyAuth

`func (o *MiddlewareSection) HasPostKeyAuth() bool`

HasPostKeyAuth returns a boolean if a field has been set.

### SetPostKeyAuthNil

`func (o *MiddlewareSection) SetPostKeyAuthNil(b bool)`

 SetPostKeyAuthNil sets the value for PostKeyAuth to be an explicit nil

### UnsetPostKeyAuth
`func (o *MiddlewareSection) UnsetPostKeyAuth()`

UnsetPostKeyAuth ensures that no value is present for PostKeyAuth, not even an explicit nil
### GetPre

`func (o *MiddlewareSection) GetPre() []MiddlewareDefinition`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *MiddlewareSection) GetPreOk() (*[]MiddlewareDefinition, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *MiddlewareSection) SetPre(v []MiddlewareDefinition)`

SetPre sets Pre field to given value.

### HasPre

`func (o *MiddlewareSection) HasPre() bool`

HasPre returns a boolean if a field has been set.

### SetPreNil

`func (o *MiddlewareSection) SetPreNil(b bool)`

 SetPreNil sets the value for Pre to be an explicit nil

### UnsetPre
`func (o *MiddlewareSection) UnsetPre()`

UnsetPre ensures that no value is present for Pre, not even an explicit nil
### GetResponse

`func (o *MiddlewareSection) GetResponse() []MiddlewareDefinition`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MiddlewareSection) GetResponseOk() (*[]MiddlewareDefinition, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MiddlewareSection) SetResponse(v []MiddlewareDefinition)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MiddlewareSection) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *MiddlewareSection) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *MiddlewareSection) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


