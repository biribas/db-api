package controller

import (
  "fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
  Name string                   `json:"name"`
  CPF int64                     `json:"cpf"`
  DataNascimento time.Time      `json:"data_nascimento"`
}

var data []User

// GetAllUsers godoc
//	@Summary		Show all users
//	@Description	Show all users
//	@Produce		json
//	@Success		200	{array}	User
//	@Router			/ [get]
func GetAllUser (ctx *fiber.Ctx) error {
  ctx.Status(fiber.StatusOK)
  return ctx.JSON(data)
}

// CreateUser godoc
//	@Summary		Create an user
//	@Description	Create an user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		User	true	"User model to be created"
//	@Success		200		{object}	User
//	@Failure		400
//	@Router			/ [post]
func CreateUser (ctx *fiber.Ctx) error {
  var newUser User
  err := ctx.BodyParser(newUser)
  if err != nil {
    fmt.Printf("unable to get user from requested body: %v\n", err)
    return ctx.SendStatus(400)
  }
  data = append(data, newUser)
  ctx.Status(fiber.StatusCreated)
  return ctx.JSON(newUser)
}

