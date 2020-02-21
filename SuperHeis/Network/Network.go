package network

import (
	"fmt"
	"net"
	"time"
	"strings"
	"../conn"
)

var localIP1 string
var port string
var protocol string
var serverIP string

type PeerUpdate struct {
	Peers []string
	New   string
	Lost  []string
}

const interval = 15 * time.Millisecond
const timeout = 50 * time.Millisecond


func InitNetwork() {

	port := ":20009"
	protocol := "udp"
	serverIP := "10.100.23.147:20009"

	done := make(chan int)

}

func BrodcastMessage(MSG Message,String ServerIP) {}

func RecieveMessage(String Port) {}

func send(port int, id string, transmitEnable <-chan bool) {

	conn := conn.DialBroadcastUDP(port)
	addr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("255.255.255.255:%d", port))

	enable := true
	for {
		select {
		case enable = <-transmitEnable:
		case <-time.After(interval):
		}
		if enable {
			conn.WriteTo([]byte(id), addr)
		}
	}
}

func updatePeers(port int, peerUpdateCh chan<- PeerUpdate) []byte{

	var buf [1024]byte
	var p PeerUpdate
	lastSeen := make(map[string]time.Time)

	conn := conn.DialBroadcastUDP(port)

	for {
		updated := false

		conn.SetReadDeadline(time.Now().Add(interval))
		n, _, _ := conn.ReadFrom(buf[0:])

		id := string(buf[:n])

		// Adding new connection
		p.New = ""
		if id != "" {
			if _, idExists := lastSeen[id]; !idExists {
				p.New = id
				updated = true
			}

			lastSeen[id] = time.Now()
		}

		// Removing dead connection
		p.Lost = make([]string, 0)
		for k, v := range lastSeen {
			if time.Now().Sub(v) > timeout {
				updated = true
				p.Lost = append(p.Lost, k)
				delete(lastSeen, k)
			}
		}

		// Sending update
		if updated {
			p.Peers = make([]string, 0, len(lastSeen))

			for k, _ := range lastSeen {
				p.Peers = append(p.Peers, k)
			}

			sort.Strings(p.Peers)
			sort.Strings(p.Lost)
			peerUpdateCh <- p
		}
	}
}

func createJSONFromMessage(){}

func createMessageFromJSON(){}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func localIP() (string, error) {
	if localIP1 == "" {
		conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
		if err != nil {
			return "", err
		}
		defer conn.Close()
		localIP1 = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return localIP1, nil
}
