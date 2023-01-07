#!/usr/bin/python
#
# Copyright 2018-2023 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    The version of the OpenAPI document: 1.21.0
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


try:
    from inspect import getfullargspec
except ImportError:
    from inspect import getargspec as getfullargspec
import pprint
import re  # noqa: F401
import six

from polyaxon_sdk.configuration import Configuration


class V1Mapping(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'kind': 'str',
        'values': 'list[object]',
        'concurrency': 'int',
        'early_stopping': 'list[object]'
    }

    attribute_map = {
        'kind': 'kind',
        'values': 'values',
        'concurrency': 'concurrency',
        'early_stopping': 'earlyStopping'
    }

    def __init__(self, kind='mapping', values=None, concurrency=None, early_stopping=None, local_vars_configuration=None):  # noqa: E501
        """V1Mapping - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._kind = None
        self._values = None
        self._concurrency = None
        self._early_stopping = None
        self.discriminator = None

        if kind is not None:
            self.kind = kind
        if values is not None:
            self.values = values
        if concurrency is not None:
            self.concurrency = concurrency
        if early_stopping is not None:
            self.early_stopping = early_stopping

    @property
    def kind(self):
        """Gets the kind of this V1Mapping.  # noqa: E501


        :return: The kind of this V1Mapping.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1Mapping.


        :param kind: The kind of this V1Mapping.  # noqa: E501
        :type kind: str
        """

        self._kind = kind

    @property
    def values(self):
        """Gets the values of this V1Mapping.  # noqa: E501


        :return: The values of this V1Mapping.  # noqa: E501
        :rtype: list[object]
        """
        return self._values

    @values.setter
    def values(self, values):
        """Sets the values of this V1Mapping.


        :param values: The values of this V1Mapping.  # noqa: E501
        :type values: list[object]
        """

        self._values = values

    @property
    def concurrency(self):
        """Gets the concurrency of this V1Mapping.  # noqa: E501


        :return: The concurrency of this V1Mapping.  # noqa: E501
        :rtype: int
        """
        return self._concurrency

    @concurrency.setter
    def concurrency(self, concurrency):
        """Sets the concurrency of this V1Mapping.


        :param concurrency: The concurrency of this V1Mapping.  # noqa: E501
        :type concurrency: int
        """

        self._concurrency = concurrency

    @property
    def early_stopping(self):
        """Gets the early_stopping of this V1Mapping.  # noqa: E501


        :return: The early_stopping of this V1Mapping.  # noqa: E501
        :rtype: list[object]
        """
        return self._early_stopping

    @early_stopping.setter
    def early_stopping(self, early_stopping):
        """Sets the early_stopping of this V1Mapping.


        :param early_stopping: The early_stopping of this V1Mapping.  # noqa: E501
        :type early_stopping: list[object]
        """

        self._early_stopping = early_stopping

    def to_dict(self, serialize=False):
        """Returns the model properties as a dict"""
        result = {}

        def convert(x):
            if hasattr(x, "to_dict"):
                args = getfullargspec(x.to_dict).args
                if len(args) == 1:
                    return x.to_dict()
                else:
                    return x.to_dict(serialize)
            else:
                return x

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            attr = self.attribute_map.get(attr, attr) if serialize else attr
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: convert(x),
                    value
                ))
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], convert(item[1])),
                    value.items()
                ))
            else:
                result[attr] = convert(value)

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1Mapping):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Mapping):
            return True

        return self.to_dict() != other.to_dict()
