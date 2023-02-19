package domain

import (
	"fmt"
	"time"
)

type Slot struct {
	VehicleID   string
	FloorID     int
	SlotID      int
	CheckInTime time.Time
}

type ParkingFee struct {
	TotalHour int
	TotalFee  int
}

func (fee *ParkingFee) Print() {
	fmt.Printf("Total Fee: %+v (for %+v hours) \n", fee.TotalFee, fee.TotalHour)
}
