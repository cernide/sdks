# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.4.2
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
from pydantic import BaseModel, Field, StrictBool, StrictFloat, StrictStr, conlist
from polyaxon_sdk.models.v1_build import V1Build
from polyaxon_sdk.models.v1_cache import V1Cache
from polyaxon_sdk.models.v1_hook import V1Hook
from polyaxon_sdk.models.v1_io import V1IO
from polyaxon_sdk.models.v1_plugins import V1Plugins
from polyaxon_sdk.models.v1_template import V1Template
from polyaxon_sdk.models.v1_termination import V1Termination


class V1Component(BaseModel):
    """
    V1Component
    """
    version: Optional[StrictFloat] = None
    kind: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    tags: Optional[conlist(StrictStr)] = None
    presets: Optional[conlist(StrictStr)] = None
    queue: Optional[StrictStr] = None
    cache: Optional[V1Cache] = None
    namespace: Optional[StrictStr] = None
    termination: Optional[V1Termination] = None
    plugins: Optional[V1Plugins] = None
    hooks: Optional[conlist(V1Hook)] = None
    inputs: Optional[conlist(V1IO)] = None
    outputs: Optional[conlist(V1IO)] = None
    build: Optional[V1Build] = None
    run: Optional[Dict[str, Any]] = None
    template: Optional[V1Template] = None
    is_approved: Optional[StrictBool] = Field(None, alias="isApproved")
    cost: Optional[StrictFloat] = None
    __properties = ["version", "kind", "name", "description", "tags", "presets", "queue", "cache", "namespace",
                    "termination", "plugins", "hooks", "inputs", "outputs", "build", "run", "template", "isApproved", "cost"]

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
    def from_json(cls, json_str: str) -> V1Component:
        """Create an instance of V1Component from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of cache
        if self.cache:
            _dict['cache'] = self.cache.to_dict()
        # override the default output from pydantic by calling `to_dict()` of termination
        if self.termination:
            _dict['termination'] = self.termination.to_dict()
        # override the default output from pydantic by calling `to_dict()` of plugins
        if self.plugins:
            _dict['plugins'] = self.plugins.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in hooks (list)
        _items = []
        if self.hooks:
            for _item in self.hooks:
                if _item:
                    _items.append(_item.to_dict())
            _dict['hooks'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in inputs (list)
        _items = []
        if self.inputs:
            for _item in self.inputs:
                if _item:
                    _items.append(_item.to_dict())
            _dict['inputs'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in outputs (list)
        _items = []
        if self.outputs:
            for _item in self.outputs:
                if _item:
                    _items.append(_item.to_dict())
            _dict['outputs'] = _items
        # override the default output from pydantic by calling `to_dict()` of build
        if self.build:
            _dict['build'] = self.build.to_dict()
        # override the default output from pydantic by calling `to_dict()` of template
        if self.template:
            _dict['template'] = self.template.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Component:
        """Create an instance of V1Component from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Component.parse_obj(obj)

        _obj = V1Component.parse_obj({
            "version": obj.get("version"),
            "kind": obj.get("kind"),
            "name": obj.get("name"),
            "description": obj.get("description"),
            "tags": obj.get("tags"),
            "presets": obj.get("presets"),
            "queue": obj.get("queue"),
            "cache": V1Cache.from_dict(obj.get("cache")) if obj.get("cache") is not None else None,
            "namespace": obj.get("namespace"),
            "termination": V1Termination.from_dict(obj.get("termination")) if obj.get("termination") is not None else None,
            "plugins": V1Plugins.from_dict(obj.get("plugins")) if obj.get("plugins") is not None else None,
            "hooks": [V1Hook.from_dict(_item) for _item in obj.get("hooks")] if obj.get("hooks") is not None else None,
            "inputs": [V1IO.from_dict(_item) for _item in obj.get("inputs")] if obj.get("inputs") is not None else None,
            "outputs": [V1IO.from_dict(_item) for _item in obj.get("outputs")] if obj.get("outputs") is not None else None,
            "build": V1Build.from_dict(obj.get("build")) if obj.get("build") is not None else None,
            "run": obj.get("run"),
            "template": V1Template.from_dict(obj.get("template")) if obj.get("template") is not None else None,
            "is_approved": obj.get("isApproved"),
            "cost": obj.get("cost")
        })
        return _obj
