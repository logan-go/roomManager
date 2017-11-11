package roomManager

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

var messageChannel map[string]chan nodeMessage
var messageChannelLock sync.RWMutex
var messageRoomCreated map[string]bool

func init() {
	if useBroadcasting {
		messageRoomCreated = make(map[string]bool)
	} else {
		messageChannel = make(map[string]chan nodeMessage)
	}
}

//给channel发送消息
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

//给广播站发送消息
func sendMessageToBroadcast(roomId string, nm nodeMessage) error {
	n, err := json.Marshal(nm.body)
	if err != nil {
		return nil
	}
	if nm.messageType == NODE_MESSAGE_TYPE_SEND_MESSAGE {
		if CheckMessage(string(n)) {
			return WordsBannedError
		}
	}
	w, err := broadcastingConnection.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	w.Write(n)
	w.Close()
	/*
		messageChannelLock.RLock()
		//如果房间不存在，创建一个房间
		if _, ok := messageRoomCreated[roomId]; ok {
			w, err := broadcastingConnection.NextWriter(websocket.TextMessage)
			if err != nil {
				messageChannelLock.RUnlock()
				return err
			}
			w.Write(n)
			w.Close()
			messageChannelLock.RUnlock()
		} else {
			messageChannelLock.RUnlock()
			messageChannelLock.Lock()
			//创建房间通道
			messageChannel[roomId] = make(chan nodeMessage, 1024)
			w, err := broadcastingConnection.NextWriter(websocket.TextMessage)
			if err != nil {
				messageChannelLock.Unlock()
				return err
			}
			w.Write(n)
			w.Close()
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
	*/
	return nil
}
