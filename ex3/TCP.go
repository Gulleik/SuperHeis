package main

import (
	"fmt"
	"net"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Defining IP-adresses and ports
	ServerIP := "10.100.23.242"
	ServerPort := ":33546"
	LocalIP := "10.100.23.223"
	LocalPort := ":20009"

	serverAddr, err := net.ResolveTCPAddr("tcp", ServerIP+ServerPort)
	checkError(err)

	serverConn, err := net.DialTCP("tcp", nil, serverAddr)
	checkError(err)

	// Make the server connect to you at LOCAL_PORT
	connectOrder := "Connect to: " + LocalIP + LocalPort + "\x00"

	localAddr, err := net.ResolveTCPAddr("tcp", LocalIP+LocalPort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", localAddr)
	checkError(err)

	// Send the connection order
	_, err = serverConn.Write([]byte(connectOrder))
	checkError(err)

	// Establish TCP connection
	clientConn, err := listener.AcceptTCP()
	checkError(err)

	for {

		msg := "This is a message from gr9, \x00"
		fmt.Println("Sending message: ", msg, "\n")
		clientConn.Write([]byte(msg))

		buffer := make([]byte, 1024)
		clientConn.Read(buffer)
		fmt.Println("Recived message: ", string(buffer), "\n")

		time.Sleep(time.Second * 1)
	}

}
