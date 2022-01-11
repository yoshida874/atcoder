package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var S, T string
	ans := "No"

	fmt.Scanf("%s\n", &S)
	fmt.Scanf("%s", &T)

	if S == T {
		ans = "Yes"
	}

	sliceS := strings.Split(S, "")
	sliceT := strings.Split(T, "")
	len := len(sliceS)

	for i := 0; i < len-1; i++ {
		sliceS[i], sliceS[i+1] = sliceS[i+1], sliceS[i]
		if reflect.DeepEqual(sliceS, sliceT) {
			ans = "Yes"
			break
		}
		sliceS[i], sliceS[i+1] = sliceS[i+1], sliceS[i]
	}
	fmt.Print(ans)
}
