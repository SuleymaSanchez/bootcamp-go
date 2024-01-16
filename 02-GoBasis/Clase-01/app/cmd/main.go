package main

/*
	primite types
		- string
		- int
		- float64
		- bool

		based types
		- byte

		type inference
		name := "John"    -> string
		married := true   -> bool
		age := 25         -> int
		height := 175.5   -> float64
*/
import (
	"fmt"
	"hola-mundo/app/calculator"

	"github.com/fatih/color"
)

const (
	FuelType = "Pertamax"
)

func main() {
	// println("Hello World!")

	// explicit declaration
	// - var

	// Deja el valor por defecto
	var firstName string
	firstName = "John"
	fmt.Println(firstName)

	var lastName string = "Doe"
	fmt.Println(lastName)

	var hasDrivenLicense bool = true

	// Implicit declaration (typr inference)
	// - var
	// - :=

	var age = 25
	height := 175.5

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Has Driver License:", hasDrivenLicense)

	// use calculator package
	// ...
	fmt.Println("Sum:", calculator.SumResult)
	fmt.Println("Sub:", calculator.SubResult)

	// Use color package
	color.Red("The program has been finished")
}
