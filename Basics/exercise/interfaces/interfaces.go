//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type vehicalDirector interface {
	drive()
}

type Truck string
type Car string
type Motorcycle string

func (t Truck) drive() {
	fmt.Println("Driving to large lifts")
}
func (c Car) drive() {
	fmt.Println("Driving to standard lifts")
}
func (m Motorcycle) drive() {
	fmt.Println("Driving to small lifts")
}

func driveVehicle(vehicals []vehicalDirector) {
	for i := 0; i < len(vehicals); i++ {
		vehicle := vehicals[i]
		fmt.Printf("--Vehicle: %v--\n", vehicle)
		vehicle.drive()
	}
	fmt.Println()
}

func main() {
	i := []vehicalDirector{Truck("Road devour"), Car("mercedes"), Motorcycle("kavasaki")}
	driveVehicle(i)
}
