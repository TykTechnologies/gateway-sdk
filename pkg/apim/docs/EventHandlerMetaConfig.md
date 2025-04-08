# EventHandlerMetaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**map[string][]EventHandlerTriggerConfig**](array.md) |  | [optional] 

## Methods

### NewEventHandlerMetaConfig

`func NewEventHandlerMetaConfig() *EventHandlerMetaConfig`

NewEventHandlerMetaConfig instantiates a new EventHandlerMetaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHandlerMetaConfigWithDefaults

`func NewEventHandlerMetaConfigWithDefaults() *EventHandlerMetaConfig`

NewEventHandlerMetaConfigWithDefaults instantiates a new EventHandlerMetaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventHandlerMetaConfig) GetEvents() map[string][]EventHandlerTriggerConfig`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventHandlerMetaConfig) GetEventsOk() (*map[string][]EventHandlerTriggerConfig, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventHandlerMetaConfig) SetEvents(v map[string][]EventHandlerTriggerConfig)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventHandlerMetaConfig) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *EventHandlerMetaConfig) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *EventHandlerMetaConfig) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


