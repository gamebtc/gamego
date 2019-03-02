package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/util"
)

func exists(x []string, y string) bool {
	for _, item := range x {
		if item == y {
			return true
		}
	}
	return false
}

var config *dnsfixConfig

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()
	defer util.PrintPanicStack()
	// open profiling
	config = InitConfig(name)

	period := time.Duration(config.Period) * time.Second
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		if lookup() {
			writeHost(config.Host)
		}
		select {
		case <-ticker.C: // 帧更新
		}
	}
}

func lookup() bool {
	changed := false
	newHost := make(map[string][]string, len(config.Host))
	for k, _ := range config.Host {
		addrs, _ := lookupHost(k)
		newHost[k] = addrs
	}
	for k, v := range newHost {
		if oldV, ok := config.Host[k]; ok {
			if len(v) != len(oldV) {
				changed = true
			}
			if len(oldV) > 0 {
				for _, addr := range oldV {
					if exists(v, addr) == false {
						v = append(v, addr)
						newHost[k] = v
						changed = true
					}
				}
			}
		} else {
			changed = true
		}
		log.Debugf("host:%v, addr:%v", k, v)
		if len(v) > 3 {
			newHost[k] = v[0:3]
		}
	}
	if changed {
		config.Host = newHost
	}
	log.Debugf("changed:%v", changed)
	return changed
}

func lookupHost(host string) (addrs []string, err error) {
	for _, a := range config.Dns {
		m := new(dns.Msg)
		m.SetQuestion(host+".", dns.TypeA)
		in, err := dns.Exchange(m, a)
		if err == nil {
			for _, answer := range in.Answer {
				if answer, ok := answer.(*dns.A); ok {
					addr := answer.A.String()
					//log.Debugf("host:%v, addr:%v", host, addr)
					if exists(addrs, addr) == false {
						addrs = append(addrs, addr)
					}
				}
			}
		}
	}
	return
}

//c:\windows\system32\drivers\etc\hosts
var hostFile = "c:\\windows\\system32\\drivers\\etc\\hosts"

func writeHost(hosts map[string][]string) {
	lines := readHost(hosts)
	for k, v := range hosts {
		if len(v) > 0 {
			for _, ip := range v {
				ipLen := 15 - len(ip)
				if ipLen >= 0 {
					lines = append(lines, strings.Repeat(" ", ipLen)+ip+"    "+k)
				}
			}
		}
	}
	f, err := os.Create(hostFile)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range lines {
		fmt.Fprintln(w, v)
	}
	w.Flush()
}

func readHost(hosts map[string][]string) []string {
	//c:\windows\system32\drivers\etc\hosts
	fi, err := os.Open(hostFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	var lines []string
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(a)
		find := false
		for k, v := range hosts {
			if len(v) > 0 && strings.Contains(line, k) {
				find = true
				break
			}
		}
		if find == false {
			lines = append(lines, line)
		}
	}
	return lines
}
