package preferredriderapp

import (
	"log"

	"github.com/zhorifiandi/LLD/ride-sharing/domain"
	"github.com/zhorifiandi/LLD/ride-sharing/usecase/mvpapp"
)

type BaseApp = mvpapp.Application
type Application struct {
	BaseApp
	OrderHistory []domain.Order
}

type ApplicationInputs struct{}

func NewApplication(inputs ApplicationInputs) (app *Application) {
	baseApp := mvpapp.NewApplication(mvpapp.ApplicationInputs{})
	return &Application{
		BaseApp:      *baseApp,
		OrderHistory: []domain.Order{},
	}
}

func (a *Application) calculateRideFeeForPreferredRider(ride domain.Ride) (totalFees float64) {
	SINGLESEAT_CONSTANT := float64(0.75)
	MULTISEAT_CONSTANT := float64(0.5)
	AMOUNT_CHARGED_PER_KM := float64(2000)
	distance := float64(ride.Destination - ride.Origin)

	if ride.NumOfSeats >= 2 {
		return distance * float64(ride.NumOfSeats) * MULTISEAT_CONSTANT * AMOUNT_CHARGED_PER_KM
	}

	return distance * SINGLESEAT_CONSTANT * AMOUNT_CHARGED_PER_KM
}

func (a *Application) getCompletedRidesByRiderID(riderID string) (count int) {
	for _, ride := range a.OrderHistory {
		if ride.RiderID == riderID {
			count += 1
		}
	}

	// log.Printf("OrderHistory: %+v\n", a.OrderHistory)
	log.Printf("Total Rides: %+v\n", count)
	return
}

func (a *Application) RiderCloseRide(rideID string) (totalFees float64) {
	ride := a.ReleaseRide(rideID)
	log.Printf("ReleaseRide: %+v\n", ride)

	totalRides := a.getCompletedRidesByRiderID(ride.RiderID)
	if totalRides >= 10 {
		return a.calculateRideFeeForPreferredRider(ride)
	}

	return a.CalculateRideFee(ride)
}

func (a *Application) RiderRequestRide(
	riderID string,
	Origin int,
	Destination int,
	NumOfSeats int) (ride domain.Ride) {
	ride = a.BaseApp.RiderRequestRide(
		riderID, Origin, Destination, NumOfSeats,
	)

	if ride.ID == "" {
		return
	}

	a.OrderHistory = append(a.OrderHistory, domain.Order{
		ID:       a.GenerateID(),
		DriverID: ride.DriverID,
		RiderID:  ride.RiderID,
	})

	return
}
