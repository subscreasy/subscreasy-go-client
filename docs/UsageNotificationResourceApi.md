# \UsageNotificationResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsageNotificationUsingPOST**](UsageNotificationResourceApi.md#CreateUsageNotificationUsingPOST) | **Post** /api/usage-notifications | createUsageNotification
[**DeleteUsageNotificationUsingDELETE**](UsageNotificationResourceApi.md#DeleteUsageNotificationUsingDELETE) | **Delete** /api/usage-notifications/{id} | deleteUsageNotification
[**GetAllUsageNotificationsUsingGET**](UsageNotificationResourceApi.md#GetAllUsageNotificationsUsingGET) | **Get** /api/usage-notifications | getAllUsageNotifications
[**GetUsageNotificationUsingGET**](UsageNotificationResourceApi.md#GetUsageNotificationUsingGET) | **Get** /api/usage-notifications/{id} | getUsageNotification
[**UpdateUsageNotificationUsingPUT**](UsageNotificationResourceApi.md#UpdateUsageNotificationUsingPUT) | **Put** /api/usage-notifications | updateUsageNotification


# **CreateUsageNotificationUsingPOST**
> UsageNotification CreateUsageNotificationUsingPOST(ctx, usageNotification)
createUsageNotification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **usageNotification** | [**UsageNotification**](UsageNotification.md)| usageNotification | 

### Return type

[**UsageNotification**](UsageNotification.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsageNotificationUsingDELETE**
> DeleteUsageNotificationUsingDELETE(ctx, id)
deleteUsageNotification

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

# **GetAllUsageNotificationsUsingGET**
> []UsageNotification GetAllUsageNotificationsUsingGET(ctx, )
getAllUsageNotifications

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UsageNotification**](UsageNotification.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageNotificationUsingGET**
> UsageNotification GetUsageNotificationUsingGET(ctx, id)
getUsageNotification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**UsageNotification**](UsageNotification.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsageNotificationUsingPUT**
> UsageNotification UpdateUsageNotificationUsingPUT(ctx, usageNotification)
updateUsageNotification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **usageNotification** | [**UsageNotification**](UsageNotification.md)| usageNotification | 

### Return type

[**UsageNotification**](UsageNotification.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

