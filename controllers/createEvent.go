package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/gin-gonic/gin"
)

// @Summary create Event Api.
// @Description This endpoint allows you to create a event with event details .
// @Tags create Invite Api.
// @Param Body body models.CreateEvent true "Accept Request Object"
// @Param token header string true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/event [post]
// @Accept json
// @Produce json
func CreateEvent(c *gin.Context) {

	logger.LogDebug("Entered in controllers.CreateEvent()", Constants.LogFields)
	// Extracting the token from the header
	token := c.GetHeader(Constants.Authorization)

	// token validation logic
	if !isValidToken(token) {
		c.JSON(http.StatusUnauthorized, gin.H{
			Constants.Status:  Constants.Failed,
			Constants.Message: Constants.InvalidExpiredToken,
			Constants.Code:    http.StatusUnauthorized,
		})
		return
	}

	var input models.GetEvent

	//taking as json object from the body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			Constants.Error: err.Error()})
		return
	}

	//checking the missing input fields
	if strings.EqualFold(input.Name, Constants.EmptyString) || strings.EqualFold(input.Location, Constants.EmptyString) || strings.EqualFold(input.Host, Constants.EmptyString) {
		c.JSON(http.StatusBadRequest, gin.H{
			Constants.Status:  Constants.Failed,
			Constants.Message: Constants.MissingFields,
			Constants.Code:    http.StatusBadRequest,
		})
		return
	}

	//setting the UTC time
	currentDate := time.Now().UTC()

	//query for inserting the fields into the users table
	query := `INSERT INTO "event" ( id,name, date, start_time, end_time, location ,host)
         VALUES (?, ?, ?, ?, ?, ?,?)`

	// Execute the SQL query with the provided data
	err := database.DB.Exec(query, input.ID, input.Name, currentDate, currentDate, currentDate, input.Location, input.Host)

	//checking the error
	if err.Error != nil {
		c.JSON(http.StatusBadGateway,
			gin.H{
				Constants.Message: Constants.UnableToCreateEvent,
			})
		return
	}

	//success validation
	c.JSON(http.StatusCreated,
		gin.H{
			Constants.Status:  Constants.Success,
			Constants.Message: Constants.CreatedSuccessfully,
			Constants.Code:    http.StatusOK,
		})

}
