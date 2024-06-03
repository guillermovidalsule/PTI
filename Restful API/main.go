package main

import (
	"os"

	//"restful-api/controllers"
	"restful-api/middleware"
	routes "restful-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	/*
		port := os.Getenv("PORT")

		if port == "" {
			port = "8082"
		}

		router := gin.New()
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"POST"},
			AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		}))


		public := router.Group("/users")
		//Route para signup
		public.POST("/signup", controllers.SignUp())
		public.OPTIONS("/signup")

		//Route para login
		public.POST("/login", controllers.Login())
		public.OPTIONS("/login")



		private := router.Group("/robot")
		private.Use(middleware.Authentication())
		//Route para añadir robot
		private.POST("/alta", controllers.CreateRobotForUser())
		private.OPTIONS("/alta")

		//Route para eliminar robot
		private.POST("/baja/:robotname", controllers.DeleteRobotForUser())
		private.OPTIONS("/baja/:robotname")

		//Listar todos los robots de un user
		private.GET("/listar", controllers.ListUserRobots())
		private.OPTIONS("/listar")

		router.Run(":" + port)
		**/

	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	/* --------------- Public routes --------------- */

	public := router

	//Rutas de signin/login de user
	routes.AuthRoutes(public)

	/* --------------- Private routes --------------- */

	private := router

	//Control de Middleware de rutas privadas
	private.Use(middleware.Authentication())

	//Rutas de gestion de robots de un user
	routes.RobotRoutes(private)

	//Rutas de gestion de user
	routes.UserManagment(private)

	router.Run(":" + port)

}
