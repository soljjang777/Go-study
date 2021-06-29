package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C,"
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P,"
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C,"
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P,"
			outPlaneChan <- plane
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	carchan1 := make(chan Car)
	carchan2 := make(chan Car)
	carchan3 := make(chan Car)

	planechan1 := make(chan Plane)
	planechan2 := make(chan Plane)
	planechan3 := make(chan Plane)

	go StartCarWork(carchan1)
	go StartPlaneWork(planechan1)
	go MakeTire(carchan1, planechan1, carchan2, planechan2)
	go MakeEngine(carchan2, planechan2, carchan3, planechan3)

	for {
		select {
		case result := <-carchan3:
			fmt.Println(result.val)
		case result := <-planechan3:
			fmt.Println(result.val)
		}
	}
}
