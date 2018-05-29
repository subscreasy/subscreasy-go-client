# \SubsriptionResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscriptionUsingPUT**](SubsriptionResourceApi.md#CancelSubscriptionUsingPUT) | **Put** /api/subscriptions/cancel | cancelSubscription
[**GetActiveSubscriptionsUsingGET**](SubsriptionResourceApi.md#GetActiveSubscriptionsUsingGET) | **Get** /api/subsriptions/subscriber/{secureId} | getActiveSubscriptions
[**GetAllCompanySubscriptionsUsingGET**](SubsriptionResourceApi.md#GetAllCompanySubscriptionsUsingGET) | **Get** /api/subscriptions/company/{id} | getAllCompanySubscriptions
[**GetSubsriptionUsingGET**](SubsriptionResourceApi.md#GetSubsriptionUsingGET) | **Get** /api/subsriptions/{id} | getSubsription
[**StartSubscriptionUsingPOST**](SubsriptionResourceApi.md#StartSubscriptionUsingPOST) | **Post** /api/subscriptions/start | startSubscription


# **CancelSubscriptionUsingPUT**
> Subsription CancelSubscriptionUsingPUT(ctx, cancellation)
cancelSubscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cancellation** | [**Cancellation**](Cancellation.md)| cancellation | 

### Return type

[**Subsription**](Subsription.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveSubscriptionsUsingGET**
> []Subsription GetActiveSubscriptionsUsingGET(ctx, secureId)
getActiveSubscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **secureId** | **string**| secureId | 

### Return type

[**[]Subsription**](Subsription.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllCompanySubscriptionsUsingGET**
> []Subsription GetAllCompanySubscriptionsUsingGET(ctx, id)
getAllCompanySubscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

[**[]Subsription**](Subsription.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubsriptionUsingGET**
> Subsription GetSubsriptionUsingGET(ctx, id)
getSubsription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Subsription**](Subsription.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSubscriptionUsingPOST**
> SubscriptionCreateResult StartSubscriptionUsingPOST(ctx, request)
startSubscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **request** | [**StartSubscriptionRequest**](StartSubscriptionRequest.md)| request | 

### Return type

[**SubscriptionCreateResult**](SubscriptionCreateResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

