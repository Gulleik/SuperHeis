package logmanagement

import (
	"time"
)

const numFloors = 4
const numButtons = 3

type Order struct {
	Floor      int //Remove
	ButtonType int //Remove
	Active     int
	// Timer?
}

func NewOrder(floor int, buttonType int, active int) Order {
	order := Order{Floor: floor, ButtonType: buttonType, Active: active}
	return order
}

var OrderQueue = &[numFloors][numButtons]Order{}

func InitializeQueue(queue *[numFloors][numButtons]Order) {
	for i := 0; i < numFloors; i++ {
		for j := 0; j < numButtons; j++ {
			//queue[i][j] = nil
			queue[i][j].Floor = i
			queue[i][j].ButtonType = j
			queue[i][j].Active = -1
		}
	}
}

/*func CheckForOrders(queue *[numFloors][numButtons]Order, receiver chan<- Order) { // Velger den første i lista, ikke den eldste ordren
	for { // Legg i orderHandler?
		time.Sleep(20 * time.Millisecond)
		for i := 0; i < numFloors; i++ {
			for j := 0; j < numButtons; j++ {
				if queue[i][j].Active == 0 {
					fmt.Printf("%+v\n", queue[i][j])
					receiver <- queue[i][j]
				}
			}
		}
	}
}*/

func CheckForOrder() Order { // Velger den første i lista, ikke den eldste ordren
	time.Sleep(20 * time.Millisecond)
	for i := 0; i < numFloors; i++ {
		for j := 0; j < numButtons; j++ {
			if OrderQueue[i][j].Active == 0 {
				return OrderQueue[i][j]
			}
		}
	}
	order := Order{Floor: -1, ButtonType: -1, Active: -1}
	return order
}

// UpdateOrderQueue updates the order queue
func UpdateOrderQueue(floor int, button int, active int) {
	OrderQueue[floor][button].Active = active
}

// GetActiveOrder returns the first found active order
func GetActiveOrder() Order {
	for i := 0; i < numFloors; i++ {
		for j := 0; j < numButtons; j++ {
			if OrderQueue[i][j].Active == 1 {
				return OrderQueue[i][j]
			}
		}
	}
	return Order{Floor: -1, ButtonType: -1, Active: -1}
}

func GetOrder(floor int, buttonType int) Order {
	return OrderQueue[floor][buttonType]
}
