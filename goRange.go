package main
import "fmt"

func main(){

	// range using array
	arr:=[5]int {11,22,33,44,55}
	for i,n:=range arr{ 					
	fmt.Printf("arr[%d]:%d\n",i,n)
	}

	// range using slice
	slc:=[]int{927,869,221}
	for i1,n1:=range slc{ 					
	fmt.Printf("slice[%d]:%d\n",i1,n1)
	}

	// range using string
	str:="Goprogramming"
	for i,n:=range str{ 					
	fmt.Printf("string[%d]:%d\n",i,n)
	}
}