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

// checks if the NewClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewClientRequest{}

// NewClientRequest struct for NewClientRequest
type NewClientRequest struct {
	ApiId       *string           `json:"api_id,omitempty"`
	ClientId    *string           `json:"client_id,omitempty"`
	Description *string           `json:"description,omitempty"`
	MetaData    map[string]string `json:"meta_data,omitempty"`
	PolicyId    *string           `json:"policy_id,omitempty"`
	RedirectUri *string           `json:"redirect_uri,omitempty"`
	Secret      *string           `json:"secret,omitempty"`
}

// NewNewClientRequest instantiates a new NewClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewClientRequest() *NewClientRequest {
	this := NewClientRequest{}
	return &this
}

// NewNewClientRequestWithDefaults instantiates a new NewClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewClientRequestWithDefaults() *NewClientRequest {
	this := NewClientRequest{}
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *NewClientRequest) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *NewClientRequest) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *NewClientRequest) SetApiId(v string) {
	o.ApiId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *NewClientRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *NewClientRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *NewClientRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewClientRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewClientRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewClientRequest) GetMetaData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewClientRequest) GetMetaDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return &o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *NewClientRequest) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]string and assigns it to the MetaData field.
func (o *NewClientRequest) SetMetaData(v map[string]string) {
	o.MetaData = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *NewClientRequest) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *NewClientRequest) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *NewClientRequest) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *NewClientRequest) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *NewClientRequest) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *NewClientRequest) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NewClientRequest) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewClientRequest) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NewClientRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NewClientRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o NewClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiId) {
		toSerialize["api_id"] = o.ApiId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policy_id"] = o.PolicyId
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableNewClientRequest struct {
	value *NewClientRequest
	isSet bool
}

func (v NullableNewClientRequest) Get() *NewClientRequest {
	return v.value
}

func (v *NullableNewClientRequest) Set(val *NewClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewClientRequest(val *NewClientRequest) *NullableNewClientRequest {
	return &NullableNewClientRequest{value: val, isSet: true}
}

func (v NullableNewClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
