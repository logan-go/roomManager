/**
	整个房间链接管理项目围绕着接收节点的管理来进行管理
**/
package roomManager

import (
	"net"
	"time"

	"github.com/gorilla/websocket"
	json "github.com/json-iterator/go"
)

//用于接收用户消息的节点
type ReciveNode struct {
	RoomID       string          //房间ID
	ClientID     int64           //客户端ID
	IP           net.IP          //当前IP地址
	UserID       string          //用户标识
	DisableRead  bool            //是否停止接收该链接内容
	Conn         *websocket.Conn //websocket链接
	UpdateTime   time.Time       //最后一次整理时间
	LastSendTime time.Time       //最后一次发送消息时间
	IsAlive      bool            //是否存活
}

func (this *ReciveNode) Add() {
	this.IsAlive = true
	this.RoomID = ""
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_IN_HALL,
		body:        this,
	}
	sendMessageToChannel(this.RoomID, nm)
	if DETAILED_LOG_FLAG {
		LinkedCounter++
	}
}

func (this *ReciveNode) ChangeRoom(roomId string) {
	if this.RoomID == roomId {
		return
	}
	if CheckIP(this.IP) {
		this.DisableRead = true
	}
	this.RoomID = roomId
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_CHANGE_ROOM,
		body:        this,
	}
	sendMessageToChannel(this.RoomID, nm)
}

//给当前节点所在房间发送消息
func (this *ReciveNode) SendMessageToRoom(message interface{}) {
	nm := nodeMessage{
		messageType: NODE_MESSAGE_TYPE_SEND_MESSAGE,
		body:        message,
	}
	if this.RoomID == "" {
		return
	}
	sendMessageToChannel(this.RoomID, nm)
	if DETAILED_LOG_FLAG {
		SendCounter++
	}
}

//给当前节点连接发送消息
func (this *ReciveNode) SendMessage(message interface{}, sendTime time.Time) {
	if this.LastSendTime == sendTime {
		return
	}
	w, err := this.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	defer w.Close()
	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	w.Write(msg)
	this.LastSendTime = sendTime
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
	if DETAILED_LOG_FLAG {
		SendCounter++
	}
}
