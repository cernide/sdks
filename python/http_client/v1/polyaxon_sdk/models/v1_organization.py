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

    The version of the OpenAPI document: 1.13.0-rc3
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


class V1Organization(object):
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
        'user_email': 'str',
        'name': 'str',
        'is_public': 'bool',
        'created_at': 'datetime',
        'updated_at': 'datetime',
        'support_revoke_at': 'datetime',
        'expiration': 'int',
        'role': 'str',
        'queue': 'str',
        'preset': 'str',
        'is_cloud_viewable': 'bool',
        'auth': 'object',
        'plan': 'object',
        'usage': 'object'
    }

    attribute_map = {
        'user': 'user',
        'user_email': 'user_email',
        'name': 'name',
        'is_public': 'is_public',
        'created_at': 'created_at',
        'updated_at': 'updated_at',
        'support_revoke_at': 'support_revoke_at',
        'expiration': 'expiration',
        'role': 'role',
        'queue': 'queue',
        'preset': 'preset',
        'is_cloud_viewable': 'is_cloud_viewable',
        'auth': 'auth',
        'plan': 'plan',
        'usage': 'usage'
    }

    def __init__(self, user=None, user_email=None, name=None, is_public=None, created_at=None, updated_at=None, support_revoke_at=None, expiration=None, role=None, queue=None, preset=None, is_cloud_viewable=None, auth=None, plan=None, usage=None, local_vars_configuration=None):  # noqa: E501
        """V1Organization - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._user = None
        self._user_email = None
        self._name = None
        self._is_public = None
        self._created_at = None
        self._updated_at = None
        self._support_revoke_at = None
        self._expiration = None
        self._role = None
        self._queue = None
        self._preset = None
        self._is_cloud_viewable = None
        self._auth = None
        self._plan = None
        self._usage = None
        self.discriminator = None

        if user is not None:
            self.user = user
        if user_email is not None:
            self.user_email = user_email
        if name is not None:
            self.name = name
        if is_public is not None:
            self.is_public = is_public
        if created_at is not None:
            self.created_at = created_at
        if updated_at is not None:
            self.updated_at = updated_at
        if support_revoke_at is not None:
            self.support_revoke_at = support_revoke_at
        if expiration is not None:
            self.expiration = expiration
        if role is not None:
            self.role = role
        if queue is not None:
            self.queue = queue
        if preset is not None:
            self.preset = preset
        if is_cloud_viewable is not None:
            self.is_cloud_viewable = is_cloud_viewable
        if auth is not None:
            self.auth = auth
        if plan is not None:
            self.plan = plan
        if usage is not None:
            self.usage = usage

    @property
    def user(self):
        """Gets the user of this V1Organization.  # noqa: E501


        :return: The user of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._user

    @user.setter
    def user(self, user):
        """Sets the user of this V1Organization.


        :param user: The user of this V1Organization.  # noqa: E501
        :type user: str
        """

        self._user = user

    @property
    def user_email(self):
        """Gets the user_email of this V1Organization.  # noqa: E501


        :return: The user_email of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._user_email

    @user_email.setter
    def user_email(self, user_email):
        """Sets the user_email of this V1Organization.


        :param user_email: The user_email of this V1Organization.  # noqa: E501
        :type user_email: str
        """

        self._user_email = user_email

    @property
    def name(self):
        """Gets the name of this V1Organization.  # noqa: E501


        :return: The name of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1Organization.


        :param name: The name of this V1Organization.  # noqa: E501
        :type name: str
        """

        self._name = name

    @property
    def is_public(self):
        """Gets the is_public of this V1Organization.  # noqa: E501


        :return: The is_public of this V1Organization.  # noqa: E501
        :rtype: bool
        """
        return self._is_public

    @is_public.setter
    def is_public(self, is_public):
        """Sets the is_public of this V1Organization.


        :param is_public: The is_public of this V1Organization.  # noqa: E501
        :type is_public: bool
        """

        self._is_public = is_public

    @property
    def created_at(self):
        """Gets the created_at of this V1Organization.  # noqa: E501


        :return: The created_at of this V1Organization.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this V1Organization.


        :param created_at: The created_at of this V1Organization.  # noqa: E501
        :type created_at: datetime
        """

        self._created_at = created_at

    @property
    def updated_at(self):
        """Gets the updated_at of this V1Organization.  # noqa: E501


        :return: The updated_at of this V1Organization.  # noqa: E501
        :rtype: datetime
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at):
        """Sets the updated_at of this V1Organization.


        :param updated_at: The updated_at of this V1Organization.  # noqa: E501
        :type updated_at: datetime
        """

        self._updated_at = updated_at

    @property
    def support_revoke_at(self):
        """Gets the support_revoke_at of this V1Organization.  # noqa: E501


        :return: The support_revoke_at of this V1Organization.  # noqa: E501
        :rtype: datetime
        """
        return self._support_revoke_at

    @support_revoke_at.setter
    def support_revoke_at(self, support_revoke_at):
        """Sets the support_revoke_at of this V1Organization.


        :param support_revoke_at: The support_revoke_at of this V1Organization.  # noqa: E501
        :type support_revoke_at: datetime
        """

        self._support_revoke_at = support_revoke_at

    @property
    def expiration(self):
        """Gets the expiration of this V1Organization.  # noqa: E501


        :return: The expiration of this V1Organization.  # noqa: E501
        :rtype: int
        """
        return self._expiration

    @expiration.setter
    def expiration(self, expiration):
        """Sets the expiration of this V1Organization.


        :param expiration: The expiration of this V1Organization.  # noqa: E501
        :type expiration: int
        """

        self._expiration = expiration

    @property
    def role(self):
        """Gets the role of this V1Organization.  # noqa: E501


        :return: The role of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._role

    @role.setter
    def role(self, role):
        """Sets the role of this V1Organization.


        :param role: The role of this V1Organization.  # noqa: E501
        :type role: str
        """

        self._role = role

    @property
    def queue(self):
        """Gets the queue of this V1Organization.  # noqa: E501


        :return: The queue of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._queue

    @queue.setter
    def queue(self, queue):
        """Sets the queue of this V1Organization.


        :param queue: The queue of this V1Organization.  # noqa: E501
        :type queue: str
        """

        self._queue = queue

    @property
    def preset(self):
        """Gets the preset of this V1Organization.  # noqa: E501


        :return: The preset of this V1Organization.  # noqa: E501
        :rtype: str
        """
        return self._preset

    @preset.setter
    def preset(self, preset):
        """Sets the preset of this V1Organization.


        :param preset: The preset of this V1Organization.  # noqa: E501
        :type preset: str
        """

        self._preset = preset

    @property
    def is_cloud_viewable(self):
        """Gets the is_cloud_viewable of this V1Organization.  # noqa: E501


        :return: The is_cloud_viewable of this V1Organization.  # noqa: E501
        :rtype: bool
        """
        return self._is_cloud_viewable

    @is_cloud_viewable.setter
    def is_cloud_viewable(self, is_cloud_viewable):
        """Sets the is_cloud_viewable of this V1Organization.


        :param is_cloud_viewable: The is_cloud_viewable of this V1Organization.  # noqa: E501
        :type is_cloud_viewable: bool
        """

        self._is_cloud_viewable = is_cloud_viewable

    @property
    def auth(self):
        """Gets the auth of this V1Organization.  # noqa: E501


        :return: The auth of this V1Organization.  # noqa: E501
        :rtype: object
        """
        return self._auth

    @auth.setter
    def auth(self, auth):
        """Sets the auth of this V1Organization.


        :param auth: The auth of this V1Organization.  # noqa: E501
        :type auth: object
        """

        self._auth = auth

    @property
    def plan(self):
        """Gets the plan of this V1Organization.  # noqa: E501


        :return: The plan of this V1Organization.  # noqa: E501
        :rtype: object
        """
        return self._plan

    @plan.setter
    def plan(self, plan):
        """Sets the plan of this V1Organization.


        :param plan: The plan of this V1Organization.  # noqa: E501
        :type plan: object
        """

        self._plan = plan

    @property
    def usage(self):
        """Gets the usage of this V1Organization.  # noqa: E501


        :return: The usage of this V1Organization.  # noqa: E501
        :rtype: object
        """
        return self._usage

    @usage.setter
    def usage(self, usage):
        """Sets the usage of this V1Organization.


        :param usage: The usage of this V1Organization.  # noqa: E501
        :type usage: object
        """

        self._usage = usage

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
        if not isinstance(other, V1Organization):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Organization):
            return True

        return self.to_dict() != other.to_dict()