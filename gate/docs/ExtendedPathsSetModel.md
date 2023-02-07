# ExtendedPathsSetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvanceCacheConfig** | Pointer to [**[]CacheMetaModelModel**](CacheMetaModelModel.md) |  | [optional] 
**BlackList** | Pointer to [**[]EndPointMetaModelModel**](EndPointMetaModelModel.md) |  | [optional] 
**Cache** | Pointer to **[]string** |  | [optional] 
**CircuitBreakers** | Pointer to [**[]CircuitBreakerMetaModelModel**](CircuitBreakerMetaModelModel.md) |  | [optional] 
**DoNotTrackEndpoints** | Pointer to [**[]TrackEndpointMetaModelModel**](TrackEndpointMetaModelModel.md) |  | [optional] 
**HardTimeouts** | Pointer to [**[]HardTimeoutMetaModelModel**](HardTimeoutMetaModelModel.md) |  | [optional] 
**Ignored** | Pointer to [**[]EndPointMetaModelModel**](EndPointMetaModelModel.md) |  | [optional] 
**Internal** | Pointer to [**[]InternalMetaModelModel**](InternalMetaModelModel.md) |  | [optional] 
**MethodTransforms** | Pointer to [**[]MethodTransformMetaModelModel**](MethodTransformMetaModelModel.md) |  | [optional] 
**SizeLimits** | Pointer to [**[]RequestSizeMetaModelModel**](RequestSizeMetaModelModel.md) |  | [optional] 
**TrackEndpoints** | Pointer to [**[]TrackEndpointMetaModelModel**](TrackEndpointMetaModelModel.md) |  | [optional] 
**Transform** | Pointer to [**[]TemplateMetaModelModel**](TemplateMetaModelModel.md) |  | [optional] 
**TransformHeaders** | Pointer to [**[]HeaderInjectionMetaModelModel**](HeaderInjectionMetaModelModel.md) |  | [optional] 
**TransformJq** | Pointer to [**[]TransformJQMetaModelModel**](TransformJQMetaModelModel.md) |  | [optional] 
**TransformJqResponse** | Pointer to [**[]TransformJQMetaModelModel**](TransformJQMetaModelModel.md) |  | [optional] 
**TransformResponse** | Pointer to [**[]TemplateMetaModelModel**](TemplateMetaModelModel.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**[]HeaderInjectionMetaModelModel**](HeaderInjectionMetaModelModel.md) |  | [optional] 
**UrlRewrites** | Pointer to [**[]URLRewriteMetaModelModel**](URLRewriteMetaModelModel.md) |  | [optional] 
**ValidateJson** | Pointer to [**[]ValidatePathMetaModelModel**](ValidatePathMetaModelModel.md) |  | [optional] 
**Virtual** | Pointer to [**[]VirtualMetaModelModel**](VirtualMetaModelModel.md) |  | [optional] 
**WhiteList** | Pointer to [**[]EndPointMetaModelModel**](EndPointMetaModelModel.md) |  | [optional] 

## Methods

### NewExtendedPathsSetModel

`func NewExtendedPathsSetModel() *ExtendedPathsSetModel`

NewExtendedPathsSetModel instantiates a new ExtendedPathsSetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedPathsSetModelWithDefaults

`func NewExtendedPathsSetModelWithDefaults() *ExtendedPathsSetModel`

NewExtendedPathsSetModelWithDefaults instantiates a new ExtendedPathsSetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanceCacheConfig

`func (o *ExtendedPathsSetModel) GetAdvanceCacheConfig() []CacheMetaModelModel`

GetAdvanceCacheConfig returns the AdvanceCacheConfig field if non-nil, zero value otherwise.

### GetAdvanceCacheConfigOk

`func (o *ExtendedPathsSetModel) GetAdvanceCacheConfigOk() (*[]CacheMetaModelModel, bool)`

GetAdvanceCacheConfigOk returns a tuple with the AdvanceCacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceCacheConfig

`func (o *ExtendedPathsSetModel) SetAdvanceCacheConfig(v []CacheMetaModelModel)`

SetAdvanceCacheConfig sets AdvanceCacheConfig field to given value.

### HasAdvanceCacheConfig

`func (o *ExtendedPathsSetModel) HasAdvanceCacheConfig() bool`

HasAdvanceCacheConfig returns a boolean if a field has been set.

### GetBlackList

`func (o *ExtendedPathsSetModel) GetBlackList() []EndPointMetaModelModel`

GetBlackList returns the BlackList field if non-nil, zero value otherwise.

### GetBlackListOk

`func (o *ExtendedPathsSetModel) GetBlackListOk() (*[]EndPointMetaModelModel, bool)`

GetBlackListOk returns a tuple with the BlackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackList

`func (o *ExtendedPathsSetModel) SetBlackList(v []EndPointMetaModelModel)`

SetBlackList sets BlackList field to given value.

### HasBlackList

`func (o *ExtendedPathsSetModel) HasBlackList() bool`

HasBlackList returns a boolean if a field has been set.

### GetCache

`func (o *ExtendedPathsSetModel) GetCache() []string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ExtendedPathsSetModel) GetCacheOk() (*[]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ExtendedPathsSetModel) SetCache(v []string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ExtendedPathsSetModel) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCircuitBreakers

`func (o *ExtendedPathsSetModel) GetCircuitBreakers() []CircuitBreakerMetaModelModel`

GetCircuitBreakers returns the CircuitBreakers field if non-nil, zero value otherwise.

### GetCircuitBreakersOk

`func (o *ExtendedPathsSetModel) GetCircuitBreakersOk() (*[]CircuitBreakerMetaModelModel, bool)`

GetCircuitBreakersOk returns a tuple with the CircuitBreakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakers

`func (o *ExtendedPathsSetModel) SetCircuitBreakers(v []CircuitBreakerMetaModelModel)`

SetCircuitBreakers sets CircuitBreakers field to given value.

### HasCircuitBreakers

`func (o *ExtendedPathsSetModel) HasCircuitBreakers() bool`

HasCircuitBreakers returns a boolean if a field has been set.

### GetDoNotTrackEndpoints

`func (o *ExtendedPathsSetModel) GetDoNotTrackEndpoints() []TrackEndpointMetaModelModel`

GetDoNotTrackEndpoints returns the DoNotTrackEndpoints field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointsOk

`func (o *ExtendedPathsSetModel) GetDoNotTrackEndpointsOk() (*[]TrackEndpointMetaModelModel, bool)`

GetDoNotTrackEndpointsOk returns a tuple with the DoNotTrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoints

`func (o *ExtendedPathsSetModel) SetDoNotTrackEndpoints(v []TrackEndpointMetaModelModel)`

SetDoNotTrackEndpoints sets DoNotTrackEndpoints field to given value.

### HasDoNotTrackEndpoints

`func (o *ExtendedPathsSetModel) HasDoNotTrackEndpoints() bool`

HasDoNotTrackEndpoints returns a boolean if a field has been set.

### GetHardTimeouts

`func (o *ExtendedPathsSetModel) GetHardTimeouts() []HardTimeoutMetaModelModel`

GetHardTimeouts returns the HardTimeouts field if non-nil, zero value otherwise.

### GetHardTimeoutsOk

`func (o *ExtendedPathsSetModel) GetHardTimeoutsOk() (*[]HardTimeoutMetaModelModel, bool)`

GetHardTimeoutsOk returns a tuple with the HardTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardTimeouts

`func (o *ExtendedPathsSetModel) SetHardTimeouts(v []HardTimeoutMetaModelModel)`

SetHardTimeouts sets HardTimeouts field to given value.

### HasHardTimeouts

`func (o *ExtendedPathsSetModel) HasHardTimeouts() bool`

HasHardTimeouts returns a boolean if a field has been set.

### GetIgnored

`func (o *ExtendedPathsSetModel) GetIgnored() []EndPointMetaModelModel`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ExtendedPathsSetModel) GetIgnoredOk() (*[]EndPointMetaModelModel, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ExtendedPathsSetModel) SetIgnored(v []EndPointMetaModelModel)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ExtendedPathsSetModel) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInternal

`func (o *ExtendedPathsSetModel) GetInternal() []InternalMetaModelModel`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtendedPathsSetModel) GetInternalOk() (*[]InternalMetaModelModel, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtendedPathsSetModel) SetInternal(v []InternalMetaModelModel)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ExtendedPathsSetModel) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMethodTransforms

`func (o *ExtendedPathsSetModel) GetMethodTransforms() []MethodTransformMetaModelModel`

GetMethodTransforms returns the MethodTransforms field if non-nil, zero value otherwise.

### GetMethodTransformsOk

`func (o *ExtendedPathsSetModel) GetMethodTransformsOk() (*[]MethodTransformMetaModelModel, bool)`

GetMethodTransformsOk returns a tuple with the MethodTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodTransforms

`func (o *ExtendedPathsSetModel) SetMethodTransforms(v []MethodTransformMetaModelModel)`

SetMethodTransforms sets MethodTransforms field to given value.

### HasMethodTransforms

`func (o *ExtendedPathsSetModel) HasMethodTransforms() bool`

HasMethodTransforms returns a boolean if a field has been set.

### GetSizeLimits

`func (o *ExtendedPathsSetModel) GetSizeLimits() []RequestSizeMetaModelModel`

GetSizeLimits returns the SizeLimits field if non-nil, zero value otherwise.

### GetSizeLimitsOk

`func (o *ExtendedPathsSetModel) GetSizeLimitsOk() (*[]RequestSizeMetaModelModel, bool)`

GetSizeLimitsOk returns a tuple with the SizeLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimits

`func (o *ExtendedPathsSetModel) SetSizeLimits(v []RequestSizeMetaModelModel)`

SetSizeLimits sets SizeLimits field to given value.

### HasSizeLimits

`func (o *ExtendedPathsSetModel) HasSizeLimits() bool`

HasSizeLimits returns a boolean if a field has been set.

### GetTrackEndpoints

`func (o *ExtendedPathsSetModel) GetTrackEndpoints() []TrackEndpointMetaModelModel`

GetTrackEndpoints returns the TrackEndpoints field if non-nil, zero value otherwise.

### GetTrackEndpointsOk

`func (o *ExtendedPathsSetModel) GetTrackEndpointsOk() (*[]TrackEndpointMetaModelModel, bool)`

GetTrackEndpointsOk returns a tuple with the TrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoints

`func (o *ExtendedPathsSetModel) SetTrackEndpoints(v []TrackEndpointMetaModelModel)`

SetTrackEndpoints sets TrackEndpoints field to given value.

### HasTrackEndpoints

`func (o *ExtendedPathsSetModel) HasTrackEndpoints() bool`

HasTrackEndpoints returns a boolean if a field has been set.

### GetTransform

`func (o *ExtendedPathsSetModel) GetTransform() []TemplateMetaModelModel`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *ExtendedPathsSetModel) GetTransformOk() (*[]TemplateMetaModelModel, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *ExtendedPathsSetModel) SetTransform(v []TemplateMetaModelModel)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *ExtendedPathsSetModel) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### GetTransformHeaders

`func (o *ExtendedPathsSetModel) GetTransformHeaders() []HeaderInjectionMetaModelModel`

GetTransformHeaders returns the TransformHeaders field if non-nil, zero value otherwise.

### GetTransformHeadersOk

`func (o *ExtendedPathsSetModel) GetTransformHeadersOk() (*[]HeaderInjectionMetaModelModel, bool)`

GetTransformHeadersOk returns a tuple with the TransformHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformHeaders

`func (o *ExtendedPathsSetModel) SetTransformHeaders(v []HeaderInjectionMetaModelModel)`

SetTransformHeaders sets TransformHeaders field to given value.

### HasTransformHeaders

`func (o *ExtendedPathsSetModel) HasTransformHeaders() bool`

HasTransformHeaders returns a boolean if a field has been set.

### GetTransformJq

`func (o *ExtendedPathsSetModel) GetTransformJq() []TransformJQMetaModelModel`

GetTransformJq returns the TransformJq field if non-nil, zero value otherwise.

### GetTransformJqOk

`func (o *ExtendedPathsSetModel) GetTransformJqOk() (*[]TransformJQMetaModelModel, bool)`

GetTransformJqOk returns a tuple with the TransformJq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJq

`func (o *ExtendedPathsSetModel) SetTransformJq(v []TransformJQMetaModelModel)`

SetTransformJq sets TransformJq field to given value.

### HasTransformJq

`func (o *ExtendedPathsSetModel) HasTransformJq() bool`

HasTransformJq returns a boolean if a field has been set.

### GetTransformJqResponse

`func (o *ExtendedPathsSetModel) GetTransformJqResponse() []TransformJQMetaModelModel`

GetTransformJqResponse returns the TransformJqResponse field if non-nil, zero value otherwise.

### GetTransformJqResponseOk

`func (o *ExtendedPathsSetModel) GetTransformJqResponseOk() (*[]TransformJQMetaModelModel, bool)`

GetTransformJqResponseOk returns a tuple with the TransformJqResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJqResponse

`func (o *ExtendedPathsSetModel) SetTransformJqResponse(v []TransformJQMetaModelModel)`

SetTransformJqResponse sets TransformJqResponse field to given value.

### HasTransformJqResponse

`func (o *ExtendedPathsSetModel) HasTransformJqResponse() bool`

HasTransformJqResponse returns a boolean if a field has been set.

### GetTransformResponse

`func (o *ExtendedPathsSetModel) GetTransformResponse() []TemplateMetaModelModel`

GetTransformResponse returns the TransformResponse field if non-nil, zero value otherwise.

### GetTransformResponseOk

`func (o *ExtendedPathsSetModel) GetTransformResponseOk() (*[]TemplateMetaModelModel, bool)`

GetTransformResponseOk returns a tuple with the TransformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponse

`func (o *ExtendedPathsSetModel) SetTransformResponse(v []TemplateMetaModelModel)`

SetTransformResponse sets TransformResponse field to given value.

### HasTransformResponse

`func (o *ExtendedPathsSetModel) HasTransformResponse() bool`

HasTransformResponse returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *ExtendedPathsSetModel) GetTransformResponseHeaders() []HeaderInjectionMetaModelModel`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *ExtendedPathsSetModel) GetTransformResponseHeadersOk() (*[]HeaderInjectionMetaModelModel, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *ExtendedPathsSetModel) SetTransformResponseHeaders(v []HeaderInjectionMetaModelModel)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *ExtendedPathsSetModel) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### GetUrlRewrites

`func (o *ExtendedPathsSetModel) GetUrlRewrites() []URLRewriteMetaModelModel`

GetUrlRewrites returns the UrlRewrites field if non-nil, zero value otherwise.

### GetUrlRewritesOk

`func (o *ExtendedPathsSetModel) GetUrlRewritesOk() (*[]URLRewriteMetaModelModel, bool)`

GetUrlRewritesOk returns a tuple with the UrlRewrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrites

`func (o *ExtendedPathsSetModel) SetUrlRewrites(v []URLRewriteMetaModelModel)`

SetUrlRewrites sets UrlRewrites field to given value.

### HasUrlRewrites

`func (o *ExtendedPathsSetModel) HasUrlRewrites() bool`

HasUrlRewrites returns a boolean if a field has been set.

### GetValidateJson

`func (o *ExtendedPathsSetModel) GetValidateJson() []ValidatePathMetaModelModel`

GetValidateJson returns the ValidateJson field if non-nil, zero value otherwise.

### GetValidateJsonOk

`func (o *ExtendedPathsSetModel) GetValidateJsonOk() (*[]ValidatePathMetaModelModel, bool)`

GetValidateJsonOk returns a tuple with the ValidateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateJson

`func (o *ExtendedPathsSetModel) SetValidateJson(v []ValidatePathMetaModelModel)`

SetValidateJson sets ValidateJson field to given value.

### HasValidateJson

`func (o *ExtendedPathsSetModel) HasValidateJson() bool`

HasValidateJson returns a boolean if a field has been set.

### GetVirtual

`func (o *ExtendedPathsSetModel) GetVirtual() []VirtualMetaModelModel`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *ExtendedPathsSetModel) GetVirtualOk() (*[]VirtualMetaModelModel, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *ExtendedPathsSetModel) SetVirtual(v []VirtualMetaModelModel)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *ExtendedPathsSetModel) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### GetWhiteList

`func (o *ExtendedPathsSetModel) GetWhiteList() []EndPointMetaModelModel`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *ExtendedPathsSetModel) GetWhiteListOk() (*[]EndPointMetaModelModel, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *ExtendedPathsSetModel) SetWhiteList(v []EndPointMetaModelModel)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *ExtendedPathsSetModel) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


