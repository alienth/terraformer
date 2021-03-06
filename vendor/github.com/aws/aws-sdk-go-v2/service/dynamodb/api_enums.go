// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

type AttributeAction string

// Enum values for AttributeAction
const (
	AttributeActionAdd    AttributeAction = "ADD"
	AttributeActionPut    AttributeAction = "PUT"
	AttributeActionDelete AttributeAction = "DELETE"
)

func (enum AttributeAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AttributeAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BackupStatus string

// Enum values for BackupStatus
const (
	BackupStatusCreating  BackupStatus = "CREATING"
	BackupStatusDeleted   BackupStatus = "DELETED"
	BackupStatusAvailable BackupStatus = "AVAILABLE"
)

func (enum BackupStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BackupStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BackupType string

// Enum values for BackupType
const (
	BackupTypeUser      BackupType = "USER"
	BackupTypeSystem    BackupType = "SYSTEM"
	BackupTypeAwsBackup BackupType = "AWS_BACKUP"
)

func (enum BackupType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BackupType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BackupTypeFilter string

// Enum values for BackupTypeFilter
const (
	BackupTypeFilterUser      BackupTypeFilter = "USER"
	BackupTypeFilterSystem    BackupTypeFilter = "SYSTEM"
	BackupTypeFilterAwsBackup BackupTypeFilter = "AWS_BACKUP"
	BackupTypeFilterAll       BackupTypeFilter = "ALL"
)

func (enum BackupTypeFilter) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BackupTypeFilter) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BillingMode string

// Enum values for BillingMode
const (
	BillingModeProvisioned   BillingMode = "PROVISIONED"
	BillingModePayPerRequest BillingMode = "PAY_PER_REQUEST"
)

func (enum BillingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BillingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq          ComparisonOperator = "EQ"
	ComparisonOperatorNe          ComparisonOperator = "NE"
	ComparisonOperatorIn          ComparisonOperator = "IN"
	ComparisonOperatorLe          ComparisonOperator = "LE"
	ComparisonOperatorLt          ComparisonOperator = "LT"
	ComparisonOperatorGe          ComparisonOperator = "GE"
	ComparisonOperatorGt          ComparisonOperator = "GT"
	ComparisonOperatorBetween     ComparisonOperator = "BETWEEN"
	ComparisonOperatorNotNull     ComparisonOperator = "NOT_NULL"
	ComparisonOperatorNull        ComparisonOperator = "NULL"
	ComparisonOperatorContains    ComparisonOperator = "CONTAINS"
	ComparisonOperatorNotContains ComparisonOperator = "NOT_CONTAINS"
	ComparisonOperatorBeginsWith  ComparisonOperator = "BEGINS_WITH"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConditionalOperator string

// Enum values for ConditionalOperator
const (
	ConditionalOperatorAnd ConditionalOperator = "AND"
	ConditionalOperatorOr  ConditionalOperator = "OR"
)

func (enum ConditionalOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConditionalOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContinuousBackupsStatus string

// Enum values for ContinuousBackupsStatus
const (
	ContinuousBackupsStatusEnabled  ContinuousBackupsStatus = "ENABLED"
	ContinuousBackupsStatusDisabled ContinuousBackupsStatus = "DISABLED"
)

func (enum ContinuousBackupsStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContinuousBackupsStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContributorInsightsAction string

// Enum values for ContributorInsightsAction
const (
	ContributorInsightsActionEnable  ContributorInsightsAction = "ENABLE"
	ContributorInsightsActionDisable ContributorInsightsAction = "DISABLE"
)

func (enum ContributorInsightsAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContributorInsightsAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContributorInsightsStatus string

// Enum values for ContributorInsightsStatus
const (
	ContributorInsightsStatusEnabling  ContributorInsightsStatus = "ENABLING"
	ContributorInsightsStatusEnabled   ContributorInsightsStatus = "ENABLED"
	ContributorInsightsStatusDisabling ContributorInsightsStatus = "DISABLING"
	ContributorInsightsStatusDisabled  ContributorInsightsStatus = "DISABLED"
	ContributorInsightsStatusFailed    ContributorInsightsStatus = "FAILED"
)

func (enum ContributorInsightsStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContributorInsightsStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GlobalTableStatus string

// Enum values for GlobalTableStatus
const (
	GlobalTableStatusCreating GlobalTableStatus = "CREATING"
	GlobalTableStatusActive   GlobalTableStatus = "ACTIVE"
	GlobalTableStatusDeleting GlobalTableStatus = "DELETING"
	GlobalTableStatusUpdating GlobalTableStatus = "UPDATING"
)

func (enum GlobalTableStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GlobalTableStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IndexStatus string

// Enum values for IndexStatus
const (
	IndexStatusCreating IndexStatus = "CREATING"
	IndexStatusUpdating IndexStatus = "UPDATING"
	IndexStatusDeleting IndexStatus = "DELETING"
	IndexStatusActive   IndexStatus = "ACTIVE"
)

func (enum IndexStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IndexStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

func (enum KeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PointInTimeRecoveryStatus string

// Enum values for PointInTimeRecoveryStatus
const (
	PointInTimeRecoveryStatusEnabled  PointInTimeRecoveryStatus = "ENABLED"
	PointInTimeRecoveryStatusDisabled PointInTimeRecoveryStatus = "DISABLED"
)

func (enum PointInTimeRecoveryStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PointInTimeRecoveryStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProjectionType string

// Enum values for ProjectionType
const (
	ProjectionTypeAll      ProjectionType = "ALL"
	ProjectionTypeKeysOnly ProjectionType = "KEYS_ONLY"
	ProjectionTypeInclude  ProjectionType = "INCLUDE"
)

func (enum ProjectionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProjectionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReplicaStatus string

// Enum values for ReplicaStatus
const (
	ReplicaStatusCreating       ReplicaStatus = "CREATING"
	ReplicaStatusCreationFailed ReplicaStatus = "CREATION_FAILED"
	ReplicaStatusUpdating       ReplicaStatus = "UPDATING"
	ReplicaStatusDeleting       ReplicaStatus = "DELETING"
	ReplicaStatusActive         ReplicaStatus = "ACTIVE"
)

func (enum ReplicaStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReplicaStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// Determines the level of detail about provisioned throughput consumption that
// is returned in the response:
//
//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
//    operation, together with ConsumedCapacity for each table and secondary
//    index that was accessed. Note that some operations, such as GetItem and
//    BatchGetItem, do not access any indexes at all. In these cases, specifying
//    INDEXES will only return ConsumedCapacity information for table(s).
//
//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
//    the operation.
//
//    * NONE - No ConsumedCapacity details are included in the response.
type ReturnConsumedCapacity string

// Enum values for ReturnConsumedCapacity
const (
	ReturnConsumedCapacityIndexes ReturnConsumedCapacity = "INDEXES"
	ReturnConsumedCapacityTotal   ReturnConsumedCapacity = "TOTAL"
	ReturnConsumedCapacityNone    ReturnConsumedCapacity = "NONE"
)

func (enum ReturnConsumedCapacity) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReturnConsumedCapacity) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReturnItemCollectionMetrics string

// Enum values for ReturnItemCollectionMetrics
const (
	ReturnItemCollectionMetricsSize ReturnItemCollectionMetrics = "SIZE"
	ReturnItemCollectionMetricsNone ReturnItemCollectionMetrics = "NONE"
)

func (enum ReturnItemCollectionMetrics) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReturnItemCollectionMetrics) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReturnValue string

// Enum values for ReturnValue
const (
	ReturnValueNone       ReturnValue = "NONE"
	ReturnValueAllOld     ReturnValue = "ALL_OLD"
	ReturnValueUpdatedOld ReturnValue = "UPDATED_OLD"
	ReturnValueAllNew     ReturnValue = "ALL_NEW"
	ReturnValueUpdatedNew ReturnValue = "UPDATED_NEW"
)

func (enum ReturnValue) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReturnValue) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReturnValuesOnConditionCheckFailure string

// Enum values for ReturnValuesOnConditionCheckFailure
const (
	ReturnValuesOnConditionCheckFailureAllOld ReturnValuesOnConditionCheckFailure = "ALL_OLD"
	ReturnValuesOnConditionCheckFailureNone   ReturnValuesOnConditionCheckFailure = "NONE"
)

func (enum ReturnValuesOnConditionCheckFailure) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReturnValuesOnConditionCheckFailure) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SSEStatus string

// Enum values for SSEStatus
const (
	SSEStatusEnabling  SSEStatus = "ENABLING"
	SSEStatusEnabled   SSEStatus = "ENABLED"
	SSEStatusDisabling SSEStatus = "DISABLING"
	SSEStatusDisabled  SSEStatus = "DISABLED"
	SSEStatusUpdating  SSEStatus = "UPDATING"
)

func (enum SSEStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SSEStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SSEType string

// Enum values for SSEType
const (
	SSETypeAes256 SSEType = "AES256"
	SSETypeKms    SSEType = "KMS"
)

func (enum SSEType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SSEType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalarAttributeType string

// Enum values for ScalarAttributeType
const (
	ScalarAttributeTypeS ScalarAttributeType = "S"
	ScalarAttributeTypeN ScalarAttributeType = "N"
	ScalarAttributeTypeB ScalarAttributeType = "B"
)

func (enum ScalarAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalarAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Select string

// Enum values for Select
const (
	SelectAllAttributes          Select = "ALL_ATTRIBUTES"
	SelectAllProjectedAttributes Select = "ALL_PROJECTED_ATTRIBUTES"
	SelectSpecificAttributes     Select = "SPECIFIC_ATTRIBUTES"
	SelectCount                  Select = "COUNT"
)

func (enum Select) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Select) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamViewType string

// Enum values for StreamViewType
const (
	StreamViewTypeNewImage        StreamViewType = "NEW_IMAGE"
	StreamViewTypeOldImage        StreamViewType = "OLD_IMAGE"
	StreamViewTypeNewAndOldImages StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewTypeKeysOnly        StreamViewType = "KEYS_ONLY"
)

func (enum StreamViewType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamViewType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TableStatus string

// Enum values for TableStatus
const (
	TableStatusCreating                          TableStatus = "CREATING"
	TableStatusUpdating                          TableStatus = "UPDATING"
	TableStatusDeleting                          TableStatus = "DELETING"
	TableStatusActive                            TableStatus = "ACTIVE"
	TableStatusInaccessibleEncryptionCredentials TableStatus = "INACCESSIBLE_ENCRYPTION_CREDENTIALS"
	TableStatusArchiving                         TableStatus = "ARCHIVING"
	TableStatusArchived                          TableStatus = "ARCHIVED"
)

func (enum TableStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TableStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TimeToLiveStatus string

// Enum values for TimeToLiveStatus
const (
	TimeToLiveStatusEnabling  TimeToLiveStatus = "ENABLING"
	TimeToLiveStatusDisabling TimeToLiveStatus = "DISABLING"
	TimeToLiveStatusEnabled   TimeToLiveStatus = "ENABLED"
	TimeToLiveStatusDisabled  TimeToLiveStatus = "DISABLED"
)

func (enum TimeToLiveStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TimeToLiveStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
