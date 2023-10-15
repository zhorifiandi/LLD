package domain

import "time"

type SeatAssignment struct {
	UserID string
}

type ShowAssignment struct {
	ID              string
	Capacity        int
	SeatAssignments []SeatAssignment
}

type Theater struct {
	ID         string
	Name       string
	Capacity   int
	Movies     map[string]Movie
	MovieShows map[string]MovieShow
}

type Movie struct {
	ID       string
	Title    string
	Genre    string
	Language string
	Director string
	Duration time.Duration
}

type MovieShow struct {
	ID         string
	Movie      Movie
	TheaterID  string
	StartTime  time.Time
	EndTime    time.Time
	Capacity   int
	Assignment ShowAssignment
}

type User struct {
	ID   string
	Name string
}
