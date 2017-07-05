package roomManager

type RowList struct {
	FrontNode   *ReciveNode //第一个节点的指针地址
	BackNode    *ReciveNode //最后一个节点的指针地址d
	Length      uint64      //当前行节点数量
	CurrentRoom *RoomInfo   //当前所在房间指针地址
}
