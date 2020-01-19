package Service

import (
	"net"
	"sync"
)

type ConnManager struct {
	connTable sync.Map
}

var once sync.Once
var instance *ConnManager

func GetConnManager()*ConnManager{
	once.Do(func() {
		instance = &ConnManager{connTable:sync.Map{}}
	})
	return instance
}

func (c *ConnManager)AddConn(userID string,conn *net.Conn){
	c.connTable.Store(userID,conn)
}

func (c *ConnManager)DeleteConn(userID string){
	c.connTable.Delete(userID)
}

func (c *ConnManager)GetConn(userID string)*net.Conn{
	res,ok:=c.connTable.Load(userID)
	if ok{
		return res.(*net.Conn)
	}
	return nil
}