# AuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaderName** | Pointer to **string** |  | [optional] 
**CookieName** | Pointer to **string** |  | [optional] 
**DisableHeader** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParamName** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**SignatureConfig**](SignatureConfig.md) |  | [optional] 
**UseCertificate** | Pointer to **bool** |  | [optional] 
**UseCookie** | Pointer to **bool** |  | [optional] 
**UseParam** | Pointer to **bool** |  | [optional] 
**ValidateSignature** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthConfig

`func NewAuthConfig() *AuthConfig`

NewAuthConfig instantiates a new AuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigWithDefaults

`func NewAuthConfigWithDefaults() *AuthConfig`

NewAuthConfigWithDefaults instantiates a new AuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaderName

`func (o *AuthConfig) GetAuthHeaderName() string`

GetAuthHeaderName returns the AuthHeaderName field if non-nil, zero value otherwise.

### GetAuthHeaderNameOk

`func (o *AuthConfig) GetAuthHeaderNameOk() (*string, bool)`

GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderName

`func (o *AuthConfig) SetAuthHeaderName(v string)`

SetAuthHeaderName sets AuthHeaderName field to given value.

### HasAuthHeaderName

`func (o *AuthConfig) HasAuthHeaderName() bool`

HasAuthHeaderName returns a boolean if a field has been set.

### GetCookieName

`func (o *AuthConfig) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *AuthConfig) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *AuthConfig) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *AuthConfig) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetDisableHeader

`func (o *AuthConfig) GetDisableHeader() bool`

GetDisableHeader returns the DisableHeader field if non-nil, zero value otherwise.

### GetDisableHeaderOk

`func (o *AuthConfig) GetDisableHeaderOk() (*bool, bool)`

GetDisableHeaderOk returns a tuple with the DisableHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHeader

`func (o *AuthConfig) SetDisableHeader(v bool)`

SetDisableHeader sets DisableHeader field to given value.

### HasDisableHeader

`func (o *AuthConfig) HasDisableHeader() bool`

HasDisableHeader returns a boolean if a field has been set.

### GetName

`func (o *AuthConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParamName

`func (o *AuthConfig) GetParamName() string`

GetParamName returns the ParamName field if non-nil, zero value otherwise.

### GetParamNameOk

`func (o *AuthConfig) GetParamNameOk() (*string, bool)`

GetParamNameOk returns a tuple with the ParamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamName

`func (o *AuthConfig) SetParamName(v string)`

SetParamName sets ParamName field to given value.

### HasParamName

`func (o *AuthConfig) HasParamName() bool`

HasParamName returns a boolean if a field has been set.

### GetSignature

`func (o *AuthConfig) GetSignature() SignatureConfig`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AuthConfig) GetSignatureOk() (*SignatureConfig, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AuthConfig) SetSignature(v SignatureConfig)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AuthConfig) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetUseCertificate

`func (o *AuthConfig) GetUseCertificate() bool`

GetUseCertificate returns the UseCertificate field if non-nil, zero value otherwise.

### GetUseCertificateOk

`func (o *AuthConfig) GetUseCertificateOk() (*bool, bool)`

GetUseCertificateOk returns a tuple with the UseCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCertificate

`func (o *AuthConfig) SetUseCertificate(v bool)`

SetUseCertificate sets UseCertificate field to given value.

### HasUseCertificate

`func (o *AuthConfig) HasUseCertificate() bool`

HasUseCertificate returns a boolean if a field has been set.

### GetUseCookie

`func (o *AuthConfig) GetUseCookie() bool`

GetUseCookie returns the UseCookie field if non-nil, zero value otherwise.

### GetUseCookieOk

`func (o *AuthConfig) GetUseCookieOk() (*bool, bool)`

GetUseCookieOk returns a tuple with the UseCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCookie

`func (o *AuthConfig) SetUseCookie(v bool)`

SetUseCookie sets UseCookie field to given value.

### HasUseCookie

`func (o *AuthConfig) HasUseCookie() bool`

HasUseCookie returns a boolean if a field has been set.

### GetUseParam

`func (o *AuthConfig) GetUseParam() bool`

GetUseParam returns the UseParam field if non-nil, zero value otherwise.

### GetUseParamOk

`func (o *AuthConfig) GetUseParamOk() (*bool, bool)`

GetUseParamOk returns a tuple with the UseParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParam

`func (o *AuthConfig) SetUseParam(v bool)`

SetUseParam sets UseParam field to given value.

### HasUseParam

`func (o *AuthConfig) HasUseParam() bool`

HasUseParam returns a boolean if a field has been set.

### GetValidateSignature

`func (o *AuthConfig) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *AuthConfig) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *AuthConfig) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *AuthConfig) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


