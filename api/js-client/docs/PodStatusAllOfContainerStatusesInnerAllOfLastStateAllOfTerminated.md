# K8SResourceApi.PodStatusAllOfContainerStatusesInnerAllOfLastStateAllOfTerminated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**containerID** | **String** | Container&#39;s ID in the format &#39;&lt;type&gt;://&lt;container_id&gt;&#39; | [optional] 
**exitCode** | **Number** | Exit status from the last termination of the container | [default to 0]
**finishedAt** | **Date** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**message** | **String** | Message regarding the last termination of the container | [optional] 
**reason** | **String** | (brief) reason from the last termination of the container | [optional] 
**signal** | **Number** | Signal from the last termination of the container | [optional] 
**startedAt** | **Date** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 


