package main

import (
	"github.com/OJ-lab/oj-lab-services/model"
	"github.com/OJ-lab/oj-lab-services/user-service/router"
	"github.com/OJ-lab/oj-lab-services/user-service/service"
	"github.com/OJ-lab/oj-lab-services/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	var configPath string
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	} else {
		configPath = "config/default.ini"
	}

	dataBaseSettings, err := utils.GetDatabaseSettings(configPath)
	if err != nil {
		panic("failed to get database settings")
	}
	jwtSettings, err := utils.GetJWTSettings(configPath)
	if err != nil {
		panic("failed to get jwt settings")
	}
	serviceSettings, err := utils.GetServiceSettings(configPath)
	if err != nil {
		panic("failed to get service settings")
	}

	model.OpenConnection(dataBaseSettings)
	utils.SetupJWTSettings(jwtSettings)
	service.SetupServiceSetting(serviceSettings)

	r := gin.Default()
	gin.SetMode(serviceSettings.Mode)
	router.SetupUserRouter(r)

	err = r.Run(serviceSettings.Port)
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
