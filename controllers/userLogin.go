package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/KiranMai5472/event-management-tool/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// @Summary UserLogin
// @Description This endpoint gives API Authentication makes a user's account more secure by adding that additional layer of authentication which asked user to enter there user name and password and check if those are valid or not.
// @Tags User Login.
// @Param Body body models.userLogin true "Accept Request Object"
// @Param token header string true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/login [post]
// @Accept json
// @Produce json
func UserLogin(c *gin.Context) {
	logger.LogDebug("Entered in controllers.UserLogin()", Constants.LogFields)
	var user models.User
	c.BindJSON(&user)

	username := user.Username
	//password := user.Password

	userData := models.User{}
	userQueryColumn := "username=?"
	if getUserData := database.DB.Select("password").Where(userQueryColumn, strings.ToLower(username)).First(&userData); getUserData.Error != nil {

		c.JSON(400,
			gin.H{
				Constants.Status:  Constants.Failed,
				Constants.Message: Constants.UserNameIncorrect,
				Constants.Code:    http.StatusBadRequest,
			})
		return
	}

	checkPasswordHash := utils.CheckPasswordHash(user.Password, userData.Password)

	if checkPasswordHash != true {
		c.JSON(400,
			gin.H{
				Constants.Status:  Constants.Failed,
				Constants.Message: Constants.PasswordIncorrect,
				Constants.Code:    http.StatusBadRequest,
			})
		return
	}

	//generation validation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		Constants.Username:   user.Username,
		Constants.Password:   user.Password,
		Constants.ExpiryTime: time.Now().Add(time.Minute * 60).UTC(),
	})

	// read the Jwt key from env file
	config, err := database.LoadConfig(".")
	if err != nil {
		logger.LogError("🚀 Could not load environment variables to get Jwt key🚀", Constants.LogFields)
	}

	var jwtKey = []byte(config.JwtKey)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(400,
			gin.H{
				Constants.Status:  Constants.Failed,
				Constants.Message: Constants.FailedToCreateToken,
				Constants.Code:    http.StatusForbidden,
			})
		return
	}

	// retrive back deta from token
	tokenretrive, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return jwtKey, nil
	})

	if err != nil {
		logger.LogError(err, Constants.LogFields)
	}

	Claims, ok := tokenretrive.Claims.(jwt.MapClaims)
	if !ok || !tokenretrive.Valid {
		c.JSON(400,
			gin.H{
				Constants.Status: http.StatusBadRequest,
				Constants.Error:  Constants.InvalidToken,
			})
		return
	}

	// Convert the map to JSON
	getdata, error := json.Marshal(Claims)
	if error != nil {
		c.JSON(400,
			gin.H{
				Constants.Status: http.StatusBadRequest,
				Constants.Error:  Constants.MarshaingError,
			})
		return
	}
	// Convert the JSON to a struct
	var userClaims *models.Claims
	jsonerror := json.Unmarshal(getdata, &userClaims)
	if jsonerror != nil {
		c.JSON(400,
			gin.H{
				Constants.Status: http.StatusBadRequest,
				Constants.Error:  Constants.UnmarshalError,
			})
		return
	}

	data := models.LoginSuccess{
		Status:   Constants.Success,
		Message:  Constants.UserLoginSuccess,
		Username: user.Username,
		Token:    tokenString,
	}

	c.JSONP(http.StatusOK, data)

}

// isValidToken validates the JWT token
func isValidToken(tokenString string) bool {

	config, err := database.LoadConfig(".")
	if err != nil {
		logger.LogError("🚀 Could not load environment variables to get Jwt key🚀", Constants.LogFields)
	}

	var jwtKey = []byte(config.JwtKey)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for parsing errors and token validity
	if err != nil || !token.Valid {
		return false
	}

	// Extract and parse the expiration time
	expirationTimeString, ok := token.Claims.(jwt.MapClaims)["expiryTime"].(string)
	if !ok {

		return false
	}

	expirationTime, err := time.Parse(time.RFC3339Nano, expirationTimeString)

	if err != nil || time.Now().UTC().After(expirationTime) {
		return false
	}

	return true
}
