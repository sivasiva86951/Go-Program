package main

import (
	"fmt"
	"math"
)

func main() {

	//print stmt
	fmt.Println("Go programming welcomes you")

	//values
	fmt.Println("go" + "language")
	fmt.Println("191+927 =", 191+927)
	fmt.Println("10.0/3.0 =", 10.0/3.0)

	//boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	//varaiables
	var a = "SIVASUBRAMANIAN"
	fmt.Println(a)

	var b int = 9
	var c int = 27
	fmt.Println(b, c)

	//operators
	fmt.Println("Addition:", b+c)
	fmt.Println("Subtraction:", b-c)
	fmt.Println("Multiply:", b/c)
	fmt.Println("Divide:", b*c)
	fmt.Println("Modulus:", b%c)

	var d = false
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "Kloudone"
	fmt.Println(f)

	//constant
	const s string = "constant"
	fmt.Println(s)

	const n = 60
	fmt.Println(math.Cos(n))

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

	//For loop
	//Example 1
	k := 5
	for i := 0; i < k; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	//switch stmt
	var z int = 3          								
	fmt.Printf("Enter the day of a week \n")
	switch z{
	case 1:
	fmt.Println("Monday")
	case 2:
	fmt.Println("Tuesday")
	case 3:
	fmt.Println("Wednesday")
	case 4:
	fmt.Println("Thursday")
	case 5:
	fmt.Println("Friday")
	case 6:
	fmt.Println("Saturday")
	case 7:
	fmt.Println("Sunday")
	default: 
	fmt.Println("Invalid")
	}
}
