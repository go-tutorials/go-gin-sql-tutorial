package handler

import (
	"github.com/core-go/core"
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
	res, err := h.service.All(c.Request.Context())
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

	res, err := h.service.Load(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, res)
	} else {
		c.JSON(http.StatusOK, res)
	}

}

func (h *UserHandler) Insert(c *gin.Context) {
	var user User
	er1 := c.ShouldBindJSON(&user)

	defer c.Request.Body.Close()
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}

	res, er2 := h.service.Insert(c.Request.Context(), &user)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	if res > 0 {
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusConflict, res)
	}
}

func (h *UserHandler) Update(c *gin.Context) {
	var user User
	er1 := c.BindJSON(&user)
	defer c.Request.Body.Close()

	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
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

	res, er2 := h.service.Update(c.Request.Context(), &user)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	if res > 0 {
		c.JSON(http.StatusOK, user)
	} else if res == 0 {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusConflict, res)
	}
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
	_, jsonMap, _ := core.BuildMapField(userType)
	body, er0 := core.BuildMapAndStruct(r, &user)
	if er0 != nil {
		c.String(http.StatusInternalServerError, er0.Error())
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		c.String(http.StatusBadRequest, "Id not match")
		return
	}
	json, er1 := core.BodyToJsonMap(r, user, body, []string{"id"}, jsonMap)
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}

	res, er2 := h.service.Patch(r.Context(), json)
	if er2 != nil {
		c.String(http.StatusInternalServerError, er2.Error())
		return
	}
	if res > 0 {
		c.JSON(http.StatusOK, json)
	} else if res == 0 {
		c.JSON(http.StatusNotFound, json)
	} else {
		c.JSON(http.StatusConflict, res)
	}
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}

	res, err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if res > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, res)
	}
}
