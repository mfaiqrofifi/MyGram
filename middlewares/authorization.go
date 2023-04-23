package middlewares

import (
	"MyGram/config"
	"MyGram/domain"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GORM.DB
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		photo := domain.Photo{}
		err = db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Status Not Found",
				"message": "Data not exist",
			})
			return
		}
		if photo.UserID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorize",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		c.Next()

	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GORM.DB
		commentID, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid Parameter",
			})
			return
		}
		CommentID := uint(commentID)
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		comment := domain.Comment{}
		err = db.Select("user_id").First(&comment, uint(CommentID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Status Not Found",
				"message": "Data not exist",
			})
			return
		}
		err = db.Select("photo_id").First(&comment, uint(CommentID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Status Not Found",
				"message": "Data not exist",
			})
			return
		}
		if comment.UserID != UserID && comment.PhotoID != CommentID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorize",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		c.Next()

	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GORM.DB
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		socialMedia := domain.SocialMedia{}
		err = db.Select("user_id").First(&socialMedia, uint(socialMediaId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Status Not Found",
				"message": "Data not exist",
			})
			return
		}
		if socialMedia.UserID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorize",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		c.Next()

	}
}
