# TraceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**TraceHttpRequest**](TraceHttpRequest.md) |  | [optional] 
**Spec** | Pointer to [**APIDefinition**](APIDefinition.md) |  | [optional] 
**Oas** | Pointer to [**NullableOneOf30XTykAPIGateway**](oneOf&lt;30,XTykAPIGateway&gt;.md) |  | [optional] 

## Methods

### NewTraceRequest

`func NewTraceRequest() *TraceRequest`

NewTraceRequest instantiates a new TraceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceRequestWithDefaults

`func NewTraceRequestWithDefaults() *TraceRequest`

NewTraceRequestWithDefaults instantiates a new TraceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *TraceRequest) GetRequest() TraceHttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *TraceRequest) GetRequestOk() (*TraceHttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *TraceRequest) SetRequest(v TraceHttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *TraceRequest) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSpec

`func (o *TraceRequest) GetSpec() APIDefinition`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TraceRequest) GetSpecOk() (*APIDefinition, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TraceRequest) SetSpec(v APIDefinition)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TraceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetOas

`func (o *TraceRequest) GetOas() OneOf30XTykAPIGateway`

GetOas returns the Oas field if non-nil, zero value otherwise.

### GetOasOk

`func (o *TraceRequest) GetOasOk() (*OneOf30XTykAPIGateway, bool)`

GetOasOk returns a tuple with the Oas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOas

`func (o *TraceRequest) SetOas(v OneOf30XTykAPIGateway)`

SetOas sets Oas field to given value.

### HasOas

`func (o *TraceRequest) HasOas() bool`

HasOas returns a boolean if a field has been set.

### SetOasNil

`func (o *TraceRequest) SetOasNil(b bool)`

 SetOasNil sets the value for Oas to be an explicit nil

### UnsetOas
`func (o *TraceRequest) UnsetOas()`

UnsetOas ensures that no value is present for Oas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


