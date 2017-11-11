package roomManager

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	broadcastingHost = "localhost"
	broadcastingPort = 6666
	broadcastingURI  = "broadcasting"
	useBroadcasting  = false

	broadcastingConnection *websocket.Conn
	err                    error
)

func ConnBroadcasting() {
	tryToConnBroadcastingStation()
	//处理读到消息之后
	go func() {
		for {
			mType, reader, err := broadcastingConnection.NextReader()
			if mType == websocket.CloseMessage || mType == -1 {
				tryToConnBroadcastingStation()
			}
			if err != nil {
				continue
			}
			msg := make([]byte, 1024)
			l, err := reader.Read(msg)
			if err != nil {
				continue
			}
			fmt.Println(string(msg))
			proMessageFromBroadcast(msg[:l])
		}
	}()
}

func tryToConnBroadcastingStation() {
	for connCount := 0; connCount < 10; connCount++ {
		fmt.Print(time.Now().Format("2006-01-02 15:04:05"), "[尝试创建广播站连接] ===========>")
		conn, err := net.Dial("tcp", broadcastingHost+":"+fmt.Sprint(broadcastingPort))
		if err != nil {
			fmt.Println("广播站连接失败：", broadcastingHost+":"+fmt.Sprint(broadcastingPort), err)
			continue
		}
		u := url.URL{}
		u.Host = broadcastingHost + ":" + fmt.Sprint(broadcastingPort)
		u.Path = broadcastingURI
		u.Scheme = "ws"

		h := http.Header{}
		broadcastingConnection, _, err = websocket.NewClient(conn, &u, h, 1024, 1024)
		if err != nil {
			fmt.Println("websocket连接失败：", err)
			conn.Close()
			continue
		}
		fmt.Println("创建成功！")
		return
	}
	os.Exit(0)
}
