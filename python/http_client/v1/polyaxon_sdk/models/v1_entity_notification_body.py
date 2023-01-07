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


class V1EntityNotificationBody(object):
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
        'namespace': 'str',
        'owner': 'str',
        'project': 'str',
        'uuid': 'str',
        'name': 'str',
        'condition': 'V1StatusCondition',
        'connections': 'list[str]'
    }

    attribute_map = {
        'namespace': 'namespace',
        'owner': 'owner',
        'project': 'project',
        'uuid': 'uuid',
        'name': 'name',
        'condition': 'condition',
        'connections': 'connections'
    }

    def __init__(self, namespace=None, owner=None, project=None, uuid=None, name=None, condition=None, connections=None, local_vars_configuration=None):  # noqa: E501
        """V1EntityNotificationBody - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._namespace = None
        self._owner = None
        self._project = None
        self._uuid = None
        self._name = None
        self._condition = None
        self._connections = None
        self.discriminator = None

        if namespace is not None:
            self.namespace = namespace
        if owner is not None:
            self.owner = owner
        if project is not None:
            self.project = project
        if uuid is not None:
            self.uuid = uuid
        if name is not None:
            self.name = name
        if condition is not None:
            self.condition = condition
        if connections is not None:
            self.connections = connections

    @property
    def namespace(self):
        """Gets the namespace of this V1EntityNotificationBody.  # noqa: E501


        :return: The namespace of this V1EntityNotificationBody.  # noqa: E501
        :rtype: str
        """
        return self._namespace

    @namespace.setter
    def namespace(self, namespace):
        """Sets the namespace of this V1EntityNotificationBody.


        :param namespace: The namespace of this V1EntityNotificationBody.  # noqa: E501
        :type namespace: str
        """

        self._namespace = namespace

    @property
    def owner(self):
        """Gets the owner of this V1EntityNotificationBody.  # noqa: E501


        :return: The owner of this V1EntityNotificationBody.  # noqa: E501
        :rtype: str
        """
        return self._owner

    @owner.setter
    def owner(self, owner):
        """Sets the owner of this V1EntityNotificationBody.


        :param owner: The owner of this V1EntityNotificationBody.  # noqa: E501
        :type owner: str
        """

        self._owner = owner

    @property
    def project(self):
        """Gets the project of this V1EntityNotificationBody.  # noqa: E501


        :return: The project of this V1EntityNotificationBody.  # noqa: E501
        :rtype: str
        """
        return self._project

    @project.setter
    def project(self, project):
        """Sets the project of this V1EntityNotificationBody.


        :param project: The project of this V1EntityNotificationBody.  # noqa: E501
        :type project: str
        """

        self._project = project

    @property
    def uuid(self):
        """Gets the uuid of this V1EntityNotificationBody.  # noqa: E501


        :return: The uuid of this V1EntityNotificationBody.  # noqa: E501
        :rtype: str
        """
        return self._uuid

    @uuid.setter
    def uuid(self, uuid):
        """Sets the uuid of this V1EntityNotificationBody.


        :param uuid: The uuid of this V1EntityNotificationBody.  # noqa: E501
        :type uuid: str
        """

        self._uuid = uuid

    @property
    def name(self):
        """Gets the name of this V1EntityNotificationBody.  # noqa: E501


        :return: The name of this V1EntityNotificationBody.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1EntityNotificationBody.


        :param name: The name of this V1EntityNotificationBody.  # noqa: E501
        :type name: str
        """

        self._name = name

    @property
    def condition(self):
        """Gets the condition of this V1EntityNotificationBody.  # noqa: E501


        :return: The condition of this V1EntityNotificationBody.  # noqa: E501
        :rtype: V1StatusCondition
        """
        return self._condition

    @condition.setter
    def condition(self, condition):
        """Sets the condition of this V1EntityNotificationBody.


        :param condition: The condition of this V1EntityNotificationBody.  # noqa: E501
        :type condition: V1StatusCondition
        """

        self._condition = condition

    @property
    def connections(self):
        """Gets the connections of this V1EntityNotificationBody.  # noqa: E501


        :return: The connections of this V1EntityNotificationBody.  # noqa: E501
        :rtype: list[str]
        """
        return self._connections

    @connections.setter
    def connections(self, connections):
        """Sets the connections of this V1EntityNotificationBody.


        :param connections: The connections of this V1EntityNotificationBody.  # noqa: E501
        :type connections: list[str]
        """

        self._connections = connections

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
        if not isinstance(other, V1EntityNotificationBody):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1EntityNotificationBody):
            return True

        return self.to_dict() != other.to_dict()
