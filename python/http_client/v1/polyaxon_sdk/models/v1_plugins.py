# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc36
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
from pydantic import BaseModel, Field, StrictBool, StrictStr, conlist
from polyaxon_sdk.models.v1_notification import V1Notification
from polyaxon_sdk.models.v1_polyaxon_sidecar_container import V1PolyaxonSidecarContainer

class V1Plugins(BaseModel):
    """
    V1Plugins
    """
    auth: Optional[StrictBool] = None
    docker: Optional[StrictBool] = None
    shm: Optional[StrictBool] = None
    mount_artifacts_store: Optional[StrictBool] = Field(None, alias="mountArtifactsStore")
    collect_artifacts: Optional[StrictBool] = Field(None, alias="collectArtifacts")
    collect_logs: Optional[StrictBool] = Field(None, alias="collectLogs")
    collect_resources: Optional[StrictBool] = Field(None, alias="collectResources")
    sync_statuses: Optional[StrictBool] = Field(None, alias="syncStatuses")
    auto_resume: Optional[StrictBool] = Field(None, alias="autoResume")
    log_level: Optional[StrictStr] = Field(None, alias="logLevel")
    external_host: Optional[StrictBool] = Field(None, alias="externalHost")
    sidecar: Optional[V1PolyaxonSidecarContainer] = None
    notifications: Optional[conlist(V1Notification)] = None
    __properties = ["auth", "docker", "shm", "mountArtifactsStore", "collectArtifacts", "collectLogs", "collectResources", "syncStatuses", "autoResume", "logLevel", "externalHost", "sidecar", "notifications"]

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
    def from_json(cls, json_str: str) -> V1Plugins:
        """Create an instance of V1Plugins from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of sidecar
        if self.sidecar:
            _dict['sidecar'] = self.sidecar.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in notifications (list)
        _items = []
        if self.notifications:
            for _item in self.notifications:
                if _item:
                    _items.append(_item.to_dict())
            _dict['notifications'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Plugins:
        """Create an instance of V1Plugins from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Plugins.parse_obj(obj)

        _obj = V1Plugins.parse_obj({
            "auth": obj.get("auth"),
            "docker": obj.get("docker"),
            "shm": obj.get("shm"),
            "mount_artifacts_store": obj.get("mountArtifactsStore"),
            "collect_artifacts": obj.get("collectArtifacts"),
            "collect_logs": obj.get("collectLogs"),
            "collect_resources": obj.get("collectResources"),
            "sync_statuses": obj.get("syncStatuses"),
            "auto_resume": obj.get("autoResume"),
            "log_level": obj.get("logLevel"),
            "external_host": obj.get("externalHost"),
            "sidecar": V1PolyaxonSidecarContainer.from_dict(obj.get("sidecar")) if obj.get("sidecar") is not None else None,
            "notifications": [V1Notification.from_dict(_item) for _item in obj.get("notifications")] if obj.get("notifications") is not None else None
        })
        return _obj

