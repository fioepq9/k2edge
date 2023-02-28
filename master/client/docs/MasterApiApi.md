# \MasterApiApi

All URIs are relative to *http://127.0.0.1:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyContainer**](MasterApiApi.md#ApplyContainer) | **Post** /container/apply | 
[**ApplyCronJob**](MasterApiApi.md#ApplyCronJob) | **Post** /cronjob/apply | 
[**ApplyDeployment**](MasterApiApi.md#ApplyDeployment) | **Post** /deployment/apply | 
[**ApplyJob**](MasterApiApi.md#ApplyJob) | **Post** /job/apply | 
[**ApplyToken**](MasterApiApi.md#ApplyToken) | **Post** /token/apply | 
[**AttachContainer**](MasterApiApi.md#AttachContainer) | **Post** /container/attach | 
[**AttachDeployment**](MasterApiApi.md#AttachDeployment) | **Post** /deployment/attach | 
[**ClusterInfo**](MasterApiApi.md#ClusterInfo) | **Get** /cluster/info | 
[**ContainerTop**](MasterApiApi.md#ContainerTop) | **Get** /container/top | 
[**Cordon**](MasterApiApi.md#Cordon) | **Post** /node/cordon | 
[**CreateContainer**](MasterApiApi.md#CreateContainer) | **Post** /container/create | 
[**CreateCronJob**](MasterApiApi.md#CreateCronJob) | **Post** /cronjob/create | 
[**CreateDeployment**](MasterApiApi.md#CreateDeployment) | **Post** /deployment/create | 
[**CreateJob**](MasterApiApi.md#CreateJob) | **Post** /job/create | 
[**CreateNamespace**](MasterApiApi.md#CreateNamespace) | **Post** /namespace/create | 
[**CreateToken**](MasterApiApi.md#CreateToken) | **Post** /token/create | 
[**DeleteContainer**](MasterApiApi.md#DeleteContainer) | **Post** /container/delete | 
[**DeleteCronJob**](MasterApiApi.md#DeleteCronJob) | **Post** /cronjob/delete | 
[**DeleteDeployment**](MasterApiApi.md#DeleteDeployment) | **Post** /deployment/delete | 
[**DeleteJob**](MasterApiApi.md#DeleteJob) | **Post** /job/delete | 
[**DeleteNamespace**](MasterApiApi.md#DeleteNamespace) | **Post** /namespace/delete | 
[**DeleteToken**](MasterApiApi.md#DeleteToken) | **Post** /token/delete | 
[**Drain**](MasterApiApi.md#Drain) | **Post** /node/drain | 
[**ExecContainer**](MasterApiApi.md#ExecContainer) | **Post** /container/exec | 
[**ExecDeployment**](MasterApiApi.md#ExecDeployment) | **Post** /deployment/exec | 
[**GetContainer**](MasterApiApi.md#GetContainer) | **Get** /container/get | 
[**GetCronJob**](MasterApiApi.md#GetCronJob) | **Get** /cronjob/get | 
[**GetDeployment**](MasterApiApi.md#GetDeployment) | **Get** /deployment/get | 
[**GetJob**](MasterApiApi.md#GetJob) | **Get** /job/get | 
[**GetNamespace**](MasterApiApi.md#GetNamespace) | **Get** /namespace/get | 
[**GetToken**](MasterApiApi.md#GetToken) | **Get** /token/get | 
[**HistoryCronJob**](MasterApiApi.md#HistoryCronJob) | **Get** /cronjob/rollout/history | 
[**HistoryDeployment**](MasterApiApi.md#HistoryDeployment) | **Get** /deployment/rollout/history | 
[**ListNamespace**](MasterApiApi.md#ListNamespace) | **Get** /namespace/list | 
[**LogsContainer**](MasterApiApi.md#LogsContainer) | **Post** /container/logs | 
[**LogsCronJob**](MasterApiApi.md#LogsCronJob) | **Post** /cronjob/logs | 
[**LogsDeployment**](MasterApiApi.md#LogsDeployment) | **Post** /deployment/logs | 
[**LogsJob**](MasterApiApi.md#LogsJob) | **Post** /job/logs | 
[**NodeTop**](MasterApiApi.md#NodeTop) | **Get** /node/top | 
[**RunContainer**](MasterApiApi.md#RunContainer) | **Post** /container/run | 
[**Scale**](MasterApiApi.md#Scale) | **Post** /deployment/scale | 
[**Uncordon**](MasterApiApi.md#Uncordon) | **Post** /node/uncordon | 
[**UndoCronJob**](MasterApiApi.md#UndoCronJob) | **Post** /cronjob/rollout/undo | 
[**UndoDeployment**](MasterApiApi.md#UndoDeployment) | **Post** /deployment/rollout/undo | 


# **ApplyContainer**
> ApplyContainerResponse ApplyContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplyContainerRequest**](ApplyContainerRequest.md)|  | 

### Return type

[**ApplyContainerResponse**](ApplyContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyCronJob**
> ApplyCronJobResponse ApplyCronJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplyCronJobRequest**](ApplyCronJobRequest.md)|  | 

### Return type

[**ApplyCronJobResponse**](ApplyCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyDeployment**
> ApplyDeploymentResponse ApplyDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplyDeploymentRequest**](ApplyDeploymentRequest.md)|  | 

### Return type

[**ApplyDeploymentResponse**](ApplyDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyJob**
> DeleteJobResponse ApplyJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteJobRequest**](DeleteJobRequest.md)|  | 

### Return type

[**DeleteJobResponse**](DeleteJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyToken**
> DeleteTokenResponse ApplyToken(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteTokenRequest**](DeleteTokenRequest.md)|  | 

### Return type

[**DeleteTokenResponse**](DeleteTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachContainer**
> AttachContainerResponse AttachContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachContainerRequest**](AttachContainerRequest.md)|  | 

### Return type

[**AttachContainerResponse**](AttachContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachDeployment**
> AttachDeploymentResponse AttachDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachDeploymentRequest**](AttachDeploymentRequest.md)|  | 

### Return type

[**AttachDeploymentResponse**](AttachDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterInfo**
> ClusterInfoResponse ClusterInfo(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterInfoResponse**](ClusterInfoResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerTop**
> ContainerTopResponse ContainerTop(ctx, selector)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | [**interface{}**](.md)|  | 

### Return type

[**ContainerTopResponse**](ContainerTopResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Cordon**
> CordonResponse Cordon(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CordonRequest**](CordonRequest.md)|  | 

### Return type

[**CordonResponse**](CordonResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContainer**
> CreateContainerResponse CreateContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateContainerRequest**](CreateContainerRequest.md)|  | 

### Return type

[**CreateContainerResponse**](CreateContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCronJob**
> CreateCronJobResponse CreateCronJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCronJobRequest**](CreateCronJobRequest.md)|  | 

### Return type

[**CreateCronJobResponse**](CreateCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDeployment**
> CreateDeploymentResponse CreateDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDeploymentRequest**](CreateDeploymentRequest.md)|  | 

### Return type

[**CreateDeploymentResponse**](CreateDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJob**
> CreateJobResponse CreateJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateJobRequest**](CreateJobRequest.md)|  | 

### Return type

[**CreateJobResponse**](CreateJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespace**
> interface{} CreateNamespace(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNamespaceRequest**](CreateNamespaceRequest.md)|  kubectl create namespace &lt;insert-namespace-name-here&gt; OR kubectl create -f ./my-namespace.yaml | 

### Return type

**interface{}**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateToken**
> CreateTokenResponse CreateToken(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTokenRequest**](CreateTokenRequest.md)|  | 

### Return type

[**CreateTokenResponse**](CreateTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContainer**
> DeleteContainerResponse DeleteContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteContainerRequest**](DeleteContainerRequest.md)|  | 

### Return type

[**DeleteContainerResponse**](DeleteContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCronJob**
> DeleteCronJobResponse DeleteCronJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteCronJobRequest**](DeleteCronJobRequest.md)|  | 

### Return type

[**DeleteCronJobResponse**](DeleteCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeployment**
> DeleteDeploymentResponse DeleteDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDeploymentRequest**](DeleteDeploymentRequest.md)|  | 

### Return type

[**DeleteDeploymentResponse**](DeleteDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJobResponse DeleteJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteJobRequest**](DeleteJobRequest.md)|  | 

### Return type

[**DeleteJobResponse**](DeleteJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespace**
> interface{} DeleteNamespace(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteNamespaceRequest**](DeleteNamespaceRequest.md)|  kubectl delete namespaces &lt;insert-some-namespace-name&gt; | 

### Return type

**interface{}**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteToken**
> DeleteTokenResponse DeleteToken(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteTokenRequest**](DeleteTokenRequest.md)|  | 

### Return type

[**DeleteTokenResponse**](DeleteTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Drain**
> DrainResponse Drain(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DrainRequest**](DrainRequest.md)|  | 

### Return type

[**DrainResponse**](DrainResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecContainer**
> ExecContainerResponse ExecContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExecContainerRequest**](ExecContainerRequest.md)|  | 

### Return type

[**ExecContainerResponse**](ExecContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecDeployment**
> ExecDeploymentResponse ExecDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExecDeploymentRequest**](ExecDeploymentRequest.md)|  | 

### Return type

[**ExecDeploymentResponse**](ExecDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainer**
> GetContainerResponse GetContainer(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**GetContainerResponse**](GetContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCronJob**
> GetCronJobResponse GetCronJob(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**GetCronJobResponse**](GetCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> GetDeploymentResponse GetDeployment(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**GetDeploymentResponse**](GetDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJob**
> GetJobResponse GetJob(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**GetJobResponse**](GetJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespace**
> GetNamespaceResponse GetNamespace(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**GetNamespaceResponse**](GetNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetToken**
> GetTokenResponse GetToken(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**GetTokenResponse**](GetTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HistoryCronJob**
> HistoryCronJobResponse HistoryCronJob(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**HistoryCronJobResponse**](HistoryCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HistoryDeployment**
> HistoryDeploymentResponse HistoryDeployment(ctx, todo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todo** | **string**|  | 

### Return type

[**HistoryDeploymentResponse**](HistoryDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespace**
> ListNamespaceResponse ListNamespace(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListNamespaceResponse**](ListNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsContainer**
> LogsContainerResponse LogsContainer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogsContainerRequest**](LogsContainerRequest.md)|  | 

### Return type

[**LogsContainerResponse**](LogsContainerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsCronJob**
> LogsCronJobResponse LogsCronJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogsCronJobRequest**](LogsCronJobRequest.md)|  | 

### Return type

[**LogsCronJobResponse**](LogsCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsDeployment**
> LogsDeploymentResponse LogsDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogsDeploymentRequest**](LogsDeploymentRequest.md)|  | 

### Return type

[**LogsDeploymentResponse**](LogsDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsJob**
> LogsJobResponse LogsJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogsJobRequest**](LogsJobRequest.md)|  | 

### Return type

[**LogsJobResponse**](LogsJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeTop**
> NodeTopResponse NodeTop(ctx, selector)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selector** | [**interface{}**](.md)|  | 

### Return type

[**NodeTopResponse**](NodeTopResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunContainer**
> RunContainerResponse RunContainer(ctx, body)


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

# **Scale**
> ScaleResponse Scale(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScaleRequest**](ScaleRequest.md)|  | 

### Return type

[**ScaleResponse**](ScaleResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Uncordon**
> UncordonResponse Uncordon(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UncordonRequest**](UncordonRequest.md)|  | 

### Return type

[**UncordonResponse**](UncordonResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UndoCronJob**
> UndoCronJobResponse UndoCronJob(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UndoCronJobRequest**](UndoCronJobRequest.md)|  | 

### Return type

[**UndoCronJobResponse**](UndoCronJobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UndoDeployment**
> UndoDeploymentResponse UndoDeployment(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UndoDeploymentRequest**](UndoDeploymentRequest.md)|  | 

### Return type

[**UndoDeploymentResponse**](UndoDeploymentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

