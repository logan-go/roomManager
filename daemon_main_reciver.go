package roomManager

import "fmt"

//守护进程，获取进程信息后工作
func daemonReciver(c chan nodeMessage, roomInfo *RoomInfo) {
	for {
		s := <-c
		fmt.Println("[房间", roomInfo.RoomID, "获得信息]:", fmt.Sprintf("%+v", s))
		switch s.messageType {
		case NODE_MESSAGE_TYPE_IN_HALL:
			changeRoom(roomInfo, s.body.(*ReciveNode))
		//case NODE_MESSAGE_TYPE_CLOSE_ROOM:
		case NODE_MESSAGE_TYPE_CHANGE_ROOM:
			changeRoom(roomInfo, s.body.(*ReciveNode))
		case NODE_MESSAGE_TYPE_SEND_MESSAGE:
			sendMessage(roomInfo, s.body)
		case NODE_MESSAGE_TYPE_CLEAN_ROOM:
			cleanRoom(roomInfo)
		}
	}
}
