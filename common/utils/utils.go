package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"redis.demo/common/message"
)

func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 4096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	// 根据buf[:4] 转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		return
	}
	// 把pkgLen反序列化成message.Message类型
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送消息的长度
	// 转成byte
	pkgLen := uint32(len(data))
	// 因为uint32是4字节，所以只需要开4字节的byte
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[:])
	if err != nil {
		return err
	} else if n != 4 {
		return errors.New("WritePkg send byte error")
	}

	// 发送消息
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return
}
