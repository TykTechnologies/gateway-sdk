# AccessDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUrls** | Pointer to [**[]AccessSpec**](AccessSpec.md) |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**ApiName** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to [**APILimit**](APILimit.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessDefinition

`func NewAccessDefinition() *AccessDefinition`

NewAccessDefinition instantiates a new AccessDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDefinitionWithDefaults

`func NewAccessDefinitionWithDefaults() *AccessDefinition`

NewAccessDefinitionWithDefaults instantiates a new AccessDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUrls

`func (o *AccessDefinition) GetAllowedUrls() []AccessSpec`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *AccessDefinition) GetAllowedUrlsOk() (*[]AccessSpec, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *AccessDefinition) SetAllowedUrls(v []AccessSpec)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *AccessDefinition) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetApiId

`func (o *AccessDefinition) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *AccessDefinition) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *AccessDefinition) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *AccessDefinition) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiName

`func (o *AccessDefinition) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *AccessDefinition) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *AccessDefinition) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *AccessDefinition) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetLimit

`func (o *AccessDefinition) GetLimit() APILimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccessDefinition) GetLimitOk() (*APILimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccessDefinition) SetLimit(v APILimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccessDefinition) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetVersions

`func (o *AccessDefinition) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AccessDefinition) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AccessDefinition) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AccessDefinition) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


