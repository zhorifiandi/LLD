package mvpapp

import (
	"time"

	"github.com/google/uuid"
	"github.com/zhorifiandi/LLD/movieticketbooking/domain"
)

type ApplicationInputs struct{}

func NewApplication(inputs ApplicationInputs) (app *Application) {
	app = &Application{
		Theaters: map[string]domain.Theater{},
	}
	return app
}

type Application struct {
	Theaters map[string]domain.Theater
}

func (a *Application) CreateTheater(theater domain.Theater) error {
	a.Theaters[theater.ID] = theater
	return nil
}

func (a *Application) AddMovie(theaterID string, movie domain.Movie) error {
	a.Theaters[theaterID].Movies[movie.ID] = movie
	return nil
}

func (a *Application) AddMovieShow(
	theaterID string,
	movieID string,
	StartTime time.Time,
	EndTime time.Time,
) error {
	a.Theaters[theaterID].MovieShows[uuid.New().String()] = domain.MovieShow{}
	return nil
}
