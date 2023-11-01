# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import List, Optional
from pydantic import BaseModel, conlist
from polyaxon_sdk.models.v1_settings_catalog import V1SettingsCatalog

class V1TeamSettings(BaseModel):
    """
    V1TeamSettings
    """
    projects: Optional[conlist(V1SettingsCatalog)] = None
    hubs: Optional[conlist(V1SettingsCatalog)] = None
    registries: Optional[conlist(V1SettingsCatalog)] = None
    __properties = ["projects", "hubs", "registries"]

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
    def from_json(cls, json_str: str) -> V1TeamSettings:
        """Create an instance of V1TeamSettings from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in projects (list)
        _items = []
        if self.projects:
            for _item in self.projects:
                if _item:
                    _items.append(_item.to_dict())
            _dict['projects'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in hubs (list)
        _items = []
        if self.hubs:
            for _item in self.hubs:
                if _item:
                    _items.append(_item.to_dict())
            _dict['hubs'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in registries (list)
        _items = []
        if self.registries:
            for _item in self.registries:
                if _item:
                    _items.append(_item.to_dict())
            _dict['registries'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1TeamSettings:
        """Create an instance of V1TeamSettings from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1TeamSettings.parse_obj(obj)

        _obj = V1TeamSettings.parse_obj({
            "projects": [V1SettingsCatalog.from_dict(_item) for _item in obj.get("projects")] if obj.get("projects") is not None else None,
            "hubs": [V1SettingsCatalog.from_dict(_item) for _item in obj.get("hubs")] if obj.get("hubs") is not None else None,
            "registries": [V1SettingsCatalog.from_dict(_item) for _item in obj.get("registries")] if obj.get("registries") is not None else None
        })
        return _obj

