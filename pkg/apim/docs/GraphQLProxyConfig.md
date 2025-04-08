# GraphQLProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaders** | Pointer to **map[string]string** |  | [optional] 
**Features** | Pointer to [**GraphQLProxyFeaturesConfig**](GraphQLProxyFeaturesConfig.md) |  | [optional] 
**RequestHeaders** | Pointer to **map[string]string** |  | [optional] 
**RequestHeadersRewrite** | Pointer to [**map[string]RequestHeadersRewriteConfig**](RequestHeadersRewriteConfig.md) |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**UseResponseExtensions** | Pointer to [**GraphQLResponseExtensions**](GraphQLResponseExtensions.md) |  | [optional] 

## Methods

### NewGraphQLProxyConfig

`func NewGraphQLProxyConfig() *GraphQLProxyConfig`

NewGraphQLProxyConfig instantiates a new GraphQLProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLProxyConfigWithDefaults

`func NewGraphQLProxyConfigWithDefaults() *GraphQLProxyConfig`

NewGraphQLProxyConfigWithDefaults instantiates a new GraphQLProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaders

`func (o *GraphQLProxyConfig) GetAuthHeaders() map[string]string`

GetAuthHeaders returns the AuthHeaders field if non-nil, zero value otherwise.

### GetAuthHeadersOk

`func (o *GraphQLProxyConfig) GetAuthHeadersOk() (*map[string]string, bool)`

GetAuthHeadersOk returns a tuple with the AuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaders

`func (o *GraphQLProxyConfig) SetAuthHeaders(v map[string]string)`

SetAuthHeaders sets AuthHeaders field to given value.

### HasAuthHeaders

`func (o *GraphQLProxyConfig) HasAuthHeaders() bool`

HasAuthHeaders returns a boolean if a field has been set.

### SetAuthHeadersNil

`func (o *GraphQLProxyConfig) SetAuthHeadersNil(b bool)`

 SetAuthHeadersNil sets the value for AuthHeaders to be an explicit nil

### UnsetAuthHeaders
`func (o *GraphQLProxyConfig) UnsetAuthHeaders()`

UnsetAuthHeaders ensures that no value is present for AuthHeaders, not even an explicit nil
### GetFeatures

`func (o *GraphQLProxyConfig) GetFeatures() GraphQLProxyFeaturesConfig`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GraphQLProxyConfig) GetFeaturesOk() (*GraphQLProxyFeaturesConfig, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GraphQLProxyConfig) SetFeatures(v GraphQLProxyFeaturesConfig)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GraphQLProxyConfig) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *GraphQLProxyConfig) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *GraphQLProxyConfig) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *GraphQLProxyConfig) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *GraphQLProxyConfig) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *GraphQLProxyConfig) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *GraphQLProxyConfig) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestHeadersRewrite

`func (o *GraphQLProxyConfig) GetRequestHeadersRewrite() map[string]RequestHeadersRewriteConfig`

GetRequestHeadersRewrite returns the RequestHeadersRewrite field if non-nil, zero value otherwise.

### GetRequestHeadersRewriteOk

`func (o *GraphQLProxyConfig) GetRequestHeadersRewriteOk() (*map[string]RequestHeadersRewriteConfig, bool)`

GetRequestHeadersRewriteOk returns a tuple with the RequestHeadersRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeadersRewrite

`func (o *GraphQLProxyConfig) SetRequestHeadersRewrite(v map[string]RequestHeadersRewriteConfig)`

SetRequestHeadersRewrite sets RequestHeadersRewrite field to given value.

### HasRequestHeadersRewrite

`func (o *GraphQLProxyConfig) HasRequestHeadersRewrite() bool`

HasRequestHeadersRewrite returns a boolean if a field has been set.

### SetRequestHeadersRewriteNil

`func (o *GraphQLProxyConfig) SetRequestHeadersRewriteNil(b bool)`

 SetRequestHeadersRewriteNil sets the value for RequestHeadersRewrite to be an explicit nil

### UnsetRequestHeadersRewrite
`func (o *GraphQLProxyConfig) UnsetRequestHeadersRewrite()`

UnsetRequestHeadersRewrite ensures that no value is present for RequestHeadersRewrite, not even an explicit nil
### GetSubscriptionType

`func (o *GraphQLProxyConfig) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *GraphQLProxyConfig) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *GraphQLProxyConfig) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *GraphQLProxyConfig) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetUseResponseExtensions

`func (o *GraphQLProxyConfig) GetUseResponseExtensions() GraphQLResponseExtensions`

GetUseResponseExtensions returns the UseResponseExtensions field if non-nil, zero value otherwise.

### GetUseResponseExtensionsOk

`func (o *GraphQLProxyConfig) GetUseResponseExtensionsOk() (*GraphQLResponseExtensions, bool)`

GetUseResponseExtensionsOk returns a tuple with the UseResponseExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponseExtensions

`func (o *GraphQLProxyConfig) SetUseResponseExtensions(v GraphQLResponseExtensions)`

SetUseResponseExtensions sets UseResponseExtensions field to given value.

### HasUseResponseExtensions

`func (o *GraphQLProxyConfig) HasUseResponseExtensions() bool`

HasUseResponseExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


