/*
 * @File: api.validator.go
 * @Description: impls custom validators for the api
 * @LastModifiedTime: 2021-01-16 14:22:35
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package api impls REST api
package api

import (
	"github.com/Akshit8/go-api/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
