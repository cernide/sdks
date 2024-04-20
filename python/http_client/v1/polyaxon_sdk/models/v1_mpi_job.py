# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.1.8
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import Optional
from pydantic import BaseModel, Field, StrictInt, StrictStr
from polyaxon_sdk.models.v1_clean_pod_policy import V1CleanPodPolicy
from polyaxon_sdk.models.v1_kf_replica import V1KFReplica
from polyaxon_sdk.models.v1_scheduling_policy import V1SchedulingPolicy

class V1MPIJob(BaseModel):
    """
    V1MPIJob
    """
    kind: Optional[StrictStr] = 'mpijob'
    clean_pod_policy: Optional[V1CleanPodPolicy] = Field(None, alias="cleanPodPolicy")
    scheduling_policy: Optional[V1SchedulingPolicy] = Field(None, alias="schedulingPolicy")
    slots_per_worker: Optional[StrictInt] = Field(None, alias="slotsPerWorker")
    worker: Optional[V1KFReplica] = None
    launcher: Optional[V1KFReplica] = None
    __properties = ["kind", "cleanPodPolicy", "schedulingPolicy", "slotsPerWorker", "worker", "launcher"]

    class Config:
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> V1MPIJob:
        """Create an instance of V1MPIJob from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of scheduling_policy
        if self.scheduling_policy:
            _dict['schedulingPolicy'] = self.scheduling_policy.to_dict()
        # override the default output from pydantic by calling `to_dict()` of worker
        if self.worker:
            _dict['worker'] = self.worker.to_dict()
        # override the default output from pydantic by calling `to_dict()` of launcher
        if self.launcher:
            _dict['launcher'] = self.launcher.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1MPIJob:
        """Create an instance of V1MPIJob from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1MPIJob.parse_obj(obj)

        _obj = V1MPIJob.parse_obj({
            "kind": obj.get("kind") if obj.get("kind") is not None else 'mpijob',
            "clean_pod_policy": obj.get("cleanPodPolicy"),
            "scheduling_policy": V1SchedulingPolicy.from_dict(obj.get("schedulingPolicy")) if obj.get("schedulingPolicy") is not None else None,
            "slots_per_worker": obj.get("slotsPerWorker"),
            "worker": V1KFReplica.from_dict(obj.get("worker")) if obj.get("worker") is not None else None,
            "launcher": V1KFReplica.from_dict(obj.get("launcher")) if obj.get("launcher") is not None else None
        })
        return _obj

