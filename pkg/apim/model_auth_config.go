/*
Tyk Gateway API

 The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway

API version: 5.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the AuthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthConfig{}

// AuthConfig struct for AuthConfig
type AuthConfig struct {
	AuthHeaderName *string `json:"auth_header_name,omitempty"`
	CookieName *string `json:"cookie_name,omitempty"`
	DisableHeader *bool `json:"disable_header,omitempty"`
	Name *string `json:"name,omitempty"`
	ParamName *string `json:"param_name,omitempty"`
	Signature *SignatureConfig `json:"signature,omitempty"`
	UseCertificate *bool `json:"use_certificate,omitempty"`
	UseCookie *bool `json:"use_cookie,omitempty"`
	UseParam *bool `json:"use_param,omitempty"`
	ValidateSignature *bool `json:"validate_signature,omitempty"`
}

// NewAuthConfig instantiates a new AuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthConfig() *AuthConfig {
	this := AuthConfig{}
	return &this
}

// NewAuthConfigWithDefaults instantiates a new AuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthConfigWithDefaults() *AuthConfig {
	this := AuthConfig{}
	return &this
}

// GetAuthHeaderName returns the AuthHeaderName field value if set, zero value otherwise.
func (o *AuthConfig) GetAuthHeaderName() string {
	if o == nil || IsNil(o.AuthHeaderName) {
		var ret string
		return ret
	}
	return *o.AuthHeaderName
}

// GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetAuthHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthHeaderName) {
		return nil, false
	}
	return o.AuthHeaderName, true
}

// HasAuthHeaderName returns a boolean if a field has been set.
func (o *AuthConfig) HasAuthHeaderName() bool {
	if o != nil && !IsNil(o.AuthHeaderName) {
		return true
	}

	return false
}

// SetAuthHeaderName gets a reference to the given string and assigns it to the AuthHeaderName field.
func (o *AuthConfig) SetAuthHeaderName(v string) {
	o.AuthHeaderName = &v
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *AuthConfig) GetCookieName() string {
	if o == nil || IsNil(o.CookieName) {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetCookieNameOk() (*string, bool) {
	if o == nil || IsNil(o.CookieName) {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *AuthConfig) HasCookieName() bool {
	if o != nil && !IsNil(o.CookieName) {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *AuthConfig) SetCookieName(v string) {
	o.CookieName = &v
}

// GetDisableHeader returns the DisableHeader field value if set, zero value otherwise.
func (o *AuthConfig) GetDisableHeader() bool {
	if o == nil || IsNil(o.DisableHeader) {
		var ret bool
		return ret
	}
	return *o.DisableHeader
}

// GetDisableHeaderOk returns a tuple with the DisableHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetDisableHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableHeader) {
		return nil, false
	}
	return o.DisableHeader, true
}

// HasDisableHeader returns a boolean if a field has been set.
func (o *AuthConfig) HasDisableHeader() bool {
	if o != nil && !IsNil(o.DisableHeader) {
		return true
	}

	return false
}

// SetDisableHeader gets a reference to the given bool and assigns it to the DisableHeader field.
func (o *AuthConfig) SetDisableHeader(v bool) {
	o.DisableHeader = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthConfig) SetName(v string) {
	o.Name = &v
}

// GetParamName returns the ParamName field value if set, zero value otherwise.
func (o *AuthConfig) GetParamName() string {
	if o == nil || IsNil(o.ParamName) {
		var ret string
		return ret
	}
	return *o.ParamName
}

// GetParamNameOk returns a tuple with the ParamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetParamNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParamName) {
		return nil, false
	}
	return o.ParamName, true
}

// HasParamName returns a boolean if a field has been set.
func (o *AuthConfig) HasParamName() bool {
	if o != nil && !IsNil(o.ParamName) {
		return true
	}

	return false
}

// SetParamName gets a reference to the given string and assigns it to the ParamName field.
func (o *AuthConfig) SetParamName(v string) {
	o.ParamName = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *AuthConfig) GetSignature() SignatureConfig {
	if o == nil || IsNil(o.Signature) {
		var ret SignatureConfig
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetSignatureOk() (*SignatureConfig, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *AuthConfig) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given SignatureConfig and assigns it to the Signature field.
func (o *AuthConfig) SetSignature(v SignatureConfig) {
	o.Signature = &v
}

// GetUseCertificate returns the UseCertificate field value if set, zero value otherwise.
func (o *AuthConfig) GetUseCertificate() bool {
	if o == nil || IsNil(o.UseCertificate) {
		var ret bool
		return ret
	}
	return *o.UseCertificate
}

// GetUseCertificateOk returns a tuple with the UseCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetUseCertificateOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCertificate) {
		return nil, false
	}
	return o.UseCertificate, true
}

// HasUseCertificate returns a boolean if a field has been set.
func (o *AuthConfig) HasUseCertificate() bool {
	if o != nil && !IsNil(o.UseCertificate) {
		return true
	}

	return false
}

// SetUseCertificate gets a reference to the given bool and assigns it to the UseCertificate field.
func (o *AuthConfig) SetUseCertificate(v bool) {
	o.UseCertificate = &v
}

// GetUseCookie returns the UseCookie field value if set, zero value otherwise.
func (o *AuthConfig) GetUseCookie() bool {
	if o == nil || IsNil(o.UseCookie) {
		var ret bool
		return ret
	}
	return *o.UseCookie
}

// GetUseCookieOk returns a tuple with the UseCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetUseCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCookie) {
		return nil, false
	}
	return o.UseCookie, true
}

// HasUseCookie returns a boolean if a field has been set.
func (o *AuthConfig) HasUseCookie() bool {
	if o != nil && !IsNil(o.UseCookie) {
		return true
	}

	return false
}

// SetUseCookie gets a reference to the given bool and assigns it to the UseCookie field.
func (o *AuthConfig) SetUseCookie(v bool) {
	o.UseCookie = &v
}

// GetUseParam returns the UseParam field value if set, zero value otherwise.
func (o *AuthConfig) GetUseParam() bool {
	if o == nil || IsNil(o.UseParam) {
		var ret bool
		return ret
	}
	return *o.UseParam
}

// GetUseParamOk returns a tuple with the UseParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetUseParamOk() (*bool, bool) {
	if o == nil || IsNil(o.UseParam) {
		return nil, false
	}
	return o.UseParam, true
}

// HasUseParam returns a boolean if a field has been set.
func (o *AuthConfig) HasUseParam() bool {
	if o != nil && !IsNil(o.UseParam) {
		return true
	}

	return false
}

// SetUseParam gets a reference to the given bool and assigns it to the UseParam field.
func (o *AuthConfig) SetUseParam(v bool) {
	o.UseParam = &v
}

// GetValidateSignature returns the ValidateSignature field value if set, zero value otherwise.
func (o *AuthConfig) GetValidateSignature() bool {
	if o == nil || IsNil(o.ValidateSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateSignature
}

// GetValidateSignatureOk returns a tuple with the ValidateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfig) GetValidateSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateSignature) {
		return nil, false
	}
	return o.ValidateSignature, true
}

// HasValidateSignature returns a boolean if a field has been set.
func (o *AuthConfig) HasValidateSignature() bool {
	if o != nil && !IsNil(o.ValidateSignature) {
		return true
	}

	return false
}

// SetValidateSignature gets a reference to the given bool and assigns it to the ValidateSignature field.
func (o *AuthConfig) SetValidateSignature(v bool) {
	o.ValidateSignature = &v
}

func (o AuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthHeaderName) {
		toSerialize["auth_header_name"] = o.AuthHeaderName
	}
	if !IsNil(o.CookieName) {
		toSerialize["cookie_name"] = o.CookieName
	}
	if !IsNil(o.DisableHeader) {
		toSerialize["disable_header"] = o.DisableHeader
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParamName) {
		toSerialize["param_name"] = o.ParamName
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.UseCertificate) {
		toSerialize["use_certificate"] = o.UseCertificate
	}
	if !IsNil(o.UseCookie) {
		toSerialize["use_cookie"] = o.UseCookie
	}
	if !IsNil(o.UseParam) {
		toSerialize["use_param"] = o.UseParam
	}
	if !IsNil(o.ValidateSignature) {
		toSerialize["validate_signature"] = o.ValidateSignature
	}
	return toSerialize, nil
}

type NullableAuthConfig struct {
	value *AuthConfig
	isSet bool
}

func (v NullableAuthConfig) Get() *AuthConfig {
	return v.value
}

func (v *NullableAuthConfig) Set(val *AuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthConfig(val *AuthConfig) *NullableAuthConfig {
	return &NullableAuthConfig{value: val, isSet: true}
}

func (v NullableAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


