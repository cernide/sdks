# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.3.3
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
from pydantic import BaseModel, StrictStr, conlist
from polyaxon_sdk.models.v1_status_condition import V1StatusCondition

class V1EntityNotificationBody(BaseModel):
    """
    V1EntityNotificationBody
    """
    namespace: Optional[StrictStr] = None
    owner: Optional[StrictStr] = None
    project: Optional[StrictStr] = None
    uuid: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    condition: Optional[V1StatusCondition] = None
    connections: Optional[conlist(StrictStr)] = None
    __properties = ["namespace", "owner", "project", "uuid", "name", "condition", "connections"]

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
    def from_json(cls, json_str: str) -> V1EntityNotificationBody:
        """Create an instance of V1EntityNotificationBody from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of condition
        if self.condition:
            _dict['condition'] = self.condition.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1EntityNotificationBody:
        """Create an instance of V1EntityNotificationBody from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1EntityNotificationBody.parse_obj(obj)

        _obj = V1EntityNotificationBody.parse_obj({
            "namespace": obj.get("namespace"),
            "owner": obj.get("owner"),
            "project": obj.get("project"),
            "uuid": obj.get("uuid"),
            "name": obj.get("name"),
            "condition": V1StatusCondition.from_dict(obj.get("condition")) if obj.get("condition") is not None else None,
            "connections": obj.get("connections")
        })
        return _obj

