# APIDefinitionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**CORS** | Pointer to [**APIDefinitionCORSModelModel**](APIDefinitionCORSModel.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowedIps** | Pointer to **[]string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Auth** | Pointer to [**AuthModelModel**](AuthModel.md) |  | [optional] 
**AuthProvider** | Pointer to [**AuthProviderMetaModelModel**](AuthProviderMetaModel.md) |  | [optional] 
**BaseIdentityProvidedBy** | Pointer to **string** |  | [optional] 
**BasicAuth** | Pointer to [**APIDefinitionBasicAuthModelModel**](APIDefinitionBasicAuthModel.md) |  | [optional] 
**BlacklistedIps** | Pointer to **[]string** |  | [optional] 
**CacheOptions** | Pointer to [**CacheOptionsModelModel**](CacheOptionsModel.md) |  | [optional] 
**Certificates** | Pointer to **[]string** |  | [optional] 
**ClientCertificates** | Pointer to **[]string** |  | [optional] 
**ConfigData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CustomMiddleware** | Pointer to [**MiddlewareSectionModelModel**](MiddlewareSectionModel.md) |  | [optional] 
**CustomMiddlewareBundle** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**APIDefinitionDefinitionModelModel**](APIDefinitionDefinitionModel.md) |  | [optional] 
**DisableQuota** | Pointer to **bool** |  | [optional] 
**DisableRateLimit** | Pointer to **bool** |  | [optional] 
**DoNotTrack** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DontSetQuotaOnCreate** | Pointer to **bool** |  | [optional] 
**EnableBatchRequestSupport** | Pointer to **bool** |  | [optional] 
**EnableContextVars** | Pointer to **bool** |  | [optional] 
**EnableCoprocessAuth** | Pointer to **bool** |  | [optional] 
**EnableIpBlacklisting** | Pointer to **bool** |  | [optional] 
**EnableIpWhitelisting** | Pointer to **bool** |  | [optional] 
**EnableJwt** | Pointer to **bool** |  | [optional] 
**EnableSignatureChecking** | Pointer to **bool** |  | [optional] 
**EventHandlers** | Pointer to [**EventHandlerMetaConfigModelModel**](EventHandlerMetaConfigModel.md) |  | [optional] 
**ExpireAnalyticsAfter** | Pointer to **int64** |  | [optional] 
**GlobalRateLimit** | Pointer to [**GlobalRateLimitModelModel**](GlobalRateLimitModel.md) |  | [optional] 
**HmacAllowedAlgorithms** | Pointer to **[]string** |  | [optional] 
**HmacAllowedClockSkew** | Pointer to **float64** |  | [optional] 
**Id** | Pointer to **string** | http://www.mongodb.org/display/DOCS/Object+IDs | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**JwtClientBaseField** | Pointer to **string** |  | [optional] 
**JwtExpiresAtValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtIdentityBaseField** | Pointer to **string** |  | [optional] 
**JwtIssuedAtValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtNotBeforeValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtPolicyFieldName** | Pointer to **string** |  | [optional] 
**JwtScopeClaimName** | Pointer to **string** |  | [optional] 
**JwtScopeToPolicyMapping** | Pointer to **map[string]string** |  | [optional] 
**JwtSigningMethod** | Pointer to **string** |  | [optional] 
**JwtSkipKid** | Pointer to **bool** |  | [optional] 
**JwtSource** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**NotificationsManagerModelModel**](NotificationsManagerModel.md) |  | [optional] 
**OauthMeta** | Pointer to [**APIDefinitionOauthMetaModelModel**](APIDefinitionOauthMetaModel.md) |  | [optional] 
**OpenidOptions** | Pointer to [**OpenIDOptionsModelModel**](OpenIDOptionsModel.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**PinnedPublicKeys** | Pointer to **map[string]string** |  | [optional] 
**Proxy** | Pointer to [**APIDefinitionProxyModelModel**](APIDefinitionProxyModel.md) |  | [optional] 
**ResponseProcessors** | Pointer to [**[]ResponseProcessorModelModel**](ResponseProcessorModelModel.md) |  | [optional] 
**SessionLifetime** | Pointer to **int64** |  | [optional] 
**SessionProvider** | Pointer to [**SessionProviderMetaModelModel**](SessionProviderMetaModel.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StripAuthData** | Pointer to **bool** |  | [optional] 
**TagHeaders** | Pointer to **[]string** |  | [optional] 
**UpstreamCertificates** | Pointer to **map[string]string** |  | [optional] 
**UptimeTests** | Pointer to [**APIDefinitionUptimeTestsModelModel**](APIDefinitionUptimeTestsModel.md) |  | [optional] 
**UseBasicAuth** | Pointer to **bool** |  | [optional] 
**UseKeyless** | Pointer to **bool** |  | [optional] 
**UseMutualTlsAuth** | Pointer to **bool** |  | [optional] 
**UseOauth2** | Pointer to **bool** |  | [optional] 
**UseOpenid** | Pointer to **bool** |  | [optional] 
**UseStandardAuth** | Pointer to **bool** |  | [optional] 
**VersionData** | Pointer to [**APIDefinitionVersionDataModelModel**](APIDefinitionVersionDataModel.md) |  | [optional] 

## Methods

### NewAPIDefinitionModel

`func NewAPIDefinitionModel() *APIDefinitionModel`

NewAPIDefinitionModel instantiates a new APIDefinitionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionModelWithDefaults

`func NewAPIDefinitionModelWithDefaults() *APIDefinitionModel`

NewAPIDefinitionModelWithDefaults instantiates a new APIDefinitionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *APIDefinitionModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *APIDefinitionModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *APIDefinitionModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *APIDefinitionModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCORS

`func (o *APIDefinitionModel) GetCORS() APIDefinitionCORSModelModel`

GetCORS returns the CORS field if non-nil, zero value otherwise.

### GetCORSOk

`func (o *APIDefinitionModel) GetCORSOk() (*APIDefinitionCORSModelModel, bool)`

GetCORSOk returns a tuple with the CORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCORS

`func (o *APIDefinitionModel) SetCORS(v APIDefinitionCORSModelModel)`

SetCORS sets CORS field to given value.

### HasCORS

`func (o *APIDefinitionModel) HasCORS() bool`

HasCORS returns a boolean if a field has been set.

### GetActive

`func (o *APIDefinitionModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *APIDefinitionModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *APIDefinitionModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *APIDefinitionModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowedIps

`func (o *APIDefinitionModel) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *APIDefinitionModel) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *APIDefinitionModel) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *APIDefinitionModel) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetApiId

`func (o *APIDefinitionModel) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *APIDefinitionModel) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *APIDefinitionModel) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *APIDefinitionModel) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAuth

`func (o *APIDefinitionModel) GetAuth() AuthModelModel`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *APIDefinitionModel) GetAuthOk() (*AuthModelModel, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *APIDefinitionModel) SetAuth(v AuthModelModel)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *APIDefinitionModel) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthProvider

`func (o *APIDefinitionModel) GetAuthProvider() AuthProviderMetaModelModel`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *APIDefinitionModel) GetAuthProviderOk() (*AuthProviderMetaModelModel, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *APIDefinitionModel) SetAuthProvider(v AuthProviderMetaModelModel)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *APIDefinitionModel) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBaseIdentityProvidedBy

`func (o *APIDefinitionModel) GetBaseIdentityProvidedBy() string`

GetBaseIdentityProvidedBy returns the BaseIdentityProvidedBy field if non-nil, zero value otherwise.

### GetBaseIdentityProvidedByOk

`func (o *APIDefinitionModel) GetBaseIdentityProvidedByOk() (*string, bool)`

GetBaseIdentityProvidedByOk returns a tuple with the BaseIdentityProvidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIdentityProvidedBy

`func (o *APIDefinitionModel) SetBaseIdentityProvidedBy(v string)`

SetBaseIdentityProvidedBy sets BaseIdentityProvidedBy field to given value.

### HasBaseIdentityProvidedBy

`func (o *APIDefinitionModel) HasBaseIdentityProvidedBy() bool`

HasBaseIdentityProvidedBy returns a boolean if a field has been set.

### GetBasicAuth

`func (o *APIDefinitionModel) GetBasicAuth() APIDefinitionBasicAuthModelModel`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *APIDefinitionModel) GetBasicAuthOk() (*APIDefinitionBasicAuthModelModel, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *APIDefinitionModel) SetBasicAuth(v APIDefinitionBasicAuthModelModel)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *APIDefinitionModel) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBlacklistedIps

`func (o *APIDefinitionModel) GetBlacklistedIps() []string`

GetBlacklistedIps returns the BlacklistedIps field if non-nil, zero value otherwise.

### GetBlacklistedIpsOk

`func (o *APIDefinitionModel) GetBlacklistedIpsOk() (*[]string, bool)`

GetBlacklistedIpsOk returns a tuple with the BlacklistedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIps

`func (o *APIDefinitionModel) SetBlacklistedIps(v []string)`

SetBlacklistedIps sets BlacklistedIps field to given value.

### HasBlacklistedIps

`func (o *APIDefinitionModel) HasBlacklistedIps() bool`

HasBlacklistedIps returns a boolean if a field has been set.

### GetCacheOptions

`func (o *APIDefinitionModel) GetCacheOptions() CacheOptionsModelModel`

GetCacheOptions returns the CacheOptions field if non-nil, zero value otherwise.

### GetCacheOptionsOk

`func (o *APIDefinitionModel) GetCacheOptionsOk() (*CacheOptionsModelModel, bool)`

GetCacheOptionsOk returns a tuple with the CacheOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptions

`func (o *APIDefinitionModel) SetCacheOptions(v CacheOptionsModelModel)`

SetCacheOptions sets CacheOptions field to given value.

### HasCacheOptions

`func (o *APIDefinitionModel) HasCacheOptions() bool`

HasCacheOptions returns a boolean if a field has been set.

### GetCertificates

`func (o *APIDefinitionModel) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *APIDefinitionModel) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *APIDefinitionModel) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *APIDefinitionModel) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetClientCertificates

`func (o *APIDefinitionModel) GetClientCertificates() []string`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *APIDefinitionModel) GetClientCertificatesOk() (*[]string, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *APIDefinitionModel) SetClientCertificates(v []string)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *APIDefinitionModel) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### GetConfigData

`func (o *APIDefinitionModel) GetConfigData() map[string]map[string]interface{}`

GetConfigData returns the ConfigData field if non-nil, zero value otherwise.

### GetConfigDataOk

`func (o *APIDefinitionModel) GetConfigDataOk() (*map[string]map[string]interface{}, bool)`

GetConfigDataOk returns a tuple with the ConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigData

`func (o *APIDefinitionModel) SetConfigData(v map[string]map[string]interface{})`

SetConfigData sets ConfigData field to given value.

### HasConfigData

`func (o *APIDefinitionModel) HasConfigData() bool`

HasConfigData returns a boolean if a field has been set.

### GetCustomMiddleware

`func (o *APIDefinitionModel) GetCustomMiddleware() MiddlewareSectionModelModel`

GetCustomMiddleware returns the CustomMiddleware field if non-nil, zero value otherwise.

### GetCustomMiddlewareOk

`func (o *APIDefinitionModel) GetCustomMiddlewareOk() (*MiddlewareSectionModelModel, bool)`

GetCustomMiddlewareOk returns a tuple with the CustomMiddleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddleware

`func (o *APIDefinitionModel) SetCustomMiddleware(v MiddlewareSectionModelModel)`

SetCustomMiddleware sets CustomMiddleware field to given value.

### HasCustomMiddleware

`func (o *APIDefinitionModel) HasCustomMiddleware() bool`

HasCustomMiddleware returns a boolean if a field has been set.

### GetCustomMiddlewareBundle

`func (o *APIDefinitionModel) GetCustomMiddlewareBundle() string`

GetCustomMiddlewareBundle returns the CustomMiddlewareBundle field if non-nil, zero value otherwise.

### GetCustomMiddlewareBundleOk

`func (o *APIDefinitionModel) GetCustomMiddlewareBundleOk() (*string, bool)`

GetCustomMiddlewareBundleOk returns a tuple with the CustomMiddlewareBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddlewareBundle

`func (o *APIDefinitionModel) SetCustomMiddlewareBundle(v string)`

SetCustomMiddlewareBundle sets CustomMiddlewareBundle field to given value.

### HasCustomMiddlewareBundle

`func (o *APIDefinitionModel) HasCustomMiddlewareBundle() bool`

HasCustomMiddlewareBundle returns a boolean if a field has been set.

### GetDefinition

`func (o *APIDefinitionModel) GetDefinition() APIDefinitionDefinitionModelModel`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *APIDefinitionModel) GetDefinitionOk() (*APIDefinitionDefinitionModelModel, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *APIDefinitionModel) SetDefinition(v APIDefinitionDefinitionModelModel)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *APIDefinitionModel) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDisableQuota

`func (o *APIDefinitionModel) GetDisableQuota() bool`

GetDisableQuota returns the DisableQuota field if non-nil, zero value otherwise.

### GetDisableQuotaOk

`func (o *APIDefinitionModel) GetDisableQuotaOk() (*bool, bool)`

GetDisableQuotaOk returns a tuple with the DisableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableQuota

`func (o *APIDefinitionModel) SetDisableQuota(v bool)`

SetDisableQuota sets DisableQuota field to given value.

### HasDisableQuota

`func (o *APIDefinitionModel) HasDisableQuota() bool`

HasDisableQuota returns a boolean if a field has been set.

### GetDisableRateLimit

`func (o *APIDefinitionModel) GetDisableRateLimit() bool`

GetDisableRateLimit returns the DisableRateLimit field if non-nil, zero value otherwise.

### GetDisableRateLimitOk

`func (o *APIDefinitionModel) GetDisableRateLimitOk() (*bool, bool)`

GetDisableRateLimitOk returns a tuple with the DisableRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRateLimit

`func (o *APIDefinitionModel) SetDisableRateLimit(v bool)`

SetDisableRateLimit sets DisableRateLimit field to given value.

### HasDisableRateLimit

`func (o *APIDefinitionModel) HasDisableRateLimit() bool`

HasDisableRateLimit returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *APIDefinitionModel) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *APIDefinitionModel) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *APIDefinitionModel) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *APIDefinitionModel) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetDomain

`func (o *APIDefinitionModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *APIDefinitionModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *APIDefinitionModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *APIDefinitionModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDontSetQuotaOnCreate

`func (o *APIDefinitionModel) GetDontSetQuotaOnCreate() bool`

GetDontSetQuotaOnCreate returns the DontSetQuotaOnCreate field if non-nil, zero value otherwise.

### GetDontSetQuotaOnCreateOk

`func (o *APIDefinitionModel) GetDontSetQuotaOnCreateOk() (*bool, bool)`

GetDontSetQuotaOnCreateOk returns a tuple with the DontSetQuotaOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontSetQuotaOnCreate

`func (o *APIDefinitionModel) SetDontSetQuotaOnCreate(v bool)`

SetDontSetQuotaOnCreate sets DontSetQuotaOnCreate field to given value.

### HasDontSetQuotaOnCreate

`func (o *APIDefinitionModel) HasDontSetQuotaOnCreate() bool`

HasDontSetQuotaOnCreate returns a boolean if a field has been set.

### GetEnableBatchRequestSupport

`func (o *APIDefinitionModel) GetEnableBatchRequestSupport() bool`

GetEnableBatchRequestSupport returns the EnableBatchRequestSupport field if non-nil, zero value otherwise.

### GetEnableBatchRequestSupportOk

`func (o *APIDefinitionModel) GetEnableBatchRequestSupportOk() (*bool, bool)`

GetEnableBatchRequestSupportOk returns a tuple with the EnableBatchRequestSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBatchRequestSupport

`func (o *APIDefinitionModel) SetEnableBatchRequestSupport(v bool)`

SetEnableBatchRequestSupport sets EnableBatchRequestSupport field to given value.

### HasEnableBatchRequestSupport

`func (o *APIDefinitionModel) HasEnableBatchRequestSupport() bool`

HasEnableBatchRequestSupport returns a boolean if a field has been set.

### GetEnableContextVars

`func (o *APIDefinitionModel) GetEnableContextVars() bool`

GetEnableContextVars returns the EnableContextVars field if non-nil, zero value otherwise.

### GetEnableContextVarsOk

`func (o *APIDefinitionModel) GetEnableContextVarsOk() (*bool, bool)`

GetEnableContextVarsOk returns a tuple with the EnableContextVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContextVars

`func (o *APIDefinitionModel) SetEnableContextVars(v bool)`

SetEnableContextVars sets EnableContextVars field to given value.

### HasEnableContextVars

`func (o *APIDefinitionModel) HasEnableContextVars() bool`

HasEnableContextVars returns a boolean if a field has been set.

### GetEnableCoprocessAuth

`func (o *APIDefinitionModel) GetEnableCoprocessAuth() bool`

GetEnableCoprocessAuth returns the EnableCoprocessAuth field if non-nil, zero value otherwise.

### GetEnableCoprocessAuthOk

`func (o *APIDefinitionModel) GetEnableCoprocessAuthOk() (*bool, bool)`

GetEnableCoprocessAuthOk returns a tuple with the EnableCoprocessAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCoprocessAuth

`func (o *APIDefinitionModel) SetEnableCoprocessAuth(v bool)`

SetEnableCoprocessAuth sets EnableCoprocessAuth field to given value.

### HasEnableCoprocessAuth

`func (o *APIDefinitionModel) HasEnableCoprocessAuth() bool`

HasEnableCoprocessAuth returns a boolean if a field has been set.

### GetEnableIpBlacklisting

`func (o *APIDefinitionModel) GetEnableIpBlacklisting() bool`

GetEnableIpBlacklisting returns the EnableIpBlacklisting field if non-nil, zero value otherwise.

### GetEnableIpBlacklistingOk

`func (o *APIDefinitionModel) GetEnableIpBlacklistingOk() (*bool, bool)`

GetEnableIpBlacklistingOk returns a tuple with the EnableIpBlacklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpBlacklisting

`func (o *APIDefinitionModel) SetEnableIpBlacklisting(v bool)`

SetEnableIpBlacklisting sets EnableIpBlacklisting field to given value.

### HasEnableIpBlacklisting

`func (o *APIDefinitionModel) HasEnableIpBlacklisting() bool`

HasEnableIpBlacklisting returns a boolean if a field has been set.

### GetEnableIpWhitelisting

`func (o *APIDefinitionModel) GetEnableIpWhitelisting() bool`

GetEnableIpWhitelisting returns the EnableIpWhitelisting field if non-nil, zero value otherwise.

### GetEnableIpWhitelistingOk

`func (o *APIDefinitionModel) GetEnableIpWhitelistingOk() (*bool, bool)`

GetEnableIpWhitelistingOk returns a tuple with the EnableIpWhitelisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpWhitelisting

`func (o *APIDefinitionModel) SetEnableIpWhitelisting(v bool)`

SetEnableIpWhitelisting sets EnableIpWhitelisting field to given value.

### HasEnableIpWhitelisting

`func (o *APIDefinitionModel) HasEnableIpWhitelisting() bool`

HasEnableIpWhitelisting returns a boolean if a field has been set.

### GetEnableJwt

`func (o *APIDefinitionModel) GetEnableJwt() bool`

GetEnableJwt returns the EnableJwt field if non-nil, zero value otherwise.

### GetEnableJwtOk

`func (o *APIDefinitionModel) GetEnableJwtOk() (*bool, bool)`

GetEnableJwtOk returns a tuple with the EnableJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJwt

`func (o *APIDefinitionModel) SetEnableJwt(v bool)`

SetEnableJwt sets EnableJwt field to given value.

### HasEnableJwt

`func (o *APIDefinitionModel) HasEnableJwt() bool`

HasEnableJwt returns a boolean if a field has been set.

### GetEnableSignatureChecking

`func (o *APIDefinitionModel) GetEnableSignatureChecking() bool`

GetEnableSignatureChecking returns the EnableSignatureChecking field if non-nil, zero value otherwise.

### GetEnableSignatureCheckingOk

`func (o *APIDefinitionModel) GetEnableSignatureCheckingOk() (*bool, bool)`

GetEnableSignatureCheckingOk returns a tuple with the EnableSignatureChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignatureChecking

`func (o *APIDefinitionModel) SetEnableSignatureChecking(v bool)`

SetEnableSignatureChecking sets EnableSignatureChecking field to given value.

### HasEnableSignatureChecking

`func (o *APIDefinitionModel) HasEnableSignatureChecking() bool`

HasEnableSignatureChecking returns a boolean if a field has been set.

### GetEventHandlers

`func (o *APIDefinitionModel) GetEventHandlers() EventHandlerMetaConfigModelModel`

GetEventHandlers returns the EventHandlers field if non-nil, zero value otherwise.

### GetEventHandlersOk

`func (o *APIDefinitionModel) GetEventHandlersOk() (*EventHandlerMetaConfigModelModel, bool)`

GetEventHandlersOk returns a tuple with the EventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlers

`func (o *APIDefinitionModel) SetEventHandlers(v EventHandlerMetaConfigModelModel)`

SetEventHandlers sets EventHandlers field to given value.

### HasEventHandlers

`func (o *APIDefinitionModel) HasEventHandlers() bool`

HasEventHandlers returns a boolean if a field has been set.

### GetExpireAnalyticsAfter

`func (o *APIDefinitionModel) GetExpireAnalyticsAfter() int64`

GetExpireAnalyticsAfter returns the ExpireAnalyticsAfter field if non-nil, zero value otherwise.

### GetExpireAnalyticsAfterOk

`func (o *APIDefinitionModel) GetExpireAnalyticsAfterOk() (*int64, bool)`

GetExpireAnalyticsAfterOk returns a tuple with the ExpireAnalyticsAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAnalyticsAfter

`func (o *APIDefinitionModel) SetExpireAnalyticsAfter(v int64)`

SetExpireAnalyticsAfter sets ExpireAnalyticsAfter field to given value.

### HasExpireAnalyticsAfter

`func (o *APIDefinitionModel) HasExpireAnalyticsAfter() bool`

HasExpireAnalyticsAfter returns a boolean if a field has been set.

### GetGlobalRateLimit

`func (o *APIDefinitionModel) GetGlobalRateLimit() GlobalRateLimitModelModel`

GetGlobalRateLimit returns the GlobalRateLimit field if non-nil, zero value otherwise.

### GetGlobalRateLimitOk

`func (o *APIDefinitionModel) GetGlobalRateLimitOk() (*GlobalRateLimitModelModel, bool)`

GetGlobalRateLimitOk returns a tuple with the GlobalRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRateLimit

`func (o *APIDefinitionModel) SetGlobalRateLimit(v GlobalRateLimitModelModel)`

SetGlobalRateLimit sets GlobalRateLimit field to given value.

### HasGlobalRateLimit

`func (o *APIDefinitionModel) HasGlobalRateLimit() bool`

HasGlobalRateLimit returns a boolean if a field has been set.

### GetHmacAllowedAlgorithms

`func (o *APIDefinitionModel) GetHmacAllowedAlgorithms() []string`

GetHmacAllowedAlgorithms returns the HmacAllowedAlgorithms field if non-nil, zero value otherwise.

### GetHmacAllowedAlgorithmsOk

`func (o *APIDefinitionModel) GetHmacAllowedAlgorithmsOk() (*[]string, bool)`

GetHmacAllowedAlgorithmsOk returns a tuple with the HmacAllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedAlgorithms

`func (o *APIDefinitionModel) SetHmacAllowedAlgorithms(v []string)`

SetHmacAllowedAlgorithms sets HmacAllowedAlgorithms field to given value.

### HasHmacAllowedAlgorithms

`func (o *APIDefinitionModel) HasHmacAllowedAlgorithms() bool`

HasHmacAllowedAlgorithms returns a boolean if a field has been set.

### GetHmacAllowedClockSkew

`func (o *APIDefinitionModel) GetHmacAllowedClockSkew() float64`

GetHmacAllowedClockSkew returns the HmacAllowedClockSkew field if non-nil, zero value otherwise.

### GetHmacAllowedClockSkewOk

`func (o *APIDefinitionModel) GetHmacAllowedClockSkewOk() (*float64, bool)`

GetHmacAllowedClockSkewOk returns a tuple with the HmacAllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedClockSkew

`func (o *APIDefinitionModel) SetHmacAllowedClockSkew(v float64)`

SetHmacAllowedClockSkew sets HmacAllowedClockSkew field to given value.

### HasHmacAllowedClockSkew

`func (o *APIDefinitionModel) HasHmacAllowedClockSkew() bool`

HasHmacAllowedClockSkew returns a boolean if a field has been set.

### GetId

`func (o *APIDefinitionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIDefinitionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIDefinitionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIDefinitionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternal

`func (o *APIDefinitionModel) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *APIDefinitionModel) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *APIDefinitionModel) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *APIDefinitionModel) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetJwtClientBaseField

`func (o *APIDefinitionModel) GetJwtClientBaseField() string`

GetJwtClientBaseField returns the JwtClientBaseField field if non-nil, zero value otherwise.

### GetJwtClientBaseFieldOk

`func (o *APIDefinitionModel) GetJwtClientBaseFieldOk() (*string, bool)`

GetJwtClientBaseFieldOk returns a tuple with the JwtClientBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClientBaseField

`func (o *APIDefinitionModel) SetJwtClientBaseField(v string)`

SetJwtClientBaseField sets JwtClientBaseField field to given value.

### HasJwtClientBaseField

`func (o *APIDefinitionModel) HasJwtClientBaseField() bool`

HasJwtClientBaseField returns a boolean if a field has been set.

### GetJwtExpiresAtValidationSkew

`func (o *APIDefinitionModel) GetJwtExpiresAtValidationSkew() int32`

GetJwtExpiresAtValidationSkew returns the JwtExpiresAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtExpiresAtValidationSkewOk

`func (o *APIDefinitionModel) GetJwtExpiresAtValidationSkewOk() (*int32, bool)`

GetJwtExpiresAtValidationSkewOk returns a tuple with the JwtExpiresAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresAtValidationSkew

`func (o *APIDefinitionModel) SetJwtExpiresAtValidationSkew(v int32)`

SetJwtExpiresAtValidationSkew sets JwtExpiresAtValidationSkew field to given value.

### HasJwtExpiresAtValidationSkew

`func (o *APIDefinitionModel) HasJwtExpiresAtValidationSkew() bool`

HasJwtExpiresAtValidationSkew returns a boolean if a field has been set.

### GetJwtIdentityBaseField

`func (o *APIDefinitionModel) GetJwtIdentityBaseField() string`

GetJwtIdentityBaseField returns the JwtIdentityBaseField field if non-nil, zero value otherwise.

### GetJwtIdentityBaseFieldOk

`func (o *APIDefinitionModel) GetJwtIdentityBaseFieldOk() (*string, bool)`

GetJwtIdentityBaseFieldOk returns a tuple with the JwtIdentityBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIdentityBaseField

`func (o *APIDefinitionModel) SetJwtIdentityBaseField(v string)`

SetJwtIdentityBaseField sets JwtIdentityBaseField field to given value.

### HasJwtIdentityBaseField

`func (o *APIDefinitionModel) HasJwtIdentityBaseField() bool`

HasJwtIdentityBaseField returns a boolean if a field has been set.

### GetJwtIssuedAtValidationSkew

`func (o *APIDefinitionModel) GetJwtIssuedAtValidationSkew() int32`

GetJwtIssuedAtValidationSkew returns the JwtIssuedAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtIssuedAtValidationSkewOk

`func (o *APIDefinitionModel) GetJwtIssuedAtValidationSkewOk() (*int32, bool)`

GetJwtIssuedAtValidationSkewOk returns a tuple with the JwtIssuedAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIssuedAtValidationSkew

`func (o *APIDefinitionModel) SetJwtIssuedAtValidationSkew(v int32)`

SetJwtIssuedAtValidationSkew sets JwtIssuedAtValidationSkew field to given value.

### HasJwtIssuedAtValidationSkew

`func (o *APIDefinitionModel) HasJwtIssuedAtValidationSkew() bool`

HasJwtIssuedAtValidationSkew returns a boolean if a field has been set.

### GetJwtNotBeforeValidationSkew

`func (o *APIDefinitionModel) GetJwtNotBeforeValidationSkew() int32`

GetJwtNotBeforeValidationSkew returns the JwtNotBeforeValidationSkew field if non-nil, zero value otherwise.

### GetJwtNotBeforeValidationSkewOk

`func (o *APIDefinitionModel) GetJwtNotBeforeValidationSkewOk() (*int32, bool)`

GetJwtNotBeforeValidationSkewOk returns a tuple with the JwtNotBeforeValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtNotBeforeValidationSkew

`func (o *APIDefinitionModel) SetJwtNotBeforeValidationSkew(v int32)`

SetJwtNotBeforeValidationSkew sets JwtNotBeforeValidationSkew field to given value.

### HasJwtNotBeforeValidationSkew

`func (o *APIDefinitionModel) HasJwtNotBeforeValidationSkew() bool`

HasJwtNotBeforeValidationSkew returns a boolean if a field has been set.

### GetJwtPolicyFieldName

`func (o *APIDefinitionModel) GetJwtPolicyFieldName() string`

GetJwtPolicyFieldName returns the JwtPolicyFieldName field if non-nil, zero value otherwise.

### GetJwtPolicyFieldNameOk

`func (o *APIDefinitionModel) GetJwtPolicyFieldNameOk() (*string, bool)`

GetJwtPolicyFieldNameOk returns a tuple with the JwtPolicyFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtPolicyFieldName

`func (o *APIDefinitionModel) SetJwtPolicyFieldName(v string)`

SetJwtPolicyFieldName sets JwtPolicyFieldName field to given value.

### HasJwtPolicyFieldName

`func (o *APIDefinitionModel) HasJwtPolicyFieldName() bool`

HasJwtPolicyFieldName returns a boolean if a field has been set.

### GetJwtScopeClaimName

`func (o *APIDefinitionModel) GetJwtScopeClaimName() string`

GetJwtScopeClaimName returns the JwtScopeClaimName field if non-nil, zero value otherwise.

### GetJwtScopeClaimNameOk

`func (o *APIDefinitionModel) GetJwtScopeClaimNameOk() (*string, bool)`

GetJwtScopeClaimNameOk returns a tuple with the JwtScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeClaimName

`func (o *APIDefinitionModel) SetJwtScopeClaimName(v string)`

SetJwtScopeClaimName sets JwtScopeClaimName field to given value.

### HasJwtScopeClaimName

`func (o *APIDefinitionModel) HasJwtScopeClaimName() bool`

HasJwtScopeClaimName returns a boolean if a field has been set.

### GetJwtScopeToPolicyMapping

`func (o *APIDefinitionModel) GetJwtScopeToPolicyMapping() map[string]string`

GetJwtScopeToPolicyMapping returns the JwtScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetJwtScopeToPolicyMappingOk

`func (o *APIDefinitionModel) GetJwtScopeToPolicyMappingOk() (*map[string]string, bool)`

GetJwtScopeToPolicyMappingOk returns a tuple with the JwtScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeToPolicyMapping

`func (o *APIDefinitionModel) SetJwtScopeToPolicyMapping(v map[string]string)`

SetJwtScopeToPolicyMapping sets JwtScopeToPolicyMapping field to given value.

### HasJwtScopeToPolicyMapping

`func (o *APIDefinitionModel) HasJwtScopeToPolicyMapping() bool`

HasJwtScopeToPolicyMapping returns a boolean if a field has been set.

### GetJwtSigningMethod

`func (o *APIDefinitionModel) GetJwtSigningMethod() string`

GetJwtSigningMethod returns the JwtSigningMethod field if non-nil, zero value otherwise.

### GetJwtSigningMethodOk

`func (o *APIDefinitionModel) GetJwtSigningMethodOk() (*string, bool)`

GetJwtSigningMethodOk returns a tuple with the JwtSigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSigningMethod

`func (o *APIDefinitionModel) SetJwtSigningMethod(v string)`

SetJwtSigningMethod sets JwtSigningMethod field to given value.

### HasJwtSigningMethod

`func (o *APIDefinitionModel) HasJwtSigningMethod() bool`

HasJwtSigningMethod returns a boolean if a field has been set.

### GetJwtSkipKid

`func (o *APIDefinitionModel) GetJwtSkipKid() bool`

GetJwtSkipKid returns the JwtSkipKid field if non-nil, zero value otherwise.

### GetJwtSkipKidOk

`func (o *APIDefinitionModel) GetJwtSkipKidOk() (*bool, bool)`

GetJwtSkipKidOk returns a tuple with the JwtSkipKid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSkipKid

`func (o *APIDefinitionModel) SetJwtSkipKid(v bool)`

SetJwtSkipKid sets JwtSkipKid field to given value.

### HasJwtSkipKid

`func (o *APIDefinitionModel) HasJwtSkipKid() bool`

HasJwtSkipKid returns a boolean if a field has been set.

### GetJwtSource

`func (o *APIDefinitionModel) GetJwtSource() string`

GetJwtSource returns the JwtSource field if non-nil, zero value otherwise.

### GetJwtSourceOk

`func (o *APIDefinitionModel) GetJwtSourceOk() (*string, bool)`

GetJwtSourceOk returns a tuple with the JwtSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSource

`func (o *APIDefinitionModel) SetJwtSource(v string)`

SetJwtSource sets JwtSource field to given value.

### HasJwtSource

`func (o *APIDefinitionModel) HasJwtSource() bool`

HasJwtSource returns a boolean if a field has been set.

### GetName

`func (o *APIDefinitionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIDefinitionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIDefinitionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIDefinitionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *APIDefinitionModel) GetNotifications() NotificationsManagerModelModel`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *APIDefinitionModel) GetNotificationsOk() (*NotificationsManagerModelModel, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *APIDefinitionModel) SetNotifications(v NotificationsManagerModelModel)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *APIDefinitionModel) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOauthMeta

`func (o *APIDefinitionModel) GetOauthMeta() APIDefinitionOauthMetaModelModel`

GetOauthMeta returns the OauthMeta field if non-nil, zero value otherwise.

### GetOauthMetaOk

`func (o *APIDefinitionModel) GetOauthMetaOk() (*APIDefinitionOauthMetaModelModel, bool)`

GetOauthMetaOk returns a tuple with the OauthMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthMeta

`func (o *APIDefinitionModel) SetOauthMeta(v APIDefinitionOauthMetaModelModel)`

SetOauthMeta sets OauthMeta field to given value.

### HasOauthMeta

`func (o *APIDefinitionModel) HasOauthMeta() bool`

HasOauthMeta returns a boolean if a field has been set.

### GetOpenidOptions

`func (o *APIDefinitionModel) GetOpenidOptions() OpenIDOptionsModelModel`

GetOpenidOptions returns the OpenidOptions field if non-nil, zero value otherwise.

### GetOpenidOptionsOk

`func (o *APIDefinitionModel) GetOpenidOptionsOk() (*OpenIDOptionsModelModel, bool)`

GetOpenidOptionsOk returns a tuple with the OpenidOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidOptions

`func (o *APIDefinitionModel) SetOpenidOptions(v OpenIDOptionsModelModel)`

SetOpenidOptions sets OpenidOptions field to given value.

### HasOpenidOptions

`func (o *APIDefinitionModel) HasOpenidOptions() bool`

HasOpenidOptions returns a boolean if a field has been set.

### GetOrgId

`func (o *APIDefinitionModel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *APIDefinitionModel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *APIDefinitionModel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *APIDefinitionModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPinnedPublicKeys

`func (o *APIDefinitionModel) GetPinnedPublicKeys() map[string]string`

GetPinnedPublicKeys returns the PinnedPublicKeys field if non-nil, zero value otherwise.

### GetPinnedPublicKeysOk

`func (o *APIDefinitionModel) GetPinnedPublicKeysOk() (*map[string]string, bool)`

GetPinnedPublicKeysOk returns a tuple with the PinnedPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedPublicKeys

`func (o *APIDefinitionModel) SetPinnedPublicKeys(v map[string]string)`

SetPinnedPublicKeys sets PinnedPublicKeys field to given value.

### HasPinnedPublicKeys

`func (o *APIDefinitionModel) HasPinnedPublicKeys() bool`

HasPinnedPublicKeys returns a boolean if a field has been set.

### GetProxy

`func (o *APIDefinitionModel) GetProxy() APIDefinitionProxyModelModel`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *APIDefinitionModel) GetProxyOk() (*APIDefinitionProxyModelModel, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *APIDefinitionModel) SetProxy(v APIDefinitionProxyModelModel)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *APIDefinitionModel) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetResponseProcessors

`func (o *APIDefinitionModel) GetResponseProcessors() []ResponseProcessorModelModel`

GetResponseProcessors returns the ResponseProcessors field if non-nil, zero value otherwise.

### GetResponseProcessorsOk

`func (o *APIDefinitionModel) GetResponseProcessorsOk() (*[]ResponseProcessorModelModel, bool)`

GetResponseProcessorsOk returns a tuple with the ResponseProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProcessors

`func (o *APIDefinitionModel) SetResponseProcessors(v []ResponseProcessorModelModel)`

SetResponseProcessors sets ResponseProcessors field to given value.

### HasResponseProcessors

`func (o *APIDefinitionModel) HasResponseProcessors() bool`

HasResponseProcessors returns a boolean if a field has been set.

### GetSessionLifetime

`func (o *APIDefinitionModel) GetSessionLifetime() int64`

GetSessionLifetime returns the SessionLifetime field if non-nil, zero value otherwise.

### GetSessionLifetimeOk

`func (o *APIDefinitionModel) GetSessionLifetimeOk() (*int64, bool)`

GetSessionLifetimeOk returns a tuple with the SessionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetime

`func (o *APIDefinitionModel) SetSessionLifetime(v int64)`

SetSessionLifetime sets SessionLifetime field to given value.

### HasSessionLifetime

`func (o *APIDefinitionModel) HasSessionLifetime() bool`

HasSessionLifetime returns a boolean if a field has been set.

### GetSessionProvider

`func (o *APIDefinitionModel) GetSessionProvider() SessionProviderMetaModelModel`

GetSessionProvider returns the SessionProvider field if non-nil, zero value otherwise.

### GetSessionProviderOk

`func (o *APIDefinitionModel) GetSessionProviderOk() (*SessionProviderMetaModelModel, bool)`

GetSessionProviderOk returns a tuple with the SessionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProvider

`func (o *APIDefinitionModel) SetSessionProvider(v SessionProviderMetaModelModel)`

SetSessionProvider sets SessionProvider field to given value.

### HasSessionProvider

`func (o *APIDefinitionModel) HasSessionProvider() bool`

HasSessionProvider returns a boolean if a field has been set.

### GetSlug

`func (o *APIDefinitionModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *APIDefinitionModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *APIDefinitionModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *APIDefinitionModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStripAuthData

`func (o *APIDefinitionModel) GetStripAuthData() bool`

GetStripAuthData returns the StripAuthData field if non-nil, zero value otherwise.

### GetStripAuthDataOk

`func (o *APIDefinitionModel) GetStripAuthDataOk() (*bool, bool)`

GetStripAuthDataOk returns a tuple with the StripAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripAuthData

`func (o *APIDefinitionModel) SetStripAuthData(v bool)`

SetStripAuthData sets StripAuthData field to given value.

### HasStripAuthData

`func (o *APIDefinitionModel) HasStripAuthData() bool`

HasStripAuthData returns a boolean if a field has been set.

### GetTagHeaders

`func (o *APIDefinitionModel) GetTagHeaders() []string`

GetTagHeaders returns the TagHeaders field if non-nil, zero value otherwise.

### GetTagHeadersOk

`func (o *APIDefinitionModel) GetTagHeadersOk() (*[]string, bool)`

GetTagHeadersOk returns a tuple with the TagHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagHeaders

`func (o *APIDefinitionModel) SetTagHeaders(v []string)`

SetTagHeaders sets TagHeaders field to given value.

### HasTagHeaders

`func (o *APIDefinitionModel) HasTagHeaders() bool`

HasTagHeaders returns a boolean if a field has been set.

### GetUpstreamCertificates

`func (o *APIDefinitionModel) GetUpstreamCertificates() map[string]string`

GetUpstreamCertificates returns the UpstreamCertificates field if non-nil, zero value otherwise.

### GetUpstreamCertificatesOk

`func (o *APIDefinitionModel) GetUpstreamCertificatesOk() (*map[string]string, bool)`

GetUpstreamCertificatesOk returns a tuple with the UpstreamCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertificates

`func (o *APIDefinitionModel) SetUpstreamCertificates(v map[string]string)`

SetUpstreamCertificates sets UpstreamCertificates field to given value.

### HasUpstreamCertificates

`func (o *APIDefinitionModel) HasUpstreamCertificates() bool`

HasUpstreamCertificates returns a boolean if a field has been set.

### GetUptimeTests

`func (o *APIDefinitionModel) GetUptimeTests() APIDefinitionUptimeTestsModelModel`

GetUptimeTests returns the UptimeTests field if non-nil, zero value otherwise.

### GetUptimeTestsOk

`func (o *APIDefinitionModel) GetUptimeTestsOk() (*APIDefinitionUptimeTestsModelModel, bool)`

GetUptimeTestsOk returns a tuple with the UptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeTests

`func (o *APIDefinitionModel) SetUptimeTests(v APIDefinitionUptimeTestsModelModel)`

SetUptimeTests sets UptimeTests field to given value.

### HasUptimeTests

`func (o *APIDefinitionModel) HasUptimeTests() bool`

HasUptimeTests returns a boolean if a field has been set.

### GetUseBasicAuth

`func (o *APIDefinitionModel) GetUseBasicAuth() bool`

GetUseBasicAuth returns the UseBasicAuth field if non-nil, zero value otherwise.

### GetUseBasicAuthOk

`func (o *APIDefinitionModel) GetUseBasicAuthOk() (*bool, bool)`

GetUseBasicAuthOk returns a tuple with the UseBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBasicAuth

`func (o *APIDefinitionModel) SetUseBasicAuth(v bool)`

SetUseBasicAuth sets UseBasicAuth field to given value.

### HasUseBasicAuth

`func (o *APIDefinitionModel) HasUseBasicAuth() bool`

HasUseBasicAuth returns a boolean if a field has been set.

### GetUseKeyless

`func (o *APIDefinitionModel) GetUseKeyless() bool`

GetUseKeyless returns the UseKeyless field if non-nil, zero value otherwise.

### GetUseKeylessOk

`func (o *APIDefinitionModel) GetUseKeylessOk() (*bool, bool)`

GetUseKeylessOk returns a tuple with the UseKeyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKeyless

`func (o *APIDefinitionModel) SetUseKeyless(v bool)`

SetUseKeyless sets UseKeyless field to given value.

### HasUseKeyless

`func (o *APIDefinitionModel) HasUseKeyless() bool`

HasUseKeyless returns a boolean if a field has been set.

### GetUseMutualTlsAuth

`func (o *APIDefinitionModel) GetUseMutualTlsAuth() bool`

GetUseMutualTlsAuth returns the UseMutualTlsAuth field if non-nil, zero value otherwise.

### GetUseMutualTlsAuthOk

`func (o *APIDefinitionModel) GetUseMutualTlsAuthOk() (*bool, bool)`

GetUseMutualTlsAuthOk returns a tuple with the UseMutualTlsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMutualTlsAuth

`func (o *APIDefinitionModel) SetUseMutualTlsAuth(v bool)`

SetUseMutualTlsAuth sets UseMutualTlsAuth field to given value.

### HasUseMutualTlsAuth

`func (o *APIDefinitionModel) HasUseMutualTlsAuth() bool`

HasUseMutualTlsAuth returns a boolean if a field has been set.

### GetUseOauth2

`func (o *APIDefinitionModel) GetUseOauth2() bool`

GetUseOauth2 returns the UseOauth2 field if non-nil, zero value otherwise.

### GetUseOauth2Ok

`func (o *APIDefinitionModel) GetUseOauth2Ok() (*bool, bool)`

GetUseOauth2Ok returns a tuple with the UseOauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOauth2

`func (o *APIDefinitionModel) SetUseOauth2(v bool)`

SetUseOauth2 sets UseOauth2 field to given value.

### HasUseOauth2

`func (o *APIDefinitionModel) HasUseOauth2() bool`

HasUseOauth2 returns a boolean if a field has been set.

### GetUseOpenid

`func (o *APIDefinitionModel) GetUseOpenid() bool`

GetUseOpenid returns the UseOpenid field if non-nil, zero value otherwise.

### GetUseOpenidOk

`func (o *APIDefinitionModel) GetUseOpenidOk() (*bool, bool)`

GetUseOpenidOk returns a tuple with the UseOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOpenid

`func (o *APIDefinitionModel) SetUseOpenid(v bool)`

SetUseOpenid sets UseOpenid field to given value.

### HasUseOpenid

`func (o *APIDefinitionModel) HasUseOpenid() bool`

HasUseOpenid returns a boolean if a field has been set.

### GetUseStandardAuth

`func (o *APIDefinitionModel) GetUseStandardAuth() bool`

GetUseStandardAuth returns the UseStandardAuth field if non-nil, zero value otherwise.

### GetUseStandardAuthOk

`func (o *APIDefinitionModel) GetUseStandardAuthOk() (*bool, bool)`

GetUseStandardAuthOk returns a tuple with the UseStandardAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStandardAuth

`func (o *APIDefinitionModel) SetUseStandardAuth(v bool)`

SetUseStandardAuth sets UseStandardAuth field to given value.

### HasUseStandardAuth

`func (o *APIDefinitionModel) HasUseStandardAuth() bool`

HasUseStandardAuth returns a boolean if a field has been set.

### GetVersionData

`func (o *APIDefinitionModel) GetVersionData() APIDefinitionVersionDataModelModel`

GetVersionData returns the VersionData field if non-nil, zero value otherwise.

### GetVersionDataOk

`func (o *APIDefinitionModel) GetVersionDataOk() (*APIDefinitionVersionDataModelModel, bool)`

GetVersionDataOk returns a tuple with the VersionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionData

`func (o *APIDefinitionModel) SetVersionData(v APIDefinitionVersionDataModelModel)`

SetVersionData sets VersionData field to given value.

### HasVersionData

`func (o *APIDefinitionModel) HasVersionData() bool`

HasVersionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


