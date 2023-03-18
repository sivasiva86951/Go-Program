package main
import "fmt"

func main() {
	
	//channels
	numCh := make(chan int)

	//GoRoutines
	go counter(numCh)
	for a := range numCh {
		fmt.Println(a)
	}
}

func counter(out chan int) {
	for i := 0; i < 10; i++{
		out <- i
	}
	close(out)
}



