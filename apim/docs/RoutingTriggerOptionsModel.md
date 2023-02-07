# RoutingTriggerOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderMatches** | Pointer to [**map[string]StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 
**PathPartMatches** | Pointer to [**map[string]StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 
**PayloadMatches** | Pointer to [**StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 
**QueryValMatches** | Pointer to [**map[string]StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 
**RequestContextMatches** | Pointer to [**map[string]StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 
**SessionMetaMatches** | Pointer to [**map[string]StringRegexMapModelModel**](StringRegexMapModel.md) |  | [optional] 

## Methods

### NewRoutingTriggerOptionsModel

`func NewRoutingTriggerOptionsModel() *RoutingTriggerOptionsModel`

NewRoutingTriggerOptionsModel instantiates a new RoutingTriggerOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingTriggerOptionsModelWithDefaults

`func NewRoutingTriggerOptionsModelWithDefaults() *RoutingTriggerOptionsModel`

NewRoutingTriggerOptionsModelWithDefaults instantiates a new RoutingTriggerOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderMatches

`func (o *RoutingTriggerOptionsModel) GetHeaderMatches() map[string]StringRegexMapModelModel`

GetHeaderMatches returns the HeaderMatches field if non-nil, zero value otherwise.

### GetHeaderMatchesOk

`func (o *RoutingTriggerOptionsModel) GetHeaderMatchesOk() (*map[string]StringRegexMapModelModel, bool)`

GetHeaderMatchesOk returns a tuple with the HeaderMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderMatches

`func (o *RoutingTriggerOptionsModel) SetHeaderMatches(v map[string]StringRegexMapModelModel)`

SetHeaderMatches sets HeaderMatches field to given value.

### HasHeaderMatches

`func (o *RoutingTriggerOptionsModel) HasHeaderMatches() bool`

HasHeaderMatches returns a boolean if a field has been set.

### GetPathPartMatches

`func (o *RoutingTriggerOptionsModel) GetPathPartMatches() map[string]StringRegexMapModelModel`

GetPathPartMatches returns the PathPartMatches field if non-nil, zero value otherwise.

### GetPathPartMatchesOk

`func (o *RoutingTriggerOptionsModel) GetPathPartMatchesOk() (*map[string]StringRegexMapModelModel, bool)`

GetPathPartMatchesOk returns a tuple with the PathPartMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPartMatches

`func (o *RoutingTriggerOptionsModel) SetPathPartMatches(v map[string]StringRegexMapModelModel)`

SetPathPartMatches sets PathPartMatches field to given value.

### HasPathPartMatches

`func (o *RoutingTriggerOptionsModel) HasPathPartMatches() bool`

HasPathPartMatches returns a boolean if a field has been set.

### GetPayloadMatches

`func (o *RoutingTriggerOptionsModel) GetPayloadMatches() StringRegexMapModelModel`

GetPayloadMatches returns the PayloadMatches field if non-nil, zero value otherwise.

### GetPayloadMatchesOk

`func (o *RoutingTriggerOptionsModel) GetPayloadMatchesOk() (*StringRegexMapModelModel, bool)`

GetPayloadMatchesOk returns a tuple with the PayloadMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadMatches

`func (o *RoutingTriggerOptionsModel) SetPayloadMatches(v StringRegexMapModelModel)`

SetPayloadMatches sets PayloadMatches field to given value.

### HasPayloadMatches

`func (o *RoutingTriggerOptionsModel) HasPayloadMatches() bool`

HasPayloadMatches returns a boolean if a field has been set.

### GetQueryValMatches

`func (o *RoutingTriggerOptionsModel) GetQueryValMatches() map[string]StringRegexMapModelModel`

GetQueryValMatches returns the QueryValMatches field if non-nil, zero value otherwise.

### GetQueryValMatchesOk

`func (o *RoutingTriggerOptionsModel) GetQueryValMatchesOk() (*map[string]StringRegexMapModelModel, bool)`

GetQueryValMatchesOk returns a tuple with the QueryValMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryValMatches

`func (o *RoutingTriggerOptionsModel) SetQueryValMatches(v map[string]StringRegexMapModelModel)`

SetQueryValMatches sets QueryValMatches field to given value.

### HasQueryValMatches

`func (o *RoutingTriggerOptionsModel) HasQueryValMatches() bool`

HasQueryValMatches returns a boolean if a field has been set.

### GetRequestContextMatches

`func (o *RoutingTriggerOptionsModel) GetRequestContextMatches() map[string]StringRegexMapModelModel`

GetRequestContextMatches returns the RequestContextMatches field if non-nil, zero value otherwise.

### GetRequestContextMatchesOk

`func (o *RoutingTriggerOptionsModel) GetRequestContextMatchesOk() (*map[string]StringRegexMapModelModel, bool)`

GetRequestContextMatchesOk returns a tuple with the RequestContextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContextMatches

`func (o *RoutingTriggerOptionsModel) SetRequestContextMatches(v map[string]StringRegexMapModelModel)`

SetRequestContextMatches sets RequestContextMatches field to given value.

### HasRequestContextMatches

`func (o *RoutingTriggerOptionsModel) HasRequestContextMatches() bool`

HasRequestContextMatches returns a boolean if a field has been set.

### GetSessionMetaMatches

`func (o *RoutingTriggerOptionsModel) GetSessionMetaMatches() map[string]StringRegexMapModelModel`

GetSessionMetaMatches returns the SessionMetaMatches field if non-nil, zero value otherwise.

### GetSessionMetaMatchesOk

`func (o *RoutingTriggerOptionsModel) GetSessionMetaMatchesOk() (*map[string]StringRegexMapModelModel, bool)`

GetSessionMetaMatchesOk returns a tuple with the SessionMetaMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMetaMatches

`func (o *RoutingTriggerOptionsModel) SetSessionMetaMatches(v map[string]StringRegexMapModelModel)`

SetSessionMetaMatches sets SessionMetaMatches field to given value.

### HasSessionMetaMatches

`func (o *RoutingTriggerOptionsModel) HasSessionMetaMatches() bool`

HasSessionMetaMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


