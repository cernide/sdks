# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc41
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
from inspect import getfullargspec
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from typing import Optional
from pydantic import BaseModel, Field, StrictFloat, StrictInt, StrictStr
from polyaxon_sdk.models.v1_event_artifact import V1EventArtifact
from polyaxon_sdk.models.v1_event_audio import V1EventAudio
from polyaxon_sdk.models.v1_event_chart import V1EventChart
from polyaxon_sdk.models.v1_event_confusion_matrix import V1EventConfusionMatrix
from polyaxon_sdk.models.v1_event_curve import V1EventCurve
from polyaxon_sdk.models.v1_event_dataframe import V1EventDataframe
from polyaxon_sdk.models.v1_event_histogram import V1EventHistogram
from polyaxon_sdk.models.v1_event_image import V1EventImage
from polyaxon_sdk.models.v1_event_model import V1EventModel
from polyaxon_sdk.models.v1_event_video import V1EventVideo

class V1Event(BaseModel):
    """
    V1Event
    """
    timestamp: Optional[datetime] = None
    step: Optional[StrictInt] = Field(None, description="Global step of the event.")
    metric: Optional[StrictFloat] = None
    image: Optional[V1EventImage] = None
    histogram: Optional[V1EventHistogram] = None
    audio: Optional[V1EventAudio] = None
    video: Optional[V1EventVideo] = None
    html: Optional[StrictStr] = None
    text: Optional[StrictStr] = None
    chart: Optional[V1EventChart] = None
    model: Optional[V1EventModel] = None
    artifact: Optional[V1EventArtifact] = None
    dataframe: Optional[V1EventDataframe] = None
    curve: Optional[V1EventCurve] = None
    confusion: Optional[V1EventConfusionMatrix] = None
    __properties = ["timestamp", "step", "metric", "image", "histogram", "audio", "video", "html", "text", "chart", "model", "artifact", "dataframe", "curve", "confusion"]

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
    def from_json(cls, json_str: str) -> V1Event:
        """Create an instance of V1Event from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of image
        if self.image:
            _dict['image'] = self.image.to_dict()
        # override the default output from pydantic by calling `to_dict()` of histogram
        if self.histogram:
            _dict['histogram'] = self.histogram.to_dict()
        # override the default output from pydantic by calling `to_dict()` of audio
        if self.audio:
            _dict['audio'] = self.audio.to_dict()
        # override the default output from pydantic by calling `to_dict()` of video
        if self.video:
            _dict['video'] = self.video.to_dict()
        # override the default output from pydantic by calling `to_dict()` of chart
        if self.chart:
            _dict['chart'] = self.chart.to_dict()
        # override the default output from pydantic by calling `to_dict()` of model
        if self.model:
            _dict['model'] = self.model.to_dict()
        # override the default output from pydantic by calling `to_dict()` of artifact
        if self.artifact:
            _dict['artifact'] = self.artifact.to_dict()
        # override the default output from pydantic by calling `to_dict()` of dataframe
        if self.dataframe:
            _dict['dataframe'] = self.dataframe.to_dict()
        # override the default output from pydantic by calling `to_dict()` of curve
        if self.curve:
            _dict['curve'] = self.curve.to_dict()
        # override the default output from pydantic by calling `to_dict()` of confusion
        if self.confusion:
            _dict['confusion'] = self.confusion.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1Event:
        """Create an instance of V1Event from a dict"""
        if obj is None:
            return None

        if type(obj) is not dict:
            return V1Event.parse_obj(obj)

        _obj = V1Event.parse_obj({
            "timestamp": obj.get("timestamp"),
            "step": obj.get("step"),
            "metric": obj.get("metric"),
            "image": V1EventImage.from_dict(obj.get("image")) if obj.get("image") is not None else None,
            "histogram": V1EventHistogram.from_dict(obj.get("histogram")) if obj.get("histogram") is not None else None,
            "audio": V1EventAudio.from_dict(obj.get("audio")) if obj.get("audio") is not None else None,
            "video": V1EventVideo.from_dict(obj.get("video")) if obj.get("video") is not None else None,
            "html": obj.get("html"),
            "text": obj.get("text"),
            "chart": V1EventChart.from_dict(obj.get("chart")) if obj.get("chart") is not None else None,
            "model": V1EventModel.from_dict(obj.get("model")) if obj.get("model") is not None else None,
            "artifact": V1EventArtifact.from_dict(obj.get("artifact")) if obj.get("artifact") is not None else None,
            "dataframe": V1EventDataframe.from_dict(obj.get("dataframe")) if obj.get("dataframe") is not None else None,
            "curve": V1EventCurve.from_dict(obj.get("curve")) if obj.get("curve") is not None else None,
            "confusion": V1EventConfusionMatrix.from_dict(obj.get("confusion")) if obj.get("confusion") is not None else None
        })
        return _obj

