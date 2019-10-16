package main

import (
	"fmt"
	"strings"
)

// Index function
func Index(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

// Include function
func Include(haystack []string, needle string) bool {
	return Index(haystack, needle) >= 0
}

// Any function
func Any(haystack []string, f func(string) bool) bool {
	for _, v := range haystack {
		if f(v) {
			return true
		}
	}
	return false
}

// All function
func All(haystack []string, f func(string) bool) bool {
	for _, v := range haystack {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter function
func Filter(haystack []string, f func(string) bool) []string {
	results := make([]string, 0)
	for _, v := range haystack {
		if f(v) {
			results = append(results, v)
		}
	}
	return results
}

// Map function
func Map(haystack []string, f func(string) string) []string {
	results := make([]string, len(haystack))
	for i, v := range haystack {
		results[i] = f(v)
	}
	return results
}

func main() {
	// var strs = []string{"peach", "apple", "pear", "plum"}
	strs := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "apple"))
	fmt.Println(Include(strs, "banana"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, func(v string) string {
		return v + "!"
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}
