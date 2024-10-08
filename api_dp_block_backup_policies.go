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


// DpBlockBackupPoliciesAPIService DpBlockBackupPoliciesAPI service
type DpBlockBackupPoliciesAPIService service

type ApiCreateDpBlockBackupPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockBackupPoliciesAPIService
	body *DpBlockBackupPolicyCreateReq
}

// policy info
func (r ApiCreateDpBlockBackupPolicyRequest) Body(body DpBlockBackupPolicyCreateReq) ApiCreateDpBlockBackupPolicyRequest {
	r.body = &body
	return r
}

func (r ApiCreateDpBlockBackupPolicyRequest) Execute() (*DpBlockBackupPolicyResp, *http.Response, error) {
	return r.ApiService.CreateDpBlockBackupPolicyExecute(r)
}

/*
CreateDpBlockBackupPolicy Method for CreateDpBlockBackupPolicy

Create dp block backup policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDpBlockBackupPolicyRequest
*/
func (a *DpBlockBackupPoliciesAPIService) CreateDpBlockBackupPolicy(ctx context.Context) ApiCreateDpBlockBackupPolicyRequest {
	return ApiCreateDpBlockBackupPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpBlockBackupPolicyResp
func (a *DpBlockBackupPoliciesAPIService) CreateDpBlockBackupPolicyExecute(r ApiCreateDpBlockBackupPolicyRequest) (*DpBlockBackupPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockBackupPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockBackupPoliciesAPIService.CreateDpBlockBackupPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-backup-policies/"

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

type ApiDeleteDpBlockBackupPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockBackupPoliciesAPIService
	policyId int64
	force *bool
}

// force delete or not
func (r ApiDeleteDpBlockBackupPolicyRequest) Force(force bool) ApiDeleteDpBlockBackupPolicyRequest {
	r.force = &force
	return r
}

func (r ApiDeleteDpBlockBackupPolicyRequest) Execute() (*DpBlockBackupPolicyResp, *http.Response, error) {
	return r.ApiService.DeleteDpBlockBackupPolicyExecute(r)
}

/*
DeleteDpBlockBackupPolicy Method for DeleteDpBlockBackupPolicy

Delete dp block backup policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiDeleteDpBlockBackupPolicyRequest
*/
func (a *DpBlockBackupPoliciesAPIService) DeleteDpBlockBackupPolicy(ctx context.Context, policyId int64) ApiDeleteDpBlockBackupPolicyRequest {
	return ApiDeleteDpBlockBackupPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockBackupPolicyResp
func (a *DpBlockBackupPoliciesAPIService) DeleteDpBlockBackupPolicyExecute(r ApiDeleteDpBlockBackupPolicyRequest) (*DpBlockBackupPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockBackupPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockBackupPoliciesAPIService.DeleteDpBlockBackupPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-backup-policies/{policy_id}"
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

type ApiGetDpBlockBackupPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockBackupPoliciesAPIService
	policyId int64
}

func (r ApiGetDpBlockBackupPolicyRequest) Execute() (*DpBlockBackupPolicyResp, *http.Response, error) {
	return r.ApiService.GetDpBlockBackupPolicyExecute(r)
}

/*
GetDpBlockBackupPolicy Method for GetDpBlockBackupPolicy

Get dp block backup policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiGetDpBlockBackupPolicyRequest
*/
func (a *DpBlockBackupPoliciesAPIService) GetDpBlockBackupPolicy(ctx context.Context, policyId int64) ApiGetDpBlockBackupPolicyRequest {
	return ApiGetDpBlockBackupPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockBackupPolicyResp
func (a *DpBlockBackupPoliciesAPIService) GetDpBlockBackupPolicyExecute(r ApiGetDpBlockBackupPolicyRequest) (*DpBlockBackupPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockBackupPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockBackupPoliciesAPIService.GetDpBlockBackupPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-backup-policies/{policy_id}"
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

type ApiListDpBlockBackupPoliciesRequest struct {
	ctx context.Context
	ApiService *DpBlockBackupPoliciesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	blockVolumeId *int64
	dpSiteId *int64
}

// paging param
func (r ApiListDpBlockBackupPoliciesRequest) Limit(limit int64) ApiListDpBlockBackupPoliciesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDpBlockBackupPoliciesRequest) Offset(offset int64) ApiListDpBlockBackupPoliciesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListDpBlockBackupPoliciesRequest) Q(q string) ApiListDpBlockBackupPoliciesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDpBlockBackupPoliciesRequest) Sort(sort string) ApiListDpBlockBackupPoliciesRequest {
	r.sort = &sort
	return r
}

// show dp block backup policies of specific block volume
func (r ApiListDpBlockBackupPoliciesRequest) BlockVolumeId(blockVolumeId int64) ApiListDpBlockBackupPoliciesRequest {
	r.blockVolumeId = &blockVolumeId
	return r
}

// related site id
func (r ApiListDpBlockBackupPoliciesRequest) DpSiteId(dpSiteId int64) ApiListDpBlockBackupPoliciesRequest {
	r.dpSiteId = &dpSiteId
	return r
}

func (r ApiListDpBlockBackupPoliciesRequest) Execute() (*DpBlockBackupPoliciesResp, *http.Response, error) {
	return r.ApiService.ListDpBlockBackupPoliciesExecute(r)
}

/*
ListDpBlockBackupPolicies Method for ListDpBlockBackupPolicies

List dp block backup policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDpBlockBackupPoliciesRequest
*/
func (a *DpBlockBackupPoliciesAPIService) ListDpBlockBackupPolicies(ctx context.Context) ApiListDpBlockBackupPoliciesRequest {
	return ApiListDpBlockBackupPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpBlockBackupPoliciesResp
func (a *DpBlockBackupPoliciesAPIService) ListDpBlockBackupPoliciesExecute(r ApiListDpBlockBackupPoliciesRequest) (*DpBlockBackupPoliciesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockBackupPoliciesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockBackupPoliciesAPIService.ListDpBlockBackupPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-backup-policies/"

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
	if r.blockVolumeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "block_volume_id", r.blockVolumeId, "form", "")
	}
	if r.dpSiteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp_site_id", r.dpSiteId, "form", "")
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

type ApiUpdateDpBlockBackupPolicyRequest struct {
	ctx context.Context
	ApiService *DpBlockBackupPoliciesAPIService
	policyId int64
	body *DpBlockBackupPolicyUpdateReq
}

// policy info
func (r ApiUpdateDpBlockBackupPolicyRequest) Body(body DpBlockBackupPolicyUpdateReq) ApiUpdateDpBlockBackupPolicyRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDpBlockBackupPolicyRequest) Execute() (*DpBlockBackupPolicyResp, *http.Response, error) {
	return r.ApiService.UpdateDpBlockBackupPolicyExecute(r)
}

/*
UpdateDpBlockBackupPolicy Method for UpdateDpBlockBackupPolicy

Update dp block backup policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiUpdateDpBlockBackupPolicyRequest
*/
func (a *DpBlockBackupPoliciesAPIService) UpdateDpBlockBackupPolicy(ctx context.Context, policyId int64) ApiUpdateDpBlockBackupPolicyRequest {
	return ApiUpdateDpBlockBackupPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpBlockBackupPolicyResp
func (a *DpBlockBackupPoliciesAPIService) UpdateDpBlockBackupPolicyExecute(r ApiUpdateDpBlockBackupPolicyRequest) (*DpBlockBackupPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpBlockBackupPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpBlockBackupPoliciesAPIService.UpdateDpBlockBackupPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-block-backup-policies/{policy_id}"
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
