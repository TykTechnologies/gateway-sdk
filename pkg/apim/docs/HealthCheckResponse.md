# HealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**map[string]HealthCheckItem**](HealthCheckItem.md) |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthCheckResponse

`func NewHealthCheckResponse() *HealthCheckResponse`

NewHealthCheckResponse instantiates a new HealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResponseWithDefaults

`func NewHealthCheckResponseWithDefaults() *HealthCheckResponse`

NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HealthCheckResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheckResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheckResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheckResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *HealthCheckResponse) GetDetails() map[string]HealthCheckItem`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HealthCheckResponse) GetDetailsOk() (*map[string]HealthCheckItem, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HealthCheckResponse) SetDetails(v map[string]HealthCheckItem)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HealthCheckResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetOutput

`func (o *HealthCheckResponse) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *HealthCheckResponse) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *HealthCheckResponse) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *HealthCheckResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *HealthCheckResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthCheckResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthCheckResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthCheckResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


