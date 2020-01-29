package main

func main() {

	port := ":20009"
	protocol := "udp"
	serverIP := "10.100.23.147:20009"

	done := make(chan int)

	go listen(port, protocol)
	go send(serverIP, protocol)

	<-done
}
