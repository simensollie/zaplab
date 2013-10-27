package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "2 Arguments required\n")
		os.Exit(1)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1202")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, sub := conn.Write([]byte(os.Args[1]))
	checkError(sub)
	go Printer(conn)
	var cmd string
	fmt.Scanf("%s", &cmd)
	_, err = conn.Write([]byte(cmd))
	checkError(err)
}

func Printer(conn net.Conn){
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
