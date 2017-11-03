package roomManager

var userList []string

func init() {
	userList = make([]string, 1024)
}

func setUserList(list []string) {
	userList = list
}

func CheckUserID(userId string) bool {
	if len(userList) == 0 {
		return false
	}
	for _, v := range userList {
		if v == userId {
			return true
		}
	}
	return false
}
