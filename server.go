package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "./ui/dist")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/map", handleMap)

	log.Fatal(app.Listen(":3000"))
}

func handleMap(c *fiber.Ctx) error {
	m := GetMap()
	PrintJson(m)
	return c.JSON(m)
}

func PrintJson(v any) {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("bind json failed")
	}
	fmt.Println(string(result))
}
