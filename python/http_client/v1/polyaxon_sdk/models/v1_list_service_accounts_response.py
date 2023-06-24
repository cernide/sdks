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


from typing import List, Optional
from pydantic import BaseModel, StrictInt, StrictStr, conlist
from polyaxon_sdk.models.v1_service_account import V1ServiceAccount

class V1ListServiceAccountsResponse(BaseModel):
    """
    V1ListServiceAccountsResponse
    """
    count: Optional[StrictInt] = None
    results: Optional[conlist(V1ServiceAccount)] = None
    previous: Optional[StrictStr] = None
    next: Optional[StrictStr] = None
    __properties = ["count", "results", "previous", "next"]

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
    def from_json(cls, json_str: str) -> V1ListServiceAccountsResponse:
        """Create an instance of V1ListServiceAccountsResponse from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in results (list)
        _items = []
        if self.results:
            for _item in self.results:
                if _item:
                    _items.append(_item.to_dict())
            _dict['results'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1ListServiceAccountsResponse:
        """Create an instance of V1ListServiceAccountsResponse from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1ListServiceAccountsResponse.parse_obj(obj)

        _obj = V1ListServiceAccountsResponse.parse_obj({
            "count": obj.get("count"),
            "results": [V1ServiceAccount.from_dict(_item) for _item in obj.get("results")] if obj.get("results") is not None else None,
            "previous": obj.get("previous"),
            "next": obj.get("next")
        })
        return _obj

