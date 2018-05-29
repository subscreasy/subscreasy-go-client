# \UserResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserUsingPOST**](UserResourceApi.md#CreateUserUsingPOST) | **Post** /api/users | createUser
[**DeleteUserUsingDELETE**](UserResourceApi.md#DeleteUserUsingDELETE) | **Delete** /api/users/{login} | deleteUser
[**GetAllUsersUsingGET**](UserResourceApi.md#GetAllUsersUsingGET) | **Get** /api/users | getAllUsers
[**GetAuthoritiesUsingGET**](UserResourceApi.md#GetAuthoritiesUsingGET) | **Get** /api/users/authorities | getAuthorities
[**GetUserUsingGET**](UserResourceApi.md#GetUserUsingGET) | **Get** /api/users/{login} | getUser
[**UpdateUserUsingPUT**](UserResourceApi.md#UpdateUserUsingPUT) | **Put** /api/users | updateUser


# **CreateUserUsingPOST**
> ResponseEntity CreateUserUsingPOST(ctx, managedUserVM)
createUser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **managedUserVM** | [**ManagedUserVm**](ManagedUserVm.md)| managedUserVM | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserUsingDELETE**
> DeleteUserUsingDELETE(ctx, login)
deleteUser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**| login | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUsersUsingGET**
> []User GetAllUsersUsingGET(ctx, optional)
getAllUsers

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

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthoritiesUsingGET**
> []string GetAuthoritiesUsingGET(ctx, )
getAuthorities

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserUsingGET**
> User GetUserUsingGET(ctx, login)
getUser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**| login | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserUsingPUT**
> User UpdateUserUsingPUT(ctx, managedUserVM)
updateUser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **managedUserVM** | [**ManagedUserVm**](ManagedUserVm.md)| managedUserVM | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

