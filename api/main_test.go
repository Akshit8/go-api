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
	"time"

	db "github.com/Akshit8/go-api/db/sqlc"
	"github.com/Akshit8/go-api/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: 5*time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
