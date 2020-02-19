package fsm

import "../elevio"

type State string

const(
	INIT = "INIT"
	IDLE = "IDLE"
	EXECUTE = "EXECUTE"
	RESET = "RESET"
)

func RunElevator() {
	state := INIT
	floor := elevio.GetFloor()
	for {
		switch state{
		case IDLE:
			//test

		case EXECUTE:
			//elevio.SetMotorDirection(1)

		case RESET:
			//reset elevator


		}
	}

}

