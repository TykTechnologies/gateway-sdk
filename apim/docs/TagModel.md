# TagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExternalDocs** | Pointer to [**ExternalDocumentationModelModel**](ExternalDocumentationModel.md) |  | [optional] 

## Methods

### NewTagModel

`func NewTagModel(name string, ) *TagModel`

NewTagModel instantiates a new TagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelWithDefaults

`func NewTagModelWithDefaults() *TagModel`

NewTagModelWithDefaults instantiates a new TagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDocs

`func (o *TagModel) GetExternalDocs() ExternalDocumentationModelModel`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *TagModel) GetExternalDocsOk() (*ExternalDocumentationModelModel, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *TagModel) SetExternalDocs(v ExternalDocumentationModelModel)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *TagModel) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


