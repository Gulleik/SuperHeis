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
	ServerIP := "10.100.23.147"
	ServerPort := ":33546"
	LocalIP := "10.100.23.223"
	LocalPort := ":20009"

	serverAddr, err := net.ResolveTCPAddr("tcp", ServerIP+ServerPort)
	checkError(err)
	//fmt.Println(serverAddr)

	localAddr, err := net.ResolveTCPAddr("tcp", LocalIP+LocalPort)
	checkError(err)
	//fmt.Println(localAddr)

	serverConn, err := net.DialTCP("tcp", localAddr, serverAddr)
	//serverConn, err := net.Dial("tcp", "10.100.23.242:33546")
	checkError(err)
	//fmt.Println(serverConn)

	// Make the server connect to you at LOCAL_PORT
	connectOrder := "Connect to: " + LocalIP + LocalPort + "\x00"

	listener, err := net.ListenTCP("tcp", nil)
	//fmt.Println("Listen: ", listener)
	checkError(err)

	// Send the connection order
	_, err = serverConn.Write([]byte(connectOrder))
	//fmt.Println("tiss")
	checkError(err)

	// Establish TCP connection
	clientConn, err := listener.AcceptTCP()
	checkError(err)
	fmt.Println(clientConn)

	for {

		msg := "This is a message from gr9, \x00"
		fmt.Println("Sending message: ", msg, "\n")
		_, err = clientConn.Write([]byte(msg))
		checkError(err)

		buffer := make([]byte, 1024)
		clientConn.Read(buffer)
		fmt.Println("Recived message: ", string(buffer), "\n")

		time.Sleep(time.Second * 1)
	}

}
