# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.1.6-rc0
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1Statuses(str, Enum):
    """
    V1Statuses
    """

    """
    allowed enum values
    """
    CREATED = 'created'
    RESUMING = 'resuming'
    ON_SCHEDULE = 'on_schedule'
    COMPILED = 'compiled'
    QUEUED = 'queued'
    SCHEDULED = 'scheduled'
    STARTING = 'starting'
    RUNNING = 'running'
    PROCESSING = 'processing'
    STOPPING = 'stopping'
    FAILED = 'failed'
    STOPPED = 'stopped'
    SUCCEEDED = 'succeeded'
    SKIPPED = 'skipped'
    WARNING = 'warning'
    UNSCHEDULABLE = 'unschedulable'
    UPSTREAM_FAILED = 'upstream_failed'
    RETRYING = 'retrying'
    UNKNOWN = 'unknown'
    DONE = 'done'

