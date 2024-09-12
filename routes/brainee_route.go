package routes

import (
	"TechnicalTestITONICS/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/brainees", controllers.CreateBrainee)
	app.Put("/brainees/:braineeId", controllers.UpdateBrainee)
	app.Delete("/brainees/:braineeId", controllers.DeleteBrainee)
	app.Get("/brainees/:braineeId", controllers.GetBraineeById)
	app.Get("/brainees_author/:braineesAuthor", controllers.GetBraineesByAuthor)
	app.Get("/brainees_brand/:braineesBrand", controllers.GetBraineesByBrand)
	app.Get("brainees", controllers.GetAllBrainees)
}
