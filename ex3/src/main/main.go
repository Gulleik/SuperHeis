package main

import "../tcp"

func main() {

	done := make(chan int)

	//go listen(port, protocol)
	//go send(serverIP, protocol)
	go tcp.TCP()

	<-done
}
