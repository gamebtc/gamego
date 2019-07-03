package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func Echo(w http.ResponseWriter, r *http.Request) {
	//允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			break
		}
		log.Printf("mt:%v, recv: %s", mt, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write err:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", Echo)
	fmt.Println("start server  port:9091")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatal("ListenAndServe:1234:", err)
	}
}
