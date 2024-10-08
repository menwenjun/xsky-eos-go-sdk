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


// DfsVipsAPIService DfsVipsAPI service
type DfsVipsAPIService service

type ApiGetDfsVIPRequest struct {
	ctx context.Context
	ApiService *DfsVipsAPIService
	dfsVipId int64
}

func (r ApiGetDfsVIPRequest) Execute() (*DfsVIPResp, *http.Response, error) {
	return r.ApiService.GetDfsVIPExecute(r)
}

/*
GetDfsVIP Method for GetDfsVIP

Get dfs vip

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsVipId vip id
 @return ApiGetDfsVIPRequest
*/
func (a *DfsVipsAPIService) GetDfsVIP(ctx context.Context, dfsVipId int64) ApiGetDfsVIPRequest {
	return ApiGetDfsVIPRequest{
		ApiService: a,
		ctx: ctx,
		dfsVipId: dfsVipId,
	}
}

// Execute executes the request
//  @return DfsVIPResp
func (a *DfsVipsAPIService) GetDfsVIPExecute(r ApiGetDfsVIPRequest) (*DfsVIPResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsVIPResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsVipsAPIService.GetDfsVIP")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-vips/{dfs_vip_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_vip_id"+"}", url.PathEscape(parameterValueToString(r.dfsVipId, "dfsVipId")), -1)

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

type ApiListDfsVIPsRequest struct {
	ctx context.Context
	ApiService *DfsVipsAPIService
	limit *int64
	offset *int64
	clusterId *string
	dfsGatewayGroupId *int64
	dfsGatewayZoneId *int64
	primaryGatewayId *int64
}

// paging param
func (r ApiListDfsVIPsRequest) Limit(limit int64) ApiListDfsVIPsRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDfsVIPsRequest) Offset(offset int64) ApiListDfsVIPsRequest {
	r.offset = &offset
	return r
}

// cluster id
func (r ApiListDfsVIPsRequest) ClusterId(clusterId string) ApiListDfsVIPsRequest {
	r.clusterId = &clusterId
	return r
}

// dfs vip group id
func (r ApiListDfsVIPsRequest) DfsGatewayGroupId(dfsGatewayGroupId int64) ApiListDfsVIPsRequest {
	r.dfsGatewayGroupId = &dfsGatewayGroupId
	return r
}

// dfs gateway zone id
func (r ApiListDfsVIPsRequest) DfsGatewayZoneId(dfsGatewayZoneId int64) ApiListDfsVIPsRequest {
	r.dfsGatewayZoneId = &dfsGatewayZoneId
	return r
}

// primary gateway id
func (r ApiListDfsVIPsRequest) PrimaryGatewayId(primaryGatewayId int64) ApiListDfsVIPsRequest {
	r.primaryGatewayId = &primaryGatewayId
	return r
}

func (r ApiListDfsVIPsRequest) Execute() (*DfsVIPsResp, *http.Response, error) {
	return r.ApiService.ListDfsVIPsExecute(r)
}

/*
ListDfsVIPs Method for ListDfsVIPs

List dfs vips

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDfsVIPsRequest
*/
func (a *DfsVipsAPIService) ListDfsVIPs(ctx context.Context) ApiListDfsVIPsRequest {
	return ApiListDfsVIPsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsVIPsResp
func (a *DfsVipsAPIService) ListDfsVIPsExecute(r ApiListDfsVIPsRequest) (*DfsVIPsResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsVIPsResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsVipsAPIService.ListDfsVIPs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-vips/"

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
	if r.dfsGatewayGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_group_id", r.dfsGatewayGroupId, "form", "")
	}
	if r.dfsGatewayZoneId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_zone_id", r.dfsGatewayZoneId, "form", "")
	}
	if r.primaryGatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "primary_gateway_id", r.primaryGatewayId, "form", "")
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

type ApiMoveDfsVIPRequest struct {
	ctx context.Context
	ApiService *DfsVipsAPIService
	dfsVipId int64
	body *DfsVIPMoveReq
}

// moving vip info
func (r ApiMoveDfsVIPRequest) Body(body DfsVIPMoveReq) ApiMoveDfsVIPRequest {
	r.body = &body
	return r
}

func (r ApiMoveDfsVIPRequest) Execute() (*DfsVIPResp, *http.Response, error) {
	return r.ApiService.MoveDfsVIPExecute(r)
}

/*
MoveDfsVIP Method for MoveDfsVIP

move vip to another dfs gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsVipId vip id
 @return ApiMoveDfsVIPRequest
*/
func (a *DfsVipsAPIService) MoveDfsVIP(ctx context.Context, dfsVipId int64) ApiMoveDfsVIPRequest {
	return ApiMoveDfsVIPRequest{
		ApiService: a,
		ctx: ctx,
		dfsVipId: dfsVipId,
	}
}

// Execute executes the request
//  @return DfsVIPResp
func (a *DfsVipsAPIService) MoveDfsVIPExecute(r ApiMoveDfsVIPRequest) (*DfsVIPResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsVIPResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsVipsAPIService.MoveDfsVIP")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-vips/{dfs_vip_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_vip_id"+"}", url.PathEscape(parameterValueToString(r.dfsVipId, "dfsVipId")), -1)

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
