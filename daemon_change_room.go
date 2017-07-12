package roomManager

import "time"

//更换房间标记
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
func changeRoom(roomInfo *RoomInfo, node *ReciveNode) {
	if len(roomInfo.Rows) > 0 {
		roomInfo.Rows[0].BackNode.NextNode = node
		node.NextNode = nil
		node.PrevNode = roomInfo.Rows[0].BackNode
		roomInfo.Rows[0].BackNode = node
		roomInfo.LastChangeTime = time.Now()
		node.CurrentList.Length++
		roomInfo.Length++
	} else {
		roomInfo.Rows[0] = &RowList{}
		roomInfo.Rows[0].BackNode = node
		roomInfo.Rows[0].FrontNode = node
		roomInfo.Rows[0].CurrentRoom = roomInfo
		roomInfo.Rows[0].Length++
		roomInfo.LastChangeTime = time.Now()
		node.NextNode = nil
		node.PrevNode = nil
		node.CurrentList.Length++
	}
}
