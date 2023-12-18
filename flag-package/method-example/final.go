/* This is a program that uses all the concepts taught in this directory.
It can be found at: https://github.com/golangbot/methods/blob/master/methods.go
*/

package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
	age      int
}

/*
displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
displaySalary() method converted to function with Employee as parameter type
*/
func displaySalary(e Employee) {
	fmt.Printf("\nSalary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

// Method Area() on Rectangle value receiver
func (r Rectangle) Area() int {
	return r.length * r.width
}

// Function perimeter with Rectangle pointer argument
func perimeter(r *Rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

// Method perimeter on Rectangle pointer receiver
func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

// Method Area() on Circle value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Function Area() on Rectangle value receiver
func area(r Rectangle) {
	fmt.Printf("\nArea Function result: %d\n", (r.length * r.width))
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("\nFull address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address   //anonymous field
}

type myInt int

func (a myInt) add(b myInt) myInt { //method on non-struct type
	return a + b
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling the displaySalary() method of Employee type

	displaySalary(emp1) //Calling the displaySalary function

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("\nArea of rectangle %d\n", r.Area()) //Calling Area method of Rectangle receiver

	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Printf("\nEmployee name before change: %s", e.name)
	e.changeName("Michael Andrew") //Calling method with value receiver
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\nEmployee age before change: %d", e.age)
	e.changeAge(51) //Calling method with pointer receiver
	fmt.Printf("\nEmployee age after change: %d", e.age)

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress() //Accessing fullAddress mehod of anonymous field address

	rect1 := Rectangle{
		length: 10,
		width:  5,
	}

	area(rect1)
	fmt.Println("Area Method Result:", rect1.Area())

	ptr1 := &rect1
	/*
		Compilation error, cannot use ptr1 (type *Rectangle) as type Rectangle in argument to area
	*/
	//area(ptr1)
	fmt.Println("Area Method Result:", ptr1.Area()) //Calling value receiver with a pointer

	rect2 := Rectangle{
		length: 10,
		width:  5,
	}

	ptr2 := &r //pointer to r
	perimeter(ptr2)
	ptr2.perimeter()

	/*
		Compilation error, cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	perimeter(r)

	rect2.perimeter() //Calling pointer receiver with a value

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2) //Calling method on non struct type
	fmt.Println("Sum is", sum)
}
