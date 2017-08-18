package roomManager

import (
	"fmt"
	"time"
)

//更换房间标记
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
func changeRoom(roomInfo *RoomInfo, node *ReciveNode) {
	if roomInfo.RoomID == "000001" {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>房间开始添加节点>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		if len(roomInfo.Rows) > 0 {
			fmt.Print("[列头节点地址]：")
			fmt.Printf("%p\n", roomInfo.Rows[0].FrontNode)
			fmt.Print("[列尾节点地址]：")
			fmt.Printf("%p\n", roomInfo.Rows[0].BackNode)
		} else {
			fmt.Println("[列头节点地址]：nil")
			fmt.Println("[列尾节点地址]：nil")
		}
		fmt.Print("[当前节点]：")
		fmt.Printf("%+v\n", node)
		fmt.Println("==========更换后=============")
	}
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
	if roomInfo.RoomID == "000001" {
		fmt.Print("[列头节点地址]：")
		fmt.Printf("%p\n", roomInfo.Rows[0].FrontNode)
		fmt.Print("[列尾节点地址]：")
		fmt.Printf("%p\n", roomInfo.Rows[0].BackNode)
		fmt.Print("[当前节点地址]：")
		fmt.Printf("%p\n", node)
		fmt.Print("[当前节点]：")
		fmt.Printf("%+v\n", node)
		fmt.Print("[前一节点]：")
		if node.PrevNode == nil {
			fmt.Println("nil")
		} else {
			fmt.Printf("%+v\n", node.PrevNode)
		}
		leng := 0
		totalLeng := 0
		for e := roomInfo.Rows[0].FrontNode; e != nil; e = e.NextNode {
			totalLeng++
			if e.IsAlive {
				leng++
			}
		}
		fmt.Println("【总长度】：", totalLeng, "【存活长度】：", leng)
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<房间添加节点结束<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	}
}
