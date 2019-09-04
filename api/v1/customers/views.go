package customers

import (
	"context"
	"fmt"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// CreateUser documentation
// @Description Customer - Create
// @Router /v1/customers/create/ [post]
func CreateCustomer(c *gin.Context) {

	var requestBody models.Customer
	ctx := context.Background()

	// Bind POST body to model struct
	c.BindJSON(&requestBody)

	if requestBody.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the first name of the user."})
		return
	}

	if requestBody.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter the last name of the user."})
		return
	}

	newUser := models.Customer{
		ID:        requestBody.ID,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		CreatedAt: requestBody.CreatedAt,
		UpdatedAt: requestBody.UpdatedAt,
	}

	if err := newUser.Insert(ctx, db.GetDB(), boil.Infer()); err != nil {
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

// ListCustomers documentation
// @Description Customer - List
// @Router /v1/customers/list/ [get]
func ListCustomers(c *gin.Context) {

	ctx := context.Background()

	customers, err := models.Customers().All(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customers)
}

// GetCustomer documentation
// @Description Customer - Get
// @Router /v1/customers/details/{id}/ [get]
func GetCustomer(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	customer, err := models.Customers(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer ID not found"})
		return
	}

	purchases, _ := customer.Books().All(ctx, db.GetDB())

	c.JSON(200, gin.H{
		"id":         customer.ID,
		"first_name": customer.FirstName,
		"last_name":  customer.LastName,
		"purchases":  purchases,
		"created_at": customer.CreatedAt,
		"updated_at": customer.UpdatedAt,
	})
}

// UpdateCustomer documentation
// @Description Customer - Update
// @Router /v1/customers/update/{id}/ [put]
func UpdateCustomer(c *gin.Context) {

	type updateCustomerData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Purchases []int  `json:"purchases"`
	}

	var (
		requestBody    updateCustomerData
		ParamFirstName string
		ParamLastName  string
		purchasesList  []int
	)
	ctx := context.Background()
	id := c.Params.ByName("id")

	c.BindJSON(&requestBody)

	oldCustomerData, err := models.Customers(qm.Where("id=?", id)).One(ctx, db.GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found"})
		return
	}

	if requestBody.FirstName == "" {
		ParamFirstName = oldCustomerData.FirstName
	} else {
		ParamFirstName = requestBody.FirstName
	}

	if requestBody.LastName == "" {
		ParamLastName = oldCustomerData.LastName
	} else {
		ParamLastName = requestBody.LastName
	}

	for n, item := range requestBody.Purchases {
		fmt.Println(requestBody.Purchases[n], item)

		purchases, _ := models.Books(qm.Where("id=?", requestBody.Purchases[n])).All(ctx, db.GetDB())

		purchasesList = append(purchasesList, item)

		err2 := oldCustomerData.AddBooks(ctx, db.GetDB(), false, purchases...)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}

	}

	_, err2 := models.Customers(qm.Where("id=?", id)).UpdateAll(ctx, db.GetDB(), models.M{
		"first_name": ParamFirstName,
		"last_name":  ParamLastName,
	})

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":         oldCustomerData.ID,
		"first_name": ParamFirstName,
		"last_name":  ParamLastName,
		"purchases":  purchasesList,
		"created_at": oldCustomerData.CreatedAt,
		"updated_at": oldCustomerData.UpdatedAt,
	})
}

// DeleteCustomer documentation
// @Description Customer - Delete
// @Router /v1/customers/delete/{id}/ [delete]
func DeleteCustomer(c *gin.Context) {
	ctx := context.Background()
	id := c.Params.ByName("id")

	// Delete Customer
	customer, _ := models.Customers(qm.Where("id=?", id)).DeleteAll(ctx, db.GetDB())
	if customer == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer ID not found"})
		return
	}

	c.JSON(200, gin.H{
		"deleted_customer": true,
	})
}
