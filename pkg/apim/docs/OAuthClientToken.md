# OAuthClientToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **int32** |  | [optional] 

## Methods

### NewOAuthClientToken

`func NewOAuthClientToken() *OAuthClientToken`

NewOAuthClientToken instantiates a new OAuthClientToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientTokenWithDefaults

`func NewOAuthClientTokenWithDefaults() *OAuthClientToken`

NewOAuthClientTokenWithDefaults instantiates a new OAuthClientToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OAuthClientToken) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthClientToken) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthClientToken) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OAuthClientToken) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExpires

`func (o *OAuthClientToken) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OAuthClientToken) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OAuthClientToken) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OAuthClientToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


