// ApplicationServer\Comd_Open_Pos_Service\config\PosgresqlConfig.go
package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"github.com/SudarshanZone/Commo_Open_Pos/common"
	"github.com/SudarshanZone/Commo_Open_Pos/internal/models"
	"gorm.io/gorm"
)

// Implmentation of Classes Defined DBConfigLoader,DBConnector,DBAccessor
type ConfigManager struct {
	database struct {
		Db *gorm.DB
	}
}
func (cm *ConfigManager) LoadPostgreSQLConfig(serviceName string, fileName string) (*common.PostgreSQLConfig, error) {
	environmentManager := &models.EnvironmentManager{}

	resultTmp := environmentManager.InitProcessSpace(serviceName, fileName, "database")
	if resultTmp != 0 {
		log.Printf("[%s] Failed to read config file: %v", serviceName, resultTmp)
		return nil, fmt.Errorf("[%s] Failed to read config file: %v", serviceName, resultTmp)
	}

	config := &common.PostgreSQLConfig{
		Host:     environmentManager.GetProcessSpaceValue("host"),
		Port:     environmentManager.GetProcessSpaceValueAsInt("port"),
		Username: environmentManager.GetProcessSpaceValue("username"),
		Password: environmentManager.GetProcessSpaceValue("password"),
		DBName:   environmentManager.GetProcessSpaceValue("dbname"),
		SSLMode:  environmentManager.GetProcessSpaceValue("sslmode"),
	}

	return config, nil
}



func (cm *ConfigManager) GetDatabaseConnection(serviceName string, cfg common.PostgreSQLConfig) int {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[%s] failed to connect database: %v", serviceName, err)
		return -1
	}
	cm.database.Db = db
	return 0
}

func (cm *ConfigManager) GetDB(serviceName string) *gorm.DB {
	return cm.database.Db
}

