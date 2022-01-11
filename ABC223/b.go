package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	n := len(s)
	v := make([]string, 0);

	for i:=0; i < n; i++ {
		v = append(v, s)
		s = s[1:] + s[:1]
	}

	sort.Strings(v)
	fmt.Println(v[0])
	fmt.Println(v[n-1])
}