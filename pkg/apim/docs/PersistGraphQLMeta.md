# PersistGraphQLMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPersistGraphQLMeta

`func NewPersistGraphQLMeta() *PersistGraphQLMeta`

NewPersistGraphQLMeta instantiates a new PersistGraphQLMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistGraphQLMetaWithDefaults

`func NewPersistGraphQLMetaWithDefaults() *PersistGraphQLMeta`

NewPersistGraphQLMetaWithDefaults instantiates a new PersistGraphQLMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PersistGraphQLMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PersistGraphQLMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PersistGraphQLMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PersistGraphQLMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOperation

`func (o *PersistGraphQLMeta) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PersistGraphQLMeta) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PersistGraphQLMeta) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PersistGraphQLMeta) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPath

`func (o *PersistGraphQLMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PersistGraphQLMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PersistGraphQLMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PersistGraphQLMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVariables

`func (o *PersistGraphQLMeta) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PersistGraphQLMeta) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PersistGraphQLMeta) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PersistGraphQLMeta) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *PersistGraphQLMeta) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *PersistGraphQLMeta) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


