package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"w8s/models"
	"w8s/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	OrderSvc services.OrderSvc
}

func (ic *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = ic.OrderSvc.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (ic *OrderController) GetOrders(ctx *gin.Context) {
	var orders []models.Order

	err := ic.OrderSvc.GetOrders(&orders)
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

	err = ic.OrderSvc.GetOrder(i, &order)
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

	err = ic.OrderSvc.GetOrder(i, &order)
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

	err = ic.OrderSvc.GetOrder(i, &order)
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
