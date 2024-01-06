package controllers

// import (
// 	"math/rand"
// 	"strings"
// )

// const (
// 	base         uint64 = 62
// 	base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// )

// func Base62Conversion() string {
// 	var shortURLLength = 6
// 	var sb strings.Builder
// 	for i := 0; i < shortURLLength; i++ {
// 		index := rand.Intn(len(base62Digits))
// 		sb.WriteByte(base62Digits[index])
// 	}

// 	return sb.String()
// }

// func toBase62(num uint64) string {
// 	encoded := ""
// 	for num > 0 {
// 		r := num % base
// 		num /= base
// 		encoded = string(base62Digits[r]) + encoded

// 	}
// 	return encoded
// }

