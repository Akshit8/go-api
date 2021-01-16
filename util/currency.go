/*
 * @File: util.currency.go
 * @Description: impls logic for supported currencies
 * @LastModifiedTime: 2021-01-16 14:26:42
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package util impls service utilities
package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, CAD:
		return true
	}
	return false
}
