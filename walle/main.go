package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	Channel string `json:"channel"`
	Key     string `json:"key"`
}

func apkProcess(w http.ResponseWriter, r *http.Request) {
	x := fmt.Sprintf("url:%v,uri:%v", r.URL, r.RequestURI)
	paths := strings.Split(r.RequestURI, "/")
	l := len(paths)
	if l >= 4 {
		channel := paths[l-3]
		key := paths[l-2]
		name := paths[l-1]
		apkFile, err := internal.NewApkFile(name)
		if err == nil {
			check := apkFile.Check()
			fmt.Printf("\r\nchannel:%v,key:%v,apkFile:%v,check:%v,len:%v", channel, key, name, check, apkFile.CommentLen())
			data, _ := json.Marshal(&ChannArgs{
				Channel: channel,
				Key:     key,
			})
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