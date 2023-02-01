package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/schema"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authController struct {
	db *gorm.DB
}

func NewAuthController() *authController {
	return &authController{
		db: db.Db,
	}
}

// Sign up a user
func (ac authController) Signup(c *gin.Context) {
	var signupSchema schema.SignupSchema
	if err := c.ShouldBindJSON(&signupSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Binding user schema")
		return
	}

	var userModel models.User

	// Check if the user already exists
	userIns := ac.db.Where("email = ?", signupSchema.Email).First(&userModel)
	if userIns.RowsAffected != 0 {
		utils.Response(c, http.StatusConflict, nil, "User already exists")
		return
	}

	// Validate the Password
	if len(signupSchema.Password) < 6 {
		utils.Response(c, http.StatusBadRequest, nil, "Password must be at least 6 characters")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupSchema.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}
	userModel.Password = string(hashedPassword)

	// Set the rest of the user fields
	userModel.Email = signupSchema.Email
	userModel.FirstName = signupSchema.FirstName
	userModel.LastName = signupSchema.LastName

	// TODO: Send the user a verification email
	log.Println("Sending verification email to: ", signupSchema.Email)

	// Create the user
	if err := ac.db.Create(&userModel).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.Response(c, http.StatusCreated, userModel, "User Signup Successful")
}
