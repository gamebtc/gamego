package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"local.com/abc/game/model"
)

var origin []byte

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "listen port")
	flag.Parse()

	model.InitGlobalFiveCard()

	log.Printf("start:%v\r\n", port)
	origin, _ = ioutil.ReadFile("bin")
	http.HandleFunc("/", license)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/pk", pk)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func license(w http.ResponseWriter, r *http.Request) {
	log.Printf("ip:%v,req:%v\r\n", r.RemoteAddr, r.RequestURI)
	w.Write(origin)
}

func favicon(w http.ResponseWriter, r *http.Request) {
}

func pk(w http.ResponseWriter, r *http.Request) {
	bin := make([]byte, 100)
	n, err := r.Body.Read(bin)
	if err == io.EOF || err == nil {
		ab := strings.Split(string(bin[:n]), ";")
		a := model.ParsePokers(ab[0], "|")
		b := model.ParsePokers(ab[1], "|")
		f1, f2 := model.FiveCardPk(a, b)
		r := fmt.Sprintf("\r\n实际胜率:%v, 表面胜率:%v\r\n", f1, f2)
		w.Write([]byte(r))
		return
	}
	w.Write([]byte("error"))
}
