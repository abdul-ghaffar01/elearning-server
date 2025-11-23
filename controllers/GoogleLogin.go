package controllers

import (
	"elearning-server/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GoogleLogin(c *gin.Context) {
	url := lib.GoogleOAuthConfig.AuthCodeURL("random-state-string")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
