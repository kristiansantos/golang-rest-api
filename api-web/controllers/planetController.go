package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kristiansantos/go-api-rest/api/contexts"
	"github.com/kristiansantos/go-api-rest/api/models"
)

func Index(c *gin.Context) {
	var planets []models.Planet
	err := contexts.ListPlanets(&planets)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, planets)
	}
}

func Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var planet models.Planet
	err := contexts.GetByPlanet(&planet, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, planet)
	}
}

func Create(c *gin.Context) {
	var planet models.Planet

	if err := c.ShouldBind(&planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
	} else {
		err := contexts.CreatePlanet(&planet)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, planet)
		}
	}
}

func Update(c *gin.Context) {
	var planet models.Planet
	id := c.Params.ByName("id")
	err := contexts.GetByPlanet(&planet, id)

	if err != nil {
		c.JSON(http.StatusNotFound, planet)
	}

	if err := c.ShouldBind(&planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
	} else {
		err = contexts.UpdatePlanet(&planet, id)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, planet)
		}
	}
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var planet models.Planet
	err := contexts.GetByPlanet(&planet, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		contexts.DeletePlanet(&planet, id)
		c.JSON(http.StatusNoContent, "")
	}
}
