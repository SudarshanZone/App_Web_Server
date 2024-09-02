package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/SudarshanZone/Fno_Ord_Dtls/common"
	"github.com/SudarshanZone/Fno_Ord_Dtls/internal/models"
)
/**********************************************************************************/
/*                                                                                 */
/*  Description       : This package defines a configuration structure for         */
/*                      connecting to a PostgreSQL database. The `PostgreSQLConfig`*/
/*                      struct holds essential details required for establishing   */
/*                      a connection, including username, password, database name, */
/*                      host, port, and SSL mode.                                  */
/*                                                                                 */
/*                      The `GetDatabaseConnection` function formats these values  */
/*                      into a Data Source Name (DSN) string, which is used to     */
/*                      open a connection to the PostgreSQL database using GORM.   */
/*                                                                                 */
/*  Functions         :                                                            */
/*                        - LoadPostgreSQLConfig: Reads the configuration from     */
/*                          an INI file and returns a `PostgreSQLConfig` instance. */
/*                                                                                 */
/*                        - GetDatabaseConnection: Constructs a DSN string from    */
/*                          the provided `PostgreSQLConfig` struct, establishes a  */
/*                          connection to the PostgreSQL database using GORM, and  */
/*                          returns an int status code.                            */
/*                                                                                 */
/*                        - GetDB: Returns the currently established database      */
/*                          connection as a `*gorm.DB` instance.                   */
/*                                                                                 */
/**********************************************************************************/


// ConfigManager implements the DBConfigLoader, DBConnector, and DBAccessor interfaces.
type ConfigManager struct {
	database struct {
		Db *gorm.DB
	}
}

// LoadPostgreSQLConfig reads the configuration from an INI file and returns a PostgreSQLConfig instance.
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
