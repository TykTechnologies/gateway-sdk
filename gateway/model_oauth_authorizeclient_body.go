/*
 * Tyk Gateway API
 *
 * The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>
 *
 * API version: 3.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gateway

type OauthAuthorizeclientBody struct {
	// Should be provided by requesting client as part of authorisation request, this should be either `code` or `token` depending on the methods you have specified for the API.
	ResponseType string `json:"response_type,omitempty"`
	// Should be provided by requesting client as part of authorisation request. The Client ID that is making the request.
	ClientId string `json:"client_id,omitempty"`
	// Should be provided by requesting client as part of authorisation request. Must match with the record stored with Tyk.
	RedirectUri string `json:"redirect_uri,omitempty"`
	// A string representation of a Session Object (form-encoded). This should be provided by your application in order to apply any quotas or rules to the key.
	KeyRules string `json:"key_rules,omitempty"`
}
