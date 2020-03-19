// Package gateway interacts with the AWS API Gateway resources built
// with the Pulumi SDK
package gateway

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// MustGetGatewayResource is like GetGatewayResource but will panic when the resource is not found, or an error occurs
func MustGetGatewayResource(ctx *pulumi.Context, restID string, path string) *apigateway.LookupResourceResult {
	resource, err := GetGatewayResource(ctx, restID, path)
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
func GetGatewayResource(ctx *pulumi.Context, restID string, path string) (*apigateway.LookupResourceResult, error) {
	return apigateway.LookupResource(ctx, &apigateway.LookupResourceArgs{
		RestApiId: restID,
		Path:      path,
	})
}
