# ServerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**map[string]ServerVariableModelModel**](ServerVariableModel.md) |  | [optional] 

## Methods

### NewServerModel

`func NewServerModel(url string, ) *ServerModel`

NewServerModel instantiates a new ServerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerModelWithDefaults

`func NewServerModelWithDefaults() *ServerModel`

NewServerModelWithDefaults instantiates a new ServerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ServerModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *ServerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariables

`func (o *ServerModel) GetVariables() map[string]ServerVariableModelModel`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ServerModel) GetVariablesOk() (*map[string]ServerVariableModelModel, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ServerModel) SetVariables(v map[string]ServerVariableModelModel)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ServerModel) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


