package controllers

import (
	"fmt"
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

// Login a user - with Gin Basic Auth
func (ac authController) Login(c *gin.Context) {
	// Bind the request body to the LoginSchema
	var loginSchema schema.LoginSchema
	if err := c.ShouldBindJSON(&loginSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Declare a user model
	var userModel models.User

	// Check if the user already exists
	userIns := ac.db.Where("email = ?", loginSchema.Email).First(&userModel)
	log.Println(userIns)
	if userIns.RowsAffected == 0 {
		utils.Response(c, http.StatusNotFound, nil, "User not found")
		return
	}

	// Check if the user is verified
	if !userModel.IsActive {
		utils.Response(c, http.StatusUnauthorized, nil, "User is not verified, Please verify your email")
		return
	}

	// Validate the Password
	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginSchema.Password))
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Incorrect Password")
		return
	}

	// Generate a Basic Auth token
	accessToken, err := utils.GenerateJWTToken(userModel.Email, utils.TokenTypeAccess)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	// Return the user
	utils.Response(c, http.StatusOK, accessToken, "User Login Successful")
}

// Verify a user
func (ac authController) Verify(c *gin.Context) {
	// Get the token from the request
	email := c.Param("email")

	// Declare a user model
	var userModel models.User

	// Check if the user already exists
	userIns := ac.db.Where("email = ?", email).First(&userModel)
	if userIns.RowsAffected == 0 {
		utils.Response(c, http.StatusNotFound, nil, "User not found")
		return
	}

	// Check if the user is already verified
	if userModel.IsActive {
		utils.Response(c, http.StatusConflict, nil, "User is already verified")
		return
	}

	// Activate the user
	userModel.IsActive = true
	if err := ac.db.Save(&userModel).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error activating user")
		return
	}

	utils.Response(c, http.StatusOK, userModel, "User activated successfully")
}

// Forgot Password
func (ac authController) ForgotPassword(c *gin.Context) {
	// Get the token from the request
	email := c.Param("email")

	// Declare a user model
	var userModel models.User

	// Check if the user already exists
	userIns := ac.db.Where("email = ?", email).First(&userModel)
	if userIns.RowsAffected == 0 {
		utils.Response(c, http.StatusNotFound, nil, "User not found")
		return
	}

	// Check if the user is verified
	if !userModel.IsActive {
		utils.Response(c, http.StatusUnauthorized, nil, "User is not verified, Please verify your email")
		return
	}

	// Generate a Basic Auth token
	resetToken, err := utils.GenerateJWTToken(userModel.Email, utils.TokenTypeReset)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	fmt.Println()
	fmt.Println(resetToken)
	fmt.Println()

	resetURL := fmt.Sprintf("http://localhost:8080/reset-password/%s", resetToken)

	// Send the user a reset password email
	log.Println("Sending reset password email to: ", email)
	log.Println("Reset URL: ", resetURL)

	utils.Response(c, http.StatusOK, nil, "Reset password email sent")
}

// Reset Password -
func (ac authController) ResetPassword(c *gin.Context) {
	// Get the token from the request
	token := c.Param("token")

	resetPass := schema.ResetPasswordSchema{}
	if err := c.ShouldBindJSON(&resetPass); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Decode the token
	email, tokenType, err := utils.VerifyJWTToken(token)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	// Check if the token is a Reset token
	if err := utils.CheckTokenType(tokenType, utils.TokenTypeReset); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	// Declare a user model
	var userModel models.User

	// Check if the user already exists
	userIns := ac.db.Where("email = ?", email).First(&userModel)
	if userIns.RowsAffected == 0 {
		utils.Response(c, http.StatusNotFound, nil, "User not found")
		return
	}

	// Validate the Password
	if len(userModel.Password) < 6 {
		utils.Response(c, http.StatusBadRequest, nil, "Password must be at least 6 characters")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}
	userModel.Password = string(hashedPassword)

	// Save the user
	if err := ac.db.Save(&userModel).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.Response(c, http.StatusOK, userModel, "User password reset successful")
}
