# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc22
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import Any, Dict, List, Optional
from pydantic import BaseModel, Field, StrictInt, StrictStr, conlist
from polyaxon_sdk.models.v1_tuner import V1Tuner

class V1Iterative(BaseModel):
    """
    V1Iterative
    """
    kind: Optional[StrictStr] = 'iterative'
    params: Optional[Dict[str, Dict[str, Any]]] = None
    max_iterations: Optional[StrictInt] = Field(None, alias="maxIterations")
    seed: Optional[StrictInt] = None
    concurrency: Optional[StrictInt] = None
    tuner: Optional[V1Tuner] = None
    early_stopping: Optional[conlist(Dict[str, Any])] = Field(None, alias="earlyStopping")
    __properties = ["kind", "params", "maxIterations", "seed", "concurrency", "tuner", "earlyStopping"]

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
    def from_json(cls, json_str: str) -> V1Iterative:
        """Create an instance of V1Iterative from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of tuner
        if self.tuner:
            _dict['tuner'] = self.tuner.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Iterative:
        """Create an instance of V1Iterative from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Iterative.parse_obj(obj)

        _obj = V1Iterative.parse_obj({
            "kind": obj.get("kind") if obj.get("kind") is not None else 'iterative',
            "params": obj.get("params"),
            "max_iterations": obj.get("maxIterations"),
            "seed": obj.get("seed"),
            "concurrency": obj.get("concurrency"),
            "tuner": V1Tuner.from_dict(obj.get("tuner")) if obj.get("tuner") is not None else None,
            "early_stopping": obj.get("earlyStopping")
        })
        return _obj

