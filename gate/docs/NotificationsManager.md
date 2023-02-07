# NotificationsManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthOnKeychangeUrl** | Pointer to **string** |  | [optional] 
**SharedSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationsManager

`func NewNotificationsManager() *NotificationsManager`

NewNotificationsManager instantiates a new NotificationsManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsManagerWithDefaults

`func NewNotificationsManagerWithDefaults() *NotificationsManager`

NewNotificationsManagerWithDefaults instantiates a new NotificationsManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthOnKeychangeUrl

`func (o *NotificationsManager) GetOauthOnKeychangeUrl() string`

GetOauthOnKeychangeUrl returns the OauthOnKeychangeUrl field if non-nil, zero value otherwise.

### GetOauthOnKeychangeUrlOk

`func (o *NotificationsManager) GetOauthOnKeychangeUrlOk() (*string, bool)`

GetOauthOnKeychangeUrlOk returns a tuple with the OauthOnKeychangeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthOnKeychangeUrl

`func (o *NotificationsManager) SetOauthOnKeychangeUrl(v string)`

SetOauthOnKeychangeUrl sets OauthOnKeychangeUrl field to given value.

### HasOauthOnKeychangeUrl

`func (o *NotificationsManager) HasOauthOnKeychangeUrl() bool`

HasOauthOnKeychangeUrl returns a boolean if a field has been set.

### GetSharedSecret

`func (o *NotificationsManager) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *NotificationsManager) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *NotificationsManager) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *NotificationsManager) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


