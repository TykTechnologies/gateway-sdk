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

// checks if the RequestSigningMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestSigningMeta{}

// RequestSigningMeta struct for RequestSigningMeta
type RequestSigningMeta struct {
	Algorithm *string `json:"algorithm,omitempty"`
	CertificateId *string `json:"certificate_id,omitempty"`
	HeaderList []string `json:"header_list,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	KeyId *string `json:"key_id,omitempty"`
	Secret *string `json:"secret,omitempty"`
	SignatureHeader *string `json:"signature_header,omitempty"`
}

// NewRequestSigningMeta instantiates a new RequestSigningMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSigningMeta() *RequestSigningMeta {
	this := RequestSigningMeta{}
	return &this
}

// NewRequestSigningMetaWithDefaults instantiates a new RequestSigningMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSigningMetaWithDefaults() *RequestSigningMeta {
	this := RequestSigningMeta{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *RequestSigningMeta) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetCertificateId() string {
	if o == nil || IsNil(o.CertificateId) {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetCertificateIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateId) {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasCertificateId() bool {
	if o != nil && !IsNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *RequestSigningMeta) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetHeaderList returns the HeaderList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSigningMeta) GetHeaderList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HeaderList
}

// GetHeaderListOk returns a tuple with the HeaderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSigningMeta) GetHeaderListOk() ([]string, bool) {
	if o == nil || IsNil(o.HeaderList) {
		return nil, false
	}
	return o.HeaderList, true
}

// HasHeaderList returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasHeaderList() bool {
	if o != nil && !IsNil(o.HeaderList) {
		return true
	}

	return false
}

// SetHeaderList gets a reference to the given []string and assigns it to the HeaderList field.
func (o *RequestSigningMeta) SetHeaderList(v []string) {
	o.HeaderList = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *RequestSigningMeta) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *RequestSigningMeta) SetKeyId(v string) {
	o.KeyId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *RequestSigningMeta) SetSecret(v string) {
	o.Secret = &v
}

// GetSignatureHeader returns the SignatureHeader field value if set, zero value otherwise.
func (o *RequestSigningMeta) GetSignatureHeader() string {
	if o == nil || IsNil(o.SignatureHeader) {
		var ret string
		return ret
	}
	return *o.SignatureHeader
}

// GetSignatureHeaderOk returns a tuple with the SignatureHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSigningMeta) GetSignatureHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureHeader) {
		return nil, false
	}
	return o.SignatureHeader, true
}

// HasSignatureHeader returns a boolean if a field has been set.
func (o *RequestSigningMeta) HasSignatureHeader() bool {
	if o != nil && !IsNil(o.SignatureHeader) {
		return true
	}

	return false
}

// SetSignatureHeader gets a reference to the given string and assigns it to the SignatureHeader field.
func (o *RequestSigningMeta) SetSignatureHeader(v string) {
	o.SignatureHeader = &v
}

func (o RequestSigningMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSigningMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.CertificateId) {
		toSerialize["certificate_id"] = o.CertificateId
	}
	if o.HeaderList != nil {
		toSerialize["header_list"] = o.HeaderList
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.KeyId) {
		toSerialize["key_id"] = o.KeyId
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SignatureHeader) {
		toSerialize["signature_header"] = o.SignatureHeader
	}
	return toSerialize, nil
}

type NullableRequestSigningMeta struct {
	value *RequestSigningMeta
	isSet bool
}

func (v NullableRequestSigningMeta) Get() *RequestSigningMeta {
	return v.value
}

func (v *NullableRequestSigningMeta) Set(val *RequestSigningMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSigningMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSigningMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSigningMeta(val *RequestSigningMeta) *NullableRequestSigningMeta {
	return &NullableRequestSigningMeta{value: val, isSet: true}
}

func (v NullableRequestSigningMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSigningMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


