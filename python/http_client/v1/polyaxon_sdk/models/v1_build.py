# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc33
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
from pydantic import BaseModel, Field, StrictStr, conlist
from polyaxon_sdk.models.v1_cache import V1Cache
from polyaxon_sdk.models.v1_param import V1Param
from polyaxon_sdk.models.v1_patch_strategy import V1PatchStrategy

class V1Build(BaseModel):
    """
    V1Build
    """
    hub_ref: Optional[StrictStr] = Field(None, alias="hubRef")
    connection: Optional[StrictStr] = None
    queue: Optional[StrictStr] = None
    presets: Optional[conlist(StrictStr)] = None
    cache: Optional[V1Cache] = None
    params: Optional[Dict[str, V1Param]] = None
    run_patch: Optional[Dict[str, Any]] = Field(None, alias="runPatch")
    patch_strategy: Optional[V1PatchStrategy] = Field(None, alias="patchStrategy")
    __properties = ["hubRef", "connection", "queue", "presets", "cache", "params", "runPatch", "patchStrategy"]

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
    def from_json(cls, json_str: str) -> V1Build:
        """Create an instance of V1Build from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of cache
        if self.cache:
            _dict['cache'] = self.cache.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each value in params (dict)
        _field_dict = {}
        if self.params:
            for _key in self.params:
                if self.params[_key]:
                    _field_dict[_key] = self.params[_key].to_dict()
            _dict['params'] = _field_dict
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Build:
        """Create an instance of V1Build from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Build.parse_obj(obj)

        _obj = V1Build.parse_obj({
            "hub_ref": obj.get("hubRef"),
            "connection": obj.get("connection"),
            "queue": obj.get("queue"),
            "presets": obj.get("presets"),
            "cache": V1Cache.from_dict(obj.get("cache")) if obj.get("cache") is not None else None,
            "params": dict((_k, V1Param.from_dict(_v)) for _k, _v in obj.get("params").items()) if obj.get("params") is not None else None,
            "run_patch": obj.get("runPatch"),
            "patch_strategy": obj.get("patchStrategy")
        })
        return _obj

