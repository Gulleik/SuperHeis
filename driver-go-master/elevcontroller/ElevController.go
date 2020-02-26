package elevcontroller

import "fmt"
import "elevio.go"


var _currentFloor int

func MoveToFloor(floor int) {
	_mtx.Lock()
	defer _mtx.Unlock()
	if(_currentFloor == floor){
		return 
	}else if(_currentFloor < floor){

	}else{

	}


	_conn.Write([]byte{3, byte(floor), 0, 0})
}