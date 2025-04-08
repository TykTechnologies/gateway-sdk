# UptimeTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckList** | Pointer to [**[]HostCheckObject**](HostCheckObject.md) |  | [optional] 
**Config** | Pointer to [**UptimeTestsConfig**](UptimeTestsConfig.md) |  | [optional] 

## Methods

### NewUptimeTests

`func NewUptimeTests() *UptimeTests`

NewUptimeTests instantiates a new UptimeTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptimeTestsWithDefaults

`func NewUptimeTestsWithDefaults() *UptimeTests`

NewUptimeTestsWithDefaults instantiates a new UptimeTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckList

`func (o *UptimeTests) GetCheckList() []HostCheckObject`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *UptimeTests) GetCheckListOk() (*[]HostCheckObject, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *UptimeTests) SetCheckList(v []HostCheckObject)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *UptimeTests) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### SetCheckListNil

`func (o *UptimeTests) SetCheckListNil(b bool)`

 SetCheckListNil sets the value for CheckList to be an explicit nil

### UnsetCheckList
`func (o *UptimeTests) UnsetCheckList()`

UnsetCheckList ensures that no value is present for CheckList, not even an explicit nil
### GetConfig

`func (o *UptimeTests) GetConfig() UptimeTestsConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UptimeTests) GetConfigOk() (*UptimeTestsConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UptimeTests) SetConfig(v UptimeTestsConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UptimeTests) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


