# Info1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TermsOfService** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**License** | Pointer to [**License**](License.md) |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewInfo1

`func NewInfo1(title string, version string, ) *Info1`

NewInfo1 instantiates a new Info1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfo1WithDefaults

`func NewInfo1WithDefaults() *Info1`

NewInfo1WithDefaults instantiates a new Info1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Info1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Info1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Info1) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Info1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Info1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Info1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Info1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTermsOfService

`func (o *Info1) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *Info1) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *Info1) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *Info1) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetContact

`func (o *Info1) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Info1) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Info1) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Info1) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLicense

`func (o *Info1) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Info1) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Info1) SetLicense(v License)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Info1) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetVersion

`func (o *Info1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Info1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Info1) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


