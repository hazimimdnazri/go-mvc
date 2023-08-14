package controllers

import (
	"go-mvc/models"

	"github.com/gofiber/fiber/v2"
)

func ViewLogin(c *fiber.Ctx) error {
	return c.Render("guest/login", fiber.Map{})
}

func SubmitLogin(c *fiber.Ctx) error {
	if c.Is("json") {
		return c.Send(c.Body())
	}

	return c.SendString("Error")
}

func ViewRegister(c *fiber.Ctx) error {
	return c.Render("guest/register", fiber.Map{})
}

func SubmitRegister(c *fiber.Ctx) error {
	p := new(models.User)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	if c.FormValue("password_confirmation") != p.Password {
		return c.Status(200).SendString("The password is not the same.")
	}

	return c.Status(200).JSON(p)

}
