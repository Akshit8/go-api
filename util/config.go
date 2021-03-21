/*
 * @File: util.config.go
 * @Description: impls function to set service config
 * @LastModifiedTime: 2021-01-16 08:36:43
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package util impls service utilities
package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are ready by viper from a config file or environment variable
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuratoion from file or environment variables.
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		// remember
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)

	return config, err
}

// setPostgresHost sets host addr for postgresdb
// func setPostgresHost(config *Config) string {
// 	if os.Getenv("GITHUB_WORKFLOW") == "CI" {
// 		return config.DBSourceTesting
// 	}
// 	return config.DBSourceTesting

// }
