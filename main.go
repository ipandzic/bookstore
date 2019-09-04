package main

import (
	"main/db"
	"net/http"
	"time"
	"main/api/v1/customers"
	"main/api/v1/authors"
	"main/api/v1/books"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "User-Agent", "Referrer", "Host", "Token"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	v1 := r.Group("/v1")
	{

		// Routes : Authors
		authorRoutes := v1.Group("/authors")
		{
			authorRoutes.POST("/create/", authors.CreateAuthor)
			authorRoutes.GET("/list/", authors.ListAuthors)
			authorRoutes.GET("/details/:id/", authors.GetAuthor)
			authorRoutes.PUT("/update/:id/", authors.UpdateAuthor)
			authorRoutes.DELETE("/delete/:id/", authors.DeleteAuthor)
		}

		// Routes : Books
		bookRoutes := v1.Group("/books")
		{
			bookRoutes.POST("/create/", books.CreateBook)
			bookRoutes.GET("/list/", books.ListBooks)
			bookRoutes.GET("/details/:id/", books.GetBook)
			bookRoutes.PUT("/update/:id/", books.UpdateBook)
			bookRoutes.DELETE("/delete/:id/", books.DeleteBook)
		}

		// Routes : Users
		userRoutes := v1.Group("/customers")
		{
			userRoutes.POST("/create/", customers.CreateCustomer)
			userRoutes.GET("/list/", customers.ListCustomers)
			userRoutes.GET("/details/:id/", customers.GetCustomer)
			userRoutes.PUT("/update/:id/", customers.UpdateCustomer)
			userRoutes.DELETE("/delete/:id/", customers.DeleteCustomer)
		}
	}

	// Routes : No-Match
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
	})

	return r
}


// @title Bookstore API
// @version 1.0
// @description This is the development version of a Bookstore API.

// @contact.name Ivan Pandzic
// @contact.email ivan.pandzic@gmail.com

// @BasePath /v1
func main() {

	db.ConnectDB()

	router := setupRouter()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	s.ListenAndServe()

	db.CloseDB()

}
