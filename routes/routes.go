package routes

import (
	"github/users/controllers"
	"github/users/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register )
	app.Post("/api/login", controllers.Login )
	app.Get("/api/users", controllers.AllUsers )

	app.Use(middlewares.IsAuthenticated)

	app.Put("/api/users/info", controllers.UpdateInfo )
	app.Put("/api/users/password", controllers.UpdatePassword )


	app.Get("/api/user", controllers.User )
	app.Post("/api/logout", controllers.Logout )

	// app.Get("/api/users", controllers.AllUsers )
	app.Post("/api/users", controllers.CreateUser )
	app.Get("/api/users/:id", controllers.GetUser )
	app.Put("/api/users/:id", controllers.UpdateUser )
	app.Delete("/api/users/:id", controllers.DeleteUser )

	app.Get("/api/roles", controllers.AllRoles )
	app.Post("/api/roles", controllers.CreateRole )
	app.Get("/api/roles/:id", controllers.GetRole )
	app.Put("/api/roles/:id", controllers.UpdateRole )
	app.Delete("/api/roles/:id", controllers.DeleteRole )

	app.Get("/api/permissions", controllers.AllPermissions )
}
