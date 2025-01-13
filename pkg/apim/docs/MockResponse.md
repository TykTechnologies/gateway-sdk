# MockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FromOASExamples** | Pointer to [**FromOASExamples**](FromOASExamples.md) |  | [optional] 
**Headers** | Pointer to [**[]Header**](Header.md) |  | [optional] 

## Methods

### NewMockResponse

`func NewMockResponse() *MockResponse`

NewMockResponse instantiates a new MockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockResponseWithDefaults

`func NewMockResponseWithDefaults() *MockResponse`

NewMockResponseWithDefaults instantiates a new MockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MockResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MockResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MockResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MockResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCode

`func (o *MockResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MockResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MockResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MockResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *MockResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MockResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MockResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MockResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFromOASExamples

`func (o *MockResponse) GetFromOASExamples() FromOASExamples`

GetFromOASExamples returns the FromOASExamples field if non-nil, zero value otherwise.

### GetFromOASExamplesOk

`func (o *MockResponse) GetFromOASExamplesOk() (*FromOASExamples, bool)`

GetFromOASExamplesOk returns a tuple with the FromOASExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOASExamples

`func (o *MockResponse) SetFromOASExamples(v FromOASExamples)`

SetFromOASExamples sets FromOASExamples field to given value.

### HasFromOASExamples

`func (o *MockResponse) HasFromOASExamples() bool`

HasFromOASExamples returns a boolean if a field has been set.

### GetHeaders

`func (o *MockResponse) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MockResponse) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MockResponse) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MockResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


