# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.4.2
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg


class V1MatrixKind(str, Enum):
    """
    V1MatrixKind
    """

    """
    allowed enum values
    """
    RANDOM = 'random'
    GRID = 'grid'
    HYPERBAND = 'hyperband'
    BAYES = 'bayes'
    HYPEROPT = 'hyperopt'
    ITERATIVE = 'iterative'
    MAPPING = 'mapping'
