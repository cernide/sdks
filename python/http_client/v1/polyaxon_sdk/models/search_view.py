# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.0-rc14
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class SearchView(str, Enum):
    """
    - any: Any view  - runs: Runs view  - selection: Selection view  - analytics: Analytics view  - components: Components view  - models: Models view  - artifacts: Artifacts view  - projects: Projects view
    """

    """
    allowed enum values
    """
    ANY = 'any'
    RUNS = 'runs'
    SELECTION = 'selection'
    ANALYTICS = 'analytics'
    COMPONENTS = 'components'
    MODELS = 'models'
    ARTIFACTS = 'artifacts'
    PROJECTS = 'projects'

