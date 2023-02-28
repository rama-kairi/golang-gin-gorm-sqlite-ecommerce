package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/category"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

type categoryController struct {
	db *ent.Client
}

func NewCategoryController(db *ent.Client) *categoryController {
	return &categoryController{
		db: db,
	}
}

// Get all Categories
func (cc categoryController) GetAll(c *gin.Context) {
	ctx := context.Background()

	categoryRes, err := cc.db.Category.Query().All(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting categories")
		return
	}

	utils.Response(c, http.StatusOK, categoryRes, "Categories found")
}

// Get a category
func (cc categoryController) Get(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the category id")
		return
	}

	category, err := cc.db.Category.Query().Where(category.ID(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.Response(c, http.StatusNotFound, nil, "Category not found")
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting category")
	}

	utils.Response(c, http.StatusNotFound, category, "category found")
}

// Create a category
func (cc categoryController) Create(c *gin.Context) {
	ctx := context.Background()

	var category ent.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error binding category")
		return
	}

	categoryCreate := cc.db.Category.Create()

	if category.Name != "" {
		categoryCreate.SetName(category.Name)
	}

	// Create the category
	categoryRes, err := categoryCreate.Save(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating category")
		return
	}

	utils.Response(c, http.StatusOK, categoryRes, "Category created")
}

// Update a category
func (cc categoryController) Update(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the category id")
		return
	}

	var category ent.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error binding category")
		return
	}

	categoryUpdate := cc.db.Category.UpdateOneID(id)

	if category.Name != "" {
		categoryUpdate.SetName(category.Name)
	}

	// Update the category
	categoryRes, err := categoryUpdate.Save(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating category")
		return
	}

	utils.Response(c, http.StatusOK, categoryRes, "Category updated")
}

// Delete a category
func (cc categoryController) Delete(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the category id")
		return
	}

	// Delete the category
	if err = cc.db.Category.DeleteOneID(id).Exec(ctx); err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting category")
		return
	}

	utils.Response(c, http.StatusOK, nil, "Category deleted")
}
