package sampolicies

// AddExecuteAPI allows the IAM role to execute API invocations
func (f *Factory) AddExecuteAPI() {
	policy := `{ "Action": "execute-api:Invoke", "Resource": "execute-api:/*", "Principal": "*", "Effect": "Allow" }`
	f.policies = append(f.policies, policy)
}

// AddAssumeRoleLambda allows AWS Lambda to assume the role and use AWS services
func (f *Factory) AddAssumeRoleLambda() {
	policy := `{ "Action": "sts:AssumeRole", "Principal": { "Service": "lambda.amazonaws.com" }, "Effect": "Allow" }`
	f.policies = append(f.policies, policy)
}

// AssumeRoleLambda returns an IAM policy document that allows the IAM role to be assumed by AWS Lambda
func AssumeRoleLambda() string {
	return `{ "Version": "2012-10-17", "Statement": [ { "Action": "sts:AssumeRole", "Principal": { "Service": "lambda.amazonaws.com" }, "Effect": "Allow" } ] }`
}
