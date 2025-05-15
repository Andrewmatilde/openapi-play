# K8SResourceApi.PodStatusAllOfContainerStatusesInnerAllOfResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**claims** | [**[PodSpecAllOfContainersInnerAllOfResourcesAllOfClaimsInner]**](PodSpecAllOfContainersInnerAllOfResourcesAllOfClaimsInner.md) | Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers. | [optional] 
**limits** | [**{String: PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue}**](PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue.md) | Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 
**requests** | [**{String: PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue}**](PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue.md) | Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 


