/*
 * @File: util.config.go
 * @Description: impls function to set service config
 * @LastModifiedTime: 2021-01-16 08:36:43
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package util impls service utilities
package util

import "os"

// SetPostgresHost sets host addr for postgresdb
func SetPostgresHost() string {
	if os.Getenv("GITHUB_WORKFLOW") == "CI" {
		return "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	}
	return "postgres://root:secret@host.docker.internal:5432/simple_bank?sslmode=disable"

}
