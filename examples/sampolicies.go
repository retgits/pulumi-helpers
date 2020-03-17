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
}
