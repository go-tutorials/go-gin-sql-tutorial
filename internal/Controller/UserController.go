package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin-sql-rest-api/internal/model"
	"net/http"
)

func GetAll(c *gin.Context) {
	var user []model.Users
	err := model.GetAllUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err ": err})
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.Users

	err := model.GetByIdUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err ": err})
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func Insert(c *gin.Context) {
	var user model.Users
	c.BindJSON(&user)

	err := model.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err ": err})
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func Update(c *gin.Context) {
	var user model.Users
	id := c.Params.ByName("id")
	err := model.GetByIdUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err})
		return
	}
	c.BindJSON(&user)
	err2 := model.UpdateUser(&user, id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err2})
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func Delete(c *gin.Context) {
	var user model.Users
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Id " + id: "is deleted"})
		return
	}
}
