# RequestHeadersRewriteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remove** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestHeadersRewriteConfig

`func NewRequestHeadersRewriteConfig() *RequestHeadersRewriteConfig`

NewRequestHeadersRewriteConfig instantiates a new RequestHeadersRewriteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestHeadersRewriteConfigWithDefaults

`func NewRequestHeadersRewriteConfigWithDefaults() *RequestHeadersRewriteConfig`

NewRequestHeadersRewriteConfigWithDefaults instantiates a new RequestHeadersRewriteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemove

`func (o *RequestHeadersRewriteConfig) GetRemove() bool`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *RequestHeadersRewriteConfig) GetRemoveOk() (*bool, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *RequestHeadersRewriteConfig) SetRemove(v bool)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *RequestHeadersRewriteConfig) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetValue

`func (o *RequestHeadersRewriteConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestHeadersRewriteConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestHeadersRewriteConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RequestHeadersRewriteConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


