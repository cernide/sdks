# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc28
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1ManagedBy(str, Enum):
    """
    V1ManagedBy
    """

    """
    allowed enum values
    """
    CLIENT = 'client'
    CLI = 'cli'
    AGENT = 'agent'
