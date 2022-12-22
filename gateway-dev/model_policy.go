/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gate

import (
	"encoding/json"
)

// Policy struct for Policy
type Policy struct {
	// http://www.mongodb.org/display/DOCS/Object+IDs
	Id *string `json:"_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Rate *float64 `json:"rate,omitempty"`
	Per *float64 `json:"per,omitempty"`
	QuotaMax *int64 `json:"quota_max,omitempty"`
	QuotaRenewalRate *int64 `json:"quota_renewal_rate,omitempty"`
	ThrottleInterval *float64 `json:"throttle_interval,omitempty"`
	ThrottleRetryLimit *float32 `json:"throttle_retry_limit,omitempty"`
	MaxQueryDepth *float32 `json:"max_query_depth,omitempty"`
	AccessRights AccessDefinition `json:"access_rights,omitempty"`
	HmacEnabled *bool `json:"hmac_enabled,omitempty"`
	EnableHttpSignatureValidation *bool `json:"enable_http_signature_validation,omitempty"`
	Active *bool `json:"active,omitempty"`
	IsInactive *bool `json:"is_inactive,omitempty"`
	Tags []string `json:"tags,omitempty"`
	KeyExpiresIn *float32 `json:"key_expires_in,omitempty"`
	Partitions PolicyPartitions `json:"partitions,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	GraphqlAccessRights map[string]interface{} `json:"graphql_access_rights,omitempty"`
}

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy() *Policy {
	this := Policy{}
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Policy) SetId(v string) {
	o.Id = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Policy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Policy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Policy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Policy) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Policy) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Policy) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Policy) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Policy) GetRate() float64 {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetRateOk() (*float64, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Policy) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *Policy) SetRate(v float64) {
	o.Rate = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Policy) GetPer() float64 {
	if o == nil || o.Per == nil {
		var ret float64
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPerOk() (*float64, bool) {
	if o == nil || o.Per == nil {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Policy) HasPer() bool {
	if o != nil && o.Per != nil {
		return true
	}

	return false
}

// SetPer gets a reference to the given float64 and assigns it to the Per field.
func (o *Policy) SetPer(v float64) {
	o.Per = &v
}

// GetQuotaMax returns the QuotaMax field value if set, zero value otherwise.
func (o *Policy) GetQuotaMax() int64 {
	if o == nil || o.QuotaMax == nil {
		var ret int64
		return ret
	}
	return *o.QuotaMax
}

// GetQuotaMaxOk returns a tuple with the QuotaMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetQuotaMaxOk() (*int64, bool) {
	if o == nil || o.QuotaMax == nil {
		return nil, false
	}
	return o.QuotaMax, true
}

// HasQuotaMax returns a boolean if a field has been set.
func (o *Policy) HasQuotaMax() bool {
	if o != nil && o.QuotaMax != nil {
		return true
	}

	return false
}

// SetQuotaMax gets a reference to the given int64 and assigns it to the QuotaMax field.
func (o *Policy) SetQuotaMax(v int64) {
	o.QuotaMax = &v
}

// GetQuotaRenewalRate returns the QuotaRenewalRate field value if set, zero value otherwise.
func (o *Policy) GetQuotaRenewalRate() int64 {
	if o == nil || o.QuotaRenewalRate == nil {
		var ret int64
		return ret
	}
	return *o.QuotaRenewalRate
}

// GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetQuotaRenewalRateOk() (*int64, bool) {
	if o == nil || o.QuotaRenewalRate == nil {
		return nil, false
	}
	return o.QuotaRenewalRate, true
}

// HasQuotaRenewalRate returns a boolean if a field has been set.
func (o *Policy) HasQuotaRenewalRate() bool {
	if o != nil && o.QuotaRenewalRate != nil {
		return true
	}

	return false
}

// SetQuotaRenewalRate gets a reference to the given int64 and assigns it to the QuotaRenewalRate field.
func (o *Policy) SetQuotaRenewalRate(v int64) {
	o.QuotaRenewalRate = &v
}

// GetThrottleInterval returns the ThrottleInterval field value if set, zero value otherwise.
func (o *Policy) GetThrottleInterval() float64 {
	if o == nil || o.ThrottleInterval == nil {
		var ret float64
		return ret
	}
	return *o.ThrottleInterval
}

// GetThrottleIntervalOk returns a tuple with the ThrottleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetThrottleIntervalOk() (*float64, bool) {
	if o == nil || o.ThrottleInterval == nil {
		return nil, false
	}
	return o.ThrottleInterval, true
}

// HasThrottleInterval returns a boolean if a field has been set.
func (o *Policy) HasThrottleInterval() bool {
	if o != nil && o.ThrottleInterval != nil {
		return true
	}

	return false
}

// SetThrottleInterval gets a reference to the given float64 and assigns it to the ThrottleInterval field.
func (o *Policy) SetThrottleInterval(v float64) {
	o.ThrottleInterval = &v
}

// GetThrottleRetryLimit returns the ThrottleRetryLimit field value if set, zero value otherwise.
func (o *Policy) GetThrottleRetryLimit() float32 {
	if o == nil || o.ThrottleRetryLimit == nil {
		var ret float32
		return ret
	}
	return *o.ThrottleRetryLimit
}

// GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetThrottleRetryLimitOk() (*float32, bool) {
	if o == nil || o.ThrottleRetryLimit == nil {
		return nil, false
	}
	return o.ThrottleRetryLimit, true
}

// HasThrottleRetryLimit returns a boolean if a field has been set.
func (o *Policy) HasThrottleRetryLimit() bool {
	if o != nil && o.ThrottleRetryLimit != nil {
		return true
	}

	return false
}

// SetThrottleRetryLimit gets a reference to the given float32 and assigns it to the ThrottleRetryLimit field.
func (o *Policy) SetThrottleRetryLimit(v float32) {
	o.ThrottleRetryLimit = &v
}

// GetMaxQueryDepth returns the MaxQueryDepth field value if set, zero value otherwise.
func (o *Policy) GetMaxQueryDepth() float32 {
	if o == nil || o.MaxQueryDepth == nil {
		var ret float32
		return ret
	}
	return *o.MaxQueryDepth
}

// GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetMaxQueryDepthOk() (*float32, bool) {
	if o == nil || o.MaxQueryDepth == nil {
		return nil, false
	}
	return o.MaxQueryDepth, true
}

// HasMaxQueryDepth returns a boolean if a field has been set.
func (o *Policy) HasMaxQueryDepth() bool {
	if o != nil && o.MaxQueryDepth != nil {
		return true
	}

	return false
}

// SetMaxQueryDepth gets a reference to the given float32 and assigns it to the MaxQueryDepth field.
func (o *Policy) SetMaxQueryDepth(v float32) {
	o.MaxQueryDepth = &v
}

// GetAccessRights returns the AccessRights field value if set, zero value otherwise.
func (o *Policy) GetAccessRights() AccessDefinition {
	if o == nil || o.AccessRights == nil {
		var ret AccessDefinition
		return ret
	}
	return o.AccessRights
}

// GetAccessRightsOk returns a tuple with the AccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetAccessRightsOk() (AccessDefinition, bool) {
	if o == nil || o.AccessRights == nil {
		return nil, false
	}
	return o.AccessRights, true
}

// HasAccessRights returns a boolean if a field has been set.
func (o *Policy) HasAccessRights() bool {
	if o != nil && o.AccessRights != nil {
		return true
	}

	return false
}

// SetAccessRights gets a reference to the given AccessDefinition and assigns it to the AccessRights field.
func (o *Policy) SetAccessRights(v AccessDefinition) {
	o.AccessRights = v
}

// GetHmacEnabled returns the HmacEnabled field value if set, zero value otherwise.
func (o *Policy) GetHmacEnabled() bool {
	if o == nil || o.HmacEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HmacEnabled
}

// GetHmacEnabledOk returns a tuple with the HmacEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetHmacEnabledOk() (*bool, bool) {
	if o == nil || o.HmacEnabled == nil {
		return nil, false
	}
	return o.HmacEnabled, true
}

// HasHmacEnabled returns a boolean if a field has been set.
func (o *Policy) HasHmacEnabled() bool {
	if o != nil && o.HmacEnabled != nil {
		return true
	}

	return false
}

// SetHmacEnabled gets a reference to the given bool and assigns it to the HmacEnabled field.
func (o *Policy) SetHmacEnabled(v bool) {
	o.HmacEnabled = &v
}

// GetEnableHttpSignatureValidation returns the EnableHttpSignatureValidation field value if set, zero value otherwise.
func (o *Policy) GetEnableHttpSignatureValidation() bool {
	if o == nil || o.EnableHttpSignatureValidation == nil {
		var ret bool
		return ret
	}
	return *o.EnableHttpSignatureValidation
}

// GetEnableHttpSignatureValidationOk returns a tuple with the EnableHttpSignatureValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetEnableHttpSignatureValidationOk() (*bool, bool) {
	if o == nil || o.EnableHttpSignatureValidation == nil {
		return nil, false
	}
	return o.EnableHttpSignatureValidation, true
}

// HasEnableHttpSignatureValidation returns a boolean if a field has been set.
func (o *Policy) HasEnableHttpSignatureValidation() bool {
	if o != nil && o.EnableHttpSignatureValidation != nil {
		return true
	}

	return false
}

// SetEnableHttpSignatureValidation gets a reference to the given bool and assigns it to the EnableHttpSignatureValidation field.
func (o *Policy) SetEnableHttpSignatureValidation(v bool) {
	o.EnableHttpSignatureValidation = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Policy) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Policy) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Policy) SetActive(v bool) {
	o.Active = &v
}

// GetIsInactive returns the IsInactive field value if set, zero value otherwise.
func (o *Policy) GetIsInactive() bool {
	if o == nil || o.IsInactive == nil {
		var ret bool
		return ret
	}
	return *o.IsInactive
}

// GetIsInactiveOk returns a tuple with the IsInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIsInactiveOk() (*bool, bool) {
	if o == nil || o.IsInactive == nil {
		return nil, false
	}
	return o.IsInactive, true
}

// HasIsInactive returns a boolean if a field has been set.
func (o *Policy) HasIsInactive() bool {
	if o != nil && o.IsInactive != nil {
		return true
	}

	return false
}

// SetIsInactive gets a reference to the given bool and assigns it to the IsInactive field.
func (o *Policy) SetIsInactive(v bool) {
	o.IsInactive = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Policy) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Policy) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Policy) SetTags(v []string) {
	o.Tags = v
}

// GetKeyExpiresIn returns the KeyExpiresIn field value if set, zero value otherwise.
func (o *Policy) GetKeyExpiresIn() float32 {
	if o == nil || o.KeyExpiresIn == nil {
		var ret float32
		return ret
	}
	return *o.KeyExpiresIn
}

// GetKeyExpiresInOk returns a tuple with the KeyExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetKeyExpiresInOk() (*float32, bool) {
	if o == nil || o.KeyExpiresIn == nil {
		return nil, false
	}
	return o.KeyExpiresIn, true
}

// HasKeyExpiresIn returns a boolean if a field has been set.
func (o *Policy) HasKeyExpiresIn() bool {
	if o != nil && o.KeyExpiresIn != nil {
		return true
	}

	return false
}

// SetKeyExpiresIn gets a reference to the given float32 and assigns it to the KeyExpiresIn field.
func (o *Policy) SetKeyExpiresIn(v float32) {
	o.KeyExpiresIn = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Policy) GetPartitions() PolicyPartitions {
	if o == nil || o.Partitions == nil {
		var ret PolicyPartitions
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPartitionsOk() (PolicyPartitions, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Policy) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given PolicyPartitions and assigns it to the Partitions field.
func (o *Policy) SetPartitions(v PolicyPartitions) {
	o.Partitions = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Policy) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Policy) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *Policy) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *Policy) GetMetaData() map[string]interface{} {
	if o == nil || o.MetaData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *Policy) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *Policy) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetGraphqlAccessRights returns the GraphqlAccessRights field value if set, zero value otherwise.
func (o *Policy) GetGraphqlAccessRights() map[string]interface{} {
	if o == nil || o.GraphqlAccessRights == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GraphqlAccessRights
}

// GetGraphqlAccessRightsOk returns a tuple with the GraphqlAccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetGraphqlAccessRightsOk() (map[string]interface{}, bool) {
	if o == nil || o.GraphqlAccessRights == nil {
		return nil, false
	}
	return o.GraphqlAccessRights, true
}

// HasGraphqlAccessRights returns a boolean if a field has been set.
func (o *Policy) HasGraphqlAccessRights() bool {
	if o != nil && o.GraphqlAccessRights != nil {
		return true
	}

	return false
}

// SetGraphqlAccessRights gets a reference to the given map[string]interface{} and assigns it to the GraphqlAccessRights field.
func (o *Policy) SetGraphqlAccessRights(v map[string]interface{}) {
	o.GraphqlAccessRights = v
}

func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Per != nil {
		toSerialize["per"] = o.Per
	}
	if o.QuotaMax != nil {
		toSerialize["quota_max"] = o.QuotaMax
	}
	if o.QuotaRenewalRate != nil {
		toSerialize["quota_renewal_rate"] = o.QuotaRenewalRate
	}
	if o.ThrottleInterval != nil {
		toSerialize["throttle_interval"] = o.ThrottleInterval
	}
	if o.ThrottleRetryLimit != nil {
		toSerialize["throttle_retry_limit"] = o.ThrottleRetryLimit
	}
	if o.MaxQueryDepth != nil {
		toSerialize["max_query_depth"] = o.MaxQueryDepth
	}
	if o.AccessRights != nil {
		toSerialize["access_rights"] = o.AccessRights
	}
	if o.HmacEnabled != nil {
		toSerialize["hmac_enabled"] = o.HmacEnabled
	}
	if o.EnableHttpSignatureValidation != nil {
		toSerialize["enable_http_signature_validation"] = o.EnableHttpSignatureValidation
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.IsInactive != nil {
		toSerialize["is_inactive"] = o.IsInactive
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.KeyExpiresIn != nil {
		toSerialize["key_expires_in"] = o.KeyExpiresIn
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.GraphqlAccessRights != nil {
		toSerialize["graphql_access_rights"] = o.GraphqlAccessRights
	}
	return json.Marshal(toSerialize)
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


