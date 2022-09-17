package routers

import (
	"net/http"
	"os"
	"w8s/controllers"
	"w8s/database"
	"w8s/repository"
	"w8s/services"

	_ "w8s/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	isValid := hasAuth && user == os.Getenv("BASIC_AUTH_USER") && password == os.Getenv("BASIC_AUTH_PASS")
	if !isValid {
		c.Abort()
		result := gin.H{
			"result": "unauthorized access",
		}
		c.JSON(http.StatusUnauthorized, result)
	}
	c.Next()
}

// @securityDefinitions.basic BasicAuth
func StartServer() *gin.Engine {
	db := database.GetDB()
	// Initialize repositories
	personRepo := repository.PersonRepo{Conn: db}
	orderRepo := repository.OrderRepo{Conn: db}

	// Initialize services
	personSvc := services.PersonSvc{PersonRepo: personRepo}
	orderSvc := services.OrderSvc{OrderRepo: orderRepo}

	// Initialize controllers
	personCtrl := controllers.PersonController{PersonSvc: personSvc}
	orderCtrl := controllers.OrderController{OrderSvc: orderSvc}

	router := gin.Default()

	// Person routes
	router.GET("/persons", personCtrl.GetPersons)
	router.GET("/person/:id", personCtrl.GetPerson)
	router.POST("/person", personCtrl.CreatePerson)
	router.PUT("/person/:id", personCtrl.UpdatePerson)
	router.DELETE("/person/:id", personCtrl.DeletePerson)

	// Order routes
	router.POST("/orders", orderCtrl.CreateOrder)
	router.GET("/orders", orderCtrl.GetOrders)
	router.GET("/order/:id", orderCtrl.GetOrder)
	router.PUT("/order/:id", orderCtrl.UpdateOrder)
	router.DELETE("/order/:id", orderCtrl.DeleteOrder)

	router.GET("/orders/person/:id", BasicAuth, orderCtrl.GetOrderAndPerson)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
