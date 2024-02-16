"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

import setuptools

try:
    with open("README.md", "r") as fh:
        long_description = fh.read()
except FileNotFoundError:
    long_description = ""

setuptools.setup(
    name="triple",
    version="0.39.8",
    author="aryan-test",
    description="Python Client SDK Generated by Speakeasy",
    long_description=long_description,
    long_description_content_type="text/markdown",
    packages=setuptools.find_packages(where="src"),
    install_requires=[
        "certifi>=2023.7.22",
        "charset-normalizer>=3.2.0",
        "dataclasses-json-speakeasy>=0.5.11",
        "idna>=3.4",
        "jsonpath-python>=1.0.6 ",
        "marshmallow>=3.19.0",
        "mypy-extensions>=1.0.0",
        "packaging>=23.1",
        "python-dateutil>=2.8.2",
        "requests>=2.31.0",
        "six>=1.16.0",
        "typing-inspect>=0.9.0",
        "typing_extensions>=4.7.1",
        "urllib3>=1.26.18",
    ],
    extras_require={
        "dev":["pylint==2.16.2"]
    },
    package_dir={'': 'src'},
    python_requires='>=3.8',
    package_data={"triple": ["py.typed"]},
)
