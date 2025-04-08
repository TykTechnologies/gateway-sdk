# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **string** |  | [optional] 
**ExtendedPaths** | Pointer to [**ExtendedPathsSet**](ExtendedPathsSet.md) |  | [optional] 
**GlobalHeaders** | Pointer to **map[string]string** |  | [optional] 
**GlobalHeadersDisabled** | Pointer to **bool** |  | [optional] 
**GlobalHeadersRemove** | Pointer to **[]string** |  | [optional] 
**GlobalResponseHeaders** | Pointer to **map[string]string** |  | [optional] 
**GlobalResponseHeadersDisabled** | Pointer to **bool** |  | [optional] 
**GlobalResponseHeadersRemove** | Pointer to **[]string** |  | [optional] 
**GlobalSizeLimit** | Pointer to **int64** |  | [optional] 
**IgnoreEndpointCase** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideTarget** | Pointer to **string** |  | [optional] 
**Paths** | Pointer to [**VersionInfoPaths**](VersionInfoPaths.md) |  | [optional] 
**UseExtendedPaths** | Pointer to **bool** |  | [optional] 

## Methods

### NewVersionInfo

`func NewVersionInfo() *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *VersionInfo) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *VersionInfo) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *VersionInfo) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *VersionInfo) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExtendedPaths

`func (o *VersionInfo) GetExtendedPaths() ExtendedPathsSet`

GetExtendedPaths returns the ExtendedPaths field if non-nil, zero value otherwise.

### GetExtendedPathsOk

`func (o *VersionInfo) GetExtendedPathsOk() (*ExtendedPathsSet, bool)`

GetExtendedPathsOk returns a tuple with the ExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPaths

`func (o *VersionInfo) SetExtendedPaths(v ExtendedPathsSet)`

SetExtendedPaths sets ExtendedPaths field to given value.

### HasExtendedPaths

`func (o *VersionInfo) HasExtendedPaths() bool`

HasExtendedPaths returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *VersionInfo) GetGlobalHeaders() map[string]string`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *VersionInfo) GetGlobalHeadersOk() (*map[string]string, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *VersionInfo) SetGlobalHeaders(v map[string]string)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *VersionInfo) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### SetGlobalHeadersNil

`func (o *VersionInfo) SetGlobalHeadersNil(b bool)`

 SetGlobalHeadersNil sets the value for GlobalHeaders to be an explicit nil

### UnsetGlobalHeaders
`func (o *VersionInfo) UnsetGlobalHeaders()`

UnsetGlobalHeaders ensures that no value is present for GlobalHeaders, not even an explicit nil
### GetGlobalHeadersDisabled

`func (o *VersionInfo) GetGlobalHeadersDisabled() bool`

GetGlobalHeadersDisabled returns the GlobalHeadersDisabled field if non-nil, zero value otherwise.

### GetGlobalHeadersDisabledOk

`func (o *VersionInfo) GetGlobalHeadersDisabledOk() (*bool, bool)`

GetGlobalHeadersDisabledOk returns a tuple with the GlobalHeadersDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeadersDisabled

`func (o *VersionInfo) SetGlobalHeadersDisabled(v bool)`

SetGlobalHeadersDisabled sets GlobalHeadersDisabled field to given value.

### HasGlobalHeadersDisabled

`func (o *VersionInfo) HasGlobalHeadersDisabled() bool`

HasGlobalHeadersDisabled returns a boolean if a field has been set.

### GetGlobalHeadersRemove

`func (o *VersionInfo) GetGlobalHeadersRemove() []string`

GetGlobalHeadersRemove returns the GlobalHeadersRemove field if non-nil, zero value otherwise.

### GetGlobalHeadersRemoveOk

`func (o *VersionInfo) GetGlobalHeadersRemoveOk() (*[]string, bool)`

GetGlobalHeadersRemoveOk returns a tuple with the GlobalHeadersRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeadersRemove

`func (o *VersionInfo) SetGlobalHeadersRemove(v []string)`

SetGlobalHeadersRemove sets GlobalHeadersRemove field to given value.

### HasGlobalHeadersRemove

`func (o *VersionInfo) HasGlobalHeadersRemove() bool`

HasGlobalHeadersRemove returns a boolean if a field has been set.

### SetGlobalHeadersRemoveNil

`func (o *VersionInfo) SetGlobalHeadersRemoveNil(b bool)`

 SetGlobalHeadersRemoveNil sets the value for GlobalHeadersRemove to be an explicit nil

### UnsetGlobalHeadersRemove
`func (o *VersionInfo) UnsetGlobalHeadersRemove()`

UnsetGlobalHeadersRemove ensures that no value is present for GlobalHeadersRemove, not even an explicit nil
### GetGlobalResponseHeaders

`func (o *VersionInfo) GetGlobalResponseHeaders() map[string]string`

GetGlobalResponseHeaders returns the GlobalResponseHeaders field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersOk

`func (o *VersionInfo) GetGlobalResponseHeadersOk() (*map[string]string, bool)`

GetGlobalResponseHeadersOk returns a tuple with the GlobalResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeaders

`func (o *VersionInfo) SetGlobalResponseHeaders(v map[string]string)`

SetGlobalResponseHeaders sets GlobalResponseHeaders field to given value.

### HasGlobalResponseHeaders

`func (o *VersionInfo) HasGlobalResponseHeaders() bool`

HasGlobalResponseHeaders returns a boolean if a field has been set.

### SetGlobalResponseHeadersNil

`func (o *VersionInfo) SetGlobalResponseHeadersNil(b bool)`

 SetGlobalResponseHeadersNil sets the value for GlobalResponseHeaders to be an explicit nil

### UnsetGlobalResponseHeaders
`func (o *VersionInfo) UnsetGlobalResponseHeaders()`

UnsetGlobalResponseHeaders ensures that no value is present for GlobalResponseHeaders, not even an explicit nil
### GetGlobalResponseHeadersDisabled

`func (o *VersionInfo) GetGlobalResponseHeadersDisabled() bool`

GetGlobalResponseHeadersDisabled returns the GlobalResponseHeadersDisabled field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersDisabledOk

`func (o *VersionInfo) GetGlobalResponseHeadersDisabledOk() (*bool, bool)`

GetGlobalResponseHeadersDisabledOk returns a tuple with the GlobalResponseHeadersDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeadersDisabled

`func (o *VersionInfo) SetGlobalResponseHeadersDisabled(v bool)`

SetGlobalResponseHeadersDisabled sets GlobalResponseHeadersDisabled field to given value.

### HasGlobalResponseHeadersDisabled

`func (o *VersionInfo) HasGlobalResponseHeadersDisabled() bool`

HasGlobalResponseHeadersDisabled returns a boolean if a field has been set.

### GetGlobalResponseHeadersRemove

`func (o *VersionInfo) GetGlobalResponseHeadersRemove() []string`

GetGlobalResponseHeadersRemove returns the GlobalResponseHeadersRemove field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersRemoveOk

`func (o *VersionInfo) GetGlobalResponseHeadersRemoveOk() (*[]string, bool)`

GetGlobalResponseHeadersRemoveOk returns a tuple with the GlobalResponseHeadersRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeadersRemove

`func (o *VersionInfo) SetGlobalResponseHeadersRemove(v []string)`

SetGlobalResponseHeadersRemove sets GlobalResponseHeadersRemove field to given value.

### HasGlobalResponseHeadersRemove

`func (o *VersionInfo) HasGlobalResponseHeadersRemove() bool`

HasGlobalResponseHeadersRemove returns a boolean if a field has been set.

### SetGlobalResponseHeadersRemoveNil

`func (o *VersionInfo) SetGlobalResponseHeadersRemoveNil(b bool)`

 SetGlobalResponseHeadersRemoveNil sets the value for GlobalResponseHeadersRemove to be an explicit nil

### UnsetGlobalResponseHeadersRemove
`func (o *VersionInfo) UnsetGlobalResponseHeadersRemove()`

UnsetGlobalResponseHeadersRemove ensures that no value is present for GlobalResponseHeadersRemove, not even an explicit nil
### GetGlobalSizeLimit

`func (o *VersionInfo) GetGlobalSizeLimit() int64`

GetGlobalSizeLimit returns the GlobalSizeLimit field if non-nil, zero value otherwise.

### GetGlobalSizeLimitOk

`func (o *VersionInfo) GetGlobalSizeLimitOk() (*int64, bool)`

GetGlobalSizeLimitOk returns a tuple with the GlobalSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSizeLimit

`func (o *VersionInfo) SetGlobalSizeLimit(v int64)`

SetGlobalSizeLimit sets GlobalSizeLimit field to given value.

### HasGlobalSizeLimit

`func (o *VersionInfo) HasGlobalSizeLimit() bool`

HasGlobalSizeLimit returns a boolean if a field has been set.

### GetIgnoreEndpointCase

`func (o *VersionInfo) GetIgnoreEndpointCase() bool`

GetIgnoreEndpointCase returns the IgnoreEndpointCase field if non-nil, zero value otherwise.

### GetIgnoreEndpointCaseOk

`func (o *VersionInfo) GetIgnoreEndpointCaseOk() (*bool, bool)`

GetIgnoreEndpointCaseOk returns a tuple with the IgnoreEndpointCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEndpointCase

`func (o *VersionInfo) SetIgnoreEndpointCase(v bool)`

SetIgnoreEndpointCase sets IgnoreEndpointCase field to given value.

### HasIgnoreEndpointCase

`func (o *VersionInfo) HasIgnoreEndpointCase() bool`

HasIgnoreEndpointCase returns a boolean if a field has been set.

### GetName

`func (o *VersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *VersionInfo) GetOverrideTarget() string`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *VersionInfo) GetOverrideTargetOk() (*string, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *VersionInfo) SetOverrideTarget(v string)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *VersionInfo) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetPaths

`func (o *VersionInfo) GetPaths() VersionInfoPaths`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *VersionInfo) GetPathsOk() (*VersionInfoPaths, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *VersionInfo) SetPaths(v VersionInfoPaths)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *VersionInfo) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetUseExtendedPaths

`func (o *VersionInfo) GetUseExtendedPaths() bool`

GetUseExtendedPaths returns the UseExtendedPaths field if non-nil, zero value otherwise.

### GetUseExtendedPathsOk

`func (o *VersionInfo) GetUseExtendedPathsOk() (*bool, bool)`

GetUseExtendedPathsOk returns a tuple with the UseExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExtendedPaths

`func (o *VersionInfo) SetUseExtendedPaths(v bool)`

SetUseExtendedPaths sets UseExtendedPaths field to given value.

### HasUseExtendedPaths

`func (o *VersionInfo) HasUseExtendedPaths() bool`

HasUseExtendedPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


