package main
import "fmt"

func vals() (int, int) {
    return 9, 27
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}