//Recieve and print out data from ZapClient(s)
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"zaplab/zapevent"
	"zaplab/ztorage"
)

//global variabel av newzapstore skal kunne brukes av main og listen
var zapstore = ztorage.NewZapStore()

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	netListen, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)
	go listen(netListen)
//	go chviewers("NRK1")
//	go chviewers("TV2 Norge")
	entries()
}

func listen(conn *net.UDPConn) {
	var data [256]byte
	for {
		n, _, err := conn.ReadFromUDP(data[0:])
		checkError(err)
		nze := *zapevent.NewZapEvent(string(data[0:n]))
//		fmt.Println(nze)
		//inn i StoreZap
		zapstore.StoreZap(nze)
	}
}

func chviewers(ch string) {
	for {
		time.Sleep(1 * time.Second)
		viewers := zapstore.ComputeViewers(ch)
		fmt.Printf("Number of viewers @ %s: %d\n", ch, viewers)
	}
}

func entries() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Number of entries in the storage: %d\n", len(zapstore))
	}
}


func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

