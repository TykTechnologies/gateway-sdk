# VersionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FallbackToDefault** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StripPath** | Pointer to **bool** |  | [optional] 
**StripVersioningData** | Pointer to **bool** |  | [optional] 
**UrlVersioningPattern** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewVersionDefinition

`func NewVersionDefinition() *VersionDefinition`

NewVersionDefinition instantiates a new VersionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDefinitionWithDefaults

`func NewVersionDefinitionWithDefaults() *VersionDefinition`

NewVersionDefinitionWithDefaults instantiates a new VersionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *VersionDefinition) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *VersionDefinition) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *VersionDefinition) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *VersionDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnabled

`func (o *VersionDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VersionDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VersionDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VersionDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbackToDefault

`func (o *VersionDefinition) GetFallbackToDefault() bool`

GetFallbackToDefault returns the FallbackToDefault field if non-nil, zero value otherwise.

### GetFallbackToDefaultOk

`func (o *VersionDefinition) GetFallbackToDefaultOk() (*bool, bool)`

GetFallbackToDefaultOk returns a tuple with the FallbackToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToDefault

`func (o *VersionDefinition) SetFallbackToDefault(v bool)`

SetFallbackToDefault sets FallbackToDefault field to given value.

### HasFallbackToDefault

`func (o *VersionDefinition) HasFallbackToDefault() bool`

HasFallbackToDefault returns a boolean if a field has been set.

### GetKey

`func (o *VersionDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VersionDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VersionDefinition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VersionDefinition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *VersionDefinition) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VersionDefinition) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VersionDefinition) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VersionDefinition) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *VersionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStripPath

`func (o *VersionDefinition) GetStripPath() bool`

GetStripPath returns the StripPath field if non-nil, zero value otherwise.

### GetStripPathOk

`func (o *VersionDefinition) GetStripPathOk() (*bool, bool)`

GetStripPathOk returns a tuple with the StripPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripPath

`func (o *VersionDefinition) SetStripPath(v bool)`

SetStripPath sets StripPath field to given value.

### HasStripPath

`func (o *VersionDefinition) HasStripPath() bool`

HasStripPath returns a boolean if a field has been set.

### GetStripVersioningData

`func (o *VersionDefinition) GetStripVersioningData() bool`

GetStripVersioningData returns the StripVersioningData field if non-nil, zero value otherwise.

### GetStripVersioningDataOk

`func (o *VersionDefinition) GetStripVersioningDataOk() (*bool, bool)`

GetStripVersioningDataOk returns a tuple with the StripVersioningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripVersioningData

`func (o *VersionDefinition) SetStripVersioningData(v bool)`

SetStripVersioningData sets StripVersioningData field to given value.

### HasStripVersioningData

`func (o *VersionDefinition) HasStripVersioningData() bool`

HasStripVersioningData returns a boolean if a field has been set.

### GetUrlVersioningPattern

`func (o *VersionDefinition) GetUrlVersioningPattern() string`

GetUrlVersioningPattern returns the UrlVersioningPattern field if non-nil, zero value otherwise.

### GetUrlVersioningPatternOk

`func (o *VersionDefinition) GetUrlVersioningPatternOk() (*string, bool)`

GetUrlVersioningPatternOk returns a tuple with the UrlVersioningPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlVersioningPattern

`func (o *VersionDefinition) SetUrlVersioningPattern(v string)`

SetUrlVersioningPattern sets UrlVersioningPattern field to given value.

### HasUrlVersioningPattern

`func (o *VersionDefinition) HasUrlVersioningPattern() bool`

HasUrlVersioningPattern returns a boolean if a field has been set.

### GetVersions

`func (o *VersionDefinition) GetVersions() map[string]string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *VersionDefinition) GetVersionsOk() (*map[string]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *VersionDefinition) SetVersions(v map[string]string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *VersionDefinition) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *VersionDefinition) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *VersionDefinition) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


