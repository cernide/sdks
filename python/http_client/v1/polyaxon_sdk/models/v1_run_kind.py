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





class V1RunKind(str, Enum):
    """
    V1RunKind
    """

    """
    allowed enum values
    """
    JOB = 'job'
    SERVICE = 'service'
    DAG = 'dag'
    SPARKJOB = 'sparkjob'
    DASKJOB = 'daskjob'
    RAYJOB = 'rayjob'
    MPIJOB = 'mpijob'
    TFJOB = 'tfjob'
    PYTORCHJOB = 'pytorchjob'
    PADDLEJOB = 'paddlejob'
    MXJOB = 'mxjob'
    XGBJOB = 'xgbjob'
    MATRIX = 'matrix'
    SCHEDULE = 'schedule'
    TUNER = 'tuner'
    WATCHDOG = 'watchdog'
    NOTIFIER = 'notifier'
    CLEANER = 'cleaner'
    BUILDER = 'builder'

