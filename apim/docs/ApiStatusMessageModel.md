# ApiStatusMessageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Response details | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStatusMessageModel

`func NewApiStatusMessageModel() *ApiStatusMessageModel`

NewApiStatusMessageModel instantiates a new ApiStatusMessageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusMessageModelWithDefaults

`func NewApiStatusMessageModelWithDefaults() *ApiStatusMessageModel`

NewApiStatusMessageModelWithDefaults instantiates a new ApiStatusMessageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiStatusMessageModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiStatusMessageModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiStatusMessageModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiStatusMessageModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ApiStatusMessageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStatusMessageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStatusMessageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiStatusMessageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


