package roomManager

import (
	"time"
)

func init() {
	//每分钟清理大厅
}

func CleanHall(roomObj *RoomInfo) {
	if roomObj.RoomID == "" {
		return
	}
	for _, rows := range roomObj.Rows {
		for e := rows.FrontNode; e != nil; e = e.NextNode {
			if e.IsAlive && time.Now().After(e.UpdateTime.Add(HALL_TIMEOUT)) {
				e.Close()
			}
		}
	}
}
