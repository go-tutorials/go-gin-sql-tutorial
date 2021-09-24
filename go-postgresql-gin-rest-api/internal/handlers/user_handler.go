package handlers

import (
	"context"
	sv "github.com/core-go/service"
	"github.com/gin-gonic/gin"
	. "go-service/internal/models"
	. "go-service/internal/services"
	"net/http"
	"reflect"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	result, err := h.service.GetAll(context.Background())
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error " : " don't get all users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data ": result})
}

func (h *UserHandler) Load(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error " : " user not found"})
		return
	}

	result, err := h.service.Load(context.Background(), id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error " : " data missing"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data " : result})
}

func (h *UserHandler) Insert(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	defer c.Request.Body.Close()
	if err != nil {
		c.Error(err)
		return
	}

	_, er2 := h.service.Insert(context.Background(), &user)
	if er2 != nil {
		c.Error(er2)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Insert new user success",
								"data" : &user})

}

func (h *UserHandler) Update(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	defer c.Request.Body.Close()

	if err != nil {
		c.Error(err)
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}

	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Id not match"})
		return
	}

	_, er2 := h.service.Update(context.Background(), &user)
	if er2 != nil {
		c.Error(er2)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Update user success",
								"data" : &user})
}

func (h *UserHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	ids := []string{"id"}

	var user User
	userType := reflect.TypeOf(user)
	_, jsonMap := sv.BuildMapField(userType)
	body, _ := sv.BuildMapAndStruct(c.Request, &user)
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Id not match"})
		return
	}
	json, er1 := sv.BodyToJson(c.Request, user, body, ids, jsonMap, nil)
	if er1 != nil {
		c.Error(er1)
		return
	}

	_, er2 := h.service.Patch(context.Background(), json)
	if er2 != nil {
		c.Error(er2)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"user has patched", "data": "" })
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}

	_, err := h.service.Delete(context.Background(), id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message ": " delete success",
								"data " : true})
}
