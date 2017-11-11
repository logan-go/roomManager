package roomManager

import (
	"errors"
	"strings"
)

var wordList []string
var WordsBannedError = errors.New("words in message is banned")

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
