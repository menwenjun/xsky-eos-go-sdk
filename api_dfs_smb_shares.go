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


// DfsSmbSharesAPIService DfsSmbSharesAPI service
type DfsSmbSharesAPIService service

type ApiAddDfsSMBShareACLsRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
	body *DfsSMBShareAddACLsReq
}

// share acls info
func (r ApiAddDfsSMBShareACLsRequest) Body(body DfsSMBShareAddACLsReq) ApiAddDfsSMBShareACLsRequest {
	r.body = &body
	return r
}

func (r ApiAddDfsSMBShareACLsRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.AddDfsSMBShareACLsExecute(r)
}

/*
AddDfsSMBShareACLs Method for AddDfsSMBShareACLs

Add dfs smb share acls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId dfs smb share id
 @return ApiAddDfsSMBShareACLsRequest
*/
func (a *DfsSmbSharesAPIService) AddDfsSMBShareACLs(ctx context.Context, dfsSmbShareId int64) ApiAddDfsSMBShareACLsRequest {
	return ApiAddDfsSMBShareACLsRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) AddDfsSMBShareACLsExecute(r ApiAddDfsSMBShareACLsRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.AddDfsSMBShareACLs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}:add-acls"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

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

type ApiCreateDfsSMBShareRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	body *DfsSMBShareCreateReq
	allowPathCreate *bool
}

// share info
func (r ApiCreateDfsSMBShareRequest) Body(body DfsSMBShareCreateReq) ApiCreateDfsSMBShareRequest {
	r.body = &body
	return r
}

// allow create path when not existed
func (r ApiCreateDfsSMBShareRequest) AllowPathCreate(allowPathCreate bool) ApiCreateDfsSMBShareRequest {
	r.allowPathCreate = &allowPathCreate
	return r
}

func (r ApiCreateDfsSMBShareRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.CreateDfsSMBShareExecute(r)
}

/*
CreateDfsSMBShare Method for CreateDfsSMBShare

Create dfs smb share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDfsSMBShareRequest
*/
func (a *DfsSmbSharesAPIService) CreateDfsSMBShare(ctx context.Context) ApiCreateDfsSMBShareRequest {
	return ApiCreateDfsSMBShareRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) CreateDfsSMBShareExecute(r ApiCreateDfsSMBShareRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.CreateDfsSMBShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/"

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

type ApiDeleteDfsSMBShareRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
	force *bool
	withDirectory *bool
}

// force delete or not
func (r ApiDeleteDfsSMBShareRequest) Force(force bool) ApiDeleteDfsSMBShareRequest {
	r.force = &force
	return r
}

// also delete directory
func (r ApiDeleteDfsSMBShareRequest) WithDirectory(withDirectory bool) ApiDeleteDfsSMBShareRequest {
	r.withDirectory = &withDirectory
	return r
}

func (r ApiDeleteDfsSMBShareRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.DeleteDfsSMBShareExecute(r)
}

/*
DeleteDfsSMBShare Method for DeleteDfsSMBShare

delete dfs smb share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId share id
 @return ApiDeleteDfsSMBShareRequest
*/
func (a *DfsSmbSharesAPIService) DeleteDfsSMBShare(ctx context.Context, dfsSmbShareId int64) ApiDeleteDfsSMBShareRequest {
	return ApiDeleteDfsSMBShareRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) DeleteDfsSMBShareExecute(r ApiDeleteDfsSMBShareRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.DeleteDfsSMBShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
	}
	if r.withDirectory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_directory", r.withDirectory, "form", "")
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

type ApiGetDfsSMBShareRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
}

func (r ApiGetDfsSMBShareRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.GetDfsSMBShareExecute(r)
}

/*
GetDfsSMBShare Method for GetDfsSMBShare

Get dfs smb share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId share id
 @return ApiGetDfsSMBShareRequest
*/
func (a *DfsSmbSharesAPIService) GetDfsSMBShare(ctx context.Context, dfsSmbShareId int64) ApiGetDfsSMBShareRequest {
	return ApiGetDfsSMBShareRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) GetDfsSMBShareExecute(r ApiGetDfsSMBShareRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.GetDfsSMBShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

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

type ApiListDfsSMBSharesRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	limit *int64
	offset *int64
	clusterId *string
	dfsRootfsId *int64
	dfsGatewayGroupId *int64
	q *string
	path *string
	sort *string
}

// paging param
func (r ApiListDfsSMBSharesRequest) Limit(limit int64) ApiListDfsSMBSharesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDfsSMBSharesRequest) Offset(offset int64) ApiListDfsSMBSharesRequest {
	r.offset = &offset
	return r
}

// cluster id
func (r ApiListDfsSMBSharesRequest) ClusterId(clusterId string) ApiListDfsSMBSharesRequest {
	r.clusterId = &clusterId
	return r
}

// dfs rootfs id
func (r ApiListDfsSMBSharesRequest) DfsRootfsId(dfsRootfsId int64) ApiListDfsSMBSharesRequest {
	r.dfsRootfsId = &dfsRootfsId
	return r
}

// dfs gateway group id
func (r ApiListDfsSMBSharesRequest) DfsGatewayGroupId(dfsGatewayGroupId int64) ApiListDfsSMBSharesRequest {
	r.dfsGatewayGroupId = &dfsGatewayGroupId
	return r
}

// query param of search
func (r ApiListDfsSMBSharesRequest) Q(q string) ApiListDfsSMBSharesRequest {
	r.q = &q
	return r
}

// related dfs path
func (r ApiListDfsSMBSharesRequest) Path(path string) ApiListDfsSMBSharesRequest {
	r.path = &path
	return r
}

// sort param of search
func (r ApiListDfsSMBSharesRequest) Sort(sort string) ApiListDfsSMBSharesRequest {
	r.sort = &sort
	return r
}

func (r ApiListDfsSMBSharesRequest) Execute() (*DfsSMBSharesResp, *http.Response, error) {
	return r.ApiService.ListDfsSMBSharesExecute(r)
}

/*
ListDfsSMBShares Method for ListDfsSMBShares

List dfs smb shares

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDfsSMBSharesRequest
*/
func (a *DfsSmbSharesAPIService) ListDfsSMBShares(ctx context.Context) ApiListDfsSMBSharesRequest {
	return ApiListDfsSMBSharesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsSMBSharesResp
func (a *DfsSmbSharesAPIService) ListDfsSMBSharesExecute(r ApiListDfsSMBSharesRequest) (*DfsSMBSharesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBSharesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.ListDfsSMBShares")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
	}
	if r.dfsRootfsId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_rootfs_id", r.dfsRootfsId, "form", "")
	}
	if r.dfsGatewayGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_group_id", r.dfsGatewayGroupId, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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

type ApiRemoveDfsSMBShareACLsRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
	body *DfsSMBShareRemoveACLsReq
}

// share acls info
func (r ApiRemoveDfsSMBShareACLsRequest) Body(body DfsSMBShareRemoveACLsReq) ApiRemoveDfsSMBShareACLsRequest {
	r.body = &body
	return r
}

func (r ApiRemoveDfsSMBShareACLsRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.RemoveDfsSMBShareACLsExecute(r)
}

/*
RemoveDfsSMBShareACLs Method for RemoveDfsSMBShareACLs

remove dfs smb share acls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId dfs smb share id
 @return ApiRemoveDfsSMBShareACLsRequest
*/
func (a *DfsSmbSharesAPIService) RemoveDfsSMBShareACLs(ctx context.Context, dfsSmbShareId int64) ApiRemoveDfsSMBShareACLsRequest {
	return ApiRemoveDfsSMBShareACLsRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) RemoveDfsSMBShareACLsExecute(r ApiRemoveDfsSMBShareACLsRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.RemoveDfsSMBShareACLs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}:remove-acls"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

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

type ApiUpdateDfsSMBShareRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
	body *DfsSMBShareUpdateReq
}

// share info
func (r ApiUpdateDfsSMBShareRequest) Body(body DfsSMBShareUpdateReq) ApiUpdateDfsSMBShareRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDfsSMBShareRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.UpdateDfsSMBShareExecute(r)
}

/*
UpdateDfsSMBShare Method for UpdateDfsSMBShare

Update dfs smb share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId share id
 @return ApiUpdateDfsSMBShareRequest
*/
func (a *DfsSmbSharesAPIService) UpdateDfsSMBShare(ctx context.Context, dfsSmbShareId int64) ApiUpdateDfsSMBShareRequest {
	return ApiUpdateDfsSMBShareRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) UpdateDfsSMBShareExecute(r ApiUpdateDfsSMBShareRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.UpdateDfsSMBShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

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

type ApiUpdateDfsSMBShareACLsRequest struct {
	ctx context.Context
	ApiService *DfsSmbSharesAPIService
	dfsSmbShareId int64
	body *DfsSMBShareUpdateACLsReq
}

// share acls info
func (r ApiUpdateDfsSMBShareACLsRequest) Body(body DfsSMBShareUpdateACLsReq) ApiUpdateDfsSMBShareACLsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDfsSMBShareACLsRequest) Execute() (*DfsSMBShareResp, *http.Response, error) {
	return r.ApiService.UpdateDfsSMBShareACLsExecute(r)
}

/*
UpdateDfsSMBShareACLs Method for UpdateDfsSMBShareACLs

Update dfs smb share ACL

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsSmbShareId share id
 @return ApiUpdateDfsSMBShareACLsRequest
*/
func (a *DfsSmbSharesAPIService) UpdateDfsSMBShareACLs(ctx context.Context, dfsSmbShareId int64) ApiUpdateDfsSMBShareACLsRequest {
	return ApiUpdateDfsSMBShareACLsRequest{
		ApiService: a,
		ctx: ctx,
		dfsSmbShareId: dfsSmbShareId,
	}
}

// Execute executes the request
//  @return DfsSMBShareResp
func (a *DfsSmbSharesAPIService) UpdateDfsSMBShareACLsExecute(r ApiUpdateDfsSMBShareACLsRequest) (*DfsSMBShareResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsSMBShareResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsSmbSharesAPIService.UpdateDfsSMBShareACLs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-smb-shares/{dfs_smb_share_id}:update-acls"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_smb_share_id"+"}", url.PathEscape(parameterValueToString(r.dfsSmbShareId, "dfsSmbShareId")), -1)

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
