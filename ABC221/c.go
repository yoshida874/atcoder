package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n string
	_, _ = fmt.Scanln(&n)
	runes := []rune(n)
	var values []int
	for i := 0; i < len(runes); i++ {
		v, _ := strconv.Atoi(string(runes[i]))
		values = append(values, v)
	}
	sort.Ints(values)

	var a, b int
	for i := len(values) - 1; i >= 0; i-- {
		if a < b {
			a, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(values[i]))
		} else {
			b, _ = strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(values[i]))
		}
	}
	fmt.Println(a * b)
}
