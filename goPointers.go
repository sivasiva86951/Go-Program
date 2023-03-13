package main
import "fmt"

func main(){

	//pointers
	a,b := 9, 27
	fmt.Println("Before swapping :" ,a,b)
	swap(&a, &b)
	fmt.Printf("After swapping : %d %d" ,a, b)
}

func swap(a,b *int){
	*a, *b = *b, *a
}