# SessionStateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**AccessRights** | Pointer to [**map[string]AccessDefinitionModelModel**](AccessDefinitionModel.md) |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **float64** |  | [optional] 
**ApplyPolicies** | Pointer to **[]string** |  | [optional] 
**ApplyPolicyId** | Pointer to **string** |  | [optional] 
**BasicAuthData** | Pointer to [**SessionStateBasicAuthDataModelModel**](SessionStateBasicAuthDataModel.md) |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**DataExpires** | Pointer to **int64** |  | [optional] 
**EnableDetailRecording** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **int64** |  | [optional] 
**HmacEnabled** | Pointer to **bool** |  | [optional] 
**HmacString** | Pointer to **string** |  | [optional] 
**IdExtractorDeadline** | Pointer to **int64** |  | [optional] 
**IsInactive** | Pointer to **bool** |  | [optional] 
**JwtData** | Pointer to [**SessionStateJwtDataModelModel**](SessionStateJwtDataModel.md) |  | [optional] 
**LastCheck** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Monitor** | Pointer to [**SessionStateMonitorModelModel**](SessionStateMonitorModel.md) |  | [optional] 
**OauthClientId** | Pointer to **string** |  | [optional] 
**OauthKeys** | Pointer to **map[string]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Per** | Pointer to **float64** |  | [optional] 
**QuotaMax** | Pointer to **int64** |  | [optional] 
**QuotaRemaining** | Pointer to **int64** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int64** |  | [optional] 
**QuotaRenews** | Pointer to **int64** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**SessionLifetime** | Pointer to **int64** |  | [optional] 
**ThrottleInterval** | Pointer to **float64** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int64** |  | [optional] 

## Methods

### NewSessionStateModel

`func NewSessionStateModel() *SessionStateModel`

NewSessionStateModel instantiates a new SessionStateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStateModelWithDefaults

`func NewSessionStateModelWithDefaults() *SessionStateModel`

NewSessionStateModelWithDefaults instantiates a new SessionStateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *SessionStateModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SessionStateModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SessionStateModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SessionStateModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccessRights

`func (o *SessionStateModel) GetAccessRights() map[string]AccessDefinitionModelModel`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *SessionStateModel) GetAccessRightsOk() (*map[string]AccessDefinitionModelModel, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *SessionStateModel) SetAccessRights(v map[string]AccessDefinitionModelModel)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *SessionStateModel) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### GetAlias

`func (o *SessionStateModel) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SessionStateModel) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SessionStateModel) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SessionStateModel) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAllowance

`func (o *SessionStateModel) GetAllowance() float64`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *SessionStateModel) GetAllowanceOk() (*float64, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *SessionStateModel) SetAllowance(v float64)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *SessionStateModel) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetApplyPolicies

`func (o *SessionStateModel) GetApplyPolicies() []string`

GetApplyPolicies returns the ApplyPolicies field if non-nil, zero value otherwise.

### GetApplyPoliciesOk

`func (o *SessionStateModel) GetApplyPoliciesOk() (*[]string, bool)`

GetApplyPoliciesOk returns a tuple with the ApplyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicies

`func (o *SessionStateModel) SetApplyPolicies(v []string)`

SetApplyPolicies sets ApplyPolicies field to given value.

### HasApplyPolicies

`func (o *SessionStateModel) HasApplyPolicies() bool`

HasApplyPolicies returns a boolean if a field has been set.

### GetApplyPolicyId

`func (o *SessionStateModel) GetApplyPolicyId() string`

GetApplyPolicyId returns the ApplyPolicyId field if non-nil, zero value otherwise.

### GetApplyPolicyIdOk

`func (o *SessionStateModel) GetApplyPolicyIdOk() (*string, bool)`

GetApplyPolicyIdOk returns a tuple with the ApplyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicyId

`func (o *SessionStateModel) SetApplyPolicyId(v string)`

SetApplyPolicyId sets ApplyPolicyId field to given value.

### HasApplyPolicyId

`func (o *SessionStateModel) HasApplyPolicyId() bool`

HasApplyPolicyId returns a boolean if a field has been set.

### GetBasicAuthData

`func (o *SessionStateModel) GetBasicAuthData() SessionStateBasicAuthDataModelModel`

GetBasicAuthData returns the BasicAuthData field if non-nil, zero value otherwise.

### GetBasicAuthDataOk

`func (o *SessionStateModel) GetBasicAuthDataOk() (*SessionStateBasicAuthDataModelModel, bool)`

GetBasicAuthDataOk returns a tuple with the BasicAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthData

`func (o *SessionStateModel) SetBasicAuthData(v SessionStateBasicAuthDataModelModel)`

SetBasicAuthData sets BasicAuthData field to given value.

### HasBasicAuthData

`func (o *SessionStateModel) HasBasicAuthData() bool`

HasBasicAuthData returns a boolean if a field has been set.

### GetCertificate

`func (o *SessionStateModel) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SessionStateModel) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SessionStateModel) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SessionStateModel) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDataExpires

`func (o *SessionStateModel) GetDataExpires() int64`

GetDataExpires returns the DataExpires field if non-nil, zero value otherwise.

### GetDataExpiresOk

`func (o *SessionStateModel) GetDataExpiresOk() (*int64, bool)`

GetDataExpiresOk returns a tuple with the DataExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExpires

`func (o *SessionStateModel) SetDataExpires(v int64)`

SetDataExpires sets DataExpires field to given value.

### HasDataExpires

`func (o *SessionStateModel) HasDataExpires() bool`

HasDataExpires returns a boolean if a field has been set.

### GetEnableDetailRecording

`func (o *SessionStateModel) GetEnableDetailRecording() bool`

GetEnableDetailRecording returns the EnableDetailRecording field if non-nil, zero value otherwise.

### GetEnableDetailRecordingOk

`func (o *SessionStateModel) GetEnableDetailRecordingOk() (*bool, bool)`

GetEnableDetailRecordingOk returns a tuple with the EnableDetailRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDetailRecording

`func (o *SessionStateModel) SetEnableDetailRecording(v bool)`

SetEnableDetailRecording sets EnableDetailRecording field to given value.

### HasEnableDetailRecording

`func (o *SessionStateModel) HasEnableDetailRecording() bool`

HasEnableDetailRecording returns a boolean if a field has been set.

### GetExpires

`func (o *SessionStateModel) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SessionStateModel) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SessionStateModel) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SessionStateModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHmacEnabled

`func (o *SessionStateModel) GetHmacEnabled() bool`

GetHmacEnabled returns the HmacEnabled field if non-nil, zero value otherwise.

### GetHmacEnabledOk

`func (o *SessionStateModel) GetHmacEnabledOk() (*bool, bool)`

GetHmacEnabledOk returns a tuple with the HmacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacEnabled

`func (o *SessionStateModel) SetHmacEnabled(v bool)`

SetHmacEnabled sets HmacEnabled field to given value.

### HasHmacEnabled

`func (o *SessionStateModel) HasHmacEnabled() bool`

HasHmacEnabled returns a boolean if a field has been set.

### GetHmacString

`func (o *SessionStateModel) GetHmacString() string`

GetHmacString returns the HmacString field if non-nil, zero value otherwise.

### GetHmacStringOk

`func (o *SessionStateModel) GetHmacStringOk() (*string, bool)`

GetHmacStringOk returns a tuple with the HmacString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacString

`func (o *SessionStateModel) SetHmacString(v string)`

SetHmacString sets HmacString field to given value.

### HasHmacString

`func (o *SessionStateModel) HasHmacString() bool`

HasHmacString returns a boolean if a field has been set.

### GetIdExtractorDeadline

`func (o *SessionStateModel) GetIdExtractorDeadline() int64`

GetIdExtractorDeadline returns the IdExtractorDeadline field if non-nil, zero value otherwise.

### GetIdExtractorDeadlineOk

`func (o *SessionStateModel) GetIdExtractorDeadlineOk() (*int64, bool)`

GetIdExtractorDeadlineOk returns a tuple with the IdExtractorDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractorDeadline

`func (o *SessionStateModel) SetIdExtractorDeadline(v int64)`

SetIdExtractorDeadline sets IdExtractorDeadline field to given value.

### HasIdExtractorDeadline

`func (o *SessionStateModel) HasIdExtractorDeadline() bool`

HasIdExtractorDeadline returns a boolean if a field has been set.

### GetIsInactive

`func (o *SessionStateModel) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *SessionStateModel) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *SessionStateModel) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.

### HasIsInactive

`func (o *SessionStateModel) HasIsInactive() bool`

HasIsInactive returns a boolean if a field has been set.

### GetJwtData

`func (o *SessionStateModel) GetJwtData() SessionStateJwtDataModelModel`

GetJwtData returns the JwtData field if non-nil, zero value otherwise.

### GetJwtDataOk

`func (o *SessionStateModel) GetJwtDataOk() (*SessionStateJwtDataModelModel, bool)`

GetJwtDataOk returns a tuple with the JwtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtData

`func (o *SessionStateModel) SetJwtData(v SessionStateJwtDataModelModel)`

SetJwtData sets JwtData field to given value.

### HasJwtData

`func (o *SessionStateModel) HasJwtData() bool`

HasJwtData returns a boolean if a field has been set.

### GetLastCheck

`func (o *SessionStateModel) GetLastCheck() int64`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *SessionStateModel) GetLastCheckOk() (*int64, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *SessionStateModel) SetLastCheck(v int64)`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *SessionStateModel) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SessionStateModel) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SessionStateModel) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SessionStateModel) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SessionStateModel) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetaData

`func (o *SessionStateModel) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *SessionStateModel) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *SessionStateModel) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *SessionStateModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMonitor

`func (o *SessionStateModel) GetMonitor() SessionStateMonitorModelModel`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *SessionStateModel) GetMonitorOk() (*SessionStateMonitorModelModel, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *SessionStateModel) SetMonitor(v SessionStateMonitorModelModel)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *SessionStateModel) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetOauthClientId

`func (o *SessionStateModel) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *SessionStateModel) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *SessionStateModel) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *SessionStateModel) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthKeys

`func (o *SessionStateModel) GetOauthKeys() map[string]string`

GetOauthKeys returns the OauthKeys field if non-nil, zero value otherwise.

### GetOauthKeysOk

`func (o *SessionStateModel) GetOauthKeysOk() (*map[string]string, bool)`

GetOauthKeysOk returns a tuple with the OauthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthKeys

`func (o *SessionStateModel) SetOauthKeys(v map[string]string)`

SetOauthKeys sets OauthKeys field to given value.

### HasOauthKeys

`func (o *SessionStateModel) HasOauthKeys() bool`

HasOauthKeys returns a boolean if a field has been set.

### GetOrgId

`func (o *SessionStateModel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SessionStateModel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SessionStateModel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SessionStateModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPer

`func (o *SessionStateModel) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *SessionStateModel) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *SessionStateModel) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *SessionStateModel) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *SessionStateModel) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *SessionStateModel) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *SessionStateModel) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *SessionStateModel) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRemaining

`func (o *SessionStateModel) GetQuotaRemaining() int64`

GetQuotaRemaining returns the QuotaRemaining field if non-nil, zero value otherwise.

### GetQuotaRemainingOk

`func (o *SessionStateModel) GetQuotaRemainingOk() (*int64, bool)`

GetQuotaRemainingOk returns a tuple with the QuotaRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRemaining

`func (o *SessionStateModel) SetQuotaRemaining(v int64)`

SetQuotaRemaining sets QuotaRemaining field to given value.

### HasQuotaRemaining

`func (o *SessionStateModel) HasQuotaRemaining() bool`

HasQuotaRemaining returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *SessionStateModel) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *SessionStateModel) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *SessionStateModel) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *SessionStateModel) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *SessionStateModel) GetQuotaRenews() int64`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *SessionStateModel) GetQuotaRenewsOk() (*int64, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *SessionStateModel) SetQuotaRenews(v int64)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *SessionStateModel) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *SessionStateModel) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SessionStateModel) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SessionStateModel) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *SessionStateModel) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSessionLifetime

`func (o *SessionStateModel) GetSessionLifetime() int64`

GetSessionLifetime returns the SessionLifetime field if non-nil, zero value otherwise.

### GetSessionLifetimeOk

`func (o *SessionStateModel) GetSessionLifetimeOk() (*int64, bool)`

GetSessionLifetimeOk returns a tuple with the SessionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetime

`func (o *SessionStateModel) SetSessionLifetime(v int64)`

SetSessionLifetime sets SessionLifetime field to given value.

### HasSessionLifetime

`func (o *SessionStateModel) HasSessionLifetime() bool`

HasSessionLifetime returns a boolean if a field has been set.

### GetThrottleInterval

`func (o *SessionStateModel) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *SessionStateModel) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *SessionStateModel) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *SessionStateModel) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *SessionStateModel) GetThrottleRetryLimit() int64`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *SessionStateModel) GetThrottleRetryLimitOk() (*int64, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *SessionStateModel) SetThrottleRetryLimit(v int64)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *SessionStateModel) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


