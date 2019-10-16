package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("sorted strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("sorted ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("ints sorted?", s)
}
