package main
import (
    "fmt"
    "time"
)

func main() {

	//channels
    done := make(chan bool, 1)

	//goroutine
    go worker(done)

    <-done
}

func worker(done chan bool) {
    fmt.Print("Work..")
    time.Sleep(2 * time.Second)
    fmt.Println("done")

    done <- true
}