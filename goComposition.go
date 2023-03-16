package main
import "fmt"

type date struct {
	day   int
	month int
	year  int
}

func (d *date) setdate() {
	fmt.Println("Enter the dob:")
	fmt.Scan(&d.day)
	fmt.Scan(&d.month)
	fmt.Scan(&d.year)
}
func (d date) getdate() {
	fmt.Print("dob:", d.day, "/", d.month, "/", d.year)

}

type employee struct {
	empName   string
	age  	  int
	phno  	  int
	dept      string
	salary    int
	empId     int
	dob       date //composition happens here
}

func (e *employee) setinfo() {
	fmt.Println("Enter the details:")
	fmt.Scan(&e.empName, &e.age, &e.phno, &e.dept, &e.salary, &e.empId)
	e.dob.setdate()

}
func (e employee) getinfo() {
	fmt.Println("name:", e.empName)
	fmt.Println("age:", e.age)
	fmt.Println("ph.no:", e.phno)
	fmt.Println("dept:", e.dept)
	fmt.Println("salary:", e.salary)
	fmt.Println("id:", e.empId)
	e.dob.getdate()
}
func main() {
	emp := employee{}
	emp.setinfo()
	emp.getinfo()
}