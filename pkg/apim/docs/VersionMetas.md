# VersionMetas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to [**[]VersionMeta**](VersionMeta.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewVersionMetas

`func NewVersionMetas() *VersionMetas`

NewVersionMetas instantiates a new VersionMetas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMetasWithDefaults

`func NewVersionMetasWithDefaults() *VersionMetas`

NewVersionMetasWithDefaults instantiates a new VersionMetas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *VersionMetas) GetApis() []VersionMeta`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *VersionMetas) GetApisOk() (*[]VersionMeta, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *VersionMetas) SetApis(v []VersionMeta)`

SetApis sets Apis field to given value.

### HasApis

`func (o *VersionMetas) HasApis() bool`

HasApis returns a boolean if a field has been set.

### SetApisNil

`func (o *VersionMetas) SetApisNil(b bool)`

 SetApisNil sets the value for Apis to be an explicit nil

### UnsetApis
`func (o *VersionMetas) UnsetApis()`

UnsetApis ensures that no value is present for Apis, not even an explicit nil
### GetStatus

`func (o *VersionMetas) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionMetas) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionMetas) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VersionMetas) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


