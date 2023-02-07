# PolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | http://www.mongodb.org/display/DOCS/Object+IDs | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**Per** | Pointer to **float64** |  | [optional] 
**QuotaMax** | Pointer to **int64** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int64** |  | [optional] 
**ThrottleInterval** | Pointer to **float64** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **float32** |  | [optional] 
**MaxQueryDepth** | Pointer to **float32** |  | [optional] 
**AccessRights** | Pointer to [**AccessDefinitionModelModel**](AccessDefinitionModel.md) |  | [optional] 
**HmacEnabled** | Pointer to **bool** |  | [optional] 
**EnableHttpSignatureValidation** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**IsInactive** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**KeyExpiresIn** | Pointer to **float32** |  | [optional] 
**Partitions** | Pointer to [**PolicyPartitionsModelModel**](PolicyPartitionsModel.md) |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**GraphqlAccessRights** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPolicyModel

`func NewPolicyModel() *PolicyModel`

NewPolicyModel instantiates a new PolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyModelWithDefaults

`func NewPolicyModelWithDefaults() *PolicyModel`

NewPolicyModelWithDefaults instantiates a new PolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *PolicyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *PolicyModel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PolicyModel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PolicyModel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PolicyModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRate

`func (o *PolicyModel) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PolicyModel) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PolicyModel) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PolicyModel) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetPer

`func (o *PolicyModel) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *PolicyModel) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *PolicyModel) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *PolicyModel) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *PolicyModel) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *PolicyModel) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *PolicyModel) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *PolicyModel) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *PolicyModel) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *PolicyModel) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *PolicyModel) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *PolicyModel) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetThrottleInterval

`func (o *PolicyModel) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *PolicyModel) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *PolicyModel) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *PolicyModel) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *PolicyModel) GetThrottleRetryLimit() float32`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *PolicyModel) GetThrottleRetryLimitOk() (*float32, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *PolicyModel) SetThrottleRetryLimit(v float32)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *PolicyModel) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.

### GetMaxQueryDepth

`func (o *PolicyModel) GetMaxQueryDepth() float32`

GetMaxQueryDepth returns the MaxQueryDepth field if non-nil, zero value otherwise.

### GetMaxQueryDepthOk

`func (o *PolicyModel) GetMaxQueryDepthOk() (*float32, bool)`

GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryDepth

`func (o *PolicyModel) SetMaxQueryDepth(v float32)`

SetMaxQueryDepth sets MaxQueryDepth field to given value.

### HasMaxQueryDepth

`func (o *PolicyModel) HasMaxQueryDepth() bool`

HasMaxQueryDepth returns a boolean if a field has been set.

### GetAccessRights

`func (o *PolicyModel) GetAccessRights() AccessDefinitionModelModel`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *PolicyModel) GetAccessRightsOk() (*AccessDefinitionModelModel, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *PolicyModel) SetAccessRights(v AccessDefinitionModelModel)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *PolicyModel) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### GetHmacEnabled

`func (o *PolicyModel) GetHmacEnabled() bool`

GetHmacEnabled returns the HmacEnabled field if non-nil, zero value otherwise.

### GetHmacEnabledOk

`func (o *PolicyModel) GetHmacEnabledOk() (*bool, bool)`

GetHmacEnabledOk returns a tuple with the HmacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacEnabled

`func (o *PolicyModel) SetHmacEnabled(v bool)`

SetHmacEnabled sets HmacEnabled field to given value.

### HasHmacEnabled

`func (o *PolicyModel) HasHmacEnabled() bool`

HasHmacEnabled returns a boolean if a field has been set.

### GetEnableHttpSignatureValidation

`func (o *PolicyModel) GetEnableHttpSignatureValidation() bool`

GetEnableHttpSignatureValidation returns the EnableHttpSignatureValidation field if non-nil, zero value otherwise.

### GetEnableHttpSignatureValidationOk

`func (o *PolicyModel) GetEnableHttpSignatureValidationOk() (*bool, bool)`

GetEnableHttpSignatureValidationOk returns a tuple with the EnableHttpSignatureValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttpSignatureValidation

`func (o *PolicyModel) SetEnableHttpSignatureValidation(v bool)`

SetEnableHttpSignatureValidation sets EnableHttpSignatureValidation field to given value.

### HasEnableHttpSignatureValidation

`func (o *PolicyModel) HasEnableHttpSignatureValidation() bool`

HasEnableHttpSignatureValidation returns a boolean if a field has been set.

### GetActive

`func (o *PolicyModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PolicyModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PolicyModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PolicyModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIsInactive

`func (o *PolicyModel) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *PolicyModel) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *PolicyModel) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.

### HasIsInactive

`func (o *PolicyModel) HasIsInactive() bool`

HasIsInactive returns a boolean if a field has been set.

### GetTags

`func (o *PolicyModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PolicyModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PolicyModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PolicyModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetKeyExpiresIn

`func (o *PolicyModel) GetKeyExpiresIn() float32`

GetKeyExpiresIn returns the KeyExpiresIn field if non-nil, zero value otherwise.

### GetKeyExpiresInOk

`func (o *PolicyModel) GetKeyExpiresInOk() (*float32, bool)`

GetKeyExpiresInOk returns a tuple with the KeyExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiresIn

`func (o *PolicyModel) SetKeyExpiresIn(v float32)`

SetKeyExpiresIn sets KeyExpiresIn field to given value.

### HasKeyExpiresIn

`func (o *PolicyModel) HasKeyExpiresIn() bool`

HasKeyExpiresIn returns a boolean if a field has been set.

### GetPartitions

`func (o *PolicyModel) GetPartitions() PolicyPartitionsModelModel`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *PolicyModel) GetPartitionsOk() (*PolicyPartitionsModelModel, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *PolicyModel) SetPartitions(v PolicyPartitionsModelModel)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *PolicyModel) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyModel) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyModel) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyModel) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyModel) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetaData

`func (o *PolicyModel) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *PolicyModel) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *PolicyModel) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *PolicyModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetGraphqlAccessRights

`func (o *PolicyModel) GetGraphqlAccessRights() map[string]interface{}`

GetGraphqlAccessRights returns the GraphqlAccessRights field if non-nil, zero value otherwise.

### GetGraphqlAccessRightsOk

`func (o *PolicyModel) GetGraphqlAccessRightsOk() (*map[string]interface{}, bool)`

GetGraphqlAccessRightsOk returns a tuple with the GraphqlAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlAccessRights

`func (o *PolicyModel) SetGraphqlAccessRights(v map[string]interface{})`

SetGraphqlAccessRights sets GraphqlAccessRights field to given value.

### HasGraphqlAccessRights

`func (o *PolicyModel) HasGraphqlAccessRights() bool`

HasGraphqlAccessRights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


