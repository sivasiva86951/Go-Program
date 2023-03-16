package main
import (
	"fmt"
)

func main () {

	//if stmt
	//Example 1
	g := 10
	if g%2 == 0 {
		fmt.Println(g, "is even")
	} else {
		fmt.Println(g, " is odd")
	}

	//Example2
	p := 50
	q := 15
	if p < q {
		fmt.Printf("%d is smaller than %d\n", p, q)
		fmt.Printf("%d is larger than %d\n", q, p)
	} else if p > q {
		fmt.Printf("%d is smaller than %d\n", q, p)
		fmt.Printf("%d is larger than %d\n", p, q)
	} else {
		fmt.Println("The two numbers are equal")
	}
}