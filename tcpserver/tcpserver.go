package tcpserver

import (
	"net"
	"fmt"
	"os"
	"time"
	//"regexp"
	//"time"
	"zaplab/topchan"
	"zaplab/ztorage"
	//"strings"
)

var stop bool

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
		conn.Close()
	}
}

func handleClient(conn net.Conn, zapstore *ztorage.Zaps, m map[string]int){
	var buf [512]byte
	for {
		/*sub, err := regexp.MatchString("sub", os.Args[1])
		checkError(err)
		unsub, err := regexp.MatchString("unsub", os.Args[1])
		checkError(err)
		refreshrate := os.Args[2] + "s"
		rate, err :=  time.ParseDuration(refreshrate)
		checkError(err)

		if sub {
			topchan.TopTen(zapstore, m, rate)
		}else if unsub {
			conn.Close()
			os.Exit(0)
		}else {
			fmt.Println("You have not entered a valid argument")
		}*/

		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		str := string(buf[0:n])
		if str == "unsub" {
			stop = true
		}

		if str == "sub" {
			fmt.Println(str)
			go subscription(conn, 1, zapstore, m)
		}
	}
}

func subscription(conn net.Conn, sleepTime int, zapstore *ztorage.Zaps, m map[string]int) {
	for {
		time.Sleep(1*time.Second)
		if stop {
			conn.Close()
		}
		_, err2 := conn.Write([]byte(topchan.TopTen(zapstore, m)))
		if err2 != nil {
			return
		}
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(-1)
	}
}
