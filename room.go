package roomManager

import (
	"sync"
	"time"
)

type RoomInfo struct {
	RoomID         string      //房间ID
	Lock           *sync.Mutex //房间操作锁
	Rows           []*RowList  //房间多行Slice
	Length         uint64      //当前房间总节点数
	LastChangeTime time.Time   //最后一次更新时间
}
