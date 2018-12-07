package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

//将unit64转为[]byte
func Uint64ToBytes(num uint64)[]byte{
	var buff bytes.Buffer
	err:=binary.Write(&buff,binary.BigEndian,num)
	if err!=nil{
		log.Panic(err)
	}
	return buff.Bytes()
}