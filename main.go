package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("I am a main function")

	app := fiber.New()

	MyHelperFunction()

	//Routes

	app.Listen("localhost:9000")
}
