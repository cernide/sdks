# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.6-rc0
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
from polyaxon_sdk.models.v1_artifacts_mount import V1ArtifactsMount
from polyaxon_sdk.models.v1_artifacts_type import V1ArtifactsType
from polyaxon_sdk.models.v1_auth_type import V1AuthType
from polyaxon_sdk.models.v1_compiled_operation import V1CompiledOperation
from polyaxon_sdk.models.v1_connection_resource import V1ConnectionResource
from polyaxon_sdk.models.v1_connection_schema import V1ConnectionSchema
from polyaxon_sdk.models.v1_connection_type import V1ConnectionType
from polyaxon_sdk.models.v1_early_stopping import V1EarlyStopping
from polyaxon_sdk.models.v1_event import V1Event
from polyaxon_sdk.models.v1_event_type import V1EventType
from polyaxon_sdk.models.v1_gcs_type import V1GcsType
from polyaxon_sdk.models.v1_hp_params import V1HpParams
from polyaxon_sdk.models.v1_matrix import V1Matrix
from polyaxon_sdk.models.v1_matrix_kind import V1MatrixKind
from polyaxon_sdk.models.v1_operation import V1Operation
from polyaxon_sdk.models.v1_polyaxon_init_container import V1PolyaxonInitContainer
from polyaxon_sdk.models.v1_polyaxon_sidecar_container import V1PolyaxonSidecarContainer
from polyaxon_sdk.models.v1_reference import V1Reference
from polyaxon_sdk.models.v1_run_schema import V1RunSchema
from polyaxon_sdk.models.v1_s3_type import V1S3Type
from polyaxon_sdk.models.v1_schedule import V1Schedule
from polyaxon_sdk.models.v1_schedule_kind import V1ScheduleKind
from polyaxon_sdk.models.v1_uri_type import V1UriType
from polyaxon_sdk.models.v1_wasb_type import V1WasbType

class V1Schemas(BaseModel):
    """
    V1Schemas
    """
    early_stopping: Optional[V1EarlyStopping] = Field(None, alias="earlyStopping")
    matrix: Optional[V1Matrix] = None
    run: Optional[V1RunSchema] = None
    operation: Optional[V1Operation] = None
    compiled_operation: Optional[V1CompiledOperation] = Field(None, alias="compiledOperation")
    schedule: Optional[V1Schedule] = None
    connection_schema: Optional[V1ConnectionSchema] = Field(None, alias="connectionSchema")
    hp_params: Optional[V1HpParams] = Field(None, alias="hpParams")
    reference: Optional[V1Reference] = None
    artifacts_mount: Optional[V1ArtifactsMount] = Field(None, alias="artifactsMount")
    polyaxon_sidecar_container: Optional[V1PolyaxonSidecarContainer] = Field(None, alias="polyaxonSidecarContainer")
    polyaxon_init_container: Optional[V1PolyaxonInitContainer] = Field(None, alias="polyaxonInitContainer")
    artifacs: Optional[V1ArtifactsType] = None
    wasb: Optional[V1WasbType] = None
    gcs: Optional[V1GcsType] = None
    s3: Optional[V1S3Type] = None
    auth: Optional[V1AuthType] = None
    uri: Optional[V1UriType] = None
    resource: Optional[V1ConnectionResource] = None
    connection: Optional[V1ConnectionType] = None
    event_type: Optional[V1EventType] = Field(None, alias="eventType")
    matrix_kind: Optional[V1MatrixKind] = Field(None, alias="matrixKind")
    schedule_kind: Optional[V1ScheduleKind] = Field(None, alias="scheduleKind")
    event: Optional[V1Event] = None
    __properties = ["earlyStopping", "matrix", "run", "operation", "compiledOperation", "schedule", "connectionSchema", "hpParams", "reference", "artifactsMount", "polyaxonSidecarContainer", "polyaxonInitContainer", "artifacs", "wasb", "gcs", "s3", "auth", "uri", "resource", "connection", "eventType", "matrixKind", "scheduleKind", "event"]

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
    def from_json(cls, json_str: str) -> V1Schemas:
        """Create an instance of V1Schemas from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of early_stopping
        if self.early_stopping:
            _dict['earlyStopping'] = self.early_stopping.to_dict()
        # override the default output from pydantic by calling `to_dict()` of matrix
        if self.matrix:
            _dict['matrix'] = self.matrix.to_dict()
        # override the default output from pydantic by calling `to_dict()` of run
        if self.run:
            _dict['run'] = self.run.to_dict()
        # override the default output from pydantic by calling `to_dict()` of operation
        if self.operation:
            _dict['operation'] = self.operation.to_dict()
        # override the default output from pydantic by calling `to_dict()` of compiled_operation
        if self.compiled_operation:
            _dict['compiledOperation'] = self.compiled_operation.to_dict()
        # override the default output from pydantic by calling `to_dict()` of schedule
        if self.schedule:
            _dict['schedule'] = self.schedule.to_dict()
        # override the default output from pydantic by calling `to_dict()` of connection_schema
        if self.connection_schema:
            _dict['connectionSchema'] = self.connection_schema.to_dict()
        # override the default output from pydantic by calling `to_dict()` of hp_params
        if self.hp_params:
            _dict['hpParams'] = self.hp_params.to_dict()
        # override the default output from pydantic by calling `to_dict()` of reference
        if self.reference:
            _dict['reference'] = self.reference.to_dict()
        # override the default output from pydantic by calling `to_dict()` of artifacts_mount
        if self.artifacts_mount:
            _dict['artifactsMount'] = self.artifacts_mount.to_dict()
        # override the default output from pydantic by calling `to_dict()` of polyaxon_sidecar_container
        if self.polyaxon_sidecar_container:
            _dict['polyaxonSidecarContainer'] = self.polyaxon_sidecar_container.to_dict()
        # override the default output from pydantic by calling `to_dict()` of polyaxon_init_container
        if self.polyaxon_init_container:
            _dict['polyaxonInitContainer'] = self.polyaxon_init_container.to_dict()
        # override the default output from pydantic by calling `to_dict()` of artifacs
        if self.artifacs:
            _dict['artifacs'] = self.artifacs.to_dict()
        # override the default output from pydantic by calling `to_dict()` of wasb
        if self.wasb:
            _dict['wasb'] = self.wasb.to_dict()
        # override the default output from pydantic by calling `to_dict()` of gcs
        if self.gcs:
            _dict['gcs'] = self.gcs.to_dict()
        # override the default output from pydantic by calling `to_dict()` of s3
        if self.s3:
            _dict['s3'] = self.s3.to_dict()
        # override the default output from pydantic by calling `to_dict()` of auth
        if self.auth:
            _dict['auth'] = self.auth.to_dict()
        # override the default output from pydantic by calling `to_dict()` of uri
        if self.uri:
            _dict['uri'] = self.uri.to_dict()
        # override the default output from pydantic by calling `to_dict()` of resource
        if self.resource:
            _dict['resource'] = self.resource.to_dict()
        # override the default output from pydantic by calling `to_dict()` of connection
        if self.connection:
            _dict['connection'] = self.connection.to_dict()
        # override the default output from pydantic by calling `to_dict()` of event_type
        if self.event_type:
            _dict['eventType'] = self.event_type.to_dict()
        # override the default output from pydantic by calling `to_dict()` of event
        if self.event:
            _dict['event'] = self.event.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Schemas:
        """Create an instance of V1Schemas from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Schemas.parse_obj(obj)

        _obj = V1Schemas.parse_obj({
            "early_stopping": V1EarlyStopping.from_dict(obj.get("earlyStopping")) if obj.get("earlyStopping") is not None else None,
            "matrix": V1Matrix.from_dict(obj.get("matrix")) if obj.get("matrix") is not None else None,
            "run": V1RunSchema.from_dict(obj.get("run")) if obj.get("run") is not None else None,
            "operation": V1Operation.from_dict(obj.get("operation")) if obj.get("operation") is not None else None,
            "compiled_operation": V1CompiledOperation.from_dict(obj.get("compiledOperation")) if obj.get("compiledOperation") is not None else None,
            "schedule": V1Schedule.from_dict(obj.get("schedule")) if obj.get("schedule") is not None else None,
            "connection_schema": V1ConnectionSchema.from_dict(obj.get("connectionSchema")) if obj.get("connectionSchema") is not None else None,
            "hp_params": V1HpParams.from_dict(obj.get("hpParams")) if obj.get("hpParams") is not None else None,
            "reference": V1Reference.from_dict(obj.get("reference")) if obj.get("reference") is not None else None,
            "artifacts_mount": V1ArtifactsMount.from_dict(obj.get("artifactsMount")) if obj.get("artifactsMount") is not None else None,
            "polyaxon_sidecar_container": V1PolyaxonSidecarContainer.from_dict(obj.get("polyaxonSidecarContainer")) if obj.get("polyaxonSidecarContainer") is not None else None,
            "polyaxon_init_container": V1PolyaxonInitContainer.from_dict(obj.get("polyaxonInitContainer")) if obj.get("polyaxonInitContainer") is not None else None,
            "artifacs": V1ArtifactsType.from_dict(obj.get("artifacs")) if obj.get("artifacs") is not None else None,
            "wasb": V1WasbType.from_dict(obj.get("wasb")) if obj.get("wasb") is not None else None,
            "gcs": V1GcsType.from_dict(obj.get("gcs")) if obj.get("gcs") is not None else None,
            "s3": V1S3Type.from_dict(obj.get("s3")) if obj.get("s3") is not None else None,
            "auth": V1AuthType.from_dict(obj.get("auth")) if obj.get("auth") is not None else None,
            "uri": V1UriType.from_dict(obj.get("uri")) if obj.get("uri") is not None else None,
            "resource": V1ConnectionResource.from_dict(obj.get("resource")) if obj.get("resource") is not None else None,
            "connection": V1ConnectionType.from_dict(obj.get("connection")) if obj.get("connection") is not None else None,
            "event_type": V1EventType.from_dict(obj.get("eventType")) if obj.get("eventType") is not None else None,
            "matrix_kind": obj.get("matrixKind"),
            "schedule_kind": obj.get("scheduleKind"),
            "event": V1Event.from_dict(obj.get("event")) if obj.get("event") is not None else None
        })
        return _obj

