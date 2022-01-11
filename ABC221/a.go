package main

import "fmt"

func main() {
	var A, B, i, k int
	k = 1

	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	for i = B; i < A; i++ {
		k *= 32
	}

	fmt.Print(k)
}
