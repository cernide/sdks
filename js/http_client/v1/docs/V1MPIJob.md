# PolyaxonSdk.V1MPIJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**kind** | **String** |  | [optional] [default to &#39;mpi_job&#39;]
**cleanPodPolicy** | [**V1CleanPodPolicy**](V1CleanPodPolicy.md) |  | [optional] 
**schedulingPolicy** | [**V1SchedulingPolicy**](V1SchedulingPolicy.md) |  | [optional] 
**sshAuthMountPath** | **String** |  | [optional] 
**implementation** | [**MPIJobImplementation**](MPIJobImplementation.md) |  | [optional] 
**slotsPerWorker** | **Number** |  | [optional] 
**template** | [**V1KFReplica**](V1KFReplica.md) |  | [optional] 
**worker** | [**V1KFReplica**](V1KFReplica.md) |  | [optional] 
**launcher** | [**V1KFReplica**](V1KFReplica.md) |  | [optional] 


