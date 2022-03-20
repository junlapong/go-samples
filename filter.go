// https://gosamples.dev/generics-filter/

package main

import (
	"fmt"
	"strings"
)

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func main() {
	websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
	httpsWebsites := filter(websites, func(v string) bool {
		return strings.HasPrefix(v, "https://")
	})
	fmt.Println(httpsWebsites)

	numbers := []int{1, 2, 3, 4, 5, 6}
	divisibleBy2 := filter(numbers, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(divisibleBy2)
}
