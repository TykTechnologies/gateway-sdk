# TraceHttpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewTraceHttpRequest

`func NewTraceHttpRequest() *TraceHttpRequest`

NewTraceHttpRequest instantiates a new TraceHttpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceHttpRequestWithDefaults

`func NewTraceHttpRequestWithDefaults() *TraceHttpRequest`

NewTraceHttpRequestWithDefaults instantiates a new TraceHttpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TraceHttpRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TraceHttpRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TraceHttpRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TraceHttpRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *TraceHttpRequest) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TraceHttpRequest) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TraceHttpRequest) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TraceHttpRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *TraceHttpRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TraceHttpRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TraceHttpRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TraceHttpRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *TraceHttpRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TraceHttpRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TraceHttpRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TraceHttpRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


