# Info

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

### NewInfo

`func NewInfo(title string, version string, ) *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Info) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Info) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Info) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Info) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Info) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Info) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Info) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTermsOfService

`func (o *Info) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *Info) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *Info) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *Info) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetContact

`func (o *Info) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Info) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Info) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Info) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLicense

`func (o *Info) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Info) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Info) SetLicense(v License)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Info) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetVersion

`func (o *Info) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Info) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Info) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


