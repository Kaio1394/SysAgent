package main

import (
	"SysAgent/internal/logger"
	"SysAgent/internal/routes"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	logger.Log.Infoln("Starting SysAgent...")
	server := gin.Default()

	routes.MemoryRoutes(server)

	_ = server.Run(":9095")
}
