package controllers

import (
	"context"
	"elearning-server/lib"
	"elearning-server/utils"
	"elearning-server/utils/auth"
	"elearning-server/utils/user"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	// Exchange code for token
	token, err := lib.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Error exchanging token:", err)
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

	// IP
	ip := c.ClientIP()

	// User Agent Parsing
	ua := user_agent.New(c.GetHeader("User-Agent"))
	browser, _ := ua.Browser()
	userOs := ua.OS()

	deviceType := "Desktop"
	if ua.Mobile() {
		deviceType = "Mobile"
	}

	// Geo (Cloudflare headers if available)
	country := c.GetHeader("CF-IPCountry")
	city := c.GetHeader("CF-IPCity")

	// Checking if user exists else creating a new user
	user, err := user.FindOrCreateUser(googleUser.Name, googleUser.Email, googleUser.Picture)

	// Generating tokens
	refreshToken, _ := utils.GenerateJWT(user.ID, "refresh")
	accessToken, _ := utils.GenerateJWT(user.ID, "access")

	if err != nil {
		log.Println("Error in FindOrCreateUser function:", err)
		c.JSON(500, gin.H{"error": "Failed to find or create user"})
		return
	}

	// saving login in the database
	err = auth.SaveLogin(
		user.ID,
		refreshToken,
		ip,
		deviceType,
		userOs,
		browser,
		country,
		city,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save login"})
		return
	}

	
	// Set refresh token cookie
	backendUrl := os.Getenv("BASE_URL")
	c.SetCookie(
		"refreshToken",   // cookie name
		refreshToken,     // value
		3600*24*7,        // maxAge (7 days)
		"/",              // path
		backendUrl, // domain (use "" for localhost)
		true,             // secure (true in production)
		true,             // httpOnly
	)
	
	// redirecting to the "frontend/login/google" with tokens and user info
	c.Redirect(302, "https://your-frontend-url.com/login/google?accessToken="+accessToken)

}
