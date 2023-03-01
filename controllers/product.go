package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rama-kairi/blog-api-golang-gin/ent"
	"github.com/rama-kairi/blog-api-golang-gin/ent/product"
	"github.com/rama-kairi/blog-api-golang-gin/ent/user"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

type productController struct {
	db *ent.Client
}

func NewProductController(db *ent.Client) *productController {
	return &productController{
		db: db,
	}
}

// Get all Products
func (p productController) GetAll(c *gin.Context) {
	ctx := context.Background()

	productRes, err := p.db.Product.
		Query().
		WithUser().
		WithCategory().
		WithSubCategory().
		All(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting products")
		return
	}

	utils.Response(c, http.StatusOK, productRes, "Products found")
}

// Get a product
func (p productController) Get(c *gin.Context) {
	ctx := context.Background()
	// Get product id from url
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the product id")
		return
	}

	// Get the product from the database
	product, err := p.db.Product.Query().
		WithCategory().
		WithSubCategory().
		WithUser().
		Where(product.ID(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.Response(c, http.StatusNotFound, nil, "Product not found")
			return
		}
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting product")
	}

	utils.Response(c, http.StatusNotFound, product, "product found")
}

// Create a product
func (p productController) Create(c *gin.Context) {
	ctx := context.Background()

	// Get the user from the context
	user := c.MustGet("user").(*ent.User)

	var productSchema ent.Product
	if err := c.ShouldBindJSON(&productSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Binding product schema")
		return
	}

	product, err := p.db.Product.Create().
		SetPrice(productSchema.Price).
		SetDescription(productSchema.Description).
		SetName(productSchema.Name).
		SetUserID(user.ID).
		SetCategoryID(productSchema.CategoryID).
		SetSubCategoryID(productSchema.SubCategoryID).
		Save(ctx)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "missing required field \"Product.user_id\"") {
			utils.Response(c, http.StatusBadRequest, nil, "`user_id` Required")
			return
		}
		if ent.IsConstraintError(err) && strings.Contains(err.Error(), "products_categories_products") {
			utils.Response(c, http.StatusBadRequest, nil, "Category id is invalid")
			return
		}
		if ent.IsConstraintError(err) && strings.Contains(err.Error(), "subcategory") {
			utils.Response(c, http.StatusBadRequest, nil, "Category id is invalid")
			return
		}

		utils.Response(c, http.StatusInternalServerError, nil, "Error creating product")
		return
	}

	utils.Response(c, http.StatusCreated, product, "Product created")
}

// Update a product
func (p productController) Update(c *gin.Context) {
	ctx := context.Background()
	// Get product id from url
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the product id")
		return
	}

	var productSchema ent.Product
	if err := c.ShouldBindJSON(&productSchema); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Binding product schema")
		return
	}

	productUpdate := p.db.Product.UpdateOneID(id)
	if productSchema.Name != "" {
		productUpdate.SetName(productSchema.Name)
	}
	if productSchema.Description != "" {
		productUpdate.SetDescription(productSchema.Description)
	}
	if productSchema.Price != 0 {
		productUpdate.SetPrice(productSchema.Price)
	}

	if productSchema.UserID != uuid.Nil {
		productUpdate.SetCategoryID(productSchema.CategoryID)
	}

	if productSchema.CategoryID != uuid.Nil {
		productUpdate.SetCategoryID(productSchema.CategoryID)
	}

	if productSchema.SubCategoryID != uuid.Nil {
		productUpdate.SetSubCategoryID(productSchema.SubCategoryID)
	}

	productRes, err := productUpdate.Save(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating product")
		return
	}

	utils.Response(c, http.StatusOK, productRes, "Product updated")
}

// Delete a product
func (p productController) Delete(c *gin.Context) {
	ctx := context.Background()
	// Get product id from url
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the product id")
		return
	}

	err = p.db.Product.DeleteOneID(id).Exec(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting product")
		return
	}

	utils.Response(c, http.StatusOK, nil, "Product deleted")
}

// Get All Products by user
func (p productController) GetAllByUser(c *gin.Context) {
	ctx := context.Background()
	// Get user id from url
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error Parsing the user id")
		return
	}

	productRes, err := p.db.Product.Query().
		Where(product.HasUserWith(
			user.ID(id),
		)).
		All(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting products")
		return
	}

	utils.Response(c, http.StatusOK, productRes, "Products found")
}
