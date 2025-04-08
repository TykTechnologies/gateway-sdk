# Upstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificatePinning** | Pointer to [**CertificatePinning**](CertificatePinning.md) |  | [optional] 
**MutualTLS** | Pointer to [**MutualTLS**](MutualTLS.md) |  | [optional] 
**RateLimit** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**ServiceDiscovery** | Pointer to [**ServiceDiscovery**](ServiceDiscovery.md) |  | [optional] 
**Test** | Pointer to [**Test**](Test.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewUpstream

`func NewUpstream() *Upstream`

NewUpstream instantiates a new Upstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamWithDefaults

`func NewUpstreamWithDefaults() *Upstream`

NewUpstreamWithDefaults instantiates a new Upstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificatePinning

`func (o *Upstream) GetCertificatePinning() CertificatePinning`

GetCertificatePinning returns the CertificatePinning field if non-nil, zero value otherwise.

### GetCertificatePinningOk

`func (o *Upstream) GetCertificatePinningOk() (*CertificatePinning, bool)`

GetCertificatePinningOk returns a tuple with the CertificatePinning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinning

`func (o *Upstream) SetCertificatePinning(v CertificatePinning)`

SetCertificatePinning sets CertificatePinning field to given value.

### HasCertificatePinning

`func (o *Upstream) HasCertificatePinning() bool`

HasCertificatePinning returns a boolean if a field has been set.

### GetMutualTLS

`func (o *Upstream) GetMutualTLS() MutualTLS`

GetMutualTLS returns the MutualTLS field if non-nil, zero value otherwise.

### GetMutualTLSOk

`func (o *Upstream) GetMutualTLSOk() (*MutualTLS, bool)`

GetMutualTLSOk returns a tuple with the MutualTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualTLS

`func (o *Upstream) SetMutualTLS(v MutualTLS)`

SetMutualTLS sets MutualTLS field to given value.

### HasMutualTLS

`func (o *Upstream) HasMutualTLS() bool`

HasMutualTLS returns a boolean if a field has been set.

### GetRateLimit

`func (o *Upstream) GetRateLimit() RateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *Upstream) GetRateLimitOk() (*RateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *Upstream) SetRateLimit(v RateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *Upstream) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *Upstream) GetServiceDiscovery() ServiceDiscovery`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *Upstream) GetServiceDiscoveryOk() (*ServiceDiscovery, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *Upstream) SetServiceDiscovery(v ServiceDiscovery)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *Upstream) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetTest

`func (o *Upstream) GetTest() Test`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Upstream) GetTestOk() (*Test, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Upstream) SetTest(v Test)`

SetTest sets Test field to given value.

### HasTest

`func (o *Upstream) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetUrl

`func (o *Upstream) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Upstream) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Upstream) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Upstream) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


