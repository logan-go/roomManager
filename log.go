package roomManager

import redis "gopkg.in/redis.v5"
import json "github.com/json-iterator/go"

var redisConn *redis.Client

func init() {
	opt := redis.Options{}
	opt.Addr = "localhost:6379"

	redisConn = redis.NewClient(&opt)
}

const (
	LOG_TYPE_SERVER_RECIVE            = iota //服务器端接收到消息数量
	LOG_TYPE_SERVER_INHALL                   //服务器端接收到连接数
	LOG_TYPE_SERVER_INROOM                   //服务器端接收到进入房间的数量（分房间）
	LOG_TYPE_SERVER_INROOM_LEFT_CONNS        //服务器端房间每次介入链接之后，剩余的链接数量总和（不是length，而是遍历之后的节点数量）
)

func Log(logType int, data interface{}) {
	switch logType {
	case LOG_TYPE_SERVER_RECIVE:
		rs, _ := json.Marshal(data)
		redisConn.SAdd("LOG_TYPE_SERVER_RECIVE", string(rs))
	case LOG_TYPE_SERVER_INHALL:
		redisConn.Incr("LOG_TYPE_SERVER_INHALL")
	case LOG_TYPE_SERVER_INROOM:
		redisConn.Incr(data.(string))
	case LOG_TYPE_SERVER_INROOM_LEFT_CONNS:
		rs, _ := json.Marshal(data)
		redisConn.SAdd("LOG_TYPE_SERVER_INROOM_LEFT_CONNS_"+data.(ReciveNode).RoomID, string(rs))
	}
}
