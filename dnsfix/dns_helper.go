package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

type dnsHeader struct {
	Id                                 uint16
	Bits                               uint16
	Qdcount, Ancount, Nscount, Arcount uint16
}

func (header *dnsHeader) SetFlag(QR uint16, OperationCode uint16, AuthoritativeAnswer uint16, Truncation uint16, RecursionDesired uint16, RecursionAvailable uint16, ResponseCode uint16) {
	header.Bits = QR<<15 + OperationCode<<11 + AuthoritativeAnswer<<10 + Truncation<<9 + RecursionDesired<<8 + RecursionAvailable<<7 + ResponseCode
}

type dnsQuery struct {
	QuestionType  uint16
	QuestionClass uint16
}

func ParseDomainName(domain string) []byte {
	var (
		buffer   bytes.Buffer
		segments []string = strings.Split(domain, ".")
	)
	for _, seg := range segments {
		binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	binary.Write(&buffer, binary.BigEndian, byte(0x00))

	return buffer.Bytes()
}
func Send(dnsServer, domain string) ([]byte, int, time.Duration) {
	requestHeader := dnsHeader{
		Id:      0x0010,
		Qdcount: 1,
		Ancount: 0,
		Nscount: 0,
		Arcount: 0,
	}
	requestHeader.SetFlag(0, 0, 0, 0, 1, 0, 0)

	requestQuery := dnsQuery{
		QuestionType:  1,
		QuestionClass: 1,
	}

	var (
		conn   net.Conn
		err    error
		buffer bytes.Buffer
	)

	if conn, err = net.Dial("udp", dnsServer); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	defer conn.Close()

	binary.Write(&buffer, binary.BigEndian, requestHeader)
	binary.Write(&buffer, binary.BigEndian, ParseDomainName(domain))
	binary.Write(&buffer, binary.BigEndian, requestQuery)

	buf := make([]byte, 2048)
	t1 := time.Now()
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	length, err := conn.Read(buf)
	t := time.Now().Sub(t1)
	return buf[0:length], length, t
}

//func main() {
//	remsg, n, _ := Send("114.114.114.114:53", "www.baidu.com")
//	fmt.Println(remsg, n)
//}