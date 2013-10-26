package tcpclient

import (
	"net"
	"fmt"
)

func something() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "2 Arguments required\n")
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp" service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
