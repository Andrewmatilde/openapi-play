# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfGitRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**directory** | **String** | directory is the target directory name. Must not contain or start with &#39;..&#39;.  If &#39;.&#39; is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name. | [optional] 
**repository** | **String** | repository is the URL | [default to &#39;&#39;]
**revision** | **String** | revision is the commit hash for the specified revision. | [optional] 


