package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/utils"

	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController() *userController {
	return &userController{
		db: db.Db,
	}
}

// Get all Users
func (u userController) GetAll(c *gin.Context) {
	var users []models.User
	// Get all users from the database
	u.db.Find(&users)

	utils.Response(c, http.StatusOK, users, "users found")
}

// Get a user
func (u userController) Get(c *gin.Context) {
	// Get user id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}

	// Get the user from the database
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		// If the user is not found, return 404
		utils.Response(c, http.StatusNotFound, nil, "user not found")
		return
	}

	utils.Response(c, http.StatusNotFound, user, "user found")
}

// Create a User
func (u userController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error creating user")
		return
	}

	// Create the user in the database
	if err := u.db.Create(&user).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			utils.Response(c, http.StatusBadRequest, nil, "Email already exists")
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	// Marshal the user into json
	utils.Response(c, http.StatusCreated, user, "user created successfully")
}

// Delete a user
func (u userController) Delete(c *gin.Context) {
	// Get user id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}
	// Delete the user from the database
	if err := u.db.Delete(&models.User{}, id).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting user")
		return
	}

	if err != nil {
		// If the user is not found, return 404
		utils.Response(c, http.StatusNoContent, nil, "user Deleted")
	}
}

// Update a user
func (u userController) Update(c *gin.Context) {
	// Get user id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}
	// Get the user payload from the request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error updating user")
		return
	}

	// Update the user in the database
	if err := u.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating user")
		return
	}

	// return the updated user
	utils.Response(c, http.StatusOK, nil, "user updated successfully")
}
