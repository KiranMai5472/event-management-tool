package controllers

import (
	//"io/ioutil"
	"fmt"
	"net/http"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/models"
	"github.com/KiranMai5472/event-management-tool/services"
	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
)

// @Summary Accept Invite in slot Api.
// @Description This endpoint gives info of the acceptence the invitation in the respective slots or time period which means if the slot is alraedy ocupied then he cannot join the other event but if the timmings is out of that event then event can be accepted.
// @Tags Accept Invite in slot.
// @Param Body body models.SlotAccept true "Accept Request Object"
// @Success 200 {object} models.OutPutOfInvite
// @Router /user/v1/acceptslot/:user_id [post]
// @Accept json
// @Produce json
func AcceptInviteInSlot(c *gin.Context) {
	var slotAccept models.SlotAccept
	Id := c.Param("user_id")
	var eventId []int64

	var newslotEvent models.GetEvent
	var input models.Invitation

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			Constants.Message: "not fould data",
			Constants.Error:   err.Error(),
		})
		return
	}

	fmt.Println(input.EventID)

	newEvent := database.DB.Select("*").Table("event").Where("id=?", input.EventID).Scan(&newslotEvent)
	if newEvent.Error != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableTOFetchData)
		return
	}

	newEventStartTime := newslotEvent.StartTime
	newEventEndTime := newslotEvent.EndTime

	fmt.Println(newEventStartTime)
	fmt.Println(newEventEndTime)

	if err := database.DB.Table("invitation").
		Where("user_id = ? AND message = ?", Id, "accepted").
		Pluck("event_id", &eventId).Error; err != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableToUpdateContent)
		return
	}

	var eventExists models.GetEvent
	for i, _ := range eventId {
		// Populate the slotAccept struct
		slotAccept.EventAccept = int(eventId[i])

		result := database.DB.Select("*").Table("event").Where("id=?", eventId[i]).Scan(&eventExists)
		if result.Error != nil {
			services.HandleError(c, http.StatusBadRequest, Constants.UnableTOFetchData)
			return
		}

	}

	// Retrieve the event start and end times from the eventExists struct

	startTime := eventExists.StartTime
	endTime := eventExists.EndTime

	fmt.Println(startTime)
	fmt.Println(endTime)

	// Check if the invitation time falls within the event time slot
	//invitationTime := time.Now() // Assuming the current time is used for invitation time

	if newEventStartTime.Before(endTime) && newEventEndTime.After(startTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot accept invitation, as it conflicts with event time slot"})
		return
	}

	Contentquery := `INSERT INTO "invitation"(user_id, event_id, message) VALUES (?, ?, ?)`
	Contentresult := database.DB.Exec(Contentquery, Id, input.EventID, "accepted")
	if Contentresult.Error != nil {
		services.HandleError(c, http.StatusBadRequest, Constants.UnableTOFetchData)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "For New Event slot accepted successfully",
	})
}
