# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**map[string]ServerVariable**](ServerVariable.md) |  | [optional] 

## Methods

### NewServer

`func NewServer(url string, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Server) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Server) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Server) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *Server) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Server) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Server) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Server) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariables

`func (o *Server) GetVariables() map[string]ServerVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Server) GetVariablesOk() (*map[string]ServerVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Server) SetVariables(v map[string]ServerVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Server) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


