package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/user"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

type userController struct {
	db *ent.Client
}

func NewUserController(db *ent.Client) *userController {
	return &userController{
		db: db,
	}
}

// Get all Users
func (u userController) GetAll(c *gin.Context) {
	ctx := context.Background()

	userRes, err := u.db.User.Query().All(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting users")
		return
	}

	utils.Response(c, http.StatusOK, userRes, "Users found")
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
	user, err := u.db.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			utils.Response(c, http.StatusNotFound, nil, "User not found")
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting user")
	}

	utils.Response(c, http.StatusNotFound, user, "user found")
}

// Create a User
func (u userController) Create(c *gin.Context) {
	var userSchema ent.User
	if err := c.ShouldBindJSON(&userSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error creating user")
		return
	}

	// Create the user in the database
	user, err := u.db.User.Create().
		SetFirstName(userSchema.FirstName).
		SetLastName(userSchema.LastName).
		SetEmail(userSchema.Email).
		SetPassword(userSchema.Password).
		Save(context.Background())
	if err != nil {
		if ent.IsConstraintError(err) && strings.Contains(err.Error(), "email") {
			utils.Response(c, http.StatusConflict, nil, "Email already exists")
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
	err = u.db.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting user")
		return
	}

	utils.Response(c, http.StatusOK, nil, "user deleted successfully")
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
	_, err = u.db.User.UpdateOneID(id).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(context.Background())

	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating user")
		return
	}

	// return the updated user
	utils.Response(c, http.StatusOK, nil, "user updated successfully")
}
