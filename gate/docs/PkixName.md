# PkixName

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
**Names** | Pointer to [**[]PkixAttributeTypeAndValue**](PkixAttributeTypeAndValue.md) |  | [optional] 
**ExtraNames** | Pointer to [**[]PkixAttributeTypeAndValueSET**](PkixAttributeTypeAndValueSET.md) |  | [optional] 

## Methods

### NewPkixName

`func NewPkixName() *PkixName`

NewPkixName instantiates a new PkixName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixNameWithDefaults

`func NewPkixNameWithDefaults() *PkixName`

NewPkixNameWithDefaults instantiates a new PkixName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PkixName) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkixName) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkixName) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PkixName) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganization

`func (o *PkixName) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkixName) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkixName) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PkixName) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *PkixName) GetOrganizationalUnit() []string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *PkixName) GetOrganizationalUnitOk() (*[]string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *PkixName) SetOrganizationalUnit(v []string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *PkixName) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetLocality

`func (o *PkixName) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkixName) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkixName) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PkixName) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetProvince

`func (o *PkixName) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkixName) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkixName) SetProvince(v []string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PkixName) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetStreetAddress

`func (o *PkixName) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkixName) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkixName) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *PkixName) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetPostalCode

`func (o *PkixName) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkixName) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkixName) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PkixName) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PkixName) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkixName) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkixName) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PkixName) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetCommonName

`func (o *PkixName) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkixName) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkixName) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PkixName) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetNames

`func (o *PkixName) GetNames() []PkixAttributeTypeAndValue`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *PkixName) GetNamesOk() (*[]PkixAttributeTypeAndValue, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *PkixName) SetNames(v []PkixAttributeTypeAndValue)`

SetNames sets Names field to given value.

### HasNames

`func (o *PkixName) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetExtraNames

`func (o *PkixName) GetExtraNames() []PkixAttributeTypeAndValueSET`

GetExtraNames returns the ExtraNames field if non-nil, zero value otherwise.

### GetExtraNamesOk

`func (o *PkixName) GetExtraNamesOk() (*[]PkixAttributeTypeAndValueSET, bool)`

GetExtraNamesOk returns a tuple with the ExtraNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraNames

`func (o *PkixName) SetExtraNames(v []PkixAttributeTypeAndValueSET)`

SetExtraNames sets ExtraNames field to given value.

### HasExtraNames

`func (o *PkixName) HasExtraNames() bool`

HasExtraNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


