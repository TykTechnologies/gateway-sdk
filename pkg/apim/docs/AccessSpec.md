# AccessSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessSpec

`func NewAccessSpec() *AccessSpec`

NewAccessSpec instantiates a new AccessSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessSpecWithDefaults

`func NewAccessSpecWithDefaults() *AccessSpec`

NewAccessSpecWithDefaults instantiates a new AccessSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *AccessSpec) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *AccessSpec) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *AccessSpec) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *AccessSpec) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### SetMethodsNil

`func (o *AccessSpec) SetMethodsNil(b bool)`

 SetMethodsNil sets the value for Methods to be an explicit nil

### UnsetMethods
`func (o *AccessSpec) UnsetMethods()`

UnsetMethods ensures that no value is present for Methods, not even an explicit nil
### GetUrl

`func (o *AccessSpec) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccessSpec) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccessSpec) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccessSpec) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


