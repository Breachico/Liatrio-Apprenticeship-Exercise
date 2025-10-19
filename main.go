package main

import (
  //"os" // Prempting for later for Docker 
  "time"
  "log"
  "github.com/gofiber/fiber/v2"
)

// Ref for JSON in GO: https://docs.gofiber.io/api/ctx#json
func main() {
  // Initialize a new Fiber app
  app := fiber.New()
  
  // Define a route for the GET method on the root path '/'
  app.Get("/", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "message": "My name is Bryce Emery",
      "timestamp": time.Now().Unix(),
    })
  })

  // add error handling

  //log.Fatal(app.Listen(":8080"))

  
  // Start the server on port 8080
  log.Fatal(app.Listen("0.0.0.0:80"))
  
}
