package main

import "fmt"
 
func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}
 
	counter := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (x[j] - x[i])*(y[k] - y[i]) - (x[k] - x[i])*(y[j] - y[i]) != 0 {
					counter++
				}
			}
		}
	}
	fmt.Printf("%d\n", counter)
}