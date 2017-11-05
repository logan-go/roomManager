package roomManager

import (
	"strings"
)

var wordList []string

func init() {
	wordList = make([]string, 0, 1024)
}

func SetWordList(list []string) {
	wordList = list
}

func CheckMessage(msg string) bool {
	if len(wordList) == 0 {
		return false
	}
	for _, v := range wordList {
		if strings.Index(msg, v) >= 0 {
			return true
		}
	}
	return false
}
