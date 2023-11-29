# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.3
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from typing import Any, Dict, List, Optional
from pydantic import BaseModel, StrictBool, StrictInt, StrictStr, conlist
from polyaxon_sdk.models.v1_statuses import V1Statuses

class V1Agent(BaseModel):
    """
    V1Agent
    """
    uuid: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    tags: Optional[conlist(StrictStr)] = None
    live_state: Optional[StrictInt] = None
    namespace: Optional[StrictStr] = None
    version_api: Optional[Dict[str, Any]] = None
    version: Optional[StrictStr] = None
    content: Optional[StrictStr] = None
    created_at: Optional[datetime] = None
    updated_at: Optional[datetime] = None
    status: Optional[V1Statuses] = None
    is_replica: Optional[StrictBool] = None
    is_ui_managed: Optional[StrictBool] = None
    settings: Optional[Dict[str, Any]] = None
    __properties = ["uuid", "name", "description", "tags", "live_state", "namespace", "version_api", "version", "content", "created_at", "updated_at", "status", "is_replica", "is_ui_managed", "settings"]

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
    def from_json(cls, json_str: str) -> V1Agent:
        """Create an instance of V1Agent from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Agent:
        """Create an instance of V1Agent from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Agent.parse_obj(obj)

        _obj = V1Agent.parse_obj({
            "uuid": obj.get("uuid"),
            "name": obj.get("name"),
            "description": obj.get("description"),
            "tags": obj.get("tags"),
            "live_state": obj.get("live_state"),
            "namespace": obj.get("namespace"),
            "version_api": obj.get("version_api"),
            "version": obj.get("version"),
            "content": obj.get("content"),
            "created_at": obj.get("created_at"),
            "updated_at": obj.get("updated_at"),
            "status": obj.get("status"),
            "is_replica": obj.get("is_replica"),
            "is_ui_managed": obj.get("is_ui_managed"),
            "settings": obj.get("settings")
        })
        return _obj

