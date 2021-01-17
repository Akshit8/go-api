/*
 * @File: util.main_test.go
 * @Description: entrypoint to all utils tests
 * @LastModifiedTime: 2021-01-14 16:47:30
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package util

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
