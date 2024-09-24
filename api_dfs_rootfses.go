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


// DfsRootfsesAPIService DfsRootfsesAPI service
type DfsRootfsesAPIService service

type ApiCreateDfsRootfsRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	body *DfsRootfsCreateReq
}

// rootfs info
func (r ApiCreateDfsRootfsRequest) Body(body DfsRootfsCreateReq) ApiCreateDfsRootfsRequest {
	r.body = &body
	return r
}

func (r ApiCreateDfsRootfsRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.CreateDfsRootfsExecute(r)
}

/*
CreateDfsRootfs Method for CreateDfsRootfs

Create dfs rootfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDfsRootfsRequest
*/
func (a *DfsRootfsesAPIService) CreateDfsRootfs(ctx context.Context) ApiCreateDfsRootfsRequest {
	return ApiCreateDfsRootfsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) CreateDfsRootfsExecute(r ApiCreateDfsRootfsRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.CreateDfsRootfs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/"

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

type ApiDeleteDfsRootfsRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
	force *bool
}

// force delete
func (r ApiDeleteDfsRootfsRequest) Force(force bool) ApiDeleteDfsRootfsRequest {
	r.force = &force
	return r
}

func (r ApiDeleteDfsRootfsRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.DeleteDfsRootfsExecute(r)
}

/*
DeleteDfsRootfs Method for DeleteDfsRootfs

delete dfs rootfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId rootfs id
 @return ApiDeleteDfsRootfsRequest
*/
func (a *DfsRootfsesAPIService) DeleteDfsRootfs(ctx context.Context, dfsRootfsId int64) ApiDeleteDfsRootfsRequest {
	return ApiDeleteDfsRootfsRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) DeleteDfsRootfsExecute(r ApiDeleteDfsRootfsRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.DeleteDfsRootfs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

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

type ApiGetDfsRootfsRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
}

func (r ApiGetDfsRootfsRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.GetDfsRootfsExecute(r)
}

/*
GetDfsRootfs Method for GetDfsRootfs

Get dfs rootfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId rootfs id
 @return ApiGetDfsRootfsRequest
*/
func (a *DfsRootfsesAPIService) GetDfsRootfs(ctx context.Context, dfsRootfsId int64) ApiGetDfsRootfsRequest {
	return ApiGetDfsRootfsRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) GetDfsRootfsExecute(r ApiGetDfsRootfsRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.GetDfsRootfs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

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

type ApiGetDfsRootfsSamplesRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
	durationBegin *string
	durationEnd *string
	period *string
}

// duration begin timestamp
func (r ApiGetDfsRootfsSamplesRequest) DurationBegin(durationBegin string) ApiGetDfsRootfsSamplesRequest {
	r.durationBegin = &durationBegin
	return r
}

// duration end timestamp
func (r ApiGetDfsRootfsSamplesRequest) DurationEnd(durationEnd string) ApiGetDfsRootfsSamplesRequest {
	r.durationEnd = &durationEnd
	return r
}

// samples period
func (r ApiGetDfsRootfsSamplesRequest) Period(period string) ApiGetDfsRootfsSamplesRequest {
	r.period = &period
	return r
}

func (r ApiGetDfsRootfsSamplesRequest) Execute() (*DfsRootfsSamplesResp, *http.Response, error) {
	return r.ApiService.GetDfsRootfsSamplesExecute(r)
}

/*
GetDfsRootfsSamples Method for GetDfsRootfsSamples

get a dfs rootfs's samples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId dfs rootfs id
 @return ApiGetDfsRootfsSamplesRequest
*/
func (a *DfsRootfsesAPIService) GetDfsRootfsSamples(ctx context.Context, dfsRootfsId int64) ApiGetDfsRootfsSamplesRequest {
	return ApiGetDfsRootfsSamplesRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsSamplesResp
func (a *DfsRootfsesAPIService) GetDfsRootfsSamplesExecute(r ApiGetDfsRootfsSamplesRequest) (*DfsRootfsSamplesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsSamplesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.GetDfsRootfsSamples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}/samples"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.durationBegin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_begin", r.durationBegin, "form", "")
	}
	if r.durationEnd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_end", r.durationEnd, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
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

type ApiListDfsRootfsesRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	clusterId *string
	poolId *int64
	fsUserId *int64
	fsUserGroupId *int64
}

// paging param
func (r ApiListDfsRootfsesRequest) Limit(limit int64) ApiListDfsRootfsesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDfsRootfsesRequest) Offset(offset int64) ApiListDfsRootfsesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListDfsRootfsesRequest) Q(q string) ApiListDfsRootfsesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDfsRootfsesRequest) Sort(sort string) ApiListDfsRootfsesRequest {
	r.sort = &sort
	return r
}

// cluster id
func (r ApiListDfsRootfsesRequest) ClusterId(clusterId string) ApiListDfsRootfsesRequest {
	r.clusterId = &clusterId
	return r
}

// pool id
func (r ApiListDfsRootfsesRequest) PoolId(poolId int64) ApiListDfsRootfsesRequest {
	r.poolId = &poolId
	return r
}

// fs user id
func (r ApiListDfsRootfsesRequest) FsUserId(fsUserId int64) ApiListDfsRootfsesRequest {
	r.fsUserId = &fsUserId
	return r
}

// fs user group id
func (r ApiListDfsRootfsesRequest) FsUserGroupId(fsUserGroupId int64) ApiListDfsRootfsesRequest {
	r.fsUserGroupId = &fsUserGroupId
	return r
}

func (r ApiListDfsRootfsesRequest) Execute() (*DfsRootfsesResp, *http.Response, error) {
	return r.ApiService.ListDfsRootfsesExecute(r)
}

/*
ListDfsRootfses Method for ListDfsRootfses

List dfs rootfses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDfsRootfsesRequest
*/
func (a *DfsRootfsesAPIService) ListDfsRootfses(ctx context.Context) ApiListDfsRootfsesRequest {
	return ApiListDfsRootfsesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsRootfsesResp
func (a *DfsRootfsesAPIService) ListDfsRootfsesExecute(r ApiListDfsRootfsesRequest) (*DfsRootfsesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.ListDfsRootfses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/"

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
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
	}
	if r.poolId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pool_id", r.poolId, "form", "")
	}
	if r.fsUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fs_user_id", r.fsUserId, "form", "")
	}
	if r.fsUserGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fs_user_group_id", r.fsUserGroupId, "form", "")
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

type ApiSetDfsWormLogPathRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
	body *DfsRootfsSetWormLogPathReq
	allowPathCreate *bool
}

// worm log path
func (r ApiSetDfsWormLogPathRequest) Body(body DfsRootfsSetWormLogPathReq) ApiSetDfsWormLogPathRequest {
	r.body = &body
	return r
}

// allow create path when not existed
func (r ApiSetDfsWormLogPathRequest) AllowPathCreate(allowPathCreate bool) ApiSetDfsWormLogPathRequest {
	r.allowPathCreate = &allowPathCreate
	return r
}

func (r ApiSetDfsWormLogPathRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.SetDfsWormLogPathExecute(r)
}

/*
SetDfsWormLogPath Method for SetDfsWormLogPath

Set dfs worm log path

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId rootfs id
 @return ApiSetDfsWormLogPathRequest
*/
func (a *DfsRootfsesAPIService) SetDfsWormLogPath(ctx context.Context, dfsRootfsId int64) ApiSetDfsWormLogPathRequest {
	return ApiSetDfsWormLogPathRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) SetDfsWormLogPathExecute(r ApiSetDfsWormLogPathRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.SetDfsWormLogPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}:set-worm-log-path"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.allowPathCreate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "allow_path_create", r.allowPathCreate, "form", "")
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

type ApiSetGCSpeedRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
	body *DfsRootfsSetGCSpeedReq
}

// gc speed
func (r ApiSetGCSpeedRequest) Body(body DfsRootfsSetGCSpeedReq) ApiSetGCSpeedRequest {
	r.body = &body
	return r
}

func (r ApiSetGCSpeedRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.SetGCSpeedExecute(r)
}

/*
SetGCSpeed Method for SetGCSpeed

Set gc speed of dfs rootfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId rootfs id
 @return ApiSetGCSpeedRequest
*/
func (a *DfsRootfsesAPIService) SetGCSpeed(ctx context.Context, dfsRootfsId int64) ApiSetGCSpeedRequest {
	return ApiSetGCSpeedRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) SetGCSpeedExecute(r ApiSetGCSpeedRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.SetGCSpeed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}:set-gc-speed"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

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

type ApiUpdateDfsRootfsActivePoolRequest struct {
	ctx context.Context
	ApiService *DfsRootfsesAPIService
	dfsRootfsId int64
	body *DfsRootfsUpdateActivePoolReq
}

// active pool ids
func (r ApiUpdateDfsRootfsActivePoolRequest) Body(body DfsRootfsUpdateActivePoolReq) ApiUpdateDfsRootfsActivePoolRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDfsRootfsActivePoolRequest) Execute() (*DfsRootfsResp, *http.Response, error) {
	return r.ApiService.UpdateDfsRootfsActivePoolExecute(r)
}

/*
UpdateDfsRootfsActivePool Method for UpdateDfsRootfsActivePool

Update dfs rootfs active pools

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsRootfsId rootfs id
 @return ApiUpdateDfsRootfsActivePoolRequest
*/
func (a *DfsRootfsesAPIService) UpdateDfsRootfsActivePool(ctx context.Context, dfsRootfsId int64) ApiUpdateDfsRootfsActivePoolRequest {
	return ApiUpdateDfsRootfsActivePoolRequest{
		ApiService: a,
		ctx: ctx,
		dfsRootfsId: dfsRootfsId,
	}
}

// Execute executes the request
//  @return DfsRootfsResp
func (a *DfsRootfsesAPIService) UpdateDfsRootfsActivePoolExecute(r ApiUpdateDfsRootfsActivePoolRequest) (*DfsRootfsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsRootfsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsRootfsesAPIService.UpdateDfsRootfsActivePool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-rootfses/{dfs_rootfs_id}:update-active-pools"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_rootfs_id"+"}", url.PathEscape(parameterValueToString(r.dfsRootfsId, "dfsRootfsId")), -1)

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
