# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.1.6-rc0
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
from pydantic import BaseModel, Field
from polyaxon_sdk.models.v1_dag_ref import V1DagRef
from polyaxon_sdk.models.v1_hub_ref import V1HubRef
from polyaxon_sdk.models.v1_path_ref import V1PathRef
from polyaxon_sdk.models.v1_url_ref import V1UrlRef

class V1Reference(BaseModel):
    """
    V1Reference
    """
    hub_ref: Optional[V1HubRef] = Field(None, alias="hubRef")
    dag_ref: Optional[V1DagRef] = Field(None, alias="dagRef")
    url_ref: Optional[V1UrlRef] = Field(None, alias="urlRef")
    path_ref: Optional[V1PathRef] = Field(None, alias="pathRef")
    __properties = ["hubRef", "dagRef", "urlRef", "pathRef"]

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
    def from_json(cls, json_str: str) -> V1Reference:
        """Create an instance of V1Reference from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of hub_ref
        if self.hub_ref:
            _dict['hubRef'] = self.hub_ref.to_dict()
        # override the default output from pydantic by calling `to_dict()` of dag_ref
        if self.dag_ref:
            _dict['dagRef'] = self.dag_ref.to_dict()
        # override the default output from pydantic by calling `to_dict()` of url_ref
        if self.url_ref:
            _dict['urlRef'] = self.url_ref.to_dict()
        # override the default output from pydantic by calling `to_dict()` of path_ref
        if self.path_ref:
            _dict['pathRef'] = self.path_ref.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Reference:
        """Create an instance of V1Reference from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Reference.parse_obj(obj)

        _obj = V1Reference.parse_obj({
            "hub_ref": V1HubRef.from_dict(obj.get("hubRef")) if obj.get("hubRef") is not None else None,
            "dag_ref": V1DagRef.from_dict(obj.get("dagRef")) if obj.get("dagRef") is not None else None,
            "url_ref": V1UrlRef.from_dict(obj.get("urlRef")) if obj.get("urlRef") is not None else None,
            "path_ref": V1PathRef.from_dict(obj.get("pathRef")) if obj.get("pathRef") is not None else None
        })
        return _obj

