package main
import "fmt"

func main () {

	//one dim array
	var a [5]int
	fmt.Println("emp:", a)

	a[2] = 50
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{21, 32, 43, 54, 65}
	fmt.Println("dcl:", b)

	//two dim array
	var twoD [3][3]int
	for k := 0; k < 2; k++ {
		for l := 0; l < 3; l++ {
			twoD[k][l] = k + l
		}
	}
	fmt.Println("2d: ", twoD)
}