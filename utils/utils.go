// Package utils - general utility functions
package utils

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
func comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return comma(s[:n-3]) + "," + s[n-3:]
}
