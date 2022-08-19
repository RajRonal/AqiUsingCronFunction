package main

import (
	"AqiWithCron/database"
	"AqiWithCron/database/utilities"
	"AqiWithCron/server"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5432", "aqi", "local", "local", database.SSLModeDisable)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	fmt.Println("connected")

	srv := server.SetupRoutes()
	utilities.RunCronJobs()
	err = srv.Run(":8080")
	if err != nil {
		logrus.Error(err)
	}
	//fmt.Println(data.Data)
}
