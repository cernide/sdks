# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.2
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
from polyaxon_sdk.models.v1_artifacts_type import V1ArtifactsType
from polyaxon_sdk.models.v1_dockerfile_type import V1DockerfileType
from polyaxon_sdk.models.v1_file_type import V1FileType
from polyaxon_sdk.models.v1_git_type import V1GitType
from polyaxon_sdk.models.v1_tensorboard_type import V1TensorboardType

class V1Init(BaseModel):
    """
    V1Init
    """
    artifacts: Optional[V1ArtifactsType] = None
    paths: Optional[conlist(Dict[str, Any])] = None
    git: Optional[V1GitType] = None
    dockerfile: Optional[V1DockerfileType] = None
    file: Optional[V1FileType] = None
    tensorboard: Optional[V1TensorboardType] = None
    lineage_ref: Optional[StrictStr] = Field(None, alias="lineageRef")
    artifact_ref: Optional[StrictStr] = Field(None, alias="artifactRef")
    model_ref: Optional[StrictStr] = Field(None, alias="modelRef")
    connection: Optional[StrictStr] = None
    path: Optional[StrictStr] = None
    container: Optional[Dict[str, Any]] = None
    __properties = ["artifacts", "paths", "git", "dockerfile", "file", "tensorboard", "lineageRef", "artifactRef", "modelRef", "connection", "path", "container"]

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
    def from_json(cls, json_str: str) -> V1Init:
        """Create an instance of V1Init from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of artifacts
        if self.artifacts:
            _dict['artifacts'] = self.artifacts.to_dict()
        # override the default output from pydantic by calling `to_dict()` of git
        if self.git:
            _dict['git'] = self.git.to_dict()
        # override the default output from pydantic by calling `to_dict()` of dockerfile
        if self.dockerfile:
            _dict['dockerfile'] = self.dockerfile.to_dict()
        # override the default output from pydantic by calling `to_dict()` of file
        if self.file:
            _dict['file'] = self.file.to_dict()
        # override the default output from pydantic by calling `to_dict()` of tensorboard
        if self.tensorboard:
            _dict['tensorboard'] = self.tensorboard.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Init:
        """Create an instance of V1Init from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Init.parse_obj(obj)

        _obj = V1Init.parse_obj({
            "artifacts": V1ArtifactsType.from_dict(obj.get("artifacts")) if obj.get("artifacts") is not None else None,
            "paths": obj.get("paths"),
            "git": V1GitType.from_dict(obj.get("git")) if obj.get("git") is not None else None,
            "dockerfile": V1DockerfileType.from_dict(obj.get("dockerfile")) if obj.get("dockerfile") is not None else None,
            "file": V1FileType.from_dict(obj.get("file")) if obj.get("file") is not None else None,
            "tensorboard": V1TensorboardType.from_dict(obj.get("tensorboard")) if obj.get("tensorboard") is not None else None,
            "lineage_ref": obj.get("lineageRef"),
            "artifact_ref": obj.get("artifactRef"),
            "model_ref": obj.get("modelRef"),
            "connection": obj.get("connection"),
            "path": obj.get("path"),
            "container": obj.get("container")
        })
        return _obj

