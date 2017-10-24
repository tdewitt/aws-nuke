package resources

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/sns"
)

type AutoScalingNuke struct {
	Service *autoscaling.AutoScaling
}

type CloudFormationNuke struct {
	Service *cloudformation.CloudFormation
}

type CloudTrailNuke struct {
	Service *cloudtrail.CloudTrail
}

type CloudWatchEventsNuke struct {
	Service *cloudwatchevents.CloudWatchEvents
}

type EC2Nuke struct {
	Service *ec2.EC2
}

type ECRNuke struct {
	Service *ecr.ECR
}

type EFSNuke struct {
	Service *efs.EFS
}

type ElasticacheNuke struct {
	Service *elasticache.ElastiCache
}

type ElbNuke struct {
	Service *elb.ELB
}

type Elbv2Nuke struct {
	Service *elbv2.ELBV2
}

type IamNuke struct {
	Service *iam.IAM
}

type LambdaNuke struct {
	Service *lambda.Lambda
}

type RDSNuke struct {
	Service *rds.RDS
}

type Route53Nuke struct {
	Service *route53.Route53
}

type SNSNuke struct {
	Service *sns.SNS
}

type KMSNuke struct {
	Service *kms.KMS
}
