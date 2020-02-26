package fsm

import (
	"fmt"

	"../elevcontroller"
	"../elevio"
	"../logmanagement"
	"../orderhandler"
)

type State string

const (
	INIT    = "INIT"
	IDLE    = "IDLE"
	EXECUTE = "EXECUTE"
	RESET   = "RESET"
)

func Initialize(numFloors int) {
	elevio.Init("localhost:15657", numFloors)
	elevcontroller.InitializeLights(numFloors)
	elevcontroller.InitializeElevator()
	elevio.SetFloorIndicator(0) // Evt fix løpende update senere
	logmanagement.InitializeQueue(logmanagement.OrderQueue)
}

func RunElevator() {
	destination := -1
	// InitElev()
	dir := 0   // declared to make code run, prob might delete later
	floor := 0 // same
	state := IDLE

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)
	//order := make(chan logmanagement.Order)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)
	//go logmanagement.CheckForOrders(logmanagement.OrderQueue, order)

	for {
		switch state {
		case IDLE:
			select {
			case a := <-drv_buttons:
				fmt.Println("in IDLE: buttons")
				elevio.SetButtonLamp(a.Button, a.Floor, true)
				newOrder := logmanagement.NewOrder(a.Floor, int(a.Button), 0)
				logmanagement.UpdateOrderQueue(newOrder.Floor, newOrder.ButtonType, 0)

			default:
				newOrder := logmanagement.CheckForOrder()
				if newOrder.Active == 0 {
					fmt.Println("Order found")
					logmanagement.UpdateOrderQueue(newOrder.Floor, newOrder.ButtonType, 1)
					//fmt.Printf("%+v\n", logmanagement.OrderQueue)
					destination = orderhandler.GetDestination()
					dir = orderhandler.GetMotorDirection(floor, destination)
					//fmt.Printf("destination in execute: %+v\n", destination)
					state = EXECUTE

				}
			}

		case EXECUTE:
			elevio.SetMotorDirection(elevio.MotorDirection(dir))
			select {
			case a := <-drv_buttons:
				fmt.Println("in EXECUTE: buttons")
				elevio.SetButtonLamp(a.Button, a.Floor, true)
				newOrder := logmanagement.NewOrder(a.Floor, int(a.Button), 0)
				logmanagement.UpdateOrderQueue(newOrder.Floor, newOrder.ButtonType, 0)

			case floorReached := <-drv_floors:
				floor = floorReached
				elevio.SetFloorIndicator(floor)
				if orderhandler.ShouldElevatorStop(floor, destination) {
					elevcontroller.ElevStopAtFloor(floor)
					orderhandler.ClearOrdersAtFloor(floor)
					//fmt.Printf("%+v\n", logmanagement.OrderQueue)
					dir = orderhandler.GetMotorDirection(floor, destination)
					elevio.SetMotorDirection(elevio.MotorDirection(dir))
					if dir == 0 {
						destination = -1
						state = IDLE
					}
				}
			default:
				fmt.Println("In EXE: default")
				if dir == 0 {
					elevcontroller.OpenCloseDoor(3)
					orderhandler.ClearOrdersAtFloor(floor)
					state = IDLE
				}
			}

		case RESET:
			//reset elevator

		}

	}

}
