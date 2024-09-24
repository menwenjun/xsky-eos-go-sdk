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


// DfsGatewayServicesAPIService DfsGatewayServicesAPI service
type DfsGatewayServicesAPIService service

type ApiBatchGetDfsGatewayServiceSamplesRequest struct {
	ctx context.Context
	ApiService *DfsGatewayServicesAPIService
	clusterId *string
	ids *string
	dfsGatewayId *int64
	durationBegin *string
	durationEnd *string
	period *string
}

// cluster id
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) ClusterId(clusterId string) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.clusterId = &clusterId
	return r
}

// dfs gateway service id, format 1,2,3
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) Ids(ids string) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.ids = &ids
	return r
}

// dfs gateway service id
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) DfsGatewayId(dfsGatewayId int64) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.dfsGatewayId = &dfsGatewayId
	return r
}

// duration begin timestamp
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) DurationBegin(durationBegin string) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.durationBegin = &durationBegin
	return r
}

// duration end timestamp
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) DurationEnd(durationEnd string) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.durationEnd = &durationEnd
	return r
}

// samples period
func (r ApiBatchGetDfsGatewayServiceSamplesRequest) Period(period string) ApiBatchGetDfsGatewayServiceSamplesRequest {
	r.period = &period
	return r
}

func (r ApiBatchGetDfsGatewayServiceSamplesRequest) Execute() (*MultiDfsGatewayServicesSamplesResp, *http.Response, error) {
	return r.ApiService.BatchGetDfsGatewayServiceSamplesExecute(r)
}

/*
BatchGetDfsGatewayServiceSamples Method for BatchGetDfsGatewayServiceSamples

Get samples of multiple dfs gateway service by ids or gateway id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBatchGetDfsGatewayServiceSamplesRequest
*/
func (a *DfsGatewayServicesAPIService) BatchGetDfsGatewayServiceSamples(ctx context.Context) ApiBatchGetDfsGatewayServiceSamplesRequest {
	return ApiBatchGetDfsGatewayServiceSamplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MultiDfsGatewayServicesSamplesResp
func (a *DfsGatewayServicesAPIService) BatchGetDfsGatewayServiceSamplesExecute(r ApiBatchGetDfsGatewayServiceSamplesRequest) (*MultiDfsGatewayServicesSamplesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MultiDfsGatewayServicesSamplesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsGatewayServicesAPIService.BatchGetDfsGatewayServiceSamples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-gateway-services/samples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
	}
	if r.ids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids, "form", "")
	}
	if r.dfsGatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_id", r.dfsGatewayId, "form", "")
	}
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

type ApiGetDfsGatewayServiceRequest struct {
	ctx context.Context
	ApiService *DfsGatewayServicesAPIService
	dfsGatewayServiceId int64
}

func (r ApiGetDfsGatewayServiceRequest) Execute() (*DfsGatewayServiceResp, *http.Response, error) {
	return r.ApiService.GetDfsGatewayServiceExecute(r)
}

/*
GetDfsGatewayService Method for GetDfsGatewayService

Get dfs gateway service

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsGatewayServiceId gateway service id
 @return ApiGetDfsGatewayServiceRequest
*/
func (a *DfsGatewayServicesAPIService) GetDfsGatewayService(ctx context.Context, dfsGatewayServiceId int64) ApiGetDfsGatewayServiceRequest {
	return ApiGetDfsGatewayServiceRequest{
		ApiService: a,
		ctx: ctx,
		dfsGatewayServiceId: dfsGatewayServiceId,
	}
}

// Execute executes the request
//  @return DfsGatewayServiceResp
func (a *DfsGatewayServicesAPIService) GetDfsGatewayServiceExecute(r ApiGetDfsGatewayServiceRequest) (*DfsGatewayServiceResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsGatewayServiceResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsGatewayServicesAPIService.GetDfsGatewayService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-gateway-services/{dfs_gateway_service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_gateway_service_id"+"}", url.PathEscape(parameterValueToString(r.dfsGatewayServiceId, "dfsGatewayServiceId")), -1)

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

type ApiGetDfsGatewayServiceSamplesRequest struct {
	ctx context.Context
	ApiService *DfsGatewayServicesAPIService
	dfsGatewayServiceId int64
	durationBegin *string
	durationEnd *string
	period *string
}

// duration begin timestamp
func (r ApiGetDfsGatewayServiceSamplesRequest) DurationBegin(durationBegin string) ApiGetDfsGatewayServiceSamplesRequest {
	r.durationBegin = &durationBegin
	return r
}

// duration end timestamp
func (r ApiGetDfsGatewayServiceSamplesRequest) DurationEnd(durationEnd string) ApiGetDfsGatewayServiceSamplesRequest {
	r.durationEnd = &durationEnd
	return r
}

// samples period
func (r ApiGetDfsGatewayServiceSamplesRequest) Period(period string) ApiGetDfsGatewayServiceSamplesRequest {
	r.period = &period
	return r
}

func (r ApiGetDfsGatewayServiceSamplesRequest) Execute() (*DfsGatewayServiceSamplesResp, *http.Response, error) {
	return r.ApiService.GetDfsGatewayServiceSamplesExecute(r)
}

/*
GetDfsGatewayServiceSamples Method for GetDfsGatewayServiceSamples

Get a dfs gateway service's samples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dfsGatewayServiceId dfs gateway service id
 @return ApiGetDfsGatewayServiceSamplesRequest
*/
func (a *DfsGatewayServicesAPIService) GetDfsGatewayServiceSamples(ctx context.Context, dfsGatewayServiceId int64) ApiGetDfsGatewayServiceSamplesRequest {
	return ApiGetDfsGatewayServiceSamplesRequest{
		ApiService: a,
		ctx: ctx,
		dfsGatewayServiceId: dfsGatewayServiceId,
	}
}

// Execute executes the request
//  @return DfsGatewayServiceSamplesResp
func (a *DfsGatewayServicesAPIService) GetDfsGatewayServiceSamplesExecute(r ApiGetDfsGatewayServiceSamplesRequest) (*DfsGatewayServiceSamplesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsGatewayServiceSamplesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsGatewayServicesAPIService.GetDfsGatewayServiceSamples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-gateway-services/{dfs_gateway_service_id}/samples"
	localVarPath = strings.Replace(localVarPath, "{"+"dfs_gateway_service_id"+"}", url.PathEscape(parameterValueToString(r.dfsGatewayServiceId, "dfsGatewayServiceId")), -1)

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

type ApiListDfsGatewayServicesRequest struct {
	ctx context.Context
	ApiService *DfsGatewayServicesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	clusterId *string
	hostId *int64
	dfsGatewayId *int64
	dfsGatewayGroupId *int64
}

// paging param
func (r ApiListDfsGatewayServicesRequest) Limit(limit int64) ApiListDfsGatewayServicesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListDfsGatewayServicesRequest) Offset(offset int64) ApiListDfsGatewayServicesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListDfsGatewayServicesRequest) Q(q string) ApiListDfsGatewayServicesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListDfsGatewayServicesRequest) Sort(sort string) ApiListDfsGatewayServicesRequest {
	r.sort = &sort
	return r
}

// cluster id
func (r ApiListDfsGatewayServicesRequest) ClusterId(clusterId string) ApiListDfsGatewayServicesRequest {
	r.clusterId = &clusterId
	return r
}

// host id
func (r ApiListDfsGatewayServicesRequest) HostId(hostId int64) ApiListDfsGatewayServicesRequest {
	r.hostId = &hostId
	return r
}

// dfs gateway id
func (r ApiListDfsGatewayServicesRequest) DfsGatewayId(dfsGatewayId int64) ApiListDfsGatewayServicesRequest {
	r.dfsGatewayId = &dfsGatewayId
	return r
}

// dfs gateway group id
func (r ApiListDfsGatewayServicesRequest) DfsGatewayGroupId(dfsGatewayGroupId int64) ApiListDfsGatewayServicesRequest {
	r.dfsGatewayGroupId = &dfsGatewayGroupId
	return r
}

func (r ApiListDfsGatewayServicesRequest) Execute() (*DfsGatewaysResp, *http.Response, error) {
	return r.ApiService.ListDfsGatewayServicesExecute(r)
}

/*
ListDfsGatewayServices Method for ListDfsGatewayServices

List dfs gateway services

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDfsGatewayServicesRequest
*/
func (a *DfsGatewayServicesAPIService) ListDfsGatewayServices(ctx context.Context) ApiListDfsGatewayServicesRequest {
	return ApiListDfsGatewayServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DfsGatewaysResp
func (a *DfsGatewayServicesAPIService) ListDfsGatewayServicesExecute(r ApiListDfsGatewayServicesRequest) (*DfsGatewaysResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DfsGatewaysResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DfsGatewayServicesAPIService.ListDfsGatewayServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dfs-gateway-services/"

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
	if r.hostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host_id", r.hostId, "form", "")
	}
	if r.dfsGatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_id", r.dfsGatewayId, "form", "")
	}
	if r.dfsGatewayGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dfs_gateway_group_id", r.dfsGatewayGroupId, "form", "")
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
