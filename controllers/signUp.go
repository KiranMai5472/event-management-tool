package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/KiranMai5472/event-management-tool/services"
	"github.com/KiranMai5472/event-management-tool/utils"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp Api.
// @Description This API is used to create a new User and a User Registration in a single request.
// @Tags SignUp.
// @Param Body body models.UserSignUp true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/signup [post]
// @Accept json
// @Produce json
func SignUpHandler(c *gin.Context) {

	var users []models.UserSignUp
	var newUser models.UserSignUp
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request body
	if newUser.Username == "" || newUser.Password == "" || newUser.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	query1 := database.DB.Select("*").Table("users").Scan(&users)

	if query1.Error != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableToUpdateContent)
		return
	}

	fmt.Println(users)
	// Check if the username already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
	}

	//var input models.UserSignUp
	hashedPassword, hashErr := utils.HashPassword(newUser.Password)
	if hashErr != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableToUpdateContent)
		return
	}

	var ID int
	// // Query to find the maximum(Last) ID in the menu table
	maxSeriesMetaIDQuery := `SELECT COALESCE(MAX(id), 0) FROM "users"`
	var LastID int
	if err := database.DB.Raw(maxSeriesMetaIDQuery).Scan(&LastID).Error; err != nil {
		// Handle the error
		logger.LogError(err, Constants.LogFields)
		return
	}

	// Increment the LastID for the new entry
	ID = LastID + 1

	//query for inserting the fields into the users table
	query := `INSERT INTO "users" ( id,username, password, fullname)
         VALUES (?, ?, ?,?)`

	// Execute the SQL query with the provided data
	err := database.DB.Exec(query, ID, newUser.Username, hashedPassword, newUser.FullName)

	//checking the error
	if err.Error != nil {
		c.JSON(http.StatusBadGateway,
			gin.H{
				Constants.Message: Constants.UnableToCreateEvent,
			})
		return
	}

	//newUser.Password = hashedPassword

	// Add the new user to the users slice
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}

// Function to hash the password using SHA-256
func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
