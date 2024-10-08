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


// MappingGroupsAPIService MappingGroupsAPI service
type MappingGroupsAPIService service

type ApiAddVolumesRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
	blockVolumeIds *MappingGroupAddVolumesReq
}

// block volume ids
func (r ApiAddVolumesRequest) BlockVolumeIds(blockVolumeIds MappingGroupAddVolumesReq) ApiAddVolumesRequest {
	r.blockVolumeIds = &blockVolumeIds
	return r
}

func (r ApiAddVolumesRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.AddVolumesExecute(r)
}

/*
AddVolumes Method for AddVolumes

add volumes to mapping group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiAddVolumesRequest
*/
func (a *MappingGroupsAPIService) AddVolumes(ctx context.Context, mappingGroupId int64) ApiAddVolumesRequest {
	return ApiAddVolumesRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) AddVolumesExecute(r ApiAddVolumesRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.AddVolumes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}/block-volumes"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.blockVolumeIds == nil {
		return localVarReturnValue, nil, reportError("blockVolumeIds is required and must be specified")
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
	localVarPostBody = r.blockVolumeIds
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

type ApiCreateMappingGroupRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroup *MappingGroupCreateReq
}

// mapping info
func (r ApiCreateMappingGroupRequest) MappingGroup(mappingGroup MappingGroupCreateReq) ApiCreateMappingGroupRequest {
	r.mappingGroup = &mappingGroup
	return r
}

func (r ApiCreateMappingGroupRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.CreateMappingGroupExecute(r)
}

/*
CreateMappingGroup Method for CreateMappingGroup

create a mapping group in access path

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateMappingGroupRequest
*/
func (a *MappingGroupsAPIService) CreateMappingGroup(ctx context.Context) ApiCreateMappingGroupRequest {
	return ApiCreateMappingGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) CreateMappingGroupExecute(r ApiCreateMappingGroupRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.CreateMappingGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mappingGroup == nil {
		return localVarReturnValue, nil, reportError("mappingGroup is required and must be specified")
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
	localVarPostBody = r.mappingGroup
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

type ApiDeleteMappingGroupRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
	force *bool
}

// force delete or not
func (r ApiDeleteMappingGroupRequest) Force(force bool) ApiDeleteMappingGroupRequest {
	r.force = &force
	return r
}

func (r ApiDeleteMappingGroupRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.DeleteMappingGroupExecute(r)
}

/*
DeleteMappingGroup Method for DeleteMappingGroup

Delete mapping group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiDeleteMappingGroupRequest
*/
func (a *MappingGroupsAPIService) DeleteMappingGroup(ctx context.Context, mappingGroupId int64) ApiDeleteMappingGroupRequest {
	return ApiDeleteMappingGroupRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) DeleteMappingGroupExecute(r ApiDeleteMappingGroupRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.DeleteMappingGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

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

type ApiGetMappingGroupRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
}

func (r ApiGetMappingGroupRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.GetMappingGroupExecute(r)
}

/*
GetMappingGroup Method for GetMappingGroup

Get mapping group by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiGetMappingGroupRequest
*/
func (a *MappingGroupsAPIService) GetMappingGroup(ctx context.Context, mappingGroupId int64) ApiGetMappingGroupRequest {
	return ApiGetMappingGroupRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) GetMappingGroupExecute(r ApiGetMappingGroupRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.GetMappingGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

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

type ApiListMappingGroupsRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	limit *int64
	offset *int64
	accessPathId *int64
	clientGroupId *int64
}

// paging param
func (r ApiListMappingGroupsRequest) Limit(limit int64) ApiListMappingGroupsRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListMappingGroupsRequest) Offset(offset int64) ApiListMappingGroupsRequest {
	r.offset = &offset
	return r
}

// access path id
func (r ApiListMappingGroupsRequest) AccessPathId(accessPathId int64) ApiListMappingGroupsRequest {
	r.accessPathId = &accessPathId
	return r
}

// client group id
func (r ApiListMappingGroupsRequest) ClientGroupId(clientGroupId int64) ApiListMappingGroupsRequest {
	r.clientGroupId = &clientGroupId
	return r
}

func (r ApiListMappingGroupsRequest) Execute() (*MappingGroupsResp, *http.Response, error) {
	return r.ApiService.ListMappingGroupsExecute(r)
}

/*
ListMappingGroups Method for ListMappingGroups

List mapping groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMappingGroupsRequest
*/
func (a *MappingGroupsAPIService) ListMappingGroups(ctx context.Context) ApiListMappingGroupsRequest {
	return ApiListMappingGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MappingGroupsResp
func (a *MappingGroupsAPIService) ListMappingGroupsExecute(r ApiListMappingGroupsRequest) (*MappingGroupsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.ListMappingGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.accessPathId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "access_path_id", r.accessPathId, "form", "")
	}
	if r.clientGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client_group_id", r.clientGroupId, "form", "")
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

type ApiRemoveVolumesRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
	blockVolumeIds *MappingGroupRemoveVolumesReq
	force *bool
}

// block volume ids
func (r ApiRemoveVolumesRequest) BlockVolumeIds(blockVolumeIds MappingGroupRemoveVolumesReq) ApiRemoveVolumesRequest {
	r.blockVolumeIds = &blockVolumeIds
	return r
}

// force delete or not
func (r ApiRemoveVolumesRequest) Force(force bool) ApiRemoveVolumesRequest {
	r.force = &force
	return r
}

func (r ApiRemoveVolumesRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.RemoveVolumesExecute(r)
}

/*
RemoveVolumes Method for RemoveVolumes

remove volumes from mapping group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiRemoveVolumesRequest
*/
func (a *MappingGroupsAPIService) RemoveVolumes(ctx context.Context, mappingGroupId int64) ApiRemoveVolumesRequest {
	return ApiRemoveVolumesRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) RemoveVolumesExecute(r ApiRemoveVolumesRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.RemoveVolumes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}/block-volumes"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.blockVolumeIds == nil {
		return localVarReturnValue, nil, reportError("blockVolumeIds is required and must be specified")
	}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
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
	localVarPostBody = r.blockVolumeIds
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

type ApiUpdateMappingGroupRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
	mappingGroup *MappingGroupUpdateReq
}

// mapping info
func (r ApiUpdateMappingGroupRequest) MappingGroup(mappingGroup MappingGroupUpdateReq) ApiUpdateMappingGroupRequest {
	r.mappingGroup = &mappingGroup
	return r
}

func (r ApiUpdateMappingGroupRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.UpdateMappingGroupExecute(r)
}

/*
UpdateMappingGroup Method for UpdateMappingGroup

update mapping group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiUpdateMappingGroupRequest
*/
func (a *MappingGroupsAPIService) UpdateMappingGroup(ctx context.Context, mappingGroupId int64) ApiUpdateMappingGroupRequest {
	return ApiUpdateMappingGroupRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) UpdateMappingGroupExecute(r ApiUpdateMappingGroupRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.UpdateMappingGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mappingGroup == nil {
		return localVarReturnValue, nil, reportError("mappingGroup is required and must be specified")
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
	localVarPostBody = r.mappingGroup
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

type ApiUpdateMappingGroupClientGroupRequest struct {
	ctx context.Context
	ApiService *MappingGroupsAPIService
	mappingGroupId int64
	clientGroupId *MappingGroupUpdateClientGroupReq
}

// client group id
func (r ApiUpdateMappingGroupClientGroupRequest) ClientGroupId(clientGroupId MappingGroupUpdateClientGroupReq) ApiUpdateMappingGroupClientGroupRequest {
	r.clientGroupId = &clientGroupId
	return r
}

func (r ApiUpdateMappingGroupClientGroupRequest) Execute() (*MappingGroupResp, *http.Response, error) {
	return r.ApiService.UpdateMappingGroupClientGroupExecute(r)
}

/*
UpdateMappingGroupClientGroup Method for UpdateMappingGroupClientGroup

update client group in mapping group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mappingGroupId mapping group id
 @return ApiUpdateMappingGroupClientGroupRequest
*/
func (a *MappingGroupsAPIService) UpdateMappingGroupClientGroup(ctx context.Context, mappingGroupId int64) ApiUpdateMappingGroupClientGroupRequest {
	return ApiUpdateMappingGroupClientGroupRequest{
		ApiService: a,
		ctx: ctx,
		mappingGroupId: mappingGroupId,
	}
}

// Execute executes the request
//  @return MappingGroupResp
func (a *MappingGroupsAPIService) UpdateMappingGroupClientGroupExecute(r ApiUpdateMappingGroupClientGroupRequest) (*MappingGroupResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MappingGroupResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingGroupsAPIService.UpdateMappingGroupClientGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mapping-groups/{mapping_group_id}/client-group"
	localVarPath = strings.Replace(localVarPath, "{"+"mapping_group_id"+"}", url.PathEscape(parameterValueToString(r.mappingGroupId, "mappingGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientGroupId == nil {
		return localVarReturnValue, nil, reportError("clientGroupId is required and must be specified")
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
	localVarPostBody = r.clientGroupId
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
