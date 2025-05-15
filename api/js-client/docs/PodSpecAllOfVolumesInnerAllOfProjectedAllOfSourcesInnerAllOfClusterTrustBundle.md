# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfProjectedAllOfSourcesInnerAllOfClusterTrustBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**labelSelector** | [**PodSpecAllOfVolumesInnerAllOfProjectedAllOfSourcesInnerAllOfClusterTrustBundleAllOfLabelSelector**](PodSpecAllOfVolumesInnerAllOfProjectedAllOfSourcesInnerAllOfClusterTrustBundleAllOfLabelSelector.md) |  | [optional] 
**name** | **String** | Select a single ClusterTrustBundle by object name.  Mutually-exclusive with signerName and labelSelector. | [optional] 
**optional** | **Boolean** | If true, don&#39;t block pod startup if the referenced ClusterTrustBundle(s) aren&#39;t available.  If using name, then the named ClusterTrustBundle is allowed not to exist.  If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles. | [optional] 
**path** | **String** | Relative path from the volume root to write the bundle. | [default to &#39;&#39;]
**signerName** | **String** | Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name.  The contents of all selected ClusterTrustBundles will be unified and deduplicated. | [optional] 


