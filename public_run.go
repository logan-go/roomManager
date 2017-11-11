package roomManager

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Run() error {
	if processFunc == nil {
		fmt.Println("信息处理方法未注册......")
		return nil
	}
	go ProcessSignals()
	if useBroadcasting {
		go ConnBroadcasting()
	}
	http.HandleFunc("/"+REQUEST_URI, handler)
	return http.ListenAndServe(LISTEN_PORT, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	node := &ReciveNode{}

	//获取目标IP地址
	node.IP = net.ParseIP(r.RemoteAddr)

	node.IsAlive = true
	node.Conn = conn
	node.Add()
	processMessage(node)
}

func processMessage(node *ReciveNode) {
	for {
		mType, reader, err := node.Conn.NextReader()

		if mType == -1 {
			node.Close()
			return
		}

		if err != nil {
			continue
		}

		//如果已经被关闭，则不允许发送消息
		if node.DisableRead {
			continue
		}

		//如果用户ID在黑名单内，则关闭发消息功能
		if CheckUserID(node.UserID) {
			node.DisableRead = true
			continue
		}

		switch mType {
		case websocket.BinaryMessage:
			continue
		case websocket.CloseMessage:
			node.Close()
			return
		case websocket.PingMessage:
			node.Conn.WriteMessage(websocket.PongMessage, []byte(""))
			continue
		case websocket.PongMessage:
			continue
		}
		//如果是文本内容的话
		//获取文本内容
		msg := make([]byte, 0, 1024)
		for {
			tmp := make([]byte, 1024)
			length, err := reader.Read(tmp)
			if err == io.EOF || length < 1024 {
				msg = append(msg, tmp[:length]...)
				break
			}
			msg = append(msg, tmp...)
		}
		//交给注册的方法处理
		processFunc(msg, node)
	}
}
