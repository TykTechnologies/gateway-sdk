# URLRewriteMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**MatchPattern** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]RoutingTrigger**](RoutingTrigger.md) |  | [optional] 

## Methods

### NewURLRewriteMeta

`func NewURLRewriteMeta() *URLRewriteMeta`

NewURLRewriteMeta instantiates a new URLRewriteMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLRewriteMetaWithDefaults

`func NewURLRewriteMetaWithDefaults() *URLRewriteMeta`

NewURLRewriteMetaWithDefaults instantiates a new URLRewriteMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *URLRewriteMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *URLRewriteMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *URLRewriteMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *URLRewriteMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMatchPattern

`func (o *URLRewriteMeta) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *URLRewriteMeta) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *URLRewriteMeta) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.

### HasMatchPattern

`func (o *URLRewriteMeta) HasMatchPattern() bool`

HasMatchPattern returns a boolean if a field has been set.

### GetMethod

`func (o *URLRewriteMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *URLRewriteMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *URLRewriteMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *URLRewriteMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *URLRewriteMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *URLRewriteMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *URLRewriteMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *URLRewriteMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRewriteTo

`func (o *URLRewriteMeta) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *URLRewriteMeta) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *URLRewriteMeta) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *URLRewriteMeta) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetTriggers

`func (o *URLRewriteMeta) GetTriggers() []RoutingTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *URLRewriteMeta) GetTriggersOk() (*[]RoutingTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *URLRewriteMeta) SetTriggers(v []RoutingTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *URLRewriteMeta) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *URLRewriteMeta) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *URLRewriteMeta) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


