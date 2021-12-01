package domain

import (
	"context"
	"time"
)

type IParkingLotApplication interface {
	CheckIn(context.Context) (ParkingTicket, error)
	RequestPayment(context.Context) (Payment, error)
	CheckOut(context.Context, string) error
}

// Service
type IParkingSpotManager interface {
	AssignParkingSpot(context.Context) (ParkingSpot, error)
	ReleaseParkingSpot(context.Context, ParkingSpot) error
}

type ITicketManager interface {
	IssueTicket(context.Context, ParkingSpot) (ParkingTicket, error)
	CloseTicket(context.Context, ParkingTicket) error
	GetTicketByID(context.Context, string) (ParkingTicket, error)
}

type IPaymentManager interface {
	CalculateFee(context.Context, ParkingTicket) (Payment, error)
	GetPaymentByID(context.Context, string) (Payment, error)
	ReceivePayment(context.Context, Payment) error
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
