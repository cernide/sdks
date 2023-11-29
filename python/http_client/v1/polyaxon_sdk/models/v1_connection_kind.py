# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.0.2
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from inspect import getfullargspec
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class V1ConnectionKind(str, Enum):
    """
    V1ConnectionKind
    """

    """
    allowed enum values
    """
    HOST_PATH = 'host_path'
    VOLUME_CLAIM = 'volume_claim'
    GCS = 'gcs'
    S3 = 's3'
    WASB = 'wasb'
    REGISTRY = 'registry'
    GIT = 'git'
    AWS = 'aws'
    GCP = 'gcp'
    AZURE = 'azure'
    MYSQL = 'mysql'
    POSTGRES = 'postgres'
    ORACLE = 'oracle'
    VERTICA = 'vertica'
    SQLITE = 'sqlite'
    MSSQL = 'mssql'
    REDIS = 'redis'
    PRESTO = 'presto'
    MONGO = 'mongo'
    CASSANDRA = 'cassandra'
    FTP = 'ftp'
    GRPC = 'grpc'
    HDFS = 'hdfs'
    HTTP = 'http'
    PIG_CLI = 'pig_cli'
    HIVE_CLI = 'hive_cli'
    HIVE_METASTORE = 'hive_metastore'
    HIVE_SERVER2 = 'hive_server2'
    JDBC = 'jdbc'
    JENKINS = 'jenkins'
    SAMBA = 'samba'
    SNOWFLAKE = 'snowflake'
    SSH = 'ssh'
    CLOUDANT = 'cloudant'
    DATABRICKS = 'databricks'
    SEGMENT = 'segment'
    SLACK = 'slack'
    DISCORD = 'discord'
    MATTERMOST = 'mattermost'
    PAGERDUTY = 'pagerduty'
    HIPCHAT = 'hipchat'
    WEBHOOK = 'webhook'
    CUSTOM = 'custom'

