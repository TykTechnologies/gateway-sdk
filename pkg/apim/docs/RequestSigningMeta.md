# RequestSigningMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**CertificateId** | Pointer to **string** |  | [optional] 
**HeaderList** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**SignatureHeader** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestSigningMeta

`func NewRequestSigningMeta() *RequestSigningMeta`

NewRequestSigningMeta instantiates a new RequestSigningMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSigningMetaWithDefaults

`func NewRequestSigningMetaWithDefaults() *RequestSigningMeta`

NewRequestSigningMetaWithDefaults instantiates a new RequestSigningMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *RequestSigningMeta) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RequestSigningMeta) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RequestSigningMeta) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RequestSigningMeta) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCertificateId

`func (o *RequestSigningMeta) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *RequestSigningMeta) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *RequestSigningMeta) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *RequestSigningMeta) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetHeaderList

`func (o *RequestSigningMeta) GetHeaderList() []string`

GetHeaderList returns the HeaderList field if non-nil, zero value otherwise.

### GetHeaderListOk

`func (o *RequestSigningMeta) GetHeaderListOk() (*[]string, bool)`

GetHeaderListOk returns a tuple with the HeaderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderList

`func (o *RequestSigningMeta) SetHeaderList(v []string)`

SetHeaderList sets HeaderList field to given value.

### HasHeaderList

`func (o *RequestSigningMeta) HasHeaderList() bool`

HasHeaderList returns a boolean if a field has been set.

### SetHeaderListNil

`func (o *RequestSigningMeta) SetHeaderListNil(b bool)`

 SetHeaderListNil sets the value for HeaderList to be an explicit nil

### UnsetHeaderList
`func (o *RequestSigningMeta) UnsetHeaderList()`

UnsetHeaderList ensures that no value is present for HeaderList, not even an explicit nil
### GetIsEnabled

`func (o *RequestSigningMeta) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RequestSigningMeta) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RequestSigningMeta) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *RequestSigningMeta) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetKeyId

`func (o *RequestSigningMeta) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *RequestSigningMeta) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *RequestSigningMeta) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *RequestSigningMeta) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetSecret

`func (o *RequestSigningMeta) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RequestSigningMeta) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RequestSigningMeta) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RequestSigningMeta) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSignatureHeader

`func (o *RequestSigningMeta) GetSignatureHeader() string`

GetSignatureHeader returns the SignatureHeader field if non-nil, zero value otherwise.

### GetSignatureHeaderOk

`func (o *RequestSigningMeta) GetSignatureHeaderOk() (*string, bool)`

GetSignatureHeaderOk returns a tuple with the SignatureHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHeader

`func (o *RequestSigningMeta) SetSignatureHeader(v string)`

SetSignatureHeader sets SignatureHeader field to given value.

### HasSignatureHeader

`func (o *RequestSigningMeta) HasSignatureHeader() bool`

HasSignatureHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


