package controller

import (
	"MyGram/domain"
	"MyGram/helpers"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	comment := domain.Comment{}
	photoId, err := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.UserID = userId
	comment.PhotoID = uint(photoId)
	comment, err = h.app.CreateComment(comment)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, comment)
}

func (h *Handler) GetAllComment(c *gin.Context) {
	// db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	res, err := h.app.GetAllComment(userId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)

}

func (h *Handler) GetCommentById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	comment := domain.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))
	comment.UserID = userId
	comment.PhotoID = uint(photoId)
	comment.ID = uint(commentId)
	res, err := h.app.GetCommentById(comment)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	comment := domain.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.UserID = userId
	comment.ID = uint(photoId)
	comment.PhotoID = uint(commentId)
	res, err := h.app.UpdateComment(comment)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, res)
}

func (h *Handler) DeleteComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// photo := domain.Photo{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	// userId := uint(userData["id"].(float64))
	err := h.app.DeletePhoto(commentId)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}
	helpers.Ok(c, "succes to delete")
}
