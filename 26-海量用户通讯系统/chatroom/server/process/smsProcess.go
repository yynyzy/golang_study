package processes

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMessage(mes message.Message) {
	//遍历服务端的onlineUsers map[int]*UserProcess，
	//将消息转发取出

	//取出mes的 smsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendToeachUser(smsMes.Content, up.Conn)
	}
}

func (this *SmsProcess) SendToeachUser(info string, conn Conn.net) {

}
