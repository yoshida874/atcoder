package main
 
import "fmt"
 
func main(){
	var n, k, a int
	fmt.Scan(&n, &k, &a)
	x := ((k % n) + (a-1)) % n
	if x == 0{
		fmt.Println(n)
	}else{
		fmt.Println(x)
	}
}