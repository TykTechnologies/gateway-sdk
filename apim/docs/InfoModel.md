# InfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TermsOfService** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to [**ContactModelModel**](ContactModel.md) |  | [optional] 
**License** | Pointer to [**LicenseModelModel**](LicenseModel.md) |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewInfoModel

`func NewInfoModel(title string, version string, ) *InfoModel`

NewInfoModel instantiates a new InfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoModelWithDefaults

`func NewInfoModelWithDefaults() *InfoModel`

NewInfoModelWithDefaults instantiates a new InfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InfoModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InfoModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InfoModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InfoModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfoModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfoModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfoModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTermsOfService

`func (o *InfoModel) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *InfoModel) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *InfoModel) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *InfoModel) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetContact

`func (o *InfoModel) GetContact() ContactModelModel`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *InfoModel) GetContactOk() (*ContactModelModel, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *InfoModel) SetContact(v ContactModelModel)`

SetContact sets Contact field to given value.

### HasContact

`func (o *InfoModel) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLicense

`func (o *InfoModel) GetLicense() LicenseModelModel`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *InfoModel) GetLicenseOk() (*LicenseModelModel, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *InfoModel) SetLicense(v LicenseModelModel)`

SetLicense sets License field to given value.

### HasLicense

`func (o *InfoModel) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetVersion

`func (o *InfoModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfoModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfoModel) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


