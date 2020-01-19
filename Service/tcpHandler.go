package Service

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/wywwwwei/IMServer/model"
	"io"
	"log"
	"net"
)

func TcpHandler(conn net.Conn){
	fmt.Println("Handling tcp connect from ",conn.RemoteAddr())

	var userID string
	reader := bufio.NewReader(conn)
	data := make([]byte,SERVER_TCP_LEN)
	defer conn.Close()
	for{
		n, err := io.ReadFull(reader, data[:4])
		data = data[:n]
		if err!=nil{
			if userID!=""{
				fmt.Println(userID," exit")
				GetConnManager().DeleteConn(userID)
			}
			if err == io.EOF{
				log.Println(conn.RemoteAddr()," disconnect...")
				return
			}
			log.Printf("%s : Read tcp data error: %s\n",conn.RemoteAddr(),err)
			return
		}
		dataLen := binary.BigEndian.Uint32(data)
		if uint64(dataLen) > uint64(cap(data)) {
			data = make([]byte, 0, dataLen)
		}
		n, err = io.ReadFull(reader, data[:dataLen])
		fmt.Println("n:",n)
		data = data[:n]
		fmt.Println("data : ",data)
		ForwardOrReply(data,conn,&userID)
	}
}

func EncodeMessage(packet model.MessagePacket)[]byte{
	var res *bytes.Buffer = new(bytes.Buffer)
	jsonData,_:= json.Marshal(packet)
	err:=binary.Write(res,binary.BigEndian,uint32(len(jsonData)))
	if err!=nil{
		log.Println("Error occurs when encoding :",err)
	}
	err = binary.Write(res,binary.BigEndian,jsonData)
	if err!=nil{
		log.Println("Error occurs when encoding :",err)
	}
	return res.Bytes()
}

func ForwardOrReply(receive []byte,conn net.Conn,userID *string){
	var message model.MessagePacket
	err := json.Unmarshal(receive,&message)
	fmt.Println("Message:",message)
	if err!=nil{
		log.Println("Error occurs when json unmarshal:",err)
		return
	}

	if message.Type == "Regist"{
		*userID = message.Sender
		GetConnManager().AddConn(message.Sender,&conn)

		UnreadList := GetUnreadManager().GetMessage(message.Sender)
		fmt.Println("Unread:",UnreadList)
		for _,v := range(UnreadList){
			conn.Write(EncodeMessage(v))
		}
		GetUnreadManager().DeleteMessage(message.Sender)
	}else if message.Type == "Message"{
		connManager := GetConnManager()
		unreadManagers := GetUnreadManager()

		replyConn := connManager.GetConn(message.Receiver)
		if replyConn == nil{
			fmt.Println("empty conn")
			unreadManagers.StoreMessage(message)
		}else{
			fmt.Println("have send..")
			(*replyConn).Write(EncodeMessage(message))
		}
	}
	fmt.Println("Handle finish")
}