# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.1.0-rc5
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
from pydantic import BaseModel, StrictBool, StrictStr, conlist
from polyaxon_sdk.models.v1_managed_by import V1ManagedBy
from polyaxon_sdk.models.v1_run_pending import V1RunPending

class V1OperationBody(BaseModel):
    """
    V1OperationBody
    """
    content: Optional[StrictStr] = None
    is_managed: Optional[StrictBool] = None
    managed_by: Optional[V1ManagedBy] = None
    pending: Optional[V1RunPending] = None
    name: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    tags: Optional[conlist(StrictStr)] = None
    meta_info: Optional[Dict[str, Any]] = None
    __properties = ["content", "is_managed", "managed_by", "pending", "name", "description", "tags", "meta_info"]

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
    def from_json(cls, json_str: str) -> V1OperationBody:
        """Create an instance of V1OperationBody from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1OperationBody:
        """Create an instance of V1OperationBody from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1OperationBody.parse_obj(obj)

        _obj = V1OperationBody.parse_obj({
            "content": obj.get("content"),
            "is_managed": obj.get("is_managed"),
            "managed_by": obj.get("managed_by"),
            "pending": obj.get("pending"),
            "name": obj.get("name"),
            "description": obj.get("description"),
            "tags": obj.get("tags"),
            "meta_info": obj.get("meta_info")
        })
        return _obj

