# URLRewriteRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Negate** | Pointer to **bool** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 

## Methods

### NewURLRewriteRule

`func NewURLRewriteRule() *URLRewriteRule`

NewURLRewriteRule instantiates a new URLRewriteRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLRewriteRuleWithDefaults

`func NewURLRewriteRuleWithDefaults() *URLRewriteRule`

NewURLRewriteRuleWithDefaults instantiates a new URLRewriteRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *URLRewriteRule) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *URLRewriteRule) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *URLRewriteRule) SetIn(v string)`

SetIn sets In field to given value.

### HasIn

`func (o *URLRewriteRule) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetName

`func (o *URLRewriteRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *URLRewriteRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *URLRewriteRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *URLRewriteRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNegate

`func (o *URLRewriteRule) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *URLRewriteRule) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *URLRewriteRule) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *URLRewriteRule) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetPattern

`func (o *URLRewriteRule) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *URLRewriteRule) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *URLRewriteRule) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *URLRewriteRule) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


