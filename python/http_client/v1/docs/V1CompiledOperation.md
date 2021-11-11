# V1CompiledOperation


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**version** | **float** |  | [optional] 
**kind** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**tags** | **list[str]** |  | [optional] 
**presets** | **list[str]** |  | [optional] 
**queue** | **str** |  | [optional] 
**cache** | [**V1Cache**](V1Cache.md) |  | [optional] 
**termination** | [**V1Termination**](V1Termination.md) |  | [optional] 
**plugins** | [**V1Plugins**](V1Plugins.md) |  | [optional] 
**schedule** | **object** |  | [optional] 
**events** | [**list[V1EventTrigger]**](V1EventTrigger.md) |  | [optional] 
**build** | [**V1Build**](V1Build.md) |  | [optional] 
**hooks** | [**list[V1Hook]**](V1Hook.md) |  | [optional] 
**dependencies** | **list[str]** |  | [optional] 
**trigger** | [**V1TriggerPolicy**](V1TriggerPolicy.md) |  | [optional] 
**conditions** | **str** |  | [optional] 
**skip_on_upstream_skip** | **bool** |  | [optional] 
**matrix** | **object** |  | [optional] 
**joins** | [**dict(str, V1Join)**](V1Join.md) |  | [optional] 
**inputs** | [**list[V1IO]**](V1IO.md) |  | [optional] 
**outputs** | [**list[V1IO]**](V1IO.md) |  | [optional] 
**contexts** | [**list[V1IO]**](V1IO.md) |  | [optional] 
**is_approved** | **bool** |  | [optional] 
**cost** | **float** |  | [optional] 
**run** | **object** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


