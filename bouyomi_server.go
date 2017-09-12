package main

import (
	_ "fmt"
	"net"
	_ "os"
	_ "time"
	"log"
	"encoding/binary"
	"net/http"
	"github.com/gorilla/websocket"
)

type postMsgInput struct {
	Msg string
}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func sendMessage(p []byte)  {
	msg := string(p)
	msg_length := uint32(len(msg))
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

	bMsglength := make([]byte, 4)
	binary.LittleEndian.PutUint32(bMsglength, msg_length)

	sData = append(sData, bMsglength...)
	sData = append(sData, msg...)

	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
	}
	_, err = conn.Write(sData)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	_ = conn.Close()
}
func sendMessageWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		sendMessage(p)
		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}
func main() {
	log.Printf("Server started")

	http.HandleFunc("/ws", sendMessageWS)
	log.Fatal(http.ListenAndServe(":9999", nil)) // ポート8080で立ち上がる
}