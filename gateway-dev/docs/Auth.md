# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaderName** | Pointer to **string** |  | [optional] 
**CookieName** | Pointer to **string** |  | [optional] 
**ParamName** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**SignatureConfig**](SignatureConfig.md) |  | [optional] 
**UseCertificate** | Pointer to **bool** |  | [optional] 
**UseCookie** | Pointer to **bool** |  | [optional] 
**UseParam** | Pointer to **bool** |  | [optional] 
**ValidateSignature** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuth

`func NewAuth() *Auth`

NewAuth instantiates a new Auth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithDefaults

`func NewAuthWithDefaults() *Auth`

NewAuthWithDefaults instantiates a new Auth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaderName

`func (o *Auth) GetAuthHeaderName() string`

GetAuthHeaderName returns the AuthHeaderName field if non-nil, zero value otherwise.

### GetAuthHeaderNameOk

`func (o *Auth) GetAuthHeaderNameOk() (*string, bool)`

GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderName

`func (o *Auth) SetAuthHeaderName(v string)`

SetAuthHeaderName sets AuthHeaderName field to given value.

### HasAuthHeaderName

`func (o *Auth) HasAuthHeaderName() bool`

HasAuthHeaderName returns a boolean if a field has been set.

### GetCookieName

`func (o *Auth) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *Auth) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *Auth) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *Auth) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetParamName

`func (o *Auth) GetParamName() string`

GetParamName returns the ParamName field if non-nil, zero value otherwise.

### GetParamNameOk

`func (o *Auth) GetParamNameOk() (*string, bool)`

GetParamNameOk returns a tuple with the ParamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamName

`func (o *Auth) SetParamName(v string)`

SetParamName sets ParamName field to given value.

### HasParamName

`func (o *Auth) HasParamName() bool`

HasParamName returns a boolean if a field has been set.

### GetSignature

`func (o *Auth) GetSignature() SignatureConfig`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Auth) GetSignatureOk() (*SignatureConfig, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Auth) SetSignature(v SignatureConfig)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Auth) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetUseCertificate

`func (o *Auth) GetUseCertificate() bool`

GetUseCertificate returns the UseCertificate field if non-nil, zero value otherwise.

### GetUseCertificateOk

`func (o *Auth) GetUseCertificateOk() (*bool, bool)`

GetUseCertificateOk returns a tuple with the UseCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCertificate

`func (o *Auth) SetUseCertificate(v bool)`

SetUseCertificate sets UseCertificate field to given value.

### HasUseCertificate

`func (o *Auth) HasUseCertificate() bool`

HasUseCertificate returns a boolean if a field has been set.

### GetUseCookie

`func (o *Auth) GetUseCookie() bool`

GetUseCookie returns the UseCookie field if non-nil, zero value otherwise.

### GetUseCookieOk

`func (o *Auth) GetUseCookieOk() (*bool, bool)`

GetUseCookieOk returns a tuple with the UseCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCookie

`func (o *Auth) SetUseCookie(v bool)`

SetUseCookie sets UseCookie field to given value.

### HasUseCookie

`func (o *Auth) HasUseCookie() bool`

HasUseCookie returns a boolean if a field has been set.

### GetUseParam

`func (o *Auth) GetUseParam() bool`

GetUseParam returns the UseParam field if non-nil, zero value otherwise.

### GetUseParamOk

`func (o *Auth) GetUseParamOk() (*bool, bool)`

GetUseParamOk returns a tuple with the UseParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParam

`func (o *Auth) SetUseParam(v bool)`

SetUseParam sets UseParam field to given value.

### HasUseParam

`func (o *Auth) HasUseParam() bool`

HasUseParam returns a boolean if a field has been set.

### GetValidateSignature

`func (o *Auth) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *Auth) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *Auth) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *Auth) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


