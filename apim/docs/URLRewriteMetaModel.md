# URLRewriteMetaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchRegexp** | Pointer to [**RegexpModelModel**](RegexpModel.md) |  | [optional] 
**MatchPattern** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]RoutingTriggerModelModel**](RoutingTriggerModelModel.md) |  | [optional] 

## Methods

### NewURLRewriteMetaModel

`func NewURLRewriteMetaModel() *URLRewriteMetaModel`

NewURLRewriteMetaModel instantiates a new URLRewriteMetaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLRewriteMetaModelWithDefaults

`func NewURLRewriteMetaModelWithDefaults() *URLRewriteMetaModel`

NewURLRewriteMetaModelWithDefaults instantiates a new URLRewriteMetaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchRegexp

`func (o *URLRewriteMetaModel) GetMatchRegexp() RegexpModelModel`

GetMatchRegexp returns the MatchRegexp field if non-nil, zero value otherwise.

### GetMatchRegexpOk

`func (o *URLRewriteMetaModel) GetMatchRegexpOk() (*RegexpModelModel, bool)`

GetMatchRegexpOk returns a tuple with the MatchRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRegexp

`func (o *URLRewriteMetaModel) SetMatchRegexp(v RegexpModelModel)`

SetMatchRegexp sets MatchRegexp field to given value.

### HasMatchRegexp

`func (o *URLRewriteMetaModel) HasMatchRegexp() bool`

HasMatchRegexp returns a boolean if a field has been set.

### GetMatchPattern

`func (o *URLRewriteMetaModel) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *URLRewriteMetaModel) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *URLRewriteMetaModel) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.

### HasMatchPattern

`func (o *URLRewriteMetaModel) HasMatchPattern() bool`

HasMatchPattern returns a boolean if a field has been set.

### GetMethod

`func (o *URLRewriteMetaModel) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *URLRewriteMetaModel) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *URLRewriteMetaModel) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *URLRewriteMetaModel) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *URLRewriteMetaModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *URLRewriteMetaModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *URLRewriteMetaModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *URLRewriteMetaModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRewriteTo

`func (o *URLRewriteMetaModel) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *URLRewriteMetaModel) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *URLRewriteMetaModel) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *URLRewriteMetaModel) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetTriggers

`func (o *URLRewriteMetaModel) GetTriggers() []RoutingTriggerModelModel`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *URLRewriteMetaModel) GetTriggersOk() (*[]RoutingTriggerModelModel, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *URLRewriteMetaModel) SetTriggers(v []RoutingTriggerModelModel)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *URLRewriteMetaModel) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


