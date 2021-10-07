package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"path/handler"
)

func main() {

	//create instance of go fiber
	app := fiber.New()

	//load the static file 
	app.Static("/", "./index.html")

	//print on the terminal
	fmt.Println("go upload files  :  ")

	//call the function
	app.Post("/upload", handler.UplaodData)


	//listen on this port no
	app.Listen(":8000")

}
