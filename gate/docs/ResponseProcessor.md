# ResponseProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseProcessor

`func NewResponseProcessor() *ResponseProcessor`

NewResponseProcessor instantiates a new ResponseProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProcessorWithDefaults

`func NewResponseProcessorWithDefaults() *ResponseProcessor`

NewResponseProcessorWithDefaults instantiates a new ResponseProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ResponseProcessor) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResponseProcessor) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResponseProcessor) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ResponseProcessor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


