package main

import "fmt"

func main() {
	
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
	s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s,"d","e","f")   
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

}
