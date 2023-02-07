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

// APIDefinitionCORS struct for APIDefinitionCORS
type APIDefinitionCORS struct {
	AllowCredentials *bool `json:"allow_credentials,omitempty"`
	AllowedHeaders []string `json:"allowed_headers,omitempty"`
	AllowedMethods []string `json:"allowed_methods,omitempty"`
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	ExposedHeaders []string `json:"exposed_headers,omitempty"`
	MaxAge *int64 `json:"max_age,omitempty"`
	OptionsPassthrough *bool `json:"options_passthrough,omitempty"`
}

// NewAPIDefinitionCORS instantiates a new APIDefinitionCORS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionCORS() *APIDefinitionCORS {
	this := APIDefinitionCORS{}
	return &this
}

// NewAPIDefinitionCORSWithDefaults instantiates a new APIDefinitionCORS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionCORSWithDefaults() *APIDefinitionCORS {
	this := APIDefinitionCORS{}
	return &this
}

// GetAllowCredentials returns the AllowCredentials field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetAllowCredentials() bool {
	if o == nil || isNil(o.AllowCredentials) {
		var ret bool
		return ret
	}
	return *o.AllowCredentials
}

// GetAllowCredentialsOk returns a tuple with the AllowCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetAllowCredentialsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowCredentials) {
    return nil, false
	}
	return o.AllowCredentials, true
}

// HasAllowCredentials returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasAllowCredentials() bool {
	if o != nil && !isNil(o.AllowCredentials) {
		return true
	}

	return false
}

// SetAllowCredentials gets a reference to the given bool and assigns it to the AllowCredentials field.
func (o *APIDefinitionCORS) SetAllowCredentials(v bool) {
	o.AllowCredentials = &v
}

// GetAllowedHeaders returns the AllowedHeaders field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetAllowedHeaders() []string {
	if o == nil || isNil(o.AllowedHeaders) {
		var ret []string
		return ret
	}
	return o.AllowedHeaders
}

// GetAllowedHeadersOk returns a tuple with the AllowedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetAllowedHeadersOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedHeaders) {
    return nil, false
	}
	return o.AllowedHeaders, true
}

// HasAllowedHeaders returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasAllowedHeaders() bool {
	if o != nil && !isNil(o.AllowedHeaders) {
		return true
	}

	return false
}

// SetAllowedHeaders gets a reference to the given []string and assigns it to the AllowedHeaders field.
func (o *APIDefinitionCORS) SetAllowedHeaders(v []string) {
	o.AllowedHeaders = v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetAllowedMethods() []string {
	if o == nil || isNil(o.AllowedMethods) {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedMethods) {
    return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasAllowedMethods() bool {
	if o != nil && !isNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *APIDefinitionCORS) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetAllowedOrigins() []string {
	if o == nil || isNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedOrigins) {
    return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasAllowedOrigins() bool {
	if o != nil && !isNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *APIDefinitionCORS) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetDebug() bool {
	if o == nil || isNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetDebugOk() (*bool, bool) {
	if o == nil || isNil(o.Debug) {
    return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasDebug() bool {
	if o != nil && !isNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *APIDefinitionCORS) SetDebug(v bool) {
	o.Debug = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetEnable() bool {
	if o == nil || isNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetEnableOk() (*bool, bool) {
	if o == nil || isNil(o.Enable) {
    return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *APIDefinitionCORS) SetEnable(v bool) {
	o.Enable = &v
}

// GetExposedHeaders returns the ExposedHeaders field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetExposedHeaders() []string {
	if o == nil || isNil(o.ExposedHeaders) {
		var ret []string
		return ret
	}
	return o.ExposedHeaders
}

// GetExposedHeadersOk returns a tuple with the ExposedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetExposedHeadersOk() ([]string, bool) {
	if o == nil || isNil(o.ExposedHeaders) {
    return nil, false
	}
	return o.ExposedHeaders, true
}

// HasExposedHeaders returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasExposedHeaders() bool {
	if o != nil && !isNil(o.ExposedHeaders) {
		return true
	}

	return false
}

// SetExposedHeaders gets a reference to the given []string and assigns it to the ExposedHeaders field.
func (o *APIDefinitionCORS) SetExposedHeaders(v []string) {
	o.ExposedHeaders = v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetMaxAge() int64 {
	if o == nil || isNil(o.MaxAge) {
		var ret int64
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetMaxAgeOk() (*int64, bool) {
	if o == nil || isNil(o.MaxAge) {
    return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasMaxAge() bool {
	if o != nil && !isNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int64 and assigns it to the MaxAge field.
func (o *APIDefinitionCORS) SetMaxAge(v int64) {
	o.MaxAge = &v
}

// GetOptionsPassthrough returns the OptionsPassthrough field value if set, zero value otherwise.
func (o *APIDefinitionCORS) GetOptionsPassthrough() bool {
	if o == nil || isNil(o.OptionsPassthrough) {
		var ret bool
		return ret
	}
	return *o.OptionsPassthrough
}

// GetOptionsPassthroughOk returns a tuple with the OptionsPassthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionCORS) GetOptionsPassthroughOk() (*bool, bool) {
	if o == nil || isNil(o.OptionsPassthrough) {
    return nil, false
	}
	return o.OptionsPassthrough, true
}

// HasOptionsPassthrough returns a boolean if a field has been set.
func (o *APIDefinitionCORS) HasOptionsPassthrough() bool {
	if o != nil && !isNil(o.OptionsPassthrough) {
		return true
	}

	return false
}

// SetOptionsPassthrough gets a reference to the given bool and assigns it to the OptionsPassthrough field.
func (o *APIDefinitionCORS) SetOptionsPassthrough(v bool) {
	o.OptionsPassthrough = &v
}

func (o APIDefinitionCORS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowCredentials) {
		toSerialize["allow_credentials"] = o.AllowCredentials
	}
	if !isNil(o.AllowedHeaders) {
		toSerialize["allowed_headers"] = o.AllowedHeaders
	}
	if !isNil(o.AllowedMethods) {
		toSerialize["allowed_methods"] = o.AllowedMethods
	}
	if !isNil(o.AllowedOrigins) {
		toSerialize["allowed_origins"] = o.AllowedOrigins
	}
	if !isNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !isNil(o.ExposedHeaders) {
		toSerialize["exposed_headers"] = o.ExposedHeaders
	}
	if !isNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !isNil(o.OptionsPassthrough) {
		toSerialize["options_passthrough"] = o.OptionsPassthrough
	}
	return json.Marshal(toSerialize)
}

type NullableAPIDefinitionCORS struct {
	value *APIDefinitionCORS
	isSet bool
}

func (v NullableAPIDefinitionCORS) Get() *APIDefinitionCORS {
	return v.value
}

func (v *NullableAPIDefinitionCORS) Set(val *APIDefinitionCORS) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionCORS) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionCORS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionCORS(val *APIDefinitionCORS) *NullableAPIDefinitionCORS {
	return &NullableAPIDefinitionCORS{value: val, isSet: true}
}

func (v NullableAPIDefinitionCORS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionCORS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


