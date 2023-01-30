# Turbo Layers for CDK

[![NPM](https://img.shields.io/npm/v/@cloudsnorkel/cdk-turbo-layers?label=npm&logo=npm)](https://www.npmjs.com/package/@cloudsnorkel/cdk-turbo-layers)
[![PyPI](https://img.shields.io/pypi/v/cloudsnorkel.cdk-turbo-layers?label=pypi&logo=pypi)](https://pypi.org/project/cloudsnorkel.cdk-turbo-layers)
[![Maven Central](https://img.shields.io/maven-central/v/com.cloudsnorkel/cdk.turbo-layers.svg?label=Maven%20Central&logo=java)](https://search.maven.org/search?q=g:%22com.cloudsnorkel%22%20AND%20a:%22cdk.turbo-layers%22)
[![Go](https://img.shields.io/github/v/tag/CloudSnorkel/cdk-turbo-layers?color=red&label=go&logo=go)](https://pkg.go.dev/github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers)
[![Nuget](https://img.shields.io/nuget/v/CloudSnorkel.Cdk.TurboLayers?color=red&&logo=nuget)](https://www.nuget.org/packages/CloudSnorkel.Cdk.TurboLayers/)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](https://github.com/CloudSnorkel/cdk-turbo-layers/blob/main/LICENSE)

Speed up deployment of Lambda functions by creating dependency layers in AWS instead of locally.

* Easily separate dependency deployment from Lambda code deployment
* Never re-package dependencies just because of a small code change
* Never download another single dependency package locally again
* Never upload oversized code packages again
* Edit your code in the browser -- no more "deployment package too large to enable inline code editing" errors
* Uninstall Docker from your laptop and extend your battery life
* Take shorter coffee breaks when deploying

Supported Lambda runtimes:

* Python
* Node.js
* Ruby

## API

The best way to browse API documentation is on [Constructs Hub](https://constructs.dev/packages/@cloudsnorkel/cdk-turbo-layers/). It is available in all supported programming languages.

## Installation

1. Confirm you're using CDK v2
2. Install the appropriate package

   1. [Python](https://pypi.org/project/cloudsnorkel.cdk-turbo-layers)

      ```
      pip install cloudsnorkel.cdk-turbo-layers
      ```
   2. [TypeScript or JavaScript](https://www.npmjs.com/package/@cloudsnorkel/cdk-turbo-layers)

      ```
      npm i @cloudsnorkel/cdk-turbo-layers
      ```
   3. [Java](https://search.maven.org/search?q=g:%22com.cloudsnorkel%22%20AND%20a:%22cdk.turbo-layers%22)

      ```xml
      <dependency>
      <groupId>com.cloudsnorkel</groupId>
      <artifactId>cdk.turbo-layers</artifactId>
      </dependency>
      ```
   4. [Go](https://pkg.go.dev/github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers)

      ```
      go get github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers
      ```
   5. [.NET](https://www.nuget.org/packages/CloudSnorkel.Cdk.TurboLayers/)

      ```
      dotnet add package CloudSnorkel.Cdk.TurboLayers
      ```

## Example

```go
 const packager = new PythonDependencyPackager(this, 'Packager', {
    runtime: lambda.Runtime.PYTHON_3_9,
    type: DependencyPackagerType.LAMBDA,
});
new Function(this, 'Function with inline requirements', {
    handler: 'index.handler',
    code: lambda.Code.fromInline('def handler(event, context):\n  import requests'),
    runtime: lambda.Runtime.PYTHON_3_9,
    // this will create a layer from with requests and Scrapy in a Lambda function instead of locally
    layers: [packager.layerFromInline('inline requirements', ['requests', 'Scrapy'])],
});
new Function(this, 'Function with external source and requirements', {
    handler: 'index.handler',
    code: lambda.Code.fromAsset('lambda-src'),
    runtime: lambda.Runtime.PYTHON_3_9,
    // this will read pyproject.toml and poetry.lock and create a layer from the requirements in a Lambda function instead of locally
    layers: [packager.layerFromInline('poetry requirements', 'lambda-src')],
});
```
