// Package slice is useful for getting slices.
package slice

const testVersion = 1

// All returns all substrings of length n of a given string.
func All(n int, s string) []string {
	var lenght int = len(s) - n + 1
	results := make([]string, lenght)
	for i := 0; i < lenght; i++ {
		results[i] = s[i:(i + n)]
	}
	return results
}

// UnsafeFirst return first substring of length n from s.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
