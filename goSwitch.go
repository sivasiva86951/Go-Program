package main
import (
	"fmt"
)

func main () {

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