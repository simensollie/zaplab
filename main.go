//Recieve and print out data from ZapClient(s)
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"
)

//global variabel av newzapstore skal kunne brukes av main og listen

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	netListen, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)
	go listen(netListen)

	for {
		time.Sleep(1 * time.Second)
		//computeviewers
		//print computed viewers
	}
}

func listen(conn *net.UDPConn) {
	var data [256]byte
	for {
		n, _, err := conn.ReadFromUDP(data[0:])
		checkError(err)
		nze := newZapEvent(string(data[0:n]))
		//fmt.Println(nze)
		//inn i StoreZap
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
