# RoutingTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**RoutingTriggerOptions**](RoutingTriggerOptions.md) |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingTrigger

`func NewRoutingTrigger() *RoutingTrigger`

NewRoutingTrigger instantiates a new RoutingTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingTriggerWithDefaults

`func NewRoutingTriggerWithDefaults() *RoutingTrigger`

NewRoutingTriggerWithDefaults instantiates a new RoutingTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *RoutingTrigger) GetOn() string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *RoutingTrigger) GetOnOk() (*string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *RoutingTrigger) SetOn(v string)`

SetOn sets On field to given value.

### HasOn

`func (o *RoutingTrigger) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetOptions

`func (o *RoutingTrigger) GetOptions() RoutingTriggerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RoutingTrigger) GetOptionsOk() (*RoutingTriggerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RoutingTrigger) SetOptions(v RoutingTriggerOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RoutingTrigger) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRewriteTo

`func (o *RoutingTrigger) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *RoutingTrigger) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *RoutingTrigger) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *RoutingTrigger) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


