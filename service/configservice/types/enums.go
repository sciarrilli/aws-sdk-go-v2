// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AggregatedSourceStatusType string

// Enum values for AggregatedSourceStatusType
const (
	AggregatedSourceStatusTypeFailed    AggregatedSourceStatusType = "FAILED"
	AggregatedSourceStatusTypeSucceeded AggregatedSourceStatusType = "SUCCEEDED"
	AggregatedSourceStatusTypeOutdated  AggregatedSourceStatusType = "OUTDATED"
)

// Values returns all known values for AggregatedSourceStatusType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AggregatedSourceStatusType) Values() []AggregatedSourceStatusType {
	return []AggregatedSourceStatusType{
		"FAILED",
		"SUCCEEDED",
		"OUTDATED",
	}
}

type AggregatedSourceType string

// Enum values for AggregatedSourceType
const (
	AggregatedSourceTypeAccount      AggregatedSourceType = "ACCOUNT"
	AggregatedSourceTypeOrganization AggregatedSourceType = "ORGANIZATION"
)

// Values returns all known values for AggregatedSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AggregatedSourceType) Values() []AggregatedSourceType {
	return []AggregatedSourceType{
		"ACCOUNT",
		"ORGANIZATION",
	}
}

type ChronologicalOrder string

// Enum values for ChronologicalOrder
const (
	ChronologicalOrderReverse ChronologicalOrder = "Reverse"
	ChronologicalOrderForward ChronologicalOrder = "Forward"
)

// Values returns all known values for ChronologicalOrder. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChronologicalOrder) Values() []ChronologicalOrder {
	return []ChronologicalOrder{
		"Reverse",
		"Forward",
	}
}

type ComplianceType string

// Enum values for ComplianceType
const (
	ComplianceTypeCompliant        ComplianceType = "COMPLIANT"
	ComplianceTypeNonCompliant     ComplianceType = "NON_COMPLIANT"
	ComplianceTypeNotApplicable    ComplianceType = "NOT_APPLICABLE"
	ComplianceTypeInsufficientData ComplianceType = "INSUFFICIENT_DATA"
)

// Values returns all known values for ComplianceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComplianceType) Values() []ComplianceType {
	return []ComplianceType{
		"COMPLIANT",
		"NON_COMPLIANT",
		"NOT_APPLICABLE",
		"INSUFFICIENT_DATA",
	}
}

type ConfigRuleComplianceSummaryGroupKey string

// Enum values for ConfigRuleComplianceSummaryGroupKey
const (
	ConfigRuleComplianceSummaryGroupKeyAccountId ConfigRuleComplianceSummaryGroupKey = "ACCOUNT_ID"
	ConfigRuleComplianceSummaryGroupKeyAwsRegion ConfigRuleComplianceSummaryGroupKey = "AWS_REGION"
)

// Values returns all known values for ConfigRuleComplianceSummaryGroupKey. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigRuleComplianceSummaryGroupKey) Values() []ConfigRuleComplianceSummaryGroupKey {
	return []ConfigRuleComplianceSummaryGroupKey{
		"ACCOUNT_ID",
		"AWS_REGION",
	}
}

type ConfigRuleState string

// Enum values for ConfigRuleState
const (
	ConfigRuleStateActive          ConfigRuleState = "ACTIVE"
	ConfigRuleStateDeleting        ConfigRuleState = "DELETING"
	ConfigRuleStateDeletingResults ConfigRuleState = "DELETING_RESULTS"
	ConfigRuleStateEvaluating      ConfigRuleState = "EVALUATING"
)

// Values returns all known values for ConfigRuleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigRuleState) Values() []ConfigRuleState {
	return []ConfigRuleState{
		"ACTIVE",
		"DELETING",
		"DELETING_RESULTS",
		"EVALUATING",
	}
}

type ConfigurationItemStatus string

// Enum values for ConfigurationItemStatus
const (
	ConfigurationItemStatusOk                         ConfigurationItemStatus = "OK"
	ConfigurationItemStatusResourceDiscovered         ConfigurationItemStatus = "ResourceDiscovered"
	ConfigurationItemStatusResourceNotRecorded        ConfigurationItemStatus = "ResourceNotRecorded"
	ConfigurationItemStatusResourceDeleted            ConfigurationItemStatus = "ResourceDeleted"
	ConfigurationItemStatusResourceDeletedNotRecorded ConfigurationItemStatus = "ResourceDeletedNotRecorded"
)

// Values returns all known values for ConfigurationItemStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationItemStatus) Values() []ConfigurationItemStatus {
	return []ConfigurationItemStatus{
		"OK",
		"ResourceDiscovered",
		"ResourceNotRecorded",
		"ResourceDeleted",
		"ResourceDeletedNotRecorded",
	}
}

type ConformancePackComplianceType string

// Enum values for ConformancePackComplianceType
const (
	ConformancePackComplianceTypeCompliant        ConformancePackComplianceType = "COMPLIANT"
	ConformancePackComplianceTypeNonCompliant     ConformancePackComplianceType = "NON_COMPLIANT"
	ConformancePackComplianceTypeInsufficientData ConformancePackComplianceType = "INSUFFICIENT_DATA"
)

// Values returns all known values for ConformancePackComplianceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConformancePackComplianceType) Values() []ConformancePackComplianceType {
	return []ConformancePackComplianceType{
		"COMPLIANT",
		"NON_COMPLIANT",
		"INSUFFICIENT_DATA",
	}
}

type ConformancePackState string

// Enum values for ConformancePackState
const (
	ConformancePackStateCreateInProgress ConformancePackState = "CREATE_IN_PROGRESS"
	ConformancePackStateCreateComplete   ConformancePackState = "CREATE_COMPLETE"
	ConformancePackStateCreateFailed     ConformancePackState = "CREATE_FAILED"
	ConformancePackStateDeleteInProgress ConformancePackState = "DELETE_IN_PROGRESS"
	ConformancePackStateDeleteFailed     ConformancePackState = "DELETE_FAILED"
)

// Values returns all known values for ConformancePackState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConformancePackState) Values() []ConformancePackState {
	return []ConformancePackState{
		"CREATE_IN_PROGRESS",
		"CREATE_COMPLETE",
		"CREATE_FAILED",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
	}
}

type DeliveryStatus string

// Enum values for DeliveryStatus
const (
	DeliveryStatusSuccess       DeliveryStatus = "Success"
	DeliveryStatusFailure       DeliveryStatus = "Failure"
	DeliveryStatusNotApplicable DeliveryStatus = "Not_Applicable"
)

// Values returns all known values for DeliveryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryStatus) Values() []DeliveryStatus {
	return []DeliveryStatus{
		"Success",
		"Failure",
		"Not_Applicable",
	}
}

type EventSource string

// Enum values for EventSource
const (
	EventSourceAwsConfig EventSource = "aws.config"
)

// Values returns all known values for EventSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EventSource) Values() []EventSource {
	return []EventSource{
		"aws.config",
	}
}

type MaximumExecutionFrequency string

// Enum values for MaximumExecutionFrequency
const (
	MaximumExecutionFrequencyOneHour         MaximumExecutionFrequency = "One_Hour"
	MaximumExecutionFrequencyThreeHours      MaximumExecutionFrequency = "Three_Hours"
	MaximumExecutionFrequencySixHours        MaximumExecutionFrequency = "Six_Hours"
	MaximumExecutionFrequencyTwelveHours     MaximumExecutionFrequency = "Twelve_Hours"
	MaximumExecutionFrequencyTwentyFourHours MaximumExecutionFrequency = "TwentyFour_Hours"
)

// Values returns all known values for MaximumExecutionFrequency. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (MaximumExecutionFrequency) Values() []MaximumExecutionFrequency {
	return []MaximumExecutionFrequency{
		"One_Hour",
		"Three_Hours",
		"Six_Hours",
		"Twelve_Hours",
		"TwentyFour_Hours",
	}
}

type MemberAccountRuleStatus string

// Enum values for MemberAccountRuleStatus
const (
	MemberAccountRuleStatusCreateSuccessful MemberAccountRuleStatus = "CREATE_SUCCESSFUL"
	MemberAccountRuleStatusCreateInProgress MemberAccountRuleStatus = "CREATE_IN_PROGRESS"
	MemberAccountRuleStatusCreateFailed     MemberAccountRuleStatus = "CREATE_FAILED"
	MemberAccountRuleStatusDeleteSuccessful MemberAccountRuleStatus = "DELETE_SUCCESSFUL"
	MemberAccountRuleStatusDeleteFailed     MemberAccountRuleStatus = "DELETE_FAILED"
	MemberAccountRuleStatusDeleteInProgress MemberAccountRuleStatus = "DELETE_IN_PROGRESS"
	MemberAccountRuleStatusUpdateSuccessful MemberAccountRuleStatus = "UPDATE_SUCCESSFUL"
	MemberAccountRuleStatusUpdateInProgress MemberAccountRuleStatus = "UPDATE_IN_PROGRESS"
	MemberAccountRuleStatusUpdateFailed     MemberAccountRuleStatus = "UPDATE_FAILED"
)

// Values returns all known values for MemberAccountRuleStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberAccountRuleStatus) Values() []MemberAccountRuleStatus {
	return []MemberAccountRuleStatus{
		"CREATE_SUCCESSFUL",
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_SUCCESSFUL",
		"DELETE_FAILED",
		"DELETE_IN_PROGRESS",
		"UPDATE_SUCCESSFUL",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeConfigurationItemChangeNotification          MessageType = "ConfigurationItemChangeNotification"
	MessageTypeConfigurationSnapshotDeliveryCompleted       MessageType = "ConfigurationSnapshotDeliveryCompleted"
	MessageTypeScheduledNotification                        MessageType = "ScheduledNotification"
	MessageTypeOversizedConfigurationItemChangeNotification MessageType = "OversizedConfigurationItemChangeNotification"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"ConfigurationItemChangeNotification",
		"ConfigurationSnapshotDeliveryCompleted",
		"ScheduledNotification",
		"OversizedConfigurationItemChangeNotification",
	}
}

type OrganizationConfigRuleTriggerType string

// Enum values for OrganizationConfigRuleTriggerType
const (
	OrganizationConfigRuleTriggerTypeConfigurationItemChangeNotification         OrganizationConfigRuleTriggerType = "ConfigurationItemChangeNotification"
	OrganizationConfigRuleTriggerTypeOversizedConfigurationItemChangeNotifcation OrganizationConfigRuleTriggerType = "OversizedConfigurationItemChangeNotification"
	OrganizationConfigRuleTriggerTypeScheduledNotification                       OrganizationConfigRuleTriggerType = "ScheduledNotification"
)

// Values returns all known values for OrganizationConfigRuleTriggerType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OrganizationConfigRuleTriggerType) Values() []OrganizationConfigRuleTriggerType {
	return []OrganizationConfigRuleTriggerType{
		"ConfigurationItemChangeNotification",
		"OversizedConfigurationItemChangeNotification",
		"ScheduledNotification",
	}
}

type OrganizationResourceDetailedStatus string

// Enum values for OrganizationResourceDetailedStatus
const (
	OrganizationResourceDetailedStatusCreateSuccessful OrganizationResourceDetailedStatus = "CREATE_SUCCESSFUL"
	OrganizationResourceDetailedStatusCreateInProgress OrganizationResourceDetailedStatus = "CREATE_IN_PROGRESS"
	OrganizationResourceDetailedStatusCreateFailed     OrganizationResourceDetailedStatus = "CREATE_FAILED"
	OrganizationResourceDetailedStatusDeleteSuccessful OrganizationResourceDetailedStatus = "DELETE_SUCCESSFUL"
	OrganizationResourceDetailedStatusDeleteFailed     OrganizationResourceDetailedStatus = "DELETE_FAILED"
	OrganizationResourceDetailedStatusDeleteInProgress OrganizationResourceDetailedStatus = "DELETE_IN_PROGRESS"
	OrganizationResourceDetailedStatusUpdateSuccessful OrganizationResourceDetailedStatus = "UPDATE_SUCCESSFUL"
	OrganizationResourceDetailedStatusUpdateInProgress OrganizationResourceDetailedStatus = "UPDATE_IN_PROGRESS"
	OrganizationResourceDetailedStatusUpdateFailed     OrganizationResourceDetailedStatus = "UPDATE_FAILED"
)

// Values returns all known values for OrganizationResourceDetailedStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OrganizationResourceDetailedStatus) Values() []OrganizationResourceDetailedStatus {
	return []OrganizationResourceDetailedStatus{
		"CREATE_SUCCESSFUL",
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_SUCCESSFUL",
		"DELETE_FAILED",
		"DELETE_IN_PROGRESS",
		"UPDATE_SUCCESSFUL",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type OrganizationResourceStatus string

// Enum values for OrganizationResourceStatus
const (
	OrganizationResourceStatusCreateSuccessful OrganizationResourceStatus = "CREATE_SUCCESSFUL"
	OrganizationResourceStatusCreateInProgress OrganizationResourceStatus = "CREATE_IN_PROGRESS"
	OrganizationResourceStatusCreateFailed     OrganizationResourceStatus = "CREATE_FAILED"
	OrganizationResourceStatusDeleteSuccessful OrganizationResourceStatus = "DELETE_SUCCESSFUL"
	OrganizationResourceStatusDeleteFailed     OrganizationResourceStatus = "DELETE_FAILED"
	OrganizationResourceStatusDeleteInProgress OrganizationResourceStatus = "DELETE_IN_PROGRESS"
	OrganizationResourceStatusUpdateSuccessful OrganizationResourceStatus = "UPDATE_SUCCESSFUL"
	OrganizationResourceStatusUpdateInProgress OrganizationResourceStatus = "UPDATE_IN_PROGRESS"
	OrganizationResourceStatusUpdateFailed     OrganizationResourceStatus = "UPDATE_FAILED"
)

// Values returns all known values for OrganizationResourceStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationResourceStatus) Values() []OrganizationResourceStatus {
	return []OrganizationResourceStatus{
		"CREATE_SUCCESSFUL",
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_SUCCESSFUL",
		"DELETE_FAILED",
		"DELETE_IN_PROGRESS",
		"UPDATE_SUCCESSFUL",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type OrganizationRuleStatus string

// Enum values for OrganizationRuleStatus
const (
	OrganizationRuleStatusCreateSuccessful OrganizationRuleStatus = "CREATE_SUCCESSFUL"
	OrganizationRuleStatusCreateInProgress OrganizationRuleStatus = "CREATE_IN_PROGRESS"
	OrganizationRuleStatusCreateFailed     OrganizationRuleStatus = "CREATE_FAILED"
	OrganizationRuleStatusDeleteSuccessful OrganizationRuleStatus = "DELETE_SUCCESSFUL"
	OrganizationRuleStatusDeleteFailed     OrganizationRuleStatus = "DELETE_FAILED"
	OrganizationRuleStatusDeleteInProgress OrganizationRuleStatus = "DELETE_IN_PROGRESS"
	OrganizationRuleStatusUpdateSuccessful OrganizationRuleStatus = "UPDATE_SUCCESSFUL"
	OrganizationRuleStatusUpdateInProgress OrganizationRuleStatus = "UPDATE_IN_PROGRESS"
	OrganizationRuleStatusUpdateFailed     OrganizationRuleStatus = "UPDATE_FAILED"
)

// Values returns all known values for OrganizationRuleStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationRuleStatus) Values() []OrganizationRuleStatus {
	return []OrganizationRuleStatus{
		"CREATE_SUCCESSFUL",
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_SUCCESSFUL",
		"DELETE_FAILED",
		"DELETE_IN_PROGRESS",
		"UPDATE_SUCCESSFUL",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type Owner string

// Enum values for Owner
const (
	OwnerCustomLambda Owner = "CUSTOM_LAMBDA"
	OwnerAws          Owner = "AWS"
)

// Values returns all known values for Owner. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Owner) Values() []Owner {
	return []Owner{
		"CUSTOM_LAMBDA",
		"AWS",
	}
}

type RecorderStatus string

// Enum values for RecorderStatus
const (
	RecorderStatusPending RecorderStatus = "Pending"
	RecorderStatusSuccess RecorderStatus = "Success"
	RecorderStatusFailure RecorderStatus = "Failure"
)

// Values returns all known values for RecorderStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecorderStatus) Values() []RecorderStatus {
	return []RecorderStatus{
		"Pending",
		"Success",
		"Failure",
	}
}

type RemediationExecutionState string

// Enum values for RemediationExecutionState
const (
	RemediationExecutionStateQueued     RemediationExecutionState = "QUEUED"
	RemediationExecutionStateInProgress RemediationExecutionState = "IN_PROGRESS"
	RemediationExecutionStateSucceeded  RemediationExecutionState = "SUCCEEDED"
	RemediationExecutionStateFailed     RemediationExecutionState = "FAILED"
)

// Values returns all known values for RemediationExecutionState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RemediationExecutionState) Values() []RemediationExecutionState {
	return []RemediationExecutionState{
		"QUEUED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type RemediationExecutionStepState string

// Enum values for RemediationExecutionStepState
const (
	RemediationExecutionStepStateSucceeded RemediationExecutionStepState = "SUCCEEDED"
	RemediationExecutionStepStatePending   RemediationExecutionStepState = "PENDING"
	RemediationExecutionStepStateFailed    RemediationExecutionStepState = "FAILED"
)

// Values returns all known values for RemediationExecutionStepState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RemediationExecutionStepState) Values() []RemediationExecutionStepState {
	return []RemediationExecutionStepState{
		"SUCCEEDED",
		"PENDING",
		"FAILED",
	}
}

type RemediationTargetType string

// Enum values for RemediationTargetType
const (
	RemediationTargetTypeSsmDocument RemediationTargetType = "SSM_DOCUMENT"
)

// Values returns all known values for RemediationTargetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RemediationTargetType) Values() []RemediationTargetType {
	return []RemediationTargetType{
		"SSM_DOCUMENT",
	}
}

type ResourceCountGroupKey string

// Enum values for ResourceCountGroupKey
const (
	ResourceCountGroupKeyResourceType ResourceCountGroupKey = "RESOURCE_TYPE"
	ResourceCountGroupKeyAccountId    ResourceCountGroupKey = "ACCOUNT_ID"
	ResourceCountGroupKeyAwsRegion    ResourceCountGroupKey = "AWS_REGION"
)

// Values returns all known values for ResourceCountGroupKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceCountGroupKey) Values() []ResourceCountGroupKey {
	return []ResourceCountGroupKey{
		"RESOURCE_TYPE",
		"ACCOUNT_ID",
		"AWS_REGION",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeCustomerGateway                  ResourceType = "AWS::EC2::CustomerGateway"
	ResourceTypeEip                              ResourceType = "AWS::EC2::EIP"
	ResourceTypeHost                             ResourceType = "AWS::EC2::Host"
	ResourceTypeInstance                         ResourceType = "AWS::EC2::Instance"
	ResourceTypeInternetGateway                  ResourceType = "AWS::EC2::InternetGateway"
	ResourceTypeNetworkAcl                       ResourceType = "AWS::EC2::NetworkAcl"
	ResourceTypeNetworkInterface                 ResourceType = "AWS::EC2::NetworkInterface"
	ResourceTypeRouteTable                       ResourceType = "AWS::EC2::RouteTable"
	ResourceTypeSecurityGroup                    ResourceType = "AWS::EC2::SecurityGroup"
	ResourceTypeSubnet                           ResourceType = "AWS::EC2::Subnet"
	ResourceTypeTrail                            ResourceType = "AWS::CloudTrail::Trail"
	ResourceTypeVolume                           ResourceType = "AWS::EC2::Volume"
	ResourceTypeVpc                              ResourceType = "AWS::EC2::VPC"
	ResourceTypeVPNConnection                    ResourceType = "AWS::EC2::VPNConnection"
	ResourceTypeVPNGateway                       ResourceType = "AWS::EC2::VPNGateway"
	ResourceTypeRegisteredHAInstance             ResourceType = "AWS::EC2::RegisteredHAInstance"
	ResourceTypeNatGateway                       ResourceType = "AWS::EC2::NatGateway"
	ResourceTypeEgressOnlyInternetGateway        ResourceType = "AWS::EC2::EgressOnlyInternetGateway"
	ResourceTypeVPCEndpoint                      ResourceType = "AWS::EC2::VPCEndpoint"
	ResourceTypeVPCEndpointService               ResourceType = "AWS::EC2::VPCEndpointService"
	ResourceTypeFlowLog                          ResourceType = "AWS::EC2::FlowLog"
	ResourceTypeVPCPeeringConnection             ResourceType = "AWS::EC2::VPCPeeringConnection"
	ResourceTypeDomain                           ResourceType = "AWS::Elasticsearch::Domain"
	ResourceTypeGroup                            ResourceType = "AWS::IAM::Group"
	ResourceTypePolicy                           ResourceType = "AWS::IAM::Policy"
	ResourceTypeRole                             ResourceType = "AWS::IAM::Role"
	ResourceTypeUser                             ResourceType = "AWS::IAM::User"
	ResourceTypeLoadBalancerV2                   ResourceType = "AWS::ElasticLoadBalancingV2::LoadBalancer"
	ResourceTypeCertificate                      ResourceType = "AWS::ACM::Certificate"
	ResourceTypeDBInstance                       ResourceType = "AWS::RDS::DBInstance"
	ResourceTypeDBSubnetGroup                    ResourceType = "AWS::RDS::DBSubnetGroup"
	ResourceTypeDBSecurityGroup                  ResourceType = "AWS::RDS::DBSecurityGroup"
	ResourceTypeDBSnapshot                       ResourceType = "AWS::RDS::DBSnapshot"
	ResourceTypeDBCluster                        ResourceType = "AWS::RDS::DBCluster"
	ResourceTypeDBClusterSnapshot                ResourceType = "AWS::RDS::DBClusterSnapshot"
	ResourceTypeEventSubscription                ResourceType = "AWS::RDS::EventSubscription"
	ResourceTypeBucket                           ResourceType = "AWS::S3::Bucket"
	ResourceTypeAccountPublicAccessBlock         ResourceType = "AWS::S3::AccountPublicAccessBlock"
	ResourceTypeCluster                          ResourceType = "AWS::Redshift::Cluster"
	ResourceTypeClusterSnapshot                  ResourceType = "AWS::Redshift::ClusterSnapshot"
	ResourceTypeClusterParameterGroup            ResourceType = "AWS::Redshift::ClusterParameterGroup"
	ResourceTypeClusterSecurityGroup             ResourceType = "AWS::Redshift::ClusterSecurityGroup"
	ResourceTypeClusterSubnetGroup               ResourceType = "AWS::Redshift::ClusterSubnetGroup"
	ResourceTypeRedshiftEventSubscription        ResourceType = "AWS::Redshift::EventSubscription"
	ResourceTypeManagedInstanceInventory         ResourceType = "AWS::SSM::ManagedInstanceInventory"
	ResourceTypeAlarm                            ResourceType = "AWS::CloudWatch::Alarm"
	ResourceTypeStack                            ResourceType = "AWS::CloudFormation::Stack"
	ResourceTypeLoadBalancer                     ResourceType = "AWS::ElasticLoadBalancing::LoadBalancer"
	ResourceTypeAutoScalingGroup                 ResourceType = "AWS::AutoScaling::AutoScalingGroup"
	ResourceTypeLaunchConfiguration              ResourceType = "AWS::AutoScaling::LaunchConfiguration"
	ResourceTypeScalingPolicy                    ResourceType = "AWS::AutoScaling::ScalingPolicy"
	ResourceTypeScheduledAction                  ResourceType = "AWS::AutoScaling::ScheduledAction"
	ResourceTypeTable                            ResourceType = "AWS::DynamoDB::Table"
	ResourceTypeProject                          ResourceType = "AWS::CodeBuild::Project"
	ResourceTypeRateBasedRule                    ResourceType = "AWS::WAF::RateBasedRule"
	ResourceTypeRule                             ResourceType = "AWS::WAF::Rule"
	ResourceTypeRuleGroup                        ResourceType = "AWS::WAF::RuleGroup"
	ResourceTypeWebACL                           ResourceType = "AWS::WAF::WebACL"
	ResourceTypeRegionalRateBasedRule            ResourceType = "AWS::WAFRegional::RateBasedRule"
	ResourceTypeRegionalRule                     ResourceType = "AWS::WAFRegional::Rule"
	ResourceTypeRegionalRuleGroup                ResourceType = "AWS::WAFRegional::RuleGroup"
	ResourceTypeRegionalWebACL                   ResourceType = "AWS::WAFRegional::WebACL"
	ResourceTypeDistribution                     ResourceType = "AWS::CloudFront::Distribution"
	ResourceTypeStreamingDistribution            ResourceType = "AWS::CloudFront::StreamingDistribution"
	ResourceTypeFunction                         ResourceType = "AWS::Lambda::Function"
	ResourceTypeNetworkFirewallFirewall          ResourceType = "AWS::NetworkFirewall::Firewall"
	ResourceTypeNetworkFirewallFirewallPolicy    ResourceType = "AWS::NetworkFirewall::FirewallPolicy"
	ResourceTypeNetworkFirewallRuleGroup         ResourceType = "AWS::NetworkFirewall::RuleGroup"
	ResourceTypeApplication                      ResourceType = "AWS::ElasticBeanstalk::Application"
	ResourceTypeApplicationVersion               ResourceType = "AWS::ElasticBeanstalk::ApplicationVersion"
	ResourceTypeEnvironment                      ResourceType = "AWS::ElasticBeanstalk::Environment"
	ResourceTypeWebACLV2                         ResourceType = "AWS::WAFv2::WebACL"
	ResourceTypeRuleGroupV2                      ResourceType = "AWS::WAFv2::RuleGroup"
	ResourceTypeIPSetV2                          ResourceType = "AWS::WAFv2::IPSet"
	ResourceTypeRegexPatternSetV2                ResourceType = "AWS::WAFv2::RegexPatternSet"
	ResourceTypeManagedRuleSetV2                 ResourceType = "AWS::WAFv2::ManagedRuleSet"
	ResourceTypeEncryptionConfig                 ResourceType = "AWS::XRay::EncryptionConfig"
	ResourceTypeAssociationCompliance            ResourceType = "AWS::SSM::AssociationCompliance"
	ResourceTypePatchCompliance                  ResourceType = "AWS::SSM::PatchCompliance"
	ResourceTypeProtection                       ResourceType = "AWS::Shield::Protection"
	ResourceTypeRegionalProtection               ResourceType = "AWS::ShieldRegional::Protection"
	ResourceTypeResourceCompliance               ResourceType = "AWS::Config::ResourceCompliance"
	ResourceTypeStage                            ResourceType = "AWS::ApiGateway::Stage"
	ResourceTypeRestApi                          ResourceType = "AWS::ApiGateway::RestApi"
	ResourceTypeStageV2                          ResourceType = "AWS::ApiGatewayV2::Stage"
	ResourceTypeApi                              ResourceType = "AWS::ApiGatewayV2::Api"
	ResourceTypePipeline                         ResourceType = "AWS::CodePipeline::Pipeline"
	ResourceTypeCloudFormationProvisionedProduct ResourceType = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"
	ResourceTypeCloudFormationProduct            ResourceType = "AWS::ServiceCatalog::CloudFormationProduct"
	ResourceTypePortfolio                        ResourceType = "AWS::ServiceCatalog::Portfolio"
	ResourceTypeQueue                            ResourceType = "AWS::SQS::Queue"
	ResourceTypeKey                              ResourceType = "AWS::KMS::Key"
	ResourceTypeQLDBLedger                       ResourceType = "AWS::QLDB::Ledger"
	ResourceTypeSecret                           ResourceType = "AWS::SecretsManager::Secret"
	ResourceTypeTopic                            ResourceType = "AWS::SNS::Topic"
	ResourceTypeFileData                         ResourceType = "AWS::SSM::FileData"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"AWS::EC2::CustomerGateway",
		"AWS::EC2::EIP",
		"AWS::EC2::Host",
		"AWS::EC2::Instance",
		"AWS::EC2::InternetGateway",
		"AWS::EC2::NetworkAcl",
		"AWS::EC2::NetworkInterface",
		"AWS::EC2::RouteTable",
		"AWS::EC2::SecurityGroup",
		"AWS::EC2::Subnet",
		"AWS::CloudTrail::Trail",
		"AWS::EC2::Volume",
		"AWS::EC2::VPC",
		"AWS::EC2::VPNConnection",
		"AWS::EC2::VPNGateway",
		"AWS::EC2::RegisteredHAInstance",
		"AWS::EC2::NatGateway",
		"AWS::EC2::EgressOnlyInternetGateway",
		"AWS::EC2::VPCEndpoint",
		"AWS::EC2::VPCEndpointService",
		"AWS::EC2::FlowLog",
		"AWS::EC2::VPCPeeringConnection",
		"AWS::Elasticsearch::Domain",
		"AWS::IAM::Group",
		"AWS::IAM::Policy",
		"AWS::IAM::Role",
		"AWS::IAM::User",
		"AWS::ElasticLoadBalancingV2::LoadBalancer",
		"AWS::ACM::Certificate",
		"AWS::RDS::DBInstance",
		"AWS::RDS::DBSubnetGroup",
		"AWS::RDS::DBSecurityGroup",
		"AWS::RDS::DBSnapshot",
		"AWS::RDS::DBCluster",
		"AWS::RDS::DBClusterSnapshot",
		"AWS::RDS::EventSubscription",
		"AWS::S3::Bucket",
		"AWS::S3::AccountPublicAccessBlock",
		"AWS::Redshift::Cluster",
		"AWS::Redshift::ClusterSnapshot",
		"AWS::Redshift::ClusterParameterGroup",
		"AWS::Redshift::ClusterSecurityGroup",
		"AWS::Redshift::ClusterSubnetGroup",
		"AWS::Redshift::EventSubscription",
		"AWS::SSM::ManagedInstanceInventory",
		"AWS::CloudWatch::Alarm",
		"AWS::CloudFormation::Stack",
		"AWS::ElasticLoadBalancing::LoadBalancer",
		"AWS::AutoScaling::AutoScalingGroup",
		"AWS::AutoScaling::LaunchConfiguration",
		"AWS::AutoScaling::ScalingPolicy",
		"AWS::AutoScaling::ScheduledAction",
		"AWS::DynamoDB::Table",
		"AWS::CodeBuild::Project",
		"AWS::WAF::RateBasedRule",
		"AWS::WAF::Rule",
		"AWS::WAF::RuleGroup",
		"AWS::WAF::WebACL",
		"AWS::WAFRegional::RateBasedRule",
		"AWS::WAFRegional::Rule",
		"AWS::WAFRegional::RuleGroup",
		"AWS::WAFRegional::WebACL",
		"AWS::CloudFront::Distribution",
		"AWS::CloudFront::StreamingDistribution",
		"AWS::Lambda::Function",
		"AWS::NetworkFirewall::Firewall",
		"AWS::NetworkFirewall::FirewallPolicy",
		"AWS::NetworkFirewall::RuleGroup",
		"AWS::ElasticBeanstalk::Application",
		"AWS::ElasticBeanstalk::ApplicationVersion",
		"AWS::ElasticBeanstalk::Environment",
		"AWS::WAFv2::WebACL",
		"AWS::WAFv2::RuleGroup",
		"AWS::WAFv2::IPSet",
		"AWS::WAFv2::RegexPatternSet",
		"AWS::WAFv2::ManagedRuleSet",
		"AWS::XRay::EncryptionConfig",
		"AWS::SSM::AssociationCompliance",
		"AWS::SSM::PatchCompliance",
		"AWS::Shield::Protection",
		"AWS::ShieldRegional::Protection",
		"AWS::Config::ResourceCompliance",
		"AWS::ApiGateway::Stage",
		"AWS::ApiGateway::RestApi",
		"AWS::ApiGatewayV2::Stage",
		"AWS::ApiGatewayV2::Api",
		"AWS::CodePipeline::Pipeline",
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct",
		"AWS::ServiceCatalog::CloudFormationProduct",
		"AWS::ServiceCatalog::Portfolio",
		"AWS::SQS::Queue",
		"AWS::KMS::Key",
		"AWS::QLDB::Ledger",
		"AWS::SecretsManager::Secret",
		"AWS::SNS::Topic",
		"AWS::SSM::FileData",
	}
}

type ResourceValueType string

// Enum values for ResourceValueType
const (
	ResourceValueTypeResourceId ResourceValueType = "RESOURCE_ID"
)

// Values returns all known values for ResourceValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceValueType) Values() []ResourceValueType {
	return []ResourceValueType{
		"RESOURCE_ID",
	}
}
