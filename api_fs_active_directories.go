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


// FsActiveDirectoriesAPIService FsActiveDirectoriesAPI service
type FsActiveDirectoriesAPIService service

type ApiCreateFSActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	body *FSActiveDirectoryCreateReq
	clusterId *string
}

// file storage active directory info
func (r ApiCreateFSActiveDirectoryRequest) Body(body FSActiveDirectoryCreateReq) ApiCreateFSActiveDirectoryRequest {
	r.body = &body
	return r
}

// cluster id
func (r ApiCreateFSActiveDirectoryRequest) ClusterId(clusterId string) ApiCreateFSActiveDirectoryRequest {
	r.clusterId = &clusterId
	return r
}

func (r ApiCreateFSActiveDirectoryRequest) Execute() (*FSActiveDirectoryResp, *http.Response, error) {
	return r.ApiService.CreateFSActiveDirectoryExecute(r)
}

/*
CreateFSActiveDirectory Method for CreateFSActiveDirectory

create file storage active directory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFSActiveDirectoryRequest
*/
func (a *FsActiveDirectoriesAPIService) CreateFSActiveDirectory(ctx context.Context) ApiCreateFSActiveDirectoryRequest {
	return ApiCreateFSActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FSActiveDirectoryResp
func (a *FsActiveDirectoriesAPIService) CreateFSActiveDirectoryExecute(r ApiCreateFSActiveDirectoryRequest) (*FSActiveDirectoryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.CreateFSActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
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

type ApiDeleteFSActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	fsActiveDirectoryId int64
}

func (r ApiDeleteFSActiveDirectoryRequest) Execute() (*FSActiveDirectoryResp, *http.Response, error) {
	return r.ApiService.DeleteFSActiveDirectoryExecute(r)
}

/*
DeleteFSActiveDirectory Method for DeleteFSActiveDirectory

Delete file storage active directory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsActiveDirectoryId file storage active directory id
 @return ApiDeleteFSActiveDirectoryRequest
*/
func (a *FsActiveDirectoriesAPIService) DeleteFSActiveDirectory(ctx context.Context, fsActiveDirectoryId int64) ApiDeleteFSActiveDirectoryRequest {
	return ApiDeleteFSActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		fsActiveDirectoryId: fsActiveDirectoryId,
	}
}

// Execute executes the request
//  @return FSActiveDirectoryResp
func (a *FsActiveDirectoriesAPIService) DeleteFSActiveDirectoryExecute(r ApiDeleteFSActiveDirectoryRequest) (*FSActiveDirectoryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.DeleteFSActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/{fs_active_directory_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_active_directory_id"+"}", url.PathEscape(parameterValueToString(r.fsActiveDirectoryId, "fsActiveDirectoryId")), -1)

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

type ApiGetFSActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	fsActiveDirectoryId int64
}

func (r ApiGetFSActiveDirectoryRequest) Execute() (*FSActiveDirectoryResp, *http.Response, error) {
	return r.ApiService.GetFSActiveDirectoryExecute(r)
}

/*
GetFSActiveDirectory Method for GetFSActiveDirectory

Get a file storage active directory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsActiveDirectoryId file storage active directory id
 @return ApiGetFSActiveDirectoryRequest
*/
func (a *FsActiveDirectoriesAPIService) GetFSActiveDirectory(ctx context.Context, fsActiveDirectoryId int64) ApiGetFSActiveDirectoryRequest {
	return ApiGetFSActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		fsActiveDirectoryId: fsActiveDirectoryId,
	}
}

// Execute executes the request
//  @return FSActiveDirectoryResp
func (a *FsActiveDirectoriesAPIService) GetFSActiveDirectoryExecute(r ApiGetFSActiveDirectoryRequest) (*FSActiveDirectoryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.GetFSActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/{fs_active_directory_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_active_directory_id"+"}", url.PathEscape(parameterValueToString(r.fsActiveDirectoryId, "fsActiveDirectoryId")), -1)

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

type ApiListFSActiveDirectoriesRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	limit *int64
	offset *int64
	clusterId *string
}

// paging param
func (r ApiListFSActiveDirectoriesRequest) Limit(limit int64) ApiListFSActiveDirectoriesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListFSActiveDirectoriesRequest) Offset(offset int64) ApiListFSActiveDirectoriesRequest {
	r.offset = &offset
	return r
}

// cluster id
func (r ApiListFSActiveDirectoriesRequest) ClusterId(clusterId string) ApiListFSActiveDirectoriesRequest {
	r.clusterId = &clusterId
	return r
}

func (r ApiListFSActiveDirectoriesRequest) Execute() (*FSActiveDirectoriesResp, *http.Response, error) {
	return r.ApiService.ListFSActiveDirectoriesExecute(r)
}

/*
ListFSActiveDirectories Method for ListFSActiveDirectories

List file storage active directories

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListFSActiveDirectoriesRequest
*/
func (a *FsActiveDirectoriesAPIService) ListFSActiveDirectories(ctx context.Context) ApiListFSActiveDirectoriesRequest {
	return ApiListFSActiveDirectoriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FSActiveDirectoriesResp
func (a *FsActiveDirectoriesAPIService) ListFSActiveDirectoriesExecute(r ApiListFSActiveDirectoriesRequest) (*FSActiveDirectoriesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoriesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.ListFSActiveDirectories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/"

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

type ApiUpdateFSActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	fsActiveDirectoryId int64
	body *FSActiveDirectoryUpdateReq
}

// file storage active directory info
func (r ApiUpdateFSActiveDirectoryRequest) Body(body FSActiveDirectoryUpdateReq) ApiUpdateFSActiveDirectoryRequest {
	r.body = &body
	return r
}

func (r ApiUpdateFSActiveDirectoryRequest) Execute() (*FSActiveDirectoryResp, *http.Response, error) {
	return r.ApiService.UpdateFSActiveDirectoryExecute(r)
}

/*
UpdateFSActiveDirectory Method for UpdateFSActiveDirectory

Update a file storage active directory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsActiveDirectoryId file storage active directory id
 @return ApiUpdateFSActiveDirectoryRequest
*/
func (a *FsActiveDirectoriesAPIService) UpdateFSActiveDirectory(ctx context.Context, fsActiveDirectoryId int64) ApiUpdateFSActiveDirectoryRequest {
	return ApiUpdateFSActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		fsActiveDirectoryId: fsActiveDirectoryId,
	}
}

// Execute executes the request
//  @return FSActiveDirectoryResp
func (a *FsActiveDirectoriesAPIService) UpdateFSActiveDirectoryExecute(r ApiUpdateFSActiveDirectoryRequest) (*FSActiveDirectoryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.UpdateFSActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/{fs_active_directory_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_active_directory_id"+"}", url.PathEscape(parameterValueToString(r.fsActiveDirectoryId, "fsActiveDirectoryId")), -1)

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

type ApiVerifyFSActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *FsActiveDirectoriesAPIService
	fsActiveDirectoryId int64
}

func (r ApiVerifyFSActiveDirectoryRequest) Execute() (*FSActiveDirectoryResp, *http.Response, error) {
	return r.ApiService.VerifyFSActiveDirectoryExecute(r)
}

/*
VerifyFSActiveDirectory Method for VerifyFSActiveDirectory

Verify a fs active directory or user/group info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsActiveDirectoryId file storage active directory id
 @return ApiVerifyFSActiveDirectoryRequest
*/
func (a *FsActiveDirectoriesAPIService) VerifyFSActiveDirectory(ctx context.Context, fsActiveDirectoryId int64) ApiVerifyFSActiveDirectoryRequest {
	return ApiVerifyFSActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		fsActiveDirectoryId: fsActiveDirectoryId,
	}
}

// Execute executes the request
//  @return FSActiveDirectoryResp
func (a *FsActiveDirectoriesAPIService) VerifyFSActiveDirectoryExecute(r ApiVerifyFSActiveDirectoryRequest) (*FSActiveDirectoryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FSActiveDirectoryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FsActiveDirectoriesAPIService.VerifyFSActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fs-active-directories/{fs_active_directory_id}:verify"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_active_directory_id"+"}", url.PathEscape(parameterValueToString(r.fsActiveDirectoryId, "fsActiveDirectoryId")), -1)

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
