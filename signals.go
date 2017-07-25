package roomManager

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ProcessSignals() {
	if DETAILED_LOG_FLAG {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

		_ = <-sig
		fmt.Println("[共建立连接]：", LinkedCounter)
		fmt.Println("[共断开连接]：", ClosedCounter)
		fmt.Println("[共发送消息]：", SendCounter)
		os.Exit(1)
	}
}
