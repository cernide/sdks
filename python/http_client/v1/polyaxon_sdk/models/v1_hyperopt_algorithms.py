# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.6-rc0
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1HyperoptAlgorithms(str, Enum):
    """
    - tpe: tpe algorithm  - rand: random algorithm  - anneal: anneal algorithm
    """

    """
    allowed enum values
    """
    TPE = 'tpe'
    RAND = 'rand'
    ANNEAL = 'anneal'

