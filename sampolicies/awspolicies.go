package sampolicies

// AssumeRoleLambda returns an IAM policy document that allows the IAM role to be assumed by AWS Lambda
func AssumeRoleLambda() string {
	return `{ "Version": "2012-10-17", "Statement": [ { "Action": "sts:AssumeRole", "Principal": { "Service": "lambda.amazonaws.com" }, "Effect": "Allow" } ] }`
}
