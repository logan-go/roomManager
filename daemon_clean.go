package roomManager

import (
	"time"
)

//清理当前房间里面不属于自己房间的节点
func cleanRoom(roomInfo *RoomInfo) {
	roomInfo.Lock.Lock()
	defer roomInfo.Lock.Unlock()

	//创建一个空的列组，准备装整理过的节点
	rows := &RowList{}
	nodeList := make([]*ReciveNode, 0, ROW_LENGTH)
	rows.Nodes = append(rows.Nodes, nodeList)

	//循环列表内的节点
	for _, row := range roomInfo.Rows {
		for _, node := range row.Nodes {

		}
	}
	roomInfo.LastChangeTime = time.Now()
}

//定时对房间进行清理
func timerForClean(c chan nodeMessage) {
	for {
		nm := nodeMessage{}
		nm.messageType = NODE_MESSAGE_TYPE_CLEAN_ROOM
		c <- nm
		time.Sleep(CLEAN_TIMER)
	}
}
