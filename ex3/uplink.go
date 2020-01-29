package main

/*
import (
	"fmt"
	"net"
	"time"
)

// Receiver
func main() {
	//Basic variables
	port := ":20009"
	protocol := "udp"
	serverIP := "10.100.23.147:20009"

	done := make(chan int)

	go listen(port, protocol)
	go send(serverIP, protocol)

	<-done
}

func send(serverIP string, protocol string) {

	Connection, err := net.Dial("udp", serverIP)
	checkError(err)

	defer Connection.Close()
	for {
		msg := []byte("Hei hei")
		time.Sleep(time.Second * 1)
		_, err := Connection.Write(msg)
		checkError(err)
	}
}

func listen(port string, protocol string) {

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	checkError(err)
	udpConn, err := net.ListenUDP(protocol, udpAddr)
	checkError(err)
	for {
		buffer := make([]byte, 1024)
		n, err := udpConn.Read(buffer)
		checkError(err)
		fmt.Println(string(buffer[0:n]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
*/
