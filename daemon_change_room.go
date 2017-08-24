package roomManager

//更换房间标记
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
//新的房间更新逻辑：
//1.扫描所有列，找到长度小于最大长度的，加入队尾
//2.修改结点中索引信息
func changeRoom(roomInfo *RoomInfo, node *ReciveNode) {
	roomInfo.Lock.Lock()
	defer roomInfo.Lock.Unlock()
	//如果房间没有列，则创建一个列
	if len(roomInfo.Rows) == 0 {
		row := &RowList{}
		row.Nodes = make([]*ReciveNode, 0, ROW_LENGTH)
		roomInfo.Rows = append(roomInfo.Rows, row)
	}
	addSuccess := false
	for k, v := range roomInfo.Rows {
		if len(v.Nodes) < ROW_LENGTH {
			node.RowIndex = k
			node.NodeIndex = len(v.Nodes)
			v.Nodes = append(v.Nodes, node)
			addSuccess = true
			break
		}
	}
	if !addSuccess {
		row := &RowList{}
		row.Nodes = make([]*ReciveNode, 0, ROW_LENGTH)
		roomInfo.Rows = append(roomInfo.Rows, row)
		node.RowIndex = len(roomInfo.Rows) - 1
		node.NodeIndex = 0
		row.Nodes = append(row.Nodes, node)
	}
}
