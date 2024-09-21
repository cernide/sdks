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
from polyaxon_sdk.models.v1_component import V1Component
from polyaxon_sdk.models.v1_event_trigger import V1EventTrigger
from polyaxon_sdk.models.v1_hook import V1Hook
from polyaxon_sdk.models.v1_join import V1Join
from polyaxon_sdk.models.v1_param import V1Param
from polyaxon_sdk.models.v1_patch_strategy import V1PatchStrategy
from polyaxon_sdk.models.v1_plugins import V1Plugins
from polyaxon_sdk.models.v1_template import V1Template
from polyaxon_sdk.models.v1_termination import V1Termination
from polyaxon_sdk.models.v1_trigger_policy import V1TriggerPolicy


class V1Operation(BaseModel):
    """
    V1Operation
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
    schedule: Optional[Dict[str, Any]] = None
    events: Optional[conlist(V1EventTrigger)] = None
    hooks: Optional[conlist(V1Hook)] = None
    dependencies: Optional[conlist(StrictStr)] = None
    trigger: Optional[V1TriggerPolicy] = None
    conditions: Optional[StrictStr] = None
    skip_on_upstream_skip: Optional[StrictBool] = Field(
        None, alias="skipOnUpstreamSkip")
    matrix: Optional[Dict[str, Any]] = None
    joins: Optional[Dict[str, V1Join]] = None
    params: Optional[Dict[str, V1Param]] = None
    run_patch: Optional[Dict[str, Any]] = Field(None, alias="runPatch")
    patch_strategy: Optional[V1PatchStrategy] = Field(
        None, alias="patchStrategy")
    is_preset: Optional[StrictBool] = Field(None, alias="isPreset")
    is_approved: Optional[StrictBool] = Field(None, alias="isApproved")
    template: Optional[V1Template] = None
    build: Optional[V1Build] = None
    cost: Optional[StrictFloat] = None
    path_ref: Optional[StrictStr] = Field(None, alias="pathRef")
    hub_ref: Optional[StrictStr] = Field(None, alias="hubRef")
    dag_ref: Optional[StrictStr] = Field(None, alias="dagRef")
    url_ref: Optional[StrictStr] = Field(None, alias="urlRef")
    component: Optional[V1Component] = None
    __properties = ["version", "kind", "name", "description", "tags", "presets", "queue", "cache", "namespace", "termination", "plugins", "schedule", "events", "hooks", "dependencies", "trigger",
                    "conditions", "skipOnUpstreamSkip", "matrix", "joins", "params", "runPatch", "patchStrategy", "isPreset", "isApproved", "template", "build", "cost", "pathRef", "hubRef", "dagRef", "urlRef", "component"]

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
    def from_json(cls, json_str: str) -> V1Operation:
        """Create an instance of V1Operation from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each value in params (dict)
        _field_dict = {}
        if self.params:
            for _key in self.params:
                if self.params[_key]:
                    _field_dict[_key] = self.params[_key].to_dict()
            _dict['params'] = _field_dict
        # override the default output from pydantic by calling `to_dict()` of template
        if self.template:
            _dict['template'] = self.template.to_dict()
        # override the default output from pydantic by calling `to_dict()` of build
        if self.build:
            _dict['build'] = self.build.to_dict()
        # override the default output from pydantic by calling `to_dict()` of component
        if self.component:
            _dict['component'] = self.component.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Operation:
        """Create an instance of V1Operation from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Operation.parse_obj(obj)

        _obj = V1Operation.parse_obj({
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
            "schedule": obj.get("schedule"),
            "events": [V1EventTrigger.from_dict(_item) for _item in obj.get("events")] if obj.get("events") is not None else None,
            "hooks": [V1Hook.from_dict(_item) for _item in obj.get("hooks")] if obj.get("hooks") is not None else None,
            "dependencies": obj.get("dependencies"),
            "trigger": obj.get("trigger"),
            "conditions": obj.get("conditions"),
            "skip_on_upstream_skip": obj.get("skipOnUpstreamSkip"),
            "matrix": obj.get("matrix"),
            "joins": dict((_k, V1Join.from_dict(_v)) for _k, _v in obj.get("joins").items()) if obj.get("joins") is not None else None,
            "params": dict((_k, V1Param.from_dict(_v)) for _k, _v in obj.get("params").items()) if obj.get("params") is not None else None,
            "run_patch": obj.get("runPatch"),
            "patch_strategy": obj.get("patchStrategy"),
            "is_preset": obj.get("isPreset"),
            "is_approved": obj.get("isApproved"),
            "template": V1Template.from_dict(obj.get("template")) if obj.get("template") is not None else None,
            "build": V1Build.from_dict(obj.get("build")) if obj.get("build") is not None else None,
            "cost": obj.get("cost"),
            "path_ref": obj.get("pathRef"),
            "hub_ref": obj.get("hubRef"),
            "dag_ref": obj.get("dagRef"),
            "url_ref": obj.get("urlRef"),
            "component": V1Component.from_dict(obj.get("component")) if obj.get("component") is not None else None
        })
        return _obj
