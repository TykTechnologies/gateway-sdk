# ResponseProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **interface{}** |  | [optional] 

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

`func (o *ResponseProcessor) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResponseProcessor) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResponseProcessor) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ResponseProcessor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ResponseProcessor) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ResponseProcessor) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


