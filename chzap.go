//Recieve and print out data from ZapClient(s)
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"
)

const timeLayout = "2006/01/02, 15:04:05"

type zapEvent struct {
	datetime time.Time
	ip string
	fromCh string
	toCh string
}

func newZapEvent(data string) *zapEvent {
	s := strings.Split(data, ", ")
	t, _ := time.Parse(timeLayout, data[0:20])
	if (len(s) < 5) {
		return &zapEvent{t, s[2], s[3], "TASTATUR"}
	} else {
		return &zapEvent{t, s[2], s[3], s[4]}
	}
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	netListen, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)
	listen(netListen)
}

func listen(conn *net.UDPConn) {
	var data [256]byte
	for {
		n, _, err := conn.ReadFromUDP(data[0:])
		checkError(err)
		nze := newZapEvent(string(data[0:n]))
		fmt.Println(nze)

//		fmt.Println(string(data[0:n])
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func (ze *zapEvent) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", ze.datetime, ze.ip, ze.fromCh, ze.toCh)
}

/*func (ze *zapEvent) Duration(provided ChZap) time.Duration {
	
	return time
}*/
