/*
 * Tyk Gateway API
 *
 * The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>
 *
 * API version: 4.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gateway

type VirtualMeta struct {
	FunctionSourceType string `json:"function_source_type,omitempty"`
	FunctionSourceUri string `json:"function_source_uri,omitempty"`
	Method string `json:"method,omitempty"`
	Path string `json:"path,omitempty"`
	ProxyOnError bool `json:"proxy_on_error,omitempty"`
	ResponseFunctionName string `json:"response_function_name,omitempty"`
	UseSession bool `json:"use_session,omitempty"`
}
