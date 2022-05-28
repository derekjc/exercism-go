package speed

type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
        battery: 100,
        speed: speed,
        batteryDrain: batteryDrain,
        distance: 0,
    }
	return car
}

type Track struct{
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
        distance: distance,
    }
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery > car.batteryDrain {
		car.battery = car.battery - car.batteryDrain
    	car.distance = car.distance + car.speed
    }
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distance_can_go := (float64(car.battery) * float64(car.speed) / float64(car.batteryDrain))
	if float64(track.distance) <= distance_can_go {
        return true
    } else {
    	return false
    }
}
