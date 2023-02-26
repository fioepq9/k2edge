# \WorkerApiApi

All URIs are relative to *http://127.0.0.1:8888*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attach**](WorkerApiApi.md#Attach) | **Post** /container/attach | prepares a streaming endpoint to attach to a running container
[**ContainerStatus**](WorkerApiApi.md#ContainerStatus) | **Get** /container/status | the status of the container
[**Exec**](WorkerApiApi.md#Exec) | **Post** /container/exec | runs a command in a container.
[**ListContainers**](WorkerApiApi.md#ListContainers) | **Get** /container/list | a list of containers
[**RemoveContainer**](WorkerApiApi.md#RemoveContainer) | **Post** /container/remove | removes the container
[**RunContainer**](WorkerApiApi.md#RunContainer) | **Post** /container/run | creates and starts a container
[**StartContainer**](WorkerApiApi.md#StartContainer) | **Post** /container/start | starts the container
[**StopContainer**](WorkerApiApi.md#StopContainer) | **Post** /container/stop | stops any running process that is part of the container
[**Version**](WorkerApiApi.md#Version) | **Get** /version | the version of api


# **Attach**
> AttachResponse Attach(ctx, body)
prepares a streaming endpoint to attach to a running container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachRequest**](AttachRequest.md)|  | 

### Return type

[**AttachResponse**](AttachResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerStatus**
> ContainerStatusResponse ContainerStatus(ctx, selector)
the status of the container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | [**interface{}**](.md)|  | 

### Return type

[**ContainerStatusResponse**](ContainerStatusResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Exec**
> ExecResponse Exec(ctx, body)
runs a command in a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExecRequest**](ExecRequest.md)|  | 

### Return type

[**ExecResponse**](ExecResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainers**
> ListContainersResponse ListContainers(ctx, optional)
a list of containers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkerApiApiListContainersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkerApiApiListContainersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Bool**|  | 
 **all** | **optional.Bool**|  | 
 **latest** | **optional.Bool**|  | 
 **since** | **optional.String**|  | 
 **before** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**ListContainersResponse**](ListContainersResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContainer**
> RemoveContainerResponse RemoveContainer(ctx, body)
removes the container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveContainerRequest**](RemoveContainerRequest.md)|  | 

### Return type

[**RemoveContainerResponse**](RemoveContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunContainer**
> RunContainerResponse RunContainer(ctx, body)
creates and starts a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunContainerRequest**](RunContainerRequest.md)|  | 

### Return type

[**RunContainerResponse**](RunContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartContainer**
> StartContainerResponse StartContainer(ctx, body)
starts the container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartContainerRequest**](StartContainerRequest.md)|  | 

### Return type

[**StartContainerResponse**](StartContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopContainer**
> StopContainerResponse StopContainer(ctx, body)
stops any running process that is part of the container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StopContainerRequest**](StopContainerRequest.md)|  | 

### Return type

[**StopContainerResponse**](StopContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Version**
> VersionResponse Version(ctx, )
the version of api

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

