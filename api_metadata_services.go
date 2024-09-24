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


// MetadataServicesAPIService MetadataServicesAPI service
type MetadataServicesAPIService service

type ApiGetMetadataServiceRequest struct {
	ctx context.Context
	ApiService *MetadataServicesAPIService
	metadataServiceId int64
}

func (r ApiGetMetadataServiceRequest) Execute() (*MetadataServiceResp, *http.Response, error) {
	return r.ApiService.GetMetadataServiceExecute(r)
}

/*
GetMetadataService Method for GetMetadataService

get a metadata service

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metadataServiceId metadata service id
 @return ApiGetMetadataServiceRequest
*/
func (a *MetadataServicesAPIService) GetMetadataService(ctx context.Context, metadataServiceId int64) ApiGetMetadataServiceRequest {
	return ApiGetMetadataServiceRequest{
		ApiService: a,
		ctx: ctx,
		metadataServiceId: metadataServiceId,
	}
}

// Execute executes the request
//  @return MetadataServiceResp
func (a *MetadataServicesAPIService) GetMetadataServiceExecute(r ApiGetMetadataServiceRequest) (*MetadataServiceResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MetadataServiceResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataServicesAPIService.GetMetadataService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metadata-services/{metadata_service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"metadata_service_id"+"}", url.PathEscape(parameterValueToString(r.metadataServiceId, "metadataServiceId")), -1)

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

type ApiGetMetadataServiceSamplesRequest struct {
	ctx context.Context
	ApiService *MetadataServicesAPIService
	metadataServiceId int64
	durationBegin *string
	durationEnd *string
	period *string
}

// duration begin timestamp
func (r ApiGetMetadataServiceSamplesRequest) DurationBegin(durationBegin string) ApiGetMetadataServiceSamplesRequest {
	r.durationBegin = &durationBegin
	return r
}

// duration end timestamp
func (r ApiGetMetadataServiceSamplesRequest) DurationEnd(durationEnd string) ApiGetMetadataServiceSamplesRequest {
	r.durationEnd = &durationEnd
	return r
}

// samples period
func (r ApiGetMetadataServiceSamplesRequest) Period(period string) ApiGetMetadataServiceSamplesRequest {
	r.period = &period
	return r
}

func (r ApiGetMetadataServiceSamplesRequest) Execute() (*MetadataServiceSamplesResp, *http.Response, error) {
	return r.ApiService.GetMetadataServiceSamplesExecute(r)
}

/*
GetMetadataServiceSamples Method for GetMetadataServiceSamples

get samples of a metadata service

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metadataServiceId metadata service id
 @return ApiGetMetadataServiceSamplesRequest
*/
func (a *MetadataServicesAPIService) GetMetadataServiceSamples(ctx context.Context, metadataServiceId int64) ApiGetMetadataServiceSamplesRequest {
	return ApiGetMetadataServiceSamplesRequest{
		ApiService: a,
		ctx: ctx,
		metadataServiceId: metadataServiceId,
	}
}

// Execute executes the request
//  @return MetadataServiceSamplesResp
func (a *MetadataServicesAPIService) GetMetadataServiceSamplesExecute(r ApiGetMetadataServiceSamplesRequest) (*MetadataServiceSamplesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MetadataServiceSamplesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataServicesAPIService.GetMetadataServiceSamples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metadata-services/{metadata_service_id}/samples"
	localVarPath = strings.Replace(localVarPath, "{"+"metadata_service_id"+"}", url.PathEscape(parameterValueToString(r.metadataServiceId, "metadataServiceId")), -1)

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

type ApiListMetadataServicesRequest struct {
	ctx context.Context
	ApiService *MetadataServicesAPIService
	limit *int64
	offset *int64
	clusterId *string
	diskIds *int64
	hostId *int64
	metadataClusterId *int64
	q *string
	sort *string
}

// paging param
func (r ApiListMetadataServicesRequest) Limit(limit int64) ApiListMetadataServicesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListMetadataServicesRequest) Offset(offset int64) ApiListMetadataServicesRequest {
	r.offset = &offset
	return r
}

// cluster id
func (r ApiListMetadataServicesRequest) ClusterId(clusterId string) ApiListMetadataServicesRequest {
	r.clusterId = &clusterId
	return r
}

// disk ids
func (r ApiListMetadataServicesRequest) DiskIds(diskIds int64) ApiListMetadataServicesRequest {
	r.diskIds = &diskIds
	return r
}

// host id
func (r ApiListMetadataServicesRequest) HostId(hostId int64) ApiListMetadataServicesRequest {
	r.hostId = &hostId
	return r
}

// metadata cluster id
func (r ApiListMetadataServicesRequest) MetadataClusterId(metadataClusterId int64) ApiListMetadataServicesRequest {
	r.metadataClusterId = &metadataClusterId
	return r
}

// query param of search
func (r ApiListMetadataServicesRequest) Q(q string) ApiListMetadataServicesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListMetadataServicesRequest) Sort(sort string) ApiListMetadataServicesRequest {
	r.sort = &sort
	return r
}

func (r ApiListMetadataServicesRequest) Execute() (*MetadataServicesResp, *http.Response, error) {
	return r.ApiService.ListMetadataServicesExecute(r)
}

/*
ListMetadataServices Method for ListMetadataServices

List all metadata services in the cluster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMetadataServicesRequest
*/
func (a *MetadataServicesAPIService) ListMetadataServices(ctx context.Context) ApiListMetadataServicesRequest {
	return ApiListMetadataServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MetadataServicesResp
func (a *MetadataServicesAPIService) ListMetadataServicesExecute(r ApiListMetadataServicesRequest) (*MetadataServicesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MetadataServicesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataServicesAPIService.ListMetadataServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metadata-services/"

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
	if r.diskIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disk_ids", r.diskIds, "form", "")
	}
	if r.hostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host_id", r.hostId, "form", "")
	}
	if r.metadataClusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metadata_cluster_id", r.metadataClusterId, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
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
