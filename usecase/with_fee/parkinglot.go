package withfee

import (
	"log"
	"time"

	"github.com/zhorifiandi/parking-lot-lld/domain"
	"github.com/zhorifiandi/parking-lot-lld/usecase/mvp"
)

type Application struct {
	AssignmentApp     mvp.Application
	TicketList        map[string]domain.Ticket
	ParkingFeePerHour int
}

func NewApplication(
	assignmentAppInput mvp.ApplicationInputs,
	parkingFeePerHour int,
) *Application {
	assignmentApp := mvp.NewApplication(assignmentAppInput)

	return &Application{
		AssignmentApp:     *assignmentApp,
		TicketList:        map[string]domain.Ticket{},
		ParkingFeePerHour: parkingFeePerHour,
	}
}

func (p *Application) AcceptCustomer(vehicleID string) domain.Ticket {
	slot := p.AssignmentApp.AcceptCustomer(vehicleID)
	ticket := domain.Ticket{
		Slot:        slot,
		CheckInTime: time.Now(),
	}

	p.TicketList[slot.VehicleID] = ticket
	return ticket
}

func (p *Application) ReleaseCustomer(vehicleID string) (ticket domain.Ticket) {
	slot := p.AssignmentApp.ReleaseCustomer(vehicleID)
	ticket = p.TicketList[slot.VehicleID]

	delete(p.TicketList, slot.VehicleID)
	return ticket
}

func (p *Application) PrintAssignment() {
	p.AssignmentApp.PrintAssignment()
}

func (p *Application) AddSlotsOnFloor(floorLevel, slotNums int) {
	p.AssignmentApp.AddSlotsOnFloor(floorLevel, slotNums)
}

func (p *Application) AddFloor(slotNums int) {
	p.AssignmentApp.AddFloor(slotNums)
}

func (p *Application) HandleCustomerExit(vehicleID string) (fee domain.ParkingFee) {
	slot := p.AssignmentApp.ReleaseCustomer(vehicleID)
	ticket := p.TicketList[slot.VehicleID]
	log.Println("TicketList", p.TicketList)
	log.Println("ticket", ticket)

	if slot.VehicleID != "" {
		// Assumed
		elapsedHour := int((time.Since(ticket.CheckInTime)).Seconds())
		log.Println("elapsedHour", elapsedHour)
		fee = domain.ParkingFee{
			TotalHour: elapsedHour,
			TotalFee:  elapsedHour * p.ParkingFeePerHour,
		}
	}

	return
}
