# SessionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**AccessRights** | Pointer to [**map[string]AccessDefinition**](AccessDefinition.md) |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Allowance** | Pointer to **float64** |  | [optional] 
**ApplyPolicies** | Pointer to **[]string** |  | [optional] 
**ApplyPolicyId** | Pointer to **string** |  | [optional] 
**BasicAuthData** | Pointer to [**SessionStateBasicAuthData**](SessionStateBasicAuthData.md) |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**DataExpires** | Pointer to **int64** |  | [optional] 
**EnableDetailRecording** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **int64** |  | [optional] 
**HmacEnabled** | Pointer to **bool** |  | [optional] 
**HmacString** | Pointer to **string** |  | [optional] 
**IdExtractorDeadline** | Pointer to **int64** |  | [optional] 
**IsInactive** | Pointer to **bool** |  | [optional] 
**JwtData** | Pointer to [**SessionStateJwtData**](SessionStateJwtData.md) |  | [optional] 
**LastCheck** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Monitor** | Pointer to [**SessionStateMonitor**](SessionStateMonitor.md) |  | [optional] 
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

### NewSessionState

`func NewSessionState() *SessionState`

NewSessionState instantiates a new SessionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStateWithDefaults

`func NewSessionStateWithDefaults() *SessionState`

NewSessionStateWithDefaults instantiates a new SessionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *SessionState) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SessionState) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SessionState) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SessionState) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccessRights

`func (o *SessionState) GetAccessRights() map[string]AccessDefinition`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *SessionState) GetAccessRightsOk() (*map[string]AccessDefinition, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *SessionState) SetAccessRights(v map[string]AccessDefinition)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *SessionState) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### GetAlias

`func (o *SessionState) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SessionState) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SessionState) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SessionState) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAllowance

`func (o *SessionState) GetAllowance() float64`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *SessionState) GetAllowanceOk() (*float64, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *SessionState) SetAllowance(v float64)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *SessionState) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetApplyPolicies

`func (o *SessionState) GetApplyPolicies() []string`

GetApplyPolicies returns the ApplyPolicies field if non-nil, zero value otherwise.

### GetApplyPoliciesOk

`func (o *SessionState) GetApplyPoliciesOk() (*[]string, bool)`

GetApplyPoliciesOk returns a tuple with the ApplyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicies

`func (o *SessionState) SetApplyPolicies(v []string)`

SetApplyPolicies sets ApplyPolicies field to given value.

### HasApplyPolicies

`func (o *SessionState) HasApplyPolicies() bool`

HasApplyPolicies returns a boolean if a field has been set.

### GetApplyPolicyId

`func (o *SessionState) GetApplyPolicyId() string`

GetApplyPolicyId returns the ApplyPolicyId field if non-nil, zero value otherwise.

### GetApplyPolicyIdOk

`func (o *SessionState) GetApplyPolicyIdOk() (*string, bool)`

GetApplyPolicyIdOk returns a tuple with the ApplyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicyId

`func (o *SessionState) SetApplyPolicyId(v string)`

SetApplyPolicyId sets ApplyPolicyId field to given value.

### HasApplyPolicyId

`func (o *SessionState) HasApplyPolicyId() bool`

HasApplyPolicyId returns a boolean if a field has been set.

### GetBasicAuthData

`func (o *SessionState) GetBasicAuthData() SessionStateBasicAuthData`

GetBasicAuthData returns the BasicAuthData field if non-nil, zero value otherwise.

### GetBasicAuthDataOk

`func (o *SessionState) GetBasicAuthDataOk() (*SessionStateBasicAuthData, bool)`

GetBasicAuthDataOk returns a tuple with the BasicAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthData

`func (o *SessionState) SetBasicAuthData(v SessionStateBasicAuthData)`

SetBasicAuthData sets BasicAuthData field to given value.

### HasBasicAuthData

`func (o *SessionState) HasBasicAuthData() bool`

HasBasicAuthData returns a boolean if a field has been set.

### GetCertificate

`func (o *SessionState) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SessionState) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SessionState) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SessionState) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDataExpires

`func (o *SessionState) GetDataExpires() int64`

GetDataExpires returns the DataExpires field if non-nil, zero value otherwise.

### GetDataExpiresOk

`func (o *SessionState) GetDataExpiresOk() (*int64, bool)`

GetDataExpiresOk returns a tuple with the DataExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExpires

`func (o *SessionState) SetDataExpires(v int64)`

SetDataExpires sets DataExpires field to given value.

### HasDataExpires

`func (o *SessionState) HasDataExpires() bool`

HasDataExpires returns a boolean if a field has been set.

### GetEnableDetailRecording

`func (o *SessionState) GetEnableDetailRecording() bool`

GetEnableDetailRecording returns the EnableDetailRecording field if non-nil, zero value otherwise.

### GetEnableDetailRecordingOk

`func (o *SessionState) GetEnableDetailRecordingOk() (*bool, bool)`

GetEnableDetailRecordingOk returns a tuple with the EnableDetailRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDetailRecording

`func (o *SessionState) SetEnableDetailRecording(v bool)`

SetEnableDetailRecording sets EnableDetailRecording field to given value.

### HasEnableDetailRecording

`func (o *SessionState) HasEnableDetailRecording() bool`

HasEnableDetailRecording returns a boolean if a field has been set.

### GetExpires

`func (o *SessionState) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SessionState) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SessionState) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SessionState) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHmacEnabled

`func (o *SessionState) GetHmacEnabled() bool`

GetHmacEnabled returns the HmacEnabled field if non-nil, zero value otherwise.

### GetHmacEnabledOk

`func (o *SessionState) GetHmacEnabledOk() (*bool, bool)`

GetHmacEnabledOk returns a tuple with the HmacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacEnabled

`func (o *SessionState) SetHmacEnabled(v bool)`

SetHmacEnabled sets HmacEnabled field to given value.

### HasHmacEnabled

`func (o *SessionState) HasHmacEnabled() bool`

HasHmacEnabled returns a boolean if a field has been set.

### GetHmacString

`func (o *SessionState) GetHmacString() string`

GetHmacString returns the HmacString field if non-nil, zero value otherwise.

### GetHmacStringOk

`func (o *SessionState) GetHmacStringOk() (*string, bool)`

GetHmacStringOk returns a tuple with the HmacString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacString

`func (o *SessionState) SetHmacString(v string)`

SetHmacString sets HmacString field to given value.

### HasHmacString

`func (o *SessionState) HasHmacString() bool`

HasHmacString returns a boolean if a field has been set.

### GetIdExtractorDeadline

`func (o *SessionState) GetIdExtractorDeadline() int64`

GetIdExtractorDeadline returns the IdExtractorDeadline field if non-nil, zero value otherwise.

### GetIdExtractorDeadlineOk

`func (o *SessionState) GetIdExtractorDeadlineOk() (*int64, bool)`

GetIdExtractorDeadlineOk returns a tuple with the IdExtractorDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractorDeadline

`func (o *SessionState) SetIdExtractorDeadline(v int64)`

SetIdExtractorDeadline sets IdExtractorDeadline field to given value.

### HasIdExtractorDeadline

`func (o *SessionState) HasIdExtractorDeadline() bool`

HasIdExtractorDeadline returns a boolean if a field has been set.

### GetIsInactive

`func (o *SessionState) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *SessionState) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *SessionState) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.

### HasIsInactive

`func (o *SessionState) HasIsInactive() bool`

HasIsInactive returns a boolean if a field has been set.

### GetJwtData

`func (o *SessionState) GetJwtData() SessionStateJwtData`

GetJwtData returns the JwtData field if non-nil, zero value otherwise.

### GetJwtDataOk

`func (o *SessionState) GetJwtDataOk() (*SessionStateJwtData, bool)`

GetJwtDataOk returns a tuple with the JwtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtData

`func (o *SessionState) SetJwtData(v SessionStateJwtData)`

SetJwtData sets JwtData field to given value.

### HasJwtData

`func (o *SessionState) HasJwtData() bool`

HasJwtData returns a boolean if a field has been set.

### GetLastCheck

`func (o *SessionState) GetLastCheck() int64`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *SessionState) GetLastCheckOk() (*int64, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *SessionState) SetLastCheck(v int64)`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *SessionState) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SessionState) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SessionState) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SessionState) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SessionState) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetaData

`func (o *SessionState) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *SessionState) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *SessionState) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *SessionState) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMonitor

`func (o *SessionState) GetMonitor() SessionStateMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *SessionState) GetMonitorOk() (*SessionStateMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *SessionState) SetMonitor(v SessionStateMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *SessionState) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetOauthClientId

`func (o *SessionState) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *SessionState) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *SessionState) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *SessionState) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthKeys

`func (o *SessionState) GetOauthKeys() map[string]string`

GetOauthKeys returns the OauthKeys field if non-nil, zero value otherwise.

### GetOauthKeysOk

`func (o *SessionState) GetOauthKeysOk() (*map[string]string, bool)`

GetOauthKeysOk returns a tuple with the OauthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthKeys

`func (o *SessionState) SetOauthKeys(v map[string]string)`

SetOauthKeys sets OauthKeys field to given value.

### HasOauthKeys

`func (o *SessionState) HasOauthKeys() bool`

HasOauthKeys returns a boolean if a field has been set.

### GetOrgId

`func (o *SessionState) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SessionState) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SessionState) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SessionState) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPer

`func (o *SessionState) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *SessionState) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *SessionState) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *SessionState) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *SessionState) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *SessionState) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *SessionState) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *SessionState) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRemaining

`func (o *SessionState) GetQuotaRemaining() int64`

GetQuotaRemaining returns the QuotaRemaining field if non-nil, zero value otherwise.

### GetQuotaRemainingOk

`func (o *SessionState) GetQuotaRemainingOk() (*int64, bool)`

GetQuotaRemainingOk returns a tuple with the QuotaRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRemaining

`func (o *SessionState) SetQuotaRemaining(v int64)`

SetQuotaRemaining sets QuotaRemaining field to given value.

### HasQuotaRemaining

`func (o *SessionState) HasQuotaRemaining() bool`

HasQuotaRemaining returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *SessionState) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *SessionState) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *SessionState) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *SessionState) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *SessionState) GetQuotaRenews() int64`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *SessionState) GetQuotaRenewsOk() (*int64, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *SessionState) SetQuotaRenews(v int64)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *SessionState) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *SessionState) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SessionState) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SessionState) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *SessionState) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSessionLifetime

`func (o *SessionState) GetSessionLifetime() int64`

GetSessionLifetime returns the SessionLifetime field if non-nil, zero value otherwise.

### GetSessionLifetimeOk

`func (o *SessionState) GetSessionLifetimeOk() (*int64, bool)`

GetSessionLifetimeOk returns a tuple with the SessionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetime

`func (o *SessionState) SetSessionLifetime(v int64)`

SetSessionLifetime sets SessionLifetime field to given value.

### HasSessionLifetime

`func (o *SessionState) HasSessionLifetime() bool`

HasSessionLifetime returns a boolean if a field has been set.

### GetThrottleInterval

`func (o *SessionState) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *SessionState) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *SessionState) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *SessionState) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *SessionState) GetThrottleRetryLimit() int64`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *SessionState) GetThrottleRetryLimitOk() (*int64, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *SessionState) SetThrottleRetryLimit(v int64)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *SessionState) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


