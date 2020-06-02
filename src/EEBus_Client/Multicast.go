package EEBus_Client

import (
	"encoding/hex"
	"log"
	"net"
	"time"
)

var srcString string

const (
	srvAddr         = "224.0.0.1:9999"
	maxDatagramSize = 8192
)

func Multicast(choice string) {
	go ping(srvAddr, choice)
	serveMulticastUDP(srvAddr, msgHandler)
}

func ping(a string, choice string) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for {
		c.Write([]byte("hello, " + choice + "\n"))
		time.Sleep(1 * time.Second)
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
	srcString = src.IP.String()
}

func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxDatagramSize)

	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		h(src, n, b)

	}
}

func GetSrcString() string {
	return srcString
}
