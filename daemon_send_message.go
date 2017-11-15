package roomManager

import (
	"fmt"
	"time"
)

func sendMessage(roomInfo *RoomInfo, message interface{}) {
	counter := 0
	realCounter := 0
	currentTime := time.Now()

	for _, rows := range roomInfo.Rows {
		for _, node := range rows.Nodes {
			counter++
			if node.RoomID == roomInfo.RoomID && node.IsAlive {
				realCounter++
				node.SendMessage(message, currentTime)
			}
		}
	}

	endTime := time.Now()
	fmt.Println("[发送消息] - 房间", roomInfo.RoomID, "：遍历节点", counter, "个，发送节点", realCounter, "个,耗时：", endTime.Sub(currentTime).String())
}
