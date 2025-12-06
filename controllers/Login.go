package controllers

import (
	"elearning-server/utils"
	"elearning-server/utils/auth"
	"elearning-server/utils/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// Login handles user authentication using email and password.
//
// Flow:
// 1. Parse JSON body (email, password)
// 2. Fetch user from database using email
// 3. Validate password using bcrypt
// 4. Parse request metadata:
//      - IP address
//      - Browser, OS, device type
//      - Country, City (Cloudflare headers)
// 5. Generate access + refresh tokens (same as Google flow)
// 6. Save login session in database (auth.SaveLogin)
// 7. Return user, tokens, and metadata in response
//
// Request Body:
//  {
//      "email": "example@gmail.com",
//      "password": "123456"
//  }
//
// Response:
//  200 OK with tokens and user object
//
func Login(c *gin.Context) {

	// ------------------ Parse Request Body ------------------
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(400, gin.H{"error": "Email and password are required"})
		return
	}

	fmt.Println(body.Email, body.Password)

	// ------------------ Fetch User ------------------
	user, err := user.FindUserByEmail(body.Email)
	if err != nil || user == nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// ------------------ Validate Password ------------------
	if !utils.CheckPasswordHash(body.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// ------------------ Device + Browser Info ------------------
	ua := user_agent.New(c.GetHeader("User-Agent"))
	browser, _ := ua.Browser()
	os := ua.OS()

	deviceType := "Desktop"
	if ua.Mobile() {
		deviceType = "Mobile"
	}

	// ------------------ Geo Data ------------------
	ip := c.ClientIP()
	country := c.GetHeader("CF-IPCountry")
	city := c.GetHeader("CF-IPCity")

	// ------------------ Generate Tokens ------------------
	refreshToken, _ := utils.GenerateJWT(user.ID, "refresh")
	accessToken, _ := utils.GenerateJWT(user.ID, "access")

	// ------------------ Save Login Session ------------------
	err = auth.SaveLogin(
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

	// ------------------ Response ------------------
	c.JSON(200, gin.H{
		"user":         user,
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
