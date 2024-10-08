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


// DpVolumeGroupSnapshotReplicationPoliciesAPIService DpVolumeGroupSnapshotReplicationPoliciesAPI service
type DpVolumeGroupSnapshotReplicationPoliciesAPIService service

type ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpVolumeGroupSnapshotReplicationPoliciesAPIService
	body *DpVolumeGroupSnapshotReplicationPolicyCreateReq
}

// policy info
func (r ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest) Body(body DpVolumeGroupSnapshotReplicationPolicyCreateReq) ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest {
	r.body = &body
	return r
}

func (r ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest) Execute() (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.CreateDpVolumeGroupSnapshotReplicationPolicyExecute(r)
}

/*
CreateDpVolumeGroupSnapshotReplicationPolicy Method for CreateDpVolumeGroupSnapshotReplicationPolicy

Create dp volume group snapshot replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest
*/
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) CreateDpVolumeGroupSnapshotReplicationPolicy(ctx context.Context) ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest {
	return ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpVolumeGroupSnapshotReplicationPolicyResp
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) CreateDpVolumeGroupSnapshotReplicationPolicyExecute(r ApiCreateDpVolumeGroupSnapshotReplicationPolicyRequest) (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpVolumeGroupSnapshotReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpVolumeGroupSnapshotReplicationPoliciesAPIService.CreateDpVolumeGroupSnapshotReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-volume-group-snapshot-replication-policies/"

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

type ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpVolumeGroupSnapshotReplicationPoliciesAPIService
	policyId int64
	force *bool
}

// force delete or not
func (r ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest) Force(force bool) ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest {
	r.force = &force
	return r
}

func (r ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest) Execute() (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.DeleteDpVolumeGroupSnapshotReplicationPolicyExecute(r)
}

/*
DeleteDpVolumeGroupSnapshotReplicationPolicy Method for DeleteDpVolumeGroupSnapshotReplicationPolicy

Delete dp volume group snapshot replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest
*/
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) DeleteDpVolumeGroupSnapshotReplicationPolicy(ctx context.Context, policyId int64) ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest {
	return ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpVolumeGroupSnapshotReplicationPolicyResp
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) DeleteDpVolumeGroupSnapshotReplicationPolicyExecute(r ApiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest) (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpVolumeGroupSnapshotReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpVolumeGroupSnapshotReplicationPoliciesAPIService.DeleteDpVolumeGroupSnapshotReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-volume-group-snapshot-replication-policies/{policy_id}"
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

type ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpVolumeGroupSnapshotReplicationPoliciesAPIService
	policyId int64
}

func (r ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest) Execute() (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.GetDpVolumeGroupSnapshotReplicationPolicyExecute(r)
}

/*
GetDpVolumeGroupSnapshotReplicationPolicy Method for GetDpVolumeGroupSnapshotReplicationPolicy

Get dp volume group snapshot replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest
*/
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) GetDpVolumeGroupSnapshotReplicationPolicy(ctx context.Context, policyId int64) ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest {
	return ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpVolumeGroupSnapshotReplicationPolicyResp
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) GetDpVolumeGroupSnapshotReplicationPolicyExecute(r ApiGetDpVolumeGroupSnapshotReplicationPolicyRequest) (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpVolumeGroupSnapshotReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpVolumeGroupSnapshotReplicationPoliciesAPIService.GetDpVolumeGroupSnapshotReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-volume-group-snapshot-replication-policies/{policy_id}"
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

type ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest struct {
	ctx context.Context
	ApiService *DpVolumeGroupSnapshotReplicationPoliciesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	dpSiteId *int64
	volumeGroupId *int64
}

// paging param
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) Limit(limit int64) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) Offset(offset int64) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) Q(q string) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) Sort(sort string) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.sort = &sort
	return r
}

// related site id
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) DpSiteId(dpSiteId int64) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.dpSiteId = &dpSiteId
	return r
}

// related volume group id
func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) VolumeGroupId(volumeGroupId int64) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	r.volumeGroupId = &volumeGroupId
	return r
}

func (r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) Execute() (*DpVolumeGroupSnapshotReplicationPoliciesResp, *http.Response, error) {
	return r.ApiService.ListDpVolumeGroupSnapshotReplicationPoliciesExecute(r)
}

/*
ListDpVolumeGroupSnapshotReplicationPolicies Method for ListDpVolumeGroupSnapshotReplicationPolicies

List dp volume group snapshot replication policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest
*/
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) ListDpVolumeGroupSnapshotReplicationPolicies(ctx context.Context) ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest {
	return ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DpVolumeGroupSnapshotReplicationPoliciesResp
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) ListDpVolumeGroupSnapshotReplicationPoliciesExecute(r ApiListDpVolumeGroupSnapshotReplicationPoliciesRequest) (*DpVolumeGroupSnapshotReplicationPoliciesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpVolumeGroupSnapshotReplicationPoliciesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpVolumeGroupSnapshotReplicationPoliciesAPIService.ListDpVolumeGroupSnapshotReplicationPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-volume-group-snapshot-replication-policies/"

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
	if r.volumeGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume_group_id", r.volumeGroupId, "form", "")
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

type ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest struct {
	ctx context.Context
	ApiService *DpVolumeGroupSnapshotReplicationPoliciesAPIService
	policyId int64
	body *DpVolumeGroupSnapshotReplicationPolicyUpdateReq
}

// policy info
func (r ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest) Body(body DpVolumeGroupSnapshotReplicationPolicyUpdateReq) ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest) Execute() (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	return r.ApiService.UpdateDpVolumeGroupSnapshotReplicationPolicyExecute(r)
}

/*
UpdateDpVolumeGroupSnapshotReplicationPolicy Method for UpdateDpVolumeGroupSnapshotReplicationPolicy

Update dp volume group snapshot replication policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId resource id
 @return ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest
*/
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) UpdateDpVolumeGroupSnapshotReplicationPolicy(ctx context.Context, policyId int64) ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest {
	return ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return DpVolumeGroupSnapshotReplicationPolicyResp
func (a *DpVolumeGroupSnapshotReplicationPoliciesAPIService) UpdateDpVolumeGroupSnapshotReplicationPolicyExecute(r ApiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest) (*DpVolumeGroupSnapshotReplicationPolicyResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DpVolumeGroupSnapshotReplicationPolicyResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DpVolumeGroupSnapshotReplicationPoliciesAPIService.UpdateDpVolumeGroupSnapshotReplicationPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dp-volume-group-snapshot-replication-policies/{policy_id}"
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
