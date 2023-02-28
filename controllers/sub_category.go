package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/subcategory"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

type subCategoryController struct {
	db *ent.Client
}

func NewSubCategoryController(db *ent.Client) *subCategoryController {
	return &subCategoryController{
		db: db,
	}
}

// Get all Categories
func (cc subCategoryController) GetAll(c *gin.Context) {
	ctx := context.Background()

	subCategoryRes, err := cc.db.SubCategory.Query().All(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting categories")
		return
	}

	utils.Response(c, http.StatusOK, subCategoryRes, "Categories found")
}

// Get a subCategory
func (cc subCategoryController) Get(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the subCategory id")
		return
	}

	subCategory, err := cc.db.SubCategory.Query().Where(subcategory.ID(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.Response(c, http.StatusNotFound, nil, "subCategory not found")
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting subCategory")
	}

	utils.Response(c, http.StatusNotFound, subCategory, "subCategory found")
}

// Create a subCategory
func (cc subCategoryController) Create(c *gin.Context) {
	ctx := context.Background()

	var subCategory ent.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error binding subCategory")
		return
	}

	subCategoryCreate := cc.db.SubCategory.Create()

	if subCategory.Name != "" {
		subCategoryCreate.SetName(subCategory.Name)
	}
	if subCategory.CategoryID != uuid.Nil {
		subCategoryCreate.SetCategoryID(subCategory.CategoryID)
	}

	// Create the subCategory
	subCategoryRes, err := subCategoryCreate.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			utils.Response(c, http.StatusBadRequest, nil, err.Error())
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating subCategory")
		return
	}

	utils.Response(c, http.StatusOK, subCategoryRes, "subCategory created")
}

// Update a subCategory
func (cc subCategoryController) Update(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the subCategory id")
		return
	}

	var subCategory ent.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error binding subCategory")
		return
	}

	subCategoryUpdate := cc.db.SubCategory.UpdateOneID(id)

	if subCategory.Name != "" {
		subCategoryUpdate.SetName(subCategory.Name)
	}

	if subCategory.CategoryID != uuid.Nil {
		subCategoryUpdate.SetCategoryID(subCategory.CategoryID)
	}

	// Update the subCategory
	subCategoryRes, err := subCategoryUpdate.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			utils.Response(c, http.StatusBadRequest, nil, err.Error())
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating subCategory")
		return
	}
	utils.Response(c, http.StatusOK, subCategoryRes, "subCategory updated")
}

// Delete a subCategory
func (cc subCategoryController) Delete(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the subCategory id")
		return
	}

	// Delete the subCategory
	if err = cc.db.SubCategory.DeleteOneID(id).Exec(ctx); err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting subCategory")
		return
	}

	utils.Response(c, http.StatusOK, nil, "subCategory deleted")
}
