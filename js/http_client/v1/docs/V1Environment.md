# PolyaxonSdk.V1Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**labels** | **{String: String}** |  | [optional] 
**annotations** | **{String: String}** |  | [optional] 
**nodeSelector** | **{String: String}** |  | [optional] 
**affinity** | **Object** |  | [optional] 
**tolerations** | **[Object]** | Optional Tolerations to apply. | [optional] 
**nodeName** | **String** | Optional NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements. | [optional] 
**serviceAccountName** | **String** |  | [optional] 
**hostAliases** | **[Object]** | Optional HostAliases is an optional list of hosts and IPs that will be injected into the pod spec. | [optional] 
**securityContext** | **Object** |  | [optional] 
**imagePullSecrets** | **[String]** |  | [optional] 
**hostNetwork** | **Boolean** | Host networking requested for this workflow pod. Default to false. | [optional] 
**hostPID** | **String** | Use the host&#39;s pid namespace. Default to false. | [optional] 
**dnsPolicy** | **String** | Set DNS policy for the pod. Defaults to \&quot;ClusterFirst\&quot;. Valid values are &#39;ClusterFirstWithHostNet&#39;, &#39;ClusterFirst&#39;, &#39;Default&#39; or &#39;None&#39;. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &#39;ClusterFirstWithHostNet&#39;. | [optional] 
**dnsConfig** | **Object** |  | [optional] 
**schedulerName** | **String** |  | [optional] 
**priorityClassName** | **String** | If specified, indicates the pod&#39;s priority. \&quot;system-node-critical\&quot; and \&quot;system-cluster-critical\&quot; are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default. | [optional] 
**priority** | **Number** | The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. | [optional] 
**restartPolicy** | **String** |  | [optional] 


