package main

import "./logmanagement"

func main() {
	logmanagement.InitNetwork(15647)

	go logmanagement.UpdateLogFromLocal()
	logmanagement.UpdateLogFromNetwork()
}
