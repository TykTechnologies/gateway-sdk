# IDExtractor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**IDExtractorConfig**](IDExtractorConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**With** | Pointer to **string** |  | [optional] 

## Methods

### NewIDExtractor

`func NewIDExtractor() *IDExtractor`

NewIDExtractor instantiates a new IDExtractor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDExtractorWithDefaults

`func NewIDExtractorWithDefaults() *IDExtractor`

NewIDExtractorWithDefaults instantiates a new IDExtractor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *IDExtractor) GetConfig() IDExtractorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IDExtractor) GetConfigOk() (*IDExtractorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IDExtractor) SetConfig(v IDExtractorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IDExtractor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *IDExtractor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IDExtractor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IDExtractor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IDExtractor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSource

`func (o *IDExtractor) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IDExtractor) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IDExtractor) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IDExtractor) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetWith

`func (o *IDExtractor) GetWith() string`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *IDExtractor) GetWithOk() (*string, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *IDExtractor) SetWith(v string)`

SetWith sets With field to given value.

### HasWith

`func (o *IDExtractor) HasWith() bool`

HasWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


