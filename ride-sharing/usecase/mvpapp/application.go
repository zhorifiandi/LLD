package mvpapp

import (
	"fmt"

	"log"

	"github.com/google/uuid"
	"github.com/zhorifiandi/LLD/ride-sharing/domain"
)

type ApplicationInputs struct{}

func NewApplication(inputs ApplicationInputs) (app *Application) {
	return &Application{
		AvailableRides: map[string]domain.Ride{},
		Drivers:        map[string]domain.Driver{},
		Riders:         map[string]domain.Rider{},
	}
}

type Application struct {
	Drivers        map[string]domain.Driver
	Riders         map[string]domain.Rider
	AvailableRides map[string]domain.Ride
}

func (a *Application) AddDriver(name string) domain.Driver {
	id := a.GenerateID()
	driver := domain.Driver{
		ID:   id,
		Name: name,
	}
	a.Drivers[id] = driver

	return driver
}

func (a *Application) AddRider(name string) domain.Rider {
	id := a.GenerateID()
	rider := domain.Rider{
		ID:   id,
		Name: name,
	}
	a.Riders[id] = rider

	return rider
}

func (a *Application) GenerateID() string {
	return (uuid.New()).String()
}

func (a *Application) DriverOffersRide(
	driverID string,
	Origin int,
	Destination int,
	NumOfSeats int,
) (ride domain.Ride) {
	driver := a.Drivers[driverID]
	if driver.ID != driverID {
		log.Println("Driver not found")
		return
	}

	id := a.GenerateID()

	ride = domain.Ride{
		ID:          id,
		DriverID:    driver.ID,
		Origin:      Origin,
		Destination: Destination,
		NumOfSeats:  NumOfSeats,
	}
	a.AvailableRides[id] = ride
	return
}

func (a *Application) RiderRequestRide(
	riderID string,
	Origin int,
	Destination int,
	NumOfSeats int,
) (ride domain.Ride) {
	rider := a.Riders[riderID]
	if rider.ID != riderID {
		log.Println("Rider not found")
		return
	}

	for id, ride := range a.AvailableRides {
		isMatched := ride.Origin == Origin && ride.Destination == Destination && ride.NumOfSeats >= NumOfSeats
		if isMatched && !ride.IsBooked {
			ride.IsBooked = true
			ride.RiderID = riderID
			a.AvailableRides[id] = ride
			return ride
		}
	}

	log.Println("No available rides")
	return domain.Ride{}
}

func (a *Application) ShowDrivers() {
	fmt.Printf("Drivers: %+v\n", a.Drivers)
}

func (a *Application) ShowRiders() {
	fmt.Printf("Riders: %+v\n", a.Riders)
}

func (a *Application) ShowAvailableRides() {
	fmt.Printf("AvailableRides: %+v\n", a.AvailableRides)
}

func (a *Application) RemoveRide(rideID string) domain.Ride {
	ride := a.AvailableRides[rideID]
	fmt.Printf("%+v %+v\n", rideID, ride)
	delete(a.AvailableRides, rideID)
	return ride
}

func (a *Application) DriverWithdrawRide(rideID string) {
	a.RemoveRide(rideID)
}

func (a *Application) CalculateRideFee(ride domain.Ride) (totalFees float64) {
	constant := float64(0.75)
	AMOUNT_CHARGED_PER_KM := float64(2000)
	distance := float64(ride.Destination - ride.Origin)

	if ride.NumOfSeats >= 2 {
		return distance * float64(ride.NumOfSeats) * constant * AMOUNT_CHARGED_PER_KM
	}

	return distance * AMOUNT_CHARGED_PER_KM
}

func (a *Application) ReleaseRide(rideID string) domain.Ride {
	ride := a.AvailableRides[rideID]
	ride.IsBooked = false
	tempRiderID := ride.RiderID
	ride.RiderID = ""
	a.AvailableRides[rideID] = ride
	ride.RiderID = tempRiderID
	return ride
}

func (a *Application) RiderCloseRide(rideID string) (totalFees float64) {
	ride := a.ReleaseRide(rideID)
	return a.CalculateRideFee(ride)
}
