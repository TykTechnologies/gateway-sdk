/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// URLRewriteMetaModel struct for URLRewriteMetaModel
type URLRewriteMetaModel struct {
	MatchRegexp *RegexpModelModel `json:"MatchRegexp,omitempty"`
	MatchPattern *string `json:"match_pattern,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
	RewriteTo *string `json:"rewrite_to,omitempty"`
	Triggers []RoutingTriggerModelModel `json:"triggers,omitempty"`
}

// NewURLRewriteMetaModel instantiates a new URLRewriteMetaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewURLRewriteMetaModel() *URLRewriteMetaModel {
	this := URLRewriteMetaModel{}
	return &this
}

// NewURLRewriteMetaModelWithDefaults instantiates a new URLRewriteMetaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewURLRewriteMetaModelWithDefaults() *URLRewriteMetaModel {
	this := URLRewriteMetaModel{}
	return &this
}

// GetMatchRegexp returns the MatchRegexp field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetMatchRegexp() RegexpModelModel {
	if o == nil || o.MatchRegexp == nil {
		var ret RegexpModelModel
		return ret
	}
	return *o.MatchRegexp
}

// GetMatchRegexpOk returns a tuple with the MatchRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetMatchRegexpOk() (*RegexpModelModel, bool) {
	if o == nil || o.MatchRegexp == nil {
		return nil, false
	}
	return o.MatchRegexp, true
}

// HasMatchRegexp returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasMatchRegexp() bool {
	if o != nil && o.MatchRegexp != nil {
		return true
	}

	return false
}

// SetMatchRegexp gets a reference to the given RegexpModelModel and assigns it to the MatchRegexp field.
func (o *URLRewriteMetaModel) SetMatchRegexp(v RegexpModelModel) {
	o.MatchRegexp = &v
}

// GetMatchPattern returns the MatchPattern field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetMatchPattern() string {
	if o == nil || o.MatchPattern == nil {
		var ret string
		return ret
	}
	return *o.MatchPattern
}

// GetMatchPatternOk returns a tuple with the MatchPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetMatchPatternOk() (*string, bool) {
	if o == nil || o.MatchPattern == nil {
		return nil, false
	}
	return o.MatchPattern, true
}

// HasMatchPattern returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasMatchPattern() bool {
	if o != nil && o.MatchPattern != nil {
		return true
	}

	return false
}

// SetMatchPattern gets a reference to the given string and assigns it to the MatchPattern field.
func (o *URLRewriteMetaModel) SetMatchPattern(v string) {
	o.MatchPattern = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *URLRewriteMetaModel) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *URLRewriteMetaModel) SetPath(v string) {
	o.Path = &v
}

// GetRewriteTo returns the RewriteTo field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetRewriteTo() string {
	if o == nil || o.RewriteTo == nil {
		var ret string
		return ret
	}
	return *o.RewriteTo
}

// GetRewriteToOk returns a tuple with the RewriteTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetRewriteToOk() (*string, bool) {
	if o == nil || o.RewriteTo == nil {
		return nil, false
	}
	return o.RewriteTo, true
}

// HasRewriteTo returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasRewriteTo() bool {
	if o != nil && o.RewriteTo != nil {
		return true
	}

	return false
}

// SetRewriteTo gets a reference to the given string and assigns it to the RewriteTo field.
func (o *URLRewriteMetaModel) SetRewriteTo(v string) {
	o.RewriteTo = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *URLRewriteMetaModel) GetTriggers() []RoutingTriggerModelModel {
	if o == nil || o.Triggers == nil {
		var ret []RoutingTriggerModelModel
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteMetaModel) GetTriggersOk() ([]RoutingTriggerModelModel, bool) {
	if o == nil || o.Triggers == nil {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *URLRewriteMetaModel) HasTriggers() bool {
	if o != nil && o.Triggers != nil {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RoutingTriggerModelModel and assigns it to the Triggers field.
func (o *URLRewriteMetaModel) SetTriggers(v []RoutingTriggerModelModel) {
	o.Triggers = v
}

func (o URLRewriteMetaModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchRegexp != nil {
		toSerialize["MatchRegexp"] = o.MatchRegexp
	}
	if o.MatchPattern != nil {
		toSerialize["match_pattern"] = o.MatchPattern
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.RewriteTo != nil {
		toSerialize["rewrite_to"] = o.RewriteTo
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}
	return json.Marshal(toSerialize)
}

type NullableURLRewriteMetaModel struct {
	value *URLRewriteMetaModel
	isSet bool
}

func (v NullableURLRewriteMetaModel) Get() *URLRewriteMetaModel {
	return v.value
}

func (v *NullableURLRewriteMetaModel) Set(val *URLRewriteMetaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableURLRewriteMetaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableURLRewriteMetaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableURLRewriteMetaModel(val *URLRewriteMetaModel) *NullableURLRewriteMetaModel {
	return &NullableURLRewriteMetaModel{value: val, isSet: true}
}

func (v NullableURLRewriteMetaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableURLRewriteMetaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


