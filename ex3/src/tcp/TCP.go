package tcp

import (
	"fmt"
	"net"
	"time"

	"../checkerror"
)

// TCP Sends and receives messages with TCP
func TCP() {
	// Defining IP-adresses and ports
	ServerIP := "10.100.23.147"
	ServerPort := ":33546"
	LocalIP := "10.100.23.223"
	LocalPort := ":20009"

	serverAddr, err := net.ResolveTCPAddr("tcp", ServerIP+ServerPort)
	checkerror.CheckError(err)

	serverConn, err := net.DialTCP("tcp", nil, serverAddr)
	checkerror.CheckError(err)

	// Make the server connect to you at LOCAL_PORT
	connectOrder := "Connect to: " + LocalIP + LocalPort + "\x00"

	localAddr, err := net.ResolveTCPAddr("tcp", LocalIP+LocalPort)
	checkerror.CheckError(err)

	listener, err := net.ListenTCP("tcp", localAddr)
	checkerror.CheckError(err)

	// Send the connection order
	_, err = serverConn.Write([]byte(connectOrder))
	checkerror.CheckError(err)

	// Establish TCP connection
	clientConn, err := listener.AcceptTCP()
	checkerror.CheckError(err)

	for {

		msg := "This is a message from gr9, \x00"
		fmt.Println("Sending message: ", msg)
		clientConn.Write([]byte(msg))

		buffer := make([]byte, 1024)
		clientConn.Read(buffer)
		fmt.Println("Recived message: ", string(buffer))

		time.Sleep(time.Second * 1)
	}

}
