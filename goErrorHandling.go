package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
    
	//ioutil.Readfile example
    content, err := ioutil.ReadFile("file.txt")
     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))

	//panic example- runtime error
	var x int = 10
    var y int = 0
    fmt.Println(x / y)

	//errorrs.new
	error := errors.New("Some error")
    if err != nil {
        fmt.Println(error)
    }
}