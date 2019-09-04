package books

import (
	"context"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// CreateBook documentation
// @Description Book - Create
// @Router /v1/books/create/ [post]
func CreateBook(c *gin.Context) {

	var requestBody models.Book
	ctx := context.Background()

	// Bind POST body to model struct
	c.BindJSON(&requestBody)

	if requestBody.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the title."})
		return
	}

	if requestBody.AuthorID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the author ID."})
		return
	}

	newBook := models.Book{
		ID:        requestBody.ID,
		Title:     requestBody.Title,
		AuthorID:  requestBody.AuthorID,
		CreatedAt: requestBody.CreatedAt,
		UpdatedAt: requestBody.UpdatedAt,
	}

	if err := newBook.Insert(ctx, db.GetDB(), boil.Infer()); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"title":      requestBody.Title,
			"author_id":  requestBody.AuthorID,
			"created_at": requestBody.CreatedAt,
			"updated_at": requestBody.UpdatedAt,
		})
	}
}

// ListBooks documentation
// @Description Book - List
// @Router /v1/books/list/ [get]
func ListBooks(c *gin.Context) {

	ctx := context.Background()

	// Here we do not check for Company ID but list all companies
	books, err := models.Books().All(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}

// GetBook documentation
// @Description Book - Get
// @Router /v1/books/details/{id}/ [get]
func GetBook(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	book, err := models.Books(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":         book.ID,
		"title":      book.Title,
		"author_id":  book.AuthorID,
		"created_at": book.CreatedAt,
		"updated_at": book.UpdatedAt,
	})
}

// UpdateBook documentation
// @Description Book - Update
// @Router /v1/books/update/{id}/ [put]
func UpdateBook(c *gin.Context) {

	var (
		requestBody   models.Book
		ParamTitle    string
		ParamAuthorID int
	)
	ctx := context.Background()
	id := c.Params.ByName("id")

	c.BindJSON(&requestBody)

	oldBookData, err := models.Books(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book ID not found"})
		return
	}

	if requestBody.Title == "" {
		ParamTitle = oldBookData.Title
	} else {
		ParamTitle = requestBody.Title
	}

	if requestBody.AuthorID == 0 {
		ParamAuthorID = oldBookData.AuthorID
	} else {
		ParamAuthorID = requestBody.AuthorID
	}

	_, err2 := models.Books(qm.Where("id=?", id)).UpdateAll(ctx, db.GetDB(), models.M{
		"title":     ParamTitle,
		"author_id": ParamAuthorID,
	})

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":         oldBookData.ID,
		"title":      ParamTitle,
		"author_id":  ParamAuthorID,
		"created_at": oldBookData.CreatedAt,
		"updated_at": oldBookData.UpdatedAt,
	})
}

// DeleteBook documentation
// @Description Book - Delete
// @Router /v1/books/delete/{id}/ [delete]
func DeleteBook(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	// Delete Book
	book, _ := models.Books(qm.Where("id=?", id)).DeleteAll(ctx, db.GetDB())
	if book == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"deleted_book": true,
	})
}
