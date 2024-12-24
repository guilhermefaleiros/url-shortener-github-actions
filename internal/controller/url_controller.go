package controller

import (
	"github.com/gin-gonic/gin"
	"url-shortener-api/internal/service"
)

type CreateUrlRequest struct {
	TargetUrl string `json:"targetUrl"`
}

type UrlController struct {
	service *service.URLService
}

func (uc *UrlController) Insert(c *gin.Context) {
	var request CreateUrlRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := uc.service.Shorten(c, request.TargetUrl)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"hash": hash,
	})
}

func (uc *UrlController) Get(c *gin.Context) {
	hash := c.Param("hash")

	targetUrl, err := uc.service.GetTargetLink(c, hash)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(302, targetUrl)
}

func NewUrlController(urlService *service.URLService) *UrlController {
	return &UrlController{service: urlService}
}
