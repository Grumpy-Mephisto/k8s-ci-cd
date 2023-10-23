package config

import (
	"os-container-project/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewMariaDBConfig() *MariaDBConfig {
	return &MariaDBConfig{
		Host:     utils.GetEnv("MARIADB_HOST", "localhost").(string),
		Port:     utils.GetEnv("MARIADB_PORT", "3306").(string),
		Username: utils.GetEnv("MARIADB_USERNAME", "root").(string),
		Password: utils.GetEnv("MARIADB_PASSWORD", "root").(string),
		Database: utils.GetEnv("MARIADB_DATABASE", "os-container-project").(string),
	}
}

func (m *MariaDBConfig) NewMariaDBClient() (*gorm.DB, error) {
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
