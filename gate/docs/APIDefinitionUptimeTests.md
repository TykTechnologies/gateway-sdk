# APIDefinitionUptimeTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckList** | Pointer to [**[]HostCheckObject**](HostCheckObject.md) |  | [optional] 
**Config** | Pointer to [**APIDefinitionUptimeTestsConfig**](APIDefinitionUptimeTestsConfig.md) |  | [optional] 

## Methods

### NewAPIDefinitionUptimeTests

`func NewAPIDefinitionUptimeTests() *APIDefinitionUptimeTests`

NewAPIDefinitionUptimeTests instantiates a new APIDefinitionUptimeTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionUptimeTestsWithDefaults

`func NewAPIDefinitionUptimeTestsWithDefaults() *APIDefinitionUptimeTests`

NewAPIDefinitionUptimeTestsWithDefaults instantiates a new APIDefinitionUptimeTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckList

`func (o *APIDefinitionUptimeTests) GetCheckList() []HostCheckObject`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *APIDefinitionUptimeTests) GetCheckListOk() (*[]HostCheckObject, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *APIDefinitionUptimeTests) SetCheckList(v []HostCheckObject)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *APIDefinitionUptimeTests) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### GetConfig

`func (o *APIDefinitionUptimeTests) GetConfig() APIDefinitionUptimeTestsConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *APIDefinitionUptimeTests) GetConfigOk() (*APIDefinitionUptimeTestsConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *APIDefinitionUptimeTests) SetConfig(v APIDefinitionUptimeTestsConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *APIDefinitionUptimeTests) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


