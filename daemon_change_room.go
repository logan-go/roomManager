package roomManager

import (
	"fmt"
	"time"
)

//更换房间标记
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
func changeRoom(roomInfo *RoomInfo, node *ReciveNode) {
	if len(roomInfo.Rows) > 0 {
		roomInfo.Rows[0].Lock.Lock()
		roomInfo.Rows[0].BackNode.NextNode = node
		node.NextNode = nil
		node.PrevNode = roomInfo.Rows[0].BackNode
		roomInfo.Rows[0].BackNode = node
		roomInfo.Rows[0].Lock.Unlock()
		roomInfo.LastChangeTime = time.Now()
		node.CurrentList = roomInfo.Rows[0]
		node.CurrentList.Length++
		roomInfo.Length++
	} else {
		roomInfo.Rows = append(roomInfo.Rows, &RowList{})
		roomInfo.Rows[0].Lock.Lock()
		roomInfo.Rows[0].BackNode = node
		roomInfo.Rows[0].FrontNode = node
		roomInfo.Rows[0].CurrentRoom = roomInfo
		roomInfo.Rows[0].Lock.Unlock()
		roomInfo.LastChangeTime = time.Now()
		node.NextNode = nil
		node.PrevNode = nil
		node.CurrentList = roomInfo.Rows[0]
		node.CurrentList.Length++
		roomInfo.Length++
	}
	//展示一下当前房间的情况
	if DETAILED_LOG_FLAG {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>房间信息>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Print("[要处理的节点地址为]：")
		fmt.Printf("%p\n", node)
		rId := "大厅"
		if roomInfo.RoomID != "" {
			rId = roomInfo.RoomID
		}
		fmt.Println("[房间ID]：", rId)
		fmt.Println("[房间人数]：", roomInfo.Length)
		fmt.Println("[房间最后一次修改时间]：", roomInfo.LastChangeTime.Format("2006-01-02 15:04:05"))
		fmt.Print("[当前列表第一节点地址为]：")
		fmt.Println("[房间连接列表]：")
		for k, v := range roomInfo.Rows {
			fmt.Println("	[列号]：", k)
			fmt.Print("[列表第一节点为]：")
			fmt.Printf("%p\n", v.FrontNode)
			fmt.Print("[列表末尾节点为]：")
			fmt.Printf("%p\n", v.BackNode)
			for e := v.FrontNode; e != nil; e = e.NextNode {
				fmt.Print("		[节点内容]：")
				fmt.Printf("%+v", e)
				fmt.Print("[当前节点地址]：")
				fmt.Printf("%p\n", e)
			}
		}
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<房间信息<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	}
}
