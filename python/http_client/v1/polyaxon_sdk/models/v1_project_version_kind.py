# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.1.0-rc5
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1ProjectVersionKind(str, Enum):
    """
    V1ProjectVersionKind
    """

    """
    allowed enum values
    """
    COMPONENT = 'component'
    MODEL = 'model'
    ARTIFACT = 'artifact'

