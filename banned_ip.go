package roomManager

import (
	"sync"
)

var ipList []string
var upListLock sync.RWMutex

func SetIPList(list []string) {
	if len(list) == 0 {
		return
	}
	upListLock.Lock()
	defer upListLock.Unlock()
	ipList = list
}

func CheckIP(ip string) bool {
	if len(ipList) == 0 {
		return false
	}
	upListLock.RLock()
	defer upListLock.RUnlock()
	for _, v := range ipList {
		if v == ip {
			return true
		}
	}
	return false
}
