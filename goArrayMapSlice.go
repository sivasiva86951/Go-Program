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

	//slice
	//Example 1
	monthOfYear := [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	months := monthOfYear[:]
	months[1] = "February"
	months[10] = "November"
	fmt.Println(monthOfYear[:])
	fmt.Println("len:", len(monthOfYear))
	fmt.Println("get:", months[2])
	fmt.Println(monthOfYear[2:7])
	fmt.Println(monthOfYear[5:11])

	//Map
	wheelsMap := map[string]int{
		"cycle":  2,
		"riksha": 3,
		"van":    6,
	}
	fmt.Println(wheelsMap)
	wheelsMap["bus"] = 8
	delete(wheelsMap, "van")
	for vehicle, wheels := range wheelsMap {
		fmt.Printf("%s:%d ", vehicle, wheels)
	}
}
