package main

import (
	"fmt"
	"gin-rest-api/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// Login logs in a user
//
// @Summary			Logs in a user
// @Description	Logs in a user
// @Tags				auth
// @Accept			json
// @Produce			json
// @Param				user	body			loginRequest	true	"User"
// @Success			200		{object}	loginResponse
// @Router			/api/v1/auth/login	[post]
func (app *application) login(c *gin.Context) {
	var auth loginRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := app.models.Users.GetByEmail(auth.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user."})
		return
	}
	if existingUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or password is invalid."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(auth.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password is incorrect."})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": existingUser.Id,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(app.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token."})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: tokenString})
}

// RegisterUser	registers a new user
// @Summary			Registers a new user
// @Description	Registers a new user
// @Tags				auth
// @Accept			json
// @Produce			json
// @Param				user	body			registerRequest	true	"User"
// @Success			201		{object}	database.User
// @Router			/api/v1/auth/register [post]
func (app *application) registerUser(c *gin.Context) {
	var register registerRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please try again later."})
		return
	}

	register.Password = string(hashedPassword)
	user := database.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: register.Password,
	}

	err = app.models.Users.Insert(&user)
	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please try again later."})
		return
	}

	c.JSON(http.StatusCreated, user)
}
