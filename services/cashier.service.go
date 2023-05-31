package Services

import (
	"strconv"
	"time"
"fmt"
	"github.com/gofiber/fiber/v2"
	config "github.com/juanfmange/fiberapp/Config"
	models "github.com/juanfmange/fiberapp/Models"
)

func CreateCashier(c *fiber.Ctx) error {
	// Implementa la lógica de creación del cajero aquí
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "bad request",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name is requiered",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier passcode is requiered",
			})
	}

	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	config.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "cashier created",
		"data":    cashier,
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	// Implementa la lógica para obtener los detalles del cajero aquí
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	config.DB.Select("id, name").Where("id=?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["created_at"] = cashier.CreatedAt
	cashierData["updated_at"] = cashier.UpdatedAt

	//Si no hay un cashier o no se dio un id
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "cashier not found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "cashier details",
		"cashier": cashierData,
	})
}

func CashierList(c *fiber.Ctx) error {
	// Implementa la lógica para obtener la lista de cajeros aquí
	var cashier []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	config.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"Success": true,
			"message": "Cashier List Data",
			"data":    cashier,
		})
}

func UpdateCashier(c *fiber.Ctx) error {
	// Implementa la lógica para actualizar el cajero aquí
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	config.DB.Find(&cashier, "id+?", cashierId)

	//validacion para checar el cashier id

	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	cashier.Name = updateCashier.Name
	config.DB.Save(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier updated",
	})

}

func DeleteCashier(c *fiber.Ctx) error {
	// Implementa la lógica para eliminar el cajero aquí
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
ar cashier models.Cashier
// Implementa la lógica para eliminar el cajero aqu
ashierId := c.Params("cashierId")
var cashier models.Caier

	config.DB.Where("id", cashierId).Fist(&cashier)

	if shier.Id == 0 {
		urn c.Status(404).JSON(fiber.Map{
"success": false,
		"message": "Cashier not found",
		})
	}

	coig.DB.Where("id=?", cashierId).Delete(cashier)
	urn c.Status(200).JSON(fiber.Map{
	"success": true,
		"message": "Cashier deleted succesfully",
	})
}
