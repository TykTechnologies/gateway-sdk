/*
 * Tyk Gateway API
 *
 * The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>
 *
 * API version: 4.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gateway

// APILimit stores quota and rate limit on ACL level (per API)
type ApiLimit struct {
	Per                float64 `json:"per,omitempty"`
	QuotaMax           int64   `json:"quota_max,omitempty"`
	QuotaRemaining     int64   `json:"quota_remaining,omitempty"`
	QuotaRenewalRate   int64   `json:"quota_renewal_rate,omitempty"`
	QuotaRenews        int64   `json:"quota_renews,omitempty"`
	Rate               float64 `json:"rate,omitempty"`
	SetByPolicy        bool    `json:"set_by_policy,omitempty"`
	ThrottleInterval   float64 `json:"throttle_interval,omitempty"`
	ThrottleRetryLimit int64   `json:"throttle_retry_limit,omitempty"`
}
