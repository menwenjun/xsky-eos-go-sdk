/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DpBlockAsyncReplicationPoliciesAPIService DpBlockAsyncReplicationPoliciesAPI service
type DpBlockAsyncReplicationPoliciesAPIService service

type ApiCreateDpBlockAsyncReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockAsyncReplicationPoliciesAPIService
	body *DpBlockAsyncReplicationPolicyCreateReq
}

// policy info
func (r ApiCreateDpBlockAsyncReplicationPolicyRequest) Body(body DpBlockAsyncReplicationPolicyCreateReq) ApiCreateDpBlockAsyncReplicationPolicyRequest {
	r.body = &body
	return r
}

func (r ApiCreateDpBlockAsyncReplicationPolicyRequest) Execute() (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.CreateDpBlockAsyncReplicationPolicyExecute(r)
}

/*
CreateDpBlockAsyncReplicationPolicy Method for CreateDpBlockAsyncReplicationPolicy

Create dp block async replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDpBlockAsyncReplicationPolicyRequest
*/
func (a *DpBlockAsyncReplicationPoliciesAPIService) CreateDpBlockAsyncReplicationPolicy(ctx context.Context) ApiCreateDpBlockAsyncReplicationPolicyRequest {
	return ApiCreateDpBlockAsyncReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpBlockAsyncReplicationPolicyResp
func (a *DpBlockAsyncReplicationPoliciesAPIService) CreateDpBlockAsyncReplicationPolicyExecute(r ApiCreateDpBlockAsyncReplicationPolicyRequest) (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockAsyncReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockAsyncReplicationPoliciesAPIService.CreateDpBlockAsyncReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-async-replication-policies/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteDpBlockAsyncReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockAsyncReplicationPoliciesAPIService
	policyId int64
	force *bool
}

// force delete or not
func (r ApiDeleteDpBlockAsyncReplicationPolicyRequest) Force(force bool) ApiDeleteDpBlockAsyncReplicationPolicyRequest {
	r.force = &force
	return r
}

func (r ApiDeleteDpBlockAsyncReplicationPolicyRequest) Execute() (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.DeleteDpBlockAsyncReplicationPolicyExecute(r)
}

/*
DeleteDpBlockAsyncReplicationPolicy Method for DeleteDpBlockAsyncReplicationPolicy

Delete dp block async replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiDeleteDpBlockAsyncReplicationPolicyRequest
*/
func (a *DpBlockAsyncReplicationPoliciesAPIService) DeleteDpBlockAsyncReplicationPolicy(ctx context.Context, policyId int64) ApiDeleteDpBlockAsyncReplicationPolicyRequest {
	return ApiDeleteDpBlockAsyncReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockAsyncReplicationPolicyResp
func (a *DpBlockAsyncReplicationPoliciesAPIService) DeleteDpBlockAsyncReplicationPolicyExecute(r ApiDeleteDpBlockAsyncReplicationPolicyRequest) (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockAsyncReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockAsyncReplicationPoliciesAPIService.DeleteDpBlockAsyncReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-async-replication-policies/{policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
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
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetDpBlockAsyncReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockAsyncReplicationPoliciesAPIService
	policyId int64
}

func (r ApiGetDpBlockAsyncReplicationPolicyRequest) Execute() (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.GetDpBlockAsyncReplicationPolicyExecute(r)
}

/*
GetDpBlockAsyncReplicationPolicy Method for GetDpBlockAsyncReplicationPolicy

Get dp block async replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiGetDpBlockAsyncReplicationPolicyRequest
*/
func (a *DpBlockAsyncReplicationPoliciesAPIService) GetDpBlockAsyncReplicationPolicy(ctx context.Context, policyId int64) ApiGetDpBlockAsyncReplicationPolicyRequest {
	return ApiGetDpBlockAsyncReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockAsyncReplicationPolicyResp
func (a *DpBlockAsyncReplicationPoliciesAPIService) GetDpBlockAsyncReplicationPolicyExecute(r ApiGetDpBlockAsyncReplicationPolicyRequest) (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockAsyncReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockAsyncReplicationPoliciesAPIService.GetDpBlockAsyncReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-async-replication-policies/{policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListDpBlockAsyncReplicationPoliciesRequest struct {
	ctx context.Context
	ApiService *DpBlockAsyncReplicationPoliciesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	dpSiteId *int64
	volumeId *int64
}

// paging param
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) Limit(limit int64) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) Offset(offset int64) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) Q(q string) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) Sort(sort string) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.sort = &sort
	return r
}

// related site id
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) DpSiteId(dpSiteId int64) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.dpSiteId = &dpSiteId
	return r
}

// related volume id
func (r ApiListDpBlockAsyncReplicationPoliciesRequest) VolumeId(volumeId int64) ApiListDpBlockAsyncReplicationPoliciesRequest {
	r.volumeId = &volumeId
	return r
}

func (r ApiListDpBlockAsyncReplicationPoliciesRequest) Execute() (*DpBlockAsyncReplicationPoliciesResp, *http.Response, error) {
	return r.ApiService.ListDpBlockAsyncReplicationPoliciesExecute(r)
}

/*
ListDpBlockAsyncReplicationPolicies Method for ListDpBlockAsyncReplicationPolicies

List dp block async replication policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDpBlockAsyncReplicationPoliciesRequest
*/
func (a *DpBlockAsyncReplicationPoliciesAPIService) ListDpBlockAsyncReplicationPolicies(ctx context.Context) ApiListDpBlockAsyncReplicationPoliciesRequest {
	return ApiListDpBlockAsyncReplicationPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpBlockAsyncReplicationPoliciesResp
func (a *DpBlockAsyncReplicationPoliciesAPIService) ListDpBlockAsyncReplicationPoliciesExecute(r ApiListDpBlockAsyncReplicationPoliciesRequest) (*DpBlockAsyncReplicationPoliciesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockAsyncReplicationPoliciesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockAsyncReplicationPoliciesAPIService.ListDpBlockAsyncReplicationPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-async-replication-policies/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.dpSiteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp_site_id", r.dpSiteId, "form", "")
	}
	if r.volumeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume_id", r.volumeId, "form", "")
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
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateDpBlockAsyncReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockAsyncReplicationPoliciesAPIService
	policyId int64
	body *DpBlockAsyncReplicationPolicyUpdateReq
}

// policy info
func (r ApiUpdateDpBlockAsyncReplicationPolicyRequest) Body(body DpBlockAsyncReplicationPolicyUpdateReq) ApiUpdateDpBlockAsyncReplicationPolicyRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDpBlockAsyncReplicationPolicyRequest) Execute() (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.UpdateDpBlockAsyncReplicationPolicyExecute(r)
}

/*
UpdateDpBlockAsyncReplicationPolicy Method for UpdateDpBlockAsyncReplicationPolicy

Update dp block async replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiUpdateDpBlockAsyncReplicationPolicyRequest
*/
func (a *DpBlockAsyncReplicationPoliciesAPIService) UpdateDpBlockAsyncReplicationPolicy(ctx context.Context, policyId int64) ApiUpdateDpBlockAsyncReplicationPolicyRequest {
	return ApiUpdateDpBlockAsyncReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockAsyncReplicationPolicyResp
func (a *DpBlockAsyncReplicationPoliciesAPIService) UpdateDpBlockAsyncReplicationPolicyExecute(r ApiUpdateDpBlockAsyncReplicationPolicyRequest) (*DpBlockAsyncReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockAsyncReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockAsyncReplicationPoliciesAPIService.UpdateDpBlockAsyncReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-async-replication-policies/{policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
