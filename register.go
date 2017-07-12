package roomManager

var processFunc func([]byte)

func RegisterProcessFunc(pFunc func([]byte)) {
	processFunc = pFunc
}
