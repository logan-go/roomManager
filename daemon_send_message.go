package roomManager

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

func sendMessage(roomInfo *RoomInfo, message interface{}) {
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
}
