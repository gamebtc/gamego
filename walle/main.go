package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"walle/internal"
)

func main() {
	var name string
	flag.StringVar(&name, "apk", "biantai-106.apk", "apk file name")
	flag.Parse()

	b, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print(err)
	}

	c, e := internal.GetCommentLength(b)
	fmt.Printf("\r\nfile:%v,%v, %v", name, c, e)
	idValues, err := internal.GetAllMap(b)

	if err != nil {
		fmt.Print(err)
	}

	channel := internal.ApkGetRaw(b)
	fmt.Printf("\r\nchannel:%v", string(channel))

	if idValues != nil {
		for i := 0; i < len(idValues.Keys); i++ {
			if internal.APK_SIGNATURE_SCHEME_V2_BLOCK_ID != idValues.Keys[i] {
				fmt.Printf("\r\nkey:%v,value:%v", idValues.Keys[i], string(idValues.Values[i]))
			}
		}
	}
	//规则
	http.HandleFunc("/apk/", apkProcess)
	http.HandleFunc("/", fuck3P)
	log.Fatal(http.ListenAndServe(":82", nil))
}

type ChannArgs struct {
	Channel string      `json:"channel"`
	Key     string      `json:"key"`
	Arg     string      `json:"arg"`
	Ip      string      `json:"ip"`
	Time    string      `json:"time"`
	Head    http.Header `json:"header"`
	//Agent   string      `json:"agent"`
}

func getClientIp(r *http.Request)string{
    if ips:= r.Header.Get("X-Forwarded-For");ips != "" {
		return strings.Split(ips, ",")[0]
	}
	if ips:= r.Header.Get("X-Real-IP");ips != "" {
		return strings.Split(ips, ",")[0]
	}
	return r.RemoteAddr
}

func apkProcess(w http.ResponseWriter, r *http.Request) {
	x := fmt.Sprintf("url:%v,uri:%v", r.URL, r.RequestURI)
	paths := strings.Split(r.RequestURI, "/")
	l := len(paths)
	if l >= 4 {
		channel := paths[l-3]
		key := paths[l-2]
		name := strings.Split(paths[l-1], "?")[0]
		r.UserAgent()
		apkFile, err := internal.NewApkFile(name)
		if err == nil {
			check := apkFile.Check()
			fmt.Printf("\r\nchannel:%v,key:%v,apkFile:%v,check:%v,len:%v", channel, key, name, check, apkFile.CommentLen())
			data, _ := json.Marshal(&ChannArgs{
				Channel: channel,
				Key:     key,
				//Agent:   r.UserAgent(),
				Arg:     r.URL.RawQuery,
				Ip:      getClientIp(r),
				Time:    time.Now().Format("2006-01-02 15:04:05"), //20060102150405
				Head:    r.Header,
			})
			w.Header().Set("Content-Type","application/apk")
			r := apkFile.CreatFile(data)
			for _, d := range r {
				w.Write(d)
			}
			return
		}
	}
	w.Write([]byte(x))
}

func fuck3P(w http.ResponseWriter, r *http.Request) {
	origin, _ := ioutil.ReadFile("1111.bin")
	w.Write(origin)
}