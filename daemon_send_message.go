package roomManager

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func sendMessage(roomInfo *RoomInfo, message interface{}) {
	startTime := time.Now()
	counter := 0
	realCounter := 0
	for _, v := range roomInfo.Rows {
		for n := v.FrontNode; n != nil; n = n.NextNode {
			counter++
			if n.RoomID == roomInfo.RoomID && n.IsAlive {
				wc, err := n.Conn.NextWriter(websocket.TextMessage)
				if err != nil {
					continue
				}
				rs, err := json.Marshal(message)
				if err != nil {
					continue
				}
				wc.Write(rs)
				wc.Close()
				realCounter++
			}
		}
	}
	endTime := time.Now()
	fmt.Println("[本次经过节点]：", counter, "[有效节点]：", realCounter, "[耗时]：", endTime.UnixNano()-startTime.UnixNano(), "纳秒")
}
