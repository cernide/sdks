#!/usr/bin/python
#
# Copyright 2018-2023 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc2
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
from pydantic import BaseModel, Field, StrictBool, StrictStr, conlist
from polyaxon_sdk.models.v1_param import V1Param
from polyaxon_sdk.models.v1_statuses import V1Statuses

class V1Hook(BaseModel):
    """
    V1Hook
    """
    hub_ref: Optional[StrictStr] = Field(None, alias="hubRef")
    connection: Optional[StrictStr] = None
    trigger: Optional[V1Statuses] = None
    conditions: Optional[StrictStr] = None
    params: Optional[Dict[str, V1Param]] = None
    queue: Optional[StrictStr] = None
    presets: Optional[conlist(StrictStr)] = None
    disable_defaults: Optional[StrictBool] = Field(None, alias="disableDefaults")
    __properties = ["hubRef", "connection", "trigger", "conditions", "params", "queue", "presets", "disableDefaults"]

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
    def from_json(cls, json_str: str) -> V1Hook:
        """Create an instance of V1Hook from a JSON string"""
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
    def from_dict(cls, obj: dict) -> V1Hook:
        """Create an instance of V1Hook from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Hook.parse_obj(obj)

        _obj = V1Hook.parse_obj({
            "hub_ref": obj.get("hubRef"),
            "connection": obj.get("connection"),
            "trigger": obj.get("trigger"),
            "conditions": obj.get("conditions"),
            "params": dict((_k, V1Param.from_dict(_v)) for _k, _v in obj.get("params").items()) if obj.get("params") is not None else None,
            "queue": obj.get("queue"),
            "presets": obj.get("presets"),
            "disable_defaults": obj.get("disableDefaults")
        })
        return _obj

