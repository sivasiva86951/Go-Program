package main
import (
	"fmt"
)

func main () {

	//For loop
	//Example 1
	k := 5
	for i := 0; i < k; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}