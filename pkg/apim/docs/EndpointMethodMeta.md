# EndpointMethodMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEndpointMethodMeta

`func NewEndpointMethodMeta() *EndpointMethodMeta`

NewEndpointMethodMeta instantiates a new EndpointMethodMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointMethodMetaWithDefaults

`func NewEndpointMethodMetaWithDefaults() *EndpointMethodMeta`

NewEndpointMethodMetaWithDefaults instantiates a new EndpointMethodMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EndpointMethodMeta) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EndpointMethodMeta) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EndpointMethodMeta) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EndpointMethodMeta) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCode

`func (o *EndpointMethodMeta) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EndpointMethodMeta) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EndpointMethodMeta) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *EndpointMethodMeta) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *EndpointMethodMeta) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndpointMethodMeta) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndpointMethodMeta) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EndpointMethodMeta) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *EndpointMethodMeta) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EndpointMethodMeta) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EndpointMethodMeta) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EndpointMethodMeta) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *EndpointMethodMeta) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *EndpointMethodMeta) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


