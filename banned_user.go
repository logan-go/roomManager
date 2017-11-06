package roomManager

import (
	"strings"
	"sync"
)

var userList []string
var userListLock sync.RWMutex

func SetUserList(list []string) {
	if len(list) == 0 {
		return
	}
	userListLock.Lock()
	defer userListLock.Unlock()
	userList = make([]string, 0, 1024)
	for _, v := range list {
		userList = append(userList, strings.ToUpper(v))
	}
}

func CheckUserID(userId string) bool {
	if len(userList) == 0 {
		return false
	}
	if userId == "" {
		return false
	}
	userListLock.RLock()
	defer userListLock.RUnlock()
	for _, v := range userList {
		if v == userId {
			return true
		}
	}
	return false
}
