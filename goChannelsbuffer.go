package main
import "fmt"

func main() {

    msg := make(chan string, 2)    
    msg <- "Channel"
	msg <- "Buffer"

    fmt.Println(<-msg)
    fmt.Println(<-msg)
}