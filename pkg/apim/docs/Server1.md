# Server1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**map[string]ServerVariable**](ServerVariable.md) |  | [optional] 

## Methods

### NewServer1

`func NewServer1(url string, ) *Server1`

NewServer1 instantiates a new Server1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServer1WithDefaults

`func NewServer1WithDefaults() *Server1`

NewServer1WithDefaults instantiates a new Server1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Server1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Server1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Server1) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *Server1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Server1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Server1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Server1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariables

`func (o *Server1) GetVariables() map[string]ServerVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Server1) GetVariablesOk() (*map[string]ServerVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Server1) SetVariables(v map[string]ServerVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Server1) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


