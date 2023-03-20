# Turbo Layers for CDK

[![NPM](https://img.shields.io/npm/v/@cloudsnorkel/cdk-turbo-layers?label=npm&logo=npm)](https://www.npmjs.com/package/@cloudsnorkel/cdk-turbo-layers)
[![PyPI](https://img.shields.io/pypi/v/cloudsnorkel.cdk-turbo-layers?label=pypi&logo=pypi)](https://pypi.org/project/cloudsnorkel.cdk-turbo-layers)
[![Maven Central](https://img.shields.io/maven-central/v/com.cloudsnorkel/cdk.turbo-layers.svg?label=Maven%20Central&logo=java)](https://search.maven.org/search?q=g:%22com.cloudsnorkel%22%20AND%20a:%22cdk.turbo-layers%22)
[![Go](https://img.shields.io/github/v/tag/CloudSnorkel/cdk-turbo-layers?color=red&label=go&logo=go)](https://pkg.go.dev/github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers)
[![Nuget](https://img.shields.io/nuget/v/CloudSnorkel.Cdk.TurboLayers?color=red&&logo=nuget)](https://www.nuget.org/packages/CloudSnorkel.Cdk.TurboLayers/)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](https://github.com/CloudSnorkel/cdk-turbo-layers/blob/main/LICENSE)

Speed up deployment of Lambda functions by creating dependency layers in AWS instead of locally.

* ⛓️ Easily separate dependency deployment from Lambda code deployment
* 🔁 Never re-package dependencies just because of a small code change
* ☁️ Never download another single dependency package locally again
* 🏋️ Never upload oversized code packages again
* 🌎 Edit your code in the browser -- no more "deployment package too large to enable inline code editing"
* ❌ Uninstall Docker from your laptop and extend your battery life
* ☕ Take shorter coffee breaks when deploying

Supported Lambda runtimes:

* 🐍 Python
* 📜 Node.js
* 💎 Ruby
* ☕ Java

## Benchmark

Below are synth and deploy times for a simple Python function with [`PythonFunction`](https://docs.aws.amazon.com/cdk/api/v2/docs/@aws-cdk_aws-lambda-python-alpha.PythonFunction.html) compared to Turbo Layers. The [benchmark](benchmark/deployment.ts) ran three times and the best time were taken for each step.

|                        | 💤 PythonFunction   | 🚀 Turbo Layers  | 💤 5x PythonFunction | 🚀 5x Functions w/ Shared Turbo Layer |
|------------------------|---------------------|------------------|----------------------|---------------------------------------|
| Initial Synth          | 1:21                | 0:06             | 2:43                 | 0:06                                  |
| Initial Deploy         | 1:18                | 2:05             | 2:10                 | 2:06                                  |
| Code Change Synth      | 0:31                | 0:06             | 1:21                 | 0:06                                  |
| Code Change Deploy     | 0:49                | 0:29             | 1:19                 | 0:36                                  |
| New Dependency Synth   | 0:33                | 0:06             | 1:30                 | 0:06                                  |
| New Dependency Deploy  | 0:52                | 1:50             | 1:31                 | 1:50                                  |

As you can see, code changes synth much faster and deploy a bit faster too. Dependency changes take longer to deploy, but are assumed to be way less frequent than code changes. The more dependencies your function uses, the better the results will be.

To run the benchmark yourself use:

```
npm run bundle && npm run benchmark
```

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
    layers: [packager.layerFromPoetry('poetry requirements', 'lambda-src')],
});
```

## Older Implementations

* [lovage](https://github.com/CloudSnorkel/lovage): standalone Python framework that uses the same trick to deploy decorated functions to AWS
* [serverless-pydeps](https://github.com/CloudSnorkel/serverless-pydeps): plugin for [Serverless Framework](https://www.serverless.com/) that speeds up deployment
