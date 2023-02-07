# AuthModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaderName** | Pointer to **string** |  | [optional] 
**CookieName** | Pointer to **string** |  | [optional] 
**ParamName** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**SignatureConfigModelModel**](SignatureConfigModel.md) |  | [optional] 
**UseCertificate** | Pointer to **bool** |  | [optional] 
**UseCookie** | Pointer to **bool** |  | [optional] 
**UseParam** | Pointer to **bool** |  | [optional] 
**ValidateSignature** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthModel

`func NewAuthModel() *AuthModel`

NewAuthModel instantiates a new AuthModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthModelWithDefaults

`func NewAuthModelWithDefaults() *AuthModel`

NewAuthModelWithDefaults instantiates a new AuthModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaderName

`func (o *AuthModel) GetAuthHeaderName() string`

GetAuthHeaderName returns the AuthHeaderName field if non-nil, zero value otherwise.

### GetAuthHeaderNameOk

`func (o *AuthModel) GetAuthHeaderNameOk() (*string, bool)`

GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderName

`func (o *AuthModel) SetAuthHeaderName(v string)`

SetAuthHeaderName sets AuthHeaderName field to given value.

### HasAuthHeaderName

`func (o *AuthModel) HasAuthHeaderName() bool`

HasAuthHeaderName returns a boolean if a field has been set.

### GetCookieName

`func (o *AuthModel) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *AuthModel) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *AuthModel) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *AuthModel) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetParamName

`func (o *AuthModel) GetParamName() string`

GetParamName returns the ParamName field if non-nil, zero value otherwise.

### GetParamNameOk

`func (o *AuthModel) GetParamNameOk() (*string, bool)`

GetParamNameOk returns a tuple with the ParamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamName

`func (o *AuthModel) SetParamName(v string)`

SetParamName sets ParamName field to given value.

### HasParamName

`func (o *AuthModel) HasParamName() bool`

HasParamName returns a boolean if a field has been set.

### GetSignature

`func (o *AuthModel) GetSignature() SignatureConfigModelModel`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AuthModel) GetSignatureOk() (*SignatureConfigModelModel, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AuthModel) SetSignature(v SignatureConfigModelModel)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AuthModel) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetUseCertificate

`func (o *AuthModel) GetUseCertificate() bool`

GetUseCertificate returns the UseCertificate field if non-nil, zero value otherwise.

### GetUseCertificateOk

`func (o *AuthModel) GetUseCertificateOk() (*bool, bool)`

GetUseCertificateOk returns a tuple with the UseCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCertificate

`func (o *AuthModel) SetUseCertificate(v bool)`

SetUseCertificate sets UseCertificate field to given value.

### HasUseCertificate

`func (o *AuthModel) HasUseCertificate() bool`

HasUseCertificate returns a boolean if a field has been set.

### GetUseCookie

`func (o *AuthModel) GetUseCookie() bool`

GetUseCookie returns the UseCookie field if non-nil, zero value otherwise.

### GetUseCookieOk

`func (o *AuthModel) GetUseCookieOk() (*bool, bool)`

GetUseCookieOk returns a tuple with the UseCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCookie

`func (o *AuthModel) SetUseCookie(v bool)`

SetUseCookie sets UseCookie field to given value.

### HasUseCookie

`func (o *AuthModel) HasUseCookie() bool`

HasUseCookie returns a boolean if a field has been set.

### GetUseParam

`func (o *AuthModel) GetUseParam() bool`

GetUseParam returns the UseParam field if non-nil, zero value otherwise.

### GetUseParamOk

`func (o *AuthModel) GetUseParamOk() (*bool, bool)`

GetUseParamOk returns a tuple with the UseParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParam

`func (o *AuthModel) SetUseParam(v bool)`

SetUseParam sets UseParam field to given value.

### HasUseParam

`func (o *AuthModel) HasUseParam() bool`

HasUseParam returns a boolean if a field has been set.

### GetValidateSignature

`func (o *AuthModel) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *AuthModel) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *AuthModel) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *AuthModel) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


