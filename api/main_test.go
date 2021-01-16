/*
 * @File: api.main_test.go
 * @Description: entry point all api tests
 * @LastModifiedTime: 2021-01-16 12:49:47
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package api impls REST api
package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
