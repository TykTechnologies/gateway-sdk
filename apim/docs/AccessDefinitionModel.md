# AccessDefinitionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUrls** | Pointer to [**[]AccessSpecModelModel**](AccessSpecModelModel.md) |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**ApiName** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to [**APILimitModelModel**](APILimitModel.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessDefinitionModel

`func NewAccessDefinitionModel() *AccessDefinitionModel`

NewAccessDefinitionModel instantiates a new AccessDefinitionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDefinitionModelWithDefaults

`func NewAccessDefinitionModelWithDefaults() *AccessDefinitionModel`

NewAccessDefinitionModelWithDefaults instantiates a new AccessDefinitionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUrls

`func (o *AccessDefinitionModel) GetAllowedUrls() []AccessSpecModelModel`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *AccessDefinitionModel) GetAllowedUrlsOk() (*[]AccessSpecModelModel, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *AccessDefinitionModel) SetAllowedUrls(v []AccessSpecModelModel)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *AccessDefinitionModel) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetApiId

`func (o *AccessDefinitionModel) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *AccessDefinitionModel) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *AccessDefinitionModel) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *AccessDefinitionModel) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiName

`func (o *AccessDefinitionModel) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *AccessDefinitionModel) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *AccessDefinitionModel) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *AccessDefinitionModel) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetLimit

`func (o *AccessDefinitionModel) GetLimit() APILimitModelModel`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccessDefinitionModel) GetLimitOk() (*APILimitModelModel, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccessDefinitionModel) SetLimit(v APILimitModelModel)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccessDefinitionModel) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetVersions

`func (o *AccessDefinitionModel) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AccessDefinitionModel) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AccessDefinitionModel) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AccessDefinitionModel) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


