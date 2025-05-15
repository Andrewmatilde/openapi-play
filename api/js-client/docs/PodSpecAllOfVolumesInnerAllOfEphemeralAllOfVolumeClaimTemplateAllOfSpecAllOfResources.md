# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfEphemeralAllOfVolumeClaimTemplateAllOfSpecAllOfResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**limits** | [**{String: PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue}**](PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue.md) | Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 
**requests** | [**{String: PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue}**](PodSpecAllOfContainersInnerAllOfResourcesAllOfLimitsValue.md) | Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 


