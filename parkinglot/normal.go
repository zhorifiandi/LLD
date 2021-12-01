package parkinglot

import (
	"context"

	"github.com/zhorifiandi/parking-lot-lld/domain"
)

type NormalParkingLot struct {
	ParkingSpotManager domain.IParkingSpotManager
	TicketManager      domain.ITicketManager
	PaymentManager     domain.IPaymentManager
}

func (p NormalParkingLot) CheckIn(ctx context.Context) (ticket domain.ParkingTicket, err error) {
	parkingSpot, err := p.ParkingSpotManager.AssignParkingSpot(ctx)
	if err != nil {
		return
	}

	ticket, err = p.TicketManager.IssueTicket(ctx, parkingSpot)
	if err != nil {
		return
	}

	return ticket, nil
}

func (p NormalParkingLot) RequestPayment(ctx context.Context, parkingSpotID string) (payment domain.Payment, err error) {
	ticket, err := p.TicketManager.GetTicketByID(ctx, parkingSpotID)
	if err != nil {
		return
	}

	payment, err = p.PaymentManager.CalculateFee(ctx, ticket)
	if err != nil {
		return
	}

	return
}

func (p NormalParkingLot) CheckOut(ctx context.Context, paymentID string) error {
	payment, err := p.PaymentManager.GetPaymentByID(ctx, paymentID)
	if err != nil {
		return err
	}

	parkingSpot := &(payment.Ticket.AssignedParkingSpot)
	err = p.ParkingSpotManager.ReleaseParkingSpot(ctx, *parkingSpot)
	if err != nil {
		return err
	}

	err = p.TicketManager.CloseTicket(ctx, payment.Ticket)
	if err != nil {
		return err
	}

	return nil
}
