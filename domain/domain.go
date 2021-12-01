package domain

import (
	"time"
)

type IParkingLotApplication interface {
	func CheckIn() error;
	func CheckOut() error;
}

// Service
type IParkingSpotManager interface {
	func AssignParkingSpot() error;
	func ReleaseParkingSpot() error;
}

type ITicketManager interface {
	func IssueTicket() error;
	func CloseTicket() error;
}

type IPaymentManager interface {
	func CalculateFee() error;
	func ReceivePayment() error;
}

// Entities
type ParkingTicket struct {
	ID string;
	CheckinTimestamp time.Time;
	CheckoutTimestamp time.Time;
	Entrance ParkingEntrance;
	Exit ParkingExit;
	AssignedParkingSpot ParkingSpot;
}

type ParkingEntrance struct {
	ID string
	Location Location
}

type ParkingExit struct {
	ID string
	Location Location
}

type ParkingSpot struct {
	ID string
	Location Location
	VehicleType string
}

type Payment struct {
	ID string
	Ticket ParkingTicket
	Fee float32;
}


// Value Object
type Location struct {
	Floor int;
	Lattitude int;
	Longitude int;
}