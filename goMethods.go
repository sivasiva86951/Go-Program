package main
import "fmt"

// Celsius is temp unit
type Celsius float32

// Fahrenheit is temp unit
type Fahrenheit float32

//Method converts cesius to fahrenheit
func (c Celsius) ConvertToFahrenheit() Fahrenheit{
	return Fahrenheit(9.0/5.0*c +32)
}

func (f Fahrenheit) String() string{
    return fmt.Sprintf("%.2f%cF", f,  0x00B0)
}

func main(){
	var c Celsius
	c = 35.8
	f := c.ConvertToFahrenheit()
	fmt.Println(f)
}
