package roomManager

import (
	"encoding/json"
	"sync"
)

var messageChannel map[string]chan nodeMessage
var messageChannelLock sync.RWMutex

func init() {
	messageChannel = make(map[string]chan nodeMessage)
}

func sendMessageToChannel(roomId string, nm nodeMessage) error {
	if nm.messageType == NODE_MESSAGE_TYPE_SEND_MESSAGE {
		n, err := json.Marshal(nm.body)
		if err != nil {
			return nil
		}
		if CheckMessage(string(n)) {
			return nil
		}
	}
	messageChannelLock.RLock()
	//如果房间不存在，创建一个房间
	if c, ok := messageChannel[roomId]; ok {
		c <- nm
		messageChannelLock.RUnlock()
	} else {
		messageChannelLock.RUnlock()
		messageChannelLock.Lock()
		//创建房间通道
		messageChannel[roomId] = make(chan nodeMessage, 1024)
		messageChannel[roomId] <- nm
		messageChannelLock.Unlock()
		//创建房间实例
		roomObj := &RoomInfo{}
		roomObj.RoomID = roomId
		roomObj.Rows = make([]*RowList, 0, 4)
		roomObj.Lock = &sync.Mutex{}

		//创建新的协程来监控房间
		go daemonReciver(messageChannel[roomId], roomObj)
		go timerForClean(messageChannel[roomId])
		//如果是大厅的话，启动大厅清理协程
		if roomId == "" {
			go CleanHall(roomObj)
		}
	}
	return nil
}
