# OASSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOASSchemaResponse

`func NewOASSchemaResponse() *OASSchemaResponse`

NewOASSchemaResponse instantiates a new OASSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOASSchemaResponseWithDefaults

`func NewOASSchemaResponseWithDefaults() *OASSchemaResponse`

NewOASSchemaResponseWithDefaults instantiates a new OASSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *OASSchemaResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OASSchemaResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OASSchemaResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OASSchemaResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSchema

`func (o *OASSchemaResponse) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OASSchemaResponse) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OASSchemaResponse) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OASSchemaResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *OASSchemaResponse) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *OASSchemaResponse) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetStatus

`func (o *OASSchemaResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OASSchemaResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OASSchemaResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OASSchemaResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


