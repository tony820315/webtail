package model

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

func TailLog(conn *websocket.Conn, fileName string) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t, err := tail.TailFile(fileName, config)
	if err != nil {
		logrus.Infof("tail.TailFile error:%v", err)
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-t.Lines
		if !ok {
			logrus.Infof("tail file close reopen, filename:%s\n", t.Filename)
			time.Sleep(time.Second)
			continue
		}
		conn.WriteMessage(websocket.TextMessage, []byte(line.Text))
		fmt.Println(line.Text)
	}
}
