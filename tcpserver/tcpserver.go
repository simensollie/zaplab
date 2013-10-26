package tcpserver

import (
	"net"
	"fmt"
)

func listenTCP() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1202")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("dette er en test")
		conn.Close()
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Errror())
		os.Exit(-1)
	}
}
