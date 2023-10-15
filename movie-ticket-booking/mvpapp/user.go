package mvpapp

import "github.com/zhorifiandi/LLD/movieticketbooking/domain"

func (a *Application) RegisterNewUser(
	theaterID string,
	user domain.User,
) error {
	return nil
}

func (a *Application) BookMovieShow(
	theaterID string,
	userID string,
	movieShowID string,
) error {
	return nil
}
