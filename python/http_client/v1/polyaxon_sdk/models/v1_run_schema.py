# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.2.0
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
from polyaxon_sdk.models.v1_dag import V1Dag
from polyaxon_sdk.models.v1_dask_job import V1DaskJob
from polyaxon_sdk.models.v1_job import V1Job
from polyaxon_sdk.models.v1_mpi_job import V1MPIJob
from polyaxon_sdk.models.v1_mx_job import V1MXJob
from polyaxon_sdk.models.v1_paddle_job import V1PaddleJob
from polyaxon_sdk.models.v1_pytorch_job import V1PytorchJob
from polyaxon_sdk.models.v1_ray_job import V1RayJob
from polyaxon_sdk.models.v1_service import V1Service
from polyaxon_sdk.models.v1_tf_job import V1TFJob
from polyaxon_sdk.models.v1_xg_boost_job import V1XGBoostJob

class V1RunSchema(BaseModel):
    """
    V1RunSchema
    """
    job: Optional[V1Job] = None
    service: Optional[V1Service] = None
    dag: Optional[V1Dag] = None
    tf_job: Optional[V1TFJob] = Field(None, alias="tfJob")
    pytorch_job: Optional[V1PytorchJob] = Field(None, alias="pytorchJob")
    mpi_job: Optional[V1MPIJob] = Field(None, alias="mpiJob")
    mx_job: Optional[V1MXJob] = Field(None, alias="mxJob")
    xgboost_job: Optional[V1XGBoostJob] = Field(None, alias="xgboostJob")
    paddle_job: Optional[V1PaddleJob] = Field(None, alias="paddleJob")
    dask_job: Optional[V1DaskJob] = Field(None, alias="daskJob")
    ray_job: Optional[V1RayJob] = Field(None, alias="rayJob")
    __properties = ["job", "service", "dag", "tfJob", "pytorchJob", "mpiJob", "mxJob", "xgboostJob", "paddleJob", "daskJob", "rayJob"]

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
    def from_json(cls, json_str: str) -> V1RunSchema:
        """Create an instance of V1RunSchema from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of job
        if self.job:
            _dict['job'] = self.job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of service
        if self.service:
            _dict['service'] = self.service.to_dict()
        # override the default output from pydantic by calling `to_dict()` of dag
        if self.dag:
            _dict['dag'] = self.dag.to_dict()
        # override the default output from pydantic by calling `to_dict()` of tf_job
        if self.tf_job:
            _dict['tfJob'] = self.tf_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of pytorch_job
        if self.pytorch_job:
            _dict['pytorchJob'] = self.pytorch_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of mpi_job
        if self.mpi_job:
            _dict['mpiJob'] = self.mpi_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of mx_job
        if self.mx_job:
            _dict['mxJob'] = self.mx_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of xgboost_job
        if self.xgboost_job:
            _dict['xgboostJob'] = self.xgboost_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of paddle_job
        if self.paddle_job:
            _dict['paddleJob'] = self.paddle_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of dask_job
        if self.dask_job:
            _dict['daskJob'] = self.dask_job.to_dict()
        # override the default output from pydantic by calling `to_dict()` of ray_job
        if self.ray_job:
            _dict['rayJob'] = self.ray_job.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1RunSchema:
        """Create an instance of V1RunSchema from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1RunSchema.parse_obj(obj)

        _obj = V1RunSchema.parse_obj({
            "job": V1Job.from_dict(obj.get("job")) if obj.get("job") is not None else None,
            "service": V1Service.from_dict(obj.get("service")) if obj.get("service") is not None else None,
            "dag": V1Dag.from_dict(obj.get("dag")) if obj.get("dag") is not None else None,
            "tf_job": V1TFJob.from_dict(obj.get("tfJob")) if obj.get("tfJob") is not None else None,
            "pytorch_job": V1PytorchJob.from_dict(obj.get("pytorchJob")) if obj.get("pytorchJob") is not None else None,
            "mpi_job": V1MPIJob.from_dict(obj.get("mpiJob")) if obj.get("mpiJob") is not None else None,
            "mx_job": V1MXJob.from_dict(obj.get("mxJob")) if obj.get("mxJob") is not None else None,
            "xgboost_job": V1XGBoostJob.from_dict(obj.get("xgboostJob")) if obj.get("xgboostJob") is not None else None,
            "paddle_job": V1PaddleJob.from_dict(obj.get("paddleJob")) if obj.get("paddleJob") is not None else None,
            "dask_job": V1DaskJob.from_dict(obj.get("daskJob")) if obj.get("daskJob") is not None else None,
            "ray_job": V1RayJob.from_dict(obj.get("rayJob")) if obj.get("rayJob") is not None else None
        })
        return _obj

