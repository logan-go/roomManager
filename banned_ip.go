package roomManager

import (
	"net"
	"sync"
)

var ipList []net.IP
var upListLock sync.RWMutex

func setIPList(list []net.IP) {
	if len(list) == 0 {
		return
	}
	upListLock.Lock()
	defer upListLock.Unlock()
	ipList = list
}

func CheckIP(ip net.IP) bool {
	if len(ipList) == 0 {
		return false
	}
	upListLock.RLock()
	defer upListLock.RUnlock()
	for _, v := range ipList {
		if v.String() == ip.String() {
			return true
		}
	}
	return false
}
