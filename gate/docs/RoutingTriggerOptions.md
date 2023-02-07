# RoutingTriggerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderMatches** | Pointer to [**map[string]StringRegexMap**](StringRegexMap.md) |  | [optional] 
**PathPartMatches** | Pointer to [**map[string]StringRegexMap**](StringRegexMap.md) |  | [optional] 
**PayloadMatches** | Pointer to [**StringRegexMap**](StringRegexMap.md) |  | [optional] 
**QueryValMatches** | Pointer to [**map[string]StringRegexMap**](StringRegexMap.md) |  | [optional] 
**RequestContextMatches** | Pointer to [**map[string]StringRegexMap**](StringRegexMap.md) |  | [optional] 
**SessionMetaMatches** | Pointer to [**map[string]StringRegexMap**](StringRegexMap.md) |  | [optional] 

## Methods

### NewRoutingTriggerOptions

`func NewRoutingTriggerOptions() *RoutingTriggerOptions`

NewRoutingTriggerOptions instantiates a new RoutingTriggerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingTriggerOptionsWithDefaults

`func NewRoutingTriggerOptionsWithDefaults() *RoutingTriggerOptions`

NewRoutingTriggerOptionsWithDefaults instantiates a new RoutingTriggerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderMatches

`func (o *RoutingTriggerOptions) GetHeaderMatches() map[string]StringRegexMap`

GetHeaderMatches returns the HeaderMatches field if non-nil, zero value otherwise.

### GetHeaderMatchesOk

`func (o *RoutingTriggerOptions) GetHeaderMatchesOk() (*map[string]StringRegexMap, bool)`

GetHeaderMatchesOk returns a tuple with the HeaderMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderMatches

`func (o *RoutingTriggerOptions) SetHeaderMatches(v map[string]StringRegexMap)`

SetHeaderMatches sets HeaderMatches field to given value.

### HasHeaderMatches

`func (o *RoutingTriggerOptions) HasHeaderMatches() bool`

HasHeaderMatches returns a boolean if a field has been set.

### GetPathPartMatches

`func (o *RoutingTriggerOptions) GetPathPartMatches() map[string]StringRegexMap`

GetPathPartMatches returns the PathPartMatches field if non-nil, zero value otherwise.

### GetPathPartMatchesOk

`func (o *RoutingTriggerOptions) GetPathPartMatchesOk() (*map[string]StringRegexMap, bool)`

GetPathPartMatchesOk returns a tuple with the PathPartMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPartMatches

`func (o *RoutingTriggerOptions) SetPathPartMatches(v map[string]StringRegexMap)`

SetPathPartMatches sets PathPartMatches field to given value.

### HasPathPartMatches

`func (o *RoutingTriggerOptions) HasPathPartMatches() bool`

HasPathPartMatches returns a boolean if a field has been set.

### GetPayloadMatches

`func (o *RoutingTriggerOptions) GetPayloadMatches() StringRegexMap`

GetPayloadMatches returns the PayloadMatches field if non-nil, zero value otherwise.

### GetPayloadMatchesOk

`func (o *RoutingTriggerOptions) GetPayloadMatchesOk() (*StringRegexMap, bool)`

GetPayloadMatchesOk returns a tuple with the PayloadMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadMatches

`func (o *RoutingTriggerOptions) SetPayloadMatches(v StringRegexMap)`

SetPayloadMatches sets PayloadMatches field to given value.

### HasPayloadMatches

`func (o *RoutingTriggerOptions) HasPayloadMatches() bool`

HasPayloadMatches returns a boolean if a field has been set.

### GetQueryValMatches

`func (o *RoutingTriggerOptions) GetQueryValMatches() map[string]StringRegexMap`

GetQueryValMatches returns the QueryValMatches field if non-nil, zero value otherwise.

### GetQueryValMatchesOk

`func (o *RoutingTriggerOptions) GetQueryValMatchesOk() (*map[string]StringRegexMap, bool)`

GetQueryValMatchesOk returns a tuple with the QueryValMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryValMatches

`func (o *RoutingTriggerOptions) SetQueryValMatches(v map[string]StringRegexMap)`

SetQueryValMatches sets QueryValMatches field to given value.

### HasQueryValMatches

`func (o *RoutingTriggerOptions) HasQueryValMatches() bool`

HasQueryValMatches returns a boolean if a field has been set.

### GetRequestContextMatches

`func (o *RoutingTriggerOptions) GetRequestContextMatches() map[string]StringRegexMap`

GetRequestContextMatches returns the RequestContextMatches field if non-nil, zero value otherwise.

### GetRequestContextMatchesOk

`func (o *RoutingTriggerOptions) GetRequestContextMatchesOk() (*map[string]StringRegexMap, bool)`

GetRequestContextMatchesOk returns a tuple with the RequestContextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContextMatches

`func (o *RoutingTriggerOptions) SetRequestContextMatches(v map[string]StringRegexMap)`

SetRequestContextMatches sets RequestContextMatches field to given value.

### HasRequestContextMatches

`func (o *RoutingTriggerOptions) HasRequestContextMatches() bool`

HasRequestContextMatches returns a boolean if a field has been set.

### GetSessionMetaMatches

`func (o *RoutingTriggerOptions) GetSessionMetaMatches() map[string]StringRegexMap`

GetSessionMetaMatches returns the SessionMetaMatches field if non-nil, zero value otherwise.

### GetSessionMetaMatchesOk

`func (o *RoutingTriggerOptions) GetSessionMetaMatchesOk() (*map[string]StringRegexMap, bool)`

GetSessionMetaMatchesOk returns a tuple with the SessionMetaMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMetaMatches

`func (o *RoutingTriggerOptions) SetSessionMetaMatches(v map[string]StringRegexMap)`

SetSessionMetaMatches sets SessionMetaMatches field to given value.

### HasSessionMetaMatches

`func (o *RoutingTriggerOptions) HasSessionMetaMatches() bool`

HasSessionMetaMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


