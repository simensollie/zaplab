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
	
	//må lage egen funksjon for denn (c1 og c2)
	for {
		time.Sleep(1 * time.Second)
		//computeviewers
		nrk := zapstore.ComputeViewers("NRK1")
		tv2 := zapstore.ComputeViewers("TV2 Norge")
		//print computed viewers
		fmt.Println("Number of viewers @ NRK1: ", nrk)
		fmt.Println("Number of viewers @ TV2: ", tv2)
	}

	//denne og må i egen funksjon (c3)
	for {
		time.Sleep(5 * time.Second)
		//print ut "the number of entries in the storage
	}
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

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

