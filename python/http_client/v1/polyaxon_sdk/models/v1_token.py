# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc30
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
from typing import List, Optional
from pydantic import BaseModel, StrictInt, StrictStr, conlist

class V1Token(BaseModel):
    """
    V1Token
    """
    uuid: Optional[StrictStr] = None
    key: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    scopes: Optional[conlist(StrictStr)] = None
    services: Optional[conlist(StrictStr)] = None
    started_at: Optional[datetime] = None
    expires_at: Optional[datetime] = None
    created_at: Optional[datetime] = None
    updated_at: Optional[datetime] = None
    expiration: Optional[StrictInt] = None
    __properties = ["uuid", "key", "name", "scopes", "services", "started_at", "expires_at", "created_at", "updated_at", "expiration"]

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
    def from_json(cls, json_str: str) -> V1Token:
        """Create an instance of V1Token from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Token:
        """Create an instance of V1Token from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Token.parse_obj(obj)

        _obj = V1Token.parse_obj({
            "uuid": obj.get("uuid"),
            "key": obj.get("key"),
            "name": obj.get("name"),
            "scopes": obj.get("scopes"),
            "services": obj.get("services"),
            "started_at": obj.get("started_at"),
            "expires_at": obj.get("expires_at"),
            "created_at": obj.get("created_at"),
            "updated_at": obj.get("updated_at"),
            "expiration": obj.get("expiration")
        })
        return _obj

