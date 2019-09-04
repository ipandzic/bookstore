package authors

import (
	"context"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// CreateAuthor documentation
// @Description Author - Create
// @Router /v1/authors/create/ [post]
func CreateAuthor(c *gin.Context) {

	var requestBody models.Author
	ctx := context.Background()

	// Bind POST body to model struct
	c.BindJSON(&requestBody)

	if requestBody.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the first name of the author."})
		return
	}

	if requestBody.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the last name of the author."})
		return
	}

	newAuthor := models.Author{
		ID:        requestBody.ID,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		CreatedAt: requestBody.CreatedAt,
		UpdatedAt: requestBody.UpdatedAt,
	}

	if err := newAuthor.Insert(ctx, db.GetDB(), boil.Infer()); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":         requestBody.ID,
			"first_name": requestBody.FirstName,
			"last_name":  requestBody.LastName,
			"created_at": requestBody.CreatedAt,
			"updated_at": requestBody.UpdatedAt,
		})
	}
}

// ListAuthors documentation
// @Description Author - List
// @Router /v1/authors/list/ [get]
func ListAuthors(c *gin.Context) {

	ctx := context.Background()

	// Here we do not check for Company ID but list all companies
	authors, err := models.Authors().All(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, authors)
}

// GetAuthor documentation
// @Description Author - Get
// @Router /v1/authors/details/{id}/ [get]
func GetAuthor(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	author, err := models.Authors(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":         author.ID,
		"first_name": author.FirstName,
		"last_name":  author.LastName,
		"created_at": author.CreatedAt,
		"updated_at": author.UpdatedAt,
	})
}

// UpdateAuthor documentation
// @Description Author - Update
// @Router /v1/authors/update/{id}/ [put]
func UpdateAuthor(c *gin.Context) {

	var (
		requestBody    models.Author
		ParamFirstName string
		ParamLastName  string
	)
	ctx := context.Background()
	id := c.Params.ByName("id")

	c.BindJSON(&requestBody)

	oldAuthorData, err := models.Customers(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author ID not found"})
		return
	}

	if requestBody.FirstName == "" {
		ParamFirstName = oldAuthorData.FirstName
	} else {
		ParamFirstName = requestBody.FirstName
	}

	if requestBody.LastName == "" {
		ParamLastName = oldAuthorData.LastName
	} else {
		ParamLastName = requestBody.LastName
	}

	_, err2 := models.Authors(qm.Where("id=?", id)).UpdateAll(ctx, db.GetDB(), models.M{
		"first_name": ParamFirstName,
		"last_name":  ParamLastName,
	})

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":         oldAuthorData.ID,
		"first_name": ParamFirstName,
		"last_name":  ParamLastName,
		"created_at": oldAuthorData.CreatedAt,
		"updated_at": oldAuthorData.UpdatedAt,
	})
}

// DeleteAuthor documentation
// @Description Author - Delete
// @Router /v1/authors/delete/{id}/ [delete]
func DeleteAuthor(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	// Delete Author
	author, _ := models.Authors(qm.Where("id=?", id)).DeleteAll(ctx, db.GetDB())
	if author == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"deleted_author": true,
	})
}
