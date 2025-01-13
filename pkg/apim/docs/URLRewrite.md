# URLRewrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]URLRewriteTrigger**](URLRewriteTrigger.md) |  | [optional] 

## Methods

### NewURLRewrite

`func NewURLRewrite() *URLRewrite`

NewURLRewrite instantiates a new URLRewrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLRewriteWithDefaults

`func NewURLRewriteWithDefaults() *URLRewrite`

NewURLRewriteWithDefaults instantiates a new URLRewrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *URLRewrite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *URLRewrite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *URLRewrite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *URLRewrite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPattern

`func (o *URLRewrite) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *URLRewrite) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *URLRewrite) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *URLRewrite) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRewriteTo

`func (o *URLRewrite) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *URLRewrite) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *URLRewrite) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *URLRewrite) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetTriggers

`func (o *URLRewrite) GetTriggers() []URLRewriteTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *URLRewrite) GetTriggersOk() (*[]URLRewriteTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *URLRewrite) SetTriggers(v []URLRewriteTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *URLRewrite) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


