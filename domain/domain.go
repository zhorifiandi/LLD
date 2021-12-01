package domain

import (
	"time"
)

type IParkingLotApplication interface {
	CheckIn() error
	CheckOut() error
}

// Service
type IParkingSpotManager interface {
	AssignParkingSpot() error
	ReleaseParkingSpot() error
}

type ITicketManager interface {
	IssueTicket() error
	CloseTicket() error
}

type IPaymentManager interface {
	CalculateFee() error
	ReceivePayment() error
}

// Entities
type ParkingTicket struct {
	ID                  string
	CheckinTimestamp    time.Time
	CheckoutTimestamp   time.Time
	Entrance            ParkingEntrance
	Exit                ParkingExit
	AssignedParkingSpot ParkingSpot
}

type ParkingEntrance struct {
	ID       string
	Location Location
}

type ParkingExit struct {
	ID       string
	Location Location
}

type ParkingSpot struct {
	ID          string
	Location    Location
	VehicleType string
}

type Payment struct {
	ID     string
	Ticket ParkingTicket
	Fee    float32
}

// Value Object
type Location struct {
	Floor     int
	Lattitude int
	Longitude int
}
