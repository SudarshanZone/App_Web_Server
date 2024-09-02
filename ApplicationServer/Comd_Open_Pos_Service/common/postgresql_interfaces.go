package common

import (
	"gorm.io/gorm"
)


type PostgreSQLConfig struct{
	Username string
	Password string
	DBName string
	Host string
	Port int
	SSLMode string
}

// load the Database
type DBConfigLoader interface{
	LoadPostgreSQLConfig(serviceName string,filename string)(*PostgreSQLConfig,error)
}

// Get the Connection
type DBConnector interface{
	GetDatabaseConnection(servicename string,cfg PostgreSQLConfig)int
}

//Access the Database
type DBAccess interface{
	GetDB(serviceName string) *gorm.DB
}