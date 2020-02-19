package main

import "./elevio"
import "fmt"
import "./logmanagement"

func Initialize(numFloors int) {
    elevio.Init("localhost:15657", numFloors)
    for i := 0; i < numFloors; i++ {
        if i != 0 {
            elevio.SetButtonLamp(elevio.BT_HallDown, i, false)
        }
        if i != numFloors {
            elevio.SetButtonLamp(elevio.BT_HallUp, i, false)
        }
        elevio.SetButtonLamp(elevio.BT_Cab, i, false)
    }
}

func main(){

    numFloors := 4

    Initialize(numFloors)
    
    var floor = 0
    var destination = 0

    var d elevio.MotorDirection = elevio.MD_Stop
    elevio.SetMotorDirection(d)
    
    drv_buttons := make(chan elevio.ButtonEvent)
    drv_floors  := make(chan int)
    drv_obstr   := make(chan bool)
    drv_stop    := make(chan bool)
    order       := make(chan logmanagement.Order)
    
    logmanagement.InitializeQueue(logmanagement.PendingOrdersQueue)
    logmanagement.InitializeQueue(logmanagement.ActiveOrdersQueue)

    go elevio.PollButtons(drv_buttons)
    go elevio.PollFloorSensor(drv_floors)
    go elevio.PollObstructionSwitch(drv_obstr)
    go elevio.PollStopButton(drv_stop)
    go logmanagement.CheckForOrders(logmanagement.PendingOrdersQueue, order)
    
    
    for {
        select {
        case a := <- drv_buttons:
            fmt.Printf("%+v\n", a)
            elevio.SetButtonLamp(a.Button, a.Floor, true)
            var newOrder = logmanagement.NewOrder(a.Floor, int(a.Button), 0)
            logmanagement.SetOrder(logmanagement.PendingOrdersQueue, newOrder)

        case a := <- drv_floors:
            fmt.Printf("%+v\n", a)
            floor = a
            if a == numFloors-1 {
                d = elevio.MD_Down
            } else if a == destination {
                d = elevio.MD_Stop
                // Remove order funksjon
            }
            elevio.SetMotorDirection(d)
            
        case a := <- drv_obstr:
            fmt.Printf("%+v\n", a)
            if a {
                elevio.SetMotorDirection(elevio.MD_Stop)
            } else {
                elevio.SetMotorDirection(d)
            }
            
        case a := <- drv_stop:
            fmt.Printf("%+v\n", a)
            for f := 0; f < numFloors; f++ {
                for b := elevio.ButtonType(0); b < 3; b++ {
                    elevio.SetButtonLamp(b, f, false)
                }
            }

        case a := <- order:
            destination = a.Floor
            if floor < a.Floor {
                d = elevio.MD_Up
            } else if floor > a.Floor {
                d = elevio.MD_Down
            }
            elevio.SetMotorDirection(d)
        }
    }    
}
