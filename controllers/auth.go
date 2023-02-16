package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/user"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/schema"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	db *ent.Client
}

func NewAuthController(db *ent.Client) *authController {
	return &authController{
		db: db,
	}
}

// Sign up a user
func (ac authController) Signup(c *gin.Context) {
	var signupSchema ent.User
	if err := c.ShouldBindJSON(&signupSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Binding user schema")
		return
	}

	// Check if the user already exists
	if isExist := ac.db.User.Query().Where(user.Email(signupSchema.Email)).ExistX(c); isExist {
		utils.Response(c, http.StatusBadRequest, nil, "User already exists")
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

	// Create the user in the database
	userModel, err := ac.db.User.Create().
		SetFirstName(signupSchema.FirstName).
		SetLastName(signupSchema.LastName).
		SetEmail(signupSchema.Email).
		SetPassword(string(hashedPassword)).
		Save(c)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	// Generate a verification token
	token, err := utils.GenerateJWTToken(userModel.Email, schema.TokenTypeVerify)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	// TODO: Send the user a verification email
	log.Println("Sending verification email to: ", token)

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

	// Check if the user exists
	userModel, err := ac.db.User.Query().Where(user.Email(loginSchema.Email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid Credentials")
		return
	}

	// Check if the user is verified
	if !userModel.IsActive {
		utils.Response(c, http.StatusUnauthorized, nil, "User is not verified, Please verify your email")
		return
	}

	// Validate the Password
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginSchema.Password))
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid Credentials")
		return
	}

	// Generate a Basic Auth token
	accessToken, err := utils.GenerateJWTToken(userModel.Email, schema.TokenTypeAccess)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}
	refreshToken, err := utils.GenerateJWTToken(userModel.Email, schema.TokenTypeRefresh)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	tr := schema.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Return the user
	utils.Response(c, http.StatusOK, tr, "User Login Successful")
}

// Verify a user
func (ac authController) Verify(c *gin.Context) {
	// Get the token from the request
	token := c.Param("token")

	// Verify the token
	email, tokenType, err := utils.VerifyJWTToken(token)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid Token")
		return
	}

	// Check if the token is of type verification
	if err := utils.CheckTokenType(tokenType, schema.TokenTypeVerify); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, err.Error())
		return
	}

	// Declare a user model
	var userModel models.User

	// Check if the user already exists
	userIns, err := ac.db.User.Query().Where(user.Email(email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid Token")
		return
	}

	// Check if the user is already verified
	if userModel.IsActive {
		utils.Response(c, http.StatusConflict, nil, "User is already verified")
		return
	}

	// Activate the user
	userModel.IsActive = true
	if err := ac.db.User.UpdateOneID(userIns.ID).SetIsActive(true).Exec(c); err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error activating user")
		return
	}

	utils.Response(c, http.StatusOK, userModel, "User activated successfully")
}

// Forgot Password
func (ac authController) ForgotPassword(c *gin.Context) {
	// Get the token from the request
	email := c.Param("email")

	// Check if the user already exists
	userModel, err := ac.db.User.Query().Where(user.Email(email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid Token")
		return
	}
	// Check if the user is verified
	if !userModel.IsActive {
		utils.Response(c, http.StatusUnauthorized, nil, "User is not verified, Please verify your email")
		// TODO: Send the user a verification email
		return
	}

	// Generate a Basic Auth token
	resetToken, err := utils.GenerateJWTToken(userModel.Email, schema.TokenTypeReset)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	resetURL := fmt.Sprintf("http://localhost:8080/reset-password/%s", resetToken)

	// Send the user a reset password email
	// TODO: Send the user a reset password email
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
	if err := utils.CheckTokenType(tokenType, schema.TokenTypeReset); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	// Declare a user model
	userModel, err := ac.db.User.Query().Where(user.Email(email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	// Validate the Password
	if len(userModel.Password) < 6 {
		utils.Response(c, http.StatusBadRequest, nil, "Password must be at least 6 characters")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPass.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}
	userModel.Password = string(hashedPassword)

	// Save the user
	if err := ac.db.User.UpdateOneID(userModel.ID).SetPassword(userModel.Password).Exec(c); err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.Response(c, http.StatusOK, userModel, "User password reset successful")
}

// Change Password
func (ac authController) ChangePassword(c *gin.Context) {
	// Get the token from the header
	token, err := utils.ParseToken(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		return
	}

	// get the email from the token
	email, tokenType, err := utils.VerifyJWTToken(token)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		log.Println(err)
		c.Abort()
		return
	}

	// Check if the token is a Access token
	if err := utils.CheckTokenType(tokenType, schema.TokenTypeAccess); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		log.Println(err)
		return
	}

	// Get the password from the request
	changePass := schema.ChangePasswordSchema{}
	if err := c.ShouldBindJSON(&changePass); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Check if the old password is correct
	user, err := ac.db.User.Query().Where(user.Email(email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePass.CurrentPassword)); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid password")
		return
	}

	// Validate the Password
	if len(changePass.NewPassword) < 6 {
		utils.Response(c, http.StatusBadRequest, nil, "Password must be at least 6 characters")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePass.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	// Update the user
	user.Password = string(hashedPassword)
	if err := ac.db.User.UpdateOneID(user.ID).SetPassword(user.Password).Exec(c); err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.Response(c, http.StatusOK, user, "User password changed successfully")
}

// Refresh Token
func (ac authController) RefreshToken(c *gin.Context) {
	// Get the token from the header
	token, err := utils.ParseToken(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		return
	}

	// get the email from the token
	email, tokenType, err := utils.VerifyJWTToken(token)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		log.Println(err)
		c.Abort()
		return
	}

	// Check if the token is a Refresh token
	if err := utils.CheckTokenType(tokenType, schema.TokenTypeRefresh); err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Unauthorized, Invalid token")
		log.Println(err)
		return
	}

	// Check if the user exists
	user, err := ac.db.User.Query().Where(user.Email(email)).Only(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, nil, "Invalid token")
		return
	}

	// Generate a new access token
	accessToken, err := utils.GenerateJWTToken(user.Email, schema.TokenTypeAccess)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	// Return the tokens
	tr := schema.TokenResponse{
		AccessToken: accessToken,
	}

	utils.Response(c, http.StatusOK, tr, "Tokens generated successfully")
}
