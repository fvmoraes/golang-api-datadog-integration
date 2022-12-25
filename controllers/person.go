package controllers

import (
	"api-sample/database"
	"api-sample/logs"
	"api-sample/models"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func getCarrierContext(ctx *gin.Context) (data string, span ddtrace.Span) {

	dataReq := ctx.Request.RemoteAddr + " " + ctx.Request.RequestURI
	spanCtx, _ := tracer.SpanFromContext(ctx.Request.Context())
	return dataReq, spanCtx
}

func ShownAllPeople(c *gin.Context) {
	db := database.GetDatabase()
	var people []models.Person
	err := db.Find(&people).Error
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Cannot find all people: " + err.Error(),
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "ERROR", "", err)
		return
	}
	c.JSON(200, people)
	data, span := getCarrierContext(c)
	logs.PopulateErrorLogFile(span, data, "INFO", "showing all person records", nil)
}

func ShownPerson(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Warning": "ID has to be number.",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "id has to be number", err)
		return
	}
	db := database.GetDatabase()
	var person models.Person
	err = db.First(&person, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Warning": "Cannot find person: " + err.Error(),
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "", err)
		return
	}
	c.JSON(200, person)
	data, span := getCarrierContext(c)
	logs.PopulateErrorLogFile(span, data, "INFO", "showing person records", nil)
}

func CreatePerson(c *gin.Context) {
	db := database.GetDatabase()
	var person models.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"Warning": "Cannot bind person to JSON.",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "cannot bind person to json", err)
		return
	}
	err = db.Create(&person).Error
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Cannot create person: " + err.Error(),
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "ERROR", "", err)
		return
	}
	c.JSON(200, person)
	data, span := getCarrierContext(c)
	logs.PopulateErrorLogFile(span, data, "INFO", "created person record", nil)
}

func UpdatePerson(c *gin.Context) {
	db := database.GetDatabase()
	var person models.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"Warning": "Cannot bind person to JSON.",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "cannot bind person to json", err)
		return
	}
	err = db.Save(&person).Error
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Cannot update person: " + err.Error(),
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "ERROR", "", err)
		return
	}
	c.JSON(200, person)
	data, span := getCarrierContext(c)
	logs.PopulateErrorLogFile(span, data, "INFO", "updated person record", nil)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Warning": "ID has to be number.",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "id has to be number", err)
		return
	}
	db := database.GetDatabase()
	var person models.Person
	err = db.Delete(&person, newid).Error
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Cannot delete person: " + err.Error(),
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "ERROR", "", err)
		return
	}
	c.JSON(204, &person)
	data, span := getCarrierContext(c)
	logs.PopulateErrorLogFile(span, data, "INFO", "deleted person record", nil)
}

func ServerInformation(c *gin.Context) {
	randomValue := rand.Intn(100)
	if randomValue > 85 {
		c.JSON(500, gin.H{
			"Error": "Server unavailable for this role",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "ERROR", "server unavailable for this role", nil)
	} else if randomValue < 15 {
		c.JSON(401, gin.H{
			"Warning": "No authorization to use this role",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "WARNING", "No authorization to use this role", nil)
	} else {
		c.JSON(200, gin.H{
			"OK": "Server available and user authorized",
		})
		data, span := getCarrierContext(c)
		logs.PopulateErrorLogFile(span, data, "INFO", "server available and user authorized", nil)
	}
}
