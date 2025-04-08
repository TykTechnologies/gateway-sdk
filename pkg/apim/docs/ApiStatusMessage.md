# ApiStatusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStatusMessage

`func NewApiStatusMessage() *ApiStatusMessage`

NewApiStatusMessage instantiates a new ApiStatusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusMessageWithDefaults

`func NewApiStatusMessageWithDefaults() *ApiStatusMessage`

NewApiStatusMessageWithDefaults instantiates a new ApiStatusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiStatusMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiStatusMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiStatusMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiStatusMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ApiStatusMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStatusMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStatusMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiStatusMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


