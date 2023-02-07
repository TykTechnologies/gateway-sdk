# PkixNameModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to **[]string** |  | [optional] 
**OrganizationalUnit** | Pointer to **[]string** |  | [optional] 
**Locality** | Pointer to **[]string** |  | [optional] 
**Province** | Pointer to **[]string** |  | [optional] 
**StreetAddress** | Pointer to **[]string** |  | [optional] 
**PostalCode** | Pointer to **[]string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Names** | Pointer to [**[]PkixAttributeTypeAndValueModelModel**](PkixAttributeTypeAndValueModelModel.md) |  | [optional] 
**ExtraNames** | Pointer to [**[]PkixAttributeTypeAndValueSETModelModel**](PkixAttributeTypeAndValueSETModelModel.md) |  | [optional] 

## Methods

### NewPkixNameModel

`func NewPkixNameModel() *PkixNameModel`

NewPkixNameModel instantiates a new PkixNameModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixNameModelWithDefaults

`func NewPkixNameModelWithDefaults() *PkixNameModel`

NewPkixNameModelWithDefaults instantiates a new PkixNameModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PkixNameModel) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkixNameModel) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkixNameModel) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PkixNameModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganization

`func (o *PkixNameModel) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkixNameModel) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkixNameModel) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PkixNameModel) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *PkixNameModel) GetOrganizationalUnit() []string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *PkixNameModel) GetOrganizationalUnitOk() (*[]string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *PkixNameModel) SetOrganizationalUnit(v []string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *PkixNameModel) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetLocality

`func (o *PkixNameModel) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkixNameModel) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkixNameModel) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PkixNameModel) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetProvince

`func (o *PkixNameModel) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkixNameModel) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkixNameModel) SetProvince(v []string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PkixNameModel) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetStreetAddress

`func (o *PkixNameModel) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkixNameModel) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkixNameModel) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *PkixNameModel) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetPostalCode

`func (o *PkixNameModel) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkixNameModel) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkixNameModel) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PkixNameModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PkixNameModel) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkixNameModel) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkixNameModel) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PkixNameModel) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetCommonName

`func (o *PkixNameModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkixNameModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkixNameModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PkixNameModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetNames

`func (o *PkixNameModel) GetNames() []PkixAttributeTypeAndValueModelModel`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *PkixNameModel) GetNamesOk() (*[]PkixAttributeTypeAndValueModelModel, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *PkixNameModel) SetNames(v []PkixAttributeTypeAndValueModelModel)`

SetNames sets Names field to given value.

### HasNames

`func (o *PkixNameModel) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetExtraNames

`func (o *PkixNameModel) GetExtraNames() []PkixAttributeTypeAndValueSETModelModel`

GetExtraNames returns the ExtraNames field if non-nil, zero value otherwise.

### GetExtraNamesOk

`func (o *PkixNameModel) GetExtraNamesOk() (*[]PkixAttributeTypeAndValueSETModelModel, bool)`

GetExtraNamesOk returns a tuple with the ExtraNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraNames

`func (o *PkixNameModel) SetExtraNames(v []PkixAttributeTypeAndValueSETModelModel)`

SetExtraNames sets ExtraNames field to given value.

### HasExtraNames

`func (o *PkixNameModel) HasExtraNames() bool`

HasExtraNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


