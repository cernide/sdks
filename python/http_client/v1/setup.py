# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

         # noqa: E501

    The version of the OpenAPI document: 2.2.0
    Contact: contact@polyaxon.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from setuptools import setup, find_packages  # noqa: H301

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools
NAME = "polyaxon-sdk"
VERSION = "2.2.0"
PYTHON_REQUIRES = ">=3.7"
REQUIRES = [
    "urllib3 >= 1.25.3",
    "python-dateutil",
    "pydantic",
    "aenum"
]

setup(
    name=NAME,
    version=VERSION,
    description="Polyaxon SDKs and REST API specification.",
    author="Polyaxon sdk",
    author_email="contact@polyaxon.com",
    url="",
    keywords=["OpenAPI", "OpenAPI-Generator", "Polyaxon SDKs and REST API specification."],
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    long_description_content_type='text/markdown',
    long_description="""\
         # noqa: E501
    """
)
