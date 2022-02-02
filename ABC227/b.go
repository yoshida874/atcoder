package main
 
import "fmt"
 
func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
 
	res := 0
 
	for i := 0; i < len(s); i++ {
		for j := 1; j <= 1000; j++ {
			b := false
			for k := 1; k <= 1000; k++ {
				ss := (4 * j * k) + (3 * j) + (3 * k)
				if ss == s[i] {
					b = true
					break
				}
			}
			if b {
				res++
				break
			}
		}
	}
	fmt.Println(len(s) - res)
}