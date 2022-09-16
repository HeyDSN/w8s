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

type PersonController struct {
	PersonSvc services.PersonSvc
}

func (ic *PersonController) GetPersons(ctx *gin.Context) {
	var persons []models.Person

	err := ic.PersonSvc.GetPersons(&persons)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
			"count":  0,
		})
	}
	if len(persons) <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"count":  0,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	}
}

func (ic *PersonController) GetPerson(ctx *gin.Context) {
	var person *models.Person
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = ic.PersonSvc.GetPerson(i, person)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
			"count":  0,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": person,
			"count":  1,
		})
	}
}

func (ic *PersonController) CreatePerson(ctx *gin.Context) {
	var person *models.Person

	first_name := ctx.PostForm("first_name")
	last_name := ctx.PostForm("last_name")
	person.FirstName = first_name
	person.LastName = last_name

	err := ic.PersonSvc.CreatePerson(person)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"result": person,
		})
	}
}

func (ic *PersonController) UpdatePerson(ctx *gin.Context) {
	id := ctx.Param("id")
	first_name := ctx.PostForm("first_name")
	last_name := ctx.PostForm("last_name")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var person *models.Person
	person.FirstName = first_name
	person.LastName = last_name

	err = ic.PersonSvc.UpdatePerson(i, person)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
			"count":  0,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": person,
			"count":  1,
		})
	}
}

func (ic *PersonController) DeletePerson(ctx *gin.Context) {
	var person models.Person
	id := ctx.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = ic.PersonSvc.DeletePerson(i)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
			"count":  0,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": person,
			"count":  1,
		})
	}
}
