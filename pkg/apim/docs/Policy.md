# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MID** | Pointer to **string** |  | [optional] 
**AccessRights** | Pointer to [**map[string]AccessDefinition**](AccessDefinition.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**EnableHttpSignatureValidation** | Pointer to **bool** |  | [optional] 
**GraphqlAccessRights** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**HmacEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsInactive** | Pointer to **bool** |  | [optional] 
**KeyExpiresIn** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**MaxQueryDepth** | Pointer to **int32** |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Partitions** | Pointer to [**PolicyPartitions**](PolicyPartitions.md) |  | [optional] 
**Per** | Pointer to **float64** |  | [optional] 
**QuotaMax** | Pointer to **int64** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int64** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**Smoothing** | Pointer to [**RateLimitSmoothing**](RateLimitSmoothing.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThrottleInterval** | Pointer to **float64** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMID

`func (o *Policy) GetMID() string`

GetMID returns the MID field if non-nil, zero value otherwise.

### GetMIDOk

`func (o *Policy) GetMIDOk() (*string, bool)`

GetMIDOk returns a tuple with the MID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMID

`func (o *Policy) SetMID(v string)`

SetMID sets MID field to given value.

### HasMID

`func (o *Policy) HasMID() bool`

HasMID returns a boolean if a field has been set.

### GetAccessRights

`func (o *Policy) GetAccessRights() map[string]AccessDefinition`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *Policy) GetAccessRightsOk() (*map[string]AccessDefinition, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *Policy) SetAccessRights(v map[string]AccessDefinition)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *Policy) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### SetAccessRightsNil

`func (o *Policy) SetAccessRightsNil(b bool)`

 SetAccessRightsNil sets the value for AccessRights to be an explicit nil

### UnsetAccessRights
`func (o *Policy) UnsetAccessRights()`

UnsetAccessRights ensures that no value is present for AccessRights, not even an explicit nil
### GetActive

`func (o *Policy) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Policy) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Policy) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Policy) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnableHttpSignatureValidation

`func (o *Policy) GetEnableHttpSignatureValidation() bool`

GetEnableHttpSignatureValidation returns the EnableHttpSignatureValidation field if non-nil, zero value otherwise.

### GetEnableHttpSignatureValidationOk

`func (o *Policy) GetEnableHttpSignatureValidationOk() (*bool, bool)`

GetEnableHttpSignatureValidationOk returns a tuple with the EnableHttpSignatureValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttpSignatureValidation

`func (o *Policy) SetEnableHttpSignatureValidation(v bool)`

SetEnableHttpSignatureValidation sets EnableHttpSignatureValidation field to given value.

### HasEnableHttpSignatureValidation

`func (o *Policy) HasEnableHttpSignatureValidation() bool`

HasEnableHttpSignatureValidation returns a boolean if a field has been set.

### GetGraphqlAccessRights

`func (o *Policy) GetGraphqlAccessRights() map[string]map[string]interface{}`

GetGraphqlAccessRights returns the GraphqlAccessRights field if non-nil, zero value otherwise.

### GetGraphqlAccessRightsOk

`func (o *Policy) GetGraphqlAccessRightsOk() (*map[string]map[string]interface{}, bool)`

GetGraphqlAccessRightsOk returns a tuple with the GraphqlAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlAccessRights

`func (o *Policy) SetGraphqlAccessRights(v map[string]map[string]interface{})`

SetGraphqlAccessRights sets GraphqlAccessRights field to given value.

### HasGraphqlAccessRights

`func (o *Policy) HasGraphqlAccessRights() bool`

HasGraphqlAccessRights returns a boolean if a field has been set.

### SetGraphqlAccessRightsNil

`func (o *Policy) SetGraphqlAccessRightsNil(b bool)`

 SetGraphqlAccessRightsNil sets the value for GraphqlAccessRights to be an explicit nil

### UnsetGraphqlAccessRights
`func (o *Policy) UnsetGraphqlAccessRights()`

UnsetGraphqlAccessRights ensures that no value is present for GraphqlAccessRights, not even an explicit nil
### GetHmacEnabled

`func (o *Policy) GetHmacEnabled() bool`

GetHmacEnabled returns the HmacEnabled field if non-nil, zero value otherwise.

### GetHmacEnabledOk

`func (o *Policy) GetHmacEnabledOk() (*bool, bool)`

GetHmacEnabledOk returns a tuple with the HmacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacEnabled

`func (o *Policy) SetHmacEnabled(v bool)`

SetHmacEnabled sets HmacEnabled field to given value.

### HasHmacEnabled

`func (o *Policy) HasHmacEnabled() bool`

HasHmacEnabled returns a boolean if a field has been set.

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInactive

`func (o *Policy) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *Policy) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *Policy) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.

### HasIsInactive

`func (o *Policy) HasIsInactive() bool`

HasIsInactive returns a boolean if a field has been set.

### GetKeyExpiresIn

`func (o *Policy) GetKeyExpiresIn() int64`

GetKeyExpiresIn returns the KeyExpiresIn field if non-nil, zero value otherwise.

### GetKeyExpiresInOk

`func (o *Policy) GetKeyExpiresInOk() (*int64, bool)`

GetKeyExpiresInOk returns a tuple with the KeyExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiresIn

`func (o *Policy) SetKeyExpiresIn(v int64)`

SetKeyExpiresIn sets KeyExpiresIn field to given value.

### HasKeyExpiresIn

`func (o *Policy) HasKeyExpiresIn() bool`

HasKeyExpiresIn returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Policy) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Policy) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Policy) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Policy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMaxQueryDepth

`func (o *Policy) GetMaxQueryDepth() int32`

GetMaxQueryDepth returns the MaxQueryDepth field if non-nil, zero value otherwise.

### GetMaxQueryDepthOk

`func (o *Policy) GetMaxQueryDepthOk() (*int32, bool)`

GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryDepth

`func (o *Policy) SetMaxQueryDepth(v int32)`

SetMaxQueryDepth sets MaxQueryDepth field to given value.

### HasMaxQueryDepth

`func (o *Policy) HasMaxQueryDepth() bool`

HasMaxQueryDepth returns a boolean if a field has been set.

### GetMetaData

`func (o *Policy) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *Policy) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *Policy) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *Policy) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *Policy) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *Policy) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Policy) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Policy) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Policy) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Policy) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPartitions

`func (o *Policy) GetPartitions() PolicyPartitions`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Policy) GetPartitionsOk() (*PolicyPartitions, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Policy) SetPartitions(v PolicyPartitions)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *Policy) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPer

`func (o *Policy) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *Policy) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *Policy) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *Policy) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *Policy) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *Policy) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *Policy) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *Policy) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *Policy) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *Policy) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *Policy) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *Policy) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetRate

`func (o *Policy) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Policy) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Policy) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Policy) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSmoothing

`func (o *Policy) GetSmoothing() RateLimitSmoothing`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *Policy) GetSmoothingOk() (*RateLimitSmoothing, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *Policy) SetSmoothing(v RateLimitSmoothing)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *Policy) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### GetTags

`func (o *Policy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Policy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Policy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Policy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Policy) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Policy) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetThrottleInterval

`func (o *Policy) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *Policy) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *Policy) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *Policy) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *Policy) GetThrottleRetryLimit() int32`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *Policy) GetThrottleRetryLimitOk() (*int32, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *Policy) SetThrottleRetryLimit(v int32)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *Policy) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


