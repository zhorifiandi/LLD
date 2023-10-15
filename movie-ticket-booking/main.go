package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/zhorifiandi/LLD/movieticketbooking/domain"
	"github.com/zhorifiandi/LLD/movieticketbooking/mvpapp"
)

type ISystemApplication interface {
	DisplayAllMovieShows(
		theaterID string,
	)
	FindMovieShowByTitle(
		theaterID string,
		title string,
	) *domain.MovieShow
}

type IAdminApplication interface {
	CreateTheater(theater domain.Theater) error
	AddMovie(theaterID string, movie domain.Movie) error
	AddMovieShow(
		theaterID string,
		movieID string,
		StartTime time.Time,
		EndTime time.Time,
	) error
}

type IUserApplication interface {
	RegisterNewUser(
		theaterID string,
		user domain.User,
	) error
	BookMovieShow(
		theaterID string,
		userID string,
		movieShowID string,
	) error
}

type IApplication interface {
	ISystemApplication
	IAdminApplication
	IUserApplication
}

func main() {
	input := mvpapp.ApplicationInputs{}
	var app IApplication = mvpapp.NewApplication(
		input,
	)
	log.Printf("App is running..... %+v\n", app)

	theaterX := domain.Theater{
		ID:       uuid.New().String(),
		Name:     "X",
		Capacity: 20,
	}
	app.CreateTheater(theaterX)

	antmanMovie := domain.Movie{
		ID:    uuid.New().String(),
		Title: "quantumania",
	}
	app.AddMovie(theaterX.ID, antmanMovie)
	app.AddMovieShow(
		theaterX.ID,
		antmanMovie.ID,
		time.Now().Add(time.Duration(1)*time.Hour),
		time.Now().Add(time.Duration(2)*time.Hour),
	)
	app.DisplayAllMovieShows(theaterX.ID)

}
