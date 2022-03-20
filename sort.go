// https://gosamples.dev/generics-sort-slice/

package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func main() {
	floatSlice := []float64{2.3, 1.2, 0.2, 51.2}
	sortSlice(floatSlice)
	fmt.Println(floatSlice)

	stringSlice := []string{"z", "a", "b"}
	sortSlice(stringSlice)
	fmt.Println(stringSlice)

	intSlice := []int{0, 3, 2, 1, 6}
	sortSlice(intSlice)
	fmt.Println(intSlice)
}
