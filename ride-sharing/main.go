package main

import (
	"log"

	"github.com/zhorifiandi/LLD/ride-sharing/domain"
	"github.com/zhorifiandi/LLD/ride-sharing/usecase/preferredriderapp"
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
		riderID string,
		Origin int,
		Destination int,
		NumOfSeats int) domain.Ride
	RiderCloseRide(rideID string) float64

	ShowRiders()
	ShowDrivers()
	ShowAvailableRides()
}

func main() {
	input := preferredriderapp.ApplicationInputs{}
	var app IApplication = preferredriderapp.NewApplication(
		input,
	)
	log.Printf("App is running.....\n")

	driver1 := app.AddDriver("Asep")
	app.ShowDrivers()

	rider2 := app.AddRider("Arizho")
	app.ShowRiders()

	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	app.DriverOffersRide(driver1.ID, 50, 60, 2)
	// app.ShowAvailableRides()

	// error case
	app.DriverOffersRide(
		"notfounddriverID",
		50, 60, 2,
	)

	ride1 := app.RiderRequestRide(
		rider2.ID,
		50, 60, 2,
	)
	// app.ShowAvailableRides()

	totalFees := app.RiderCloseRide(ride1.ID)
	log.Printf("Ride ID: %+v, Total Fees: %+v\n", ride1.ID, totalFees)
	// app.ShowAvailableRides()

	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	app.RiderRequestRide(rider2.ID, 50, 60, 2)
	rideX := app.RiderRequestRide(rider2.ID, 50, 60, 2)
	// app.ShowAvailableRides()

	totalFees = app.RiderCloseRide(rideX.ID)
	log.Printf("Ride ID: %+v, Total Fees: %+v\n", rideX.ID, totalFees)
}
