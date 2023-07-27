# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc28
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json


from typing import Any, Dict, Optional
from pydantic import BaseModel, StrictBool, StrictStr
from polyaxon_sdk.models.v1_artifact_kind import V1ArtifactKind

class V1RunArtifact(BaseModel):
    """
    V1RunArtifact
    """
    name: Optional[StrictStr] = None
    state: Optional[StrictStr] = None
    kind: Optional[V1ArtifactKind] = None
    path: Optional[StrictStr] = None
    connection: Optional[StrictStr] = None
    run: Optional[StrictStr] = None
    summary: Optional[Dict[str, Any]] = None
    is_input: Optional[StrictBool] = None
    meta_info: Optional[Dict[str, Any]] = None
    __properties = ["name", "state", "kind", "path", "connection", "run", "summary", "is_input", "meta_info"]

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
    def from_json(cls, json_str: str) -> V1RunArtifact:
        """Create an instance of V1RunArtifact from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1RunArtifact:
        """Create an instance of V1RunArtifact from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1RunArtifact.parse_obj(obj)

        _obj = V1RunArtifact.parse_obj({
            "name": obj.get("name"),
            "state": obj.get("state"),
            "kind": obj.get("kind"),
            "path": obj.get("path"),
            "connection": obj.get("connection"),
            "run": obj.get("run"),
            "summary": obj.get("summary"),
            "is_input": obj.get("is_input"),
            "meta_info": obj.get("meta_info")
        })
        return _obj

