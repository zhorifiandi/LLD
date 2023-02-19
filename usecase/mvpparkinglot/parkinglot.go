package mvpparkinglot

import "log"

type Slot struct {
	AssignedTicketID string
	FloorIndex       int
	SlotIndex        int
}

func NewSlot(
	FloorIndex int,
	SlotIndex int,
) Slot {
	return Slot{
		FloorIndex: FloorIndex,
		SlotIndex:  SlotIndex,
	}
}

type Floor struct {
	Slots []Slot
}

func NewFloor(FloorIndex int, SlotsNumbersForEachFloor int) Floor {
	var slots []Slot
	for i := 0; i < SlotsNumbersForEachFloor; i++ {
		slot := NewSlot(
			FloorIndex,
			i,
		)
		slots = append(slots, slot)
	}

	return Floor{
		Slots: slots,
	}
}

type Assignment struct {
	Floors []Floor
}

func (a *Assignment) Assign() {

}

func NewAssignment(
	FloorNumbers int,
	SlotsNumbersForEachFloor int,
) Assignment {
	var floors []Floor
	for i := 0; i < FloorNumbers; i++ {
		log.Printf("run %+v\n", i)
		floor := NewFloor(i, SlotsNumbersForEachFloor)
		floors = append(floors, floor)
	}

	return Assignment{
		Floors: floors,
	}
}

type Application struct {
	Assignment Assignment
}

type ApplicationInputs struct {
	FloorNumbers             int
	SlotsNumbersForEachFloor int
}

func NewApplication(inputs ApplicationInputs) *Application {
	assignment := NewAssignment(
		inputs.FloorNumbers,
		inputs.SlotsNumbersForEachFloor,
	)

	log.Printf("Assignment: %+v \n", assignment)

	return &Application{
		Assignment: assignment,
	}
}

func (p *Application) AcceptCustomer() {
	// Assignment.Assign()
	floors := p.Assignment.Floors
	for i, floor := range floors {
		log.Printf("%+v %+v\n", i, floor)

	}
}

func (p *Application) ReleaseCustomer() {
}
