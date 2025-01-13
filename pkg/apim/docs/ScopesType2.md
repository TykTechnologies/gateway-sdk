# ScopesType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | Pointer to **string** |  | [optional] 
**ScopeToPolicyMapping** | Pointer to [**[]ScopeToPolicy**](ScopeToPolicy.md) |  | [optional] 

## Methods

### NewScopesType2

`func NewScopesType2() *ScopesType2`

NewScopesType2 instantiates a new ScopesType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesType2WithDefaults

`func NewScopesType2WithDefaults() *ScopesType2`

NewScopesType2WithDefaults instantiates a new ScopesType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *ScopesType2) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ScopesType2) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ScopesType2) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ScopesType2) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetScopeToPolicyMapping

`func (o *ScopesType2) GetScopeToPolicyMapping() []ScopeToPolicy`

GetScopeToPolicyMapping returns the ScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetScopeToPolicyMappingOk

`func (o *ScopesType2) GetScopeToPolicyMappingOk() (*[]ScopeToPolicy, bool)`

GetScopeToPolicyMappingOk returns a tuple with the ScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeToPolicyMapping

`func (o *ScopesType2) SetScopeToPolicyMapping(v []ScopeToPolicy)`

SetScopeToPolicyMapping sets ScopeToPolicyMapping field to given value.

### HasScopeToPolicyMapping

`func (o *ScopesType2) HasScopeToPolicyMapping() bool`

HasScopeToPolicyMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


