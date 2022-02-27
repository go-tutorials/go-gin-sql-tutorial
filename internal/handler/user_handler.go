package handler

import (
	"context"
	sv "github.com/core-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"

	. "go-service/internal/model"
	. "go-service/internal/service"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) All(c *gin.Context) {
	res, err := h.service.All(context.Background())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Load(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}

	res, err := h.service.Load(context.Background(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Insert(c *gin.Context) {
	var user User
	er1 := c.ShouldBindJSON(&user)

	defer c.Request.Body.Close()
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}

	res, er2 := h.service.Insert(context.Background(), &user)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(c *gin.Context) {
	var user User
	er1 := c.BindJSON(&user)
	defer c.Request.Body.Close()

	if er1 != nil {
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		c.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}

	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		c.String(http.StatusBadRequest, "Id not match")
		return
	}

	res, er2 := h.service.Update(context.Background(), &user)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}

	r := c.Request
	var user User
	userType := reflect.TypeOf(user)
	_, jsonMap, _ := sv.BuildMapField(userType)
	body, _ := sv.BuildMapAndStruct(r, &user)
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		c.String(http.StatusBadRequest, "Id not match")
		return
	}
	json, er1 := sv.BodyToJsonMap(r, user, body, []string{"id"}, jsonMap)
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}

	res, er2 := h.service.Patch(context.Background(), json)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}

	res, err := h.service.Delete(context.Background(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
