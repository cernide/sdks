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


class V1CompiledOperation(object):
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
        'version': 'float',
        'kind': 'str',
        'name': 'str',
        'description': 'str',
        'tags': 'list[str]',
        'presets': 'list[str]',
        'queue': 'str',
        'cache': 'V1Cache',
        'termination': 'V1Termination',
        'plugins': 'V1Plugins',
        'schedule': 'object',
        'events': 'list[V1EventTrigger]',
        'build': 'V1Build',
        'hooks': 'list[V1Hook]',
        'dependencies': 'list[str]',
        'trigger': 'V1TriggerPolicy',
        'conditions': 'str',
        'skip_on_upstream_skip': 'bool',
        'matrix': 'object',
        'joins': 'dict(str, V1Join)',
        'inputs': 'list[V1IO]',
        'outputs': 'list[V1IO]',
        'contexts': 'list[V1IO]',
        'is_approved': 'bool',
        'cost': 'float',
        'run': 'object'
    }

    attribute_map = {
        'version': 'version',
        'kind': 'kind',
        'name': 'name',
        'description': 'description',
        'tags': 'tags',
        'presets': 'presets',
        'queue': 'queue',
        'cache': 'cache',
        'termination': 'termination',
        'plugins': 'plugins',
        'schedule': 'schedule',
        'events': 'events',
        'build': 'build',
        'hooks': 'hooks',
        'dependencies': 'dependencies',
        'trigger': 'trigger',
        'conditions': 'conditions',
        'skip_on_upstream_skip': 'skipOnUpstreamSkip',
        'matrix': 'matrix',
        'joins': 'joins',
        'inputs': 'inputs',
        'outputs': 'outputs',
        'contexts': 'contexts',
        'is_approved': 'isApproved',
        'cost': 'cost',
        'run': 'run'
    }

    def __init__(self, version=None, kind=None, name=None, description=None, tags=None, presets=None, queue=None, cache=None, termination=None, plugins=None, schedule=None, events=None, build=None, hooks=None, dependencies=None, trigger=None, conditions=None, skip_on_upstream_skip=None, matrix=None, joins=None, inputs=None, outputs=None, contexts=None, is_approved=None, cost=None, run=None, local_vars_configuration=None):  # noqa: E501
        """V1CompiledOperation - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._version = None
        self._kind = None
        self._name = None
        self._description = None
        self._tags = None
        self._presets = None
        self._queue = None
        self._cache = None
        self._termination = None
        self._plugins = None
        self._schedule = None
        self._events = None
        self._build = None
        self._hooks = None
        self._dependencies = None
        self._trigger = None
        self._conditions = None
        self._skip_on_upstream_skip = None
        self._matrix = None
        self._joins = None
        self._inputs = None
        self._outputs = None
        self._contexts = None
        self._is_approved = None
        self._cost = None
        self._run = None
        self.discriminator = None

        if version is not None:
            self.version = version
        if kind is not None:
            self.kind = kind
        if name is not None:
            self.name = name
        if description is not None:
            self.description = description
        if tags is not None:
            self.tags = tags
        if presets is not None:
            self.presets = presets
        if queue is not None:
            self.queue = queue
        if cache is not None:
            self.cache = cache
        if termination is not None:
            self.termination = termination
        if plugins is not None:
            self.plugins = plugins
        if schedule is not None:
            self.schedule = schedule
        if events is not None:
            self.events = events
        if build is not None:
            self.build = build
        if hooks is not None:
            self.hooks = hooks
        if dependencies is not None:
            self.dependencies = dependencies
        if trigger is not None:
            self.trigger = trigger
        if conditions is not None:
            self.conditions = conditions
        if skip_on_upstream_skip is not None:
            self.skip_on_upstream_skip = skip_on_upstream_skip
        if matrix is not None:
            self.matrix = matrix
        if joins is not None:
            self.joins = joins
        if inputs is not None:
            self.inputs = inputs
        if outputs is not None:
            self.outputs = outputs
        if contexts is not None:
            self.contexts = contexts
        if is_approved is not None:
            self.is_approved = is_approved
        if cost is not None:
            self.cost = cost
        if run is not None:
            self.run = run

    @property
    def version(self):
        """Gets the version of this V1CompiledOperation.  # noqa: E501


        :return: The version of this V1CompiledOperation.  # noqa: E501
        :rtype: float
        """
        return self._version

    @version.setter
    def version(self, version):
        """Sets the version of this V1CompiledOperation.


        :param version: The version of this V1CompiledOperation.  # noqa: E501
        :type version: float
        """

        self._version = version

    @property
    def kind(self):
        """Gets the kind of this V1CompiledOperation.  # noqa: E501


        :return: The kind of this V1CompiledOperation.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1CompiledOperation.


        :param kind: The kind of this V1CompiledOperation.  # noqa: E501
        :type kind: str
        """

        self._kind = kind

    @property
    def name(self):
        """Gets the name of this V1CompiledOperation.  # noqa: E501


        :return: The name of this V1CompiledOperation.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1CompiledOperation.


        :param name: The name of this V1CompiledOperation.  # noqa: E501
        :type name: str
        """

        self._name = name

    @property
    def description(self):
        """Gets the description of this V1CompiledOperation.  # noqa: E501


        :return: The description of this V1CompiledOperation.  # noqa: E501
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description):
        """Sets the description of this V1CompiledOperation.


        :param description: The description of this V1CompiledOperation.  # noqa: E501
        :type description: str
        """

        self._description = description

    @property
    def tags(self):
        """Gets the tags of this V1CompiledOperation.  # noqa: E501


        :return: The tags of this V1CompiledOperation.  # noqa: E501
        :rtype: list[str]
        """
        return self._tags

    @tags.setter
    def tags(self, tags):
        """Sets the tags of this V1CompiledOperation.


        :param tags: The tags of this V1CompiledOperation.  # noqa: E501
        :type tags: list[str]
        """

        self._tags = tags

    @property
    def presets(self):
        """Gets the presets of this V1CompiledOperation.  # noqa: E501


        :return: The presets of this V1CompiledOperation.  # noqa: E501
        :rtype: list[str]
        """
        return self._presets

    @presets.setter
    def presets(self, presets):
        """Sets the presets of this V1CompiledOperation.


        :param presets: The presets of this V1CompiledOperation.  # noqa: E501
        :type presets: list[str]
        """

        self._presets = presets

    @property
    def queue(self):
        """Gets the queue of this V1CompiledOperation.  # noqa: E501


        :return: The queue of this V1CompiledOperation.  # noqa: E501
        :rtype: str
        """
        return self._queue

    @queue.setter
    def queue(self, queue):
        """Sets the queue of this V1CompiledOperation.


        :param queue: The queue of this V1CompiledOperation.  # noqa: E501
        :type queue: str
        """

        self._queue = queue

    @property
    def cache(self):
        """Gets the cache of this V1CompiledOperation.  # noqa: E501


        :return: The cache of this V1CompiledOperation.  # noqa: E501
        :rtype: V1Cache
        """
        return self._cache

    @cache.setter
    def cache(self, cache):
        """Sets the cache of this V1CompiledOperation.


        :param cache: The cache of this V1CompiledOperation.  # noqa: E501
        :type cache: V1Cache
        """

        self._cache = cache

    @property
    def termination(self):
        """Gets the termination of this V1CompiledOperation.  # noqa: E501


        :return: The termination of this V1CompiledOperation.  # noqa: E501
        :rtype: V1Termination
        """
        return self._termination

    @termination.setter
    def termination(self, termination):
        """Sets the termination of this V1CompiledOperation.


        :param termination: The termination of this V1CompiledOperation.  # noqa: E501
        :type termination: V1Termination
        """

        self._termination = termination

    @property
    def plugins(self):
        """Gets the plugins of this V1CompiledOperation.  # noqa: E501


        :return: The plugins of this V1CompiledOperation.  # noqa: E501
        :rtype: V1Plugins
        """
        return self._plugins

    @plugins.setter
    def plugins(self, plugins):
        """Sets the plugins of this V1CompiledOperation.


        :param plugins: The plugins of this V1CompiledOperation.  # noqa: E501
        :type plugins: V1Plugins
        """

        self._plugins = plugins

    @property
    def schedule(self):
        """Gets the schedule of this V1CompiledOperation.  # noqa: E501


        :return: The schedule of this V1CompiledOperation.  # noqa: E501
        :rtype: object
        """
        return self._schedule

    @schedule.setter
    def schedule(self, schedule):
        """Sets the schedule of this V1CompiledOperation.


        :param schedule: The schedule of this V1CompiledOperation.  # noqa: E501
        :type schedule: object
        """

        self._schedule = schedule

    @property
    def events(self):
        """Gets the events of this V1CompiledOperation.  # noqa: E501


        :return: The events of this V1CompiledOperation.  # noqa: E501
        :rtype: list[V1EventTrigger]
        """
        return self._events

    @events.setter
    def events(self, events):
        """Sets the events of this V1CompiledOperation.


        :param events: The events of this V1CompiledOperation.  # noqa: E501
        :type events: list[V1EventTrigger]
        """

        self._events = events

    @property
    def build(self):
        """Gets the build of this V1CompiledOperation.  # noqa: E501


        :return: The build of this V1CompiledOperation.  # noqa: E501
        :rtype: V1Build
        """
        return self._build

    @build.setter
    def build(self, build):
        """Sets the build of this V1CompiledOperation.


        :param build: The build of this V1CompiledOperation.  # noqa: E501
        :type build: V1Build
        """

        self._build = build

    @property
    def hooks(self):
        """Gets the hooks of this V1CompiledOperation.  # noqa: E501


        :return: The hooks of this V1CompiledOperation.  # noqa: E501
        :rtype: list[V1Hook]
        """
        return self._hooks

    @hooks.setter
    def hooks(self, hooks):
        """Sets the hooks of this V1CompiledOperation.


        :param hooks: The hooks of this V1CompiledOperation.  # noqa: E501
        :type hooks: list[V1Hook]
        """

        self._hooks = hooks

    @property
    def dependencies(self):
        """Gets the dependencies of this V1CompiledOperation.  # noqa: E501


        :return: The dependencies of this V1CompiledOperation.  # noqa: E501
        :rtype: list[str]
        """
        return self._dependencies

    @dependencies.setter
    def dependencies(self, dependencies):
        """Sets the dependencies of this V1CompiledOperation.


        :param dependencies: The dependencies of this V1CompiledOperation.  # noqa: E501
        :type dependencies: list[str]
        """

        self._dependencies = dependencies

    @property
    def trigger(self):
        """Gets the trigger of this V1CompiledOperation.  # noqa: E501


        :return: The trigger of this V1CompiledOperation.  # noqa: E501
        :rtype: V1TriggerPolicy
        """
        return self._trigger

    @trigger.setter
    def trigger(self, trigger):
        """Sets the trigger of this V1CompiledOperation.


        :param trigger: The trigger of this V1CompiledOperation.  # noqa: E501
        :type trigger: V1TriggerPolicy
        """

        self._trigger = trigger

    @property
    def conditions(self):
        """Gets the conditions of this V1CompiledOperation.  # noqa: E501


        :return: The conditions of this V1CompiledOperation.  # noqa: E501
        :rtype: str
        """
        return self._conditions

    @conditions.setter
    def conditions(self, conditions):
        """Sets the conditions of this V1CompiledOperation.


        :param conditions: The conditions of this V1CompiledOperation.  # noqa: E501
        :type conditions: str
        """

        self._conditions = conditions

    @property
    def skip_on_upstream_skip(self):
        """Gets the skip_on_upstream_skip of this V1CompiledOperation.  # noqa: E501


        :return: The skip_on_upstream_skip of this V1CompiledOperation.  # noqa: E501
        :rtype: bool
        """
        return self._skip_on_upstream_skip

    @skip_on_upstream_skip.setter
    def skip_on_upstream_skip(self, skip_on_upstream_skip):
        """Sets the skip_on_upstream_skip of this V1CompiledOperation.


        :param skip_on_upstream_skip: The skip_on_upstream_skip of this V1CompiledOperation.  # noqa: E501
        :type skip_on_upstream_skip: bool
        """

        self._skip_on_upstream_skip = skip_on_upstream_skip

    @property
    def matrix(self):
        """Gets the matrix of this V1CompiledOperation.  # noqa: E501


        :return: The matrix of this V1CompiledOperation.  # noqa: E501
        :rtype: object
        """
        return self._matrix

    @matrix.setter
    def matrix(self, matrix):
        """Sets the matrix of this V1CompiledOperation.


        :param matrix: The matrix of this V1CompiledOperation.  # noqa: E501
        :type matrix: object
        """

        self._matrix = matrix

    @property
    def joins(self):
        """Gets the joins of this V1CompiledOperation.  # noqa: E501


        :return: The joins of this V1CompiledOperation.  # noqa: E501
        :rtype: dict(str, V1Join)
        """
        return self._joins

    @joins.setter
    def joins(self, joins):
        """Sets the joins of this V1CompiledOperation.


        :param joins: The joins of this V1CompiledOperation.  # noqa: E501
        :type joins: dict(str, V1Join)
        """

        self._joins = joins

    @property
    def inputs(self):
        """Gets the inputs of this V1CompiledOperation.  # noqa: E501


        :return: The inputs of this V1CompiledOperation.  # noqa: E501
        :rtype: list[V1IO]
        """
        return self._inputs

    @inputs.setter
    def inputs(self, inputs):
        """Sets the inputs of this V1CompiledOperation.


        :param inputs: The inputs of this V1CompiledOperation.  # noqa: E501
        :type inputs: list[V1IO]
        """

        self._inputs = inputs

    @property
    def outputs(self):
        """Gets the outputs of this V1CompiledOperation.  # noqa: E501


        :return: The outputs of this V1CompiledOperation.  # noqa: E501
        :rtype: list[V1IO]
        """
        return self._outputs

    @outputs.setter
    def outputs(self, outputs):
        """Sets the outputs of this V1CompiledOperation.


        :param outputs: The outputs of this V1CompiledOperation.  # noqa: E501
        :type outputs: list[V1IO]
        """

        self._outputs = outputs

    @property
    def contexts(self):
        """Gets the contexts of this V1CompiledOperation.  # noqa: E501


        :return: The contexts of this V1CompiledOperation.  # noqa: E501
        :rtype: list[V1IO]
        """
        return self._contexts

    @contexts.setter
    def contexts(self, contexts):
        """Sets the contexts of this V1CompiledOperation.


        :param contexts: The contexts of this V1CompiledOperation.  # noqa: E501
        :type contexts: list[V1IO]
        """

        self._contexts = contexts

    @property
    def is_approved(self):
        """Gets the is_approved of this V1CompiledOperation.  # noqa: E501


        :return: The is_approved of this V1CompiledOperation.  # noqa: E501
        :rtype: bool
        """
        return self._is_approved

    @is_approved.setter
    def is_approved(self, is_approved):
        """Sets the is_approved of this V1CompiledOperation.


        :param is_approved: The is_approved of this V1CompiledOperation.  # noqa: E501
        :type is_approved: bool
        """

        self._is_approved = is_approved

    @property
    def cost(self):
        """Gets the cost of this V1CompiledOperation.  # noqa: E501


        :return: The cost of this V1CompiledOperation.  # noqa: E501
        :rtype: float
        """
        return self._cost

    @cost.setter
    def cost(self, cost):
        """Sets the cost of this V1CompiledOperation.


        :param cost: The cost of this V1CompiledOperation.  # noqa: E501
        :type cost: float
        """

        self._cost = cost

    @property
    def run(self):
        """Gets the run of this V1CompiledOperation.  # noqa: E501


        :return: The run of this V1CompiledOperation.  # noqa: E501
        :rtype: object
        """
        return self._run

    @run.setter
    def run(self, run):
        """Sets the run of this V1CompiledOperation.


        :param run: The run of this V1CompiledOperation.  # noqa: E501
        :type run: object
        """

        self._run = run

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
        if not isinstance(other, V1CompiledOperation):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1CompiledOperation):
            return True

        return self.to_dict() != other.to_dict()
