package controllers

import (
	"net/http"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/KiranMai5472/event-management-tool/services"
	"github.com/gin-gonic/gin"
)

// @Summary Accept Invite.
// @Description This endpoint gives info of the acceptence the invitation.
// @Tags Accept Invite.
// @Param Body body models.AcceptRequest true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/accept [post]
// @Accept json
// @Produce json
func AcceptInviteHandler(c *gin.Context) {
	// var acceptRequest struct {
	// 	EventID uint `json:"event_id" binding:"required"`
	// 	UserID  uint `json:"user_id" binding:"required"`
	// }
	var invitations []models.Invitation
	var acceptRequest models.AcceptRequest
	if err := c.ShouldBindJSON(&acceptRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the invitation exists
	var found bool
	for i, inv := range invitations {
		if inv.EventID == acceptRequest.EventID && inv.UserID == acceptRequest.UserID && inv.Message == "pending" {
			invitations[i].Message = "accepted"
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation not found or already accepted"})
		return
	}

	Contentquery := `INSERT INTO "invitation"(user_id, event_id, message) VALUES (?, ?, ?)`
	Contentresult := database.DB.Exec(Contentquery, invitations[0].UserID, invitations[0].EventID, invitations[0].Message)

	// Checking the error
	if Contentresult.Error != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableToUpdateContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation accepted successfully"})
}
