package Service

import (
	"github.com/gin-gonic/gin"
	"github.com/wywwwwei/IMServer/model"
	"net/http"
	"strconv"
)

func ProfileHandler(c *gin.Context){
	id,_:= strconv.Atoi(c.Params.ByName("uid"))
	userRWM.RLock()
	if id > len(Users){
		c.JSON(400,gin.H{"msg":"Out of range"})
		return
	}
	c.JSON(200,Users[id - 1])
	userRWM.RUnlock()
}

func ListHandler(c *gin.Context){
	id:=c.Params.ByName("uid")
	friendRWM.RLock()
	friendList := Friend[id]
	friend :=  make([]model.Friend,0)

	for _,val := range friendList{
		user := Users[val-1]
		friend = append(friend,model.Friend{
			Id:user.UserID,
			Username:user.Name,
		})
	}
	friendRWM.RUnlock()
	c.JSON(200,friend)
}

func UsernameHandler(c *gin.Context){
	id,_:=strconv.Atoi(c.Params.ByName("uid"))
	userRWM.RLock()
	if id > len(Users){
		c.JSON(400,gin.H{"msg":"Out of range"})
		return
	}
	c.JSON(200,gin.H{"name":Users[id-1].Name})
	userRWM.RUnlock()
}

func LoginHandler(c *gin.Context){
	var login struct {
		User     string `json:"user"  binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err:=c.ShouldBindJSON(&login);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	id,_:= strconv.Atoi(login.User)
	userRWM.RLock()
	if id > len(Users){
		c.JSON(400,gin.H{"msg":"Out of range"})
		return
	}
	user := Users[id-1]
	userRWM.RUnlock()
	if user.Password != login.Password{
		c.JSON(400,gin.H{"msg":"Wrong userID or password"})
		return
	}
	c.JSON(200, user)
}

func RegistHandler(c *gin.Context){
	var user model.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	user.UserID = strconv.Itoa(userCount)
	userRWM.Lock()
	Users = append(Users,user)
	userCount++;
	userRWM.Unlock()
	c.JSON(200,user)
}