package main

import (
	"crypto/sha512"
	"crypto/md5"
	"hash/crc64"
	"fmt"
	"hash/crc32"
)

func main(){
	str:="这是一个字符串"
	crc32Table:=crc32.MakeTable(crc32.IEEE)
	hashVal:=crc32.Checksum([]byte(str), crc32Table)
	fmt.Println(hashVal)
	crc64Table:=crc64.MakeTable(crc64.ECMA)
	hash64Val := crc64.Checksum([]byte(str), crc64Table)
	fmt.Println(hash64Val)
	//hashMd5 := md5.New().Sum([]byte(str))
	md5Str:=md5.Sum([]byte(str))
	fmt.Printf("%x\n",md5Str)
	sha512Str := sha512.Sum512([]byte(str))
	//%x转成16进制打印
	fmt.Printf("%x\n",sha512Str)
	
}