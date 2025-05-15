# K8SResourceApi.PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStartAllOfHttpGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**host** | **String** | Host name to connect to, defaults to the pod IP. You probably want to set \&quot;Host\&quot; in httpHeaders instead. | [optional] 
**httpHeaders** | [**[PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStartAllOfHttpGetAllOfHttpHeadersInner]**](PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStartAllOfHttpGetAllOfHttpHeadersInner.md) | Custom headers to set in the request. HTTP allows repeated headers. | [optional] 
**path** | **String** | Path to access on the HTTP server. | [optional] 
**port** | [**PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStartAllOfHttpGetAllOfPort**](PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStartAllOfHttpGetAllOfPort.md) |  | 
**scheme** | **String** | Scheme to use for connecting to the host. Defaults to HTTP. | [optional] 


