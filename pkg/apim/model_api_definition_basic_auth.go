/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation). * Managing and listing policies. * Managing and listing API Definitions (only when not using the Tyk Dashboard). * Hot reloads / reloading a cluster configuration. * OAuth client creation (only when not using the Tyk Dashboard).  In order to use the Gateway API, you'll need to set the **secret** parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  **x-tyk-authorization: <your-secret>*** <br/>  <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the APIDefinitionBasicAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDefinitionBasicAuth{}

// APIDefinitionBasicAuth struct for APIDefinitionBasicAuth
type APIDefinitionBasicAuth struct {
	BodyPasswordRegexp *string `json:"body_password_regexp,omitempty"`
	BodyUserRegexp     *string `json:"body_user_regexp,omitempty"`
	CacheTtl           *int32  `json:"cache_ttl,omitempty"`
	DisableCaching     *bool   `json:"disable_caching,omitempty"`
	ExtractFromBody    *bool   `json:"extract_from_body,omitempty"`
}

// NewAPIDefinitionBasicAuth instantiates a new APIDefinitionBasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionBasicAuth() *APIDefinitionBasicAuth {
	this := APIDefinitionBasicAuth{}
	return &this
}

// NewAPIDefinitionBasicAuthWithDefaults instantiates a new APIDefinitionBasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionBasicAuthWithDefaults() *APIDefinitionBasicAuth {
	this := APIDefinitionBasicAuth{}
	return &this
}

// GetBodyPasswordRegexp returns the BodyPasswordRegexp field value if set, zero value otherwise.
func (o *APIDefinitionBasicAuth) GetBodyPasswordRegexp() string {
	if o == nil || IsNil(o.BodyPasswordRegexp) {
		var ret string
		return ret
	}
	return *o.BodyPasswordRegexp
}

// GetBodyPasswordRegexpOk returns a tuple with the BodyPasswordRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionBasicAuth) GetBodyPasswordRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.BodyPasswordRegexp) {
		return nil, false
	}
	return o.BodyPasswordRegexp, true
}

// HasBodyPasswordRegexp returns a boolean if a field has been set.
func (o *APIDefinitionBasicAuth) HasBodyPasswordRegexp() bool {
	if o != nil && !IsNil(o.BodyPasswordRegexp) {
		return true
	}

	return false
}

// SetBodyPasswordRegexp gets a reference to the given string and assigns it to the BodyPasswordRegexp field.
func (o *APIDefinitionBasicAuth) SetBodyPasswordRegexp(v string) {
	o.BodyPasswordRegexp = &v
}

// GetBodyUserRegexp returns the BodyUserRegexp field value if set, zero value otherwise.
func (o *APIDefinitionBasicAuth) GetBodyUserRegexp() string {
	if o == nil || IsNil(o.BodyUserRegexp) {
		var ret string
		return ret
	}
	return *o.BodyUserRegexp
}

// GetBodyUserRegexpOk returns a tuple with the BodyUserRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionBasicAuth) GetBodyUserRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.BodyUserRegexp) {
		return nil, false
	}
	return o.BodyUserRegexp, true
}

// HasBodyUserRegexp returns a boolean if a field has been set.
func (o *APIDefinitionBasicAuth) HasBodyUserRegexp() bool {
	if o != nil && !IsNil(o.BodyUserRegexp) {
		return true
	}

	return false
}

// SetBodyUserRegexp gets a reference to the given string and assigns it to the BodyUserRegexp field.
func (o *APIDefinitionBasicAuth) SetBodyUserRegexp(v string) {
	o.BodyUserRegexp = &v
}

// GetCacheTtl returns the CacheTtl field value if set, zero value otherwise.
func (o *APIDefinitionBasicAuth) GetCacheTtl() int32 {
	if o == nil || IsNil(o.CacheTtl) {
		var ret int32
		return ret
	}
	return *o.CacheTtl
}

// GetCacheTtlOk returns a tuple with the CacheTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionBasicAuth) GetCacheTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTtl) {
		return nil, false
	}
	return o.CacheTtl, true
}

// HasCacheTtl returns a boolean if a field has been set.
func (o *APIDefinitionBasicAuth) HasCacheTtl() bool {
	if o != nil && !IsNil(o.CacheTtl) {
		return true
	}

	return false
}

// SetCacheTtl gets a reference to the given int32 and assigns it to the CacheTtl field.
func (o *APIDefinitionBasicAuth) SetCacheTtl(v int32) {
	o.CacheTtl = &v
}

// GetDisableCaching returns the DisableCaching field value if set, zero value otherwise.
func (o *APIDefinitionBasicAuth) GetDisableCaching() bool {
	if o == nil || IsNil(o.DisableCaching) {
		var ret bool
		return ret
	}
	return *o.DisableCaching
}

// GetDisableCachingOk returns a tuple with the DisableCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionBasicAuth) GetDisableCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCaching) {
		return nil, false
	}
	return o.DisableCaching, true
}

// HasDisableCaching returns a boolean if a field has been set.
func (o *APIDefinitionBasicAuth) HasDisableCaching() bool {
	if o != nil && !IsNil(o.DisableCaching) {
		return true
	}

	return false
}

// SetDisableCaching gets a reference to the given bool and assigns it to the DisableCaching field.
func (o *APIDefinitionBasicAuth) SetDisableCaching(v bool) {
	o.DisableCaching = &v
}

// GetExtractFromBody returns the ExtractFromBody field value if set, zero value otherwise.
func (o *APIDefinitionBasicAuth) GetExtractFromBody() bool {
	if o == nil || IsNil(o.ExtractFromBody) {
		var ret bool
		return ret
	}
	return *o.ExtractFromBody
}

// GetExtractFromBodyOk returns a tuple with the ExtractFromBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionBasicAuth) GetExtractFromBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.ExtractFromBody) {
		return nil, false
	}
	return o.ExtractFromBody, true
}

// HasExtractFromBody returns a boolean if a field has been set.
func (o *APIDefinitionBasicAuth) HasExtractFromBody() bool {
	if o != nil && !IsNil(o.ExtractFromBody) {
		return true
	}

	return false
}

// SetExtractFromBody gets a reference to the given bool and assigns it to the ExtractFromBody field.
func (o *APIDefinitionBasicAuth) SetExtractFromBody(v bool) {
	o.ExtractFromBody = &v
}

func (o APIDefinitionBasicAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDefinitionBasicAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BodyPasswordRegexp) {
		toSerialize["body_password_regexp"] = o.BodyPasswordRegexp
	}
	if !IsNil(o.BodyUserRegexp) {
		toSerialize["body_user_regexp"] = o.BodyUserRegexp
	}
	if !IsNil(o.CacheTtl) {
		toSerialize["cache_ttl"] = o.CacheTtl
	}
	if !IsNil(o.DisableCaching) {
		toSerialize["disable_caching"] = o.DisableCaching
	}
	if !IsNil(o.ExtractFromBody) {
		toSerialize["extract_from_body"] = o.ExtractFromBody
	}
	return toSerialize, nil
}

type NullableAPIDefinitionBasicAuth struct {
	value *APIDefinitionBasicAuth
	isSet bool
}

func (v NullableAPIDefinitionBasicAuth) Get() *APIDefinitionBasicAuth {
	return v.value
}

func (v *NullableAPIDefinitionBasicAuth) Set(val *APIDefinitionBasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionBasicAuth(val *APIDefinitionBasicAuth) *NullableAPIDefinitionBasicAuth {
	return &NullableAPIDefinitionBasicAuth{value: val, isSet: true}
}

func (v NullableAPIDefinitionBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
