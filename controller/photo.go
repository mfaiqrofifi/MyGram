package controller

import (
	"MyGram/domain"
	"MyGram/helpers"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photo := domain.Photo{}
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}
	photo.UserID = userId
	photo, err := h.app.CreatePhoto(photo)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, photo)
}

func (h *Handler) GetAllPhoto(c *gin.Context) {
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	res, err := h.app.GetAllPhoto(userId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)

}

func (h *Handler) GetPhotoById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	photo := domain.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))
	photo.UserID = userId
	photo.ID = uint(photoId)
	res, err := h.app.GetPhotoById(photo)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photo := domain.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}
	photo.UserID = userId
	photo.ID = uint(photoId)
	res, err := h.app.UpdatePhoto(photo)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) DeletePhoto(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// photo := domain.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	// userId := uint(userData["id"].(float64))
	err := h.app.DeletePhoto(photoId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, "succes to delete")
}
