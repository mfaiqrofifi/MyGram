package controller

import (
	"MyGram/domain"
	"MyGram/helpers"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socialMedia := domain.SocialMedia{}
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}
	socialMedia.UserID = userId
	scMedia, err := h.app.CreateSocialMedia(socialMedia)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, scMedia)
}

func (h *Handler) GetAllSocialMedia(c *gin.Context) {
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	res, err := h.app.GetAllSocialMedia(userId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)

}

func (h *Handler) GetSocialMediaById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	socialMedia := domain.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))
	socialMedia.UserID = userId
	socialMedia.ID = uint(socialMediaId)
	res, err := h.app.GetSocialMediaById(socialMedia)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socialMedia := domain.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}
	socialMedia.UserID = userId
	socialMedia.ID = uint(socialMediaId)
	res, err := h.app.UpdateSocialMedia(socialMedia)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) DeleteSocialMedia(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// photo := domain.Photo{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	// userId := uint(userData["id"].(float64))
	err := h.app.DeleteSocialMedia(socialMediaId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, "succes to delete")
}
