# BasicAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashType** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewBasicAuthData

`func NewBasicAuthData() *BasicAuthData`

NewBasicAuthData instantiates a new BasicAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicAuthDataWithDefaults

`func NewBasicAuthDataWithDefaults() *BasicAuthData`

NewBasicAuthDataWithDefaults instantiates a new BasicAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashType

`func (o *BasicAuthData) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *BasicAuthData) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *BasicAuthData) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *BasicAuthData) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetPassword

`func (o *BasicAuthData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicAuthData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicAuthData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BasicAuthData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


