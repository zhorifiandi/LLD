package mvp

import (
	"fmt"

	"github.com/zhorifiandi/parking-lot-lld/domain"
)

type ApplicationInputs struct {
	FloorNumbers             int
	SlotsNumbersForEachFloor int
}

type AssignmentType = [][]domain.Slot

type Application struct {
	Assignment AssignmentType
}

func NewApplication(inputs ApplicationInputs) *Application {
	assignment := AssignmentType{}
	for i := 0; i < inputs.FloorNumbers; i++ {
		assignment = append(assignment, []domain.Slot{})
		for j := 0; j < inputs.SlotsNumbersForEachFloor; j++ {
			assignment[i] = append(assignment[i], domain.Slot{})
		}
	}

	return &Application{
		Assignment: assignment,
	}
}

func (p *Application) AcceptCustomer(vehicleID string) domain.Slot {
	for i, floor := range p.Assignment {
		for j, slot := range floor {
			if slot.VehicleID == "" {
				newSlot := domain.Slot{
					VehicleID: vehicleID,
					FloorID:   i,
					SlotID:    j,
				}
				p.Assignment[i][j] = newSlot
				return newSlot
			}
		}
	}

	return domain.Slot{}
}

func (p *Application) ReleaseCustomer(vehicleID string) (slot domain.Slot) {
	for i, floor := range p.Assignment {
		for j, slot := range floor {
			if slot.VehicleID == vehicleID {
				releasedSlot := p.Assignment[i][j]
				p.Assignment[i][j] = domain.Slot{}
				return releasedSlot
			}
		}
	}

	return domain.Slot{}
}

func (p *Application) PrintAssignment() {
	fmt.Println("Current Assignment:")
	for i, floor := range p.Assignment {
		fmt.Printf("Floor Level %+v: ", i)
		for _, slot := range floor {
			if slot.VehicleID == "" {
				fmt.Printf("xxxxx ")
			} else {
				fmt.Printf("%+v ", slot.VehicleID)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Requirement 2
func (p *Application) AddSlotsOnFloor(floorLevel, slotNums int) {
	for i := 0; i < slotNums; i++ {
		p.Assignment[floorLevel] = append(p.Assignment[floorLevel], domain.Slot{})
	}
}

// Requirement 3
func (p *Application) AddFloor(slotNums int) {
	p.Assignment = append(p.Assignment, []domain.Slot{})
	lastFloorID := len(p.Assignment) - 1
	for i := 0; i < slotNums; i++ {
		p.Assignment[lastFloorID] = append(p.Assignment[lastFloorID], domain.Slot{})
	}
}
