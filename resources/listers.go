package resources

import (
	"github.com/aws/aws-sdk-go/aws/session"
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

func GetListers(sess *session.Session) []ResourceLister {
	var (
		autoscaling      = AutoScalingNuke{autoscaling.New(sess)}
		cloudformation   = CloudFormationNuke{cloudformation.New(sess)}
		cloudtrail       = CloudTrailNuke{cloudtrail.New(sess)}
		cloudwatchevents = CloudWatchEventsNuke{cloudwatchevents.New(sess)}
		ec2              = EC2Nuke{ec2.New(sess)}
		ecr              = ECRNuke{ecr.New(sess)}
		efs              = EFSNuke{efs.New(sess)}
		elasticache      = ElasticacheNuke{elasticache.New(sess)}
		elb              = ElbNuke{elb.New(sess)}
		elbv2            = Elbv2Nuke{elbv2.New(sess)}
		iam              = IamNuke{iam.New(sess)}
		kms              = KMSNuke{kms.New(sess)}
		lambda           = LambdaNuke{lambda.New(sess)}
		rds              = RDSNuke{rds.New(sess)}
		route53          = Route53Nuke{route53.New(sess)}
		sns              = SNSNuke{sns.New(sess)}
	)

	return []ResourceLister{
		autoscaling.ListGroups,
		autoscaling.ListLaunchConfigurations,
		cloudformation.ListStacks,
		cloudtrail.ListTrails,
		cloudwatchevents.ListRules,
		cloudwatchevents.ListTargets,
		ec2.ListAddresses,
		ec2.ListCustomerGateways,
		ec2.ListDhcpOptions,
		ec2.ListInstances,
		ec2.ListInternetGatewayAttachements,
		ec2.ListInternetGateways,
		ec2.ListKeyPairs,
		ec2.ListNatGateways,
		ec2.ListNetworkACLs,
		ec2.ListRouteTables,
		ec2.ListSecurityGroups,
		ec2.ListSpotFleetRequests,
		ec2.ListSubnets,
		ec2.ListVolumes,
		ec2.ListVpcEndpoints,
		ec2.ListVpcs,
		ec2.ListVpnConnections,
		ec2.ListVpnGatewayAttachements,
		ec2.ListVpnGateways,
		ecr.ListRepos,
		efs.ListFileSystems,
		efs.ListMountTargets,
		elasticache.ListCacheClusters,
		elasticache.ListSubnetGroups,
		elb.ListELBs,
		elbv2.ListELBs,
		elbv2.ListTargetGroups,
		iam.ListGroupPolicyAttachements,
		iam.ListGroups,
		iam.ListInstanceProfileRoles,
		iam.ListInstanceProfiles,
		iam.ListPolicies,
		iam.ListRolePolicyAttachements,
		iam.ListRoles,
		iam.ListServerCertificates,
		iam.ListUserAccessKeys,
		iam.ListUserGroupAttachements,
		iam.ListUserGroupAttachements,
		iam.ListUserPolicyAttachements,
		iam.ListUsers,
		kms.ListAliases,
		kms.ListKeys,
		lambda.ListFunctions,
		rds.ListInstances,
		rds.ListParameterGroups,
		rds.ListSnapshots,
		rds.ListSubnetGroups,
		route53.ListHostedZones,
		route53.ListResourceRecords,
		sns.ListSubscriptions,
		sns.ListTopics,
	}
}
