package roomManager

import "fmt"

func sendMessage(roomInfo *RoomInfo, message interface{}) {
	counter := 0
	realCounter := 0
	if DETAILED_LOG_FLAG {
		fmt.Println("[发送消息] - 房间", roomInfo.RoomID, "：遍历节点", counter, "个，发送节点", realCounter, "个")
	}
}
