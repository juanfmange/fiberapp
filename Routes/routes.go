package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanfmange/fiberapp/Controllers"
)

func Setup(app *fiber.App) {
	// app.Post("/cashiers/:cashierId/login", controller.Login)
	// app.Get("/cashiers/:cashierId/logout",controller.Logout)
	// app.Post("/cashiers/cashierId/passcode",controller.Passcode)

	// Cashier routes
	app.Post("/cashiers", Controllers.CreateCashier)
	app.Get("/cashiers/:cashierId", Controllers.GetCashierDetails)
	app.Get("/cashiers", Controllers.CashierList)
	app.Delete("/cashiers/:cashierId", Controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", Controllers.UpdateCashier)
}
