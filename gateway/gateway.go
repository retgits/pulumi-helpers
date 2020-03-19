// Package gateway interacts with the AWS API Gateway resources built
// with the Pulumi SDK
package gateway

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// MustGetGatewayResource is like GetGatewayResource but will panic when the resource is not found, or an error occurs
func MustGetGatewayResource(ctx *pulumi.Context, gw *apigateway.RestApi, path string) *apigateway.LookupResourceResult {
	resource, err := GetGatewayResource(ctx, gw, path)
	if err != nil {
		panic(err)
	}

	if resource == nil {
		panic(fmt.Errorf("resource %s does not exist in gateway", path))
	}

	return resource
}

// GetGatewayResource gets the API resource based on the path. If no resource can be found
// an error is thrown.
func GetGatewayResource(ctx *pulumi.Context, gw *apigateway.RestApi, path string) (*apigateway.LookupResourceResult, error) {
	// Return a mock result for preview and dry run
	if ctx.DryRun() {
		return &apigateway.LookupResourceResult{
			Id: "1",
		}, nil
	}

	c := make(chan string)
	gw.ID().ToStringOutput().ApplyString(func(id string) string {
		c <- id
		return id
	})

	restID := <-c

	return apigateway.LookupResource(ctx, &apigateway.LookupResourceArgs{
		RestApiId: restID,
		Path:      path,
	})
}
