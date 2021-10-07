package handler

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UplaodData(c *fiber.Ctx) error {
	fmt.Println("file uploading start")

	//MultipartForm files can be retrieved by name,
	//the first file from the given key is returned.
	// Get first file from form field "myfile":

	file, err := c.FormFile("myFile")

	// Check for errors:
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	// Save file to root directory
	c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	//create the temp named file and upload the data
	tempFile, err := ioutil.TempFile("temp", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}

	//close temp folder at the last of executuion
	defer tempFile.Close()

	fmt.Println("Successfully Uploaded File\n")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "completed data",
	})
}
