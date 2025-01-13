# Middleware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**Global**](Global.md) |  | [optional] 
**Operations** | Pointer to [**map[string]Operation**](Operation.md) |  | [optional] 

## Methods

### NewMiddleware

`func NewMiddleware() *Middleware`

NewMiddleware instantiates a new Middleware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareWithDefaults

`func NewMiddlewareWithDefaults() *Middleware`

NewMiddlewareWithDefaults instantiates a new Middleware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *Middleware) GetGlobal() Global`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *Middleware) GetGlobalOk() (*Global, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *Middleware) SetGlobal(v Global)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *Middleware) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetOperations

`func (o *Middleware) GetOperations() map[string]Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Middleware) GetOperationsOk() (*map[string]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Middleware) SetOperations(v map[string]Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Middleware) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


