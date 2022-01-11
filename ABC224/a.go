package main

import (
	"fmt"
)

func main() {
	var s string
	
	fmt.Scanf("%s", &s)

	if s[len(s)-2:] == "er" {
		fmt.Print("er")
	} else {
		fmt.Print("ist")
	}
}