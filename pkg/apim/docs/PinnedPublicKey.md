# PinnedPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**PublicKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPinnedPublicKey

`func NewPinnedPublicKey() *PinnedPublicKey`

NewPinnedPublicKey instantiates a new PinnedPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedPublicKeyWithDefaults

`func NewPinnedPublicKeyWithDefaults() *PinnedPublicKey`

NewPinnedPublicKeyWithDefaults instantiates a new PinnedPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PinnedPublicKey) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PinnedPublicKey) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PinnedPublicKey) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PinnedPublicKey) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPublicKeys

`func (o *PinnedPublicKey) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *PinnedPublicKey) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *PinnedPublicKey) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *PinnedPublicKey) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### SetPublicKeysNil

`func (o *PinnedPublicKey) SetPublicKeysNil(b bool)`

 SetPublicKeysNil sets the value for PublicKeys to be an explicit nil

### UnsetPublicKeys
`func (o *PinnedPublicKey) UnsetPublicKeys()`

UnsetPublicKeys ensures that no value is present for PublicKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


