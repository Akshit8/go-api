/*
 * @File: util.random.go
 * @Description: generates random data
 * @LastModifiedTime: 2021-01-14 17:33:06
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package util impls service utilities
package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency type
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "INR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
