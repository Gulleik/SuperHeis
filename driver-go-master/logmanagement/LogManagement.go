package logmanagement

//import "fmt"

import (
	"fmt"
	"time"

	network "../Network"
)

const numFloors = 4
const numButtons = 3

var id string

/*State enum*/
type State int

const (
	Idle = 0
	Exec = 1
	Lost = 2
)

/*OrderStruct*/
type Order struct {
	Floor      int
	ButtonType int
	Pending    int
	// Timer?
}

/*Elevstruct for keeping info about ther elevs*/
type Elev struct {
	ID    string
	Floor int
	//Lastseen time
	state int
}

/*Log to be sendt over the network*/
type log struct {
	//Orders []Order
	//Elevs []Elev
	Message string
	Iter    int
	//version time
}

/*Declaration of local log*/
var log1 log

/*Broadcast and recieve channel*/
var RcvChannel chan log
var bcastChannel chan log

func NewOrder(floor int, buttonType int, pending int) Order {
	order := Order{Floor: floor, ButtonType: buttonType, Pending: 0}
	return order
}

var PendingOrdersQueue = &[numFloors][numButtons]Order{}
var ActiveOrdersQueue = &[numFloors][numButtons]Order{}

func InitializeQueue(queue *[numFloors][numButtons]Order) {
	for i := 0; i < numFloors; i++ {
		for j := 0; j < numButtons; j++ {
			//queue[i][j] = nil
			queue[i][j].Floor = -1
			queue[i][j].ButtonType = -1
			queue[i][j].Pending = 0
		}
	}
}

func CheckForOrders(queue *[numFloors][numButtons]Order, receiver chan<- Order) { // Velger den første i lista, ikke den eldste ordren
	for {
		time.Sleep(20 * time.Millisecond)
		for i := 0; i < numFloors; i++ {
			for j := 0; j < numButtons; j++ {
				if queue[i][j].Floor != -1 {
					receiver <- queue[i][j]
				}
			}
		}
	}

}

func SetOrder(queue *[numFloors][numButtons]Order, order Order) {
	queue[order.Floor][order.ButtonType] = order
}

/**
 * @brief puts message on bcastChannel
 * @param Message; message to be transmitted
 */
func UpdateLogFromLocal( /*Message log*/ ) {
	log1.Iter = 0
	log1.Message = "lil boy repus"
	for {
		fmt.Printf("updating\n")
		log1.Iter++
		bcastChannel <- log1
		fmt.Printf("ok?\n")
		time.Sleep(1 * time.Second)
	}
}

/**
 * @brief reads message from RcvChannel and does hit width it
 */
func UpdateLogFromNetwork() {
	for {
		a := <-RcvChannel
		fmt.Printf("Received: %#v\n", a)
	}
}

/**
 * @brief initiates channels and creates coroutines for brodcasting and recieving
 * @param port; port to listen and read on
 */
func InitNetwork(port int) {
	RcvChannel = make(chan log)
	bcastChannel = make(chan log)
	go network.BrodcastMessage(port, bcastChannel)
	go network.RecieveMessage(port, RcvChannel)
	fmt.Printf("log initialized\n")
}

/**
 * @brief Set id of elev.
 */ /*
func setElevID() {

	// Our id can be anything. Here we pass it on the command line, using
	//  `go run main.go -id=our_id`
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.Parse()

	// ... or alternatively, we can use the local IP address.
	// (But since we can run multiple programs on the same PC, we also append the
	//  process ID)
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}
}*/
