# RequestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**RelativeUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestDefinition

`func NewRequestDefinition() *RequestDefinition`

NewRequestDefinition instantiates a new RequestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDefinitionWithDefaults

`func NewRequestDefinitionWithDefaults() *RequestDefinition`

NewRequestDefinitionWithDefaults instantiates a new RequestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *RequestDefinition) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RequestDefinition) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RequestDefinition) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *RequestDefinition) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *RequestDefinition) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RequestDefinition) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RequestDefinition) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RequestDefinition) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *RequestDefinition) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *RequestDefinition) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *RequestDefinition) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestDefinition) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestDefinition) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestDefinition) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRelativeUrl

`func (o *RequestDefinition) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *RequestDefinition) GetRelativeUrlOk() (*string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeUrl

`func (o *RequestDefinition) SetRelativeUrl(v string)`

SetRelativeUrl sets RelativeUrl field to given value.

### HasRelativeUrl

`func (o *RequestDefinition) HasRelativeUrl() bool`

HasRelativeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


