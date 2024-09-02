# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.4.0
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import Optional
from pydantic import BaseModel, Field, StrictInt, StrictStr

class V1EventVideo(BaseModel):
    """
    V1EventVideo
    """
    height: Optional[StrictInt] = Field(None, description="Height of the video.")
    width: Optional[StrictInt] = Field(None, description="Width of the video.")
    colorspace: Optional[StrictInt] = None
    path: Optional[StrictStr] = None
    content_type: Optional[StrictStr] = None
    __properties = ["height", "width", "colorspace", "path", "content_type"]

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
    def from_json(cls, json_str: str) -> V1EventVideo:
        """Create an instance of V1EventVideo from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1EventVideo:
        """Create an instance of V1EventVideo from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1EventVideo.parse_obj(obj)

        _obj = V1EventVideo.parse_obj({
            "height": obj.get("height"),
            "width": obj.get("width"),
            "colorspace": obj.get("colorspace"),
            "path": obj.get("path"),
            "content_type": obj.get("content_type")
        })
        return _obj

