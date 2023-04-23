package controller

import (
	"MyGram/domain"
	"MyGram/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func (h Handler) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	user := domain.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}
	res, err := h.app.UserRegister(user)
	if err != nil {
		helpers.InternalServer(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h Handler) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	user := domain.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}
	passwordClient := user.Password
	user, err := h.app.UserLogin(user)
	if err != nil {
		fmt.Println(err.Error())
		helpers.BadRequest(c, err.Error())
		return
	}
	isValid := helpers.ComparePass([]byte(user.Password), []byte(passwordClient))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err":     "unaothorize",
			"message": "invalid email or password",
		})
		return
	}
	token := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"data":  user,
	})
}
