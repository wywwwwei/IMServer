package Service

import (
	"github.com/wywwwwei/IMServer/model"
	"sync"
)

type UnreadManager struct {
	unread sync.Map
}
var onces sync.Once
var unreadManager *UnreadManager

func GetUnreadManager()*UnreadManager{
	onces.Do(func() {
		unreadManager = &UnreadManager{unread:sync.Map{}}
	})
	return unreadManager
}

func (um *UnreadManager)StoreMessage(packet model.MessagePacket){
	UnreadList,ok:= um.unread.Load(packet.Receiver)
	if ok{
		NewUnreadList := UnreadList.([]model.MessagePacket)
		NewUnreadList = append(NewUnreadList,packet)
		um.unread.Store(packet.Receiver,NewUnreadList)
		return
	}
	unreads := []model.MessagePacket{packet}
	um.unread.Store(packet.Receiver,unreads)
	return
}

func (um *UnreadManager)GetMessage(userID string)[]model.MessagePacket{
	var result []model.MessagePacket
	UnreadList,ok:= um.unread.Load(userID)
	if ok{
		result = UnreadList.([]model.MessagePacket)
	}
	return result
}

func (um *UnreadManager)DeleteMessage(userID string){
	um.unread.Delete(userID)
}