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


class V1PasswordChange(object):
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
        'old_password': 'str',
        'new_password1': 'str',
        'new_password2': 'str'
    }

    attribute_map = {
        'old_password': 'old_password',
        'new_password1': 'new_password1',
        'new_password2': 'new_password2'
    }

    def __init__(self, old_password=None, new_password1=None, new_password2=None, local_vars_configuration=None):  # noqa: E501
        """V1PasswordChange - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._old_password = None
        self._new_password1 = None
        self._new_password2 = None
        self.discriminator = None

        if old_password is not None:
            self.old_password = old_password
        if new_password1 is not None:
            self.new_password1 = new_password1
        if new_password2 is not None:
            self.new_password2 = new_password2

    @property
    def old_password(self):
        """Gets the old_password of this V1PasswordChange.  # noqa: E501


        :return: The old_password of this V1PasswordChange.  # noqa: E501
        :rtype: str
        """
        return self._old_password

    @old_password.setter
    def old_password(self, old_password):
        """Sets the old_password of this V1PasswordChange.


        :param old_password: The old_password of this V1PasswordChange.  # noqa: E501
        :type old_password: str
        """

        self._old_password = old_password

    @property
    def new_password1(self):
        """Gets the new_password1 of this V1PasswordChange.  # noqa: E501


        :return: The new_password1 of this V1PasswordChange.  # noqa: E501
        :rtype: str
        """
        return self._new_password1

    @new_password1.setter
    def new_password1(self, new_password1):
        """Sets the new_password1 of this V1PasswordChange.


        :param new_password1: The new_password1 of this V1PasswordChange.  # noqa: E501
        :type new_password1: str
        """

        self._new_password1 = new_password1

    @property
    def new_password2(self):
        """Gets the new_password2 of this V1PasswordChange.  # noqa: E501


        :return: The new_password2 of this V1PasswordChange.  # noqa: E501
        :rtype: str
        """
        return self._new_password2

    @new_password2.setter
    def new_password2(self, new_password2):
        """Sets the new_password2 of this V1PasswordChange.


        :param new_password2: The new_password2 of this V1PasswordChange.  # noqa: E501
        :type new_password2: str
        """

        self._new_password2 = new_password2

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
        if not isinstance(other, V1PasswordChange):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1PasswordChange):
            return True

        return self.to_dict() != other.to_dict()
