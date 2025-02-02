/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation). * Managing and listing policies. * Managing and listing API Definitions (only when not using the Tyk Dashboard). * Hot reloads / reloading a cluster configuration. * OAuth client creation (only when not using the Tyk Dashboard).  In order to use the Gateway API, you'll need to set the **secret** parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  **x-tyk-authorization: <your-secret>*** <br/>  <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the ServiceDiscovery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDiscovery{}

// ServiceDiscovery struct for ServiceDiscovery
type ServiceDiscovery struct {
	Cache               *ServiceDiscoveryCache `json:"cache,omitempty"`
	CacheTimeout        *int32                 `json:"cacheTimeout,omitempty"`
	DataPath            *string                `json:"dataPath,omitempty"`
	Enabled             *bool                  `json:"enabled,omitempty"`
	EndpointReturnsList *bool                  `json:"endpointReturnsList,omitempty"`
	ParentDataPath      *string                `json:"parentDataPath,omitempty"`
	PortDataPath        *string                `json:"portDataPath,omitempty"`
	QueryEndpoint       *string                `json:"queryEndpoint,omitempty"`
	TargetPath          *string                `json:"targetPath,omitempty"`
	UseNestedQuery      *bool                  `json:"useNestedQuery,omitempty"`
	UseTargetList       *bool                  `json:"useTargetList,omitempty"`
}

// NewServiceDiscovery instantiates a new ServiceDiscovery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDiscovery() *ServiceDiscovery {
	this := ServiceDiscovery{}
	return &this
}

// NewServiceDiscoveryWithDefaults instantiates a new ServiceDiscovery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDiscoveryWithDefaults() *ServiceDiscovery {
	this := ServiceDiscovery{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetCache() ServiceDiscoveryCache {
	if o == nil || IsNil(o.Cache) {
		var ret ServiceDiscoveryCache
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetCacheOk() (*ServiceDiscoveryCache, bool) {
	if o == nil || IsNil(o.Cache) {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasCache() bool {
	if o != nil && !IsNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given ServiceDiscoveryCache and assigns it to the Cache field.
func (o *ServiceDiscovery) SetCache(v ServiceDiscoveryCache) {
	o.Cache = &v
}

// GetCacheTimeout returns the CacheTimeout field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetCacheTimeout() int32 {
	if o == nil || IsNil(o.CacheTimeout) {
		var ret int32
		return ret
	}
	return *o.CacheTimeout
}

// GetCacheTimeoutOk returns a tuple with the CacheTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetCacheTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTimeout) {
		return nil, false
	}
	return o.CacheTimeout, true
}

// HasCacheTimeout returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasCacheTimeout() bool {
	if o != nil && !IsNil(o.CacheTimeout) {
		return true
	}

	return false
}

// SetCacheTimeout gets a reference to the given int32 and assigns it to the CacheTimeout field.
func (o *ServiceDiscovery) SetCacheTimeout(v int32) {
	o.CacheTimeout = &v
}

// GetDataPath returns the DataPath field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetDataPath() string {
	if o == nil || IsNil(o.DataPath) {
		var ret string
		return ret
	}
	return *o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.DataPath) {
		return nil, false
	}
	return o.DataPath, true
}

// HasDataPath returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasDataPath() bool {
	if o != nil && !IsNil(o.DataPath) {
		return true
	}

	return false
}

// SetDataPath gets a reference to the given string and assigns it to the DataPath field.
func (o *ServiceDiscovery) SetDataPath(v string) {
	o.DataPath = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceDiscovery) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEndpointReturnsList returns the EndpointReturnsList field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetEndpointReturnsList() bool {
	if o == nil || IsNil(o.EndpointReturnsList) {
		var ret bool
		return ret
	}
	return *o.EndpointReturnsList
}

// GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetEndpointReturnsListOk() (*bool, bool) {
	if o == nil || IsNil(o.EndpointReturnsList) {
		return nil, false
	}
	return o.EndpointReturnsList, true
}

// HasEndpointReturnsList returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasEndpointReturnsList() bool {
	if o != nil && !IsNil(o.EndpointReturnsList) {
		return true
	}

	return false
}

// SetEndpointReturnsList gets a reference to the given bool and assigns it to the EndpointReturnsList field.
func (o *ServiceDiscovery) SetEndpointReturnsList(v bool) {
	o.EndpointReturnsList = &v
}

// GetParentDataPath returns the ParentDataPath field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetParentDataPath() string {
	if o == nil || IsNil(o.ParentDataPath) {
		var ret string
		return ret
	}
	return *o.ParentDataPath
}

// GetParentDataPathOk returns a tuple with the ParentDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetParentDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDataPath) {
		return nil, false
	}
	return o.ParentDataPath, true
}

// HasParentDataPath returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasParentDataPath() bool {
	if o != nil && !IsNil(o.ParentDataPath) {
		return true
	}

	return false
}

// SetParentDataPath gets a reference to the given string and assigns it to the ParentDataPath field.
func (o *ServiceDiscovery) SetParentDataPath(v string) {
	o.ParentDataPath = &v
}

// GetPortDataPath returns the PortDataPath field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetPortDataPath() string {
	if o == nil || IsNil(o.PortDataPath) {
		var ret string
		return ret
	}
	return *o.PortDataPath
}

// GetPortDataPathOk returns a tuple with the PortDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetPortDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.PortDataPath) {
		return nil, false
	}
	return o.PortDataPath, true
}

// HasPortDataPath returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasPortDataPath() bool {
	if o != nil && !IsNil(o.PortDataPath) {
		return true
	}

	return false
}

// SetPortDataPath gets a reference to the given string and assigns it to the PortDataPath field.
func (o *ServiceDiscovery) SetPortDataPath(v string) {
	o.PortDataPath = &v
}

// GetQueryEndpoint returns the QueryEndpoint field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetQueryEndpoint() string {
	if o == nil || IsNil(o.QueryEndpoint) {
		var ret string
		return ret
	}
	return *o.QueryEndpoint
}

// GetQueryEndpointOk returns a tuple with the QueryEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetQueryEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.QueryEndpoint) {
		return nil, false
	}
	return o.QueryEndpoint, true
}

// HasQueryEndpoint returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasQueryEndpoint() bool {
	if o != nil && !IsNil(o.QueryEndpoint) {
		return true
	}

	return false
}

// SetQueryEndpoint gets a reference to the given string and assigns it to the QueryEndpoint field.
func (o *ServiceDiscovery) SetQueryEndpoint(v string) {
	o.QueryEndpoint = &v
}

// GetTargetPath returns the TargetPath field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetTargetPath() string {
	if o == nil || IsNil(o.TargetPath) {
		var ret string
		return ret
	}
	return *o.TargetPath
}

// GetTargetPathOk returns a tuple with the TargetPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetTargetPathOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPath) {
		return nil, false
	}
	return o.TargetPath, true
}

// HasTargetPath returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasTargetPath() bool {
	if o != nil && !IsNil(o.TargetPath) {
		return true
	}

	return false
}

// SetTargetPath gets a reference to the given string and assigns it to the TargetPath field.
func (o *ServiceDiscovery) SetTargetPath(v string) {
	o.TargetPath = &v
}

// GetUseNestedQuery returns the UseNestedQuery field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetUseNestedQuery() bool {
	if o == nil || IsNil(o.UseNestedQuery) {
		var ret bool
		return ret
	}
	return *o.UseNestedQuery
}

// GetUseNestedQueryOk returns a tuple with the UseNestedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetUseNestedQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNestedQuery) {
		return nil, false
	}
	return o.UseNestedQuery, true
}

// HasUseNestedQuery returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasUseNestedQuery() bool {
	if o != nil && !IsNil(o.UseNestedQuery) {
		return true
	}

	return false
}

// SetUseNestedQuery gets a reference to the given bool and assigns it to the UseNestedQuery field.
func (o *ServiceDiscovery) SetUseNestedQuery(v bool) {
	o.UseNestedQuery = &v
}

// GetUseTargetList returns the UseTargetList field value if set, zero value otherwise.
func (o *ServiceDiscovery) GetUseTargetList() bool {
	if o == nil || IsNil(o.UseTargetList) {
		var ret bool
		return ret
	}
	return *o.UseTargetList
}

// GetUseTargetListOk returns a tuple with the UseTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscovery) GetUseTargetListOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTargetList) {
		return nil, false
	}
	return o.UseTargetList, true
}

// HasUseTargetList returns a boolean if a field has been set.
func (o *ServiceDiscovery) HasUseTargetList() bool {
	if o != nil && !IsNil(o.UseTargetList) {
		return true
	}

	return false
}

// SetUseTargetList gets a reference to the given bool and assigns it to the UseTargetList field.
func (o *ServiceDiscovery) SetUseTargetList(v bool) {
	o.UseTargetList = &v
}

func (o ServiceDiscovery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDiscovery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cache) {
		toSerialize["cache"] = o.Cache
	}
	if !IsNil(o.CacheTimeout) {
		toSerialize["cacheTimeout"] = o.CacheTimeout
	}
	if !IsNil(o.DataPath) {
		toSerialize["dataPath"] = o.DataPath
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EndpointReturnsList) {
		toSerialize["endpointReturnsList"] = o.EndpointReturnsList
	}
	if !IsNil(o.ParentDataPath) {
		toSerialize["parentDataPath"] = o.ParentDataPath
	}
	if !IsNil(o.PortDataPath) {
		toSerialize["portDataPath"] = o.PortDataPath
	}
	if !IsNil(o.QueryEndpoint) {
		toSerialize["queryEndpoint"] = o.QueryEndpoint
	}
	if !IsNil(o.TargetPath) {
		toSerialize["targetPath"] = o.TargetPath
	}
	if !IsNil(o.UseNestedQuery) {
		toSerialize["useNestedQuery"] = o.UseNestedQuery
	}
	if !IsNil(o.UseTargetList) {
		toSerialize["useTargetList"] = o.UseTargetList
	}
	return toSerialize, nil
}

type NullableServiceDiscovery struct {
	value *ServiceDiscovery
	isSet bool
}

func (v NullableServiceDiscovery) Get() *ServiceDiscovery {
	return v.value
}

func (v *NullableServiceDiscovery) Set(val *ServiceDiscovery) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDiscovery) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDiscovery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDiscovery(val *ServiceDiscovery) *NullableServiceDiscovery {
	return &NullableServiceDiscovery{value: val, isSet: true}
}

func (v NullableServiceDiscovery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDiscovery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
