package controller

import (
	"net/http"
	"path/filepath"
	"websocket-tail/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func WSHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// conn, err := upgrader.Upgrade(r, w, w.Header())
	if err != nil {
		logrus.Infof("upgrader.Upgrade error:%v", err)
		// http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	filename := filepath.Clean(model.Args.Path)

	go model.TailLog(conn, filename)
	// c.Request.WriteHeader(http.StatusUnauthorized)
}
