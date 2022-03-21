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

    The version of the OpenAPI document: 1.17.0
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


class AgentStateResponseAgentState(object):
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
        'schedules': 'object',
        'hooks': 'object',
        'watchdogs': 'object',
        'tuners': 'object',
        'queued': 'object',
        'stopping': 'object',
        'deleting': 'object',
        'apply': 'object',
        'checks': 'object',
        'full': 'bool'
    }

    attribute_map = {
        'schedules': 'schedules',
        'hooks': 'hooks',
        'watchdogs': 'watchdogs',
        'tuners': 'tuners',
        'queued': 'queued',
        'stopping': 'stopping',
        'deleting': 'deleting',
        'apply': 'apply',
        'checks': 'checks',
        'full': 'full'
    }

    def __init__(self, schedules=None, hooks=None, watchdogs=None, tuners=None, queued=None, stopping=None, deleting=None, apply=None, checks=None, full=None, local_vars_configuration=None):  # noqa: E501
        """AgentStateResponseAgentState - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._schedules = None
        self._hooks = None
        self._watchdogs = None
        self._tuners = None
        self._queued = None
        self._stopping = None
        self._deleting = None
        self._apply = None
        self._checks = None
        self._full = None
        self.discriminator = None

        if schedules is not None:
            self.schedules = schedules
        if hooks is not None:
            self.hooks = hooks
        if watchdogs is not None:
            self.watchdogs = watchdogs
        if tuners is not None:
            self.tuners = tuners
        if queued is not None:
            self.queued = queued
        if stopping is not None:
            self.stopping = stopping
        if deleting is not None:
            self.deleting = deleting
        if apply is not None:
            self.apply = apply
        if checks is not None:
            self.checks = checks
        if full is not None:
            self.full = full

    @property
    def schedules(self):
        """Gets the schedules of this AgentStateResponseAgentState.  # noqa: E501


        :return: The schedules of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._schedules

    @schedules.setter
    def schedules(self, schedules):
        """Sets the schedules of this AgentStateResponseAgentState.


        :param schedules: The schedules of this AgentStateResponseAgentState.  # noqa: E501
        :type schedules: object
        """

        self._schedules = schedules

    @property
    def hooks(self):
        """Gets the hooks of this AgentStateResponseAgentState.  # noqa: E501


        :return: The hooks of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._hooks

    @hooks.setter
    def hooks(self, hooks):
        """Sets the hooks of this AgentStateResponseAgentState.


        :param hooks: The hooks of this AgentStateResponseAgentState.  # noqa: E501
        :type hooks: object
        """

        self._hooks = hooks

    @property
    def watchdogs(self):
        """Gets the watchdogs of this AgentStateResponseAgentState.  # noqa: E501


        :return: The watchdogs of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._watchdogs

    @watchdogs.setter
    def watchdogs(self, watchdogs):
        """Sets the watchdogs of this AgentStateResponseAgentState.


        :param watchdogs: The watchdogs of this AgentStateResponseAgentState.  # noqa: E501
        :type watchdogs: object
        """

        self._watchdogs = watchdogs

    @property
    def tuners(self):
        """Gets the tuners of this AgentStateResponseAgentState.  # noqa: E501


        :return: The tuners of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._tuners

    @tuners.setter
    def tuners(self, tuners):
        """Sets the tuners of this AgentStateResponseAgentState.


        :param tuners: The tuners of this AgentStateResponseAgentState.  # noqa: E501
        :type tuners: object
        """

        self._tuners = tuners

    @property
    def queued(self):
        """Gets the queued of this AgentStateResponseAgentState.  # noqa: E501


        :return: The queued of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._queued

    @queued.setter
    def queued(self, queued):
        """Sets the queued of this AgentStateResponseAgentState.


        :param queued: The queued of this AgentStateResponseAgentState.  # noqa: E501
        :type queued: object
        """

        self._queued = queued

    @property
    def stopping(self):
        """Gets the stopping of this AgentStateResponseAgentState.  # noqa: E501


        :return: The stopping of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._stopping

    @stopping.setter
    def stopping(self, stopping):
        """Sets the stopping of this AgentStateResponseAgentState.


        :param stopping: The stopping of this AgentStateResponseAgentState.  # noqa: E501
        :type stopping: object
        """

        self._stopping = stopping

    @property
    def deleting(self):
        """Gets the deleting of this AgentStateResponseAgentState.  # noqa: E501


        :return: The deleting of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._deleting

    @deleting.setter
    def deleting(self, deleting):
        """Sets the deleting of this AgentStateResponseAgentState.


        :param deleting: The deleting of this AgentStateResponseAgentState.  # noqa: E501
        :type deleting: object
        """

        self._deleting = deleting

    @property
    def apply(self):
        """Gets the apply of this AgentStateResponseAgentState.  # noqa: E501


        :return: The apply of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._apply

    @apply.setter
    def apply(self, apply):
        """Sets the apply of this AgentStateResponseAgentState.


        :param apply: The apply of this AgentStateResponseAgentState.  # noqa: E501
        :type apply: object
        """

        self._apply = apply

    @property
    def checks(self):
        """Gets the checks of this AgentStateResponseAgentState.  # noqa: E501


        :return: The checks of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: object
        """
        return self._checks

    @checks.setter
    def checks(self, checks):
        """Sets the checks of this AgentStateResponseAgentState.


        :param checks: The checks of this AgentStateResponseAgentState.  # noqa: E501
        :type checks: object
        """

        self._checks = checks

    @property
    def full(self):
        """Gets the full of this AgentStateResponseAgentState.  # noqa: E501


        :return: The full of this AgentStateResponseAgentState.  # noqa: E501
        :rtype: bool
        """
        return self._full

    @full.setter
    def full(self, full):
        """Sets the full of this AgentStateResponseAgentState.


        :param full: The full of this AgentStateResponseAgentState.  # noqa: E501
        :type full: bool
        """

        self._full = full

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
        if not isinstance(other, AgentStateResponseAgentState):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, AgentStateResponseAgentState):
            return True

        return self.to_dict() != other.to_dict()
