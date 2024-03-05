package controllers

import (
	"net/http"
	"time"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// @Summary Get Token.
// @Description This endpoint gives an access token for use with the API.
// Each array item contains an object with a category which has both the "Enabled" and "Enable on Homepage" flags set to true,
// as well as an array of channel data.
// @Tags Get Token.
// @Success 200 {object} models.Token
// @Param Body body models.UserToken true "Accept Request Object"
// @Router /user/v1/token [post]
// @Accept json
// @Produce json
func GetToken(c *gin.Context) {

	logger.LogDebug("Entered in controllers.GetToken()", Constants.LogFields)

	// assign the user struct to variable
	var user *models.UserToken
	// binding the request body to get data in struct
	//c.BindJSON(&user)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			Constants.Error: err.Error()})
		return
	}

	// request body data assign to variable
	username := user.UserName
	password := user.Password

	//generation of token with claims and expiry of 30 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		Constants.Username:   username,
		Constants.Password:   password,
		Constants.ExpiryTime: time.Now().Add(time.Minute * 60).UTC(),
	})

	// read the Jwt key from env file
	config, err := database.LoadConfig(".")
	if err != nil {
		logger.LogError("🚀 Could not load environment variables to get Jwt key🚀", Constants.LogFields)
	}

	var jwtKey = []byte(config.JwtKey)
	// convertion of token in string with secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(400,
			gin.H{
				Constants.Status: http.StatusBadRequest,
				Constants.Error:  Constants.FailedToCreateToken,
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		Constants.Status: http.StatusOK,
		Constants.Token:  tokenString,
	})

}
