package main

import (
	"fmt"
	_ "net"
	_ "os"
	_ "time"
	"encoding/binary"
)

type Data struct {
	msg string
	msg_length uint32
}

func main() {
	var data Data
	//var headData [15]byte
	data.msg = "やあ(´・ω・｀)ようこそ、バーボンハウスへ。このテキーラはサービスだから、まず飲んで落ち着いて欲しい。うん、「また」なんだ。済まない。仏の顔もって言うしね、謝って許してもらおうとも思っていない。でも、このタイトルを見たとき、君は、きっと言葉では言い表せない「ときめき」みたいなものを感じてくれたと思う。殺伐とした世の中で、そういう気持ちを忘れないで欲しい。そう思って、このブロマガを作ったんだ。ちなみにタイトルには255文字まで使えるらしいんだがどういうふうに表示されるんだろうか。じゃあ、注文を聞こうか。あ"
	data.msg_length = uint32(len(data.msg))
	iCommand := []byte{1, 0}
	iSpeed := []byte{100, 0}
	iTone := []byte{255, 255}
	iVolume := []byte{255, 255}
	iVoice := []byte{0, 0}
	bCode := []byte{0}

	sData := append(iCommand, iSpeed...)
	sData = append(sData, iTone...)
	sData = append(sData, iVolume...)
	sData = append(sData, iVoice...)
	sData = append(sData, bCode...)
	//sData = append(sData, []uint8(data.msg_length)...)

	h := data.msg_length
	bMsglength := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, h)

	//y := uint8(data.msg_length & 0xff)

	fmt.Println(sData)
	fmt.Println(data.msg_length)
	fmt.Println(a)
}




