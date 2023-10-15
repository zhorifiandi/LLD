package mvpapp

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/zhorifiandi/LLD/movieticketbooking/domain"
)

func (a *Application) findTheaterByID(theaterID string) (domain.Theater, error) {
	theater, ok := a.Theaters[theaterID]
	if ok {
		return theater, nil
	}

	return theater, errors.New("theater not found")
}

func (a *Application) DisplayAllMovieShows(theaterID string) {
	theater, _ := a.findTheaterByID(theaterID)
	fmt.Printf("==== Start: Theater: %s ====\n", theater.Name)
	for _, movieShow := range theater.MovieShows {
		fmt.Printf("> Title: %+s; (%s/%s)\n - %s until %s",
			movieShow.Movie.Title,
			strconv.Itoa(len(movieShow.Assignment.SeatAssignments)),
			strconv.Itoa(movieShow.Capacity),
			movieShow.StartTime.Format(time.Kitchen),
			movieShow.EndTime.Format(time.Kitchen),
		)
	}

	fmt.Printf("==== End: Theater: %s ====\n", theater.Name)
}

func (a *Application) FindMovieShowByTitle(theaterID string, title string) *domain.MovieShow {
	theater, _ := a.findTheaterByID(theaterID)

	for _, movieShow := range theater.MovieShows {
		if movieShow.Movie.Title == title {
			return &movieShow
		}
	}

	return nil
}
