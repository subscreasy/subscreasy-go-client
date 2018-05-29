# \HistoryResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryUsingPOST**](HistoryResourceApi.md#CreateHistoryUsingPOST) | **Post** /api/histories | createHistory
[**DeleteHistoryUsingDELETE**](HistoryResourceApi.md#DeleteHistoryUsingDELETE) | **Delete** /api/histories/{id} | deleteHistory
[**GetAllHistoriesUsingGET**](HistoryResourceApi.md#GetAllHistoriesUsingGET) | **Get** /api/histories | getAllHistories
[**GetHistoryUsingGET**](HistoryResourceApi.md#GetHistoryUsingGET) | **Get** /api/histories/{id} | getHistory
[**UpdateHistoryUsingPUT**](HistoryResourceApi.md#UpdateHistoryUsingPUT) | **Put** /api/histories | updateHistory


# **CreateHistoryUsingPOST**
> History CreateHistoryUsingPOST(ctx, history)
createHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **history** | [**History**](History.md)| history | 

### Return type

[**History**](History.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHistoryUsingDELETE**
> DeleteHistoryUsingDELETE(ctx, id)
deleteHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHistoriesUsingGET**
> []History GetAllHistoriesUsingGET(ctx, )
getAllHistories

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]History**](History.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryUsingGET**
> History GetHistoryUsingGET(ctx, id)
getHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**History**](History.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHistoryUsingPUT**
> History UpdateHistoryUsingPUT(ctx, history)
updateHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **history** | [**History**](History.md)| history | 

### Return type

[**History**](History.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

