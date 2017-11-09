package roomManager

var processFunc func([]byte, *ReciveNode)
var proMessageFromBroadcast func([]byte)

//注册消息处理方法
func RegisterProcessFunc(pFunc func([]byte, *ReciveNode)) {
	processFunc = pFunc
}

//注册广播站地址
func RegisterBroadcastStation(domain string, port int, uri string) {
	broadcastingHost = domain
	broadcastingPort = port
	broadcastingURI = uri
	useBroadcasting = true
}

//注册从广播站收到消息后处理的
func RegisterProcessMessageFromBroadcast(pFunc func([]byte)) {
	proMessageFromBroadcast = pFunc
}
