# OAuthClientTokenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **int64** |  | [optional] 

## Methods

### NewOAuthClientTokenModel

`func NewOAuthClientTokenModel() *OAuthClientTokenModel`

NewOAuthClientTokenModel instantiates a new OAuthClientTokenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientTokenModelWithDefaults

`func NewOAuthClientTokenModelWithDefaults() *OAuthClientTokenModel`

NewOAuthClientTokenModelWithDefaults instantiates a new OAuthClientTokenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OAuthClientTokenModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthClientTokenModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthClientTokenModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OAuthClientTokenModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExpires

`func (o *OAuthClientTokenModel) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OAuthClientTokenModel) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OAuthClientTokenModel) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OAuthClientTokenModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


