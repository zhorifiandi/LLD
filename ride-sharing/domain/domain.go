package domain

type Driver struct {
	ID   string
	Name string
}

type Rider struct {
	ID   string
	Name string
}

type Ride struct {
	ID          string
	DriverID    string
	RiderID     string
	Origin      int
	Destination int
	NumOfSeats  int
	IsBooked    bool
}

type Order struct {
	ID        string
	DriverID  string
	RiderID   string
	TotalFees float64
}
