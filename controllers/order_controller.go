package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"w8s/models"
	"w8s/repository"
	"w8s/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	OrderSvc services.OrderSvc
}

var IOrderSvc services.IOrderSvc

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param order body models.Order true "Order object"
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders [post]
func (ic *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	IOrderSvc = &ic.OrderSvc
	err = IOrderSvc.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders [get]
func (ic *OrderController) GetOrders(ctx *gin.Context) {
	var orders []models.Order

	IOrderSvc = &ic.OrderSvc
	err := IOrderSvc.GetOrders(&orders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
			"count":  0,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
		"count":  len(orders),
	})
}

// GetOrder godoc
// @Summary Get a order
// @Description Get a order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [get]
func (ic *OrderController) GetOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	IOrderSvc = &ic.OrderSvc
	err = IOrderSvc.GetOrder(i, &order)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"result": gin.H{},
			"count":  0,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": order,
			"count":  1,
		})
	}
}

// UpdateOrder godoc
// @Summary Update a order
// @Description Update a order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Order object"
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [put]
func (ic *OrderController) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	IOrderSvc = &ic.OrderSvc
	err = IOrderSvc.GetOrder(i, &order)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = ic.OrderSvc.UpdateOrder(i, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// DeleteOrder godoc
// @Summary Delete a order
// @Description Delete a order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [delete]
func (ic *OrderController) DeleteOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	IOrderSvc = &ic.OrderSvc
	err = IOrderSvc.GetOrder(i, &order)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = ic.OrderSvc.DeleteOrder(i)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}

// GetOrderWithPerson godoc
// @Summary Get a order with person
// @Description Get a order with person
// @Tags Order
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/person/{id} [get]
func (ic *OrderController) GetOrderAndPerson(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	IOrderSvc = &ic.OrderSvc
	err = IOrderSvc.GetOrder(i, &order)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"result": gin.H{},
			"count":  0,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	client := &http.Client{}

	req, errReq := repository.GetPersonsFromHTTP("/data.php?qty=1&apikey=7f8fc96e-de1f-4aab-9c62-3dd1de365e66")
	if errReq != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errReq,
		})
	}

	res, errRes := client.Do(req)
	if errRes != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errRes,
		})
		return
	}

	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errBody,
		})
		return
	}

	var personResp *models.PersonResultHTTP
	err = json.Unmarshal(body, &personResp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order":  order,
		"person": personResp.Result[0],
	})
}
