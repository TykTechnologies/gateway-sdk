/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gate

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type CertsApi interface {

	/*
	AddCert Add a certificate

	Add a certificate to the Tyk Gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCertRequest
	*/
	AddCert(ctx context.Context) ApiAddCertRequest

	// AddCertExecute executes the request
	//  @return APICertificateStatusMessage
	AddCertExecute(r ApiAddCertRequest) (*APICertificateStatusMessage, *http.Response, error)

	/*
	DeleteCerts Delete Certificate

	Delete certificate by id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteCertsRequest
	*/
	DeleteCerts(ctx context.Context) ApiDeleteCertsRequest

	// DeleteCertsExecute executes the request
	//  @return ApiStatusMessage
	DeleteCertsExecute(r ApiDeleteCertsRequest) (*ApiStatusMessage, *http.Response, error)

	/*
	ListCerts List Certificates

	List All Certificates in the Tyk Gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCertsRequest
	*/
	ListCerts(ctx context.Context) ApiListCertsRequest

	// ListCertsExecute executes the request
	//  @return ListCerts200Response
	ListCertsExecute(r ApiListCertsRequest) (*ListCerts200Response, *http.Response, error)
}

// CertsApiService CertsApi service
type CertsApiService service

type ApiAddCertRequest struct {
	ctx context.Context
	ApiService CertsApi
	orgId *string
	body *string
}

// Organisation ID to list the certificates
func (r ApiAddCertRequest) OrgId(orgId string) ApiAddCertRequest {
	r.orgId = &orgId
	return r
}

func (r ApiAddCertRequest) Body(body string) ApiAddCertRequest {
	r.body = &body
	return r
}

func (r ApiAddCertRequest) Execute() (*APICertificateStatusMessage, *http.Response, error) {
	return r.ApiService.AddCertExecute(r)
}

/*
AddCert Add a certificate

Add a certificate to the Tyk Gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddCertRequest
*/
func (a *CertsApiService) AddCert(ctx context.Context) ApiAddCertRequest {
	return ApiAddCertRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return APICertificateStatusMessage
func (a *CertsApiService) AddCertExecute(r ApiAddCertRequest) (*APICertificateStatusMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *APICertificateStatusMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertsApiService.AddCert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/certs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	localVarQueryParams.Add("org_id", parameterToString(*r.orgId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiDeleteCertsRequest struct {
	ctx context.Context
	ApiService CertsApi
	certID *string
	orgId *string
}

// Certifiicate ID to be deleted
func (r ApiDeleteCertsRequest) CertID(certID string) ApiDeleteCertsRequest {
	r.certID = &certID
	return r
}

// Organisation ID to list the certificates
func (r ApiDeleteCertsRequest) OrgId(orgId string) ApiDeleteCertsRequest {
	r.orgId = &orgId
	return r
}

func (r ApiDeleteCertsRequest) Execute() (*ApiStatusMessage, *http.Response, error) {
	return r.ApiService.DeleteCertsExecute(r)
}

/*
DeleteCerts Delete Certificate

Delete certificate by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteCertsRequest
*/
func (a *CertsApiService) DeleteCerts(ctx context.Context) ApiDeleteCertsRequest {
	return ApiDeleteCertsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStatusMessage
func (a *CertsApiService) DeleteCertsExecute(r ApiDeleteCertsRequest) (*ApiStatusMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertsApiService.DeleteCerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/certs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certID == nil {
		return localVarReturnValue, nil, reportError("certID is required and must be specified")
	}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	localVarQueryParams.Add("certID", parameterToString(*r.certID, ""))
	localVarQueryParams.Add("org_id", parameterToString(*r.orgId, ""))
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

type ApiListCertsRequest struct {
	ctx context.Context
	ApiService CertsApi
	orgId *string
	mode *string
	certID *string
}

// Organisation ID to list the certificates
func (r ApiListCertsRequest) OrgId(orgId string) ApiListCertsRequest {
	r.orgId = &orgId
	return r
}

// Mode to list the certificate details
func (r ApiListCertsRequest) Mode(mode string) ApiListCertsRequest {
	r.mode = &mode
	return r
}

// Comma separated list of certificates to list
func (r ApiListCertsRequest) CertID(certID string) ApiListCertsRequest {
	r.certID = &certID
	return r
}

func (r ApiListCertsRequest) Execute() (*ListCerts200Response, *http.Response, error) {
	return r.ApiService.ListCertsExecute(r)
}

/*
ListCerts List Certificates

List All Certificates in the Tyk Gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCertsRequest
*/
func (a *CertsApiService) ListCerts(ctx context.Context) ApiListCertsRequest {
	return ApiListCertsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCerts200Response
func (a *CertsApiService) ListCertsExecute(r ApiListCertsRequest) (*ListCerts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCerts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertsApiService.ListCerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/certs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	localVarQueryParams.Add("org_id", parameterToString(*r.orgId, ""))
	if r.mode != nil {
		localVarQueryParams.Add("mode", parameterToString(*r.mode, ""))
	}
	if r.certID != nil {
		localVarQueryParams.Add("certID", parameterToString(*r.certID, ""))
	}
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
