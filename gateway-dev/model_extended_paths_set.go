/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gate

import (
	"encoding/json"
)

// ExtendedPathsSet struct for ExtendedPathsSet
type ExtendedPathsSet struct {
	AdvanceCacheConfig []CacheMeta `json:"advance_cache_config,omitempty"`
	BlackList []EndPointMeta `json:"black_list,omitempty"`
	Cache []string `json:"cache,omitempty"`
	CircuitBreakers []CircuitBreakerMeta `json:"circuit_breakers,omitempty"`
	DoNotTrackEndpoints []TrackEndpointMeta `json:"do_not_track_endpoints,omitempty"`
	HardTimeouts []HardTimeoutMeta `json:"hard_timeouts,omitempty"`
	Ignored []EndPointMeta `json:"ignored,omitempty"`
	Internal []InternalMeta `json:"internal,omitempty"`
	MethodTransforms []MethodTransformMeta `json:"method_transforms,omitempty"`
	SizeLimits []RequestSizeMeta `json:"size_limits,omitempty"`
	TrackEndpoints []TrackEndpointMeta `json:"track_endpoints,omitempty"`
	Transform []TemplateMeta `json:"transform,omitempty"`
	TransformHeaders []HeaderInjectionMeta `json:"transform_headers,omitempty"`
	TransformJq []TransformJQMeta `json:"transform_jq,omitempty"`
	TransformJqResponse []TransformJQMeta `json:"transform_jq_response,omitempty"`
	TransformResponse []TemplateMeta `json:"transform_response,omitempty"`
	TransformResponseHeaders []HeaderInjectionMeta `json:"transform_response_headers,omitempty"`
	UrlRewrites []URLRewriteMeta `json:"url_rewrites,omitempty"`
	ValidateJson []ValidatePathMeta `json:"validate_json,omitempty"`
	Virtual []VirtualMeta `json:"virtual,omitempty"`
	WhiteList []EndPointMeta `json:"white_list,omitempty"`
}

// NewExtendedPathsSet instantiates a new ExtendedPathsSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedPathsSet() *ExtendedPathsSet {
	this := ExtendedPathsSet{}
	return &this
}

// NewExtendedPathsSetWithDefaults instantiates a new ExtendedPathsSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedPathsSetWithDefaults() *ExtendedPathsSet {
	this := ExtendedPathsSet{}
	return &this
}

// GetAdvanceCacheConfig returns the AdvanceCacheConfig field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetAdvanceCacheConfig() []CacheMeta {
	if o == nil || o.AdvanceCacheConfig == nil {
		var ret []CacheMeta
		return ret
	}
	return o.AdvanceCacheConfig
}

// GetAdvanceCacheConfigOk returns a tuple with the AdvanceCacheConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetAdvanceCacheConfigOk() ([]CacheMeta, bool) {
	if o == nil || o.AdvanceCacheConfig == nil {
		return nil, false
	}
	return o.AdvanceCacheConfig, true
}

// HasAdvanceCacheConfig returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasAdvanceCacheConfig() bool {
	if o != nil && o.AdvanceCacheConfig != nil {
		return true
	}

	return false
}

// SetAdvanceCacheConfig gets a reference to the given []CacheMeta and assigns it to the AdvanceCacheConfig field.
func (o *ExtendedPathsSet) SetAdvanceCacheConfig(v []CacheMeta) {
	o.AdvanceCacheConfig = v
}

// GetBlackList returns the BlackList field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetBlackList() []EndPointMeta {
	if o == nil || o.BlackList == nil {
		var ret []EndPointMeta
		return ret
	}
	return o.BlackList
}

// GetBlackListOk returns a tuple with the BlackList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetBlackListOk() ([]EndPointMeta, bool) {
	if o == nil || o.BlackList == nil {
		return nil, false
	}
	return o.BlackList, true
}

// HasBlackList returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasBlackList() bool {
	if o != nil && o.BlackList != nil {
		return true
	}

	return false
}

// SetBlackList gets a reference to the given []EndPointMeta and assigns it to the BlackList field.
func (o *ExtendedPathsSet) SetBlackList(v []EndPointMeta) {
	o.BlackList = v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetCache() []string {
	if o == nil || o.Cache == nil {
		var ret []string
		return ret
	}
	return o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetCacheOk() ([]string, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given []string and assigns it to the Cache field.
func (o *ExtendedPathsSet) SetCache(v []string) {
	o.Cache = v
}

// GetCircuitBreakers returns the CircuitBreakers field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetCircuitBreakers() []CircuitBreakerMeta {
	if o == nil || o.CircuitBreakers == nil {
		var ret []CircuitBreakerMeta
		return ret
	}
	return o.CircuitBreakers
}

// GetCircuitBreakersOk returns a tuple with the CircuitBreakers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetCircuitBreakersOk() ([]CircuitBreakerMeta, bool) {
	if o == nil || o.CircuitBreakers == nil {
		return nil, false
	}
	return o.CircuitBreakers, true
}

// HasCircuitBreakers returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasCircuitBreakers() bool {
	if o != nil && o.CircuitBreakers != nil {
		return true
	}

	return false
}

// SetCircuitBreakers gets a reference to the given []CircuitBreakerMeta and assigns it to the CircuitBreakers field.
func (o *ExtendedPathsSet) SetCircuitBreakers(v []CircuitBreakerMeta) {
	o.CircuitBreakers = v
}

// GetDoNotTrackEndpoints returns the DoNotTrackEndpoints field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetDoNotTrackEndpoints() []TrackEndpointMeta {
	if o == nil || o.DoNotTrackEndpoints == nil {
		var ret []TrackEndpointMeta
		return ret
	}
	return o.DoNotTrackEndpoints
}

// GetDoNotTrackEndpointsOk returns a tuple with the DoNotTrackEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetDoNotTrackEndpointsOk() ([]TrackEndpointMeta, bool) {
	if o == nil || o.DoNotTrackEndpoints == nil {
		return nil, false
	}
	return o.DoNotTrackEndpoints, true
}

// HasDoNotTrackEndpoints returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasDoNotTrackEndpoints() bool {
	if o != nil && o.DoNotTrackEndpoints != nil {
		return true
	}

	return false
}

// SetDoNotTrackEndpoints gets a reference to the given []TrackEndpointMeta and assigns it to the DoNotTrackEndpoints field.
func (o *ExtendedPathsSet) SetDoNotTrackEndpoints(v []TrackEndpointMeta) {
	o.DoNotTrackEndpoints = v
}

// GetHardTimeouts returns the HardTimeouts field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetHardTimeouts() []HardTimeoutMeta {
	if o == nil || o.HardTimeouts == nil {
		var ret []HardTimeoutMeta
		return ret
	}
	return o.HardTimeouts
}

// GetHardTimeoutsOk returns a tuple with the HardTimeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetHardTimeoutsOk() ([]HardTimeoutMeta, bool) {
	if o == nil || o.HardTimeouts == nil {
		return nil, false
	}
	return o.HardTimeouts, true
}

// HasHardTimeouts returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasHardTimeouts() bool {
	if o != nil && o.HardTimeouts != nil {
		return true
	}

	return false
}

// SetHardTimeouts gets a reference to the given []HardTimeoutMeta and assigns it to the HardTimeouts field.
func (o *ExtendedPathsSet) SetHardTimeouts(v []HardTimeoutMeta) {
	o.HardTimeouts = v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetIgnored() []EndPointMeta {
	if o == nil || o.Ignored == nil {
		var ret []EndPointMeta
		return ret
	}
	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetIgnoredOk() ([]EndPointMeta, bool) {
	if o == nil || o.Ignored == nil {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasIgnored() bool {
	if o != nil && o.Ignored != nil {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given []EndPointMeta and assigns it to the Ignored field.
func (o *ExtendedPathsSet) SetIgnored(v []EndPointMeta) {
	o.Ignored = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetInternal() []InternalMeta {
	if o == nil || o.Internal == nil {
		var ret []InternalMeta
		return ret
	}
	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetInternalOk() ([]InternalMeta, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given []InternalMeta and assigns it to the Internal field.
func (o *ExtendedPathsSet) SetInternal(v []InternalMeta) {
	o.Internal = v
}

// GetMethodTransforms returns the MethodTransforms field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetMethodTransforms() []MethodTransformMeta {
	if o == nil || o.MethodTransforms == nil {
		var ret []MethodTransformMeta
		return ret
	}
	return o.MethodTransforms
}

// GetMethodTransformsOk returns a tuple with the MethodTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetMethodTransformsOk() ([]MethodTransformMeta, bool) {
	if o == nil || o.MethodTransforms == nil {
		return nil, false
	}
	return o.MethodTransforms, true
}

// HasMethodTransforms returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasMethodTransforms() bool {
	if o != nil && o.MethodTransforms != nil {
		return true
	}

	return false
}

// SetMethodTransforms gets a reference to the given []MethodTransformMeta and assigns it to the MethodTransforms field.
func (o *ExtendedPathsSet) SetMethodTransforms(v []MethodTransformMeta) {
	o.MethodTransforms = v
}

// GetSizeLimits returns the SizeLimits field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetSizeLimits() []RequestSizeMeta {
	if o == nil || o.SizeLimits == nil {
		var ret []RequestSizeMeta
		return ret
	}
	return o.SizeLimits
}

// GetSizeLimitsOk returns a tuple with the SizeLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetSizeLimitsOk() ([]RequestSizeMeta, bool) {
	if o == nil || o.SizeLimits == nil {
		return nil, false
	}
	return o.SizeLimits, true
}

// HasSizeLimits returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasSizeLimits() bool {
	if o != nil && o.SizeLimits != nil {
		return true
	}

	return false
}

// SetSizeLimits gets a reference to the given []RequestSizeMeta and assigns it to the SizeLimits field.
func (o *ExtendedPathsSet) SetSizeLimits(v []RequestSizeMeta) {
	o.SizeLimits = v
}

// GetTrackEndpoints returns the TrackEndpoints field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTrackEndpoints() []TrackEndpointMeta {
	if o == nil || o.TrackEndpoints == nil {
		var ret []TrackEndpointMeta
		return ret
	}
	return o.TrackEndpoints
}

// GetTrackEndpointsOk returns a tuple with the TrackEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTrackEndpointsOk() ([]TrackEndpointMeta, bool) {
	if o == nil || o.TrackEndpoints == nil {
		return nil, false
	}
	return o.TrackEndpoints, true
}

// HasTrackEndpoints returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTrackEndpoints() bool {
	if o != nil && o.TrackEndpoints != nil {
		return true
	}

	return false
}

// SetTrackEndpoints gets a reference to the given []TrackEndpointMeta and assigns it to the TrackEndpoints field.
func (o *ExtendedPathsSet) SetTrackEndpoints(v []TrackEndpointMeta) {
	o.TrackEndpoints = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransform() []TemplateMeta {
	if o == nil || o.Transform == nil {
		var ret []TemplateMeta
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformOk() ([]TemplateMeta, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given []TemplateMeta and assigns it to the Transform field.
func (o *ExtendedPathsSet) SetTransform(v []TemplateMeta) {
	o.Transform = v
}

// GetTransformHeaders returns the TransformHeaders field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransformHeaders() []HeaderInjectionMeta {
	if o == nil || o.TransformHeaders == nil {
		var ret []HeaderInjectionMeta
		return ret
	}
	return o.TransformHeaders
}

// GetTransformHeadersOk returns a tuple with the TransformHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformHeadersOk() ([]HeaderInjectionMeta, bool) {
	if o == nil || o.TransformHeaders == nil {
		return nil, false
	}
	return o.TransformHeaders, true
}

// HasTransformHeaders returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransformHeaders() bool {
	if o != nil && o.TransformHeaders != nil {
		return true
	}

	return false
}

// SetTransformHeaders gets a reference to the given []HeaderInjectionMeta and assigns it to the TransformHeaders field.
func (o *ExtendedPathsSet) SetTransformHeaders(v []HeaderInjectionMeta) {
	o.TransformHeaders = v
}

// GetTransformJq returns the TransformJq field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransformJq() []TransformJQMeta {
	if o == nil || o.TransformJq == nil {
		var ret []TransformJQMeta
		return ret
	}
	return o.TransformJq
}

// GetTransformJqOk returns a tuple with the TransformJq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformJqOk() ([]TransformJQMeta, bool) {
	if o == nil || o.TransformJq == nil {
		return nil, false
	}
	return o.TransformJq, true
}

// HasTransformJq returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransformJq() bool {
	if o != nil && o.TransformJq != nil {
		return true
	}

	return false
}

// SetTransformJq gets a reference to the given []TransformJQMeta and assigns it to the TransformJq field.
func (o *ExtendedPathsSet) SetTransformJq(v []TransformJQMeta) {
	o.TransformJq = v
}

// GetTransformJqResponse returns the TransformJqResponse field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransformJqResponse() []TransformJQMeta {
	if o == nil || o.TransformJqResponse == nil {
		var ret []TransformJQMeta
		return ret
	}
	return o.TransformJqResponse
}

// GetTransformJqResponseOk returns a tuple with the TransformJqResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformJqResponseOk() ([]TransformJQMeta, bool) {
	if o == nil || o.TransformJqResponse == nil {
		return nil, false
	}
	return o.TransformJqResponse, true
}

// HasTransformJqResponse returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransformJqResponse() bool {
	if o != nil && o.TransformJqResponse != nil {
		return true
	}

	return false
}

// SetTransformJqResponse gets a reference to the given []TransformJQMeta and assigns it to the TransformJqResponse field.
func (o *ExtendedPathsSet) SetTransformJqResponse(v []TransformJQMeta) {
	o.TransformJqResponse = v
}

// GetTransformResponse returns the TransformResponse field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransformResponse() []TemplateMeta {
	if o == nil || o.TransformResponse == nil {
		var ret []TemplateMeta
		return ret
	}
	return o.TransformResponse
}

// GetTransformResponseOk returns a tuple with the TransformResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformResponseOk() ([]TemplateMeta, bool) {
	if o == nil || o.TransformResponse == nil {
		return nil, false
	}
	return o.TransformResponse, true
}

// HasTransformResponse returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransformResponse() bool {
	if o != nil && o.TransformResponse != nil {
		return true
	}

	return false
}

// SetTransformResponse gets a reference to the given []TemplateMeta and assigns it to the TransformResponse field.
func (o *ExtendedPathsSet) SetTransformResponse(v []TemplateMeta) {
	o.TransformResponse = v
}

// GetTransformResponseHeaders returns the TransformResponseHeaders field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetTransformResponseHeaders() []HeaderInjectionMeta {
	if o == nil || o.TransformResponseHeaders == nil {
		var ret []HeaderInjectionMeta
		return ret
	}
	return o.TransformResponseHeaders
}

// GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetTransformResponseHeadersOk() ([]HeaderInjectionMeta, bool) {
	if o == nil || o.TransformResponseHeaders == nil {
		return nil, false
	}
	return o.TransformResponseHeaders, true
}

// HasTransformResponseHeaders returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasTransformResponseHeaders() bool {
	if o != nil && o.TransformResponseHeaders != nil {
		return true
	}

	return false
}

// SetTransformResponseHeaders gets a reference to the given []HeaderInjectionMeta and assigns it to the TransformResponseHeaders field.
func (o *ExtendedPathsSet) SetTransformResponseHeaders(v []HeaderInjectionMeta) {
	o.TransformResponseHeaders = v
}

// GetUrlRewrites returns the UrlRewrites field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetUrlRewrites() []URLRewriteMeta {
	if o == nil || o.UrlRewrites == nil {
		var ret []URLRewriteMeta
		return ret
	}
	return o.UrlRewrites
}

// GetUrlRewritesOk returns a tuple with the UrlRewrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetUrlRewritesOk() ([]URLRewriteMeta, bool) {
	if o == nil || o.UrlRewrites == nil {
		return nil, false
	}
	return o.UrlRewrites, true
}

// HasUrlRewrites returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasUrlRewrites() bool {
	if o != nil && o.UrlRewrites != nil {
		return true
	}

	return false
}

// SetUrlRewrites gets a reference to the given []URLRewriteMeta and assigns it to the UrlRewrites field.
func (o *ExtendedPathsSet) SetUrlRewrites(v []URLRewriteMeta) {
	o.UrlRewrites = v
}

// GetValidateJson returns the ValidateJson field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetValidateJson() []ValidatePathMeta {
	if o == nil || o.ValidateJson == nil {
		var ret []ValidatePathMeta
		return ret
	}
	return o.ValidateJson
}

// GetValidateJsonOk returns a tuple with the ValidateJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetValidateJsonOk() ([]ValidatePathMeta, bool) {
	if o == nil || o.ValidateJson == nil {
		return nil, false
	}
	return o.ValidateJson, true
}

// HasValidateJson returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasValidateJson() bool {
	if o != nil && o.ValidateJson != nil {
		return true
	}

	return false
}

// SetValidateJson gets a reference to the given []ValidatePathMeta and assigns it to the ValidateJson field.
func (o *ExtendedPathsSet) SetValidateJson(v []ValidatePathMeta) {
	o.ValidateJson = v
}

// GetVirtual returns the Virtual field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetVirtual() []VirtualMeta {
	if o == nil || o.Virtual == nil {
		var ret []VirtualMeta
		return ret
	}
	return o.Virtual
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetVirtualOk() ([]VirtualMeta, bool) {
	if o == nil || o.Virtual == nil {
		return nil, false
	}
	return o.Virtual, true
}

// HasVirtual returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasVirtual() bool {
	if o != nil && o.Virtual != nil {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given []VirtualMeta and assigns it to the Virtual field.
func (o *ExtendedPathsSet) SetVirtual(v []VirtualMeta) {
	o.Virtual = v
}

// GetWhiteList returns the WhiteList field value if set, zero value otherwise.
func (o *ExtendedPathsSet) GetWhiteList() []EndPointMeta {
	if o == nil || o.WhiteList == nil {
		var ret []EndPointMeta
		return ret
	}
	return o.WhiteList
}

// GetWhiteListOk returns a tuple with the WhiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPathsSet) GetWhiteListOk() ([]EndPointMeta, bool) {
	if o == nil || o.WhiteList == nil {
		return nil, false
	}
	return o.WhiteList, true
}

// HasWhiteList returns a boolean if a field has been set.
func (o *ExtendedPathsSet) HasWhiteList() bool {
	if o != nil && o.WhiteList != nil {
		return true
	}

	return false
}

// SetWhiteList gets a reference to the given []EndPointMeta and assigns it to the WhiteList field.
func (o *ExtendedPathsSet) SetWhiteList(v []EndPointMeta) {
	o.WhiteList = v
}

func (o ExtendedPathsSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdvanceCacheConfig != nil {
		toSerialize["advance_cache_config"] = o.AdvanceCacheConfig
	}
	if o.BlackList != nil {
		toSerialize["black_list"] = o.BlackList
	}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.CircuitBreakers != nil {
		toSerialize["circuit_breakers"] = o.CircuitBreakers
	}
	if o.DoNotTrackEndpoints != nil {
		toSerialize["do_not_track_endpoints"] = o.DoNotTrackEndpoints
	}
	if o.HardTimeouts != nil {
		toSerialize["hard_timeouts"] = o.HardTimeouts
	}
	if o.Ignored != nil {
		toSerialize["ignored"] = o.Ignored
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.MethodTransforms != nil {
		toSerialize["method_transforms"] = o.MethodTransforms
	}
	if o.SizeLimits != nil {
		toSerialize["size_limits"] = o.SizeLimits
	}
	if o.TrackEndpoints != nil {
		toSerialize["track_endpoints"] = o.TrackEndpoints
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	if o.TransformHeaders != nil {
		toSerialize["transform_headers"] = o.TransformHeaders
	}
	if o.TransformJq != nil {
		toSerialize["transform_jq"] = o.TransformJq
	}
	if o.TransformJqResponse != nil {
		toSerialize["transform_jq_response"] = o.TransformJqResponse
	}
	if o.TransformResponse != nil {
		toSerialize["transform_response"] = o.TransformResponse
	}
	if o.TransformResponseHeaders != nil {
		toSerialize["transform_response_headers"] = o.TransformResponseHeaders
	}
	if o.UrlRewrites != nil {
		toSerialize["url_rewrites"] = o.UrlRewrites
	}
	if o.ValidateJson != nil {
		toSerialize["validate_json"] = o.ValidateJson
	}
	if o.Virtual != nil {
		toSerialize["virtual"] = o.Virtual
	}
	if o.WhiteList != nil {
		toSerialize["white_list"] = o.WhiteList
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedPathsSet struct {
	value *ExtendedPathsSet
	isSet bool
}

func (v NullableExtendedPathsSet) Get() *ExtendedPathsSet {
	return v.value
}

func (v *NullableExtendedPathsSet) Set(val *ExtendedPathsSet) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedPathsSet) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedPathsSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedPathsSet(val *ExtendedPathsSet) *NullableExtendedPathsSet {
	return &NullableExtendedPathsSet{value: val, isSet: true}
}

func (v NullableExtendedPathsSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedPathsSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


