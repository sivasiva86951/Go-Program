package main
import (
    "fmt"
)

func main() {

	//execute fun in sequence
	hi("Siva")
    hi("Ammu")
    hi("Theja")
    hi("Vicky")
    hi("Jashwin")
	fmt.Println("------------------")

	//execute fun concurrently
    go hello("Martin")
    go hello("Lucia")
    go hello("Michal")
    go hello("Jozef")
    go hello("Peter")
	fmt.Println("------------------")

    fmt.Scanln()
}

func hi(name string) {
    fmt.Printf("Hi %s!\n", name)	
}
 
func hello(name string) {
    fmt.Printf("Hello %s!\n", name)
}