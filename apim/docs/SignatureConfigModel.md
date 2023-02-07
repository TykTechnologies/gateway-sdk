# SignatureConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**AllowedClockSkew** | Pointer to **int64** |  | [optional] 
**ErrorCode** | Pointer to **int64** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewSignatureConfigModel

`func NewSignatureConfigModel() *SignatureConfigModel`

NewSignatureConfigModel instantiates a new SignatureConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureConfigModelWithDefaults

`func NewSignatureConfigModelWithDefaults() *SignatureConfigModel`

NewSignatureConfigModelWithDefaults instantiates a new SignatureConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *SignatureConfigModel) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *SignatureConfigModel) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *SignatureConfigModel) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *SignatureConfigModel) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetAllowedClockSkew

`func (o *SignatureConfigModel) GetAllowedClockSkew() int64`

GetAllowedClockSkew returns the AllowedClockSkew field if non-nil, zero value otherwise.

### GetAllowedClockSkewOk

`func (o *SignatureConfigModel) GetAllowedClockSkewOk() (*int64, bool)`

GetAllowedClockSkewOk returns a tuple with the AllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClockSkew

`func (o *SignatureConfigModel) SetAllowedClockSkew(v int64)`

SetAllowedClockSkew sets AllowedClockSkew field to given value.

### HasAllowedClockSkew

`func (o *SignatureConfigModel) HasAllowedClockSkew() bool`

HasAllowedClockSkew returns a boolean if a field has been set.

### GetErrorCode

`func (o *SignatureConfigModel) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SignatureConfigModel) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SignatureConfigModel) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SignatureConfigModel) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SignatureConfigModel) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SignatureConfigModel) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SignatureConfigModel) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SignatureConfigModel) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetHeader

`func (o *SignatureConfigModel) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SignatureConfigModel) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SignatureConfigModel) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SignatureConfigModel) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSecret

`func (o *SignatureConfigModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignatureConfigModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignatureConfigModel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SignatureConfigModel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


