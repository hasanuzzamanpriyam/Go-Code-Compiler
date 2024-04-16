package main

import (
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Serve to the index file
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	// Endpoint for submitting code
	app.Post("/compile", func(c *fiber.Ctx) error {
		// Get the code from the request body
		code := c.Body()

		// compile and  run the code
		output, err := CompileAndRun(code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Return the output
		return c.SendString(output)
	})
	app.Listen(":8080")
}

func CompileAndRun(code []byte) (string, error) {
	// Write code to a temporary file
	tempFile := "temp.go"
	err := writeFile(tempFile, code)
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile)

	// Compile the code
	cmd := exec.Command("go", "run", tempFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

func writeFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
