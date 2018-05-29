# \EndpointsApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeUsingPUT**](EndpointsApi.md#AuthorizeUsingPUT) | **Put** /api/authorize | authorize
[**DeductUsingPUT**](EndpointsApi.md#DeductUsingPUT) | **Put** /api/deduct | deduct
[**GetAuthorizedServicesUsingGET**](EndpointsApi.md#GetAuthorizedServicesUsingGET) | **Get** /api/service/subscriber/{secureId} | getAuthorizedServices
[**GetChargingLogBySubscriptionUsingGET**](EndpointsApi.md#GetChargingLogBySubscriptionUsingGET) | **Get** /api/charging-logs/subscription/{id} | getChargingLogBySubscription
[**GetCustomerTotalAmountUsingGET**](EndpointsApi.md#GetCustomerTotalAmountUsingGET) | **Get** /api/customer-totalAmountCharge/{id} | getCustomerTotalAmount
[**GetInvoiceDetailsUsingGET**](EndpointsApi.md#GetInvoiceDetailsUsingGET) | **Get** /api/getInvoiceDetails | getInvoiceDetails
[**GetMessageTemplateUsingGET**](EndpointsApi.md#GetMessageTemplateUsingGET) | **Get** /api/message-templates/email/{lifecycleEventName} | getMessageTemplate
[**GetServiceInstancesBySubscriptionUsingGET**](EndpointsApi.md#GetServiceInstancesBySubscriptionUsingGET) | **Get** /api/service-instances/subscription/{id} | getServiceInstancesBySubscription
[**GetServiceOfferingsBySubscriptionPlanUsingGET**](EndpointsApi.md#GetServiceOfferingsBySubscriptionPlanUsingGET) | **Get** /api/service-offerings/offer/{id} | getServiceOfferingsBySubscriptionPlan
[**GetTotalRevenuePerMonthUsingGET**](EndpointsApi.md#GetTotalRevenuePerMonthUsingGET) | **Get** /api/charging-logs-totalamount-customer/{id} | getTotalRevenuePerMonth


# **AuthorizeUsingPUT**
> bool AuthorizeUsingPUT(ctx, authorization)
authorize

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authorization** | [**Authorization**](Authorization.md)| authorization | 

### Return type

**bool**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeductUsingPUT**
> DeductionResult DeductUsingPUT(ctx, deduction)
deduct

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deduction** | [**Deduction**](Deduction.md)| deduction | 

### Return type

[**DeductionResult**](DeductionResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizedServicesUsingGET**
> AuthorizedServicesResponse GetAuthorizedServicesUsingGET(ctx, secureId)
getAuthorizedServices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **secureId** | **string**| secureId | 

### Return type

[**AuthorizedServicesResponse**](AuthorizedServicesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChargingLogBySubscriptionUsingGET**
> []ChargingLog GetChargingLogBySubscriptionUsingGET(ctx, id)
getChargingLogBySubscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**[]ChargingLog**](ChargingLog.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerTotalAmountUsingGET**
> int64 GetCustomerTotalAmountUsingGET(ctx, id)
getCustomerTotalAmount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

**int64**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceDetailsUsingGET**
> interface{} GetInvoiceDetailsUsingGET(ctx, invoiceRequest)
getInvoiceDetails

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **invoiceRequest** | [**InvoiceRequest**](InvoiceRequest.md)| invoiceRequest | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageTemplateUsingGET**
> MessageTemplate GetMessageTemplateUsingGET(ctx, lifecycleEventName)
getMessageTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lifecycleEventName** | **string**| lifecycleEventName | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstancesBySubscriptionUsingGET**
> []ServiceInstanceResult GetServiceInstancesBySubscriptionUsingGET(ctx, id)
getServiceInstancesBySubscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**[]ServiceInstanceResult**](ServiceInstanceResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceOfferingsBySubscriptionPlanUsingGET**
> []ServiceOfferingResult GetServiceOfferingsBySubscriptionPlanUsingGET(ctx, id)
getServiceOfferingsBySubscriptionPlan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**[]ServiceOfferingResult**](ServiceOfferingResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTotalRevenuePerMonthUsingGET**
> []interface{} GetTotalRevenuePerMonthUsingGET(ctx, id)
getTotalRevenuePerMonth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

