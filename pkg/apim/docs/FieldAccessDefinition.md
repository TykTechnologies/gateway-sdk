# FieldAccessDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**FieldLimits**](FieldLimits.md) |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewFieldAccessDefinition

`func NewFieldAccessDefinition() *FieldAccessDefinition`

NewFieldAccessDefinition instantiates a new FieldAccessDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldAccessDefinitionWithDefaults

`func NewFieldAccessDefinitionWithDefaults() *FieldAccessDefinition`

NewFieldAccessDefinitionWithDefaults instantiates a new FieldAccessDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FieldAccessDefinition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FieldAccessDefinition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FieldAccessDefinition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FieldAccessDefinition) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetLimits

`func (o *FieldAccessDefinition) GetLimits() FieldLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *FieldAccessDefinition) GetLimitsOk() (*FieldLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *FieldAccessDefinition) SetLimits(v FieldLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *FieldAccessDefinition) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTypeName

`func (o *FieldAccessDefinition) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FieldAccessDefinition) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FieldAccessDefinition) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *FieldAccessDefinition) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


