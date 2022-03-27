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

    The version of the OpenAPI document: 1.17.1
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


class V1Hook(object):
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
        'hub_ref': 'str',
        'connection': 'str',
        'trigger': 'V1Statuses',
        'conditions': 'str',
        'params': 'dict(str, V1Param)',
        'queue': 'str',
        'presets': 'list[str]',
        'disable_defaults': 'bool'
    }

    attribute_map = {
        'hub_ref': 'hubRef',
        'connection': 'connection',
        'trigger': 'trigger',
        'conditions': 'conditions',
        'params': 'params',
        'queue': 'queue',
        'presets': 'presets',
        'disable_defaults': 'disableDefaults'
    }

    def __init__(self, hub_ref=None, connection=None, trigger=None, conditions=None, params=None, queue=None, presets=None, disable_defaults=None, local_vars_configuration=None):  # noqa: E501
        """V1Hook - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._hub_ref = None
        self._connection = None
        self._trigger = None
        self._conditions = None
        self._params = None
        self._queue = None
        self._presets = None
        self._disable_defaults = None
        self.discriminator = None

        if hub_ref is not None:
            self.hub_ref = hub_ref
        if connection is not None:
            self.connection = connection
        if trigger is not None:
            self.trigger = trigger
        if conditions is not None:
            self.conditions = conditions
        if params is not None:
            self.params = params
        if queue is not None:
            self.queue = queue
        if presets is not None:
            self.presets = presets
        if disable_defaults is not None:
            self.disable_defaults = disable_defaults

    @property
    def hub_ref(self):
        """Gets the hub_ref of this V1Hook.  # noqa: E501


        :return: The hub_ref of this V1Hook.  # noqa: E501
        :rtype: str
        """
        return self._hub_ref

    @hub_ref.setter
    def hub_ref(self, hub_ref):
        """Sets the hub_ref of this V1Hook.


        :param hub_ref: The hub_ref of this V1Hook.  # noqa: E501
        :type hub_ref: str
        """

        self._hub_ref = hub_ref

    @property
    def connection(self):
        """Gets the connection of this V1Hook.  # noqa: E501


        :return: The connection of this V1Hook.  # noqa: E501
        :rtype: str
        """
        return self._connection

    @connection.setter
    def connection(self, connection):
        """Sets the connection of this V1Hook.


        :param connection: The connection of this V1Hook.  # noqa: E501
        :type connection: str
        """

        self._connection = connection

    @property
    def trigger(self):
        """Gets the trigger of this V1Hook.  # noqa: E501


        :return: The trigger of this V1Hook.  # noqa: E501
        :rtype: V1Statuses
        """
        return self._trigger

    @trigger.setter
    def trigger(self, trigger):
        """Sets the trigger of this V1Hook.


        :param trigger: The trigger of this V1Hook.  # noqa: E501
        :type trigger: V1Statuses
        """

        self._trigger = trigger

    @property
    def conditions(self):
        """Gets the conditions of this V1Hook.  # noqa: E501


        :return: The conditions of this V1Hook.  # noqa: E501
        :rtype: str
        """
        return self._conditions

    @conditions.setter
    def conditions(self, conditions):
        """Sets the conditions of this V1Hook.


        :param conditions: The conditions of this V1Hook.  # noqa: E501
        :type conditions: str
        """

        self._conditions = conditions

    @property
    def params(self):
        """Gets the params of this V1Hook.  # noqa: E501


        :return: The params of this V1Hook.  # noqa: E501
        :rtype: dict(str, V1Param)
        """
        return self._params

    @params.setter
    def params(self, params):
        """Sets the params of this V1Hook.


        :param params: The params of this V1Hook.  # noqa: E501
        :type params: dict(str, V1Param)
        """

        self._params = params

    @property
    def queue(self):
        """Gets the queue of this V1Hook.  # noqa: E501


        :return: The queue of this V1Hook.  # noqa: E501
        :rtype: str
        """
        return self._queue

    @queue.setter
    def queue(self, queue):
        """Sets the queue of this V1Hook.


        :param queue: The queue of this V1Hook.  # noqa: E501
        :type queue: str
        """

        self._queue = queue

    @property
    def presets(self):
        """Gets the presets of this V1Hook.  # noqa: E501


        :return: The presets of this V1Hook.  # noqa: E501
        :rtype: list[str]
        """
        return self._presets

    @presets.setter
    def presets(self, presets):
        """Sets the presets of this V1Hook.


        :param presets: The presets of this V1Hook.  # noqa: E501
        :type presets: list[str]
        """

        self._presets = presets

    @property
    def disable_defaults(self):
        """Gets the disable_defaults of this V1Hook.  # noqa: E501


        :return: The disable_defaults of this V1Hook.  # noqa: E501
        :rtype: bool
        """
        return self._disable_defaults

    @disable_defaults.setter
    def disable_defaults(self, disable_defaults):
        """Sets the disable_defaults of this V1Hook.


        :param disable_defaults: The disable_defaults of this V1Hook.  # noqa: E501
        :type disable_defaults: bool
        """

        self._disable_defaults = disable_defaults

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
        if not isinstance(other, V1Hook):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Hook):
            return True

        return self.to_dict() != other.to_dict()
