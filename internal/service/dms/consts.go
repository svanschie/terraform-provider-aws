// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dms

const (
	endpointStatusDeleting = "deleting"

	replicationInstanceStatusAvailable = "available"
	replicationInstanceStatusCreating  = "creating"
	replicationInstanceStatusDeleting  = "deleting"
	replicationInstanceStatusModifying = "modifying"
	replicationInstanceStatusUpgrading = "upgrading"

	replicationTaskStatusCreating  = "creating"
	replicationTaskStatusDeleting  = "deleting"
	replicationTaskStatusFailed    = "failed"
	replicationTaskStatusModifying = "modifying"
	replicationTaskStatusReady     = "ready"
	replicationTaskStatusStopped   = "stopped"
	replicationTaskStatusStopping  = "stopping"
	replicationTaskStatusRunning   = "running"
	replicationTaskStatusStarting  = "starting"
)

const (
	engineNameAurora                     = "aurora"
	engineNameAuroraPostgresql           = "aurora-postgresql"
	engineNameAuroraPostgresqlServerless = "aurora-postgresql-serverless"
	engineNameAuroraServerless           = "aurora-serverless"
	engineNameAzuredb                    = "azuredb"
	engineNameAzureSQLManagedInstance    = "azure-sql-managed-instance"
	engineNameBabelfish                  = "babelfish"
	engineNameDB2                        = "db2"
	engineNameDB2zOS                     = "db2-zos"
	engineNameTransfer                   = "dms-transfer"
	engineNameDocDB                      = "docdb"
	engineNameDynamoDB                   = "dynamodb"
	engineNameElasticsearch              = "elasticsearch"
	engineNameKafka                      = "kafka"
	engineNameKinesis                    = "kinesis"
	engineNameMariadb                    = "mariadb"
	engineNameMongodb                    = "mongodb"
	engineNameMySQL                      = "mysql"
	engineNameNeptune                    = "neptune"
	engineNameOpenSearch                 = "opensearch"
	engineNameOracle                     = "oracle"
	engineNamePostgres                   = "postgres"
	engineNameRedis                      = "redis"
	engineNameRedshift                   = "redshift"
	engineNameS3                         = "s3"
	engineNameSQLServer                  = "sqlserver"
	engineNameSybase                     = "sybase"
)

func engineName_Values() []string {
	return []string{
		engineNameAurora,
		engineNameAuroraPostgresql,
		engineNameAuroraPostgresqlServerless,
		engineNameAuroraServerless,
		engineNameAzuredb,
		engineNameAzureSQLManagedInstance,
		engineNameBabelfish,
		engineNameDB2,
		engineNameDB2zOS,
		engineNameTransfer,
		engineNameDocDB,
		engineNameDynamoDB,
		engineNameElasticsearch,
		engineNameKafka,
		engineNameKinesis,
		engineNameMariadb,
		engineNameMongodb,
		engineNameMySQL,
		engineNameNeptune,
		engineNameOpenSearch,
		engineNameOracle,
		engineNamePostgres,
		engineNameRedis,
		engineNameRedshift,
		engineNameS3,
		engineNameSQLServer,
		engineNameSybase,
	}
}

const (
	kafkaDefaultTopic = "kafka-default-topic"
)

// https://github.com/aws/aws-sdk-go/issues/2522.
const (
	mongoDBAuthMechanismValueDefault   = "default"
	mongoDBAuthMechanismValueMongodbCr = "mongodb-cr"
	mongoDBAuthMechanismValueScramSha1 = "scram-sha-1"
)

func mongoDBAuthMechanismValue_Values() []string {
	return []string{
		mongoDBAuthMechanismValueDefault,
		mongoDBAuthMechanismValueMongodbCr,
		mongoDBAuthMechanismValueScramSha1,
	}
}

const (
	mongoDBAuthSourceAdmin = "admin"
)

const (
	s3SettingsCompressionTypeGzip = "GZIP"
	s3SettingsCompressionTypeNone = "NONE"
)

func s3SettingsCompressionType_Values() []string {
	return []string{
		s3SettingsCompressionTypeGzip,
		s3SettingsCompressionTypeNone,
	}
}

const (
	encryptionModeSseKMS = "SSE_KMS"
	encryptionModeSseS3  = "SSE_S3"
)

func encryptionMode_Values() []string {
	return []string{
		encryptionModeSseKMS,
		encryptionModeSseS3,
	}
}

const (
	networkTypeDual = "DUAL"
	networkTypeIPv4 = "IPV4"
)

func networkType_Values() []string {
	return []string{
		networkTypeDual,
		networkTypeIPv4,
	}
}
