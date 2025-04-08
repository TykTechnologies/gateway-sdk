# URLRewriteTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]URLRewriteRule**](URLRewriteRule.md) |  | [optional] 

## Methods

### NewURLRewriteTrigger

`func NewURLRewriteTrigger() *URLRewriteTrigger`

NewURLRewriteTrigger instantiates a new URLRewriteTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLRewriteTriggerWithDefaults

`func NewURLRewriteTriggerWithDefaults() *URLRewriteTrigger`

NewURLRewriteTriggerWithDefaults instantiates a new URLRewriteTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *URLRewriteTrigger) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *URLRewriteTrigger) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *URLRewriteTrigger) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *URLRewriteTrigger) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetRewriteTo

`func (o *URLRewriteTrigger) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *URLRewriteTrigger) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *URLRewriteTrigger) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *URLRewriteTrigger) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetRules

`func (o *URLRewriteTrigger) GetRules() []URLRewriteRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *URLRewriteTrigger) GetRulesOk() (*[]URLRewriteRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *URLRewriteTrigger) SetRules(v []URLRewriteRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *URLRewriteTrigger) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


