# APILimitModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Per** | Pointer to **float64** |  | [optional] 
**QuotaMax** | Pointer to **int64** |  | [optional] 
**QuotaRemaining** | Pointer to **int64** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int64** |  | [optional] 
**QuotaRenews** | Pointer to **int64** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**SetByPolicy** | Pointer to **bool** |  | [optional] 
**ThrottleInterval** | Pointer to **float64** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int64** |  | [optional] 

## Methods

### NewAPILimitModel

`func NewAPILimitModel() *APILimitModel`

NewAPILimitModel instantiates a new APILimitModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPILimitModelWithDefaults

`func NewAPILimitModelWithDefaults() *APILimitModel`

NewAPILimitModelWithDefaults instantiates a new APILimitModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPer

`func (o *APILimitModel) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *APILimitModel) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *APILimitModel) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *APILimitModel) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *APILimitModel) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *APILimitModel) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *APILimitModel) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *APILimitModel) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRemaining

`func (o *APILimitModel) GetQuotaRemaining() int64`

GetQuotaRemaining returns the QuotaRemaining field if non-nil, zero value otherwise.

### GetQuotaRemainingOk

`func (o *APILimitModel) GetQuotaRemainingOk() (*int64, bool)`

GetQuotaRemainingOk returns a tuple with the QuotaRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRemaining

`func (o *APILimitModel) SetQuotaRemaining(v int64)`

SetQuotaRemaining sets QuotaRemaining field to given value.

### HasQuotaRemaining

`func (o *APILimitModel) HasQuotaRemaining() bool`

HasQuotaRemaining returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *APILimitModel) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *APILimitModel) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *APILimitModel) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *APILimitModel) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *APILimitModel) GetQuotaRenews() int64`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *APILimitModel) GetQuotaRenewsOk() (*int64, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *APILimitModel) SetQuotaRenews(v int64)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *APILimitModel) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *APILimitModel) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *APILimitModel) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *APILimitModel) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *APILimitModel) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSetByPolicy

`func (o *APILimitModel) GetSetByPolicy() bool`

GetSetByPolicy returns the SetByPolicy field if non-nil, zero value otherwise.

### GetSetByPolicyOk

`func (o *APILimitModel) GetSetByPolicyOk() (*bool, bool)`

GetSetByPolicyOk returns a tuple with the SetByPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetByPolicy

`func (o *APILimitModel) SetSetByPolicy(v bool)`

SetSetByPolicy sets SetByPolicy field to given value.

### HasSetByPolicy

`func (o *APILimitModel) HasSetByPolicy() bool`

HasSetByPolicy returns a boolean if a field has been set.

### GetThrottleInterval

`func (o *APILimitModel) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *APILimitModel) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *APILimitModel) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *APILimitModel) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *APILimitModel) GetThrottleRetryLimit() int64`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *APILimitModel) GetThrottleRetryLimitOk() (*int64, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *APILimitModel) SetThrottleRetryLimit(v int64)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *APILimitModel) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


