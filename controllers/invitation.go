package controllers

import (
	"net/http"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/KiranMai5472/event-management-tool/services"
	"github.com/gin-gonic/gin"
)

// @Summary Send Invite Api.
// @Description This endpoint gives info of the sending the invitation.
// @Tags Send Invite.
// @Param Body body models.AcceptRequest true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/invite [post]
// @Accept json
// @Produce json
func InviteHandler(c *gin.Context) {

	var invitations []models.Invitation
	var newInvitation models.Invitation
	if err := c.ShouldBindJSON(&newInvitation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the event and user exist (mock implementation)
	if !eventExists(newInvitation.EventID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event does not exist"})
		return
	}

	if !userExists(newInvitation.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	// Add the new invitation to the invitations slice
	newInvitation.Message = "pending"
	invitations = append(invitations, newInvitation)

	Contentquery := `INSERT INTO "invitation"(user_id, event_id, message) VALUES (?, ?, ?)`
	Contentresult := database.DB.Exec(Contentquery, newInvitation.UserID, newInvitation.EventID, newInvitation.Message)

	// Checking the error
	if Contentresult.Error != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableToUpdateContent)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invitation sent successfully"})
}

func eventExists(eventID uint) bool {
	return true
}

func userExists(userID uint) bool {
	return true
}
