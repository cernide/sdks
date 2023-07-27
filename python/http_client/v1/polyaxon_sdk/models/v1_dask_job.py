# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc28
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
from pydantic import BaseModel, StrictStr
from polyaxon_sdk.models.v1_dask_replica import V1DaskReplica

class V1DaskJob(BaseModel):
    """
    V1DaskJob
    """
    kind: Optional[StrictStr] = 'daskjob'
    job: Optional[V1DaskReplica] = None
    worker: Optional[V1DaskReplica] = None
    scheduler: Optional[V1DaskReplica] = None
    __properties = ["kind", "job", "worker", "scheduler"]

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
    def from_json(cls, json_str: str) -> V1DaskJob:
        """Create an instance of V1DaskJob from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of job
        if self.job:
            _dict['job'] = self.job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of worker
        if self.worker:
            _dict['worker'] = self.worker.to_dict()
        # override the default output from pydantic by calling `to_dict()` of scheduler
        if self.scheduler:
            _dict['scheduler'] = self.scheduler.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1DaskJob:
        """Create an instance of V1DaskJob from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1DaskJob.parse_obj(obj)

        _obj = V1DaskJob.parse_obj({
            "kind": obj.get("kind") if obj.get("kind") is not None else 'daskjob',
            "job": V1DaskReplica.from_dict(obj.get("job")) if obj.get("job") is not None else None,
            "worker": V1DaskReplica.from_dict(obj.get("worker")) if obj.get("worker") is not None else None,
            "scheduler": V1DaskReplica.from_dict(obj.get("scheduler")) if obj.get("scheduler") is not None else None
        })
        return _obj

