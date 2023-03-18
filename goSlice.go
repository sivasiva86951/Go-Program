package main
import (
	"fmt"
)

func main () {

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
}