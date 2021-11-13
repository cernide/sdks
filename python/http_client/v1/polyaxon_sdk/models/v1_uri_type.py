#!/usr/bin/python
#
# Copyright 2018-2021 Polyaxon, Inc.
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

    The version of the OpenAPI document: 1.13.0-rc2
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


class V1UriType(object):
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
        'user': 'str',
        'password': 'str',
        'host': 'bool'
    }

    attribute_map = {
        'user': 'user',
        'password': 'password',
        'host': 'host'
    }

    def __init__(self, user=None, password=None, host=None, local_vars_configuration=None):  # noqa: E501
        """V1UriType - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._user = None
        self._password = None
        self._host = None
        self.discriminator = None

        if user is not None:
            self.user = user
        if password is not None:
            self.password = password
        if host is not None:
            self.host = host

    @property
    def user(self):
        """Gets the user of this V1UriType.  # noqa: E501


        :return: The user of this V1UriType.  # noqa: E501
        :rtype: str
        """
        return self._user

    @user.setter
    def user(self, user):
        """Sets the user of this V1UriType.


        :param user: The user of this V1UriType.  # noqa: E501
        :type user: str
        """

        self._user = user

    @property
    def password(self):
        """Gets the password of this V1UriType.  # noqa: E501


        :return: The password of this V1UriType.  # noqa: E501
        :rtype: str
        """
        return self._password

    @password.setter
    def password(self, password):
        """Sets the password of this V1UriType.


        :param password: The password of this V1UriType.  # noqa: E501
        :type password: str
        """

        self._password = password

    @property
    def host(self):
        """Gets the host of this V1UriType.  # noqa: E501


        :return: The host of this V1UriType.  # noqa: E501
        :rtype: bool
        """
        return self._host

    @host.setter
    def host(self, host):
        """Sets the host of this V1UriType.


        :param host: The host of this V1UriType.  # noqa: E501
        :type host: bool
        """

        self._host = host

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
        if not isinstance(other, V1UriType):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1UriType):
            return True

        return self.to_dict() != other.to_dict()
