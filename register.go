package roomManager

var processFunc func([]byte, *ReciveNode)

func RegisterProcessFunc(pFunc func([]byte, *ReciveNode)) {
	processFunc = pFunc
}
