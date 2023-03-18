package main
import "fmt"

func main () {
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