# ProxyConfigTransport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyUrl** | Pointer to **string** |  | [optional] 
**SslCiphers** | Pointer to **[]string** |  | [optional] 
**SslForceCommonNameCheck** | Pointer to **bool** |  | [optional] 
**SslInsecureSkipVerify** | Pointer to **bool** |  | [optional] 
**SslMaxVersion** | Pointer to **int32** |  | [optional] 
**SslMinVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewProxyConfigTransport

`func NewProxyConfigTransport() *ProxyConfigTransport`

NewProxyConfigTransport instantiates a new ProxyConfigTransport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigTransportWithDefaults

`func NewProxyConfigTransportWithDefaults() *ProxyConfigTransport`

NewProxyConfigTransportWithDefaults instantiates a new ProxyConfigTransport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxyUrl

`func (o *ProxyConfigTransport) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *ProxyConfigTransport) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *ProxyConfigTransport) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *ProxyConfigTransport) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetSslCiphers

`func (o *ProxyConfigTransport) GetSslCiphers() []string`

GetSslCiphers returns the SslCiphers field if non-nil, zero value otherwise.

### GetSslCiphersOk

`func (o *ProxyConfigTransport) GetSslCiphersOk() (*[]string, bool)`

GetSslCiphersOk returns a tuple with the SslCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCiphers

`func (o *ProxyConfigTransport) SetSslCiphers(v []string)`

SetSslCiphers sets SslCiphers field to given value.

### HasSslCiphers

`func (o *ProxyConfigTransport) HasSslCiphers() bool`

HasSslCiphers returns a boolean if a field has been set.

### SetSslCiphersNil

`func (o *ProxyConfigTransport) SetSslCiphersNil(b bool)`

 SetSslCiphersNil sets the value for SslCiphers to be an explicit nil

### UnsetSslCiphers
`func (o *ProxyConfigTransport) UnsetSslCiphers()`

UnsetSslCiphers ensures that no value is present for SslCiphers, not even an explicit nil
### GetSslForceCommonNameCheck

`func (o *ProxyConfigTransport) GetSslForceCommonNameCheck() bool`

GetSslForceCommonNameCheck returns the SslForceCommonNameCheck field if non-nil, zero value otherwise.

### GetSslForceCommonNameCheckOk

`func (o *ProxyConfigTransport) GetSslForceCommonNameCheckOk() (*bool, bool)`

GetSslForceCommonNameCheckOk returns a tuple with the SslForceCommonNameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForceCommonNameCheck

`func (o *ProxyConfigTransport) SetSslForceCommonNameCheck(v bool)`

SetSslForceCommonNameCheck sets SslForceCommonNameCheck field to given value.

### HasSslForceCommonNameCheck

`func (o *ProxyConfigTransport) HasSslForceCommonNameCheck() bool`

HasSslForceCommonNameCheck returns a boolean if a field has been set.

### GetSslInsecureSkipVerify

`func (o *ProxyConfigTransport) GetSslInsecureSkipVerify() bool`

GetSslInsecureSkipVerify returns the SslInsecureSkipVerify field if non-nil, zero value otherwise.

### GetSslInsecureSkipVerifyOk

`func (o *ProxyConfigTransport) GetSslInsecureSkipVerifyOk() (*bool, bool)`

GetSslInsecureSkipVerifyOk returns a tuple with the SslInsecureSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslInsecureSkipVerify

`func (o *ProxyConfigTransport) SetSslInsecureSkipVerify(v bool)`

SetSslInsecureSkipVerify sets SslInsecureSkipVerify field to given value.

### HasSslInsecureSkipVerify

`func (o *ProxyConfigTransport) HasSslInsecureSkipVerify() bool`

HasSslInsecureSkipVerify returns a boolean if a field has been set.

### GetSslMaxVersion

`func (o *ProxyConfigTransport) GetSslMaxVersion() int32`

GetSslMaxVersion returns the SslMaxVersion field if non-nil, zero value otherwise.

### GetSslMaxVersionOk

`func (o *ProxyConfigTransport) GetSslMaxVersionOk() (*int32, bool)`

GetSslMaxVersionOk returns a tuple with the SslMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMaxVersion

`func (o *ProxyConfigTransport) SetSslMaxVersion(v int32)`

SetSslMaxVersion sets SslMaxVersion field to given value.

### HasSslMaxVersion

`func (o *ProxyConfigTransport) HasSslMaxVersion() bool`

HasSslMaxVersion returns a boolean if a field has been set.

### GetSslMinVersion

`func (o *ProxyConfigTransport) GetSslMinVersion() int32`

GetSslMinVersion returns the SslMinVersion field if non-nil, zero value otherwise.

### GetSslMinVersionOk

`func (o *ProxyConfigTransport) GetSslMinVersionOk() (*int32, bool)`

GetSslMinVersionOk returns a tuple with the SslMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMinVersion

`func (o *ProxyConfigTransport) SetSslMinVersion(v int32)`

SetSslMinVersion sets SslMinVersion field to given value.

### HasSslMinVersion

`func (o *ProxyConfigTransport) HasSslMinVersion() bool`

HasSslMinVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


