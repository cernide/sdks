# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc21
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import Dict, List, Optional
from pydantic import BaseModel, Field, StrictStr, conlist
from polyaxon_sdk.models.v1_param import V1Param

class V1Tuner(BaseModel):
    """
    V1Tuner
    """
    hub_ref: Optional[StrictStr] = Field(None, alias="hubRef")
    queue: Optional[StrictStr] = None
    presets: Optional[conlist(StrictStr)] = None
    params: Optional[Dict[str, V1Param]] = None
    __properties = ["hubRef", "queue", "presets", "params"]

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
    def from_json(cls, json_str: str) -> V1Tuner:
        """Create an instance of V1Tuner from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each value in params (dict)
        _field_dict = {}
        if self.params:
            for _key in self.params:
                if self.params[_key]:
                    _field_dict[_key] = self.params[_key].to_dict()
            _dict['params'] = _field_dict
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Tuner:
        """Create an instance of V1Tuner from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Tuner.parse_obj(obj)

        _obj = V1Tuner.parse_obj({
            "hub_ref": obj.get("hubRef"),
            "queue": obj.get("queue"),
            "presets": obj.get("presets"),
            "params": dict((_k, V1Param.from_dict(_v)) for _k, _v in obj.get("params").items()) if obj.get("params") is not None else None
        })
        return _obj

