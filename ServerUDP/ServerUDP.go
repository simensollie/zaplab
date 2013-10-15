//Recieve and print out data from ZapClient(s)
package main

import (
	"fmt"
	"net"
	"os"
)

type zapEvent struct {
	date string
	time string
	ip string
	fromCh string
	toCh string
}

func newZapEvent(data string) *zapEvent {
	//z := zapEvent(string(data))
	return nil
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
		fmt.Println(string(data[0:n]))
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func (ze *zapEvent) String() string {
	
	return "Ikke implementert"
}
