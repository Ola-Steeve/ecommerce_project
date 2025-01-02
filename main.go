package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "go-ecommerce-app/config"
)

func main() {

	app := fiber.New()

	//Basic Types: int, float64, string, boolean
	//composite types: array, slice, map, struct
	//Pointer Types: *

	//var age int
	//var height float64
	//var firstName string
	//var isEmployed bool

	age := 29
	height := 179.87
	firstName := "Ola"
	isEmployed := true

	//fmt.Println(age, height, firstName, isEmployed)
	fmt.Printf("Age: %d\n", age) //for s
	fmt.Printf("Height: %f\n", height)
	fmt.Printf("First Name: %s\n", firstName)
	fmt.Printf("Is Employed: %t\n", isEmployed)

	app.Listen("localhost:9000")
}
