// Package sampolicies allows you to choose from a list of AWS SAM policy templates to scope
// the permissions of your Lambda functions to the resources that are used by your application.
package sampolicies

import (
	"errors"
	"fmt"
	"strings"
)

const (
	iamTemplate = `{ "Version": "2012-10-17", "Statement": [ %s ] }`
)

const (
	// AccountIDMissingErr is returned when GetPolicyStatement is executed without an accountID
	AccountIDMissingErr = "factory is missing required variable accountID"

	// PartitionMissingErr is returned when GetPolicyStatement is executed without a partition
	PartitionMissingErr = "factory is missing required variable partition"

	// RegionMissingErr is returned when GetPolicyStatement is executed without a region
	RegionMissingErr = "factory is missing required variable region"
)

// Factory is the main struct to create all new policies.
// It also has methods to get the IAM statement and add new
// policies to the array.
type Factory struct {
	policies  []string
	partition string
	region    string
	accountID string
}

// NewFactory returns a new Factory pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewFactory() *Factory {
	return &Factory{}
}

// WithPartition sets the AWS partition to use and returns a pointer to the
// existing resource to allow chaining.
func (f *Factory) WithPartition(partition string) *Factory {
	f.partition = partition
	return f
}

// WithRegion sets the AWS region to use and returns a pointer to the
// existing resource to allow chaining.
func (f *Factory) WithRegion(region string) *Factory {
	f.region = region
	return f
}

// WithAccountID sets the AWS AccountID to use and returns a pointer to the
// existing resource to allow chaining.
func (f *Factory) WithAccountID(accountID string) *Factory {
	f.accountID = accountID
	return f
}

// GetPolicyStatement creates the AWS IAM policy statement by linking
// together the policies that have been added so far and substituting the
// partition, region, and accountID. If any of the fields are missing, an
// error will be thrown.
func (f *Factory) GetPolicyStatement() (string, error) {
	// Perform checks
	if len(f.accountID) == 0 {
		return "", errors.New(AccountIDMissingErr)
	}

	if len(f.region) == 0 {
		return "", errors.New(RegionMissingErr)
	}

	if len(f.partition) == 0 {
		return "", errors.New(PartitionMissingErr)
	}

	// Replace AWS placeholders
	t := fmt.Sprintf(iamTemplate, strings.Join(f.policies, ","))
	t = strings.ReplaceAll(t, "${AWS::Partition}", f.partition)
	t = strings.ReplaceAll(t, "${AWS::Region}", f.region)
	t = strings.ReplaceAll(t, "${AWS::AccountId}", f.accountID)

	// Return the policy document
	return t, nil
}
