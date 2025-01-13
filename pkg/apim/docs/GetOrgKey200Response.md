# GetOrgKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRights** | Pointer to [**map[string]GetOrgKey200ResponseAccessRightsValue**](GetOrgKey200ResponseAccessRightsValue.md) |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **int32** |  | [optional] 
**ApplyPolicies** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**EnableDetailedRecording** | Pointer to **bool** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Per** | Pointer to **int32** |  | [optional] 
**QuotaMax** | Pointer to **int32** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int32** |  | [optional] 
**QuotaRenews** | Pointer to **int32** |  | [optional] 
**Rate** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThrottleInterval** | Pointer to **int32** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetOrgKey200Response

`func NewGetOrgKey200Response() *GetOrgKey200Response`

NewGetOrgKey200Response instantiates a new GetOrgKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgKey200ResponseWithDefaults

`func NewGetOrgKey200ResponseWithDefaults() *GetOrgKey200Response`

NewGetOrgKey200ResponseWithDefaults instantiates a new GetOrgKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRights

`func (o *GetOrgKey200Response) GetAccessRights() map[string]GetOrgKey200ResponseAccessRightsValue`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *GetOrgKey200Response) GetAccessRightsOk() (*map[string]GetOrgKey200ResponseAccessRightsValue, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *GetOrgKey200Response) SetAccessRights(v map[string]GetOrgKey200ResponseAccessRightsValue)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *GetOrgKey200Response) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### SetAccessRightsNil

`func (o *GetOrgKey200Response) SetAccessRightsNil(b bool)`

 SetAccessRightsNil sets the value for AccessRights to be an explicit nil

### UnsetAccessRights
`func (o *GetOrgKey200Response) UnsetAccessRights()`

UnsetAccessRights ensures that no value is present for AccessRights, not even an explicit nil
### GetAlias

`func (o *GetOrgKey200Response) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetOrgKey200Response) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetOrgKey200Response) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetOrgKey200Response) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAllowance

`func (o *GetOrgKey200Response) GetAllowance() int32`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *GetOrgKey200Response) GetAllowanceOk() (*int32, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *GetOrgKey200Response) SetAllowance(v int32)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *GetOrgKey200Response) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetApplyPolicies

`func (o *GetOrgKey200Response) GetApplyPolicies() []string`

GetApplyPolicies returns the ApplyPolicies field if non-nil, zero value otherwise.

### GetApplyPoliciesOk

`func (o *GetOrgKey200Response) GetApplyPoliciesOk() (*[]string, bool)`

GetApplyPoliciesOk returns a tuple with the ApplyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicies

`func (o *GetOrgKey200Response) SetApplyPolicies(v []string)`

SetApplyPolicies sets ApplyPolicies field to given value.

### HasApplyPolicies

`func (o *GetOrgKey200Response) HasApplyPolicies() bool`

HasApplyPolicies returns a boolean if a field has been set.

### SetApplyPoliciesNil

`func (o *GetOrgKey200Response) SetApplyPoliciesNil(b bool)`

 SetApplyPoliciesNil sets the value for ApplyPolicies to be an explicit nil

### UnsetApplyPolicies
`func (o *GetOrgKey200Response) UnsetApplyPolicies()`

UnsetApplyPolicies ensures that no value is present for ApplyPolicies, not even an explicit nil
### GetDateCreated

`func (o *GetOrgKey200Response) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetOrgKey200Response) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetOrgKey200Response) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetOrgKey200Response) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetEnableDetailedRecording

`func (o *GetOrgKey200Response) GetEnableDetailedRecording() bool`

GetEnableDetailedRecording returns the EnableDetailedRecording field if non-nil, zero value otherwise.

### GetEnableDetailedRecordingOk

`func (o *GetOrgKey200Response) GetEnableDetailedRecordingOk() (*bool, bool)`

GetEnableDetailedRecordingOk returns a tuple with the EnableDetailedRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDetailedRecording

`func (o *GetOrgKey200Response) SetEnableDetailedRecording(v bool)`

SetEnableDetailedRecording sets EnableDetailedRecording field to given value.

### HasEnableDetailedRecording

`func (o *GetOrgKey200Response) HasEnableDetailedRecording() bool`

HasEnableDetailedRecording returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetOrgKey200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetOrgKey200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetOrgKey200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetOrgKey200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetaData

`func (o *GetOrgKey200Response) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *GetOrgKey200Response) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *GetOrgKey200Response) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *GetOrgKey200Response) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *GetOrgKey200Response) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *GetOrgKey200Response) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
### GetOrgId

`func (o *GetOrgKey200Response) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GetOrgKey200Response) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GetOrgKey200Response) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GetOrgKey200Response) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPer

`func (o *GetOrgKey200Response) GetPer() int32`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *GetOrgKey200Response) GetPerOk() (*int32, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *GetOrgKey200Response) SetPer(v int32)`

SetPer sets Per field to given value.

### HasPer

`func (o *GetOrgKey200Response) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *GetOrgKey200Response) GetQuotaMax() int32`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *GetOrgKey200Response) GetQuotaMaxOk() (*int32, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *GetOrgKey200Response) SetQuotaMax(v int32)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *GetOrgKey200Response) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *GetOrgKey200Response) GetQuotaRenewalRate() int32`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *GetOrgKey200Response) GetQuotaRenewalRateOk() (*int32, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *GetOrgKey200Response) SetQuotaRenewalRate(v int32)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *GetOrgKey200Response) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *GetOrgKey200Response) GetQuotaRenews() int32`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *GetOrgKey200Response) GetQuotaRenewsOk() (*int32, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *GetOrgKey200Response) SetQuotaRenews(v int32)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *GetOrgKey200Response) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *GetOrgKey200Response) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetOrgKey200Response) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetOrgKey200Response) SetRate(v int32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetOrgKey200Response) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTags

`func (o *GetOrgKey200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrgKey200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrgKey200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrgKey200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GetOrgKey200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GetOrgKey200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetThrottleInterval

`func (o *GetOrgKey200Response) GetThrottleInterval() int32`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *GetOrgKey200Response) GetThrottleIntervalOk() (*int32, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *GetOrgKey200Response) SetThrottleInterval(v int32)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *GetOrgKey200Response) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *GetOrgKey200Response) GetThrottleRetryLimit() int32`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *GetOrgKey200Response) GetThrottleRetryLimitOk() (*int32, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *GetOrgKey200Response) SetThrottleRetryLimit(v int32)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *GetOrgKey200Response) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


