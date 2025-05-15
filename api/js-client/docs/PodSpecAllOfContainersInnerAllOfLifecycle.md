# K8SResourceApi.PodSpecAllOfContainersInnerAllOfLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**postStart** | [**PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStart**](PodSpecAllOfContainersInnerAllOfLifecycleAllOfPostStart.md) |  | [optional] 
**preStop** | [**PodSpecAllOfContainersInnerAllOfLifecycleAllOfPreStop**](PodSpecAllOfContainersInnerAllOfLifecycleAllOfPreStop.md) |  | [optional] 
**stopSignal** | **String** | StopSignal defines which signal will be sent to a container when it is being stopped. If not specified, the default is defined by the container runtime in use. StopSignal can only be set for Pods with a non-empty .spec.os.name | [optional] 


