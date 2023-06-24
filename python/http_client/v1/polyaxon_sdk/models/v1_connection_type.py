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


from typing import Any, Dict, List, Optional
from pydantic import BaseModel, Field, StrictStr, conlist
from polyaxon_sdk.models.v1_connection_kind import V1ConnectionKind
from polyaxon_sdk.models.v1_connection_resource import V1ConnectionResource

class V1ConnectionType(BaseModel):
    """
    V1ConnectionType
    """
    name: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    tags: Optional[StrictStr] = None
    kind: Optional[V1ConnectionKind] = None
    var_schema: Optional[Dict[str, Any]] = Field(None, alias="schema")
    secret: Optional[V1ConnectionResource] = None
    config_map: Optional[V1ConnectionResource] = Field(None, alias="configMap")
    env: Optional[conlist(Dict[str, Any])] = None
    annotations: Optional[Dict[str, StrictStr]] = None
    __properties = ["name", "description", "tags", "kind", "schema", "secret", "configMap", "env", "annotations"]

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
    def from_json(cls, json_str: str) -> V1ConnectionType:
        """Create an instance of V1ConnectionType from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of secret
        if self.secret:
            _dict['secret'] = self.secret.to_dict()
        # override the default output from pydantic by calling `to_dict()` of config_map
        if self.config_map:
            _dict['configMap'] = self.config_map.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1ConnectionType:
        """Create an instance of V1ConnectionType from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1ConnectionType.parse_obj(obj)

        _obj = V1ConnectionType.parse_obj({
            "name": obj.get("name"),
            "description": obj.get("description"),
            "tags": obj.get("tags"),
            "kind": obj.get("kind"),
            "var_schema": obj.get("schema"),
            "secret": V1ConnectionResource.from_dict(obj.get("secret")) if obj.get("secret") is not None else None,
            "config_map": V1ConnectionResource.from_dict(obj.get("configMap")) if obj.get("configMap") is not None else None,
            "env": obj.get("env"),
            "annotations": obj.get("annotations")
        })
        return _obj

