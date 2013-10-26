package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
//	"zaplab/mapsort"
	"zaplab/topchan"
	"zaplab/zapevent"
	"zaplab/ztorage"
	"runtime/pprof"
	"flag"
)

var zapstore = ztorage.NewZapStore()
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var m  = make(map[string]int)

func main() {
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)

	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	netListen, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)
	go listen(netListen)
	//go chviewers("NRK1")
	//go chviewers("TV2 Norge")
	//go entries(zapstore)
	//go topchan.TopTen(zapstore, m)

	<-c
	memProfile()
}

func listen(conn *net.UDPConn) {
	var data [256]byte
	for {
		n, _, err := conn.ReadFromUDP(data[0:])
		checkError(err)
		nze := *zapevent.NewZapEvent(string(data[0:n]))
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

func entries(zaps *ztorage.Zaps) {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Number of entries in the storage: %d\n", len(*zaps))
	}
}


func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func memProfile(){
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		checkError(err)
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

