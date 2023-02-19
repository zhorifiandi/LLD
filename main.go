package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zhorifiandi/parking-lot-lld/domain"
	"github.com/zhorifiandi/parking-lot-lld/usecase/mvp"
	"github.com/zhorifiandi/parking-lot-lld/usecase/sold"
)

type IParkingLotApplication interface {
	// Customer side
	PrintAssignment()
	AcceptCustomer(vehicleID string) domain.Slot
	ReleaseCustomer(vehicleID string) domain.Slot

	// Requirement 2
	AddSlotsOnFloor(int, int)

	// Requirement 3
	AddFloor(slotNums int)

	// Requirement 4
	SetParkingFeePerHour(fee int)
	HandleCustomerExit(vehicleID string) domain.ParkingFee
}

func main() {
	var app IParkingLotApplication = mvp.NewApplication(
		mvp.ApplicationInputs{
			FloorNumbers:             2,
			SlotsNumbersForEachFloor: 3,
		},
	)
	log.Printf("App is running..... \n")

	// Requirement #1
	// A Company has a Parking Building, with entrance and exit gate placed side by side
	// The Parking Building has 2 Floors with 3 Vehicle slot for each flors
	// They want to have a system to handle ticket entrance and ticket exit
	// Customer will be assigned to a nearest available slot from the entrance gate
	// Each customer will be identified by its Vehicle ID
	// A slot become unavailable for customer if it's assigned
	// A slot become available if previous customer has leave with their vehicle
	// Admin must be able to see current slot assignment visually

	_ = app.AcceptCustomer("D1234")
	_ = app.AcceptCustomer("D1235")
	_ = app.AcceptCustomer("E1236")
	_ = app.AcceptCustomer("E1237")
	_ = app.AcceptCustomer("F1238")

	app.PrintAssignment()

	_ = app.ReleaseCustomer("D1235")
	_ = app.ReleaseCustomer("E1237")

	app.PrintAssignment()

	_ = app.AcceptCustomer("G1239")
	_ = app.AcceptCustomer("G1249")
	_ = app.AcceptCustomer("G1259")
	_ = app.AcceptCustomer("G1269")

	app.PrintAssignment()

	// Requirement 2:
	// A company has done minimal renovattion on the building by adding 3 new slots on floor level 1
	// They want to reflect this change on the systems
	fmt.Println(">>>>>>>>>>>>>")
	fmt.Println("Requirement 2")
	fmt.Println(">>>>>>>>>>>>>")
	time.Sleep(1 * time.Second)
	floorLevel := 1
	slotNums := 3
	app.AddSlotsOnFloor(floorLevel, slotNums)
	_ = app.AcceptCustomer("H1279")
	_ = app.AcceptCustomer("H1289")
	_ = app.AcceptCustomer("H1299")
	app.PrintAssignment()

	// Requirement 3:
	// A company has done a quite renovation on the building by adding 2 floors with 7 slots and 6 slots each
	// They want to reflect this change on the systems
	fmt.Println(">>>>>>>>>>>>>")
	fmt.Println("Requirement 3")
	fmt.Println(">>>>>>>>>>>>>")
	time.Sleep(1 * time.Second)
	slotNums = 7
	app.AddFloor(slotNums)

	slotNums = 6
	app.AddFloor(slotNums)

	_ = app.AcceptCustomer("I1300")
	app.PrintAssignment()

	// Requirement #4:
	// The company want to start automate the Parking Fee Calculation
	// Parking Fee per hour is 2000
	// At the time customer exit, systems needs to show
	// - How many time has elapsed (1 sec ~= 1 hour)
	// - Total Fee
	fmt.Println(">>>>>>>>>>>>>")
	fmt.Println("Requirement 4")
	fmt.Println(">>>>>>>>>>>>>")
	time.Sleep(1 * time.Second)
	app.SetParkingFeePerHour(2000)
	fee := app.HandleCustomerExit("D1234")
	fee.Print()

	app.PrintAssignment()

	// Requirement #5:
	// The company decides to sell the software to other company (Company B)
	// Company B building specs:
	// Floor 0: 6 slots
	// Floor 1: 3 slots
	// Floor 2: 4 slots
	// Floor 3: 5 slots
	fmt.Println(">>>>>>>>>>>>>")
	fmt.Println("Requirement 5")
	fmt.Println(">>>>>>>>>>>>>")

	// Clue: Proxy Design Pattern
	soldApp := sold.NewApplication([]int{
		5, 6, 7, 8, 5,
	})
	soldApp.BaseApplication.PrintAssignment()
}
