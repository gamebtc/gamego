package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	//"local.com/abc/game/walle"
)

func main() {
	var name string
	flag.StringVar(&name, "apk", "biantai-106.apk", "apk file name")
	flag.Parse()

	b, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print(err)
	}

	c, e := GetCommentLength(b)
	fmt.Printf("\r\nfile:%v,%v, %v", name, c, e)
	idValues, err := GetAllMap(b)

	if err != nil {
		fmt.Print(err)
	}

	channel := ApkGetRaw(b)
	fmt.Printf("\r\nchannel:%v", string(channel))

	if idValues != nil {
		for i := 0; i < len(idValues.Keys); i++ {
			if APK_SIGNATURE_SCHEME_V2_BLOCK_ID != idValues.Keys[i] {
				fmt.Printf("\r\nkey:%v,value:%v", idValues.Keys[i], string(idValues.Values[i]))
			}
		}
	}
	//规则
	http.HandleFunc("/apk/", apkProcess)
	http.HandleFunc("/", fuck3P)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func apkProcess(w http.ResponseWriter, r *http.Request) {
	x := fmt.Sprintf("url:%v,uri:%v", r.URL, r.RequestURI)
	paths := strings.Split(r.RequestURI, "/")
	l := len(paths)
	if l >= 4 {
		channel := paths[l-3]
		key := paths[l-2]
		name := paths[l-1]
		apkFile, err := NewApkFile(name)
		if err == nil {
			check := apkFile.Check()
			fmt.Printf("\r\nchannel:%v,key:%v,apkFile:%v,check:%v,len:%v", channel, key, name,check,apkFile.CommentLen())
			r := apkFile.CreatFile([]byte(channel + "." + key))
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