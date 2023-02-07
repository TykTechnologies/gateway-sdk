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

// MiddlewareSectionModel struct for MiddlewareSectionModel
type MiddlewareSectionModel struct {
	AuthCheck *MiddlewareDefinitionModelModel `json:"auth_check,omitempty"`
	Driver *string `json:"driver,omitempty"`
	IdExtractor *MiddlewareIdExtractorModelModel `json:"id_extractor,omitempty"`
	Post []MiddlewareDefinitionModelModel `json:"post,omitempty"`
	PostKeyAuth []MiddlewareDefinitionModelModel `json:"post_key_auth,omitempty"`
	Pre []MiddlewareDefinitionModelModel `json:"pre,omitempty"`
	Response []MiddlewareDefinitionModelModel `json:"response,omitempty"`
}

// NewMiddlewareSectionModel instantiates a new MiddlewareSectionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareSectionModel() *MiddlewareSectionModel {
	this := MiddlewareSectionModel{}
	return &this
}

// NewMiddlewareSectionModelWithDefaults instantiates a new MiddlewareSectionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareSectionModelWithDefaults() *MiddlewareSectionModel {
	this := MiddlewareSectionModel{}
	return &this
}

// GetAuthCheck returns the AuthCheck field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetAuthCheck() MiddlewareDefinitionModelModel {
	if o == nil || o.AuthCheck == nil {
		var ret MiddlewareDefinitionModelModel
		return ret
	}
	return *o.AuthCheck
}

// GetAuthCheckOk returns a tuple with the AuthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetAuthCheckOk() (*MiddlewareDefinitionModelModel, bool) {
	if o == nil || o.AuthCheck == nil {
		return nil, false
	}
	return o.AuthCheck, true
}

// HasAuthCheck returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasAuthCheck() bool {
	if o != nil && o.AuthCheck != nil {
		return true
	}

	return false
}

// SetAuthCheck gets a reference to the given MiddlewareDefinitionModelModel and assigns it to the AuthCheck field.
func (o *MiddlewareSectionModel) SetAuthCheck(v MiddlewareDefinitionModelModel) {
	o.AuthCheck = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *MiddlewareSectionModel) SetDriver(v string) {
	o.Driver = &v
}

// GetIdExtractor returns the IdExtractor field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetIdExtractor() MiddlewareIdExtractorModelModel {
	if o == nil || o.IdExtractor == nil {
		var ret MiddlewareIdExtractorModelModel
		return ret
	}
	return *o.IdExtractor
}

// GetIdExtractorOk returns a tuple with the IdExtractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetIdExtractorOk() (*MiddlewareIdExtractorModelModel, bool) {
	if o == nil || o.IdExtractor == nil {
		return nil, false
	}
	return o.IdExtractor, true
}

// HasIdExtractor returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasIdExtractor() bool {
	if o != nil && o.IdExtractor != nil {
		return true
	}

	return false
}

// SetIdExtractor gets a reference to the given MiddlewareIdExtractorModelModel and assigns it to the IdExtractor field.
func (o *MiddlewareSectionModel) SetIdExtractor(v MiddlewareIdExtractorModelModel) {
	o.IdExtractor = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetPost() []MiddlewareDefinitionModelModel {
	if o == nil || o.Post == nil {
		var ret []MiddlewareDefinitionModelModel
		return ret
	}
	return o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetPostOk() ([]MiddlewareDefinitionModelModel, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given []MiddlewareDefinitionModelModel and assigns it to the Post field.
func (o *MiddlewareSectionModel) SetPost(v []MiddlewareDefinitionModelModel) {
	o.Post = v
}

// GetPostKeyAuth returns the PostKeyAuth field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetPostKeyAuth() []MiddlewareDefinitionModelModel {
	if o == nil || o.PostKeyAuth == nil {
		var ret []MiddlewareDefinitionModelModel
		return ret
	}
	return o.PostKeyAuth
}

// GetPostKeyAuthOk returns a tuple with the PostKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetPostKeyAuthOk() ([]MiddlewareDefinitionModelModel, bool) {
	if o == nil || o.PostKeyAuth == nil {
		return nil, false
	}
	return o.PostKeyAuth, true
}

// HasPostKeyAuth returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasPostKeyAuth() bool {
	if o != nil && o.PostKeyAuth != nil {
		return true
	}

	return false
}

// SetPostKeyAuth gets a reference to the given []MiddlewareDefinitionModelModel and assigns it to the PostKeyAuth field.
func (o *MiddlewareSectionModel) SetPostKeyAuth(v []MiddlewareDefinitionModelModel) {
	o.PostKeyAuth = v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetPre() []MiddlewareDefinitionModelModel {
	if o == nil || o.Pre == nil {
		var ret []MiddlewareDefinitionModelModel
		return ret
	}
	return o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetPreOk() ([]MiddlewareDefinitionModelModel, bool) {
	if o == nil || o.Pre == nil {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasPre() bool {
	if o != nil && o.Pre != nil {
		return true
	}

	return false
}

// SetPre gets a reference to the given []MiddlewareDefinitionModelModel and assigns it to the Pre field.
func (o *MiddlewareSectionModel) SetPre(v []MiddlewareDefinitionModelModel) {
	o.Pre = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *MiddlewareSectionModel) GetResponse() []MiddlewareDefinitionModelModel {
	if o == nil || o.Response == nil {
		var ret []MiddlewareDefinitionModelModel
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSectionModel) GetResponseOk() ([]MiddlewareDefinitionModelModel, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *MiddlewareSectionModel) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []MiddlewareDefinitionModelModel and assigns it to the Response field.
func (o *MiddlewareSectionModel) SetResponse(v []MiddlewareDefinitionModelModel) {
	o.Response = v
}

func (o MiddlewareSectionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthCheck != nil {
		toSerialize["auth_check"] = o.AuthCheck
	}
	if o.Driver != nil {
		toSerialize["driver"] = o.Driver
	}
	if o.IdExtractor != nil {
		toSerialize["id_extractor"] = o.IdExtractor
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	if o.PostKeyAuth != nil {
		toSerialize["post_key_auth"] = o.PostKeyAuth
	}
	if o.Pre != nil {
		toSerialize["pre"] = o.Pre
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableMiddlewareSectionModel struct {
	value *MiddlewareSectionModel
	isSet bool
}

func (v NullableMiddlewareSectionModel) Get() *MiddlewareSectionModel {
	return v.value
}

func (v *NullableMiddlewareSectionModel) Set(val *MiddlewareSectionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareSectionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareSectionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareSectionModel(val *MiddlewareSectionModel) *NullableMiddlewareSectionModel {
	return &NullableMiddlewareSectionModel{value: val, isSet: true}
}

func (v NullableMiddlewareSectionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareSectionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


