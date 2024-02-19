package main

import (
	"github.com/biribas/db-api/controller"
	_ "github.com/biribas/db-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			Databases College Homeowrk API
//	@version		0.1
//	@description	API created for college homework 
func main() {
  app := fiber.New()
  app.Use(logger.New())

  app.Get("/swagger/*", swagger.HandlerDefault)

  app.Get("/", controller.GetAllUser)
  app.Post("/", controller.CreateUser)

  app.Listen(":2814")
}

