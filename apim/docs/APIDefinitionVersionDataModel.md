# APIDefinitionVersionDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **string** |  | [optional] 
**NotVersioned** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**map[string]VersionInfoModelModel**](VersionInfoModel.md) |  | [optional] 

## Methods

### NewAPIDefinitionVersionDataModel

`func NewAPIDefinitionVersionDataModel() *APIDefinitionVersionDataModel`

NewAPIDefinitionVersionDataModel instantiates a new APIDefinitionVersionDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionVersionDataModelWithDefaults

`func NewAPIDefinitionVersionDataModelWithDefaults() *APIDefinitionVersionDataModel`

NewAPIDefinitionVersionDataModelWithDefaults instantiates a new APIDefinitionVersionDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *APIDefinitionVersionDataModel) GetDefaultVersion() string`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *APIDefinitionVersionDataModel) GetDefaultVersionOk() (*string, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *APIDefinitionVersionDataModel) SetDefaultVersion(v string)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *APIDefinitionVersionDataModel) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetNotVersioned

`func (o *APIDefinitionVersionDataModel) GetNotVersioned() bool`

GetNotVersioned returns the NotVersioned field if non-nil, zero value otherwise.

### GetNotVersionedOk

`func (o *APIDefinitionVersionDataModel) GetNotVersionedOk() (*bool, bool)`

GetNotVersionedOk returns a tuple with the NotVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotVersioned

`func (o *APIDefinitionVersionDataModel) SetNotVersioned(v bool)`

SetNotVersioned sets NotVersioned field to given value.

### HasNotVersioned

`func (o *APIDefinitionVersionDataModel) HasNotVersioned() bool`

HasNotVersioned returns a boolean if a field has been set.

### GetVersions

`func (o *APIDefinitionVersionDataModel) GetVersions() map[string]VersionInfoModelModel`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *APIDefinitionVersionDataModel) GetVersionsOk() (*map[string]VersionInfoModelModel, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *APIDefinitionVersionDataModel) SetVersions(v map[string]VersionInfoModelModel)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *APIDefinitionVersionDataModel) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


