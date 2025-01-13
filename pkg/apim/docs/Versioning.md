# Versioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FallbackToDefault** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StripVersioningData** | Pointer to **bool** |  | [optional] 
**UrlVersioningPattern** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to [**[]VersionToID**](VersionToID.md) |  | [optional] 

## Methods

### NewVersioning

`func NewVersioning() *Versioning`

NewVersioning instantiates a new Versioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersioningWithDefaults

`func NewVersioningWithDefaults() *Versioning`

NewVersioningWithDefaults instantiates a new Versioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *Versioning) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Versioning) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Versioning) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Versioning) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnabled

`func (o *Versioning) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Versioning) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Versioning) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Versioning) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbackToDefault

`func (o *Versioning) GetFallbackToDefault() bool`

GetFallbackToDefault returns the FallbackToDefault field if non-nil, zero value otherwise.

### GetFallbackToDefaultOk

`func (o *Versioning) GetFallbackToDefaultOk() (*bool, bool)`

GetFallbackToDefaultOk returns a tuple with the FallbackToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToDefault

`func (o *Versioning) SetFallbackToDefault(v bool)`

SetFallbackToDefault sets FallbackToDefault field to given value.

### HasFallbackToDefault

`func (o *Versioning) HasFallbackToDefault() bool`

HasFallbackToDefault returns a boolean if a field has been set.

### GetKey

`func (o *Versioning) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Versioning) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Versioning) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Versioning) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *Versioning) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Versioning) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Versioning) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Versioning) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *Versioning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Versioning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Versioning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Versioning) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStripVersioningData

`func (o *Versioning) GetStripVersioningData() bool`

GetStripVersioningData returns the StripVersioningData field if non-nil, zero value otherwise.

### GetStripVersioningDataOk

`func (o *Versioning) GetStripVersioningDataOk() (*bool, bool)`

GetStripVersioningDataOk returns a tuple with the StripVersioningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripVersioningData

`func (o *Versioning) SetStripVersioningData(v bool)`

SetStripVersioningData sets StripVersioningData field to given value.

### HasStripVersioningData

`func (o *Versioning) HasStripVersioningData() bool`

HasStripVersioningData returns a boolean if a field has been set.

### GetUrlVersioningPattern

`func (o *Versioning) GetUrlVersioningPattern() string`

GetUrlVersioningPattern returns the UrlVersioningPattern field if non-nil, zero value otherwise.

### GetUrlVersioningPatternOk

`func (o *Versioning) GetUrlVersioningPatternOk() (*string, bool)`

GetUrlVersioningPatternOk returns a tuple with the UrlVersioningPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlVersioningPattern

`func (o *Versioning) SetUrlVersioningPattern(v string)`

SetUrlVersioningPattern sets UrlVersioningPattern field to given value.

### HasUrlVersioningPattern

`func (o *Versioning) HasUrlVersioningPattern() bool`

HasUrlVersioningPattern returns a boolean if a field has been set.

### GetVersions

`func (o *Versioning) GetVersions() []VersionToID`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Versioning) GetVersionsOk() (*[]VersionToID, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Versioning) SetVersions(v []VersionToID)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Versioning) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *Versioning) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *Versioning) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


