package main

import (
	"log"

	"github.com/zhorifiandi/LLD/ride-sharing/domain"
	"github.com/zhorifiandi/LLD/ride-sharing/mvpapp"
)

type IApplication interface {
	AddDriver(name string) domain.Driver
	AddRider(name string) domain.Rider
	DriverOffersRide(
		driverID string,
		Origin int,
		Destination int,
		NumOfSeats int) domain.Ride
	DriverWithdrawRide(rideID string)

	RiderRequestRide(
		Origin int,
		Destination int,
		NumOfSeats int) domain.Ride
	RiderCloseRide(rideID string) float64

	ShowRiders()
	ShowDrivers()
	ShowAvailableRides()
}

func main() {
	input := mvpapp.ApplicationInputs{}
	var app IApplication = mvpapp.NewApplication(
		input,
	)
	log.Printf("App is running.....\n")

	driver1 := app.AddDriver("Asep")
	app.ShowDrivers()

	app.DriverOffersRide(
		driver1.ID,
		50, 60, 2,
	)
	app.ShowAvailableRides()

	// error case
	app.DriverOffersRide(
		"notfounddriverID",
		50, 60, 2,
	)

	ride1 := app.RiderRequestRide(
		50, 60, 2,
	)
	app.ShowAvailableRides()

	totalFees := app.RiderCloseRide(ride1.ID)
	log.Printf("Ride ID: %+v, Total Fees: %+v\n", ride1.ID, totalFees)
	app.ShowAvailableRides()

}
