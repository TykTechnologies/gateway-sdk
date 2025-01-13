# AuthProviderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StorageEngine** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthProviderMeta

`func NewAuthProviderMeta() *AuthProviderMeta`

NewAuthProviderMeta instantiates a new AuthProviderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthProviderMetaWithDefaults

`func NewAuthProviderMetaWithDefaults() *AuthProviderMeta`

NewAuthProviderMetaWithDefaults instantiates a new AuthProviderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AuthProviderMeta) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthProviderMeta) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthProviderMeta) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthProviderMeta) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *AuthProviderMeta) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *AuthProviderMeta) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetName

`func (o *AuthProviderMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthProviderMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthProviderMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthProviderMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageEngine

`func (o *AuthProviderMeta) GetStorageEngine() string`

GetStorageEngine returns the StorageEngine field if non-nil, zero value otherwise.

### GetStorageEngineOk

`func (o *AuthProviderMeta) GetStorageEngineOk() (*string, bool)`

GetStorageEngineOk returns a tuple with the StorageEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEngine

`func (o *AuthProviderMeta) SetStorageEngine(v string)`

SetStorageEngine sets StorageEngine field to given value.

### HasStorageEngine

`func (o *AuthProviderMeta) HasStorageEngine() bool`

HasStorageEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


