# \ChargingLogResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChargingLogUsingPOST**](ChargingLogResourceApi.md#CreateChargingLogUsingPOST) | **Post** /api/charging-logs | createChargingLog
[**DeleteChargingLogUsingDELETE**](ChargingLogResourceApi.md#DeleteChargingLogUsingDELETE) | **Delete** /api/charging-logs/{id} | deleteChargingLog
[**GetAllChargingLogsUsingGET**](ChargingLogResourceApi.md#GetAllChargingLogsUsingGET) | **Get** /api/charging-logs | getAllChargingLogs
[**GetChargingLogUsingGET**](ChargingLogResourceApi.md#GetChargingLogUsingGET) | **Get** /api/charging-logs/{id} | getChargingLog
[**RefundUsingPOST**](ChargingLogResourceApi.md#RefundUsingPOST) | **Post** /api/charging-logs/refund/{chargingLogId} | refund
[**UpdateChargingLogUsingPUT**](ChargingLogResourceApi.md#UpdateChargingLogUsingPUT) | **Put** /api/charging-logs | updateChargingLog


# **CreateChargingLogUsingPOST**
> ChargingLog CreateChargingLogUsingPOST(ctx, chargingLog)
createChargingLog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **chargingLog** | [**ChargingLog**](ChargingLog.md)| chargingLog | 

### Return type

[**ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChargingLogUsingDELETE**
> DeleteChargingLogUsingDELETE(ctx, id)
deleteChargingLog

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

# **GetAllChargingLogsUsingGET**
> []ChargingLog GetAllChargingLogsUsingGET(ctx, optional)
getAllChargingLogs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number of the requested page | 
 **size** | **int32**| Size of a page | 
 **sort** | [**[]string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChargingLogUsingGET**
> ChargingLog GetChargingLogUsingGET(ctx, id)
getChargingLog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundUsingPOST**
> ChargingLog RefundUsingPOST(ctx, chargingLogId)
refund

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **chargingLogId** | **int64**| chargingLogId | 

### Return type

[**ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChargingLogUsingPUT**
> ChargingLog UpdateChargingLogUsingPUT(ctx, chargingLog)
updateChargingLog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **chargingLog** | [**ChargingLog**](ChargingLog.md)| chargingLog | 

### Return type

[**ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

