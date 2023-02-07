# VersionInfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**VersionInfoPathsModelModel**](VersionInfoPathsModel.md) |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 
**ExtendedPaths** | Pointer to [**ExtendedPathsSetModelModel**](ExtendedPathsSetModel.md) |  | [optional] 
**GlobalHeaders** | Pointer to **map[string]string** |  | [optional] 
**GlobalHeadersRemove** | Pointer to **[]string** |  | [optional] 
**GlobalSizeLimit** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideTarget** | Pointer to **string** |  | [optional] 
**UseExtendedPaths** | Pointer to **bool** |  | [optional] 

## Methods

### NewVersionInfoModel

`func NewVersionInfoModel() *VersionInfoModel`

NewVersionInfoModel instantiates a new VersionInfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoModelWithDefaults

`func NewVersionInfoModelWithDefaults() *VersionInfoModel`

NewVersionInfoModelWithDefaults instantiates a new VersionInfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *VersionInfoModel) GetPaths() VersionInfoPathsModelModel`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *VersionInfoModel) GetPathsOk() (*VersionInfoPathsModelModel, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *VersionInfoModel) SetPaths(v VersionInfoPathsModelModel)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *VersionInfoModel) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetExpires

`func (o *VersionInfoModel) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *VersionInfoModel) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *VersionInfoModel) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *VersionInfoModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExtendedPaths

`func (o *VersionInfoModel) GetExtendedPaths() ExtendedPathsSetModelModel`

GetExtendedPaths returns the ExtendedPaths field if non-nil, zero value otherwise.

### GetExtendedPathsOk

`func (o *VersionInfoModel) GetExtendedPathsOk() (*ExtendedPathsSetModelModel, bool)`

GetExtendedPathsOk returns a tuple with the ExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPaths

`func (o *VersionInfoModel) SetExtendedPaths(v ExtendedPathsSetModelModel)`

SetExtendedPaths sets ExtendedPaths field to given value.

### HasExtendedPaths

`func (o *VersionInfoModel) HasExtendedPaths() bool`

HasExtendedPaths returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *VersionInfoModel) GetGlobalHeaders() map[string]string`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *VersionInfoModel) GetGlobalHeadersOk() (*map[string]string, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *VersionInfoModel) SetGlobalHeaders(v map[string]string)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *VersionInfoModel) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetGlobalHeadersRemove

`func (o *VersionInfoModel) GetGlobalHeadersRemove() []string`

GetGlobalHeadersRemove returns the GlobalHeadersRemove field if non-nil, zero value otherwise.

### GetGlobalHeadersRemoveOk

`func (o *VersionInfoModel) GetGlobalHeadersRemoveOk() (*[]string, bool)`

GetGlobalHeadersRemoveOk returns a tuple with the GlobalHeadersRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeadersRemove

`func (o *VersionInfoModel) SetGlobalHeadersRemove(v []string)`

SetGlobalHeadersRemove sets GlobalHeadersRemove field to given value.

### HasGlobalHeadersRemove

`func (o *VersionInfoModel) HasGlobalHeadersRemove() bool`

HasGlobalHeadersRemove returns a boolean if a field has been set.

### GetGlobalSizeLimit

`func (o *VersionInfoModel) GetGlobalSizeLimit() int64`

GetGlobalSizeLimit returns the GlobalSizeLimit field if non-nil, zero value otherwise.

### GetGlobalSizeLimitOk

`func (o *VersionInfoModel) GetGlobalSizeLimitOk() (*int64, bool)`

GetGlobalSizeLimitOk returns a tuple with the GlobalSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSizeLimit

`func (o *VersionInfoModel) SetGlobalSizeLimit(v int64)`

SetGlobalSizeLimit sets GlobalSizeLimit field to given value.

### HasGlobalSizeLimit

`func (o *VersionInfoModel) HasGlobalSizeLimit() bool`

HasGlobalSizeLimit returns a boolean if a field has been set.

### GetName

`func (o *VersionInfoModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionInfoModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionInfoModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionInfoModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *VersionInfoModel) GetOverrideTarget() string`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *VersionInfoModel) GetOverrideTargetOk() (*string, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *VersionInfoModel) SetOverrideTarget(v string)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *VersionInfoModel) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetUseExtendedPaths

`func (o *VersionInfoModel) GetUseExtendedPaths() bool`

GetUseExtendedPaths returns the UseExtendedPaths field if non-nil, zero value otherwise.

### GetUseExtendedPathsOk

`func (o *VersionInfoModel) GetUseExtendedPathsOk() (*bool, bool)`

GetUseExtendedPathsOk returns a tuple with the UseExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExtendedPaths

`func (o *VersionInfoModel) SetUseExtendedPaths(v bool)`

SetUseExtendedPaths sets UseExtendedPaths field to given value.

### HasUseExtendedPaths

`func (o *VersionInfoModel) HasUseExtendedPaths() bool`

HasUseExtendedPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


