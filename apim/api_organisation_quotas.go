/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type OrganisationQuotasAPI interface {

	/*
	AddOrgKey Create an organisation key

	This work similar to Keys API except that Key ID is always equals Organisation ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The Key ID
	@return ApiAddOrgKeyRequest
	*/
	AddOrgKey(ctx context.Context, keyID string) ApiAddOrgKeyRequest

	// AddOrgKeyExecute executes the request
	//  @return ApiModifyKeySuccessModelModel
	AddOrgKeyExecute(r ApiAddOrgKeyRequest) (*ApiModifyKeySuccessModelModel, *http.Response, error)

	/*
	DeleteOrgKey Delete Organisation Key

	Deleting a key will remove all limits from organisation. It does not affects regualar keys created within organisation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The Key ID
	@return ApiDeleteOrgKeyRequest
	*/
	DeleteOrgKey(ctx context.Context, keyID string) ApiDeleteOrgKeyRequest

	// DeleteOrgKeyExecute executes the request
	//  @return ApiStatusMessageModelModel
	DeleteOrgKeyExecute(r ApiDeleteOrgKeyRequest) (*ApiStatusMessageModelModel, *http.Response, error)

	/*
	GetOrgKey Get an Organisation Key

	Get session info about specified orgnanisation key. Should return up to date rate limit and quota usage numbers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The Key ID
	@return ApiGetOrgKeyRequest
	*/
	GetOrgKey(ctx context.Context, keyID string) ApiGetOrgKeyRequest

	// GetOrgKeyExecute executes the request
	//  @return SessionStateModelModel
	GetOrgKeyExecute(r ApiGetOrgKeyRequest) (*SessionStateModelModel, *http.Response, error)

	/*
	ListOrgKeys List Organisation Keys

	You can now set rate limits at the organisation level by using the following fields - allowance and rate. These are the number of allowed requests for the specified per value, and need to be set to the same value. If you don't want to have organisation level rate limiting, set 'rate' or 'per' to zero, or don't add them to your request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListOrgKeysRequest
	*/
	ListOrgKeys(ctx context.Context) ApiListOrgKeysRequest

	// ListOrgKeysExecute executes the request
	//  @return ListOrgKeys200ResponseModelModel
	ListOrgKeysExecute(r ApiListOrgKeysRequest) (*ListOrgKeys200ResponseModelModel, *http.Response, error)

	/*
	UpdateOrgKey Update Organisation Key

	This work similar to Keys API except that Key ID is always equals Organisation ID

For Gateway v2.6.0 onwards, you can now set rate limits at the organisation level by using the following fields - allowance and rate. These are the number of allowed requests for the specified per value, and need to be set to the same value. If you don't want to have organisation level rate limiting, set `rate` or `per` to zero, or don't add them to your request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The Key ID
	@return ApiUpdateOrgKeyRequest
	*/
	UpdateOrgKey(ctx context.Context, keyID string) ApiUpdateOrgKeyRequest

	// UpdateOrgKeyExecute executes the request
	//  @return ApiModifyKeySuccessModelModel
	UpdateOrgKeyExecute(r ApiUpdateOrgKeyRequest) (*ApiModifyKeySuccessModelModel, *http.Response, error)
}

// OrganisationQuotasAPIService OrganisationQuotasAPI service
type OrganisationQuotasAPIService service

type ApiAddOrgKeyRequest struct {
	ctx context.Context
	ApiService OrganisationQuotasAPI
	keyID string
	sessionStateModel *SessionStateModel
}

func (r ApiAddOrgKeyRequest) SessionStateModel(sessionStateModel SessionStateModel) ApiAddOrgKeyRequest {
	r.sessionStateModel = &sessionStateModel
	return r
}

func (r ApiAddOrgKeyRequest) Execute() (*ApiModifyKeySuccessModelModel, *http.Response, error) {
	return r.ApiService.AddOrgKeyExecute(r)
}

/*
AddOrgKey Create an organisation key

This work similar to Keys API except that Key ID is always equals Organisation ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiAddOrgKeyRequest
*/
func (a *OrganisationQuotasAPIService) AddOrgKey(ctx context.Context, keyID string) ApiAddOrgKeyRequest {
	return ApiAddOrgKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiModifyKeySuccessModelModel
func (a *OrganisationQuotasAPIService) AddOrgKeyExecute(r ApiAddOrgKeyRequest) (*ApiModifyKeySuccessModelModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiModifyKeySuccessModelModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganisationQuotasAPIService.AddOrgKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/org/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sessionStateModel
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteOrgKeyRequest struct {
	ctx context.Context
	ApiService OrganisationQuotasAPI
	keyID string
}

func (r ApiDeleteOrgKeyRequest) Execute() (*ApiStatusMessageModelModel, *http.Response, error) {
	return r.ApiService.DeleteOrgKeyExecute(r)
}

/*
DeleteOrgKey Delete Organisation Key

Deleting a key will remove all limits from organisation. It does not affects regualar keys created within organisation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiDeleteOrgKeyRequest
*/
func (a *OrganisationQuotasAPIService) DeleteOrgKey(ctx context.Context, keyID string) ApiDeleteOrgKeyRequest {
	return ApiDeleteOrgKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiStatusMessageModelModel
func (a *OrganisationQuotasAPIService) DeleteOrgKeyExecute(r ApiDeleteOrgKeyRequest) (*ApiStatusMessageModelModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusMessageModelModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganisationQuotasAPIService.DeleteOrgKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/org/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrgKeyRequest struct {
	ctx context.Context
	ApiService OrganisationQuotasAPI
	keyID string
}

func (r ApiGetOrgKeyRequest) Execute() (*SessionStateModelModel, *http.Response, error) {
	return r.ApiService.GetOrgKeyExecute(r)
}

/*
GetOrgKey Get an Organisation Key

Get session info about specified orgnanisation key. Should return up to date rate limit and quota usage numbers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiGetOrgKeyRequest
*/
func (a *OrganisationQuotasAPIService) GetOrgKey(ctx context.Context, keyID string) ApiGetOrgKeyRequest {
	return ApiGetOrgKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return SessionStateModelModel
func (a *OrganisationQuotasAPIService) GetOrgKeyExecute(r ApiGetOrgKeyRequest) (*SessionStateModelModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionStateModelModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganisationQuotasAPIService.GetOrgKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/org/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgKeysRequest struct {
	ctx context.Context
	ApiService OrganisationQuotasAPI
}

func (r ApiListOrgKeysRequest) Execute() (*ListOrgKeys200ResponseModelModel, *http.Response, error) {
	return r.ApiService.ListOrgKeysExecute(r)
}

/*
ListOrgKeys List Organisation Keys

You can now set rate limits at the organisation level by using the following fields - allowance and rate. These are the number of allowed requests for the specified per value, and need to be set to the same value. If you don't want to have organisation level rate limiting, set 'rate' or 'per' to zero, or don't add them to your request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOrgKeysRequest
*/
func (a *OrganisationQuotasAPIService) ListOrgKeys(ctx context.Context) ApiListOrgKeysRequest {
	return ApiListOrgKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOrgKeys200ResponseModelModel
func (a *OrganisationQuotasAPIService) ListOrgKeysExecute(r ApiListOrgKeysRequest) (*ListOrgKeys200ResponseModelModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOrgKeys200ResponseModelModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganisationQuotasAPIService.ListOrgKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/org/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateOrgKeyRequest struct {
	ctx context.Context
	ApiService OrganisationQuotasAPI
	keyID string
	resetQuota *string
	sessionStateModel *SessionStateModel
}

// Adding the &#x60;reset_quota&#x60; parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place.
func (r ApiUpdateOrgKeyRequest) ResetQuota(resetQuota string) ApiUpdateOrgKeyRequest {
	r.resetQuota = &resetQuota
	return r
}

func (r ApiUpdateOrgKeyRequest) SessionStateModel(sessionStateModel SessionStateModel) ApiUpdateOrgKeyRequest {
	r.sessionStateModel = &sessionStateModel
	return r
}

func (r ApiUpdateOrgKeyRequest) Execute() (*ApiModifyKeySuccessModelModel, *http.Response, error) {
	return r.ApiService.UpdateOrgKeyExecute(r)
}

/*
UpdateOrgKey Update Organisation Key

This work similar to Keys API except that Key ID is always equals Organisation ID

For Gateway v2.6.0 onwards, you can now set rate limits at the organisation level by using the following fields - allowance and rate. These are the number of allowed requests for the specified per value, and need to be set to the same value. If you don't want to have organisation level rate limiting, set `rate` or `per` to zero, or don't add them to your request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiUpdateOrgKeyRequest
*/
func (a *OrganisationQuotasAPIService) UpdateOrgKey(ctx context.Context, keyID string) ApiUpdateOrgKeyRequest {
	return ApiUpdateOrgKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiModifyKeySuccessModelModel
func (a *OrganisationQuotasAPIService) UpdateOrgKeyExecute(r ApiUpdateOrgKeyRequest) (*ApiModifyKeySuccessModelModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiModifyKeySuccessModelModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganisationQuotasAPIService.UpdateOrgKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/org/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resetQuota != nil {
		localVarQueryParams.Add("reset_quota", parameterToString(*r.resetQuota, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sessionStateModel
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
