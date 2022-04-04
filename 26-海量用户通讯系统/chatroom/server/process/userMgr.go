package processes

import "fmt"

var (
	userMgr *UserMgr
)

// 定义一个 UserMgr，用来维护在线用户列表
//因为整个服务端只维护一个列表，需要在全局使用，所以设为全局
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//添加一个在线用户
func (this *UserMgr) AddOnlineUsers(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除一个下线用户
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)

}

//返回所有在线用户的列表
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据指定Id返回在线用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId] //用带检测的发生返回指定用户
	if !ok {
		fmt.Errorf("用户%d不存在 err=", userId, err)
		return
	}
	return
}
