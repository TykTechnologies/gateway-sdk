# BatchRequestStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]RequestDefinition**](RequestDefinition.md) |  | [optional] 
**SuppressParallelExecution** | Pointer to **bool** |  | [optional] 

## Methods

### NewBatchRequestStructure

`func NewBatchRequestStructure() *BatchRequestStructure`

NewBatchRequestStructure instantiates a new BatchRequestStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestStructureWithDefaults

`func NewBatchRequestStructureWithDefaults() *BatchRequestStructure`

NewBatchRequestStructureWithDefaults instantiates a new BatchRequestStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *BatchRequestStructure) GetRequests() []RequestDefinition`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BatchRequestStructure) GetRequestsOk() (*[]RequestDefinition, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BatchRequestStructure) SetRequests(v []RequestDefinition)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BatchRequestStructure) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *BatchRequestStructure) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *BatchRequestStructure) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil
### GetSuppressParallelExecution

`func (o *BatchRequestStructure) GetSuppressParallelExecution() bool`

GetSuppressParallelExecution returns the SuppressParallelExecution field if non-nil, zero value otherwise.

### GetSuppressParallelExecutionOk

`func (o *BatchRequestStructure) GetSuppressParallelExecutionOk() (*bool, bool)`

GetSuppressParallelExecutionOk returns a tuple with the SuppressParallelExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressParallelExecution

`func (o *BatchRequestStructure) SetSuppressParallelExecution(v bool)`

SetSuppressParallelExecution sets SuppressParallelExecution field to given value.

### HasSuppressParallelExecution

`func (o *BatchRequestStructure) HasSuppressParallelExecution() bool`

HasSuppressParallelExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


