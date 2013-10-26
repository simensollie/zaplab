package tcpserver

import (
	"net"
	"fmt"
	"os"
	"regexp"
	"zaplab/topchan"
	"zaplab/ztorage"
)

func ListenTCP(zapstore *ztorage.Zaps, m map[string]int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1202")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn, zapstore, m)
	}
}

func handleClient(conn net.Conn, zapstore *ztorage.Zaps, m map[string]int){
	//var buf [512]byte
	for {
		sub, err := regexp.MatchString("sub", os.Args[1])
		unsub, err := regexp.MatchString("unsub", os.Args[1])
		rate := os.Args[2]

		if sub {
			topchan.TopTen(zapstore, m, rate)
		}else if unsub {
			conn.Close()
			os.Exit()
		}else {
			fmt.Println("You have not entered a valid argument")
		}

		/*n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}*/
	}
}
func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(-1)
	}
}
