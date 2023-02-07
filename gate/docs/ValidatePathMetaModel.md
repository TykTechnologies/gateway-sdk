# ValidatePathMetaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorResponseCode** | Pointer to **int64** | Allows override of default 422 Unprocessible Entity response code for validation errors. | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SchemaB64** | Pointer to **string** |  | [optional] 

## Methods

### NewValidatePathMetaModel

`func NewValidatePathMetaModel() *ValidatePathMetaModel`

NewValidatePathMetaModel instantiates a new ValidatePathMetaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePathMetaModelWithDefaults

`func NewValidatePathMetaModelWithDefaults() *ValidatePathMetaModel`

NewValidatePathMetaModelWithDefaults instantiates a new ValidatePathMetaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorResponseCode

`func (o *ValidatePathMetaModel) GetErrorResponseCode() int64`

GetErrorResponseCode returns the ErrorResponseCode field if non-nil, zero value otherwise.

### GetErrorResponseCodeOk

`func (o *ValidatePathMetaModel) GetErrorResponseCodeOk() (*int64, bool)`

GetErrorResponseCodeOk returns a tuple with the ErrorResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponseCode

`func (o *ValidatePathMetaModel) SetErrorResponseCode(v int64)`

SetErrorResponseCode sets ErrorResponseCode field to given value.

### HasErrorResponseCode

`func (o *ValidatePathMetaModel) HasErrorResponseCode() bool`

HasErrorResponseCode returns a boolean if a field has been set.

### GetMethod

`func (o *ValidatePathMetaModel) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ValidatePathMetaModel) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ValidatePathMetaModel) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ValidatePathMetaModel) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ValidatePathMetaModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidatePathMetaModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidatePathMetaModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ValidatePathMetaModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSchema

`func (o *ValidatePathMetaModel) GetSchema() map[string]map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ValidatePathMetaModel) GetSchemaOk() (*map[string]map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ValidatePathMetaModel) SetSchema(v map[string]map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ValidatePathMetaModel) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaB64

`func (o *ValidatePathMetaModel) GetSchemaB64() string`

GetSchemaB64 returns the SchemaB64 field if non-nil, zero value otherwise.

### GetSchemaB64Ok

`func (o *ValidatePathMetaModel) GetSchemaB64Ok() (*string, bool)`

GetSchemaB64Ok returns a tuple with the SchemaB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaB64

`func (o *ValidatePathMetaModel) SetSchemaB64(v string)`

SetSchemaB64 sets SchemaB64 field to given value.

### HasSchemaB64

`func (o *ValidatePathMetaModel) HasSchemaB64() bool`

HasSchemaB64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


