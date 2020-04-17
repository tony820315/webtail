package main

import (
	"fmt"
	"os"
	"path"
	"websocket-tail/controller"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logPath := "./webtail.log"
	file, err := os.OpenFile(path.Join(logPath), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err", err)
	}
	logrus.SetOutput(file)
}

func main() {
	bindAddress := "localhost:9780"
	router := gin.Default()
	router.GET("/log", controller.WSHandler)
	router.Run(bindAddress)
}
