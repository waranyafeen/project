package routers

import(
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/project/backend/controllers"
	"github.com/waranyafeen/project/backend/middlewares"
)

func SetupRouter() *gin.Engine {
	return gin.Default()
}

func InitRouter(route *gin.Engine) {
	route.Use(middlewares.Authorizes())

	authRouter := route.Group("/")
	initRequiredAuthRouter(authRouter)

}

func initRequiredAuthRouter(route *gin.RouterGroup) {
	route.Use(middlewares. Authorizes())

	user := middlewares. Authorizes()

	// employee system
	route.GET("/employees", controllers.GetAllEmployees)
	route.GET("/employees/:id", controllers.GetEmployees)
	route.POST("/employees", controllers.CreateEmployees)
	route.PUT("/employees/:id", controllers.UpdateEmployees)
	route.DELETE("/employees/:id", controllers.DeleteEmployees)

	route.GET("/employees/position", user, controllers.GetAllPositions)
	route.GET("/employees/precede", controllers.GetAllPrecedes)
	route.GET("/employees/gender", controllers.GetAllGendersForEmployee)

	// payment system
	route.GET("/payments", controllers.GetAllPayments)
	route.GET("/payments/:id", controllers.GetPayments)
	route.POST("/payments", controllers.CreatePayments)
	route.PUT("/payment/:id", controllers.UpdatePayments)
	route.DELETE("/payment/:id", controllers.DeletePayments)

	// ticket system
	route.GET("/tickets", controllers.GetAllTickets)
	route.GET("/tickets/:id", controllers.GetTickets)
	route.POST("/tickets", controllers.CreateTickets)
	route.PUT("/tickets/:id", controllers.UpdateTickets)
	route.DELETE("/tickets/:id", controllers.DeleteTickets)

	route.GET("/tickets/car", controllers.GetAllCars)
	route.GET("/tickets/province", controllers.GetAllProvinces)
	route.GET("/tickets/departure", controllers.GetAllDepartures)

	// user system
	route.GET("/users", controllers.GetAllUsers)
	route.GET("/users/:id", controllers.GetUsers)
	route.POST("/users", controllers.CreateUsers)
	route.PUT("/users/:id", controllers.UpdateUsers)
	route.DELETE("/users/:id", controllers.DeleteUsers)

	route.GET("/users/role", controllers.GetAllRoles)
	route.GET("/users/gender", controllers.GetAllGendersForUser)
}
