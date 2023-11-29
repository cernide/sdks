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


from typing import Optional
from pydantic import BaseModel, Field, StrictFloat, StrictInt, StrictStr

class V1DiffStoppingPolicy(BaseModel):
    """
    Early stopping with diff factor stopping, this policy computes checks runs against the best run and stops those whose performance is worse than the best by the factor defined by the user.
    """
    kind: Optional[StrictStr] = None
    percent: Optional[StrictFloat] = None
    evaluation_interval: Optional[StrictInt] = Field(None, alias="evaluationInterval", description="Interval/Frequency for applying the policy.")
    min_interval: Optional[StrictInt] = Field(None, alias="minInterval")
    min_samples: Optional[StrictInt] = Field(None, alias="minSamples")
    __properties = ["kind", "percent", "evaluationInterval", "minInterval", "minSamples"]

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
    def from_json(cls, json_str: str) -> V1DiffStoppingPolicy:
        """Create an instance of V1DiffStoppingPolicy from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1DiffStoppingPolicy:
        """Create an instance of V1DiffStoppingPolicy from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1DiffStoppingPolicy.parse_obj(obj)

        _obj = V1DiffStoppingPolicy.parse_obj({
            "kind": obj.get("kind"),
            "percent": obj.get("percent"),
            "evaluation_interval": obj.get("evaluationInterval"),
            "min_interval": obj.get("minInterval"),
            "min_samples": obj.get("minSamples")
        })
        return _obj

