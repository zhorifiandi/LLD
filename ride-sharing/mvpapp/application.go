package mvpapp

import (
	"fmt"
	"strconv"
	"time"

	"log"

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
	id := a.generateID()
	driver := domain.Driver{
		ID:   id,
		Name: name,
	}
	a.Drivers[id] = driver

	return driver
}

func (a *Application) AddRider(name string) domain.Rider {
	id := a.generateID()
	rider := domain.Rider{
		ID:   id,
		Name: name,
	}
	a.Riders[id] = rider

	return rider
}

func (a *Application) generateID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
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

	id := a.generateID()

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
	Origin int,
	Destination int,
	NumOfSeats int,
) domain.Ride {
	for id, ride := range a.AvailableRides {
		isMatched := ride.Origin == Origin && ride.Destination == Destination && ride.NumOfSeats >= NumOfSeats
		if isMatched && !ride.IsBooked {
			ride.IsBooked = true
			a.AvailableRides[id] = ride
			return ride
		}
	}

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

func (a *Application) removeRide(rideID string) domain.Ride {
	ride := a.AvailableRides[rideID]
	fmt.Printf("%+v %+v\n", rideID, ride)
	delete(a.AvailableRides, rideID)
	return ride
}

func (a *Application) DriverWithdrawRide(rideID string) {
	a.removeRide(rideID)
}

func (a *Application) RiderCloseRide(rideID string) (totalFees float64) {
	ride := a.removeRide(rideID)
	constant := float64(0.75)
	AMOUNT_CHARGED_PER_KM := float64(2000)
	distance := float64(ride.Destination - ride.Origin)

	if ride.NumOfSeats >= 2 {
		return distance * float64(ride.NumOfSeats) * constant * AMOUNT_CHARGED_PER_KM
	}

	return distance * AMOUNT_CHARGED_PER_KM
}
