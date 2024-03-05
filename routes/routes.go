package routes

import (
	"net/http"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/controllers"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter function is used for the to get router and end points
func SetupRouter() *gin.Engine {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"x-client-token", "Authorization", "Content-Type", "token"}
	r.Use(cors.New(config))

	// health checker api to test the health
	r.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	//user group apis used to get the users data like role,permissions,login and logout etc.
	userGrp := r.Group("/user/v1")
	{

		//creating the user login
		userGrp.POST("/login", controllers.UserLogin)
		// generating the required token for validation
		userGrp.POST("/token", controllers.GetToken)
		//creating the event
		userGrp.POST("/event", controllers.CreateEvent)
		//creating the event
		userGrp.POST("/signup", controllers.SignUpHandler)
		//invite users
		userGrp.POST("/invite", controllers.InviteHandler)
		//accept invite
		userGrp.POST("/accept", controllers.AcceptInviteHandler)
		//accepthandler within time
		userGrp.POST("/acceptslot/:user_id", controllers.AcceptInviteInSlot)
	}

	logger.LogDebug("Entered in SetupRouter()", Constants.LogFields)
	return r
}
