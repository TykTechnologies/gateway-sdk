# RoutingTriggerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**RoutingTriggerOptionsModelModel**](RoutingTriggerOptionsModel.md) |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingTriggerModel

`func NewRoutingTriggerModel() *RoutingTriggerModel`

NewRoutingTriggerModel instantiates a new RoutingTriggerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingTriggerModelWithDefaults

`func NewRoutingTriggerModelWithDefaults() *RoutingTriggerModel`

NewRoutingTriggerModelWithDefaults instantiates a new RoutingTriggerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *RoutingTriggerModel) GetOn() string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *RoutingTriggerModel) GetOnOk() (*string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *RoutingTriggerModel) SetOn(v string)`

SetOn sets On field to given value.

### HasOn

`func (o *RoutingTriggerModel) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetOptions

`func (o *RoutingTriggerModel) GetOptions() RoutingTriggerOptionsModelModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RoutingTriggerModel) GetOptionsOk() (*RoutingTriggerOptionsModelModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RoutingTriggerModel) SetOptions(v RoutingTriggerOptionsModelModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RoutingTriggerModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRewriteTo

`func (o *RoutingTriggerModel) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *RoutingTriggerModel) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *RoutingTriggerModel) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *RoutingTriggerModel) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


