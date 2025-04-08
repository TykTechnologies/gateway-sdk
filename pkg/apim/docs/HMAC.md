# HMAC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSources** | Pointer to [**AuthSources**](AuthSources.md) |  | [optional] 
**AllowedAlgorithms** | Pointer to **[]string** |  | [optional] 
**AllowedClockSkew** | Pointer to **float32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewHMAC

`func NewHMAC() *HMAC`

NewHMAC instantiates a new HMAC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHMACWithDefaults

`func NewHMACWithDefaults() *HMAC`

NewHMACWithDefaults instantiates a new HMAC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSources

`func (o *HMAC) GetAuthSources() AuthSources`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *HMAC) GetAuthSourcesOk() (*AuthSources, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *HMAC) SetAuthSources(v AuthSources)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *HMAC) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetAllowedAlgorithms

`func (o *HMAC) GetAllowedAlgorithms() []string`

GetAllowedAlgorithms returns the AllowedAlgorithms field if non-nil, zero value otherwise.

### GetAllowedAlgorithmsOk

`func (o *HMAC) GetAllowedAlgorithmsOk() (*[]string, bool)`

GetAllowedAlgorithmsOk returns a tuple with the AllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAlgorithms

`func (o *HMAC) SetAllowedAlgorithms(v []string)`

SetAllowedAlgorithms sets AllowedAlgorithms field to given value.

### HasAllowedAlgorithms

`func (o *HMAC) HasAllowedAlgorithms() bool`

HasAllowedAlgorithms returns a boolean if a field has been set.

### GetAllowedClockSkew

`func (o *HMAC) GetAllowedClockSkew() float32`

GetAllowedClockSkew returns the AllowedClockSkew field if non-nil, zero value otherwise.

### GetAllowedClockSkewOk

`func (o *HMAC) GetAllowedClockSkewOk() (*float32, bool)`

GetAllowedClockSkewOk returns a tuple with the AllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClockSkew

`func (o *HMAC) SetAllowedClockSkew(v float32)`

SetAllowedClockSkew sets AllowedClockSkew field to given value.

### HasAllowedClockSkew

`func (o *HMAC) HasAllowedClockSkew() bool`

HasAllowedClockSkew returns a boolean if a field has been set.

### GetEnabled

`func (o *HMAC) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HMAC) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HMAC) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HMAC) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


