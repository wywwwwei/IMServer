package Service

import (
	"github.com/wywwwwei/IMServer/model"
	"sync"
)

var userRWM *sync.RWMutex = new(sync.RWMutex)
var userCount int = 1
var Users []model.User = make([]model.User,0)

var friendRWM *sync.RWMutex = new(sync.RWMutex)
var Friend map[string][]int = map[string][]int{}

func UserInit()  {
	userRWM.Lock()
	Users = append(Users,model.User{
		UserID:    "1",
		Password:  "123",
		Name:      "Wu",
		Sex:       "Boy",
		Email:     "1191448318@qq.com",
		Signature: "Hello",
	})
	userCount++
	Users = append(Users,model.User{
		UserID:    "2",
		Password:  "123",
		Name:      "WuYW",
		Sex:       "Boy",
		Email:     "1191448318@qq.com",
		Signature: "World",
	})
	userCount++;
	userRWM.Unlock()

	friendRWM.Lock()
	Friend["1"]=[]int{2}
	Friend["2"]=[]int{1}
	friendRWM.Unlock()
}
