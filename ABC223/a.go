package main

import "fmt"

func main() {
	var x int

	fmt.Scan(x)
	fmt.Scanf("%d", &x)

	if x != 0 && x % 100 == 0 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}