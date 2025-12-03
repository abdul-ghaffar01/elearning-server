package controllers

import (
	"context"
	"elearning-server/lib"
	"elearning-server/utils"
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	// Exchange code for token
	token, err := lib.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to exchange token"})
		return
	}

	// Fetch Google profile
	client := lib.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch Google profile"})
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	var googleUser struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
		Name    string `json:"name"`
	}

	json.Unmarshal(data, &googleUser)

	// ---------------------------------------
	// 🔍 Extract Metadata
	// ---------------------------------------

	// IP
	ip := c.ClientIP()

	// User Agent Parsing
	ua := user_agent.New(c.GetHeader("User-Agent"))
	browser, _ := ua.Browser()
	os := ua.OS()

	deviceType := "Desktop"
	if ua.Mobile() {
		deviceType = "Mobile"
	}

	// Geo (Cloudflare headers if available)
	country := c.GetHeader("CF-IPCountry")
	city := c.GetHeader("CF-IPCity")

	refreshToken, _ := utils.GenerateJWT(googleUser.ID, "refresh")
	accessToken, _ := utils.GenerateJWT(googleUser.ID, "access")

	// Checking if user exists else creating a new user
	user, err := lib.FindOrCreateUser(googleUser.Name, googleUser.Email, googleUser.Picture)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to find or create user"})
		return
	}

	// saving login in the database
	err = lib.SaveLogin(
		user.ID,
		refreshToken,
		ip,
		deviceType,
		os,
		browser,
		country,
		city,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save login"})
		return
	}

	log.Println(user)
	c.JSON(200, gin.H{
		"user":         googleUser,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"ip":           ip,
		"deviceType":   deviceType,
		"browser":      browser,
		"os":           os,
		"country":      country,
		"city":         city,
	})
}
