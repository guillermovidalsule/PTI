package routes

import (
	controller "restful-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	//Route para signup
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.OPTIONS("/users/signup")

	//Route para login
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.OPTIONS("/users/login")

	incomingRoutes.POST("/users/existe", controller.Existe())
	incomingRoutes.OPTIONS("/users/existe")
	//Route para test de cpu_freq (tarde o temprano lo quitaremos)
	incomingRoutes.POST("/robot/info", controller.AddParam())
	incomingRoutes.OPTIONS("/robot/info")

}

func UserManagment(incomingRoutes *gin.Engine) {
	//Route para dar de baja un user
	incomingRoutes.POST("/users/baja", controller.DeleteUser())
	incomingRoutes.OPTIONS("/users/baja")

}

func RobotRoutes(incomingRoutes *gin.Engine) {

	//Route para añadir robot
	incomingRoutes.POST("/robot/alta", controller.CreateRobotForUser())
	incomingRoutes.OPTIONS("/robot/alta")

	//Route para eliminar robot
	incomingRoutes.POST("/robot/baja/:robotname", controller.DeleteRobotForUser())
	incomingRoutes.OPTIONS("/robot/baja/:robotname")

	//Listar todos los robots de un user
	incomingRoutes.GET("/robot/listar", controller.ListUserRobots())
	incomingRoutes.OPTIONS("/robot/listar")

	//Consultar UN robot
	incomingRoutes.GET("/robot/consulta/:robotname", controller.ConsultUserRobot())
	incomingRoutes.OPTIONS("/robot/consulta/:robotname")

	//Establecer macros robot
	incomingRoutes.POST("robot/refresh_macros/:robotname", controller.RefreshMacrosRobot())
	incomingRoutes.OPTIONS("robot/refresh_macros/:robotname")

}
