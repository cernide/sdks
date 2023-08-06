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


from typing import Any, Dict, List, Optional
from pydantic import BaseModel, Field, StrictBool, StrictFloat, StrictStr, conlist
from polyaxon_sdk.models.v1_build import V1Build
from polyaxon_sdk.models.v1_cache import V1Cache
from polyaxon_sdk.models.v1_event_trigger import V1EventTrigger
from polyaxon_sdk.models.v1_hook import V1Hook
from polyaxon_sdk.models.v1_io import V1IO
from polyaxon_sdk.models.v1_join import V1Join
from polyaxon_sdk.models.v1_plugins import V1Plugins
from polyaxon_sdk.models.v1_termination import V1Termination
from polyaxon_sdk.models.v1_trigger_policy import V1TriggerPolicy

class V1CompiledOperation(BaseModel):
    """
    V1CompiledOperation
    """
    version: Optional[StrictFloat] = None
    kind: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    tags: Optional[conlist(StrictStr)] = None
    presets: Optional[conlist(StrictStr)] = None
    queue: Optional[StrictStr] = None
    cache: Optional[V1Cache] = None
    termination: Optional[V1Termination] = None
    plugins: Optional[V1Plugins] = None
    schedule: Optional[Dict[str, Any]] = None
    events: Optional[conlist(V1EventTrigger)] = None
    build: Optional[V1Build] = None
    hooks: Optional[conlist(V1Hook)] = None
    dependencies: Optional[conlist(StrictStr)] = None
    trigger: Optional[V1TriggerPolicy] = None
    conditions: Optional[StrictStr] = None
    skip_on_upstream_skip: Optional[StrictBool] = Field(None, alias="skipOnUpstreamSkip")
    matrix: Optional[Dict[str, Any]] = None
    joins: Optional[Dict[str, V1Join]] = None
    inputs: Optional[conlist(V1IO)] = None
    outputs: Optional[conlist(V1IO)] = None
    contexts: Optional[conlist(V1IO)] = None
    is_approved: Optional[StrictBool] = Field(None, alias="isApproved")
    cost: Optional[StrictFloat] = None
    run: Optional[Dict[str, Any]] = None
    __properties = ["version", "kind", "name", "description", "tags", "presets", "queue", "cache", "termination", "plugins", "schedule", "events", "build", "hooks", "dependencies", "trigger", "conditions", "skipOnUpstreamSkip", "matrix", "joins", "inputs", "outputs", "contexts", "isApproved", "cost", "run"]

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
    def from_json(cls, json_str: str) -> V1CompiledOperation:
        """Create an instance of V1CompiledOperation from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in events (list)
        _items = []
        if self.events:
            for _item in self.events:
                if _item:
                    _items.append(_item.to_dict())
            _dict['events'] = _items
        # override the default output from pydantic by calling `to_dict()` of build
        if self.build:
            _dict['build'] = self.build.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in hooks (list)
        _items = []
        if self.hooks:
            for _item in self.hooks:
                if _item:
                    _items.append(_item.to_dict())
            _dict['hooks'] = _items
        # override the default output from pydantic by calling `to_dict()` of each value in joins (dict)
        _field_dict = {}
        if self.joins:
            for _key in self.joins:
                if self.joins[_key]:
                    _field_dict[_key] = self.joins[_key].to_dict()
            _dict['joins'] = _field_dict
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
        # override the default output from pydantic by calling `to_dict()` of each item in contexts (list)
        _items = []
        if self.contexts:
            for _item in self.contexts:
                if _item:
                    _items.append(_item.to_dict())
            _dict['contexts'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1CompiledOperation:
        """Create an instance of V1CompiledOperation from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1CompiledOperation.parse_obj(obj)

        _obj = V1CompiledOperation.parse_obj({
            "version": obj.get("version"),
            "kind": obj.get("kind"),
            "name": obj.get("name"),
            "description": obj.get("description"),
            "tags": obj.get("tags"),
            "presets": obj.get("presets"),
            "queue": obj.get("queue"),
            "cache": V1Cache.from_dict(obj.get("cache")) if obj.get("cache") is not None else None,
            "termination": V1Termination.from_dict(obj.get("termination")) if obj.get("termination") is not None else None,
            "plugins": V1Plugins.from_dict(obj.get("plugins")) if obj.get("plugins") is not None else None,
            "schedule": obj.get("schedule"),
            "events": [V1EventTrigger.from_dict(_item) for _item in obj.get("events")] if obj.get("events") is not None else None,
            "build": V1Build.from_dict(obj.get("build")) if obj.get("build") is not None else None,
            "hooks": [V1Hook.from_dict(_item) for _item in obj.get("hooks")] if obj.get("hooks") is not None else None,
            "dependencies": obj.get("dependencies"),
            "trigger": obj.get("trigger"),
            "conditions": obj.get("conditions"),
            "skip_on_upstream_skip": obj.get("skipOnUpstreamSkip"),
            "matrix": obj.get("matrix"),
            "joins": dict((_k, V1Join.from_dict(_v)) for _k, _v in obj.get("joins").items()) if obj.get("joins") is not None else None,
            "inputs": [V1IO.from_dict(_item) for _item in obj.get("inputs")] if obj.get("inputs") is not None else None,
            "outputs": [V1IO.from_dict(_item) for _item in obj.get("outputs")] if obj.get("outputs") is not None else None,
            "contexts": [V1IO.from_dict(_item) for _item in obj.get("contexts")] if obj.get("contexts") is not None else None,
            "is_approved": obj.get("isApproved"),
            "cost": obj.get("cost"),
            "run": obj.get("run")
        })
        return _obj

