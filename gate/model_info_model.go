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

// InfoModel struct for InfoModel
type InfoModel struct {
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
	TermsOfService *string `json:"termsOfService,omitempty"`
	Contact *ContactModelModel `json:"contact,omitempty"`
	License *LicenseModelModel `json:"license,omitempty"`
	Version string `json:"version"`
}

// NewInfoModel instantiates a new InfoModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoModel(title string, version string) *InfoModel {
	this := InfoModel{}
	this.Title = title
	this.Version = version
	return &this
}

// NewInfoModelWithDefaults instantiates a new InfoModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoModelWithDefaults() *InfoModel {
	this := InfoModel{}
	return &this
}

// GetTitle returns the Title field value
func (o *InfoModel) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InfoModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InfoModel) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfoModel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoModel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfoModel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfoModel) SetDescription(v string) {
	o.Description = &v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise.
func (o *InfoModel) GetTermsOfService() string {
	if o == nil || o.TermsOfService == nil {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoModel) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || o.TermsOfService == nil {
		return nil, false
	}
	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *InfoModel) HasTermsOfService() bool {
	if o != nil && o.TermsOfService != nil {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *InfoModel) SetTermsOfService(v string) {
	o.TermsOfService = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *InfoModel) GetContact() ContactModelModel {
	if o == nil || o.Contact == nil {
		var ret ContactModelModel
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoModel) GetContactOk() (*ContactModelModel, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *InfoModel) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactModelModel and assigns it to the Contact field.
func (o *InfoModel) SetContact(v ContactModelModel) {
	o.Contact = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *InfoModel) GetLicense() LicenseModelModel {
	if o == nil || o.License == nil {
		var ret LicenseModelModel
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoModel) GetLicenseOk() (*LicenseModelModel, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *InfoModel) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given LicenseModelModel and assigns it to the License field.
func (o *InfoModel) SetLicense(v LicenseModelModel) {
	o.License = &v
}

// GetVersion returns the Version field value
func (o *InfoModel) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *InfoModel) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *InfoModel) SetVersion(v string) {
	o.Version = v
}

func (o InfoModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TermsOfService != nil {
		toSerialize["termsOfService"] = o.TermsOfService
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInfoModel struct {
	value *InfoModel
	isSet bool
}

func (v NullableInfoModel) Get() *InfoModel {
	return v.value
}

func (v *NullableInfoModel) Set(val *InfoModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoModel(val *InfoModel) *NullableInfoModel {
	return &NullableInfoModel{value: val, isSet: true}
}

func (v NullableInfoModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


