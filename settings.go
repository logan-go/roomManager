package roomManager

import (
	"os"
	"time"
)

var (
	DETAILED_LOG_FLAG = true                                             //详细日志开关
	NORMAL_LOG_FLAG   = true                                             //常规日志开关
	TRACE_FLAG        = true                                             //是否打开trace开关
	TRACE_LOG_PATH    = os.ExpandEnv("$GOPATH/trace_logs/trace_log.out") //trace日志地址
)

const (
	REQUEST_URI = "websocket"
	CLEAN_TIMER = 5 * time.Minute //房间清理定时器
)
