# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to [**Allowance**](Allowance.md) |  | [optional] 
**Block** | Pointer to [**Allowance**](Allowance.md) |  | [optional] 
**Cache** | Pointer to [**CachePlugin**](CachePlugin.md) |  | [optional] 
**CircuitBreaker** | Pointer to [**CircuitBreaker**](CircuitBreaker.md) |  | [optional] 
**DoNotTrackEndpoint** | Pointer to [**TrackEndpoint**](TrackEndpoint.md) |  | [optional] 
**EnforceTimeout** | Pointer to [**EnforceTimeout**](EnforceTimeout.md) |  | [optional] 
**IgnoreAuthentication** | Pointer to [**Allowance**](Allowance.md) |  | [optional] 
**Internal** | Pointer to [**Internal**](Internal.md) |  | [optional] 
**MockResponse** | Pointer to [**MockResponse**](MockResponse.md) |  | [optional] 
**PostPlugins** | Pointer to [**[]EndpointPostPlugin**](EndpointPostPlugin.md) |  | [optional] 
**RateLimit** | Pointer to [**RateLimitEndpoint**](RateLimitEndpoint.md) |  | [optional] 
**RequestSizeLimit** | Pointer to [**RequestSizeLimit**](RequestSizeLimit.md) |  | [optional] 
**TrackEndpoint** | Pointer to [**TrackEndpoint**](TrackEndpoint.md) |  | [optional] 
**TransformRequestBody** | Pointer to [**TransformBody**](TransformBody.md) |  | [optional] 
**TransformRequestHeaders** | Pointer to [**TransformHeaders**](TransformHeaders.md) |  | [optional] 
**TransformRequestMethod** | Pointer to [**TransformRequestMethod**](TransformRequestMethod.md) |  | [optional] 
**TransformResponseBody** | Pointer to [**TransformBody**](TransformBody.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**TransformHeaders**](TransformHeaders.md) |  | [optional] 
**UrlRewrite** | Pointer to [**URLRewrite**](URLRewrite.md) |  | [optional] 
**ValidateRequest** | Pointer to [**ValidateRequest**](ValidateRequest.md) |  | [optional] 
**VirtualEndpoint** | Pointer to [**VirtualEndpoint**](VirtualEndpoint.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *Operation) GetAllow() Allowance`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *Operation) GetAllowOk() (*Allowance, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *Operation) SetAllow(v Allowance)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *Operation) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetBlock

`func (o *Operation) GetBlock() Allowance`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Operation) GetBlockOk() (*Allowance, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Operation) SetBlock(v Allowance)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *Operation) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetCache

`func (o *Operation) GetCache() CachePlugin`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *Operation) GetCacheOk() (*CachePlugin, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *Operation) SetCache(v CachePlugin)`

SetCache sets Cache field to given value.

### HasCache

`func (o *Operation) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCircuitBreaker

`func (o *Operation) GetCircuitBreaker() CircuitBreaker`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *Operation) GetCircuitBreakerOk() (*CircuitBreaker, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *Operation) SetCircuitBreaker(v CircuitBreaker)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *Operation) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetDoNotTrackEndpoint

`func (o *Operation) GetDoNotTrackEndpoint() TrackEndpoint`

GetDoNotTrackEndpoint returns the DoNotTrackEndpoint field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointOk

`func (o *Operation) GetDoNotTrackEndpointOk() (*TrackEndpoint, bool)`

GetDoNotTrackEndpointOk returns a tuple with the DoNotTrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoint

`func (o *Operation) SetDoNotTrackEndpoint(v TrackEndpoint)`

SetDoNotTrackEndpoint sets DoNotTrackEndpoint field to given value.

### HasDoNotTrackEndpoint

`func (o *Operation) HasDoNotTrackEndpoint() bool`

HasDoNotTrackEndpoint returns a boolean if a field has been set.

### GetEnforceTimeout

`func (o *Operation) GetEnforceTimeout() EnforceTimeout`

GetEnforceTimeout returns the EnforceTimeout field if non-nil, zero value otherwise.

### GetEnforceTimeoutOk

`func (o *Operation) GetEnforceTimeoutOk() (*EnforceTimeout, bool)`

GetEnforceTimeoutOk returns a tuple with the EnforceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTimeout

`func (o *Operation) SetEnforceTimeout(v EnforceTimeout)`

SetEnforceTimeout sets EnforceTimeout field to given value.

### HasEnforceTimeout

`func (o *Operation) HasEnforceTimeout() bool`

HasEnforceTimeout returns a boolean if a field has been set.

### GetIgnoreAuthentication

`func (o *Operation) GetIgnoreAuthentication() Allowance`

GetIgnoreAuthentication returns the IgnoreAuthentication field if non-nil, zero value otherwise.

### GetIgnoreAuthenticationOk

`func (o *Operation) GetIgnoreAuthenticationOk() (*Allowance, bool)`

GetIgnoreAuthenticationOk returns a tuple with the IgnoreAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAuthentication

`func (o *Operation) SetIgnoreAuthentication(v Allowance)`

SetIgnoreAuthentication sets IgnoreAuthentication field to given value.

### HasIgnoreAuthentication

`func (o *Operation) HasIgnoreAuthentication() bool`

HasIgnoreAuthentication returns a boolean if a field has been set.

### GetInternal

`func (o *Operation) GetInternal() Internal`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Operation) GetInternalOk() (*Internal, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Operation) SetInternal(v Internal)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Operation) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMockResponse

`func (o *Operation) GetMockResponse() MockResponse`

GetMockResponse returns the MockResponse field if non-nil, zero value otherwise.

### GetMockResponseOk

`func (o *Operation) GetMockResponseOk() (*MockResponse, bool)`

GetMockResponseOk returns a tuple with the MockResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockResponse

`func (o *Operation) SetMockResponse(v MockResponse)`

SetMockResponse sets MockResponse field to given value.

### HasMockResponse

`func (o *Operation) HasMockResponse() bool`

HasMockResponse returns a boolean if a field has been set.

### GetPostPlugins

`func (o *Operation) GetPostPlugins() []EndpointPostPlugin`

GetPostPlugins returns the PostPlugins field if non-nil, zero value otherwise.

### GetPostPluginsOk

`func (o *Operation) GetPostPluginsOk() (*[]EndpointPostPlugin, bool)`

GetPostPluginsOk returns a tuple with the PostPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugins

`func (o *Operation) SetPostPlugins(v []EndpointPostPlugin)`

SetPostPlugins sets PostPlugins field to given value.

### HasPostPlugins

`func (o *Operation) HasPostPlugins() bool`

HasPostPlugins returns a boolean if a field has been set.

### GetRateLimit

`func (o *Operation) GetRateLimit() RateLimitEndpoint`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *Operation) GetRateLimitOk() (*RateLimitEndpoint, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *Operation) SetRateLimit(v RateLimitEndpoint)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *Operation) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRequestSizeLimit

`func (o *Operation) GetRequestSizeLimit() RequestSizeLimit`

GetRequestSizeLimit returns the RequestSizeLimit field if non-nil, zero value otherwise.

### GetRequestSizeLimitOk

`func (o *Operation) GetRequestSizeLimitOk() (*RequestSizeLimit, bool)`

GetRequestSizeLimitOk returns a tuple with the RequestSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSizeLimit

`func (o *Operation) SetRequestSizeLimit(v RequestSizeLimit)`

SetRequestSizeLimit sets RequestSizeLimit field to given value.

### HasRequestSizeLimit

`func (o *Operation) HasRequestSizeLimit() bool`

HasRequestSizeLimit returns a boolean if a field has been set.

### GetTrackEndpoint

`func (o *Operation) GetTrackEndpoint() TrackEndpoint`

GetTrackEndpoint returns the TrackEndpoint field if non-nil, zero value otherwise.

### GetTrackEndpointOk

`func (o *Operation) GetTrackEndpointOk() (*TrackEndpoint, bool)`

GetTrackEndpointOk returns a tuple with the TrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoint

`func (o *Operation) SetTrackEndpoint(v TrackEndpoint)`

SetTrackEndpoint sets TrackEndpoint field to given value.

### HasTrackEndpoint

`func (o *Operation) HasTrackEndpoint() bool`

HasTrackEndpoint returns a boolean if a field has been set.

### GetTransformRequestBody

`func (o *Operation) GetTransformRequestBody() TransformBody`

GetTransformRequestBody returns the TransformRequestBody field if non-nil, zero value otherwise.

### GetTransformRequestBodyOk

`func (o *Operation) GetTransformRequestBodyOk() (*TransformBody, bool)`

GetTransformRequestBodyOk returns a tuple with the TransformRequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestBody

`func (o *Operation) SetTransformRequestBody(v TransformBody)`

SetTransformRequestBody sets TransformRequestBody field to given value.

### HasTransformRequestBody

`func (o *Operation) HasTransformRequestBody() bool`

HasTransformRequestBody returns a boolean if a field has been set.

### GetTransformRequestHeaders

`func (o *Operation) GetTransformRequestHeaders() TransformHeaders`

GetTransformRequestHeaders returns the TransformRequestHeaders field if non-nil, zero value otherwise.

### GetTransformRequestHeadersOk

`func (o *Operation) GetTransformRequestHeadersOk() (*TransformHeaders, bool)`

GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestHeaders

`func (o *Operation) SetTransformRequestHeaders(v TransformHeaders)`

SetTransformRequestHeaders sets TransformRequestHeaders field to given value.

### HasTransformRequestHeaders

`func (o *Operation) HasTransformRequestHeaders() bool`

HasTransformRequestHeaders returns a boolean if a field has been set.

### GetTransformRequestMethod

`func (o *Operation) GetTransformRequestMethod() TransformRequestMethod`

GetTransformRequestMethod returns the TransformRequestMethod field if non-nil, zero value otherwise.

### GetTransformRequestMethodOk

`func (o *Operation) GetTransformRequestMethodOk() (*TransformRequestMethod, bool)`

GetTransformRequestMethodOk returns a tuple with the TransformRequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestMethod

`func (o *Operation) SetTransformRequestMethod(v TransformRequestMethod)`

SetTransformRequestMethod sets TransformRequestMethod field to given value.

### HasTransformRequestMethod

`func (o *Operation) HasTransformRequestMethod() bool`

HasTransformRequestMethod returns a boolean if a field has been set.

### GetTransformResponseBody

`func (o *Operation) GetTransformResponseBody() TransformBody`

GetTransformResponseBody returns the TransformResponseBody field if non-nil, zero value otherwise.

### GetTransformResponseBodyOk

`func (o *Operation) GetTransformResponseBodyOk() (*TransformBody, bool)`

GetTransformResponseBodyOk returns a tuple with the TransformResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseBody

`func (o *Operation) SetTransformResponseBody(v TransformBody)`

SetTransformResponseBody sets TransformResponseBody field to given value.

### HasTransformResponseBody

`func (o *Operation) HasTransformResponseBody() bool`

HasTransformResponseBody returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *Operation) GetTransformResponseHeaders() TransformHeaders`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *Operation) GetTransformResponseHeadersOk() (*TransformHeaders, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *Operation) SetTransformResponseHeaders(v TransformHeaders)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *Operation) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### GetUrlRewrite

`func (o *Operation) GetUrlRewrite() URLRewrite`

GetUrlRewrite returns the UrlRewrite field if non-nil, zero value otherwise.

### GetUrlRewriteOk

`func (o *Operation) GetUrlRewriteOk() (*URLRewrite, bool)`

GetUrlRewriteOk returns a tuple with the UrlRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrite

`func (o *Operation) SetUrlRewrite(v URLRewrite)`

SetUrlRewrite sets UrlRewrite field to given value.

### HasUrlRewrite

`func (o *Operation) HasUrlRewrite() bool`

HasUrlRewrite returns a boolean if a field has been set.

### GetValidateRequest

`func (o *Operation) GetValidateRequest() ValidateRequest`

GetValidateRequest returns the ValidateRequest field if non-nil, zero value otherwise.

### GetValidateRequestOk

`func (o *Operation) GetValidateRequestOk() (*ValidateRequest, bool)`

GetValidateRequestOk returns a tuple with the ValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRequest

`func (o *Operation) SetValidateRequest(v ValidateRequest)`

SetValidateRequest sets ValidateRequest field to given value.

### HasValidateRequest

`func (o *Operation) HasValidateRequest() bool`

HasValidateRequest returns a boolean if a field has been set.

### GetVirtualEndpoint

`func (o *Operation) GetVirtualEndpoint() VirtualEndpoint`

GetVirtualEndpoint returns the VirtualEndpoint field if non-nil, zero value otherwise.

### GetVirtualEndpointOk

`func (o *Operation) GetVirtualEndpointOk() (*VirtualEndpoint, bool)`

GetVirtualEndpointOk returns a tuple with the VirtualEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEndpoint

`func (o *Operation) SetVirtualEndpoint(v VirtualEndpoint)`

SetVirtualEndpoint sets VirtualEndpoint field to given value.

### HasVirtualEndpoint

`func (o *Operation) HasVirtualEndpoint() bool`

HasVirtualEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


