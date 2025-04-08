# ScopeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeClaimName** | Pointer to **string** |  | [optional] 
**ScopeToPolicy** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScopeClaim

`func NewScopeClaim() *ScopeClaim`

NewScopeClaim instantiates a new ScopeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeClaimWithDefaults

`func NewScopeClaimWithDefaults() *ScopeClaim`

NewScopeClaimWithDefaults instantiates a new ScopeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeClaimName

`func (o *ScopeClaim) GetScopeClaimName() string`

GetScopeClaimName returns the ScopeClaimName field if non-nil, zero value otherwise.

### GetScopeClaimNameOk

`func (o *ScopeClaim) GetScopeClaimNameOk() (*string, bool)`

GetScopeClaimNameOk returns a tuple with the ScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClaimName

`func (o *ScopeClaim) SetScopeClaimName(v string)`

SetScopeClaimName sets ScopeClaimName field to given value.

### HasScopeClaimName

`func (o *ScopeClaim) HasScopeClaimName() bool`

HasScopeClaimName returns a boolean if a field has been set.

### GetScopeToPolicy

`func (o *ScopeClaim) GetScopeToPolicy() map[string]string`

GetScopeToPolicy returns the ScopeToPolicy field if non-nil, zero value otherwise.

### GetScopeToPolicyOk

`func (o *ScopeClaim) GetScopeToPolicyOk() (*map[string]string, bool)`

GetScopeToPolicyOk returns a tuple with the ScopeToPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeToPolicy

`func (o *ScopeClaim) SetScopeToPolicy(v map[string]string)`

SetScopeToPolicy sets ScopeToPolicy field to given value.

### HasScopeToPolicy

`func (o *ScopeClaim) HasScopeToPolicy() bool`

HasScopeToPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


