package roomManager

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Run() {
	if processFunc == nil {
		fmt.Println("信息处理方法未注册......")
		return
	}
	http.HandleFunc("/"+REQUEST_URI, handler)
	http.ListenAndServe(LISTEN_PORT, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	node := &ReciveNode{}
	node.IsAlive = true
	node.Conn = conn
	node.Add()
	processMessage(node)
}

func processMessage(node *ReciveNode) {
	for {
		mType, reader, err := node.Conn.NextReader()
		if err != nil {
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
			if err == io.EOF {
				msg = append(msg, tmp[:length]...)
				break
			}
			msg = append(msg, tmp...)
		}
		//交给注册的方法处理
		processFunc(msg, node)
	}
}
