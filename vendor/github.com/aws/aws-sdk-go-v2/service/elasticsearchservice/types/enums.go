// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AutoTuneDesiredState string

// Enum values for AutoTuneDesiredState
const (
	AutoTuneDesiredStateEnabled  AutoTuneDesiredState = "ENABLED"
	AutoTuneDesiredStateDisabled AutoTuneDesiredState = "DISABLED"
)

// Values returns all known values for AutoTuneDesiredState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoTuneDesiredState) Values() []AutoTuneDesiredState {
	return []AutoTuneDesiredState{
		"ENABLED",
		"DISABLED",
	}
}

type AutoTuneState string

// Enum values for AutoTuneState
const (
	AutoTuneStateEnabled                       AutoTuneState = "ENABLED"
	AutoTuneStateDisabled                      AutoTuneState = "DISABLED"
	AutoTuneStateEnableInProgress              AutoTuneState = "ENABLE_IN_PROGRESS"
	AutoTuneStateDisableInProgress             AutoTuneState = "DISABLE_IN_PROGRESS"
	AutoTuneStateDisabledAndRollbackScheduled  AutoTuneState = "DISABLED_AND_ROLLBACK_SCHEDULED"
	AutoTuneStateDisabledAndRollbackInProgress AutoTuneState = "DISABLED_AND_ROLLBACK_IN_PROGRESS"
	AutoTuneStateDisabledAndRollbackComplete   AutoTuneState = "DISABLED_AND_ROLLBACK_COMPLETE"
	AutoTuneStateDisabledAndRollbackError      AutoTuneState = "DISABLED_AND_ROLLBACK_ERROR"
	AutoTuneStateError                         AutoTuneState = "ERROR"
)

// Values returns all known values for AutoTuneState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoTuneState) Values() []AutoTuneState {
	return []AutoTuneState{
		"ENABLED",
		"DISABLED",
		"ENABLE_IN_PROGRESS",
		"DISABLE_IN_PROGRESS",
		"DISABLED_AND_ROLLBACK_SCHEDULED",
		"DISABLED_AND_ROLLBACK_IN_PROGRESS",
		"DISABLED_AND_ROLLBACK_COMPLETE",
		"DISABLED_AND_ROLLBACK_ERROR",
		"ERROR",
	}
}

type AutoTuneType string

// Enum values for AutoTuneType
const (
	AutoTuneTypeScheduledAction AutoTuneType = "SCHEDULED_ACTION"
)

// Values returns all known values for AutoTuneType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AutoTuneType) Values() []AutoTuneType {
	return []AutoTuneType{
		"SCHEDULED_ACTION",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusPendingUpdate DeploymentStatus = "PENDING_UPDATE"
	DeploymentStatusInProgress    DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusCompleted     DeploymentStatus = "COMPLETED"
	DeploymentStatusNotEligible   DeploymentStatus = "NOT_ELIGIBLE"
	DeploymentStatusEligible      DeploymentStatus = "ELIGIBLE"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"PENDING_UPDATE",
		"IN_PROGRESS",
		"COMPLETED",
		"NOT_ELIGIBLE",
		"ELIGIBLE",
	}
}

type DescribePackagesFilterName string

// Enum values for DescribePackagesFilterName
const (
	DescribePackagesFilterNamePackageID     DescribePackagesFilterName = "PackageID"
	DescribePackagesFilterNamePackageName   DescribePackagesFilterName = "PackageName"
	DescribePackagesFilterNamePackageStatus DescribePackagesFilterName = "PackageStatus"
)

// Values returns all known values for DescribePackagesFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DescribePackagesFilterName) Values() []DescribePackagesFilterName {
	return []DescribePackagesFilterName{
		"PackageID",
		"PackageName",
		"PackageStatus",
	}
}

type DomainPackageStatus string

// Enum values for DomainPackageStatus
const (
	DomainPackageStatusAssociating        DomainPackageStatus = "ASSOCIATING"
	DomainPackageStatusAssociationFailed  DomainPackageStatus = "ASSOCIATION_FAILED"
	DomainPackageStatusActive             DomainPackageStatus = "ACTIVE"
	DomainPackageStatusDissociating       DomainPackageStatus = "DISSOCIATING"
	DomainPackageStatusDissociationFailed DomainPackageStatus = "DISSOCIATION_FAILED"
)

// Values returns all known values for DomainPackageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainPackageStatus) Values() []DomainPackageStatus {
	return []DomainPackageStatus{
		"ASSOCIATING",
		"ASSOCIATION_FAILED",
		"ACTIVE",
		"DISSOCIATING",
		"DISSOCIATION_FAILED",
	}
}

type EngineType string

// Enum values for EngineType
const (
	EngineTypeOpenSearch    EngineType = "OpenSearch"
	EngineTypeElasticsearch EngineType = "Elasticsearch"
)

// Values returns all known values for EngineType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EngineType) Values() []EngineType {
	return []EngineType{
		"OpenSearch",
		"Elasticsearch",
	}
}

type ESPartitionInstanceType string

// Enum values for ESPartitionInstanceType
const (
	ESPartitionInstanceTypeM3MediumElasticsearch         ESPartitionInstanceType = "m3.medium.elasticsearch"
	ESPartitionInstanceTypeM3LargeElasticsearch          ESPartitionInstanceType = "m3.large.elasticsearch"
	ESPartitionInstanceTypeM3XlargeElasticsearch         ESPartitionInstanceType = "m3.xlarge.elasticsearch"
	ESPartitionInstanceTypeM32xlargeElasticsearch        ESPartitionInstanceType = "m3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM4LargeElasticsearch          ESPartitionInstanceType = "m4.large.elasticsearch"
	ESPartitionInstanceTypeM4XlargeElasticsearch         ESPartitionInstanceType = "m4.xlarge.elasticsearch"
	ESPartitionInstanceTypeM42xlargeElasticsearch        ESPartitionInstanceType = "m4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM44xlargeElasticsearch        ESPartitionInstanceType = "m4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeM410xlargeElasticsearch       ESPartitionInstanceType = "m4.10xlarge.elasticsearch"
	ESPartitionInstanceTypeM5LargeElasticsearch          ESPartitionInstanceType = "m5.large.elasticsearch"
	ESPartitionInstanceTypeM5XlargeElasticsearch         ESPartitionInstanceType = "m5.xlarge.elasticsearch"
	ESPartitionInstanceTypeM52xlargeElasticsearch        ESPartitionInstanceType = "m5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM54xlargeElasticsearch        ESPartitionInstanceType = "m5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeM512xlargeElasticsearch       ESPartitionInstanceType = "m5.12xlarge.elasticsearch"
	ESPartitionInstanceTypeR5LargeElasticsearch          ESPartitionInstanceType = "r5.large.elasticsearch"
	ESPartitionInstanceTypeR5XlargeElasticsearch         ESPartitionInstanceType = "r5.xlarge.elasticsearch"
	ESPartitionInstanceTypeR52xlargeElasticsearch        ESPartitionInstanceType = "r5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR54xlargeElasticsearch        ESPartitionInstanceType = "r5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR512xlargeElasticsearch       ESPartitionInstanceType = "r5.12xlarge.elasticsearch"
	ESPartitionInstanceTypeC5LargeElasticsearch          ESPartitionInstanceType = "c5.large.elasticsearch"
	ESPartitionInstanceTypeC5XlargeElasticsearch         ESPartitionInstanceType = "c5.xlarge.elasticsearch"
	ESPartitionInstanceTypeC52xlargeElasticsearch        ESPartitionInstanceType = "c5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeC54xlargeElasticsearch        ESPartitionInstanceType = "c5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeC59xlargeElasticsearch        ESPartitionInstanceType = "c5.9xlarge.elasticsearch"
	ESPartitionInstanceTypeC518xlargeElasticsearch       ESPartitionInstanceType = "c5.18xlarge.elasticsearch"
	ESPartitionInstanceTypeUltrawarm1MediumElasticsearch ESPartitionInstanceType = "ultrawarm1.medium.elasticsearch"
	ESPartitionInstanceTypeUltrawarm1LargeElasticsearch  ESPartitionInstanceType = "ultrawarm1.large.elasticsearch"
	ESPartitionInstanceTypeT2MicroElasticsearch          ESPartitionInstanceType = "t2.micro.elasticsearch"
	ESPartitionInstanceTypeT2SmallElasticsearch          ESPartitionInstanceType = "t2.small.elasticsearch"
	ESPartitionInstanceTypeT2MediumElasticsearch         ESPartitionInstanceType = "t2.medium.elasticsearch"
	ESPartitionInstanceTypeR3LargeElasticsearch          ESPartitionInstanceType = "r3.large.elasticsearch"
	ESPartitionInstanceTypeR3XlargeElasticsearch         ESPartitionInstanceType = "r3.xlarge.elasticsearch"
	ESPartitionInstanceTypeR32xlargeElasticsearch        ESPartitionInstanceType = "r3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR34xlargeElasticsearch        ESPartitionInstanceType = "r3.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR38xlargeElasticsearch        ESPartitionInstanceType = "r3.8xlarge.elasticsearch"
	ESPartitionInstanceTypeI2XlargeElasticsearch         ESPartitionInstanceType = "i2.xlarge.elasticsearch"
	ESPartitionInstanceTypeI22xlargeElasticsearch        ESPartitionInstanceType = "i2.2xlarge.elasticsearch"
	ESPartitionInstanceTypeD2XlargeElasticsearch         ESPartitionInstanceType = "d2.xlarge.elasticsearch"
	ESPartitionInstanceTypeD22xlargeElasticsearch        ESPartitionInstanceType = "d2.2xlarge.elasticsearch"
	ESPartitionInstanceTypeD24xlargeElasticsearch        ESPartitionInstanceType = "d2.4xlarge.elasticsearch"
	ESPartitionInstanceTypeD28xlargeElasticsearch        ESPartitionInstanceType = "d2.8xlarge.elasticsearch"
	ESPartitionInstanceTypeC4LargeElasticsearch          ESPartitionInstanceType = "c4.large.elasticsearch"
	ESPartitionInstanceTypeC4XlargeElasticsearch         ESPartitionInstanceType = "c4.xlarge.elasticsearch"
	ESPartitionInstanceTypeC42xlargeElasticsearch        ESPartitionInstanceType = "c4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeC44xlargeElasticsearch        ESPartitionInstanceType = "c4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeC48xlargeElasticsearch        ESPartitionInstanceType = "c4.8xlarge.elasticsearch"
	ESPartitionInstanceTypeR4LargeElasticsearch          ESPartitionInstanceType = "r4.large.elasticsearch"
	ESPartitionInstanceTypeR4XlargeElasticsearch         ESPartitionInstanceType = "r4.xlarge.elasticsearch"
	ESPartitionInstanceTypeR42xlargeElasticsearch        ESPartitionInstanceType = "r4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR44xlargeElasticsearch        ESPartitionInstanceType = "r4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR48xlargeElasticsearch        ESPartitionInstanceType = "r4.8xlarge.elasticsearch"
	ESPartitionInstanceTypeR416xlargeElasticsearch       ESPartitionInstanceType = "r4.16xlarge.elasticsearch"
	ESPartitionInstanceTypeI3LargeElasticsearch          ESPartitionInstanceType = "i3.large.elasticsearch"
	ESPartitionInstanceTypeI3XlargeElasticsearch         ESPartitionInstanceType = "i3.xlarge.elasticsearch"
	ESPartitionInstanceTypeI32xlargeElasticsearch        ESPartitionInstanceType = "i3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeI34xlargeElasticsearch        ESPartitionInstanceType = "i3.4xlarge.elasticsearch"
	ESPartitionInstanceTypeI38xlargeElasticsearch        ESPartitionInstanceType = "i3.8xlarge.elasticsearch"
	ESPartitionInstanceTypeI316xlargeElasticsearch       ESPartitionInstanceType = "i3.16xlarge.elasticsearch"
)

// Values returns all known values for ESPartitionInstanceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ESPartitionInstanceType) Values() []ESPartitionInstanceType {
	return []ESPartitionInstanceType{
		"m3.medium.elasticsearch",
		"m3.large.elasticsearch",
		"m3.xlarge.elasticsearch",
		"m3.2xlarge.elasticsearch",
		"m4.large.elasticsearch",
		"m4.xlarge.elasticsearch",
		"m4.2xlarge.elasticsearch",
		"m4.4xlarge.elasticsearch",
		"m4.10xlarge.elasticsearch",
		"m5.large.elasticsearch",
		"m5.xlarge.elasticsearch",
		"m5.2xlarge.elasticsearch",
		"m5.4xlarge.elasticsearch",
		"m5.12xlarge.elasticsearch",
		"r5.large.elasticsearch",
		"r5.xlarge.elasticsearch",
		"r5.2xlarge.elasticsearch",
		"r5.4xlarge.elasticsearch",
		"r5.12xlarge.elasticsearch",
		"c5.large.elasticsearch",
		"c5.xlarge.elasticsearch",
		"c5.2xlarge.elasticsearch",
		"c5.4xlarge.elasticsearch",
		"c5.9xlarge.elasticsearch",
		"c5.18xlarge.elasticsearch",
		"ultrawarm1.medium.elasticsearch",
		"ultrawarm1.large.elasticsearch",
		"t2.micro.elasticsearch",
		"t2.small.elasticsearch",
		"t2.medium.elasticsearch",
		"r3.large.elasticsearch",
		"r3.xlarge.elasticsearch",
		"r3.2xlarge.elasticsearch",
		"r3.4xlarge.elasticsearch",
		"r3.8xlarge.elasticsearch",
		"i2.xlarge.elasticsearch",
		"i2.2xlarge.elasticsearch",
		"d2.xlarge.elasticsearch",
		"d2.2xlarge.elasticsearch",
		"d2.4xlarge.elasticsearch",
		"d2.8xlarge.elasticsearch",
		"c4.large.elasticsearch",
		"c4.xlarge.elasticsearch",
		"c4.2xlarge.elasticsearch",
		"c4.4xlarge.elasticsearch",
		"c4.8xlarge.elasticsearch",
		"r4.large.elasticsearch",
		"r4.xlarge.elasticsearch",
		"r4.2xlarge.elasticsearch",
		"r4.4xlarge.elasticsearch",
		"r4.8xlarge.elasticsearch",
		"r4.16xlarge.elasticsearch",
		"i3.large.elasticsearch",
		"i3.xlarge.elasticsearch",
		"i3.2xlarge.elasticsearch",
		"i3.4xlarge.elasticsearch",
		"i3.8xlarge.elasticsearch",
		"i3.16xlarge.elasticsearch",
	}
}

type ESWarmPartitionInstanceType string

// Enum values for ESWarmPartitionInstanceType
const (
	ESWarmPartitionInstanceTypeUltrawarm1MediumElasticsearch ESWarmPartitionInstanceType = "ultrawarm1.medium.elasticsearch"
	ESWarmPartitionInstanceTypeUltrawarm1LargeElasticsearch  ESWarmPartitionInstanceType = "ultrawarm1.large.elasticsearch"
)

// Values returns all known values for ESWarmPartitionInstanceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ESWarmPartitionInstanceType) Values() []ESWarmPartitionInstanceType {
	return []ESWarmPartitionInstanceType{
		"ultrawarm1.medium.elasticsearch",
		"ultrawarm1.large.elasticsearch",
	}
}

type InboundCrossClusterSearchConnectionStatusCode string

// Enum values for InboundCrossClusterSearchConnectionStatusCode
const (
	InboundCrossClusterSearchConnectionStatusCodePendingAcceptance InboundCrossClusterSearchConnectionStatusCode = "PENDING_ACCEPTANCE"
	InboundCrossClusterSearchConnectionStatusCodeApproved          InboundCrossClusterSearchConnectionStatusCode = "APPROVED"
	InboundCrossClusterSearchConnectionStatusCodeRejecting         InboundCrossClusterSearchConnectionStatusCode = "REJECTING"
	InboundCrossClusterSearchConnectionStatusCodeRejected          InboundCrossClusterSearchConnectionStatusCode = "REJECTED"
	InboundCrossClusterSearchConnectionStatusCodeDeleting          InboundCrossClusterSearchConnectionStatusCode = "DELETING"
	InboundCrossClusterSearchConnectionStatusCodeDeleted           InboundCrossClusterSearchConnectionStatusCode = "DELETED"
)

// Values returns all known values for
// InboundCrossClusterSearchConnectionStatusCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (InboundCrossClusterSearchConnectionStatusCode) Values() []InboundCrossClusterSearchConnectionStatusCode {
	return []InboundCrossClusterSearchConnectionStatusCode{
		"PENDING_ACCEPTANCE",
		"APPROVED",
		"REJECTING",
		"REJECTED",
		"DELETING",
		"DELETED",
	}
}

type LogType string

// Enum values for LogType
const (
	LogTypeIndexSlowLogs     LogType = "INDEX_SLOW_LOGS"
	LogTypeSearchSlowLogs    LogType = "SEARCH_SLOW_LOGS"
	LogTypeEsApplicationLogs LogType = "ES_APPLICATION_LOGS"
	LogTypeAuditLogs         LogType = "AUDIT_LOGS"
)

// Values returns all known values for LogType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogType) Values() []LogType {
	return []LogType{
		"INDEX_SLOW_LOGS",
		"SEARCH_SLOW_LOGS",
		"ES_APPLICATION_LOGS",
		"AUDIT_LOGS",
	}
}

type OptionState string

// Enum values for OptionState
const (
	OptionStateRequiresIndexDocuments OptionState = "RequiresIndexDocuments"
	OptionStateProcessing             OptionState = "Processing"
	OptionStateActive                 OptionState = "Active"
)

// Values returns all known values for OptionState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OptionState) Values() []OptionState {
	return []OptionState{
		"RequiresIndexDocuments",
		"Processing",
		"Active",
	}
}

type OutboundCrossClusterSearchConnectionStatusCode string

// Enum values for OutboundCrossClusterSearchConnectionStatusCode
const (
	OutboundCrossClusterSearchConnectionStatusCodePendingAcceptance OutboundCrossClusterSearchConnectionStatusCode = "PENDING_ACCEPTANCE"
	OutboundCrossClusterSearchConnectionStatusCodeValidating        OutboundCrossClusterSearchConnectionStatusCode = "VALIDATING"
	OutboundCrossClusterSearchConnectionStatusCodeValidationFailed  OutboundCrossClusterSearchConnectionStatusCode = "VALIDATION_FAILED"
	OutboundCrossClusterSearchConnectionStatusCodeProvisioning      OutboundCrossClusterSearchConnectionStatusCode = "PROVISIONING"
	OutboundCrossClusterSearchConnectionStatusCodeActive            OutboundCrossClusterSearchConnectionStatusCode = "ACTIVE"
	OutboundCrossClusterSearchConnectionStatusCodeRejected          OutboundCrossClusterSearchConnectionStatusCode = "REJECTED"
	OutboundCrossClusterSearchConnectionStatusCodeDeleting          OutboundCrossClusterSearchConnectionStatusCode = "DELETING"
	OutboundCrossClusterSearchConnectionStatusCodeDeleted           OutboundCrossClusterSearchConnectionStatusCode = "DELETED"
)

// Values returns all known values for
// OutboundCrossClusterSearchConnectionStatusCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutboundCrossClusterSearchConnectionStatusCode) Values() []OutboundCrossClusterSearchConnectionStatusCode {
	return []OutboundCrossClusterSearchConnectionStatusCode{
		"PENDING_ACCEPTANCE",
		"VALIDATING",
		"VALIDATION_FAILED",
		"PROVISIONING",
		"ACTIVE",
		"REJECTED",
		"DELETING",
		"DELETED",
	}
}

type OverallChangeStatus string

// Enum values for OverallChangeStatus
const (
	OverallChangeStatusPending    OverallChangeStatus = "PENDING"
	OverallChangeStatusProcessing OverallChangeStatus = "PROCESSING"
	OverallChangeStatusCompleted  OverallChangeStatus = "COMPLETED"
	OverallChangeStatusFailed     OverallChangeStatus = "FAILED"
)

// Values returns all known values for OverallChangeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OverallChangeStatus) Values() []OverallChangeStatus {
	return []OverallChangeStatus{
		"PENDING",
		"PROCESSING",
		"COMPLETED",
		"FAILED",
	}
}

type PackageStatus string

// Enum values for PackageStatus
const (
	PackageStatusCopying          PackageStatus = "COPYING"
	PackageStatusCopyFailed       PackageStatus = "COPY_FAILED"
	PackageStatusValidating       PackageStatus = "VALIDATING"
	PackageStatusValidationFailed PackageStatus = "VALIDATION_FAILED"
	PackageStatusAvailable        PackageStatus = "AVAILABLE"
	PackageStatusDeleting         PackageStatus = "DELETING"
	PackageStatusDeleted          PackageStatus = "DELETED"
	PackageStatusDeleteFailed     PackageStatus = "DELETE_FAILED"
)

// Values returns all known values for PackageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageStatus) Values() []PackageStatus {
	return []PackageStatus{
		"COPYING",
		"COPY_FAILED",
		"VALIDATING",
		"VALIDATION_FAILED",
		"AVAILABLE",
		"DELETING",
		"DELETED",
		"DELETE_FAILED",
	}
}

type PackageType string

// Enum values for PackageType
const (
	PackageTypeTxtDictionary PackageType = "TXT-DICTIONARY"
)

// Values returns all known values for PackageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PackageType) Values() []PackageType {
	return []PackageType{
		"TXT-DICTIONARY",
	}
}

type ReservedElasticsearchInstancePaymentOption string

// Enum values for ReservedElasticsearchInstancePaymentOption
const (
	ReservedElasticsearchInstancePaymentOptionAllUpfront     ReservedElasticsearchInstancePaymentOption = "ALL_UPFRONT"
	ReservedElasticsearchInstancePaymentOptionPartialUpfront ReservedElasticsearchInstancePaymentOption = "PARTIAL_UPFRONT"
	ReservedElasticsearchInstancePaymentOptionNoUpfront      ReservedElasticsearchInstancePaymentOption = "NO_UPFRONT"
)

// Values returns all known values for ReservedElasticsearchInstancePaymentOption.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReservedElasticsearchInstancePaymentOption) Values() []ReservedElasticsearchInstancePaymentOption {
	return []ReservedElasticsearchInstancePaymentOption{
		"ALL_UPFRONT",
		"PARTIAL_UPFRONT",
		"NO_UPFRONT",
	}
}

type RollbackOnDisable string

// Enum values for RollbackOnDisable
const (
	RollbackOnDisableNoRollback      RollbackOnDisable = "NO_ROLLBACK"
	RollbackOnDisableDefaultRollback RollbackOnDisable = "DEFAULT_ROLLBACK"
)

// Values returns all known values for RollbackOnDisable. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RollbackOnDisable) Values() []RollbackOnDisable {
	return []RollbackOnDisable{
		"NO_ROLLBACK",
		"DEFAULT_ROLLBACK",
	}
}

type ScheduledAutoTuneActionType string

// Enum values for ScheduledAutoTuneActionType
const (
	ScheduledAutoTuneActionTypeJvmHeapSizeTuning ScheduledAutoTuneActionType = "JVM_HEAP_SIZE_TUNING"
	ScheduledAutoTuneActionTypeJvmYoungGenTuning ScheduledAutoTuneActionType = "JVM_YOUNG_GEN_TUNING"
)

// Values returns all known values for ScheduledAutoTuneActionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledAutoTuneActionType) Values() []ScheduledAutoTuneActionType {
	return []ScheduledAutoTuneActionType{
		"JVM_HEAP_SIZE_TUNING",
		"JVM_YOUNG_GEN_TUNING",
	}
}

type ScheduledAutoTuneSeverityType string

// Enum values for ScheduledAutoTuneSeverityType
const (
	ScheduledAutoTuneSeverityTypeLow    ScheduledAutoTuneSeverityType = "LOW"
	ScheduledAutoTuneSeverityTypeMedium ScheduledAutoTuneSeverityType = "MEDIUM"
	ScheduledAutoTuneSeverityTypeHigh   ScheduledAutoTuneSeverityType = "HIGH"
)

// Values returns all known values for ScheduledAutoTuneSeverityType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ScheduledAutoTuneSeverityType) Values() []ScheduledAutoTuneSeverityType {
	return []ScheduledAutoTuneSeverityType{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitHours TimeUnit = "HOURS"
)

// Values returns all known values for TimeUnit. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TimeUnit) Values() []TimeUnit {
	return []TimeUnit{
		"HOURS",
	}
}

type TLSSecurityPolicy string

// Enum values for TLSSecurityPolicy
const (
	TLSSecurityPolicyPolicyMinTls10201907 TLSSecurityPolicy = "Policy-Min-TLS-1-0-2019-07"
	TLSSecurityPolicyPolicyMinTls12201907 TLSSecurityPolicy = "Policy-Min-TLS-1-2-2019-07"
)

// Values returns all known values for TLSSecurityPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TLSSecurityPolicy) Values() []TLSSecurityPolicy {
	return []TLSSecurityPolicy{
		"Policy-Min-TLS-1-0-2019-07",
		"Policy-Min-TLS-1-2-2019-07",
	}
}

type UpgradeStatus string

// Enum values for UpgradeStatus
const (
	UpgradeStatusInProgress          UpgradeStatus = "IN_PROGRESS"
	UpgradeStatusSucceeded           UpgradeStatus = "SUCCEEDED"
	UpgradeStatusSucceededWithIssues UpgradeStatus = "SUCCEEDED_WITH_ISSUES"
	UpgradeStatusFailed              UpgradeStatus = "FAILED"
)

// Values returns all known values for UpgradeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpgradeStatus) Values() []UpgradeStatus {
	return []UpgradeStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"SUCCEEDED_WITH_ISSUES",
		"FAILED",
	}
}

type UpgradeStep string

// Enum values for UpgradeStep
const (
	UpgradeStepPreUpgradeCheck UpgradeStep = "PRE_UPGRADE_CHECK"
	UpgradeStepSnapshot        UpgradeStep = "SNAPSHOT"
	UpgradeStepUpgrade         UpgradeStep = "UPGRADE"
)

// Values returns all known values for UpgradeStep. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UpgradeStep) Values() []UpgradeStep {
	return []UpgradeStep{
		"PRE_UPGRADE_CHECK",
		"SNAPSHOT",
		"UPGRADE",
	}
}

type VolumeType string

// Enum values for VolumeType
const (
	VolumeTypeStandard VolumeType = "standard"
	VolumeTypeGp2      VolumeType = "gp2"
	VolumeTypeIo1      VolumeType = "io1"
	VolumeTypeGp3      VolumeType = "gp3"
)

// Values returns all known values for VolumeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VolumeType) Values() []VolumeType {
	return []VolumeType{
		"standard",
		"gp2",
		"io1",
		"gp3",
	}
}
