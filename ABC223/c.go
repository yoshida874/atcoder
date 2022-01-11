package main

import (
	"fmt"
)

func main() {
	var n int
	
	fmt.Scanf("%d", &n)
	a := make([]float64, n)
	b := make([]float64, n)

	for i:=0; i < n; i++ {
		fmt.Scanf("%f %f", &a[i], &b[i])
	}

	var x float64 = 0;
	for i:=0; i < n; i++ {
		x += a[i] / b[i];
	}

	x /= 2;
	var ans float64 = 0;

	for i:= 0; i < n; i++ {
		if x >= a[i] / b[i] {
			ans += a[i]
			x -= a[i] / b[i]
		} else {
			ans += b[i] * x
			break
		}
	}

	fmt.Print(ans)
}