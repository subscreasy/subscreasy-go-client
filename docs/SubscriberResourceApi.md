# \SubscriberResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriberUsingPOST**](SubscriberResourceApi.md#CreateSubscriberUsingPOST) | **Post** /api/subscribers | createSubscriber
[**DeleteSubscriberUsingDELETE**](SubscriberResourceApi.md#DeleteSubscriberUsingDELETE) | **Delete** /api/subscribers/{id} | deleteSubscriber
[**GetAllSubscribersUsingGET**](SubscriberResourceApi.md#GetAllSubscribersUsingGET) | **Get** /api/subscribers | getAllSubscribers
[**GetSubscriberByEmailUsingGET**](SubscriberResourceApi.md#GetSubscriberByEmailUsingGET) | **Get** /api/subscribers/email/{email} | getSubscriberByEmail
[**GetSubscriberByNameUsingGET**](SubscriberResourceApi.md#GetSubscriberByNameUsingGET) | **Get** /api/subscribers/name/{name} | getSubscriberByName
[**GetSubscriberUsingGET**](SubscriberResourceApi.md#GetSubscriberUsingGET) | **Get** /api/subscribers/{id} | getSubscriber
[**UpdateSubscriberUsingPUT**](SubscriberResourceApi.md#UpdateSubscriberUsingPUT) | **Put** /api/subscribers | updateSubscriber


# **CreateSubscriberUsingPOST**
> Subscriber CreateSubscriberUsingPOST(ctx, subscriber)
createSubscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subscriber** | [**Subscriber**](Subscriber.md)| subscriber | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriberUsingDELETE**
> DeleteSubscriberUsingDELETE(ctx, id)
deleteSubscriber

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

# **GetAllSubscribersUsingGET**
> []Subscriber GetAllSubscribersUsingGET(ctx, )
getAllSubscribers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriberByEmailUsingGET**
> []Subscriber GetSubscriberByEmailUsingGET(ctx, email)
getSubscriberByEmail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**| email | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriberByNameUsingGET**
> []Subscriber GetSubscriberByNameUsingGET(ctx, name)
getSubscriberByName

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| name | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriberUsingGET**
> Subscriber GetSubscriberUsingGET(ctx, id)
getSubscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscriberUsingPUT**
> Subscriber UpdateSubscriberUsingPUT(ctx, subscriber)
updateSubscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subscriber** | [**Subscriber**](Subscriber.md)| subscriber | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

