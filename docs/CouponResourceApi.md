# \CouponResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCouponUsingPOST**](CouponResourceApi.md#CreateCouponUsingPOST) | **Post** /api/coupons | createCoupon
[**DeleteCouponUsingDELETE**](CouponResourceApi.md#DeleteCouponUsingDELETE) | **Delete** /api/coupons/{id} | deleteCoupon
[**GetAllCouponsUsingGET**](CouponResourceApi.md#GetAllCouponsUsingGET) | **Get** /api/coupons | getAllCoupons
[**GetCouponUsingGET**](CouponResourceApi.md#GetCouponUsingGET) | **Get** /api/coupons/{id} | getCoupon
[**UpdateCouponUsingPUT**](CouponResourceApi.md#UpdateCouponUsingPUT) | **Put** /api/coupons | updateCoupon


# **CreateCouponUsingPOST**
> Coupon CreateCouponUsingPOST(ctx, coupon)
createCoupon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **coupon** | [**Coupon**](Coupon.md)| coupon | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCouponUsingDELETE**
> DeleteCouponUsingDELETE(ctx, id)
deleteCoupon

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

# **GetAllCouponsUsingGET**
> []Coupon GetAllCouponsUsingGET(ctx, )
getAllCoupons

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Coupon**](Coupon.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponUsingGET**
> Coupon GetCouponUsingGET(ctx, id)
getCoupon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCouponUsingPUT**
> Coupon UpdateCouponUsingPUT(ctx, coupon)
updateCoupon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **coupon** | [**Coupon**](Coupon.md)| coupon | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

