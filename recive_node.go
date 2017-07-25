/**
	整个房间链接管理项目围绕着接收节点的管理来进行管理
**/
package roomManager

import (
	"github.com/gorilla/websocket"
)

//用于接收用户消息的节点
type ReciveNode struct {
	RoomID      string
	NextNode    *ReciveNode
	PrevNode    *ReciveNode
	CurrentList *RowList
	Conn        *websocket.Conn
	IsAlive     bool
}

func (this *ReciveNode) Add() {
	this.IsAlive = true
	this.RoomID = ""
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_IN_HALL,
		body:        this,
	}
	sendMessageToChannel(this.RoomID, nm)
}

func (this *ReciveNode) ChangeRoom(roomId string) {
	if this.RoomID == roomId {
		return
	}
	this.RoomID = roomId
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_CHANGE_ROOM,
		body:        this,
	}
	sendMessageToChannel(this.RoomID, nm)
}

func (this *ReciveNode) SendMessage(message interface{}) {
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_SEND_MESSAGE,
		body:        message,
	}
	if this.RoomID == "" {
		return
	}
	sendMessageToChannel(this.RoomID, nm)
}

func (this *ReciveNode) Close() {
	this.IsAlive = false
	this.Conn.Close()
}

func SendMessageFromOuter(roomID string, message interface{}) {
	if roomID == "" {
		return
	}
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_SEND_MESSAGE,
		body:        message,
	}
	sendMessageToChannel(roomID, nm)
}
