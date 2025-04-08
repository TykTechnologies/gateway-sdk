# ExtendedPathsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvanceCacheConfig** | Pointer to [**[]CacheMeta**](CacheMeta.md) |  | [optional] 
**BlackList** | Pointer to [**[]EndPointMeta**](EndPointMeta.md) |  | [optional] 
**Cache** | Pointer to **[]string** |  | [optional] 
**CircuitBreakers** | Pointer to [**[]CircuitBreakerMeta**](CircuitBreakerMeta.md) |  | [optional] 
**DoNotTrackEndpoints** | Pointer to [**[]TrackEndpointMeta**](TrackEndpointMeta.md) |  | [optional] 
**GoPlugin** | Pointer to [**[]GoPluginMeta**](GoPluginMeta.md) |  | [optional] 
**HardTimeouts** | Pointer to [**[]HardTimeoutMeta**](HardTimeoutMeta.md) |  | [optional] 
**Ignored** | Pointer to [**[]EndPointMeta**](EndPointMeta.md) |  | [optional] 
**Internal** | Pointer to [**[]InternalMeta**](InternalMeta.md) |  | [optional] 
**MethodTransforms** | Pointer to [**[]MethodTransformMeta**](MethodTransformMeta.md) |  | [optional] 
**MockResponse** | Pointer to [**[]MockResponseMeta**](MockResponseMeta.md) |  | [optional] 
**PersistGraphql** | Pointer to [**[]PersistGraphQLMeta**](PersistGraphQLMeta.md) |  | [optional] 
**RateLimit** | Pointer to [**[]RateLimitMeta**](RateLimitMeta.md) |  | [optional] 
**SizeLimits** | Pointer to [**[]RequestSizeMeta**](RequestSizeMeta.md) |  | [optional] 
**TrackEndpoints** | Pointer to [**[]TrackEndpointMeta**](TrackEndpointMeta.md) |  | [optional] 
**Transform** | Pointer to [**[]TemplateMeta**](TemplateMeta.md) |  | [optional] 
**TransformHeaders** | Pointer to [**[]HeaderInjectionMeta**](HeaderInjectionMeta.md) |  | [optional] 
**TransformJq** | Pointer to [**[]TransformJQMeta**](TransformJQMeta.md) |  | [optional] 
**TransformJqResponse** | Pointer to [**[]TransformJQMeta**](TransformJQMeta.md) |  | [optional] 
**TransformResponse** | Pointer to [**[]TemplateMeta**](TemplateMeta.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**[]HeaderInjectionMeta**](HeaderInjectionMeta.md) |  | [optional] 
**UrlRewrites** | Pointer to [**[]URLRewriteMeta**](URLRewriteMeta.md) |  | [optional] 
**ValidateJson** | Pointer to [**[]ValidatePathMeta**](ValidatePathMeta.md) |  | [optional] 
**ValidateRequest** | Pointer to [**[]ValidateRequestMeta**](ValidateRequestMeta.md) |  | [optional] 
**Virtual** | Pointer to [**[]VirtualMeta**](VirtualMeta.md) |  | [optional] 
**WhiteList** | Pointer to [**[]EndPointMeta**](EndPointMeta.md) |  | [optional] 

## Methods

### NewExtendedPathsSet

`func NewExtendedPathsSet() *ExtendedPathsSet`

NewExtendedPathsSet instantiates a new ExtendedPathsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedPathsSetWithDefaults

`func NewExtendedPathsSetWithDefaults() *ExtendedPathsSet`

NewExtendedPathsSetWithDefaults instantiates a new ExtendedPathsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanceCacheConfig

`func (o *ExtendedPathsSet) GetAdvanceCacheConfig() []CacheMeta`

GetAdvanceCacheConfig returns the AdvanceCacheConfig field if non-nil, zero value otherwise.

### GetAdvanceCacheConfigOk

`func (o *ExtendedPathsSet) GetAdvanceCacheConfigOk() (*[]CacheMeta, bool)`

GetAdvanceCacheConfigOk returns a tuple with the AdvanceCacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceCacheConfig

`func (o *ExtendedPathsSet) SetAdvanceCacheConfig(v []CacheMeta)`

SetAdvanceCacheConfig sets AdvanceCacheConfig field to given value.

### HasAdvanceCacheConfig

`func (o *ExtendedPathsSet) HasAdvanceCacheConfig() bool`

HasAdvanceCacheConfig returns a boolean if a field has been set.

### GetBlackList

`func (o *ExtendedPathsSet) GetBlackList() []EndPointMeta`

GetBlackList returns the BlackList field if non-nil, zero value otherwise.

### GetBlackListOk

`func (o *ExtendedPathsSet) GetBlackListOk() (*[]EndPointMeta, bool)`

GetBlackListOk returns a tuple with the BlackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackList

`func (o *ExtendedPathsSet) SetBlackList(v []EndPointMeta)`

SetBlackList sets BlackList field to given value.

### HasBlackList

`func (o *ExtendedPathsSet) HasBlackList() bool`

HasBlackList returns a boolean if a field has been set.

### GetCache

`func (o *ExtendedPathsSet) GetCache() []string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ExtendedPathsSet) GetCacheOk() (*[]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ExtendedPathsSet) SetCache(v []string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ExtendedPathsSet) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCircuitBreakers

`func (o *ExtendedPathsSet) GetCircuitBreakers() []CircuitBreakerMeta`

GetCircuitBreakers returns the CircuitBreakers field if non-nil, zero value otherwise.

### GetCircuitBreakersOk

`func (o *ExtendedPathsSet) GetCircuitBreakersOk() (*[]CircuitBreakerMeta, bool)`

GetCircuitBreakersOk returns a tuple with the CircuitBreakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakers

`func (o *ExtendedPathsSet) SetCircuitBreakers(v []CircuitBreakerMeta)`

SetCircuitBreakers sets CircuitBreakers field to given value.

### HasCircuitBreakers

`func (o *ExtendedPathsSet) HasCircuitBreakers() bool`

HasCircuitBreakers returns a boolean if a field has been set.

### GetDoNotTrackEndpoints

`func (o *ExtendedPathsSet) GetDoNotTrackEndpoints() []TrackEndpointMeta`

GetDoNotTrackEndpoints returns the DoNotTrackEndpoints field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointsOk

`func (o *ExtendedPathsSet) GetDoNotTrackEndpointsOk() (*[]TrackEndpointMeta, bool)`

GetDoNotTrackEndpointsOk returns a tuple with the DoNotTrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoints

`func (o *ExtendedPathsSet) SetDoNotTrackEndpoints(v []TrackEndpointMeta)`

SetDoNotTrackEndpoints sets DoNotTrackEndpoints field to given value.

### HasDoNotTrackEndpoints

`func (o *ExtendedPathsSet) HasDoNotTrackEndpoints() bool`

HasDoNotTrackEndpoints returns a boolean if a field has been set.

### GetGoPlugin

`func (o *ExtendedPathsSet) GetGoPlugin() []GoPluginMeta`

GetGoPlugin returns the GoPlugin field if non-nil, zero value otherwise.

### GetGoPluginOk

`func (o *ExtendedPathsSet) GetGoPluginOk() (*[]GoPluginMeta, bool)`

GetGoPluginOk returns a tuple with the GoPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPlugin

`func (o *ExtendedPathsSet) SetGoPlugin(v []GoPluginMeta)`

SetGoPlugin sets GoPlugin field to given value.

### HasGoPlugin

`func (o *ExtendedPathsSet) HasGoPlugin() bool`

HasGoPlugin returns a boolean if a field has been set.

### GetHardTimeouts

`func (o *ExtendedPathsSet) GetHardTimeouts() []HardTimeoutMeta`

GetHardTimeouts returns the HardTimeouts field if non-nil, zero value otherwise.

### GetHardTimeoutsOk

`func (o *ExtendedPathsSet) GetHardTimeoutsOk() (*[]HardTimeoutMeta, bool)`

GetHardTimeoutsOk returns a tuple with the HardTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardTimeouts

`func (o *ExtendedPathsSet) SetHardTimeouts(v []HardTimeoutMeta)`

SetHardTimeouts sets HardTimeouts field to given value.

### HasHardTimeouts

`func (o *ExtendedPathsSet) HasHardTimeouts() bool`

HasHardTimeouts returns a boolean if a field has been set.

### GetIgnored

`func (o *ExtendedPathsSet) GetIgnored() []EndPointMeta`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ExtendedPathsSet) GetIgnoredOk() (*[]EndPointMeta, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ExtendedPathsSet) SetIgnored(v []EndPointMeta)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ExtendedPathsSet) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInternal

`func (o *ExtendedPathsSet) GetInternal() []InternalMeta`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtendedPathsSet) GetInternalOk() (*[]InternalMeta, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtendedPathsSet) SetInternal(v []InternalMeta)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ExtendedPathsSet) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMethodTransforms

`func (o *ExtendedPathsSet) GetMethodTransforms() []MethodTransformMeta`

GetMethodTransforms returns the MethodTransforms field if non-nil, zero value otherwise.

### GetMethodTransformsOk

`func (o *ExtendedPathsSet) GetMethodTransformsOk() (*[]MethodTransformMeta, bool)`

GetMethodTransformsOk returns a tuple with the MethodTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodTransforms

`func (o *ExtendedPathsSet) SetMethodTransforms(v []MethodTransformMeta)`

SetMethodTransforms sets MethodTransforms field to given value.

### HasMethodTransforms

`func (o *ExtendedPathsSet) HasMethodTransforms() bool`

HasMethodTransforms returns a boolean if a field has been set.

### GetMockResponse

`func (o *ExtendedPathsSet) GetMockResponse() []MockResponseMeta`

GetMockResponse returns the MockResponse field if non-nil, zero value otherwise.

### GetMockResponseOk

`func (o *ExtendedPathsSet) GetMockResponseOk() (*[]MockResponseMeta, bool)`

GetMockResponseOk returns a tuple with the MockResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockResponse

`func (o *ExtendedPathsSet) SetMockResponse(v []MockResponseMeta)`

SetMockResponse sets MockResponse field to given value.

### HasMockResponse

`func (o *ExtendedPathsSet) HasMockResponse() bool`

HasMockResponse returns a boolean if a field has been set.

### GetPersistGraphql

`func (o *ExtendedPathsSet) GetPersistGraphql() []PersistGraphQLMeta`

GetPersistGraphql returns the PersistGraphql field if non-nil, zero value otherwise.

### GetPersistGraphqlOk

`func (o *ExtendedPathsSet) GetPersistGraphqlOk() (*[]PersistGraphQLMeta, bool)`

GetPersistGraphqlOk returns a tuple with the PersistGraphql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistGraphql

`func (o *ExtendedPathsSet) SetPersistGraphql(v []PersistGraphQLMeta)`

SetPersistGraphql sets PersistGraphql field to given value.

### HasPersistGraphql

`func (o *ExtendedPathsSet) HasPersistGraphql() bool`

HasPersistGraphql returns a boolean if a field has been set.

### SetPersistGraphqlNil

`func (o *ExtendedPathsSet) SetPersistGraphqlNil(b bool)`

 SetPersistGraphqlNil sets the value for PersistGraphql to be an explicit nil

### UnsetPersistGraphql
`func (o *ExtendedPathsSet) UnsetPersistGraphql()`

UnsetPersistGraphql ensures that no value is present for PersistGraphql, not even an explicit nil
### GetRateLimit

`func (o *ExtendedPathsSet) GetRateLimit() []RateLimitMeta`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ExtendedPathsSet) GetRateLimitOk() (*[]RateLimitMeta, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ExtendedPathsSet) SetRateLimit(v []RateLimitMeta)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ExtendedPathsSet) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *ExtendedPathsSet) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *ExtendedPathsSet) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetSizeLimits

`func (o *ExtendedPathsSet) GetSizeLimits() []RequestSizeMeta`

GetSizeLimits returns the SizeLimits field if non-nil, zero value otherwise.

### GetSizeLimitsOk

`func (o *ExtendedPathsSet) GetSizeLimitsOk() (*[]RequestSizeMeta, bool)`

GetSizeLimitsOk returns a tuple with the SizeLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimits

`func (o *ExtendedPathsSet) SetSizeLimits(v []RequestSizeMeta)`

SetSizeLimits sets SizeLimits field to given value.

### HasSizeLimits

`func (o *ExtendedPathsSet) HasSizeLimits() bool`

HasSizeLimits returns a boolean if a field has been set.

### GetTrackEndpoints

`func (o *ExtendedPathsSet) GetTrackEndpoints() []TrackEndpointMeta`

GetTrackEndpoints returns the TrackEndpoints field if non-nil, zero value otherwise.

### GetTrackEndpointsOk

`func (o *ExtendedPathsSet) GetTrackEndpointsOk() (*[]TrackEndpointMeta, bool)`

GetTrackEndpointsOk returns a tuple with the TrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoints

`func (o *ExtendedPathsSet) SetTrackEndpoints(v []TrackEndpointMeta)`

SetTrackEndpoints sets TrackEndpoints field to given value.

### HasTrackEndpoints

`func (o *ExtendedPathsSet) HasTrackEndpoints() bool`

HasTrackEndpoints returns a boolean if a field has been set.

### GetTransform

`func (o *ExtendedPathsSet) GetTransform() []TemplateMeta`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *ExtendedPathsSet) GetTransformOk() (*[]TemplateMeta, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *ExtendedPathsSet) SetTransform(v []TemplateMeta)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *ExtendedPathsSet) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### GetTransformHeaders

`func (o *ExtendedPathsSet) GetTransformHeaders() []HeaderInjectionMeta`

GetTransformHeaders returns the TransformHeaders field if non-nil, zero value otherwise.

### GetTransformHeadersOk

`func (o *ExtendedPathsSet) GetTransformHeadersOk() (*[]HeaderInjectionMeta, bool)`

GetTransformHeadersOk returns a tuple with the TransformHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformHeaders

`func (o *ExtendedPathsSet) SetTransformHeaders(v []HeaderInjectionMeta)`

SetTransformHeaders sets TransformHeaders field to given value.

### HasTransformHeaders

`func (o *ExtendedPathsSet) HasTransformHeaders() bool`

HasTransformHeaders returns a boolean if a field has been set.

### GetTransformJq

`func (o *ExtendedPathsSet) GetTransformJq() []TransformJQMeta`

GetTransformJq returns the TransformJq field if non-nil, zero value otherwise.

### GetTransformJqOk

`func (o *ExtendedPathsSet) GetTransformJqOk() (*[]TransformJQMeta, bool)`

GetTransformJqOk returns a tuple with the TransformJq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJq

`func (o *ExtendedPathsSet) SetTransformJq(v []TransformJQMeta)`

SetTransformJq sets TransformJq field to given value.

### HasTransformJq

`func (o *ExtendedPathsSet) HasTransformJq() bool`

HasTransformJq returns a boolean if a field has been set.

### GetTransformJqResponse

`func (o *ExtendedPathsSet) GetTransformJqResponse() []TransformJQMeta`

GetTransformJqResponse returns the TransformJqResponse field if non-nil, zero value otherwise.

### GetTransformJqResponseOk

`func (o *ExtendedPathsSet) GetTransformJqResponseOk() (*[]TransformJQMeta, bool)`

GetTransformJqResponseOk returns a tuple with the TransformJqResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJqResponse

`func (o *ExtendedPathsSet) SetTransformJqResponse(v []TransformJQMeta)`

SetTransformJqResponse sets TransformJqResponse field to given value.

### HasTransformJqResponse

`func (o *ExtendedPathsSet) HasTransformJqResponse() bool`

HasTransformJqResponse returns a boolean if a field has been set.

### GetTransformResponse

`func (o *ExtendedPathsSet) GetTransformResponse() []TemplateMeta`

GetTransformResponse returns the TransformResponse field if non-nil, zero value otherwise.

### GetTransformResponseOk

`func (o *ExtendedPathsSet) GetTransformResponseOk() (*[]TemplateMeta, bool)`

GetTransformResponseOk returns a tuple with the TransformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponse

`func (o *ExtendedPathsSet) SetTransformResponse(v []TemplateMeta)`

SetTransformResponse sets TransformResponse field to given value.

### HasTransformResponse

`func (o *ExtendedPathsSet) HasTransformResponse() bool`

HasTransformResponse returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *ExtendedPathsSet) GetTransformResponseHeaders() []HeaderInjectionMeta`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *ExtendedPathsSet) GetTransformResponseHeadersOk() (*[]HeaderInjectionMeta, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *ExtendedPathsSet) SetTransformResponseHeaders(v []HeaderInjectionMeta)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *ExtendedPathsSet) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### GetUrlRewrites

`func (o *ExtendedPathsSet) GetUrlRewrites() []URLRewriteMeta`

GetUrlRewrites returns the UrlRewrites field if non-nil, zero value otherwise.

### GetUrlRewritesOk

`func (o *ExtendedPathsSet) GetUrlRewritesOk() (*[]URLRewriteMeta, bool)`

GetUrlRewritesOk returns a tuple with the UrlRewrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrites

`func (o *ExtendedPathsSet) SetUrlRewrites(v []URLRewriteMeta)`

SetUrlRewrites sets UrlRewrites field to given value.

### HasUrlRewrites

`func (o *ExtendedPathsSet) HasUrlRewrites() bool`

HasUrlRewrites returns a boolean if a field has been set.

### GetValidateJson

`func (o *ExtendedPathsSet) GetValidateJson() []ValidatePathMeta`

GetValidateJson returns the ValidateJson field if non-nil, zero value otherwise.

### GetValidateJsonOk

`func (o *ExtendedPathsSet) GetValidateJsonOk() (*[]ValidatePathMeta, bool)`

GetValidateJsonOk returns a tuple with the ValidateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateJson

`func (o *ExtendedPathsSet) SetValidateJson(v []ValidatePathMeta)`

SetValidateJson sets ValidateJson field to given value.

### HasValidateJson

`func (o *ExtendedPathsSet) HasValidateJson() bool`

HasValidateJson returns a boolean if a field has been set.

### GetValidateRequest

`func (o *ExtendedPathsSet) GetValidateRequest() []ValidateRequestMeta`

GetValidateRequest returns the ValidateRequest field if non-nil, zero value otherwise.

### GetValidateRequestOk

`func (o *ExtendedPathsSet) GetValidateRequestOk() (*[]ValidateRequestMeta, bool)`

GetValidateRequestOk returns a tuple with the ValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRequest

`func (o *ExtendedPathsSet) SetValidateRequest(v []ValidateRequestMeta)`

SetValidateRequest sets ValidateRequest field to given value.

### HasValidateRequest

`func (o *ExtendedPathsSet) HasValidateRequest() bool`

HasValidateRequest returns a boolean if a field has been set.

### GetVirtual

`func (o *ExtendedPathsSet) GetVirtual() []VirtualMeta`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *ExtendedPathsSet) GetVirtualOk() (*[]VirtualMeta, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *ExtendedPathsSet) SetVirtual(v []VirtualMeta)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *ExtendedPathsSet) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### GetWhiteList

`func (o *ExtendedPathsSet) GetWhiteList() []EndPointMeta`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *ExtendedPathsSet) GetWhiteListOk() (*[]EndPointMeta, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *ExtendedPathsSet) SetWhiteList(v []EndPointMeta)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *ExtendedPathsSet) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


