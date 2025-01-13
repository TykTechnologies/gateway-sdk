# AccessDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowanceScope** | Pointer to **string** |  | [optional] 
**AllowedTypes** | Pointer to [**[]GraphqlType**](GraphqlType.md) |  | [optional] 
**AllowedUrls** | Pointer to [**[]AccessSpec**](AccessSpec.md) |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**ApiName** | Pointer to **string** |  | [optional] 
**DisableIntrospection** | Pointer to **bool** |  | [optional] 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) |  | [optional] 
**FieldAccessRights** | Pointer to [**[]FieldAccessDefinition**](FieldAccessDefinition.md) |  | [optional] 
**Limit** | Pointer to [**APILimit**](APILimit.md) |  | [optional] 
**RestrictedTypes** | Pointer to [**[]GraphqlType**](GraphqlType.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessDefinition

`func NewAccessDefinition() *AccessDefinition`

NewAccessDefinition instantiates a new AccessDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDefinitionWithDefaults

`func NewAccessDefinitionWithDefaults() *AccessDefinition`

NewAccessDefinitionWithDefaults instantiates a new AccessDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowanceScope

`func (o *AccessDefinition) GetAllowanceScope() string`

GetAllowanceScope returns the AllowanceScope field if non-nil, zero value otherwise.

### GetAllowanceScopeOk

`func (o *AccessDefinition) GetAllowanceScopeOk() (*string, bool)`

GetAllowanceScopeOk returns a tuple with the AllowanceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceScope

`func (o *AccessDefinition) SetAllowanceScope(v string)`

SetAllowanceScope sets AllowanceScope field to given value.

### HasAllowanceScope

`func (o *AccessDefinition) HasAllowanceScope() bool`

HasAllowanceScope returns a boolean if a field has been set.

### GetAllowedTypes

`func (o *AccessDefinition) GetAllowedTypes() []GraphqlType`

GetAllowedTypes returns the AllowedTypes field if non-nil, zero value otherwise.

### GetAllowedTypesOk

`func (o *AccessDefinition) GetAllowedTypesOk() (*[]GraphqlType, bool)`

GetAllowedTypesOk returns a tuple with the AllowedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTypes

`func (o *AccessDefinition) SetAllowedTypes(v []GraphqlType)`

SetAllowedTypes sets AllowedTypes field to given value.

### HasAllowedTypes

`func (o *AccessDefinition) HasAllowedTypes() bool`

HasAllowedTypes returns a boolean if a field has been set.

### SetAllowedTypesNil

`func (o *AccessDefinition) SetAllowedTypesNil(b bool)`

 SetAllowedTypesNil sets the value for AllowedTypes to be an explicit nil

### UnsetAllowedTypes
`func (o *AccessDefinition) UnsetAllowedTypes()`

UnsetAllowedTypes ensures that no value is present for AllowedTypes, not even an explicit nil
### GetAllowedUrls

`func (o *AccessDefinition) GetAllowedUrls() []AccessSpec`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *AccessDefinition) GetAllowedUrlsOk() (*[]AccessSpec, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *AccessDefinition) SetAllowedUrls(v []AccessSpec)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *AccessDefinition) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### SetAllowedUrlsNil

`func (o *AccessDefinition) SetAllowedUrlsNil(b bool)`

 SetAllowedUrlsNil sets the value for AllowedUrls to be an explicit nil

### UnsetAllowedUrls
`func (o *AccessDefinition) UnsetAllowedUrls()`

UnsetAllowedUrls ensures that no value is present for AllowedUrls, not even an explicit nil
### GetApiId

`func (o *AccessDefinition) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *AccessDefinition) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *AccessDefinition) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *AccessDefinition) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiName

`func (o *AccessDefinition) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *AccessDefinition) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *AccessDefinition) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *AccessDefinition) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetDisableIntrospection

`func (o *AccessDefinition) GetDisableIntrospection() bool`

GetDisableIntrospection returns the DisableIntrospection field if non-nil, zero value otherwise.

### GetDisableIntrospectionOk

`func (o *AccessDefinition) GetDisableIntrospectionOk() (*bool, bool)`

GetDisableIntrospectionOk returns a tuple with the DisableIntrospection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIntrospection

`func (o *AccessDefinition) SetDisableIntrospection(v bool)`

SetDisableIntrospection sets DisableIntrospection field to given value.

### HasDisableIntrospection

`func (o *AccessDefinition) HasDisableIntrospection() bool`

HasDisableIntrospection returns a boolean if a field has been set.

### GetEndpoints

`func (o *AccessDefinition) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *AccessDefinition) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *AccessDefinition) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *AccessDefinition) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetFieldAccessRights

`func (o *AccessDefinition) GetFieldAccessRights() []FieldAccessDefinition`

GetFieldAccessRights returns the FieldAccessRights field if non-nil, zero value otherwise.

### GetFieldAccessRightsOk

`func (o *AccessDefinition) GetFieldAccessRightsOk() (*[]FieldAccessDefinition, bool)`

GetFieldAccessRightsOk returns a tuple with the FieldAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAccessRights

`func (o *AccessDefinition) SetFieldAccessRights(v []FieldAccessDefinition)`

SetFieldAccessRights sets FieldAccessRights field to given value.

### HasFieldAccessRights

`func (o *AccessDefinition) HasFieldAccessRights() bool`

HasFieldAccessRights returns a boolean if a field has been set.

### SetFieldAccessRightsNil

`func (o *AccessDefinition) SetFieldAccessRightsNil(b bool)`

 SetFieldAccessRightsNil sets the value for FieldAccessRights to be an explicit nil

### UnsetFieldAccessRights
`func (o *AccessDefinition) UnsetFieldAccessRights()`

UnsetFieldAccessRights ensures that no value is present for FieldAccessRights, not even an explicit nil
### GetLimit

`func (o *AccessDefinition) GetLimit() APILimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccessDefinition) GetLimitOk() (*APILimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccessDefinition) SetLimit(v APILimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccessDefinition) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRestrictedTypes

`func (o *AccessDefinition) GetRestrictedTypes() []GraphqlType`

GetRestrictedTypes returns the RestrictedTypes field if non-nil, zero value otherwise.

### GetRestrictedTypesOk

`func (o *AccessDefinition) GetRestrictedTypesOk() (*[]GraphqlType, bool)`

GetRestrictedTypesOk returns a tuple with the RestrictedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedTypes

`func (o *AccessDefinition) SetRestrictedTypes(v []GraphqlType)`

SetRestrictedTypes sets RestrictedTypes field to given value.

### HasRestrictedTypes

`func (o *AccessDefinition) HasRestrictedTypes() bool`

HasRestrictedTypes returns a boolean if a field has been set.

### SetRestrictedTypesNil

`func (o *AccessDefinition) SetRestrictedTypesNil(b bool)`

 SetRestrictedTypesNil sets the value for RestrictedTypes to be an explicit nil

### UnsetRestrictedTypes
`func (o *AccessDefinition) UnsetRestrictedTypes()`

UnsetRestrictedTypes ensures that no value is present for RestrictedTypes, not even an explicit nil
### GetVersions

`func (o *AccessDefinition) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AccessDefinition) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AccessDefinition) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AccessDefinition) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *AccessDefinition) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *AccessDefinition) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


