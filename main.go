package main

import (
	"SysAgent/config"
	"SysAgent/database"
	"SysAgent/internal/logger"
	"SysAgent/internal/routes"
	"github.com/gin-gonic/gin"
	"strconv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	port := config.ConfigViper.Port

	logger.Log.Infoln("Starting SysAgent...")
	server := gin.Default()

	conn, err := database.ConnectDatabase()
	if err != nil {
		logger.Log.Errorln(err)
		return
	}

	routes.RegisterAgentRoutes(conn, server)

	routes.RegisterMetricsRoute(server)

	_ = server.Run(":" + strconv.Itoa(port))
}
