package roomManager

import "net"

var ipList []net.IP

func init() {
	ipList = make([]net.IP, 1024)
}

func setIPList(list []net.IP) {
	ipList = list
}

func CheckIP(ip net.IP) bool {
	if len(ipList) == 0 {
		return false
	}
	for _, v := range ipList {
		if v.String() == ip.String {
			return true
		}
	}
	return false
}
