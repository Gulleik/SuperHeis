package logmanagement

import "time"
//import "fmt"

const numFloors = 4
const numButtons = 3

type Order struct {
	Floor int
	ButtonType int
	Pending int
	// Timer?
}

func NewOrder(floor int, buttonType int, pending int) Order{
	order := Order{Floor: floor, ButtonType: buttonType, Pending: 0}
	return order
}

var PendingOrdersQueue = &[numFloors][numButtons]Order{}
var ActiveOrdersQueue = &[numFloors][numButtons]Order{}

func InitializeQueue(queue *[numFloors][numButtons]Order){
	for i := 0; i < numFloors; i++ {
		for j := 0; j < numButtons; j++ {
			//queue[i][j] = nil
			queue[i][j].Floor = -1
			queue[i][j].ButtonType = -1
			queue[i][j].Pending = 0
		}
	}
}

func CheckForOrders(queue *[numFloors][numButtons]Order, receiver chan<- Order){ // Velger den første i lista, ikke den eldste ordren
	for {
		time.Sleep(20 * time.Millisecond)
		for i := 0; i < numFloors; i++ {	
			for j := 0; j < numButtons; j++ {
				if queue[i][j].Floor != -1{
					receiver <- queue[i][j]
				}
			}
		}
	}
	
}

func SetOrder(queue *[numFloors][numButtons]Order, order Order){
	queue[order.Floor][order.ButtonType] = order
}
