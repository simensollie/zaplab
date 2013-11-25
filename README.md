zaplab
======
School project by Kristine Svaboe & Simen Sollie

Lines of code that need to be commented in to work with the different assignments:
task: a
	go listen(netListen)
task c1:
	go chViewers("NRK1")
task c2:
	go chViewers("TV2 Norge")
task c3:
	go entries()
task f:
	"zaplab/tcpserver"
	go tcpserver.ListenTCP(zapstore, m)
	/* How to run:
		go run tcpclient/tcpclient.go sub 	//subscribe to service
		unsub					//unsubscribe and close connection
	*/
task g:
	go listen(netListen)
	ztorage.ComputeDuration(nze) //in method listen()
