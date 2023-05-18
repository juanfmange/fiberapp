package Controllers

import (
	// Importa otros paquetes necesarios

	"github.com/gofiber/fiber/v2"
	Services "github.com/juanfmange/fiberapp/Services"
)

func CreateCashier(c *fiber.Ctx) error {
	return Services.CreateCashier(c)
}

func GetCashierDetails(c *fiber.Ctx) error {
	return Services.GetCashierDetails(c)
}

func CashierList(c *fiber.Ctx) error {
	return Services.CashierList(c)
}

func UpdateCashier(c *fiber.Ctx) error {
	return Services.UpdateCashier(c)
}

func DeleteCashier(c *fiber.Ctx) error {
	return Services.DeleteCashier(c)
}
