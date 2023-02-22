// Package utils - general utility functions
package utils

import (
	"bytes"
	"fmt"
	"strings"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// comma inserts commas in a non-negative decimal integer string
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c/go => c, a/b.c.go => b.c
func Basename(s string) string {
	slash := strings.LastIndex(s, "/") // --1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// intsToString is like fmt.Sprint(values) but adds commas.
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
