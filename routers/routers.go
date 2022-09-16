package routers

import (
	"w8s/controllers"
	"w8s/database"
	"w8s/repository"
	"w8s/services"

	"github.com/gin-gonic/gin"
)

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

	return router
}
