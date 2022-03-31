package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //正式传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据")
	//conn.Read() 在conn没有关闭的情况下才会阻塞，如果客户端关闭要

	_, err = this.Conn.Read(this.Buf[:4]) //读取消息的头四个字节，获取应得到的长度
	//这里的buf[:4] 是指：从 conn套接字中读取了 4个字节的数据到 buf中
	//返回的n是指读到了多少个字节
	if err != nil {
		return
	}

	//根据 buf[:4] 转成一个 uint32类型
	var pkglen uint32 = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据  从套接字中读取 pkglen 长度的数据到 buf 中去
	n, err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		return
	}

	//根据 pkglen 读取相应长度的数据-> 反序列化为 message.Message
	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkglen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkglen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		err = errors.New("conn.Write(bytes) fail err")
		return
	}
	// fmt.Printf("客户端发送信息的长度=%d,内容是=%s\n", pkglen, string(data))

	//发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		err = errors.New("conn.Write(data) fail err")
		return
	}
	return
}
