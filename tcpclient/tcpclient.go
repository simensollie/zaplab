package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "3 Arguments required\n")
		os.Exit(1)
	}
	//service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1202")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, sub = conn.Write([]byte(os.Args[1]))
	checkError(sub)
	_, rate = conn.Write([]byte(os.Args[2]))
	checkError(rate)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
