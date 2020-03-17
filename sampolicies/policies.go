package sampolicies

import "strings"

// AddLambdaInvokePolicy Gives permission to invoke a Lambda Function, Alias or Version
func (f *Factory) AddLambdaInvokePolicy(functionName string) {
	policy := `{ "Action": ["lambda:InvokeFunction"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:lambda:${AWS::Region}:${AWS::AccountId}:function:${functionName}*" }`
	policy = strings.ReplaceAll(policy, "${functionName}", functionName)
	f.policies = append(f.policies, policy)
}

// AddAMIDescribePolicy Gives permissions to describe AMIs
func (f *Factory) AddAMIDescribePolicy() {
	policy := `{ "Action": ["ec2:DescribeImages"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:ec2:${AWS::Region}:${AWS::AccountId}:image/*" }`
	f.policies = append(f.policies, policy)
}

// AddMobileAnalyticsWriteOnlyAccessPolicy Gives write only permissions to put event data for all application resources
func (f *Factory) AddMobileAnalyticsWriteOnlyAccessPolicy() {
	policy := `{ "Action": ["mobileanalytics:PutEvents"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddDynamoDBCrudPolicy Gives CRUD access to a DynamoDB Table
func (f *Factory) AddDynamoDBCrudPolicy(tableName string) {
	policy := `{ "Effect": "Allow", "Action": [ "dynamodb:GetItem", "dynamodb:DeleteItem", "dynamodb:PutItem", "dynamodb:Scan", "dynamodb:Query", "dynamodb:UpdateItem", "dynamodb:BatchWriteItem", "dynamodb:BatchGetItem", "dynamodb:DescribeTable", "dynamodb:ConditionCheckItem" ], "Resource": [ "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}", "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}/index/*" ] }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddElasticsearchHttpPostPolicy Gives POST and PUT permissions to Elasticsearch
func (f *Factory) AddElasticsearchHttpPostPolicy(domainName string) {
	policy := `{ "Action": ["es:ESHttpPost","es:ESHttpPut"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:es:${AWS::Region}:${AWS::AccountId}:domain/${domainName}/*" }`
	policy = strings.ReplaceAll(policy, "${domainName}", domainName)
	f.policies = append(f.policies, policy)
}

// AddRekognitionLabelsPolicy Gives permission to detect object and moderation labels
func (f *Factory) AddRekognitionLabelsPolicy() {
	policy := `{ "Action": ["rekognition:DetectLabels","rekognition:DetectModerationLabels"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddAWSSecretsManagerGetSecretValuePolicy Grants permissions to GetSecretValue for the specified AWS Secrets Manager secret
func (f *Factory) AddAWSSecretsManagerGetSecretValuePolicy(secretArn string) {
	policy := `{ "Action": ["secretsmanager:GetSecretValue"], "Effect": "Allow", "Resource": "${secretArn}" }`
	policy = strings.ReplaceAll(policy, "${secretArn}", secretArn)
	f.policies = append(f.policies, policy)
}

// AddRekognitionNoDataAccessPolicy Gives permission to compare and detect faces and labels
func (f *Factory) AddRekognitionNoDataAccessPolicy(collectionId string) {
	policy := `{ "Action": ["rekognition:CompareFaces","rekognition:DetectFaces","rekognition:DetectLabels","rekognition:DetectModerationLabels"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:rekognition:${AWS::Region}:${AWS::AccountId}:collection/${collectionId}" }`
	policy = strings.ReplaceAll(policy, "${collectionId}", collectionId)
	f.policies = append(f.policies, policy)
}

// AddSQSSendMessagePolicy Gives permission to send message to SQS Queue
func (f *Factory) AddSQSSendMessagePolicy(queueName string) {
	policy := `{ "Action": ["sqs:SendMessage*"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:sqs:${AWS::Region}:${AWS::AccountId}:${queueName}" }`
	policy = strings.ReplaceAll(policy, "${queueName}", queueName)
	f.policies = append(f.policies, policy)
}

// AddCodePipelineReadOnlyPolicy Gives read permissions to get details about a CodePipeline pipeline
func (f *Factory) AddCodePipelineReadOnlyPolicy(pipelinename string) {
	policy := `{ "Action": ["codepipeline:ListPipelineExecutions"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:codepipeline:${AWS::Region}:${AWS::AccountId}:${pipelinename}" }`
	policy = strings.ReplaceAll(policy, "${pipelinename}", pipelinename)
	f.policies = append(f.policies, policy)
}

// AddSESBulkTemplatedCrudPolicy Gives permission to send email, templated email, templated bulk emails and verify identity
func (f *Factory) AddSESBulkTemplatedCrudPolicy(identityName string) {
	policy := `{ "Action": ["ses:GetIdentityVerificationAttributes","ses:SendEmail","ses:SendRawEmail","ses:SendTemplatedEmail","ses:SendBulkTemplatedEmail","ses:VerifyEmailIdentity"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:ses:${AWS::Region}:${AWS::AccountId}:identity/${identityName}" }`
	policy = strings.ReplaceAll(policy, "${identityName}", identityName)
	f.policies = append(f.policies, policy)
}

// AddSQSPollerPolicy Gives permissions to poll an SQS Queue
func (f *Factory) AddSQSPollerPolicy(queueName string) {
	policy := `{ "Action": ["sqs:ChangeMessageVisibility","sqs:ChangeMessageVisibilityBatch","sqs:DeleteMessage","sqs:DeleteMessageBatch","sqs:GetQueueAttributes","sqs:ReceiveMessage"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:sqs:${AWS::Region}:${AWS::AccountId}:${queueName}" }`
	policy = strings.ReplaceAll(policy, "${queueName}", queueName)
	f.policies = append(f.policies, policy)
}

// AddRekognitionReadPolicy Gives permission to list and search faces
func (f *Factory) AddRekognitionReadPolicy(collectionId string) {
	policy := `{ "Action": ["rekognition:ListCollections","rekognition:ListFaces","rekognition:SearchFaces","rekognition:SearchFacesByImage"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:rekognition:${AWS::Region}:${AWS::AccountId}:collection/${collectionId}" }`
	policy = strings.ReplaceAll(policy, "${collectionId}", collectionId)
	f.policies = append(f.policies, policy)
}

// AddDynamoDBWritePolicy Gives write only access to a DynamoDB Table
func (f *Factory) AddDynamoDBWritePolicy(tableName string) {
	policy := `{ "Effect": "Allow", "Action": [ "dynamodb:PutItem", "dynamodb:UpdateItem", "dynamodb:BatchWriteItem" ], "Resource": [ ""arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}"", ""arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}/index/*"" ] }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddSNSPublishMessagePolicy Gives permission to publish message to SNS Topic
func (f *Factory) AddSNSPublishMessagePolicy(topicName string) {
	policy := `{ "Action": ["sns:Publish"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:sns:${AWS::Region}:${AWS::AccountId}:${topicName}" }`
	policy = strings.ReplaceAll(policy, "${topicName}", topicName)
	f.policies = append(f.policies, policy)
}

// AddKinesisStreamReadPolicy Gives permission to list and read a Kinesis stream
func (f *Factory) AddKinesisStreamReadPolicy() {
	policy := `{ "Action": ["kinesis:ListStreams","kinesis:DescribeLimits"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:kinesis:${AWS::Region}:${AWS::AccountId}:stream/*" }`
	f.policies = append(f.policies, policy)
}

// AddTextractGetResultPolicy Gives access to get detected and analyzed documents from Textract
func (f *Factory) AddTextractGetResultPolicy() {
	policy := `{ "Action": ["textract:GetDocumentTextDetection","textract:GetDocumentAnalysis"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddSESCrudPolicy Gives permission to send email and verify identity
func (f *Factory) AddSESCrudPolicy(identityName string) {
	policy := `{ "Action": ["ses:GetIdentityVerificationAttributes","ses:SendEmail","ses:SendRawEmail","ses:VerifyEmailIdentity"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:ses:${AWS::Region}:${AWS::AccountId}:identity/${identityName}" }`
	policy = strings.ReplaceAll(policy, "${identityName}", identityName)
	f.policies = append(f.policies, policy)
}

// AddS3FullAccessPolicy Gives full access permissions to objects in the S3 Bucket
func (f *Factory) AddS3FullAccessPolicy(bucketName string) {
	policy := `{ "Effect": "Allow", "Action": [ "s3:GetObject", "s3:GetObjectAcl", "s3:GetObjectVersion", "s3:PutObject", "s3:PutObjectAcl", "s3:DeleteObject", "s3:DeleteObjectTagging", "s3:DeleteObjectVersionTagging", "s3:GetObjectTagging", "s3:GetObjectVersionTagging", "s3:PutObjectTagging", "s3:PutObjectVersionTagging" ], "Resource": [ { "arn:${AWS::Partition}:s3:::${bucketName}/*", } ] }, { "Effect": "Allow", "Action": [ "s3:ListBucket", "s3:GetBucketLocation", "s3:GetLifecycleConfiguration", "s3:PutLifecycleConfiguration" ], "Resource": [ "arn:${AWS::Partition}:s3:::${bucketName}"] } ] }`
	policy = strings.ReplaceAll(policy, "${bucketName}", bucketName)
	f.policies = append(f.policies, policy)
}

// AddEC2CopyImagePolicy Gives permission top copy EC2 Images
func (f *Factory) AddEC2CopyImagePolicy(imageId string) {
	policy := `{ "Action": ["ec2:CopyImage"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:ec2:${AWS::Region}:${AWS::AccountId}:image/${imageId}" }`
	policy = strings.ReplaceAll(policy, "${imageId}", imageId)
	f.policies = append(f.policies, policy)
}

// AddPinpointEndpointAccessPolicy Gives permissions to get and update endpoints for a Pinpoint application
func (f *Factory) AddPinpointEndpointAccessPolicy(pinpointApplicationId string) {
	policy := `{ "Action": ["mobiletargeting:GetEndpoint","mobiletargeting:UpdateEndpoint","mobiletargeting:UpdateEndpointsBatch"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:mobiletargeting:${AWS::Region}:${AWS::AccountId}:apps/${pinpointApplicationId}/endpoints/*" }`
	policy = strings.ReplaceAll(policy, "${pinpointApplicationId}", pinpointApplicationId)
	f.policies = append(f.policies, policy)
}

// AddRekognitionFacesPolicy Gives permission to compare and detect faces and labels
func (f *Factory) AddRekognitionFacesPolicy() {
	policy := `{ "Action": ["rekognition:CompareFaces","rekognition:DetectFaces"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddFilterLogEventsPolicy Gives permission to filter Log Events from a specified Log Group
func (f *Factory) AddFilterLogEventsPolicy(logGroupName string) {
	policy := `{ "Action": ["logs:FilterLogEvents"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:logs:${AWS::Region}:${AWS::AccountId}:log-group:${logGroupName}:log-stream:*" }`
	policy = strings.ReplaceAll(policy, "${logGroupName}", logGroupName)
	f.policies = append(f.policies, policy)
}

// AddAthenaQueryPolicy Gives permissions to execute Athena queries
func (f *Factory) AddAthenaQueryPolicy() {
	policy := `{ "Action": ["athena:ListWorkGroups","athena:GetExecutionEngine","athena:GetExecutionEngines","athena:GetNamespace","athena:GetCatalogs","athena:GetNamespaces","athena:GetTables","athena:GetTable"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddS3WritePolicy Gives write permissions to objects in the S3 Bucket
func (f *Factory) AddS3WritePolicy(bucketName string) {
	policy := `{ "Effect": "Allow", "Action": [ "s3:PutObject", "s3:PutObjectAcl", "s3:PutLifecycleConfiguration" ], "Resource": [ "arn:${AWS::Partition}:s3:::${bucketName}","arn:${AWS::Partition}:s3:::${bucketName}/*" ] }`
	policy = strings.ReplaceAll(policy, "${bucketName}", bucketName)
	f.policies = append(f.policies, policy)
}

// AddVPCAccessPolicy Gives access to create, delete, describe and detach ENIs
func (f *Factory) AddVPCAccessPolicy() {
	policy := `{ "Action": ["ec2:CreateNetworkInterface","ec2:DeleteNetworkInterface","ec2:DescribeNetworkInterfaces","ec2:DetachNetworkInterface"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddSNSCrudPolicy Gives permissions to create, publish and subscribe to SNS topics
func (f *Factory) AddSNSCrudPolicy(topicName string) {
	policy := `{ "Action": ["sns:ListSubscriptionsByTopic","sns:CreateTopic","sns:SetTopicAttributes","sns:Subscribe","sns:Publish"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:sns:${AWS::Region}:${AWS::AccountId}:${topicName}*" }`
	policy = strings.ReplaceAll(policy, "${topicName}", topicName)
	f.policies = append(f.policies, policy)
}

// AddKinesisCrudPolicy Gives permission to create, publish and delete Kinesis Stream
func (f *Factory) AddKinesisCrudPolicy(streamName string) {
	policy := `{ "Action": ["kinesis:AddTagsToStream","kinesis:CreateStream","kinesis:DecreaseStreamRetentionPeriod","kinesis:DeleteStream","kinesis:DescribeStream","kinesis:DescribeStreamSummary","kinesis:GetShardIterator","kinesis:IncreaseStreamRetentionPeriod","kinesis:ListTagsForStream","kinesis:MergeShards","kinesis:PutRecord","kinesis:PutRecords","kinesis:SplitShard","kinesis:RemoveTagsFromStream"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:kinesis:${AWS::Region}:${AWS::AccountId}:stream/${streamName}" }`
	policy = strings.ReplaceAll(policy, "${streamName}", streamName)
	f.policies = append(f.policies, policy)
}

// AddCostExplorerReadOnlyPolicy Gives access to the readonly Cost Explorer APIs for billing history
func (f *Factory) AddCostExplorerReadOnlyPolicy() {
	policy := `{ "Action": ["ce:GetCostAndUsage","ce:GetDimensionValues","ce:GetReservationCoverage","ce:GetReservationPurchaseRecommendation","ce:GetReservationUtilization","ce:GetTags"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddStepFunctionsExecutionPolicy Gives permission to start a Step Functions state machine execution
func (f *Factory) AddStepFunctionsExecutionPolicy(stateMachineName string) {
	policy := `{ "Action": ["states:StartExecution"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:states:${AWS::Region}:${AWS::AccountId}:stateMachine:${stateMachineName}" }`
	policy = strings.ReplaceAll(policy, "${stateMachineName}", stateMachineName)
	f.policies = append(f.policies, policy)
}

// AddS3CrudPolicy Gives CRUD permissions to objects in the S3 Bucket
func (f *Factory) AddS3CrudPolicy(bucketName string) {
	policy := `{ "Effect": "Allow", "Action": [ "s3:GetObject", "s3:ListBucket", "s3:GetBucketLocation", "s3:GetObjectVersion", "s3:PutObject", "s3:PutObjectAcl", "s3:GetLifecycleConfiguration", "s3:PutLifecycleConfiguration", "s3:DeleteObject" ], "Resource": ["arn:${AWS::Partition}:s3:::${bucketName}", "arn:${AWS::Partition}:s3:::${bucketName}/*" ] }`
	policy = strings.ReplaceAll(policy, "${bucketName}", bucketName)
	f.policies = append(f.policies, policy)
}

// AddDynamoDBStreamReadPolicy Gives permission to describe and read a DynamoDB Stream and Records
func (f *Factory) AddDynamoDBStreamReadPolicy() {
	policy := `{ "Action": ["dynamodb:DescribeStream","dynamodb:GetRecords","dynamodb:GetShardIterator"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}/stream/${streamName}" }`
	f.policies = append(f.policies, policy)
}

// AddDynamoDBBackupFullAccessPolicy Gives read/write permissions to DynamoDB on-demand backups for a table
func (f *Factory) AddDynamoDBBackupFullAccessPolicy(tableName string) {
	policy := `{ "Action": ["dynamodb:CreateBackup","dynamodb:DescribeContinuousBackups"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}" }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddAWSSecretsManagerRotationPolicy Grants permissions to APIs required to rotate a secret in AWS Secrets Manager
func (f *Factory) AddAWSSecretsManagerRotationPolicy() {
	policy := `{ "Action": ["secretsmanager:DescribeSecret","secretsmanager:GetSecretValue","secretsmanager:PutSecretValue","secretsmanager:UpdateSecretVersionStage"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:*" }`
	f.policies = append(f.policies, policy)
}

// AddS3ReadPolicy Gives read permissions to objects in the S3 Bucket
func (f *Factory) AddS3ReadPolicy(bucketName string) {
	policy := `{ "Effect": "Allow", "Action": [ "s3:GetObject", "s3:ListBucket", "s3:GetBucketLocation", "s3:GetObjectVersion", "s3:GetLifecycleConfiguration" ], "Resource": [ "arn:${AWS::Partition}:s3:::${bucketName}","arn:${AWS::Partition}:s3:::${bucketName}/*"] }`
	policy = strings.ReplaceAll(policy, "${bucketName}", bucketName)
	f.policies = append(f.policies, policy)
}

// AddEKSDescribePolicy Gives permission to describe or list Amazon EKS clusters
func (f *Factory) AddEKSDescribePolicy() {
	policy := `{ "Action": ["eks:DescribeCluster","eks:ListClusters"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddCloudWatchDescribeAlarmHistoryPolicy Gives permissions to describe CloudWatch alarm history
func (f *Factory) AddCloudWatchDescribeAlarmHistoryPolicy() {
	policy := `{ "Action": ["cloudwatch:DescribeAlarmHistory"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddRekognitionDetectOnlyPolicy Gives permission to detect faces, labels and text
func (f *Factory) AddRekognitionDetectOnlyPolicy() {
	policy := `{ "Action": ["rekognition:DetectFaces","rekognition:DetectLabels","rekognition:DetectModerationLabels","rekognition:DetectText"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddPollyFullAccessPolicy Gives full access permissions to Polly lexicon resources
func (f *Factory) AddPollyFullAccessPolicy(lexiconName string) {
	policy := `{ "Effect": "Allow", "Action": [ "polly:GetLexicon", "polly:DeleteLexicon" ], "Resource": [ { "arn:${AWS::Partition}:polly:${AWS::Region}:${AWS::AccountId}:lexicon/${lexiconName}", } ] }, { "Effect": "Allow", "Action": [ "polly:DescribeVoices", "polly:ListLexicons", "polly:PutLexicon", "polly:SynthesizeSpeech" ], "Resource": [ "arn:${AWS::Partition}:polly:${AWS::Region}:${AWS::AccountId}:lexicon/*" ] }`
	policy = strings.ReplaceAll(policy, "${lexiconName}", lexiconName)
	f.policies = append(f.policies, policy)
}

// AddEventBridgePutEventsPolicy Gives permissions to send events to EventBridge
func (f *Factory) AddEventBridgePutEventsPolicy(eventBusName string) {
	policy := `{ "Action": ["events:PutEvents"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:events:${AWS::Region}:${AWS::AccountId}:event-bus/${eventBusName}" }`
	policy = strings.ReplaceAll(policy, "${eventBusName}", eventBusName)
	f.policies = append(f.policies, policy)
}

// AddSSMParameterReadPolicy Gives access to a parameter to load secrets in this account. If not using default key, KMSDecryptPolicy will also be needed.
func (f *Factory) AddSSMParameterReadPolicy() {
	policy := `{ "Action": ["ssm:DescribeParameters"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddCodeCommitCrudPolicy Gives permissions to create/read/update/delete objects within a specific codecommit repository
func (f *Factory) AddCodeCommitCrudPolicy(repositoryName string) {
	policy := `{ "Action": ["codecommit:GitPull","codecommit:GitPush","codecommit:CreateBranch","codecommit:DeleteBranch","codecommit:GetBranch","codecommit:ListBranches","codecommit:MergeBranchesByFastForward","codecommit:MergeBranchesBySquash","codecommit:MergeBranchesByThreeWay","codecommit:UpdateDefaultBranch","codecommit:BatchDescribeMergeConflicts","codecommit:CreateUnreferencedMergeCommit","codecommit:DescribeMergeConflicts","codecommit:GetMergeCommit","codecommit:GetMergeOptions","codecommit:BatchGetPullRequests","codecommit:CreatePullRequest","codecommit:DescribePullRequestEvents","codecommit:GetCommentsForPullRequest","codecommit:GetCommitsFromMergeBase","codecommit:GetMergeConflicts","codecommit:GetPullRequest","codecommit:ListPullRequests","codecommit:MergePullRequestByFastForward","codecommit:MergePullRequestBySquash","codecommit:MergePullRequestByThreeWay","codecommit:PostCommentForPullRequest","codecommit:UpdatePullRequestDescription","codecommit:UpdatePullRequestStatus","codecommit:UpdatePullRequestTitle","codecommit:DeleteFile","codecommit:GetBlob","codecommit:GetFile","codecommit:GetFolder","codecommit:PutFile","codecommit:DeleteCommentContent","codecommit:GetComment","codecommit:GetCommentsForComparedCommit","codecommit:PostCommentForComparedCommit","codecommit:PostCommentReply","codecommit:UpdateComment","codecommit:BatchGetCommits","codecommit:CreateCommit","codecommit:GetCommit","codecommit:GetCommitHistory","codecommit:GetDifferences","codecommit:GetObjectIdentifier","codecommit:GetReferences","codecommit:GetTree","codecommit:GetRepository","codecommit:UpdateRepositoryDescription","codecommit:ListTagsForResource","codecommit:TagResource","codecommit:UntagResource","codecommit:GetRepositoryTriggers","codecommit:PutRepositoryTriggers","codecommit:TestRepositoryTriggers","codecommit:GetBranch","codecommit:GetCommit","codecommit:UploadArchive","codecommit:GetUploadArchiveStatus","codecommit:CancelUploadArchive"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:codecommit:${AWS::Region}:${AWS::AccountId}:${repositoryName}" }`
	policy = strings.ReplaceAll(policy, "${repositoryName}", repositoryName)
	f.policies = append(f.policies, policy)
}

// AddCloudWatchPutMetricPolicy Gives permissions to put metrics to CloudWatch
func (f *Factory) AddCloudWatchPutMetricPolicy() {
	policy := `{ "Action": ["cloudwatch:PutMetricData"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddEC2DescribePolicy Gives permission to describe EC2 instances
func (f *Factory) AddEC2DescribePolicy() {
	policy := `{ "Action": ["ec2:DescribeRegions","ec2:DescribeInstances"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddRekognitionWriteOnlyAccessPolicy Gives permission to create collection and index faces
func (f *Factory) AddRekognitionWriteOnlyAccessPolicy(collectionId string) {
	policy := `{ "Action": ["rekognition:CreateCollection","rekognition:IndexFaces"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:rekognition:${AWS::Region}:${AWS::AccountId}:collection/${collectionId}" }`
	policy = strings.ReplaceAll(policy, "${collectionId}", collectionId)
	f.policies = append(f.policies, policy)
}

// AddServerlessRepoReadWriteAccessPolicy Gives access permissions to create and list applications in the AWS Serverless Application Repository service
func (f *Factory) AddServerlessRepoReadWriteAccessPolicy() {
	policy := `{ "Effect": "Allow", "Action": [ "serverlessrepo:CreateApplication", "serverlessrepo:CreateApplicationVersion", "serverlessrepo:UpdateApplication", "serverlessrepo:GetApplication", "serverlessrepo:ListApplications", "serverlessrepo:ListApplicationVersions", "serverlessrepo:ListApplicationDependencies" ], "Resource": [ "arn:${AWS::Partition}:serverlessrepo:${AWS::Region}:${AWS::AccountId}:applications/*" ] }`
	f.policies = append(f.policies, policy)
}

// AddRekognitionFacesManagementPolicy Gives permission to add, delete and search faces in a collection
func (f *Factory) AddRekognitionFacesManagementPolicy(collectionId string) {
	policy := `{ "Action": ["rekognition:IndexFaces","rekognition:DeleteFaces","rekognition:SearchFaces","rekognition:SearchFacesByImage","rekognition:ListFaces"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:rekognition:${AWS::Region}:${AWS::AccountId}:collection/${collectionId}" }`
	policy = strings.ReplaceAll(policy, "${collectionId}", collectionId)
	f.policies = append(f.policies, policy)
}

// AddFirehoseWritePolicy Gives permission to write to a Kinesis Firehose Delivery Stream
func (f *Factory) AddFirehoseWritePolicy(deliveryStreamName string) {
	policy := `{ "Action": ["firehose:PutRecord","firehose:PutRecordBatch"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:firehose:${AWS::Region}:${AWS::AccountId}:deliverystream/${deliveryStreamName}" }`
	policy = strings.ReplaceAll(policy, "${deliveryStreamName}", deliveryStreamName)
	f.policies = append(f.policies, policy)
}

// AddSESSendBouncePolicy Gives SendBounce permission to a SES identity
func (f *Factory) AddSESSendBouncePolicy(identityName string) {
	policy := `{ "Action": ["ses:SendBounce"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:ses:${AWS::Region}:${AWS::AccountId}:identity/${identityName}" }`
	policy = strings.ReplaceAll(policy, "${identityName}", identityName)
	f.policies = append(f.policies, policy)
}

// AddCloudFormationDescribeStacksPolicy Gives permission to describe CloudFormation stacks
func (f *Factory) AddCloudFormationDescribeStacksPolicy() {
	policy := `{ "Action": ["cloudformation:DescribeStacks"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:cloudformation:${AWS::Region}:${AWS::AccountId}:stack/*" }`
	f.policies = append(f.policies, policy)
}

// AddKMSDecryptPolicy Gives permission to decrypt with KMS Key
func (f *Factory) AddKMSDecryptPolicy(keyId string) {
	policy := `{ "Action": ["kms:Decrypt"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:kms:${AWS::Region}:${AWS::AccountId}:key/${keyId}" }`
	policy = strings.ReplaceAll(policy, "${keyId}", keyId)
	f.policies = append(f.policies, policy)
}

// AddKMSEncryptPolicy Gives permission to encrypt with KMS Key
func (f *Factory) AddKMSEncryptPolicy(keyId string) {
	policy := `{ "Action": ["kms:Encrypt"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:kms:${AWS::Region}:${AWS::AccountId}:key/${keyId}" }`
	policy = strings.ReplaceAll(policy, "${keyId}", keyId)
	f.policies = append(f.policies, policy)
}

// AddDynamoDBReadPolicy Gives read only access to a DynamoDB Table
func (f *Factory) AddDynamoDBReadPolicy(tableName string) {
	policy := `{ "Effect": "Allow", "Action": [ "dynamodb:GetItem", "dynamodb:Scan", "dynamodb:Query", "dynamodb:BatchGetItem", "dynamodb:DescribeTable" ], "Resource": [ "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}", "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}/index/*"] }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddDynamoDBRestoreFromBackupPolicy Gives permissions to restore a table from backup
func (f *Factory) AddDynamoDBRestoreFromBackupPolicy(tableName string) {
	policy := `{ "Action": ["dynamodb:RestoreTableFromBackup"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}/backup/*" }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddSESEmailTemplateCrudPolicy Gives permission to create, get, list, update and delete SES Email Templates
func (f *Factory) AddSESEmailTemplateCrudPolicy() {
	policy := `{ "Action": ["ses:CreateTemplate","ses:GetTemplate","ses:ListTemplates","ses:UpdateTemplate","ses:DeleteTemplate","ses:TestRenderTemplate"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddTextractDetectAnalyzePolicy Gives access to detect and analyze documents with Textract
func (f *Factory) AddTextractDetectAnalyzePolicy() {
	policy := `{ "Action": ["textract:DetectDocumentText","textract:StartDocumentTextDetection","textract:StartDocumentAnalysis","textract:AnalyzeDocument"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddCloudWatchDashboardPolicy Gives permissions to put metrics to operate on CloudWatch Dashboards
func (f *Factory) AddCloudWatchDashboardPolicy() {
	policy := `{ "Action": ["cloudwatch:GetDashboard","cloudwatch:ListDashboards","cloudwatch:PutDashboard","cloudwatch:ListMetrics"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddComprehendBasicAccessPolicy Gives access to Amazon Comprehend APIs for detecting entities, key phrases, languages and sentiments
func (f *Factory) AddComprehendBasicAccessPolicy() {
	policy := `{ "Action": ["comprehend:BatchDetectKeyPhrases","comprehend:DetectDominantLanguage","comprehend:DetectEntities","comprehend:BatchDetectEntities","comprehend:DetectKeyPhrases","comprehend:DetectSentiment","comprehend:BatchDetectDominantLanguage","comprehend:BatchDetectSentiment"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddFirehoseCrudPolicy Gives permission to create, write to, update, and delete a Kinesis Firehose Delivery Stream
func (f *Factory) AddFirehoseCrudPolicy(deliveryStreamName string) {
	policy := `{ "Action": ["firehose:CreateDeliveryStream","firehose:DeleteDeliveryStream","firehose:DescribeDeliveryStream","firehose:PutRecord","firehose:PutRecordBatch","firehose:UpdateDestination"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:firehose:${AWS::Region}:${AWS::AccountId}:deliverystream/${deliveryStreamName}" }`
	policy = strings.ReplaceAll(policy, "${deliveryStreamName}", deliveryStreamName)
	f.policies = append(f.policies, policy)
}

// AddTextractPolicy Gives full access to Textract
func (f *Factory) AddTextractPolicy() {
	policy := `{ "Action": ["textract:*"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddDynamoDBReconfigurePolicy Gives access reconfigure to a DynamoDB Table
func (f *Factory) AddDynamoDBReconfigurePolicy(tableName string) {
	policy := `{ "Action": ["dynamodb:UpdateTable"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:dynamodb:${AWS::Region}:${AWS::AccountId}:table/${tableName}" }`
	policy = strings.ReplaceAll(policy, "${tableName}", tableName)
	f.policies = append(f.policies, policy)
}

// AddCodePipelineLambdaExecutionPolicy Gives permission for a Lambda function invoked by AWS CodePipeline to report back status of the job
func (f *Factory) AddCodePipelineLambdaExecutionPolicy() {
	policy := `{ "Action": ["codepipeline:PutJobSuccessResult","codepipeline:PutJobFailureResult"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddOrganizationsListAccountsPolicy Gives readonly permission to list child account names and ids
func (f *Factory) AddOrganizationsListAccountsPolicy() {
	policy := `{ "Action": ["organizations:ListAccounts"], "Effect": "Allow", "Resource": "*" }`
	f.policies = append(f.policies, policy)
}

// AddCodeCommitReadPolicy Gives permissions to read objects within a specific codecommit repository
func (f *Factory) AddCodeCommitReadPolicy(repositoryName string) {
	policy := `{ "Action": ["codecommit:GitPull","codecommit:GetBranch","codecommit:ListBranches","codecommit:BatchDescribeMergeConflicts","codecommit:DescribeMergeConflicts","codecommit:GetMergeCommit","codecommit:GetMergeOptions","codecommit:BatchGetPullRequests","codecommit:DescribePullRequestEvents","codecommit:GetCommentsForPullRequest","codecommit:GetCommitsFromMergeBase","codecommit:GetMergeConflicts","codecommit:GetPullRequest","codecommit:ListPullRequests","codecommit:GetBlob","codecommit:GetFile","codecommit:GetFolder","codecommit:GetComment","codecommit:GetCommentsForComparedCommit","codecommit:BatchGetCommits","codecommit:GetCommit","codecommit:GetCommitHistory","codecommit:GetDifferences","codecommit:GetObjectIdentifier","codecommit:GetReferences","codecommit:GetTree","codecommit:GetRepository","codecommit:ListTagsForResource","codecommit:GetRepositoryTriggers","codecommit:TestRepositoryTriggers","codecommit:GetBranch","codecommit:GetCommit","codecommit:GetUploadArchiveStatus"], "Effect": "Allow", "Resource": "arn:${AWS::Partition}:codecommit:${AWS::Region}:${AWS::AccountId}:${repositoryName}" }`
	policy = strings.ReplaceAll(policy, "${repositoryName}", repositoryName)
	f.policies = append(f.policies, policy)
}
