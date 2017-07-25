package roomManager

import "time"

//清理当前房间里面不属于自己房间的节点
func cleanRoom(roomInfo *RoomInfo) {
	for _, v := range roomInfo.Rows {
		for n := v.FrontNode; n != nil; n = n.NextNode {
			if n.RoomID != roomInfo.RoomID || n.IsAlive == false {
				//如果是第一个节点
				if n == v.FrontNode {
					v.FrontNode = n.NextNode
					if v.FrontNode != nil {//如果只有一个节点的话
						v.FrontNode.PrevNode = nil
					}
				} else if n.NextNode == nil {//如果是最后一个节点
					n.PrevNode.NextNode = nil
				} else {
					n.PrevNode.NextNode = n.NextNode
					n.NextNode.PrevNode = n.PrevNode
				}
				roomInfo.Length--
				n.CurrentList.Length--
			}
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
