package main
import "fmt"


//normal function
func add(a int, b int) int {
	return a + b
}

//omit func - same datatype of parameters
func sub(a, b int) int {
	return a - b
}

//multiple return values
func values() (int, int) {
	return 7, 8
}

//variadic function
func Sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


//recursive function
func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

//closures or anonymous function
func intSequence() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}



func main() {
	
	//normal function
	fmt.Println("Addition:", add(9, 27))

	//omit func - same datatype of parameters
	fmt.Println("Subtraction:", sub(81, 35))

	//multiple return values
	p, q := values()
	fmt.Println(p, q)
	_, r := values()
	fmt.Println(r)

	//variadic function
	Sum(1, 2)
	nums := []int{1, 2, 3, 4}
	Sum(nums...)

	//recursive function
	fmt.Println("Factorial of number 5:", fact(5))

	//closures or anonymous function
	nextInt := intSequence()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSequence()
    fmt.Println(newInts())
}
