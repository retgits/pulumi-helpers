# pulumi-helpers

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/retgits/pulumi-helpers?tab=doc)

## Why build this

[Pulumi](pulumi.com) is **Modern Infrastructure as Code**. With the Pulumi Platform, you can declare cloud infrastructure using real programming languages, like Go. Using Go, rather than needing to learn yet-another YAML or DSL dialect enables modern approaches to cloud applications and infrastructure as code. Such an approach allows developers to master the Go toolchain and use their set of frameworks, and go to any cloud — AWS, Azure, GCP, or Kubernetes — with ease. However, in those YAML and DSL dialects, there are definitely a few awesome helper methods that allow developers to save a little time. This Go library has a collection of those helper methods to allow developers to use Pulumi, while also saving developers a bit of time.

## AWS SAM Policy Templates

AWS SAM allows you to choose from a list of policy templates to scope the permissions of your Lambda functions to the resources that are used by your application. The `sampolicy` package wraps the [policy templates](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-policy-template-list.html) in a format to be used with Pulumi (or virtually any other type of Go app where you need AWS IAM policy templates).

### Usage

To use the templates, you need to import the `sampolicies` package, create a new `Factory`, and add templates to factory.

```go
package main

import (
	"fmt"
	// Import the sampolicies package
	"github.com/retgits/pulumi-helpers/sampolicies"
)

func main() {
	// Create a new factory with your AWS Account ID, AWS Partition, and the region you want to create policies for
	iamFactory := sampolicies.NewFactory().WithAccountID("01234567890").WithPartition("aws").WithRegion("us-west-2")

	// Add policies to the factory
	iamFactory.AddAthenaQueryPolicy()

	// Create the policy document (this method returns an error if AWS Account ID, AWS Partition, or region is not set)
	policy, _ := iamFactory.GetPolicyStatement()

	// Print the policy statement
	fmt.Println(policy)

	// Remove all policies so you can begin with a clean slate
	iamFactory.ClearPolicies()

	// Add a new policy to the factory
	iamFactory.AddAWSSecretsManagerRotationPolicy()

	// Create the policy document
	policy, _ = iamFactory.GetPolicyStatement()
	
	// Print the policy statement
	fmt.Println(policy)
}

```

## Generating policies

The policies are generated from the develop branch of the [Serverless Application Model](https://github.com/awslabs/serverless-application-model) repository, using the [generator](./cmd/generator.go). This little app will create the [policies.go](./policies.go) file and tell which policies need to be updated manually because they resulted in an error.

Right now, there are nine (9) policies that result in some error processing:

* DynamoDBCrudPolicy
* DynamoDBWritePolicy
* S3FullAccessPolicy
* S3WritePolicy
* S3CrudPolicy
* S3ReadPolicy
* PollyFullAccessPolicy
* ServerlessRepoReadWriteAccessPolicy
* DynamoDBReadPolicy

## Contributing

Pull requests are welcome in their individual repositories. For major changes or questions, please open [an issue](https://github.com/retgits/go-sam-policy-templates/issues) first to discuss what you would like to change.

## License

See the [LICENSE](./LICENSE) file in the repository.
