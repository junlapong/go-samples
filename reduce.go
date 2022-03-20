// https://gosamples.dev/generics-reduce/

package main

import (
    "fmt"
)

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    sum := reduce(numbers, func(acc, current int) int {
        return acc + current
    }, 0)
    fmt.Println(sum)
	
    divided := reduce(numbers, func(acc float64, current int) float64 {
        return acc + float64(current)/10.0
	}, 0)
    fmt.Println(divided)
}

