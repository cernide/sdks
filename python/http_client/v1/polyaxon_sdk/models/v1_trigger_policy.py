# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc21
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1TriggerPolicy(str, Enum):
    """
    V1TriggerPolicy
    """

    """
    allowed enum values
    """
    ALL_SUCCEEDED = 'all_succeeded'
    ALL_FAILED = 'all_failed'
    ALL_DONE = 'all_done'
    ONE_SUCCEEDED = 'one_succeeded'
    ONE_FAILED = 'one_failed'
    ONE_DONE = 'one_done'

